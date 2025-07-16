package main

import (
	"fmt"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

func appendFreeTextAnnotation(pdf_api *asposepdfcloud.PdfApiService, document string, output_document string, page_num int32, text string, subject string, contents string, description string, rect *asposepdfcloud.Rectangle, remote_folder string) {
	uploadFile(pdf_api, document)

	args := map[string]interface{}{
		"folder": remote_folder,
	}

	text_style := asposepdfcloud.TextStyle{
		FontSize:        20,
		Font:            "Arial",
		ForegroundColor: &asposepdfcloud.Color{A: 0xFF, R: 0x00, G: 0xFF, B: 0x00},
		BackgroundColor: &asposepdfcloud.Color{A: 0xFF, R: 0xFF, G: 0x00, B: 0x00},
	}

	new_annotation := asposepdfcloud.FreeTextAnnotation{
		TextStyle:           &text_style,
		Flags:               []asposepdfcloud.AnnotationFlags{asposepdfcloud.AnnotationFlagsDefault},
		HorizontalAlignment: asposepdfcloud.HorizontalAlignmentCenter,
		Intent:              asposepdfcloud.FreeTextIntentFreeTextTypeWriter,
		RichText:            text,
		Subject:             subject,
		Contents:            contents,
		ZIndex:              1,
		Title:               description,
		Justification:       asposepdfcloud.JustificationCenter,
		Rect:                rect,
		Name:                "Free_Text_Annotation",
		Modified:            "05/22/2025 11:30:00.000 AM",
	}

	result, httpResponse, err := pdf_api.PostPageFreeTextAnnotations(
		document, page_num, []asposepdfcloud.FreeTextAnnotation{new_annotation}, args,
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
