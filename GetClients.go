package whmcsgo

import "encoding/json"

/*
GetClients Obtain the Clients that match passed criteria

WHMCS API docs

https://developers.whmcs.com/api-reference/getclients/

Request Parameters

limitstart
	The offset for the returned log data (default: 0) Optional
limitnum
	The number of records to return (default: 25) Optional
sorting
	The direction to sort the results. ASC or DESC. Default: ASC Optional
search
	The search term to look for at the start of email, firstname, lastname, fullname or companyname Optional
*/
func (s *AccountsService) GetClients(parms map[string]string) (*WHMCSclients, *Response, error) {
	obj := new(WHMCSclients)
	resp, err := do(s.client, Params{parms: parms, u: "GetClients"}, obj)
	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
