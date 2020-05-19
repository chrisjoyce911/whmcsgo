package whmcsgo

import "encoding/json"

/*
GetClientsProducts Obtain a list of Client Purchased Products matching the provided criteria

WHMCs API docs

https://developers.whmcs.com/api-reference/getclientsproducts/

Request Parameters

limitstart
	int	The offset for the returned log data (default: 0)	Optional
limitnum
	int	The number of records to return (default: 25)	Optional
clientid
	int	The client id to obtain the details for.	Optional
serviceid
	int	The specific service id to obtain the details for	Optional
pid
	int	The specific product id to obtain the details for	Optional
domain
	string	The specific domain to obtain the service details for	Optional
username2
	string	The specific username to obtain the details for	Optional

Response Parameters

result
	string	The result of the operation: success or error
clientid
	int	The specific client id searched for
serviceid
	int	The specific service id searched for
pid
	int	The specific product id searched for
domain
	string	The specific domain searched for
totalresults
	int	The total number of results available
startnumber
	int	The starting number for the returned results
numreturned
	int	The total number of results returned
products
	array	The products returned matching the criteria passed

*/

func (s *AccountsService) GetClientsProducts(parms map[string]string) (*ClientsProduct, *Response, error) {
	p := new(ClientsProduct)
	resp, err := do(s.client, Params{parms: parms, u: "GetClientsProducts"}, p)
	if err != nil {
		return nil, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &p)

	return p, resp, err
}

type ClientsProduct struct {
	Clientid    string      `json:"clientid"`
	Domain      interface{} `json:"domain"`
	Numreturned int64       `json:"numreturned"`
	Pid         interface{} `json:"pid"`
	Products    struct {
		Product []struct {
			Assignedips   string `json:"assignedips"`
			Billingcycle  string `json:"billingcycle"`
			Bwlimit       int64  `json:"bwlimit"`
			Bwusage       int64  `json:"bwusage"`
			Clientid      int64  `json:"clientid"`
			Configoptions struct {
				Configoption []struct {
					ID     int64  `json:"id"`
					Option string `json:"option"`
					Type   string `json:"type"`
					Value  int64  `json:"value"`
				} `json:"configoption"`
			} `json:"configoptions"`
			Customfields struct {
				Customfield []struct {
					ID             int64  `json:"id"`
					Name           string `json:"name"`
					TranslatedName string `json:"translated_name"`
					Value          string `json:"value"`
				} `json:"customfield"`
			} `json:"customfields"`
			Dedicatedip         string      `json:"dedicatedip"`
			Disklimit           int64       `json:"disklimit"`
			Diskusage           int64       `json:"diskusage"`
			Domain              string      `json:"domain"`
			Firstpaymentamount  string      `json:"firstpaymentamount"`
			Groupname           string      `json:"groupname"`
			ID                  int64       `json:"id"`
			Lastupdate          string      `json:"lastupdate"`
			Name                string      `json:"name"`
			Nextduedate         string      `json:"nextduedate"`
			Notes               string      `json:"notes"`
			Ns1                 string      `json:"ns1"`
			Ns2                 string      `json:"ns2"`
			Orderid             int64       `json:"orderid"`
			Overideautosuspend  int64       `json:"overideautosuspend"`
			Overidesuspenduntil string      `json:"overidesuspenduntil"`
			Password            string      `json:"password"`
			Paymentmethod       string      `json:"paymentmethod"`
			Paymentmethodname   string      `json:"paymentmethodname"`
			Pid                 int64       `json:"pid"`
			Promoid             int64       `json:"promoid"`
			Recurringamount     string      `json:"recurringamount"`
			Regdate             string      `json:"regdate"`
			Serverhostname      interface{} `json:"serverhostname"`
			Serverid            int64       `json:"serverid"`
			Serverip            interface{} `json:"serverip"`
			Servername          string      `json:"servername"`
			Status              string      `json:"status"`
			Subscriptionid      string      `json:"subscriptionid"`
			Suspensionreason    string      `json:"suspensionreason"`
			TranslatedGroupname string      `json:"translated_groupname"`
			TranslatedName      string      `json:"translated_name"`
			Username            string      `json:"username"`
		} `json:"product"`
	} `json:"products"`
	Result       string      `json:"result"`
	Serviceid    interface{} `json:"serviceid"`
	Startnumber  int64       `json:"startnumber"`
	Totalresults int64       `json:"totalresults"`
}
