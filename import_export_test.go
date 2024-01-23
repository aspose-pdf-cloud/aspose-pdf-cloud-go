/**
 *
 * Copyright (c) 2024 Aspose.PDF Cloud
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
	"os"
	"testing"
)

func TestGetExportFieldsFromPdfToXmlInStorage(t *testing.T) {
	name := "FormData.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetExportFieldsFromPdfToXmlInStorage(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetExportFieldsFromPdfToXmlInStorage - %d\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestGetExportFieldsFromPdfToFdfInStorage(t *testing.T) {
	name := "FormData.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetExportFieldsFromPdfToFdfInStorage(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetExportFieldsFromPdfToFdfInStorage - %d\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestGetExportFieldsFromPdfToXfdfInStorage(t *testing.T) {
	name := "FormData.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetExportFieldsFromPdfToXfdfInStorage(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetExportFieldsFromPdfToXfdfInStorage - %d\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutExportFieldsFromPdfToXmlInStorage(t *testing.T) {
	name := "FormData.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	outPath := GetBaseTest().remoteFolder + "/exportData.xml"

	response, httpResponse, err := GetBaseTest().PdfAPI.PutExportFieldsFromPdfToXmlInStorage(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetExportFieldsFromPdfToXmlInStorage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutExportFieldsFromPdfToFdfInStorage(t *testing.T) {
	name := "FormData.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	outPath := GetBaseTest().remoteFolder + "/exportData.fdf"

	response, httpResponse, err := GetBaseTest().PdfAPI.PutExportFieldsFromPdfToFdfInStorage(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetExportFieldsFromPdfToFdfInStorage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutExportFieldsFromPdfToXfdfInStorage(t *testing.T) {
	name := "FormData.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	outPath := GetBaseTest().remoteFolder + "/exportData.xfdf"

	response, httpResponse, err := GetBaseTest().PdfAPI.PutExportFieldsFromPdfToXfdfInStorage(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetExportFieldsFromPdfToXfdfInStorage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetImportFieldsFromFdfInStorage(t *testing.T) {
	name := "FormData.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	dataFile := "FormData.fdf"
	if err := GetBaseTest().UploadFile(dataFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	dataPath := GetBaseTest().remoteFolder + "/" + dataFile

	response, httpResponse, err := GetBaseTest().PdfAPI.GetImportFieldsFromFdfInStorage(name, dataPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetImportFieldsFromFdfInStorage - %d\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestGetImportFieldsFromXfdfInStorage(t *testing.T) {
	name := "FormDataXfdf_in.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	dataFile := "FormDataXfdf_in.xfdf"
	if err := GetBaseTest().UploadFile(dataFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	dataPath := GetBaseTest().remoteFolder + "/" + dataFile

	response, httpResponse, err := GetBaseTest().PdfAPI.GetImportFieldsFromXfdfInStorage(name, dataPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetImportFieldsFromXfdfInStorage - %d\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestGetImportFieldsFromXmlInStorage(t *testing.T) {
	name := "FormDataXfa_in.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	dataFile := "FormDataXfa_in.xml"
	if err := GetBaseTest().UploadFile(dataFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	dataPath := GetBaseTest().remoteFolder + "/" + dataFile

	response, httpResponse, err := GetBaseTest().PdfAPI.GetImportFieldsFromXmlInStorage(name, dataPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetImportFieldsFromXmlInStorage - %d\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutImportFieldsFromFdfInStorage(t *testing.T) {
	name := "FormData.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	dataFile := "FormData.fdf"
	if err := GetBaseTest().UploadFile(dataFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	dataPath := GetBaseTest().remoteFolder + "/" + dataFile

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImportFieldsFromFdfInStorage(name, dataPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImportFieldsFromFdfInStorage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutImportFieldsFromXfdfInStorage(t *testing.T) {
	name := "FormDataXfdf_in.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	dataFile := "FormDataXfdf_in.xfdf"
	if err := GetBaseTest().UploadFile(dataFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	dataPath := GetBaseTest().remoteFolder + "/" + dataFile

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImportFieldsFromXfdfInStorage(name, dataPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImportFieldsFromXfdfInStorage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutImportFieldsFromXmlInStorage(t *testing.T) {
	name := "FormDataXfa_in.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	dataFile := "FormDataXfa_in.xml"
	if err := GetBaseTest().UploadFile(dataFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	dataPath := GetBaseTest().remoteFolder + "/" + dataFile

	response, httpResponse, err := GetBaseTest().PdfAPI.PutImportFieldsFromXmlInStorage(name, dataPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutImportFieldsFromXmlInStorage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostImportFieldsFromFdf(t *testing.T) {
	name := "FormData.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	dataFile := "FormData.fdf"

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + dataFile)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder":  GetBaseTest().remoteFolder,
		"fdfData": file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostImportFieldsFromFdf(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostImportFieldsFromFdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostImportFieldsFromXfdf(t *testing.T) {
	name := "FormDataXfdf_in.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	dataFile := "FormDataXfdf_in.xfdf"

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + dataFile)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder":   GetBaseTest().remoteFolder,
		"xfdfData": file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostImportFieldsFromXfdf(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostImportFieldsFromXfdf - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostImportFieldsFromXml(t *testing.T) {
	name := "FormDataXfa_in.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	dataFile := "FormDataXfa_in.xml"

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + dataFile)
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder":  GetBaseTest().remoteFolder,
		"xmlData": file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostImportFieldsFromXml(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostImportFieldsFromXml - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
