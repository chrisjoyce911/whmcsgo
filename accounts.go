package whmcsgo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AccountsService handles communication with the Client related
// methods of the WHMCS API.
//
// WHMCS API docs: https://developers.whmcs.com/api/api-index/
type AccountsService struct {
	client *Client
}

// Account represents an WHMCS user.
type Account struct {
	Address1          string      `json:"address1"`
	Address2          string      `json:"address2"`
	AllowSingleSignOn int         `json:"allowSingleSignOn"`
	Billingcid        int         `json:"billingcid"`
	Cclastfour        interface{} `json:"cclastfour"`
	Cctype            interface{} `json:"cctype"`
	City              string      `json:"city"`
	// Client            struct {
	// 	Address1          string `json:"address1"`
	// 	Address2          string `json:"address2"`
	// 	AllowSingleSignOn int    `json:"allowSingleSignOn"`
	// 	Billingcid        int    `json:"billingcid"`
	// 	City              string `json:"city"`
	// 	Companyname       string `json:"companyname"`
	// 	Country           string `json:"country"`
	// 	Countrycode       string `json:"countrycode"`
	// 	Countryname       string `json:"countryname"`
	// 	Credit            string `json:"credit"`
	// 	Currency          int    `json:"currency"`
	// 	CurrencyCode      string `json:"currency_code"`
	// 	Customfields      []struct {
	// 		ID    int    `json:"id"`
	// 		Value string `json:"value"`
	// 	} `json:"customfields"`
	// 	Customfields1              string `json:"customfields1"`
	// 	Customfields10             string `json:"customfields10"`
	// 	Customfields11             string `json:"customfields11"`
	// 	Customfields12             string `json:"customfields12"`
	// 	Customfields13             string `json:"customfields13"`
	// 	Customfields14             string `json:"customfields14"`
	// 	Customfields15             string `json:"customfields15"`
	// 	Customfields16             string `json:"customfields16"`
	// 	Customfields17             string `json:"customfields17"`
	// 	Customfields2              string `json:"customfields2"`
	// 	Customfields3              string `json:"customfields3"`
	// 	Customfields4              string `json:"customfields4"`
	// 	Customfields5              string `json:"customfields5"`
	// 	Customfields6              string `json:"customfields6"`
	// 	Customfields7              string `json:"customfields7"`
	// 	Customfields8              string `json:"customfields8"`
	// 	Customfields9              string `json:"customfields9"`
	// 	Defaultgateway             string `json:"defaultgateway"`
	// 	Disableautocc              bool   `json:"disableautocc"`
	// 	Email                      string `json:"email"`
	// 	Emailoptout                bool   `json:"emailoptout"`
	// 	Firstname                  string `json:"firstname"`
	// 	Fullname                   string `json:"fullname"`
	// 	Fullstate                  string `json:"fullstate"`
	// 	Groupid                    int    `json:"groupid"`
	// 	ID                         int    `json:"id"`
	// 	IsOptedInToMarketingEmails bool   `json:"isOptedInToMarketingEmails"`
	// 	Language                   string `json:"language"`
	// 	Lastlogin                  string `json:"lastlogin"`
	// 	Lastname                   string `json:"lastname"`
	// 	Latefeeoveride             bool   `json:"latefeeoveride"`
	// 	MarketingEmailsOptIn       bool   `json:"marketing_emails_opt_in"`
	// 	Notes                      string `json:"notes"`
	// 	Overideduenotices          bool   `json:"overideduenotices"`
	// 	Overrideautoclose          bool   `json:"overrideautoclose"`
	// 	Password                   string `json:"password"`
	// 	Phonecc                    int    `json:"phonecc"`
	// 	Phonenumber                string `json:"phonenumber"`
	// 	Phonenumberformatted       string `json:"phonenumberformatted"`
	// 	Postcode                   string `json:"postcode"`
	// 	Securityqans               string `json:"securityqans"`
	// 	Securityqid                int    `json:"securityqid"`
	// 	Separateinvoices           bool   `json:"separateinvoices"`
	// 	State                      string `json:"state"`
	// 	Statecode                  string `json:"statecode"`
	// 	Status                     string `json:"status"`
	// 	TaxID                      string `json:"tax_id"`
	// 	Taxexempt                  bool   `json:"taxexempt"`
	// 	TelephoneNumber            string `json:"telephoneNumber"`
	// 	Twofaenabled               bool   `json:"twofaenabled"`
	// 	Userid                     int    `json:"userid"`
	// 	UUID                       string `json:"uuid"`
	// } `json:"client"`
	Companyname  string `json:"companyname"`
	Country      string `json:"country"`
	Countrycode  string `json:"countrycode"`
	Countryname  string `json:"countryname"`
	Credit       string `json:"credit"`
	Currency     int    `json:"currency"`
	CurrencyCode string `json:"currency_code"`
	// Customfields []struct {
	// 	ID    int    `json:"id"`
	// 	Value string `json:"value"`
	// } `json:"customfields"`
	// Customfields1              string      `json:"customfields1"`
	// Customfields10             string      `json:"customfields10"`
	// Customfields11             string      `json:"customfields11"`
	// Customfields12             string      `json:"customfields12"`
	// Customfields13             string      `json:"customfields13"`
	// Customfields14             string      `json:"customfields14"`
	// Customfields15             string      `json:"customfields15"`
	// Customfields16             string      `json:"customfields16"`
	// Customfields17             string      `json:"customfields17"`
	// Customfields2              string      `json:"customfields2"`
	// Customfields3              string      `json:"customfields3"`
	// Customfields4              string      `json:"customfields4"`
	// Customfields5              string      `json:"customfields5"`
	// Customfields6              string      `json:"customfields6"`
	// Customfields7              string      `json:"customfields7"`
	// Customfields8              string      `json:"customfields8"`
	// Customfields9              string      `json:"customfields9"`
	Defaultgateway             string      `json:"defaultgateway"`
	Disableautocc              bool        `json:"disableautocc"`
	Email                      string      `json:"email"`
	Emailoptout                bool        `json:"emailoptout"`
	Firstname                  string      `json:"firstname"`
	Fullname                   string      `json:"fullname"`
	Fullstate                  string      `json:"fullstate"`
	Gatewayid                  interface{} `json:"gatewayid"`
	Groupid                    int         `json:"groupid"`
	ID                         int         `json:"id"`
	IsOptedInToMarketingEmails bool        `json:"isOptedInToMarketingEmails"`
	Language                   string      `json:"language"`
	Lastlogin                  string      `json:"lastlogin"`
	Lastname                   string      `json:"lastname"`
	Latefeeoveride             bool        `json:"latefeeoveride"`
	MarketingEmailsOptIn       bool        `json:"marketing_emails_opt_in"`
	Notes                      string      `json:"notes"`
	Overideduenotices          bool        `json:"overideduenotices"`
	Overrideautoclose          bool        `json:"overrideautoclose"`
	Password                   string      `json:"password"`
	Phonecc                    int         `json:"phonecc"`
	Phonenumber                string      `json:"phonenumber"`
	Phonenumberformatted       string      `json:"phonenumberformatted"`
	Postcode                   string      `json:"postcode"`
	Result                     string      `json:"result"`
	Securityqans               string      `json:"securityqans"`
	Securityqid                int         `json:"securityqid"`
	Separateinvoices           bool        `json:"separateinvoices"`
	State                      string      `json:"state"`
	Statecode                  string      `json:"statecode"`
	Status                     string      `json:"status"`
	TaxID                      string      `json:"tax_id"`
	Taxexempt                  bool        `json:"taxexempt"`
	TelephoneNumber            string      `json:"telephoneNumber"`
	Twofaenabled               bool        `json:"twofaenabled"`
	Userid                     int         `json:"userid"`
	UUID                       string      `json:"uuid"`
}

