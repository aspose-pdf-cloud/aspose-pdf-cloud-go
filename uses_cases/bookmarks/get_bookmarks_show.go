package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func extractDocumentBookmarks(pdf_api *asposepdfcloud.PdfApiService, document_name string, remote_folder string) {
	args := map[string]interface{}{
		"folder": remote_folder,
	}
	result, _, _ := pdf_api.GetDocumentBookmarks(document_name, args)
	fmt.Println(result.Bookmarks)
}
