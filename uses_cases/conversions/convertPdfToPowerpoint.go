package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func convertPdfToPowerpoint(pdf_api *asposepdfcloud.PdfApiService, pdf_name string, pptx_name string, remote_folder string) {
	outPath := path.Join(remote_folder, pptx_name)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.PutPdfInStorageToPptx(pdf_name, outPath, args)

	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)
	}
}
