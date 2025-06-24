package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

const (
	REMOTE_FOLDER             = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER              = "c:\\Samples"
	PDF_DOCUMENT              = "sample.pdf"
	PDF_DOCUMENT_WITH_ATTACH  = "sample_file_with_attachment.pdf"
	PDF_OUTPUT                = "output_sample.pdf"
	NEW_ATTACHMENT_FILE       = "sample_video.mp4"
	NEW_ATTACHMENT_MIME       = "video/mp4"
	NEW_ATTACHMENT_DECRIPTION = "This is a sample attachment"
	PAGE_NUMBER               = 2
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	// Initialize Credentials and create Pdf.Cloud service object
	AppSID := "******" // Your Application SID
	AppKey := "******" // Your Application Key

	pdfApi := asposepdfcloud.NewPdfApiService(AppSID, AppKey, "")
	return pdfApi
}

func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	// Upload local file to the Pdf.Cloud folder
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, _ := os.Open(filepath.Join(LOCAL_FOLDER, name))
	_, _, _ = pdf_api.UploadFile(filepath.Join(REMOTE_FOLDER, name), file, args)
}

func saveByteArrayToFile(local_folder string, file_name string, data []byte) {
	// Save byte array data to the local folder
	fileName := path.Join(local_folder, file_name)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		size, err := f.Write(data)
		if err != nil || size == 0 {
			fmt.Println("Failures in downloading result!")
		} else {
			fmt.Println("Result file'" + fileName + "' successfully downloaded!")
		}
	}
}

func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	// Download modified Pdf document to local folder from the Pdf.Cloud folder
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, _, _ := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	saveByteArrayToFile(LOCAL_FOLDER, PDF_OUTPUT, result_data)
}
