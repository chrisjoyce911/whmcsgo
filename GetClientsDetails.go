package whmcsgo

import (
	"encoding/json"
)

/*
GetClientsDetails Obtain the Clients Details for a specific client

WARNING

Please use GetContacts , GetClientsDetails may be deprecated and may be removed in a future version of WHMCS.

Note this function returns the client information in the top level array.

WHMCS API docs

https://developers.whmcs.com/api-reference/getclientsdetails/

Request Parameters

to many to list see WHMCS API docs
*/
func (s *AccountsService) GetClientsDetails(parms map[string]string) (*Account, *Response, error) {
	a := new(Account)
	resp, err := do(s.client, Params{parms: parms, u: "GetClientsDetails"}, a)
	if err != nil {
		return nil, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &a)

	return a, resp, err
}