func (u Account) String() string {
	return Stringify(u)
}

/*
GetContacts Obtain the Client Contacts that match passed criteria

WHMCS API docs

https://developers.whmcs.com/api-reference/getcontacts/

Request Parameters

limitstart
	The offset for the returned log data (default: 0) Optional
limitnum
	The number of records to return (default: 25) Optional
userid
	Find contacts for a specific client id Optional
firstname
	Find contacts with a specific first name Optional
lastname
	Find contacts with a specific last name Optional
companyname
	Find contacts with a specific company name Optional
email
	Find contacts with a specific email address Optional
address1
	Find contacts with a specific address line 1 Optional
address2
	Find contacts with a specific address line 2 Optional
city
	Find contacts with a specific city Optional
state
	Find contacts with a specific state Optional
postcode
	Find contacts with a specific post/zip code Optional
country
	Find contacts with a specific country Optional
phonenumber
	Find contacts with a specific phone number Optional
subaccount
	Search for sub-accounts Optional
*/
func (s *AccountsService) GetContacts(parms map[string]string) (*Account, *Response, error) {
	a := new(Account)
	resp, err := do(s.client, Params{parms: parms, u: "GetContacts"}, a)
	if err != nil {
		return nil, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &a)

	return a, resp, err
}

// WHMCSclient opbect
type WHMCSclient struct {
	Companyname string    `json:"companyname"`
	Datecreated WHCMSdate `json:"datecreated"`
	Email       string    `json:"email"`
	Firstname   string    `json:"firstname"`
	Groupid     int       `json:"groupid"`
	ID          int       `json:"id"`
	Lastname    string    `json:"lastname"`
	Status      string    `json:"status"`
}

