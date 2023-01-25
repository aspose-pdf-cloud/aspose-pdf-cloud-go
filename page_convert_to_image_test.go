/**
 *
 * Copyright (c) 2023 Aspose.PDF Cloud
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

func TestGetPageConvertToTiff(t *testing.T) {

	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	pageNumber := int32(2)

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageConvertToTiff(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageConvertToTiff - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPageConvertToTiff(t *testing.T) {

	name := "4pages.pdf"
	resFileName := "page.tiff"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	pageNumber := int32(2)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPageConvertToTiff(name, pageNumber, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPageConvertToTiff - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageConvertToJpeg(t *testing.T) {

	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	pageNumber := int32(2)

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageConvertToJpeg(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageConvertToJpeg - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPageConvertToJpeg(t *testing.T) {

	name := "4pages.pdf"
	resFileName := "page.jpg"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	pageNumber := int32(2)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPageConvertToJpeg(name, pageNumber, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPageConvertToJpeg - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageConvertToPng(t *testing.T) {

	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	pageNumber := int32(2)

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageConvertToPng(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageConvertToPng - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPageConvertToPng(t *testing.T) {

	name := "4pages.pdf"
	resFileName := "page.png"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	pageNumber := int32(2)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPageConvertToPng(name, pageNumber, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPageConvertToPng - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageConvertToEmf(t *testing.T) {

	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	pageNumber := int32(2)

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageConvertToEmf(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageConvertToEmf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPageConvertToEmf(t *testing.T) {

	name := "4pages.pdf"
	resFileName := "page.emf"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	pageNumber := int32(2)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPageConvertToEmf(name, pageNumber, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPageConvertToEmf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageConvertToBmp(t *testing.T) {

	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	pageNumber := int32(2)

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageConvertToBmp(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageConvertToBmp - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPageConvertToBmp(t *testing.T) {

	name := "4pages.pdf"
	resFileName := "page.bmp"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	pageNumber := int32(2)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPageConvertToBmp(name, pageNumber, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPageConvertToBmp - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageConvertToGif(t *testing.T) {

	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	pageNumber := int32(2)

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageConvertToGif(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageConvertToGif - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPageConvertToGif(t *testing.T) {

	name := "4pages.pdf"
	resFileName := "page.gif"
	outPath := GetBaseTest().remoteFolder + "/" + resFileName
	pageNumber := int32(2)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPageConvertToGif(name, pageNumber, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPageConvertToGif - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
