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

func TestGetDocumentTextAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentTextAnnotations(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentTextAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageTextAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageTextAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageTextAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetTextAnnotation(t *testing.T) {

	name := "PdfWithAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseAnnotations, _, err := GetBaseTest().PdfAPI.GetDocumentTextAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetTextAnnotation(name, annotationID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetTextAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPageTextAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	annotation := TextAnnotation{
		Name:                "Name",
		Rect:                &Rectangle{LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags:               []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		RichText:            "Rich Text",
		Subject:             "Text Box Subj",
		ZIndex:              1,
		Title:               "Title",
		State:               AnnotationStateUndefined,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostPageTextAnnotations(name, pageNumber, []TextAnnotation{annotation}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostPageTextAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutTextAnnotation(t *testing.T) {

	name := "PdfWithAnnotations.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	annotation := TextAnnotation{
		Name:                "Updated Name",
		Rect:                &Rectangle{LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags:               []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		RichText:            "Updated Rich Text",
		Subject:             "Updated Text Box Subj",
		ZIndex:              1,
		Title:               "Updated Title",
		State:               AnnotationStateUndefined,
	}

	responseAnnotations, _, err := GetBaseTest().PdfAPI.GetDocumentTextAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutTextAnnotation(name, annotationID, annotation, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutTextAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
