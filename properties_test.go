/**
 *
 * Copyright (c) 2026 Aspose.PDF Cloud
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

func TestDeleteProperties(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	propertyNames := []string{"prop1", "prop2"}
	propertyValues := []string{"val1", "val2"}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	_, _, err := GetBaseTest().PdfAPI.PutSetProperty(name, propertyNames[0], propertyValues[0], args)
	if err != nil {
		t.Error(err)
	}
	_, _, err = GetBaseTest().PdfAPI.PutSetProperty(name, propertyNames[1], propertyValues[1], args)
	if err != nil {
		t.Error(err)
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteProperties(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteProperties - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteProperty(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	propertyNames := []string{"prop1", "prop2"}
	propertyValues := []string{"val1", "val2"}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	_, _, err := GetBaseTest().PdfAPI.PutSetProperty(name, propertyNames[0], propertyValues[0], args)
	if err != nil {
		t.Error(err)
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteProperty(name, propertyNames[0], args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteProperty - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetDocumentProperties(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	propertyNames := []string{"prop1", "prop2"}
	propertyValues := []string{"val1", "val2"}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	_, _, err := GetBaseTest().PdfAPI.PutSetProperty(name, propertyNames[0], propertyValues[0], args)
	if err != nil {
		t.Error(err)
	}
	_, _, err = GetBaseTest().PdfAPI.PutSetProperty(name, propertyNames[1], propertyValues[1], args)
	if err != nil {
		t.Error(err)
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentProperties(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentProperties - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetDocumentProperty(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	propertyNames := []string{"prop1", "prop2"}
	propertyValues := []string{"val1", "val2"}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	_, _, err := GetBaseTest().PdfAPI.PutSetProperty(name, propertyNames[0], propertyValues[0], args)
	if err != nil {
		t.Error(err)
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentProperty(name, propertyNames[0], args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentProperty - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutSetProperty(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	propertyNames := []string{"prop1", "prop2"}
	propertyValues := []string{"val1", "val2"}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutSetProperty(name, propertyNames[0], propertyValues[0], args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutSetProperty - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
