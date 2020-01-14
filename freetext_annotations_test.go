 /**
 *
 *   Copyright (c) 2020 Aspose.PDF Cloud
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
 
 func TestGetDocumentFreeTextAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"	

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentFreeTextAnnotations(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentFreeTextAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageFreeTextAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"	
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageFreeTextAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageFreeTextAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetFreeTextAnnotation(t *testing.T) {
 
	name := "PdfWithAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	responseAnnotations, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentFreeTextAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetFreeTextAnnotation(name, annotationID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			t.Fail()
	} else {
		fmt.Printf("%d\tTestGetFreeTextAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPageFreeTextAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"	
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	textStyle := TextStyle {
		FontSize: float64(12), 
		Font: "Arial",
		ForegroundColor: &Color { A: int32(0xFF), R: 0, G: int32(0xFF), B: 0},
		BackgroundColor: &Color { A: int32(0xFF), R: int32(0xFF), G: 0, B: 0},
	}

	annotation := FreeTextAnnotation {
		Name: "Name",
		TextStyle: &textStyle,
		Rect: &Rectangle{ LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags: []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		Intent: FreeTextIntentFreeTextTypeWriter,
		RichText: "Rich Text",
		Subject: "Text Box Subj",
		ZIndex: 1,
		Justification: JustificationCenter,
		Title: "Title",
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostPageFreeTextAnnotations(name, pageNumber, []FreeTextAnnotation{annotation}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostPageFreeTextAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutFreeTextAnnotation(t *testing.T) {

	name := "PdfWithAnnotations.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	textStyle := TextStyle {
		FontSize: float64(12), 
		Font: "Arial",
		ForegroundColor: &Color { A: int32(0xFF), R: 0, G: int32(0xFF), B: 0},
		BackgroundColor: &Color { A: int32(0xFF), R: int32(0xFF), G: 0, B: 0},
	}

	annotation := FreeTextAnnotation {
		Name: "Updated Name",
		TextStyle: &textStyle,
		Rect: &Rectangle{ LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags: []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		Intent: FreeTextIntentFreeTextTypeWriter,
		RichText: "Updated Rich Text",
		Subject: "Updated Text Box Subj",
		ZIndex: 1,
		Justification: JustificationCenter,
		Title: "Updated Title",
	}

	responseAnnotations, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentFreeTextAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutFreeTextAnnotation(name, annotationID, annotation, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutFreeTextAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}