// WHMCSclients object
type WHMCSclients struct {
	Clients struct {
		Client []WHMCSclient `json:"client"`
	} `json:"clients"`
	Numreturned  int    `json:"numreturned"`
	Result       string `json:"result"`
	Startnumber  int    `json:"startnumber"`
	Totalresults int    `json:"totalresults"`
}

// ContactList quick contact list
type ContactList struct {
	CompanyName    string `json:"companyname"`
	FullName       string `json:"fullname"`
	Phone          string `json:"phonenumberformatted"`
	VendorSoftware string `json:"customfields3"`
	VMRef          string `json:"customfields1"`
	AlertPrimary   string `json:"customfields5"`
	Status         string `json:"status"`
	UserID         int    `json:"userid"`
	State          string `json:"state"`
	Email          string `json:"email"`
	GroupID        int    `json:"groupid"`
}
}

/*
ClientContactList list of contact for a given status
*/
func (s *AccountsService) ClientContactList(status string) ([]ContactList, error) {
	params := map[string]string{"sorting": "ASC", "limitstart": "0", "limitnum": "2500"}

	var contactList []ContactList

	//	obj := new(WHMCSclients)
	obj, _, err := s.GetClients(params)

	if err != nil {
		return nil, err
	}

	for _, c := range obj.Clients.Client {

		if c.Status == status {
			cl := ContactList{}
			clientid := fmt.Sprintf("%d", c.ID)
			clParams := map[string]string{"clientid": clientid}

			resp, err := do(s.client, Params{parms: clParams, u: "GetClientsDetails"}, cl)
			if err != nil {
				fmt.Println(err)
			}
			json.Unmarshal([]byte(resp.Body), &cl)

			cl.VendorSoftware = strings.TrimSpace(cl.VendorSoftware)
			cl.Phone = strings.Replace(cl.Phone, "+61.", "0", -1)
			cl.Phone = strings.Replace(cl.Phone, " ", "", -1)
			if len(cl.Phone) == 10 {
				cl.Phone = fmt.Sprintf("%s %s %s", cl.Phone[0:4], cl.Phone[4:7], cl.Phone[7:10])
			}

			cl.AlertPrimary = strings.Replace(cl.AlertPrimary, " ", "", -1)
			if len(cl.AlertPrimary) == 10 {
				cl.AlertPrimary = fmt.Sprintf("%s %s %s", cl.AlertPrimary[0:4], cl.AlertPrimary[4:7], cl.AlertPrimary[7:10])
			}

			contactList = append(contactList, cl)
		}
	}

	return contactList, nil
}

// ClientLastBilledList quick contact list
type ClientLastBilledList struct {
	CompanyName string `json:"companyname"`
	Date        string `json:"date"`
	Total       string `json:"total"`
	Status      string `json:"status"`
}

/*
ClientLastBilled list of the last invoice date
*/
func (s *AccountsService) ClientLastBilled(status string) ([]ClientLastBilledList, error) {
	params := map[string]string{"sorting": "ASC", "limitstart": "0", "limitnum": "2500"}

	var contactList []ClientLastBilledList

	// obj := new(WHMCSclients)
	obj, _, err := s.GetClients(params)

	if err != nil {
		return nil, err
	}

	for _, c := range obj.Clients.Client {

		if c.Status == status {
			cl := ClientLastBilledList{}
			clientid := fmt.Sprintf("%d", c.ID)
			clParams := map[string]string{"clientid": clientid}

			resp, err := do(s.client, Params{parms: clParams, u: "GetClientsDetails"}, cl)
			if err != nil {
				fmt.Println(err)
			}
			json.Unmarshal([]byte(resp.Body), &cl)
			// cl.Phone = strings.Replace(cl.Phone, "+61.", "0", -1)
			// cl.Phone = strings.Replace(cl.Phone, " ", "", -1)
			// if len(cl.Phone) == 10 {
			// 	cl.Phone = fmt.Sprintf("%s %s %s", cl.Phone[0:4], cl.Phone[4:7], cl.Phone[7:10])
			// }

			// cl.AlertPrimary = strings.Replace(cl.AlertPrimary, " ", "", -1)
			// if len(cl.AlertPrimary) == 10 {
			// 	cl.AlertPrimary = fmt.Sprintf("%s %s %s", cl.AlertPrimary[0:4], cl.AlertPrimary[4:7], cl.AlertPrimary[7:10])
			// }

			contactList = append(contactList, cl)
		}
	}

	return contactList, nil
}
