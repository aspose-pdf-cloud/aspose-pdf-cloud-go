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

func TestGetField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"
    fieldName := "textField"
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetField(name, fieldName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetFields(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostCreateField(t *testing.T) {

	name := "Hello world.pdf"
	pageNumber := int32(1)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	field := Field {
		Name: "checkboxfield",
		Type_: FieldTypeBoolean,
		Values: []string{"1"},
		Rect: &RectanglePdf {
			LLX: float64(50),
			LLY: float64(200),
			URX: float64(200),
			URY: float64(400),
		},
	}
			
	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"field": field,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostCreateField(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostCreateField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutUpdateField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"
	fieldName := "textField"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	field := Field {
		Name: fieldName,
		Type_: FieldTypeText,
		Values: []string{"Text field updated value."},
		Rect: nil,
		Links: nil,
	}
			
	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"field": field,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutUpdateField(name, field.Name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutUpdateField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutUpdateFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"
	fieldName := "textField"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	field := Field {
		Name: fieldName,
		Type_: FieldTypeText,
		Values: []string{"Text field updated value."},
		Rect: nil,
		Links: nil,
	}

	fields := Fields {
		List: []Field{field},
	}
	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"fields": fields,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutUpdateFields(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutUpdateFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestDeleteField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"
	fieldName := "textField"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.DeleteField(name, fieldName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutFieldsFlatten(t *testing.T) {

	name := "PdfWithAcroForm.pdf"
	
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutFieldsFlatten(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutFieldsFlatten - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
