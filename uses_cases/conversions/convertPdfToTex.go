package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func convertPdfToTex(pdf_api *asposepdfcloud.PdfApiService, pdf_name string, tex_name string, remote_folder string) {
	outPath := path.Join(remote_folder, tex_name)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.PutPdfInStorageToTeX(pdf_name, outPath, args)

	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)
	}
}
