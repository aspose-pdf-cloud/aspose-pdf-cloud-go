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
	"encoding/base64"
	"fmt"
	"path/filepath"
	"testing"
)

func TestGetDocument(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocument(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostOptimizeDocument(t *testing.T) {
	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	optimizeOptions := OptimizeOptions{
		AllowReusePageContent: false,
		CompressImages:        true,
		ImageQuality:          int32(100),
		LinkDuplcateStreams:   true,
		RemoveUnusedObjects:   true,
		RemoveUnusedStreams:   true,
		UnembedFonts:          true,
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PostOptimizeDocument(name, optimizeOptions, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostOptimizeDocument - %db\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostOptimizeDocumentWithPassword(t *testing.T) {
	name := "4pagesEncrypted.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	optimizeOptions := OptimizeOptions{
		Password:              base64.StdEncoding.EncodeToString([]byte("user $^Password!&")),
		AllowReusePageContent: false,
		CompressImages:        true,
		ImageQuality:          int32(100),
		LinkDuplcateStreams:   true,
		RemoveUnusedObjects:   true,
		RemoveUnusedStreams:   true,
		UnembedFonts:          true,
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PostOptimizeDocument(name, optimizeOptions, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostOptimizeDocumentWithPassword - %db\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostSplitDocument(t *testing.T) {
	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PostSplitDocument(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostSplitDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostSplitRangePdfDocument(t *testing.T) {
	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	options := SplitRangePdfOptions{
		PageRanges: []PageRange{
			{To: 2},
			{From: 3},
			{From: 2, To: 3},
		},
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PostSplitRangePdfDocument(name, options, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else if len(response.Result.Documents) != 3 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostSplitRangePdfDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutCreateDocument(t *testing.T) {
	name := "pdf_go.pdf"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutCreateDocument(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutCreateDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostCreateDocument(t *testing.T) {
	name := "pdf_go_post.pdf"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	config := DocumentConfig{
		PagesCount: 2,
		DocumentProperties: &DocumentProperties{
			List: []DocumentProperty{
				{
					BuiltIn: false,
					Name:    "prop1",
					Value:   "Val1",
				},
			},
		},
		DisplayProperties: &DisplayProperties{
			CenterWindow: true,
			HideMenuBar:  true,
		},
		DefaultPageConfig: &DefaultPageConfig{
			Height: 100,
			Width:  100,
		},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostCreateDocument(name, config, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostCreateDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetDocumentDisplayProperties(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentDisplayProperties(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentDisplayProperties - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutDocumentDisplayProperties(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	displayProperties := DisplayProperties{
		CenterWindow:          true,
		Direction:             DirectionL2R,
		DisplayDocTitle:       true,
		HideMenuBar:           true,
		HideToolBar:           true,
		HideWindowUI:          true,
		NonFullScreenPageMode: PageModeUseNone,
		PageLayout:            PageLayoutTwoPageLeft,
		PageMode:              PageModeUseOC,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutDocumentDisplayProperties(name, displayProperties, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutDocumentDisplayProperties - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostOrganizeDocument(t *testing.T) {
	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	response, httpResponse, err := GetBaseTest().PdfAPI.PostOrganizeDocument(
		name, "1,4-2", filepath.Join(GetBaseTest().remoteFolder, name), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostOrganizeDocument - %db\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostOrganizeDocuments(t *testing.T) {
	name1 := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name1); err != nil {
		t.Error(err)
	}
	name2 := "marketing.pdf"
	if err := GetBaseTest().UploadFile(name2); err != nil {
		t.Error(err)
	}
	request := OrganizeDocumentRequest{
		List: []OrganizeDocumentData{
			{Path: filepath.Join(GetBaseTest().remoteFolder, name1), Pages: "4-2"},
			{Path: filepath.Join(GetBaseTest().remoteFolder, name2), Pages: "2"},
			{Path: filepath.Join(GetBaseTest().remoteFolder, name1), Pages: "3,1"},
		},
	}
	args := make(map[string]interface{})
	response, httpResponse, err := GetBaseTest().PdfAPI.PostOrganizeDocuments(
		request, filepath.Join(GetBaseTest().remoteFolder, "OrganizeMany.pdf"), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostOrganizeDocuments - %db\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
