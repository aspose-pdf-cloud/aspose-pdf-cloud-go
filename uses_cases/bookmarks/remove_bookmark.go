package main

import (
	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func removeBookmark(pdf_api *asposepdfcloud.PdfApiService, document_name string, bookmark_path string, remote_folder string) {
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	_, _, _ = pdf_api.DeleteBookmark(document_name, bookmark_path, args)
}
