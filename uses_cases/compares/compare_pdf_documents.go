package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

const (
	REMOTE_FOLDER  = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER   = "c:\\Samples"
	PDF_DOCUMENT_1 = "sample_compare_1.pdf"
	PDF_DOCUMENT_2 = "sample_compare_2.pdf"
	PDF_OUTPUT     = "output_compare.pdf"

	PDF_API_SID = "******"
	PDF_API_KEY = "******"
)

func СomparePdf(pdf_api *asposepdfcloud.PdfApiService, document1 string, document2 string, output_document string, remote_folder string) {
	uploadFile(pdf_api, document1)
	uploadFile(pdf_api, document2)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.PostComparePdf(path.Join(REMOTE_FOLDER, document1), path.Join(REMOTE_FOLDER, document2), path.Join(REMOTE_FOLDER, output_document), args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)
		downloadFile(pdf_api, PDF_OUTPUT, PDF_OUTPUT)
	}
}
