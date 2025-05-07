package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func extractBookmark(pdf_api *asposepdfcloud.PdfApiService, document_name string, bookmark_path string, remote_folder string) {
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	result, _, _ := pdf_api.GetBookmarks(document_name, bookmark_path, args)
	fmt.Println(result.Bookmarks)
}
