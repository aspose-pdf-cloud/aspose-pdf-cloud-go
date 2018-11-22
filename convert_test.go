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
	"os"
	"fmt"
	"testing"
)

// To DOC
func TestGetPdfInStorageToDoc(t *testing.T) {

	name := "4pages.pdf"	

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,

	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToDoc(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToDoc - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToDoc(t *testing.T) {

	name := "4pages.pdf"	

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.doc";
    outPath := GetBaseTest().remoteFolder + "/" + resFileName;

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToDoc(name, outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToDoc - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToDoc(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.doc";
    outPath := GetBaseTest().remoteFolder + "/" + resFileName;

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name) 
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"file": file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToDoc(outPath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToDoc - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// To PDFA
func TestGetPdfInStorageToPdfA(t *testing.T) {

	name := "4pages.pdf"	

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,

	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPdfInStorageToPdfA(name, string(PdfATypePDFA1A), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPdfInStorageToPdfA - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}

func TestPutPdfInStorageToPdfA(t *testing.T) {

	name := "4pages.pdf"	

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	resFileName := "result.pdf";
    outPath := GetBaseTest().remoteFolder + "/" + resFileName;

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInStorageToPdfA(name, outPath, string(PdfATypePDFA1A), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInStorageToPdfA - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutPdfInRequestToPdfA(t *testing.T) {
	name := "4pages.pdf"
	resFileName := "result.pdf";
    outPath := GetBaseTest().remoteFolder + "/" + resFileName;

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name) 
	if err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"file": file,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutPdfInRequestToPdfA(outPath, string(PdfATypePDFA1A), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutPdfInRequestToPdfA - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}