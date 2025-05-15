package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func convertMarkdownToPdf(pdf_api *asposepdfcloud.PdfApiService, md_document string, pdf_name string, remote_folder string) {
	args := map[string]interface{}{
		"folder": remote_folder,
	}

	src_path := path.Join(remote_folder, md_document)

	result, httpResponse, err := pdf_api.PutMarkdownInStorageToPdf(pdf_name, src_path, args)

	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)
	}
}
