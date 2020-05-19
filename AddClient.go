package whmcsgo

/*
AddClient Adds a client

WHMCs API docs

https://developers.whmcs.com/api-reference/addclient/

*/
func (s *AccountsService) AddClient(parms map[string]string) (*Account, *Response, error) {
	a := new(Account)
	resp, err := do(s.client, Params{parms: parms, u: "AddClient"}, a)
	if err != nil {
		return nil, resp, err
	}
	return a, resp, err
}
