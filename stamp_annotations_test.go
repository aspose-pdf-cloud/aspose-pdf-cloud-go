/**
 *
 * Copyright (c) 2024 Aspose.PDF Cloud
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */
package asposepdfcloud

import (
	"fmt"
	"testing"
)

func TestGetDocumentStampAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentStampAnnotations(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentStampAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageStampAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageStampAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageStampAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetStampAnnotation(t *testing.T) {

	name := "PdfWithAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseAnnotations, _, err := GetBaseTest().PdfAPI.GetDocumentStampAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetStampAnnotation(name, annotationID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetStampAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPageStampAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	attachmentFile := "4pages.pdf"
	if err := GetBaseTest().UploadFile(attachmentFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	annotation := StampAnnotation{
		Name:                "Name",
		Rect:                &Rectangle{LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags:               []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		RichText:            "Rich Text",
		Subject:             "Text Box Subj",
		ZIndex:              1,
		Title:               "Title",
		FilePath:            GetBaseTest().remoteFolder + "/" + attachmentFile,
		Modified:            "02/02/2018 12:00:00.000 AM",
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostPageStampAnnotations(name, pageNumber, []StampAnnotation{annotation}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostPageStampAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutStampAnnotation(t *testing.T) {

	name := "PdfWithAnnotations.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	attachmentFile := "4pages.pdf"
	if err := GetBaseTest().UploadFile(attachmentFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	annotation := StampAnnotation{
		Name:                "Name Updated",
		Rect:                &Rectangle{LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags:               []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		RichText:            "Rich Text Updated",
		Subject:             "Text Box Subj Updated",
		ZIndex:              1,
		Title:               "Title Updated",
		FilePath:            GetBaseTest().remoteFolder + "/" + attachmentFile,
		Modified:            "02/02/2018 12:01:02.000 AM",
	}

	responseAnnotations, _, err := GetBaseTest().PdfAPI.GetDocumentStampAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutStampAnnotation(name, annotationID, annotation, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutStampAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetStampAnnotationData(t *testing.T) {

	name := "PdfWithAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseAnnotations, _, err := GetBaseTest().PdfAPI.GetDocumentStampAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetStampAnnotationData(name, annotationID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetStampAnnotationData - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutStampAnnotationDataExtract(t *testing.T) {

	name := "PdfWithAnnotations.pdf"
	outFilePath := GetBaseTest().remoteFolder + "/stamp.dat"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseAnnotations, _, err := GetBaseTest().PdfAPI.GetDocumentStampAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutStampAnnotationDataExtract(name, annotationID, outFilePath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetStampAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
