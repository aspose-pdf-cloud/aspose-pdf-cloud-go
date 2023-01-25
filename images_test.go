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

func TestGetImage(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetImage(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetImage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetImages(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetImages - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteImage(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteImage(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteImage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutReplaceImage(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	imageFileName := "Koala.jpg"
	if err := GetBaseTest().UploadFile(imageFileName); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder":        GetBaseTest().remoteFolder,
		"imageFilePath": GetBaseTest().remoteFolder + "/" + imageFileName,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutReplaceImage(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutReplaceImage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostInsertImage(t *testing.T) {

	name := "PdfWithImages2.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	imageFileName := "Koala.jpg"
	if err := GetBaseTest().UploadFile(imageFileName); err != nil {
		t.Error(err)
	}

	pageNumber := int32(1)
	llx := float64(10)
	lly := float64(10)
	urx := float64(100)
	ury := float64(100)

	args := map[string]interface{}{
		"folder":        GetBaseTest().remoteFolder,
		"imageFilePath": GetBaseTest().remoteFolder + "/" + imageFileName,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostInsertImage(name, pageNumber, llx, lly, urx, ury, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostInsertImage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutImagesExtractAsJpeg(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	destFolder := "extract_jpg"

	args := map[string]interface{}{
		"folder":     GetBaseTest().remoteFolder,
		"destFolder": GetBaseTest().remoteFolder + "/" + destFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImagesExtractAsJpeg(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImagesExtractAsJpeg - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutImagesExtractAsTiff(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	destFolder := "extract_tiff"

	args := map[string]interface{}{
		"folder":     GetBaseTest().remoteFolder,
		"destFolder": GetBaseTest().remoteFolder + "/" + destFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImagesExtractAsTiff(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImagesExtractAsTiff - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutImagesExtractAsGif(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	destFolder := "extract_gif"

	args := map[string]interface{}{
		"folder":     GetBaseTest().remoteFolder,
		"destFolder": GetBaseTest().remoteFolder + "/" + destFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImagesExtractAsGif(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImagesExtractAsGif - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutImagesExtractAsPng(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	destFolder := "extract_png"

	args := map[string]interface{}{
		"folder":     GetBaseTest().remoteFolder,
		"destFolder": GetBaseTest().remoteFolder + "/" + destFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImagesExtractAsPng(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImagesExtractAsPng - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutImageExtractAsJpeg(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	destFolder := "extract_jpg"

	args := map[string]interface{}{
		"folder":     GetBaseTest().remoteFolder,
		"destFolder": GetBaseTest().remoteFolder + "/" + destFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImageExtractAsJpeg(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImageExtractAsJpeg - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetImageExtractAsJpeg(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetImageExtractAsJpeg(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetImageExtractAsJpeg - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutImageExtractAsTiff(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	destFolder := "extract_tiff"

	args := map[string]interface{}{
		"folder":     GetBaseTest().remoteFolder,
		"destFolder": GetBaseTest().remoteFolder + "/" + destFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImageExtractAsTiff(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImageExtractAsTiff - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetImageExtractAsTiff(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetImageExtractAsTiff(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetImageExtractAsTiff - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutImageExtractAsGif(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	destFolder := "extract_gif"

	args := map[string]interface{}{
		"folder":     GetBaseTest().remoteFolder,
		"destFolder": GetBaseTest().remoteFolder + "/" + destFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImageExtractAsGif(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImageExtractAsGif - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetImageExtractAsGif(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetImageExtractAsGif(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetImageExtractAsGif - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutImageExtractAsPng(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	destFolder := "extract_png"

	args := map[string]interface{}{
		"folder":     GetBaseTest().remoteFolder,
		"destFolder": GetBaseTest().remoteFolder + "/" + destFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImageExtractAsPng(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImageExtractAsPng - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetImageExtractAsPng(t *testing.T) {

	name := "PdfWithImages2.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	responseImages, _, err := GetBaseTest().PdfAPI.GetImages(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	}
	imageID := responseImages.Images.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetImageExtractAsPng(name, imageID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetImageExtractAsPng - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}
