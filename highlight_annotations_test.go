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
 
 func TestGetDocumentHighlightAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"	

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentHighlightAnnotations(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentHighlightAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageHighlightAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"	
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageHighlightAnnotations(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageHighlightAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetHighlightAnnotation(t *testing.T) {
 
	name := "PdfWithAnnotations.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	responseAnnotations, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentHighlightAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetHighlightAnnotation(name, annotationID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			t.Fail()
	} else {
		fmt.Printf("%d\tTestGetHighlightAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPageHighlightAnnotations(t *testing.T) {

	name := "PdfWithAnnotations.pdf"	
	var pageNumber int32 = 2
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	annotation := HighlightAnnotation {
		Name: "Name",
		Rect: &Rectangle{ LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags: []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		RichText: "Rich Text",
		Subject: "Text Box Subj",
		ZIndex: 1,
		Title: "Title",
		QuadPoints: []Point{
			Point{X: 10, Y: 10},
			Point{X: 20, Y: 10},
			Point{X: 10, Y: 20},
			Point{X: 10, Y: 10},
		},
		Modified: "02/02/2018 00:00:00.000 AM",
	}


	response, httpResponse, err := GetBaseTest().PdfAPI.PostPageHighlightAnnotations(name, pageNumber, []HighlightAnnotation{annotation}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostPageHighlightAnnotations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutHighlightAnnotation(t *testing.T) {

	name := "PdfWithAnnotations.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	annotation := HighlightAnnotation {
		Name: "Name Updated",
		Rect: &Rectangle{ LLX: 100, LLY: 100, URX: 200, URY: 200},
		Flags: []AnnotationFlags{AnnotationFlagsDefault},
		HorizontalAlignment: HorizontalAlignmentCenter,
		RichText: "Rich Text Updated",
		Subject: "Text Box Subj Updated",
		ZIndex: 1,
		Title: "Title Updated",
		QuadPoints: []Point{
			Point{X: 10, Y: 10},
			Point{X: 20, Y: 10},
			Point{X: 10, Y: 20},
			Point{X: 10, Y: 10},
		},
		Modified: "02/02/2018 00:00:00.000 AM",
	}

	responseAnnotations, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentHighlightAnnotations(name, args)
	if err != nil {
		t.Error(err)
	}
	annotationID := responseAnnotations.Annotations.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutHighlightAnnotation(name, annotationID, annotation, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutHighlightAnnotation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}