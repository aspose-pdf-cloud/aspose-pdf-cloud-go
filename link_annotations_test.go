 /**
 *
 *   Copyright (c) 2019 Aspose.PDF Cloud
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

func TestGetPageLinkAnnotation(t *testing.T) {

	name := "PdfWithLinks.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	responseLinks, httpResponse, err := GetBaseTest().PdfAPI.GetPageLinkAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	linkID := responseLinks.Links.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageLinkAnnotation(name, pageNumber, linkID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageLinkAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteLinkAnnotation(t *testing.T) {

	name := "PdfWithLinks.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	responseLinks, httpResponse, err := GetBaseTest().PdfAPI.GetPageLinkAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	linkID := responseLinks.Links.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteLinkAnnotation(name, linkID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteLinkAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageLinkAnnotations(t *testing.T) {

	name := "PdfWithLinks.pdf"	
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageLinkAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageLinkAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPageLinkAnnotations(t *testing.T) {

	name := "PdfWithLinks.pdf"	
	pageNumber := int32(2)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	annotation := LinkAnnotation {
		Rect: &Rectangle{ LLX: 100, LLY: 100, URX: 500, URY: 500},
		ActionType: LinkActionTypeGoToURIAction,
		Action: "https://products.aspose.cloud/pdf",
		Highlighting: LinkHighlightingModeInvert,
		Color: &Color{A: 0xFF, R: 0xAA, G: 0xAA, B: 0xAA},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostPageLinkAnnotations(name, pageNumber, []LinkAnnotation{annotation}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostPageLinkAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutLinkAnnotation(t *testing.T) {

	name := "PdfWithLinks.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	annotation := LinkAnnotation {
		Rect: &Rectangle{ LLX: 100, LLY: 100, URX: 500, URY: 500},
		ActionType: LinkActionTypeGoToURIAction,
		Action: "https://products.aspose.cloud/pdf",
		Highlighting: LinkHighlightingModeInvert,
		Color: &Color{A: 0xFF, R: 0xAA, G: 0xAA, B: 0xAA},
	}

	responseLinks, httpResponse, err := GetBaseTest().PdfAPI.GetPageLinkAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	linkID := responseLinks.Links.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutLinkAnnotation(name, linkID, annotation, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutLinkAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeletePageLinkAnnotations(t *testing.T) {

	name := "PdfWithLinks.pdf"	
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeletePageLinkAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeletePageLinkAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteDocumentLinkAnnotations(t *testing.T) {

	name := "PdfWithLinks.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteDocumentLinkAnnotations(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteDocumentLinkAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

 func TestGetLinkAnnotation(t *testing.T) {

	name := "PdfWithLinks.pdf"	
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	responseLinks, httpResponse, err := GetBaseTest().PdfAPI.GetPageLinkAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	linkID := responseLinks.Links.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetLinkAnnotation(name, linkID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetLinkAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}






