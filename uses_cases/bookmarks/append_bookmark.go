package main

import (
	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func appendBookmark(pdf_api *asposepdfcloud.PdfApiService, document_name string, title string, remote_folder string) {
	bookmarkPath := "1"
	args := map[string]interface{}{
		"folder": remote_folder,
	}

	bookmark := asposepdfcloud.Bookmark{
		Action:            "GoTo",
		Bold:              true,
		Italic:            false,
		Title:             title,
		PageDisplay:       "XYZ",
		PageDisplayBottom: 10,
		PageDisplayLeft:   10,
		PageDisplayRight:  10,
		PageDisplayTop:    10,
		PageDisplayZoom:   2,
		PageNumber:        1,
		Color:             &asposepdfcloud.Color{A: 0x00, R: 0x00, G: 0xFF, B: 0x00},
	}
	_, _, _ = pdf_api.PostBookmark(document_name, bookmarkPath, []asposepdfcloud.Bookmark{bookmark}, args)
}
