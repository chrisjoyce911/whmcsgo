package whmcsgo

/*
UpdateClient Updates a client with the passed parameters.

WHMCS API docs

https://developers.whmcs.com/api-reference/updateclient/
*/
func (s *AccountsService) UpdateClient() (*Account, *Response, error) {
	a := new(Account)
	var parms map[string]string

	resp, err := do(s.client, Params{parms: parms, u: "updateclient"}, a)
	if err != nil {
		return nil, resp, err
	}
	return a, resp, err
}
