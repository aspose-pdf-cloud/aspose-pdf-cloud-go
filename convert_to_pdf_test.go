 /**
 *
 *   Copyright (c) 2018 Aspose.PDF Cloud
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

// from EPUB
func TestGetEpubInStorageToPdf(t *testing.T) {

	name := "4pages.epub"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetEpubInStorageToPdf(srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetEpubInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutEpubInStorageToPdf(t *testing.T) {

	name := "4pages.epub"	
	resFileName := "fromEpub.pdf"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutEpubInStorageToPdf(resFileName, srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutEpubInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from WEB
func TestGetWebInStorageToPdf(t *testing.T) {

	sourceURL := "http://google.com"
	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetWebInStorageToPdf(sourceURL, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetWebInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutWebInStorageToPdf(t *testing.T) {

	sourceURL := "http://google.com"
    resFileName := "fromWeb.pdf"

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutWebInStorageToPdf(resFileName, sourceURL, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutWebInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from LaTeX
func TestGetLaTeXInStorageToPdf(t *testing.T) {

	name := "TexExample.tex"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetLaTeXInStorageToPdf(srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetLaTeXInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutLaTeXInStorageToPdf(t *testing.T) {

	name := "TexExample.tex"	
	resFileName := "fromLaTeX.pdf"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutLaTeXInStorageToPdf(resFileName, srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutLaTeXInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from MHT
func TestGetMhtInStorageToPdf(t *testing.T) {

	name := "MhtExample.mht"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetMhtInStorageToPdf(srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetMhtInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutMhtInStorageToPdf(t *testing.T) {

	name := "MhtExample.mht"	
	resFileName := "fromMht.pdf"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutMhtInStorageToPdf(resFileName, srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutMhtInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from HTML
func TestGetHtmlInStorageToPdf(t *testing.T) {

	name := "HtmlWithImage.zip"
	htmlFileName := "HtmlWithImage.html"
	height := float64(650)
    width := float64(250)
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"height": height,
		"width": width,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetHtmlInStorageToPdf(srcPath, htmlFileName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetHtmlInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutHtmlInStorageToPdf(t *testing.T) {

	name := "HtmlWithImage.zip"
	htmlFileName := "HtmlWithImage.html"
	resFileName := "fromHtml.pdf"
	height := float64(650)
    width := float64(250)
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"height": height,
		"width": width,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutHtmlInStorageToPdf(resFileName, srcPath, htmlFileName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutHtmlInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from XLS FO
func TestGetXslFoInStorageToPdf(t *testing.T) {

	name := "XslfoExample.xslfo"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetXslFoInStorageToPdf(srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetXslFoInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutXslFoInStorageToPdf(t *testing.T) {

	name := "XslfoExample.xslfo"	
	resFileName := "fromXlsFo.pdf"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutXslFoInStorageToPdf(resFileName, srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutXslFoInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from XPS
func TestGetXpsInStorageToPdf(t *testing.T) {

	name := "Simple.xps"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetXpsInStorageToPdf(srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetXpsInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutXpsInStorageToPdf(t *testing.T) {

	name := "Simple.xps"	
	resFileName := "fromXps.pdf"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutXpsInStorageToPdf(resFileName, srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutXpsInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from SVG
func TestGetSvgInStorageToPdf(t *testing.T) {

	name := "Simple.svg"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetSvgInStorageToPdf(srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetSvgInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutSvgInStorageToPdf(t *testing.T) {

	name := "Simple.svg"	
	resFileName := "fromSvg.pdf"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutSvgInStorageToPdf(resFileName, srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutSvgInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from PCL
func TestGetPclInStorageToPdf(t *testing.T) {

	name := "template.pcl"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPclInStorageToPdf(srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPclInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPclInStorageToPdf(t *testing.T) {

	name := "template.pcl"
	resFileName := "fromPcl.pdf"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPclInStorageToPdf(resFileName, srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPclInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from XML
func TestGetXmlInStorageToPdf(t *testing.T) {

	name := "template.xml"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetXmlInStorageToPdf(srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetXmlInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutXmlInStorageToPdf(t *testing.T) {

	name := "template.xml"
	resFileName := "fromXml.pdf"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutXmlInStorageToPdf(resFileName, srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutXmlInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// from PS
func TestGetPsInStorageToPdf(t *testing.T) {

	name := "Typography.PS"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := make(map[string]interface{})

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPsInStorageToPdf(srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPsInStorageToPdf - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPsInStorageToPdf(t *testing.T) {

	name := "Typography.PS"
	resFileName := "fromPs.pdf"
	srcPath := GetBaseTest().remoteFolder + "/" + name

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPsInStorageToPdf(resFileName, srcPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPsInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutImageInStorageToPdf(t *testing.T) {

	images := []string{"33539.jpg", "44781.jpg"}
    resFileName := "fromImages.pdf"
	imageTemplateList := []ImageTemplate {}

	for _, fileName := range images {
		imageTemplate := ImageTemplate {
			ImagePath: GetBaseTest().remoteFolder + "/" + fileName,
			ImageSrcType: ImageSrcTypeCommon,
		}
		imageTemplateList = append(imageTemplateList, imageTemplate)

		if err := GetBaseTest().UploadFile(fileName); err != nil {
			t.Error(err)
		}
	}

	imageTemplatesRequest := ImageTemplatesRequest {
		IsOCR: true,
		OCRLangs: "eng",
		ImagesList: imageTemplateList,
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImageInStorageToPdf(resFileName, imageTemplatesRequest, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImageInStorageToPdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
