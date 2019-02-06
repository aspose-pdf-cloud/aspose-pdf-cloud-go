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
	"os"
	"fmt"
	"testing"
)

func TestPutCreate(t *testing.T) {
	name := "PdfWithBookmarks.pdf"

	args := make(map[string]interface{})

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name) 
	if err != nil {
		t.Error(err)
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutCreate(GetBaseTest().remoteFolder + "/" + name, file, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestPutCreate - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
	
}

func TestGetDownload(t *testing.T) {
	name := "PdfWithBookmarks.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDownload(GetBaseTest().remoteFolder + "/" + name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDownload - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
	
}

func TestPostMoveFile(t *testing.T) {
	name := "4pages.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	src := GetBaseTest().remoteFolder + "/" + name
	dest := GetBaseTest().remoteFolder + "/4pages_renamed.pdf"

	response, httpResponse, err := GetBaseTest().PdfAPI.PostMoveFile(src, dest, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestPostMoveFile - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteFile(t *testing.T) {
	name := "4pages.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	path := GetBaseTest().remoteFolder + "/" + name

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteFile(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tDeleteFile - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetListFiles(t *testing.T) {

	args := map[string]interface{} {
		"path":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetListFiles(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestGetListFiles - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutCreateFolder(t *testing.T) {

	args := make(map[string]interface{})

	path := GetBaseTest().remoteFolder + "/testFolder"

	response, httpResponse, err := GetBaseTest().PdfAPI.PutCreateFolder(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tPutCreateFolder - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostMoveFolder(t *testing.T) {

	args := make(map[string]interface{})

	src := GetBaseTest().remoteFolder + "/testFolder"
	_, _, err := GetBaseTest().PdfAPI.PutCreateFolder(src, args)
	if err != nil {
		t.Error(err)
	}

	dest :=  GetBaseTest().remoteFolder + "/testFolderRenamed"

	response, httpResponse, err := GetBaseTest().PdfAPI.PostMoveFolder(src, dest, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tPostMoveFolder - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteFolder(t *testing.T) {

	args := make(map[string]interface{})

	path := GetBaseTest().remoteFolder + "/testFolder"
	_, _, err := GetBaseTest().PdfAPI.PutCreateFolder(path, args)
	if err != nil {
		t.Error(err)
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteFolder(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tDeleteFolder - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetIsStorageExist(t *testing.T) {

	name := "PDF-CI"

	response, httpResponse, err := GetBaseTest().PdfAPI.GetIsStorageExist(name)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tGetIsStorageExist - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetIsExist(t *testing.T) {
	name := "4pages.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	path := GetBaseTest().remoteFolder + "/" + name

	response, httpResponse, err := GetBaseTest().PdfAPI.GetIsExist(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tGetIsExist - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetDiscUsage(t *testing.T) {
	response, httpResponse, err := GetBaseTest().PdfAPI.GetDiscUsage(nil)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tGetDiscUsage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetListFileVersions(t *testing.T) {
	name := "4pages.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	path := GetBaseTest().remoteFolder + "/" + name

	response, httpResponse, err := GetBaseTest().PdfAPI.GetListFileVersions(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tGetListFileVersions - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}