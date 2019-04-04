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
 
func TestGetDocumentTables(t *testing.T) {

	name := "PdfWithTable.pdf"
	
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentTables(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentTables - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteDocumentTables(t *testing.T) {

	name := "PdfWithTable.pdf"
	
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteDocumentTables(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteDocumentTables - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageTables(t *testing.T) {

	name := "PdfWithTable.pdf"
	var pageNumber int32 = 1
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageTables(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageTables - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeletePageTables(t *testing.T) {

	name := "PdfWithTable.pdf"
	var pageNumber int32 = 1
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeletePageTables(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeletePageTables - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetTable(t *testing.T) {
 
	name := "PdfWithTable.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	responseTables, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentTables(name, args)
	if err != nil {
		t.Error(err)
	}
	stampID := responseTables.Tables.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.GetTable(name, stampID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			t.Fail()
	} else {
		fmt.Printf("%d\tTestGetTable - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}


func TestDeleteTable(t *testing.T) {
 
	name := "PdfWithTable.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	responseTables, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentTables(name, args)
	if err != nil {
		t.Error(err)
	}
	stampID := responseTables.Tables.List[0].Id

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteTable(name, stampID, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteTable - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
