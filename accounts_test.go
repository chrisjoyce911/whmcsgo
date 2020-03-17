package whmcsgo

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestAccountsService_GetClients(t *testing.T) {

	type fields struct {
		StatusCode int
		Body       string
	}
	type args struct {
		parms map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WHMCSclients
		want1   *Response
		wantErr bool
	}{
		{
			name: "First test",
			fields: fields{
				StatusCode: 200,
				Body:       `{"result":"success","totalresults":683,"startnumber":1,"numreturned":2,"clients":{"client":[{"id":638,"firstname":"Darren","lastname":"","companyname":"360  Financial Group","email":"darren@360fg.com.au","datecreated":"2017-10-26","groupid":2,"status":"Inactive"},{"id":632,"firstname":"Fadi","lastname":"","companyname":"Eureka Accounting Group","email":"fadi.said@eurekaaccountinggroup.com.au","datecreated":"2017-10-12","groupid":2,"status":"Inactive"}]}}`,
			},
			args: args{parms: map[string]string{"sorting": "ASC", "limitstart": "0", "limitnum": "2500"}},
			want: &WHMCSclients{},
			want1: &Response{
				StatusCode:    200,
				Body:          `{"result":"success","totalresults":683,"startnumber":1,"numreturned":2,"clients":{"client":[{"id":638,"firstname":"Darren","lastname":"","companyname":"360  Financial Group","email":"darren@360fg.com.au","datecreated":"2017-10-26","groupid":2,"status":"Inactive"},{"id":632,"firstname":"Fadi","lastname":"","companyname":"Eureka Accounting Group","email":"fadi.said@eurekaaccountinggroup.com.au","datecreated":"2017-10-12","groupid":2,"status":"Inactive"}]}}`,
				ContentLength: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tclient := NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: tt.fields.StatusCode,
					Body:       ioutil.NopCloser(bytes.NewBufferString(tt.fields.Body)),
					Header:     make(http.Header),
				}
			})

			s := &AccountsService{
				client: NewClient(tclient, Authentication{}, "defaultBaseURL string"),
			}
			got, got1, err := s.GetClients(tt.args.parms)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountsService.GetClients() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountsService.GetClients() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AccountsService.GetClients() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
