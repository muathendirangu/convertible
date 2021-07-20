package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)
var addConversionChoices= []struct{
	body []byte
	wantedCode int
}{
  {
   
    body: []byte(`{"from":"NGN","to":"KES","amount":200}`),
    wantedCode: http.StatusBadRequest,
  },
  {
    body: []byte("{}"),
    wantedCode: http.StatusBadRequest,
  },
  {
	body: []byte(`{"amount": "5","from": "KES","to": "GHC"}`),
	wantedCode: http.StatusOK,
  },
}
func TestCurrencyConverterHandler(t *testing.T)  {
	for _, testCase := range addConversionChoices{
		recorder := httptest.NewRecorder()
		request,_:=http.NewRequest(http.MethodPost, "/",bytes.NewBuffer(testCase.body))
		currencyConverterHandler(recorder,request)
		if recorder.Code != testCase.wantedCode {
			t.Errorf("got %d wanted %d", recorder.Code, testCase.wantedCode)
		}		
	}
}