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
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestPutEncryptDocument(t *testing.T) {
	name := "4pages.pdf"
	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"file": file,
	}
	outPath := GetBaseTest().remoteFolder + "/" + name
	userPassword := "user $^Password!&"
	ownerPassword := "owner\\//? $12^Password!&"
	response, httpResponse, err := GetBaseTest().PdfAPI.PutEncryptDocument(outPath,
		base64.StdEncoding.EncodeToString([]byte(userPassword)),
		base64.StdEncoding.EncodeToString([]byte(ownerPassword)),
		string(CryptoAlgorithmAESx128),
		args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutEncryptDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostEncryptDocumentInStorage(t *testing.T) {
	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	userPassword := "user $^Password!&"
	ownerPassword := "owner\\//? $12^Password!&"
	response, httpResponse, err := GetBaseTest().PdfAPI.PostEncryptDocumentInStorage(name,
		base64.StdEncoding.EncodeToString([]byte(userPassword)),
		base64.StdEncoding.EncodeToString([]byte(ownerPassword)),
		string(CryptoAlgorithmAESx128),
		args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostEncryptDocumentInStorage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutDecryptDocument(t *testing.T) {
	name := "4pagesEncrypted.pdf"
	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"file": file,
	}
	outPath := GetBaseTest().remoteFolder + "/" + name
	userPassword := "user $^Password!&"
	response, httpResponse, err := GetBaseTest().PdfAPI.PutDecryptDocument(outPath,
		base64.StdEncoding.EncodeToString([]byte(userPassword)),
		args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutDecryptDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostDecryptDocumentInStorage(t *testing.T) {
	name := "4pagesEncrypted.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	userPassword := "user $^Password!&"
	response, httpResponse, err := GetBaseTest().PdfAPI.PostDecryptDocumentInStorage(name,
		base64.StdEncoding.EncodeToString([]byte(userPassword)),
		args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostDecryptDocumentInStorage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutChangePasswordDocument(t *testing.T) {

	name := "4pagesEncrypted.pdf"
	file, err := os.Open(GetBaseTest().localTestDataFolder + "/" + name)
	if err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"file": file,
	}
	outPath := GetBaseTest().remoteFolder + "/" + name
	ownerPassword := "owner\\//? $12^Password!&"
	newUserPassword := "user new\\//? $12^Password!&"
	newOwnerPassword := "owner new\\//? $12^Password!&"
	response, httpResponse, err := GetBaseTest().PdfAPI.PutChangePasswordDocument(outPath,
		base64.StdEncoding.EncodeToString([]byte(ownerPassword)),
		base64.StdEncoding.EncodeToString([]byte(newUserPassword)),
		base64.StdEncoding.EncodeToString([]byte(newOwnerPassword)),
		args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutChangePasswordDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostChangePasswordDocumentInStorage(t *testing.T) {
	name := "4pagesEncrypted.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}
	ownerPassword := "owner\\//? $12^Password!&"
	newUserPassword := "user new\\//? $12^Password!&"
	newOwnerPassword := "owner new\\//? $12^Password!&"
	response, httpResponse, err := GetBaseTest().PdfAPI.PostChangePasswordDocumentInStorage(name,
		base64.StdEncoding.EncodeToString([]byte(ownerPassword)),
		base64.StdEncoding.EncodeToString([]byte(newUserPassword)),
		base64.StdEncoding.EncodeToString([]byte(newOwnerPassword)),
		args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tPostChangePasswordDocumentInStorage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
