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

func TestPostDocumentTextHeader(t *testing.T) {

	name := "4pages.pdf"
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"startPageNumber": int32(2),
		"endPageNumber": int32(3),
	}

	header := TextHeader {
		Background: true,
		LeftMargin: 1,
		RightMargin: 2,
		TopMargin: 3,
		HorizontalAlignment: HorizontalAlignmentCenter,
		TextAlignment: HorizontalAlignmentCenter,
		Opacity: 1,
		Rotate: RotationNone,
		RotateAngle: 0,
		XIndent: 0,
		YIndent: 0,
		Zoom: 1,
		Value: "Header",
		TextState: &TextState{FontSize: 14, FontStyle: FontStylesRegular},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentTextHeader(name, header, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostDocumentTextHeader - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostDocumentTextFooter(t *testing.T) {

	name := "4pages.pdf"
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"startPageNumber": int32(2),
		"endPageNumber": int32(3),
	}

	footer := TextFooter {
		Background: true,
		LeftMargin: 1,
		RightMargin: 2,
		BottomMargin: 3,
		HorizontalAlignment: HorizontalAlignmentCenter,
		TextAlignment: HorizontalAlignmentCenter,
		Opacity: 1,
		Rotate: RotationNone,
		RotateAngle: 0,
		XIndent: 0,
		YIndent: 0,
		Zoom: 1,
		Value: "Header",
		TextState: &TextState{FontSize: 14, FontStyle: FontStylesRegular},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentTextFooter(name, footer, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostDocumentTextFooter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostDocumentImageHeader(t *testing.T) {

	name := "4pages.pdf"
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	image := "Koala.jpg"
	
	if err := GetBaseTest().UploadFile(image); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"startPageNumber": int32(2),
		"endPageNumber": int32(3),
	}

	header := ImageHeader {
		Background: true,
		LeftMargin: 1,
		RightMargin: 2,
		TopMargin: 3,
		HorizontalAlignment: HorizontalAlignmentCenter,
		Opacity: 1,
		Rotate: RotationNone,
		RotateAngle: 0,
		XIndent: 0,
		YIndent: 0,
		Zoom: 1,
		FileName: GetBaseTest().remoteFolder + "/" + image,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentImageHeader(name, header, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostDocumentImageHeader - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostDocumentImageFooter(t *testing.T) {

	name := "4pages.pdf"
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	image := "Koala.jpg"
	
	if err := GetBaseTest().UploadFile(image); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"startPageNumber": int32(2),
		"endPageNumber": int32(3),
	}

	footer := ImageFooter {
		Background: true,
		LeftMargin: 1,
		RightMargin: 2,
		BottomMargin: 3,
		HorizontalAlignment: HorizontalAlignmentCenter,
		Opacity: 1,
		Rotate: RotationNone,
		RotateAngle: 0,
		XIndent: 0,
		YIndent: 0,
		Zoom: 1,
		FileName: GetBaseTest().remoteFolder + "/" + image,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentImageFooter(name, footer, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostDocumentImageFooter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}