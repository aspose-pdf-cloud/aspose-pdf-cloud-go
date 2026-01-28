/**
 *
 * Copyright (c) 2026 Aspose.PDF Cloud
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

func TestGetDocumentStamps(t *testing.T) {

	name := "PageNumberStamp.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentStamps(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteDocumentStamps(t *testing.T) {

	name := "PageNumberStamp.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteDocumentStamps(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteDocumentStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageStamps(t *testing.T) {

	name := "PageNumberStamp.pdf"
	var pageNumber int32 = 1

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageStamps(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeletePageStamps(t *testing.T) {

	name := "PageNumberStamp.pdf"
	var pageNumber int32 = 1

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeletePageStamps(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeletePageStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPageTextStamps(t *testing.T) {

	name := "PageNumberStamp.pdf"
	var pageNumber int32 = 1

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	stamp := TextStamp{
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		TextAlignment:       HorizontalAlignmentCenter,
		Value:               "Text Stamp",
		TextState:           &TextState{FontSize: 14, FontStyle: FontStylesRegular, Font: "Arial"},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostPageTextStamps(name, pageNumber, []TextStamp{stamp}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostPageTextStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostDocumentTextStamps(t *testing.T) {

	name := "PageNumberStamp.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	stamp := TextStamp{
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		TextAlignment:       HorizontalAlignmentCenter,
		Value:               "Text Stamp",
		TextState:           &TextState{FontSize: 14, FontStyle: FontStylesRegular, Font: "Arial"},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentTextStamps(name, []TextStamp{stamp}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostDocumentTextStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostDocumentTextStampsPageSpecified(t *testing.T) {

	name := "PageNumberStamp.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	stamp1 := TextStampPageSpecified{
		PageNumber:          2,
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		TextAlignment:       HorizontalAlignmentCenter,
		Value:               "Text Stamp 1",
		TextState:           &TextState{FontSize: 14, FontStyle: FontStylesRegular, Font: "Arial"},
	}

	stamp2 := TextStampPageSpecified{
		PageNumber:          4,
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		TextAlignment:       HorizontalAlignmentCenter,
		Value:               "Text Stamp 2",
		TextState:           &TextState{FontSize: 14, FontStyle: FontStylesRegular, Font: "Arial"},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentTextStampsPageSpecified(name, []TextStampPageSpecified{stamp1, stamp2}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostDocumentTextStampsPageSpecified - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPageImageStamps(t *testing.T) {

	name := "PageNumberStamp.pdf"
	var pageNumber int32 = 1
	image := "Koala.jpg"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	if err := GetBaseTest().UploadFile(image); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	stamp := ImageStamp{
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		FileName:            GetBaseTest().remoteFolder + "/" + image,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostPageImageStamps(name, pageNumber, []ImageStamp{stamp}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostPageImageStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostDocumentImageStamps(t *testing.T) {

	name := "PageNumberStamp.pdf"
	image := "Koala.jpg"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	if err := GetBaseTest().UploadFile(image); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	stamp := ImageStamp{
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		FileName:            GetBaseTest().remoteFolder + "/" + image,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentImageStamps(name, []ImageStamp{stamp}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostDocumentImageStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostDocumentImageStampsPageSpecified(t *testing.T) {

	name := "PageNumberStamp.pdf"
	image := "Koala.jpg"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	if err := GetBaseTest().UploadFile(image); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	stamp1 := ImageStampPageSpecified{
		PageNumber:          2,
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		FileName:            GetBaseTest().remoteFolder + "/" + image,
	}

	stamp2 := ImageStampPageSpecified{
		PageNumber:          4,
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		FileName:            GetBaseTest().remoteFolder + "/" + image,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentImageStampsPageSpecified(name, []ImageStampPageSpecified{stamp1, stamp2}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostDocumentImageStampsPageSpecified - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPagePdfPageStamps(t *testing.T) {

	name := "PageNumberStamp.pdf"
	var pageNumber int32 = 1
	pdf := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	if err := GetBaseTest().UploadFile(pdf); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	stamp := PdfPageStamp{
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		FileName:            GetBaseTest().remoteFolder + "/" + pdf,
		PageIndex:           2,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostPagePdfPageStamps(name, pageNumber, []PdfPageStamp{stamp}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostPagePdfPageStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteStamp(t *testing.T) {

	name := "PageNumberStamp.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseStamps, _, err := GetBaseTest().PdfAPI.GetDocumentStamps(name, args)
	if err != nil {
		t.Error(err)
	}
	stampID := responseStamps.Stamps.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteStamp(name, stampID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteStamp - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostDocumentPageNumberStamps(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder":          GetBaseTest().remoteFolder,
		"startPageNumber": int32(2),
		"endPageNumber":   int32(3),
	}

	stamp := PageNumberStamp{
		Background:          true,
		LeftMargin:          1,
		RightMargin:         2,
		TopMargin:           3,
		BottomMargin:        4,
		HorizontalAlignment: HorizontalAlignmentCenter,
		VerticalAlignment:   VerticalAlignmentCenter,
		Opacity:             1,
		Rotate:              RotationNone,
		RotateAngle:         0,
		XIndent:             0,
		YIndent:             0,
		Zoom:                1,
		StartingNumber:      int32(3),
		Value:               "Page #",
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentPageNumberStamps(name, stamp, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostDocumentPageNumberStamps - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
