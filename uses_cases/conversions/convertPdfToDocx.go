package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func convertPdfToWord(pdf_api *asposepdfcloud.PdfApiService, pdf_name string, doc_name string, remote_folder string) {
	outPath := path.Join(remote_folder, doc_name)

	args := map[string]interface{}{
		"folder": remote_folder,
		"format": string(asposepdfcloud.DocFormatDocX),
	}

	result, httpResponse, err := pdf_api.PutPdfInStorageToDoc(pdf_name, outPath, args)

	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)
	}
}
