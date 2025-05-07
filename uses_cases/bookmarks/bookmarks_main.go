package main

import (
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

const (
	REMOTE_FOLDER  = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER   = "c:\\Samples"
	PDF_DOCUMENT   = "sample.pdf"
	PDF_OUTPUT     = "output_sample.pdf"
	BOOKMARK_TITLE = "NEW Bookmark Title XYZ"
	BOOKMARK_PATH  = "/1"
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	AppSID := "*************"
	AppKey := "*************"
	pdfApi := asposepdfcloud.NewPdfApiService(AppSID, AppKey, "")
	return pdfApi
}

func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, _ := os.Open(filepath.Join(LOCAL_FOLDER, name))
	_, _, _ = pdf_api.UploadFile(filepath.Join(REMOTE_FOLDER, name), file, args)
}
func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, _, _ := pdf_api.DownloadFile(path.Join(REMOTE_FOLDER, name), args)
	fileName := path.Join(LOCAL_FOLDER, PDF_OUTPUT)
	f, _ := os.Create(fileName)
	_, _ = f.Write(result_data)
}

func main() {
	pdfApi := initPdfApi()
	uploadFile(pdfApi, PDF_DOCUMENT)
	extractDocumentBookmarks(pdfApi, PDF_DOCUMENT, REMOTE_FOLDER)
	extractBookmark(pdfApi, PDF_DOCUMENT, BOOKMARK_PATH, REMOTE_FOLDER)
	appendBookmark(pdfApi, PDF_DOCUMENT, BOOKMARK_PATH, BOOKMARK_TITLE, REMOTE_FOLDER)
	replaceBookmark(pdfApi, PDF_DOCUMENT, BOOKMARK_PATH, BOOKMARK_TITLE, REMOTE_FOLDER)
	removeBookmark(pdfApi, PDF_DOCUMENT, BOOKMARK_PATH, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_DOCUMENT)
}
