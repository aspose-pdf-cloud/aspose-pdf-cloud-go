package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

const (
	REMOTE_FOLDER        = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER         = "c:\\Samples"
	PDF_DOCUMENT         = "sample.pdf"
	PDF_OUTPUT_HIGHLIGHT = "output_highlight.pdf"
	PDF_OUTPUT_STRIKEOUT = "output_strikeout.pdf"
	PDF_OUTPUT_FREETEXT  = "output_freetext.pdf"
	PDF_OUTPUT_UNDERLINE = "output_underline.pdf"

	NEW_HIGHLIGHT_TEXT        = "NEW HIGHLIGHT TEXT ANNOTATION"
	NEW_HIGHLIGHT_DESCRIPTION = "This is a sample highlight annotation"
	NEW_HIGHLIGHT_SUBJECT     = "Highlight Text Box Subject"
	NEW_HIGHLIGHT_CONTENTS    = "Highlight annotation sample contents"

	NEW_STRIKEOUT_TEXT        = "NEW STRIKEOUT TEXT ANNOTATION"
	NEW_STRIKEOUT_DESCRIPTION = "This is a sample strikeout annotation"
	NEW_STRIKEOUT_SUBJECT     = "Strikeout Text Box Subject"
	NEW_STRIKEOUT_CONTENTS    = "Strikeout annotation sample contents"

	NEW_FREETEXT_TEXT        = "NEW FREE TEXT ANNOTATION"
	NEW_FREETEXT_DESCRIPTION = "This is a sample free text annotation"
	NEW_FREETEXT_SUBJECT     = "Free Text Box Subject"
	NEW_FREETEXT_CONTENTS    = "Free text annotation sample contents"

	NEW_UNDERLINE_TEXT        = "NEW UNDERLINE TEXT ANNOTATION"
	NEW_UNDERLINE_DESCRIPTION = "This is a sample underline annotation"
	NEW_UNDERLINE_SUBJECT     = "Underline Text Box Subject"
	NEW_UNDERLINE_CONTENTS    = "Underline annotation sample contents"

	ANNOTATION_ID = "GI5TAOZVG42CYOBRG4WDKOJVFQ4DGOA"

	PAGE_NUMBER         = 1
	PAGE_NUMBER_EXTRACT = 2
)

var (
	LINK_RECT = asposepdfcloud.Rectangle{LLX: 270, LLY: 510, URX: 380, URY: 530}
)

func initPdfApi() *asposepdfcloud.PdfApiService {
	// Initialize Credentials and create Pdf.Cloud service object
	AppSID := "******" // Your Application SID
	AppKey := "******" // Your Application Key

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
		result, httpResponse, err := pdf_api.UploadFile(path.Join(REMOTE_FOLDER, name), file, args)
		if err != nil {
			fmt.Println(err.Error())
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			fmt.Println("Unexpected error!")
		} else {
			fmt.Println(result)
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
		fmt.Println("File '" + fileName + "'successfully downloaded.")
	}
}

// Show Annotations from array
func showAnnotations(links *[]asposepdfcloud.AnnotationInfo) {
	for i := 0; i < len(*links); i++ {
		fmt.Print("annotation #")
		fmt.Println(i)
		fmt.Print("\tId == '")
		fmt.Print((*links)[i].Id)
		fmt.Println("'")
		fmt.Print("\tName == '")
		fmt.Print((*links)[i].Name)
		fmt.Println("'")
		fmt.Print("\tContents == '")
		fmt.Print((*links)[i].Contents)
		fmt.Println("'")
	}
}

// Show Annotation
func showAnnotation(annotation *asposepdfcloud.TextAnnotation) {
	fmt.Print("annotation '")
	fmt.Println(annotation.Title)
	fmt.Println("'")
	fmt.Print("\tId == '")
	fmt.Print(annotation.Id)
	fmt.Println("'")
	fmt.Print("\tName == '")
	fmt.Print(annotation.Name)
	fmt.Println("'")
	fmt.Print("\tContents == '")
	fmt.Print(annotation.Contents)
	fmt.Println("'")

	fmt.Print("\tIcon == '")
	fmt.Print(annotation.Icon)
	fmt.Println("'")
}
