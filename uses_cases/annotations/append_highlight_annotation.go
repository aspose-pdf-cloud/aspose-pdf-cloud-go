package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func appendHighlightAnnotation(pdf_api *asposepdfcloud.PdfApiService, document string, output_document string, page_num int32, text string, subject string, contents string, description string, rect *asposepdfcloud.Rectangle, remote_folder string) {
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	new_annotation := asposepdfcloud.HighlightAnnotation{
		Rect:                rect,
		Name:                "Highlight_Text_Annotation",
		Flags:               []asposepdfcloud.AnnotationFlags{asposepdfcloud.AnnotationFlagsDefault},
		HorizontalAlignment: asposepdfcloud.HorizontalAlignmentLeft,
		VerticalAlignment:   asposepdfcloud.VerticalAlignmentTop,
		RichText:            text,
		Subject:             subject,
		Contents:            contents,
		ZIndex:              1,
		Title:               description,
		Color:               &asposepdfcloud.Color{A: 0x00, R: 0x00, G: 0xFF, B: 0x00},
		QuadPoints: []asposepdfcloud.Point{
			{X: 10, Y: 10},
			{X: 20, Y: 10},
			{X: 10, Y: 20},
			{X: 10, Y: 10},
		},
		Modified: "03/27/2025 00:00:00.000 AM",
		//Icon:     "Conment",
	}
	result, httpResponse, err := pdf_api.PostPageHighlightAnnotations(
		document, page_num, []asposepdfcloud.HighlightAnnotation{new_annotation}, args,
	)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println(result)

		downloadFile(pdf_api, document, output_document)
	}
}
