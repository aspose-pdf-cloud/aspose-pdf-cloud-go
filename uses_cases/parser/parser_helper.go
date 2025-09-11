package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

const (
	REMOTE_FOLDER   = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER    = "c:\\Samples"
	PDF_DOCUMENT    = "sample.pdf"
	PDF_OUTPUT      = "output_pages.pdf"
	XML_OUTPUT_FILE = "output_sample.xml"
	FDF_OUTPUT_FILE = "output_sample.fdf"
	PAGE_NUMBER     = 1

	AppSID = "**********"     // Your Application SID
	AppKey = "**********"     // Your Application Key
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	pdfApi := asposepdfcloud.NewPdfApiService(AppSID, AppKey, "")
	return pdfApi
}

// Upload local file to the remote folder with check errors
func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, err := os.Open(filepath.Join(LOCAL_FOLDER, name))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, httpResponse, err := pdf_api.UploadFile(path.Join(REMOTE_FOLDER, name), file, args)
		if err != nil {
			fmt.Println(err.Error())
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			fmt.Println("Unexpected error!")
		} else {
			fmt.Println("File '" + name + " ' successfully uploaded.")
		}
	}
}

// Download file from remote folder and save it locally with check errors
func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string, output_name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, httpResponse, err := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fileName := path.Join(LOCAL_FOLDER, output_name)
		f, _ := os.Create(fileName)
		_, _ = f.Write(result_data)
		fmt.Println("File '" + fileName + "' successfully downloaded.")
	}
}
