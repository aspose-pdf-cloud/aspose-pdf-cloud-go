package main

import (
	"fmt"
	"path"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func converJpegToPdf(pdf_api *asposepdfcloud.PdfApiService, jpg_file string, pdf_name string, remote_folder string) {
	imageTemplateList := []asposepdfcloud.ImageTemplate{}

	imageTemplate := asposepdfcloud.ImageTemplate{
		ImagePath:    path.Join(remote_folder, jpg_file),
		ImageSrcType: asposepdfcloud.ImageSrcTypeCommon,
	}
	imageTemplateList = append(imageTemplateList, imageTemplate)

	uploadFile(pdf_api, jpg_file)

	imageTemplatesRequest := asposepdfcloud.ImageTemplatesRequest{
		IsOCR:      true,
		OCRLangs:   "eng",
		ImagesList: imageTemplateList,
	}

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	result, httpResponse, err := pdf_api.PutImageInStorageToPdf(pdf_name, imageTemplatesRequest, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)
	}
}
