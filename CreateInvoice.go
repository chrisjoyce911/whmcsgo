package whmcsgo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/*
CreateInvoice Create an invoice using the provided parameters.

WHMCS API docs

https://developers.whmcs.com/api-reference/getclients/

Request Parameters

draft
	bool	Should the invoice be created in draft status (No need to pass $status also)	Optional
paymentmethod
	string	The payment method of the created invoice in system format	Optional
taxrate
	float	The first level tax rate to apply to the invoice to override the system default	Optional
taxrate2
	float	The second level tax rate to apply to the invoice to override the system default	Optional

Response Parameters

result
	string	The result of the operation: success or error
invoiceid
	int	The ID of the newly created invoice
status
	string	The status of the newly created invoice

*/
func (s *BillingService) CreateInvoice(userID int, invoice CreateInvoiceRequest) (int, *Response, error) {
	//a := new(BillingItem)

	userid := fmt.Sprintf("%d", userID)
	invoiceParams := map[string]string{"userid": userid}

	switch invoice.Status {
	case "Draft", "Unpaid", "Paid":
		invoiceParams["status"] = invoice.Status
	default:
		return 0, &Response{}, fmt.Errorf("unsupported status value: %s", invoice.Status)
	}

	invoiceParams["sendinvoice"] = FormatBool(invoice.SendInvoice)
	invoiceParams["autoapplycredit"] = FormatBool(invoice.AutoApplyCredit)

	layout := "2006-01-02"

	if !invoice.Date.IsZero() {
		invoiceParams["date"] = invoice.Date.Format(layout)
	}

	if !invoice.DueDate.IsZero() {
		invoiceParams["duedate"] = invoice.DueDate.Format(layout)
	}

	if len(invoice.Notes) > 0 {
		invoiceParams["notes"] = invoice.Notes
	}

	if len(invoice.LineItems) < 1 {
		return 0, &Response{}, fmt.Errorf("No Line items for invoice found")
	}

	li := lineItemstoParams(invoice.LineItems)
	for k, v := range li {
		invoiceParams[k] = v
	}

	resp, err := do(s.client, Params{parms: invoiceParams, u: "CreateInvoice"}, nil)
	if err != nil {
		return 0, resp, err
	}

	// WHMCS returns a error sometimes that is not in JSON !
	r := strings.Replace(resp.Body, `<div class="alert alert-error">Module credit_purchase_improvement: Module error occured. Please contact with support.</div>`, ``, -1)
	ir := InvoiceReply{}
	err = json.Unmarshal([]byte(r), &ir)

	return ir.InvoiceID, resp, err

}

func lineItemstoParams(items []InvoiceLineItems) map[string]string {
	lineItems := map[string]string{}
	for _, li := range items {
		lineItems[fmt.Sprintf("itemdescription%d", li.ItemOrder)] = li.ItemDescription
		lineItems[fmt.Sprintf("itemamount%d", li.ItemOrder)] = fmt.Sprintf("%.2f", li.ItemAmount)
		lineItems[fmt.Sprintf("itemtaxed%d", li.ItemOrder)] = FormatBool(li.ItemTaxed)
	}

	return lineItems
}

// CreateInvoiceRequest the new invoice to be created for a client
type CreateInvoiceRequest struct {
	// The ID of the client to charge
	Status          string             // The status of the invoice being created Paid,Unpaid,Draft
	SendInvoice     bool               // Should the Invoice Created Email be sent to the client
	Date            time.Time          // The date that the invoice should show as created
	DueDate         time.Time          // The due date of the newly created invoice
	Notes           string             // The notes to appear on the created invoice
	AutoApplyCredit bool               // Should credit on the client account be automatically applied to the invoice
	LineItems       []InvoiceLineItems // Invoice Line Items
}

// InvoiceLineItems the new invoice to be created for a client
type InvoiceLineItems struct {
	ItemOrder       int     // The Order to be show on the invoice
	ItemDescription string  // The line items description
	ItemAmount      float32 // The line items amount
	ItemTaxed       bool    // The line items is taxed value

}

// InvoiceReply the status after creating or updating an invoice
type InvoiceReply struct {
	InvoiceID int    `json:"invoiceid"` // The ID of the invoice
	Result    string `json:"result"`    // The result of the operation: success or error
	Status    string `json:"status"`    // The status of the invoice
}
