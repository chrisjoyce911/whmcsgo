package whmcsgo

import (
	"encoding/json"
	"errors"
	"fmt"
)

// BillingService handles communication with the billableitem related
// methods of the WHMCS API.
//
// WHMCS API docs: http://docs.whmcs.com/API
type BillingService struct {
	client *Client
}

// BillingItem represents a
type BillingItem struct {
	ClientID      *string `json:"clientid"`
	Description   *string `json:"description"`
	Hours         *string `json:"hours"`
	Amount        *string `json:"amount"`
	InvoiceAction *string `json:"invoiceaction"`
}

func (r BillingItem) String() string {
	return Stringify(r)
}

// AddBillableItem a new billable item.
//
// WHMCS API docs: https://developers.whmcs.com/api-reference/addbillableitem/
func (s *BillingService) AddBillableItem(parms map[string]string) (*BillingItem, *Response, error) {
	a := new(BillingItem)
	resp, err := do(s.client, Params{parms: parms, u: "AddBillableItem"}, a)
	if err != nil {
		return nil, resp, err
	}
	return a, resp, err
}

// InvoicesReply object from WHMCS
type InvoicesReply struct {
	Invoices struct {
		Invoice []Invoice `json:"invoice"`
	} `json:"invoices"`
	Numreturned  int    `json:"numreturned"`
	Result       string `json:"result"`
	Startnumber  int    `json:"startnumber"`
	Totalresults int    `json:"totalresults"`
}

// Invoices from WHCMS
type Invoices struct {
	Invoice []Invoice `json:"invoice"`
}

// Invoice from WHCMS
type Invoice struct {
	Companyname        string      `json:"companyname"`
	CreatedAt          WHCMSdate   `json:"created_at"`
	Credit             string      `json:"credit"`
	Currencycode       string      `json:"currencycode"`
	Currencyprefix     string      `json:"currencyprefix"`
	Currencysuffix     string      `json:"currencysuffix"`
	Date               WHCMSdate   `json:"date"`
	DateCancelled      WHCMSdate   `json:"date_cancelled"`
	DateRefunded       WHCMSdate   `json:"date_refunded"`
	Datepaid           WHCMSdate   `json:"datepaid"`
	Duedate            WHCMSdate   `json:"duedate"`
	Firstname          string      `json:"firstname"`
	ID                 int         `json:"id"`
	Invoicenum         string      `json:"invoicenum"`
	LastCaptureAttempt WHCMSdate   `json:"last_capture_attempt"`
	Lastname           string      `json:"lastname"`
	Notes              string      `json:"notes"`
	Paymentmethod      string      `json:"paymentmethod"`
	Paymethodid        interface{} `json:"paymethodid"`
	Status             string      `json:"status"`
	Subtotal           string      `json:"subtotal"`
	Tax                string      `json:"tax"`
	Tax2               string      `json:"tax2"`
	Taxrate            string      `json:"taxrate"`
	Taxrate2           string      `json:"taxrate2"`
	Total              string      `json:"total"`
	UpdatedAt          WHCMSdate   `json:"updated_at"`
	Userid             int         `json:"userid"`
}

func (i Invoices) String() string {
	return Stringify(i)
}

/*
GetLastInvoice retrieve the task invoice for a given status

Request Parameters

userid
	Find invoices for a specific client id - Requited

status
	Find invoices for a specific status. Standard Invoice statuses plus Overdue - Requited

*/
func (s *BillingService) GetLastInvoice(userid int, status string) (Invoice, error) {
	invoices := new(InvoicesReply)
	uid := fmt.Sprintf("%d", userid)

	parms := map[string]string{"status": status, "userid": uid, "limitnum": "1", "orderby": "date", "order": "desc"}

	resp, err := do(s.client, Params{u: "GetInvoices", parms: parms}, invoices)

	if err != nil {
		return Invoice{}, err
	}

	json.Unmarshal([]byte(resp.Body), &invoices)

	if invoices.Numreturned > 0 {
		return invoices.Invoices.Invoice[0], nil
	}
	return Invoice{}, errors.New("No invoice found")

	// return invoices.Invoices.Invoice[0], err
}

/*
GetInvoices retrieve a list of invoices.

WHMCS API docs

https://developers.whmcs.com/api-reference/getinvoices/

Request Parameters

limitstart
	The offset for the returned invoice data (default: 0) - Optional

limitnum
	The number of records to return (default: 25) - Optional

userid
	Find invoices for a specific client id - Optional

status
	Find invoices for a specific status. Standard Invoice statuses plus Overdue - Optional

orderby
	The field to sort results by. Accepted values are: id, invoicenumber, date, duedate, total, status - Optional

order
	Order sort attribute. Accepted values are: asc or desc - Optional
*/
func (s *BillingService) GetInvoices(parms map[string]string) ([]Invoice, *Response, error) {
	invoices := new(InvoicesReply)
	resp, err := do(s.client, Params{parms: parms, u: "GetInvoices"}, invoices)

	if err != nil {
		return nil, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &invoices)

	var r []Invoice
	for _, i := range invoices.Invoices.Invoice {
		r = append(r, i)
	}

	return r, resp, err
}

// CaptureResult from a payment capture attempt
// Possible error condition responses include:
// Invoice Not Found or Not Unpaid
// Payment Attempt Failed
type CaptureResult struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

/*
CapturePayment attempt to capture a payment on an unpaid Invoice

WHMCS API docs: https://developers.whmcs.com/api-reference/capturepayment/

Request Parameters

invoiceid
	The ID of the pending order	Required
cvv
	The CVV Number for the card being attempted	Optional
*/
func (s *BillingService) CapturePayment(invoice int) (*CaptureResult, *Response, error) {
	if invoice < 1 {
		return nil, nil, errors.New("Invoice ID required to attempt payment capture")
	}
	result := new(CaptureResult)

	invoiceid := fmt.Sprintf("%d", invoice)

	parms := map[string]string{"invoiceid": invoiceid}

	resp, err := do(s.client, Params{parms: parms, u: "CapturePayment"}, result)

	if err != nil {
		return nil, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &result)
	return result, resp, err
}
