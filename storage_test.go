 /**
 *
 *   Copyright (c) 2020 Aspose.PDF Cloud
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

func TestUploadFile(t *testing.T) {
	name := "PdfWithBookmarks.pdf"

	args := make(map[string]interface{})

	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name) 
	if err != nil {
		t.Error(err)
	}

	_, httpResponse, err := GetBaseTest().PdfAPI.UploadFile(GetBaseTest().remoteFolder + "/" + name, file, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestUploadFile - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
	
}

func TestDownloadFile(t *testing.T) {
	name := "PdfWithBookmarks.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DownloadFile(GetBaseTest().remoteFolder + "/" + name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestDownloadFile - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
	
}

func TestMoveFile(t *testing.T) {
	name := "4pages.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	src := GetBaseTest().remoteFolder + "/" + name
	dest := GetBaseTest().remoteFolder + "/4pages_renamed.pdf"

	httpResponse, err := GetBaseTest().PdfAPI.MoveFile(src, dest, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestMoveFile - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestDeleteFile(t *testing.T) {
	name := "4pages.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	path := GetBaseTest().remoteFolder + "/" + name

	httpResponse, err := GetBaseTest().PdfAPI.DeleteFile(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteFile - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestGetFilesList(t *testing.T) {

	path := GetBaseTest().remoteFolder
	args := make(map[string]interface{})

	_, httpResponse, err := GetBaseTest().PdfAPI.GetFilesList(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestGetFilesList - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestPutCreateFolder(t *testing.T) {

	args := make(map[string]interface{})

	path := GetBaseTest().remoteFolder + "/testFolder"

	httpResponse, err := GetBaseTest().PdfAPI.CreateFolder(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tPutCreateFolder - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestMoveFolder(t *testing.T) {

	args := make(map[string]interface{})

	src := GetBaseTest().remoteFolder + "/testFolder"
	_, err := GetBaseTest().PdfAPI.CreateFolder(src, args)
	if err != nil {
		t.Error(err)
	}

	dest :=  GetBaseTest().remoteFolder + "/testFolderRenamed"

	httpResponse, err := GetBaseTest().PdfAPI.MoveFolder(src, dest, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tPostMoveFolder - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestDeleteFolder(t *testing.T) {

	args := make(map[string]interface{})

	path := GetBaseTest().remoteFolder + "/testFolder"
	_, err := GetBaseTest().PdfAPI.CreateFolder(path, args)
	if err != nil {
		t.Error(err)
	}

	httpResponse, err := GetBaseTest().PdfAPI.DeleteFolder(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tDeleteFolder - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestStorageExists(t *testing.T) {

	name := "PDF-CI"

	_, httpResponse, err := GetBaseTest().PdfAPI.StorageExists(name)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestStorageExists - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestObjectExists(t *testing.T) {
	name := "4pages.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	path := GetBaseTest().remoteFolder + "/" + name

	_, httpResponse, err := GetBaseTest().PdfAPI.ObjectExists(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestObjectExists - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestGetDiscUsage(t *testing.T) {
	_, httpResponse, err := GetBaseTest().PdfAPI.GetDiscUsage(nil)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tGetDiscUsage - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestGetFileVersions(t *testing.T) {
	name := "4pages.pdf"

	args := make(map[string]interface{})

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	path := GetBaseTest().remoteFolder + "/" + name

	_, httpResponse, err := GetBaseTest().PdfAPI.GetFileVersions(path, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		 t.Fail()
	} else {
		fmt.Printf("%d\tTestGetFileVersions - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}