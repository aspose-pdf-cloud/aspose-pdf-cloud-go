package asposepdfcloud

import (
	"fmt"
	"net/http"
	"testing"
)

var (
	remoteFolder = "TempPdfCloud"
)
func TestPutCreateDocument(t *testing.T) {
	var (
		response DocumentResponse
		httpResponse *http.Response
		err error
	)
	
	args := map[string]interface{} {
		"folder": remoteFolder,
	}

	pdfAPI := NewPdfApiService("AppSid", "AppKey", "https://billing.cloud.saltov.dynabic.com/v2.0")
	

	response, httpResponse, err =  pdfAPI.PutCreateDocument("pdf_go.pdf", args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		t.Log(fmt.Sprint(response.Code))
	}
}