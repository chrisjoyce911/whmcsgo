package whmcsgo

import (
	"encoding/json"
	"fmt"
)

/*
UpdateInvoice Update an invoice using the provided parameters.

https://developers.whmcs.com/api-reference/updateinvoice/

Request Parameters

invoiceid
	int	The ID of the invoice to update	Required
status
	string	The status of the invoice being	Optional
paymentmethod
	string	The payment method of the invoice in system format	Optional
taxrate
	float	The first level tax rate to apply to the invoice to override the system default	Optional
taxrate2
	float	The second level tax rate to apply to the invoice to override the system default	Optional
credit
	float	Update the credit applied to the invoice	Optional
date
	\Carbon\Carbon	The date that the invoice should show as created YYYY-mm-dd	Optional
duedate
	\Carbon\Carbon	The due date of the invoice YYYY-mm-dd	Optional
datepaid
	\Carbon\Carbon	The date paid of the invoice YYYY-mm-dd	Optional
notes
	string	The notes to appear on the invoice	Optional
itemdescription
	string[]	An array of lineItemId => Description of items to change. The lineItemId is the id of the item from the GetInvoice API command.	Optional
itemamount
	float[]	An array of lineItemId => amount of items to change. Required if itemdescription is provided.	Optional
itemtaxed
	bool[]	An array of lineItemId => taxed of items to change Required if itemdescription is provided.	Optional
newitemdescription
	string[]	The line items description. This should be a numerically indexed array of new line item descriptions.	Optional
newitemamount
	float[]	The line items amount. This should be a numerically indexed array of new line item amounts.	Optional
newitemtaxed
	bool[]	Should the new line items be taxed. This should be a numerically indexed array of new line item taxed values.	Optional
deletelineids
	int[]	An array of line item ids to remove from the invoice. This is the id of the line item, from GetInvoice API command.	Optional
publish
	bool	Publish the invoice	Optional
publishandsendemail
	bool	Publish and email the invoice	Optional

Response Parameters

result
	string	The result of the operation: success or error
invoiceid
	int	The ID of the invoice

*/
func (s *BillingService) UpdateInvoice(invoiceID int, items []InvoiceLineItems) (*InvoiceReply, *Response, error) {

	i := new(InvoiceReply)
	parms := map[string]string{}

	for _, li := range items {
		parms[fmt.Sprintf("newitemdescription[%d]", li.ItemOrder)] = li.ItemDescription
		parms[fmt.Sprintf("newitemamount[%d]", li.ItemOrder)] = fmt.Sprintf("%.2f", li.ItemAmount)
		parms[fmt.Sprintf("newitemtaxed[%d]", li.ItemOrder)] = FormatBool(li.ItemTaxed)
	}

	parms["invoiceid"] = fmt.Sprintf("%d", invoiceID)

	resp, err := do(s.client, Params{parms: parms, u: "UpdateInvoice"}, i)
	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &i)
	return i, resp, err
}
