package whmcsgo

import "encoding/json"

/*
UpdateClientProduct Updates a Client Service

WHMCS API docs

https://developers.whmcs.com/api-reference/updateclientproduct/

Request Parameters

serviceid
	int	The id of the client service to update	Required
pid
	int	The package id to associate with the service	Optional
serverid
	int	The server id to associate with the service	Optional
regdate
	\Carbon\Carbon	The registration date of the service (Y-m-d)	Optional
nextduedate
	\Carbon\Carbon	The next due date of the service (Y-m-d)	Optional
terminationDate
	\Carbon\Carbon	Update the termination date of the service (Y-m-d)	Optional
completedDate
	\Carbon\Carbon	Update the completed date of the service (Y-m-d)	Optional
domain
	string	The domain name to be changed to	Optional
firstpaymentamount
	float	The first payment amount on the service	Optional
recurringamount
	float	The recurring amount for automatic renewal invoices	Optional
paymentmethod
	string	The payment method to associate in system format (eg paypal)	Optional
billingcycle
	string	The term in which the product is billed on (eg One-Time, Monthly, Quarterly, etc)	Optional
subscriptionid
	string	The subscription ID to associate with the service	Optional
status
	string	The status to change the service to	Optional
notes
	string	The admin notes for the service	Optional
serviceusername
	string	The service username	Optional
servicepassword
	string	The service password	Optional
overideautosuspend
	string	Should override auto suspend be provided (‘on’ or ‘off’)	Optional
overidesuspenduntil
	\Carbon\Carbon	Update the Override Suspend date of the service (Y-m-d)	Optional
ns1
	string	(VPS/Dedicated servers only)	Optional
ns2
	string	(VPS/Dedicated servers only)	Optional
dedicatedip
	string		Optional
assignedips
	string	(VPS/Dedicated servers only)	Optional
diskusage
	int	The disk usage in megabytes	Optional
disklimit
	int	The disk limit in megabytes	Optional
bwusage
	int	The bandwidth usage in megabytes	Optional
bwlimit
	int	The bandwidth limit in megabytes	Optional
overidesuspenduntil
	\Carbon\Carbon		Optional
suspendreason
	string		Optional
promoid
	int	The promotion Id to associate	Optional
unset
	array	An array of items to unset. Can be one of: ‘domain’, ‘serviceusername’, ‘servicepassword’, ‘subscriptionid’, ‘ns1’, ‘ns2’, ‘dedicatedip’, ‘assignedips’, ‘notes’, ‘suspendreason’	Optional
autorecalc
	bool	Should the recurring amount of the service be automatically recalculated (this will ignore any passed $recurringamount)	Optional
customfields
	string	Base64 encoded serialized array of custom field values - base64_encode(serialize(array(“1”=>“Yahoo”)));	Optional
configoptions
	string	Base64 encoded serialized array of configurable option field values - base64_encode(serialize(array(configoptionid => dropdownoptionid, XXX => array(‘optionid’ => YYY, ‘qty’ => ZZZ)))) - XXX is the ID of the configurable option - YYY is the optionid found in tblhostingconfigoption.optionid - ZZZ is the quantity you want to use for that option	Optional

Response Parameters

result
	string	The result of the operation: success or error
serviceid
	int	The Id of the updated service
*/
func (s *SystemService) UpdateClientProduct(parms map[string]string) (*UpdateClientProductReply, *Response, error) {
	obj := new(UpdateClientProductReply)
	resp, err := do(s.client, Params{parms: parms, u: "UpdateClientProduct"}, obj)
	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}

type UpdateClientProductReply struct {
	Result    string `json:"result"`
	Serviceid string `json:"serviceid"`
}
