/**
 *
 * Copyright (c) 2025 Aspose.PDF Cloud
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

func TestGetDocumentBookmarks(t *testing.T) {
	name := "PdfWithBookmarks.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentBookmarks(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentBookmarks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetBookmarks(t *testing.T) {
	name := "PdfWithBookmarks.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	bookmarkPath := "1"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetBookmarks(name, bookmarkPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetBookmarks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetBookmark(t *testing.T) {
	name := "PdfWithBookmarks.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	bookmarkPath := "1"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetBookmark(name, bookmarkPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetBookmark - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteDocumentBookmarks(t *testing.T) {
	name := "PdfWithBookmarks.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteDocumentBookmarks(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteDocumentBookmarks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteBookmark(t *testing.T) {
	name := "PdfWithBookmarks.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	bookmarkPath := "1"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteBookmark(name, bookmarkPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteBookmark - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostBookmark(t *testing.T) {
	name := "PdfWithBookmarks.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	bookmarkPath := "1"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	bookmark := Bookmark{
		Action:            "GoTo",
		Bold:              true,
		Italic:            false,
		Title:             "New Bookmark XYZ",
		PageDisplay:       "XYZ",
		PageDisplayBottom: 10,
		PageDisplayLeft:   10,
		PageDisplayRight:  10,
		PageDisplayTop:    10,
		PageDisplayZoom:   2,
		PageNumber:        2,
		Color:             &Color{A: 0x00, R: 0x00, G: 0xFF, B: 0x00},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostBookmark(name, bookmarkPath, []Bookmark{bookmark}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostBookmark - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutBookmark(t *testing.T) {
	name := "PdfWithBookmarks.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	bookmarkPath := "1"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	bookmark := Bookmark{
		Action:            "GoTo",
		Bold:              true,
		Italic:            false,
		Title:             "New Bookmark XYZ",
		PageDisplay:       "XYZ",
		PageDisplayBottom: 10,
		PageDisplayLeft:   10,
		PageDisplayRight:  10,
		PageDisplayTop:    10,
		PageDisplayZoom:   2,
		PageNumber:        2,
		Color:             &Color{A: 0x00, R: 0x00, G: 0xFF, B: 0x00},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutBookmark(name, bookmarkPath, bookmark, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutBookmark - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
