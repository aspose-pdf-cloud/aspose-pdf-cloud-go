/**
 *
 * Copyright (c) 2021 Aspose.PDF Cloud
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

func TestGetDocumentScreenAnnotations(t *testing.T) {

	name := "PdfWithScreenAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentScreenAnnotations(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentScreenAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageScreenAnnotations(t *testing.T) {

	name := "PdfWithScreenAnnotations.pdf"
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageScreenAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageScreenAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetScreenAnnotation(t *testing.T) {

	name := "PdfWithScreenAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseAnnotations, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentScreenAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetScreenAnnotation(name, annotationID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetScreenAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPageScreenAnnotations(t *testing.T) {

	name := "PdfWithScreenAnnotations.pdf"
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	attachmentFile := "ScreenMovie.swf"
	if err := GetBaseTest().UploadFile(attachmentFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	annotation := ScreenAnnotation{
		Name:                "Name",
		Rect:                &Rectangle{LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags:               []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		ZIndex:              1,
		Title:               "Title",
		FilePath:            GetBaseTest().remoteFolder + "/" + attachmentFile,
		Modified:            "02/02/2018 12:00:00.000 AM",
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostPageScreenAnnotations(name, pageNumber, []ScreenAnnotation{annotation}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostPageScreenAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutScreenAnnotation(t *testing.T) {

	name := "PdfWithScreenAnnotations.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	attachmentFile := "ScreenMovie.swf"
	if err := GetBaseTest().UploadFile(attachmentFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	annotation := ScreenAnnotation{
		Name:                "Name Updated",
		Rect:                &Rectangle{LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags:               []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		ZIndex:              1,
		Title:               "Title Updated",
		FilePath:            GetBaseTest().remoteFolder + "/" + attachmentFile,
		Modified:            "02/02/2018 12:01:02.000 AM",
	}

	responseAnnotations, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentScreenAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutScreenAnnotation(name, annotationID, annotation, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutScreenAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetScreenAnnotationData(t *testing.T) {

	name := "PdfWithScreenAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseAnnotations, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentScreenAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetScreenAnnotationData(name, annotationID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetScreenAnnotationData - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutScreenAnnotationDataExtract(t *testing.T) {

	name := "PdfWithScreenAnnotations.pdf"
	outFilePath := GetBaseTest().remoteFolder + "/screen.dat"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseAnnotations, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentScreenAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutScreenAnnotationDataExtract(name, annotationID, outFilePath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetScreenAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
