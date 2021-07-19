package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)
var addConversionChoices= []struct{
	body []byte
	wantedCode int
}{
  {
   
    body: []byte(`{"NGN","KES",200}`),
    wantedCode: http.StatusOK,
  },
  {
    body: []byte("{}"),
    wantedCode: http.StatusBadRequest,
  },
}
func TestConvertCurrency(t *testing.T)  {
	for _, testCase := range addConversionChoices{
		recorder := httptest.NewRecorder()
		request,_:=http.NewRequest(http.MethodPost, "/",bytes.NewBuffer(testCase.body))
		currencyConverterHandler(recorder,request)
		log.Fatal(recorder.Code)
		if recorder.Code != testCase.wantedCode {
			t.Errorf("got %d wanted %d", recorder.Code, testCase.wantedCode)
		}		
	}
}