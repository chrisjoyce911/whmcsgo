package whmcsgo

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// A Client manages communication with the WHMCS API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.  Defaults to https://www.yourdomain.com/billing/, but can be
	// set to a domain endpoint to use with your billing at your enterprise.  BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// APIEndSux endpoint suffix used when communicating with the WHMCS API.
	APIEndSux string

	Auth Authentication
	// Services used for talking to different parts of the WHMCS API.
	Orders     *OrdersService
	Billing    *BillingService
	Module     *ModuleService
	Support    *SupportService
	System     *SystemService
	Products   *ProductsService
	Project    *ProjectService
	Affiliates *AffiliatesService
	Accounts   *AccountsService // Client

	Domains *DomainsService
	Servers *ServersService
	Tickets *TicketsService
	Service *ServiceService
	Addons  *AddonsService
}

// NewClient returns a new WHMCS API client.  If a nil httpClient is
// provided, http.DefaultClient will be used.  To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, authentication Authentication, defaultBaseURL string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, APIEndSux: "includes/api.php", Auth: authentication}

	c.Orders = &OrdersService{client: c}
	c.Billing = &BillingService{client: c}
	c.Module = &ModuleService{client: c}
	c.Support = &SupportService{client: c}
	c.System = &SystemService{client: c}
	c.Products = &ProductsService{client: c}
	c.Project = &ProjectService{client: c}
	c.Affiliates = &AffiliatesService{client: c}
	c.Accounts = &AccountsService{client: c}

	c.Domains = &DomainsService{client: c}
	c.Servers = &ServersService{client: c}
	c.Tickets = &TicketsService{client: c}
	c.Service = &ServiceService{client: c}
	c.Addons = &AddonsService{client: c}
	return c
}

// Authentication provides authentication information for the client
type Authentication struct {
	authentication url.Values
}

// NewAuth provides authentication information for the client
//
// map[string]string{"identifier": "xxxxx", "secret": "xxxxx", "accesskey": "access key"}
//
// map[string]string{"identifier": "xxxxx", "secret": "xxxxx"}
//
// map[string]string{"username": "yyyy", "password": "yyyy"}
//
// As well, the associated admin user must have the API Access permission granted to their admin role group.
// API Authentication Credentials can be generated for an admin user within the Admin area as described in the WHMCS Documentation.
// Authentication is required for each API request.
//
// Authenticating With API Credentials (http://docs.whmcs.com/API_Authentication_Credentials#Creating_Admin_API_Authentication_Credentials)
// API requests will be authenticated based on the request parameters identifier and secret as provisioned when Creating Admin API Authentication Credentials within the WHMCS Admin Area.
//
// Authenticating With Login Credentials
// Prior to WHMCS verison 7.2, authentication was validated based on admin login credentials, and not API Authentication Credentials.
// This method of authentication is still supported for backwards compatibility but may be deprecated in a future version of WHMCS.
//
// To authenticate with the admin login credentials, pass the admin username and the MD5 hashed value of the respective admin’s password.
//
// Access to the API by default is restricted by IP. (WHMCS admin area and navigate to Setup > General Settings > Security.)
//For situations where IP access control is not feasible, an Access Key can also be configured.
func NewAuth(authInfo map[string]string) Authentication {

	var auth Authentication
	a := url.Values{}

	if len(authInfo["accesskey"]) > 0 {
		a.Add("accesskey", authInfo["accesskey"])
	}

	// have user/pass
	if len(authInfo["identifier"]) > 0 && len(authInfo["secret"]) > 0 {
		a.Add("identifier", authInfo["identifier"])
		a.Add("secret", authInfo["secret"])
	}

	if len(authInfo["username"]) > 0 && len(authInfo["password"]) > 0 {
		a.Add("username", authInfo["username"])
		a.Add("password", authInfo["password"])
	}

	auth.authentication = a
	return auth
}

// A WRequest manages communication with the WHMCS API.
type WRequest struct {
	data *url.Values

	url *url.URL
}

// Params specifies the optional parameters to various List methods that
// support pagination.
type Params struct {
	parms map[string]string
	u     string
}

// addFormValues adds the parameters in opt as URL values parameters.
func addFormValues(opt map[string]string) *url.Values {
	uv := url.Values{}
	for k, v := range opt {
		uv.Set(k, v)
	}
	return &uv
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(dat map[string]string, action string) (*WRequest, error) {
	rel, err := url.Parse(c.APIEndSux)
	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)

	if len(strings.TrimSpace(action)) > 0 {
		dat["action"] = action
		dat["responsetype"] = "json"
	}

	for key, value := range c.Auth.authentication {
		dat[key] = value[0]
	}

	return &WRequest{url: u, data: addFormValues(dat)}, nil
}

// Response is a WHMCS API response.  This wraps the standard http.Response
// returned from WHMCS and provides convenient access to things like
// pagination links.
type Response struct {
	Status        string // e.g. "200 OK"
	StatusCode    int    // e.g. 200
	Body          string
	ContentLength int64
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	body, _ := ioutil.ReadAll(r.Body)
	response := &Response{
		Status:        r.Status,
		StatusCode:    r.StatusCode,
		Body:          string(body),
		ContentLength: r.ContentLength,
	}
	return response
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req WRequest, v interface{}) (*Response, error) {

	// 	fmt.Println("--- " + req.url.String())
	resp, err := c.client.PostForm(req.url.String(), *req.data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)
	err = CheckResponse(response)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	return response, err
}

func do(c *Client, p Params, a interface{}) (*Response, error) {
	req, err := c.NewRequest(p.parms, p.u)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(*req, a)
	if err != nil {
		return resp, err
	}

	return resp, err
}

/*
ErrorResponse reports one or more errors caused by an API request.
*/
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
	Errors   []Error        `json:"errors"`  // more detail on individual errors
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message, r.Errors)
}

// sanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user, specifically in the ErrorResponse error message.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("accesskey")) > 0 {
		params.Set("accesskey", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

/*
Error reports more details on an individual error in an ErrorResponse.
 These are the possible validation error codes:

	 missing:
		 resource does not exist
	 missing_field:
		 a required field on a resource has not been set
	 invalid:
		 the formatting of a field is invalid
	 already_exists:
		 another resource has the same valid as this field
*/
type Error struct {
	Resource string `json:"resource"` // resource on which the error occurred
	Field    string `json:"field"`    // field on which the error occurred
	Code     string `json:"code"`     // validation error code
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused by %v field on %v resource",
		e.Code, e.Field, e.Resource)
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func CheckResponse(r *Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return errors.New(r.Body)
}

// parseBoolResponse determines the boolean result from a WHMCS API response.
// Several WHMCS API methods return boolean responses indicated by the HTTP
// status code in the response (true indicated by a 204, false indicated by a
// 404).  This helper function will determine that result and hide the 404
// error if present.  Any other error will be returned through as-is.
func parseBoolResponse(err error) (bool, error) {
	if err == nil {
		return true, nil
	}

	if err, ok := err.(*ErrorResponse); ok && err.Response.StatusCode == http.StatusNotFound {
		// Simply false.  In this one case, we do not pass the error through.
		return false, nil
	}

	// some other real error occurred
	return false, err
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.
func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	p := new(string)
	*p = v
	return p
}

// FormatBool returns "1" or "0" according to the value of b
func FormatBool(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
