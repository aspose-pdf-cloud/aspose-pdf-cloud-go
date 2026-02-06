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

func TestGetField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"
	fieldName := "textField"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
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

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
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

	field := Field{
		Name:   "checkboxfield",
		Type_:  FieldTypeBoolean,
		Values: []string{"1"},
		Rect: &Rectangle{
			LLX: float64(50),
			LLY: float64(200),
			URX: float64(200),
			URY: float64(400),
		},
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostCreateField(name, pageNumber, field, args)
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

	field := Field{
		Name:   fieldName,
		Type_:  FieldTypeText,
		Values: []string{"Text field updated value."},
		Rect:   nil,
		Links:  nil,
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutUpdateField(name, field.Name, field, args)
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

	field := Field{
		Name:   fieldName,
		Type_:  FieldTypeText,
		Values: []string{"Text field updated value."},
		Rect:   nil,
		Links:  nil,
	}

	fields := Fields{
		List: []Field{field},
	}
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutUpdateFields(name, fields, args)
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

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
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

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
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

func TestPostFlattenDocument(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"updateAppearances": true,
		"hideButtons":       true,
		"folder":            GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostFlattenDocument(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostFlattenDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetDocumentSignatureFields(t *testing.T) {

	name := "adbe.x509.rsa_sha1.valid.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentSignatureFields(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentSignatureFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageSignatureFields(t *testing.T) {

	name := "adbe.x509.rsa_sha1.valid.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	pageNumber := int32(1)

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageSignatureFields(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageSignatureFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetSignatureField(t *testing.T) {

	name := "adbe.x509.rsa_sha1.valid.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	signatureName := "Signature1"

	response, httpResponse, err := GetBaseTest().PdfAPI.GetSignatureField(name, signatureName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetSignatureField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostSignatureField(t *testing.T) {

	name := "4pages.pdf"
	signatureName := "33226.p12"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	if err := GetBaseTest().UploadFile(signatureName); err != nil {
		t.Error(err)
	}

	signature := Signature{
		Authority:      "Sergey Smal",
		Contact:        "test@mail.ru",
		Date:           "08/01/2012 12:15:00.000 PM",
		FormFieldName:  "Signature1",
		Location:       "Ukraine",
		Password:       "sIikZSmz",
		Rectangle:      &Rectangle{LLX: 100, LLY: 100, URX: 0, URY: 0},
		SignaturePath:  GetBaseTest().remoteFolder + "/" + signatureName,
		SignatureType:  SignatureTypePKCS7,
		Visible:        true,
		ShowProperties: false,
	}

	field := SignatureField{
		PageIndex:   1,
		Signature:   &signature,
		Rect:        &Rectangle{LLX: 100, LLY: 100, URX: 0, URY: 0},
		PartialName: "sign1",
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostSignatureField(name, field, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostSignatureField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutSignatureField(t *testing.T) {

	name := "4pages.pdf"
	signatureName := "33226.p12"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	if err := GetBaseTest().UploadFile(signatureName); err != nil {
		t.Error(err)
	}

	signature := Signature{
		Authority:      "Sergey Smal",
		Contact:        "test@mail.ru",
		Date:           "08/01/2012 12:15:00.000 PM",
		FormFieldName:  "Signature1",
		Location:       "Ukraine",
		Password:       "sIikZSmz",
		Rectangle:      &Rectangle{LLX: 100, LLY: 100, URX: 0, URY: 0},
		SignaturePath:  GetBaseTest().remoteFolder + "/" + signatureName,
		SignatureType:  SignatureTypePKCS7,
		Visible:        true,
		ShowProperties: false,
	}

	field := SignatureField{
		PageIndex:   1,
		Signature:   &signature,
		Rect:        &Rectangle{LLX: 100, LLY: 100, URX: 0, URY: 0},
		PartialName: "sign1",
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutSignatureField(name, "Signature1", field, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutSignatureField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetDocumentTextBoxFields(t *testing.T) {

	name := "FormDataTextBox.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentTextBoxFields(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentTextBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageTextBoxFields(t *testing.T) {

	name := "FormDataTextBox.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	pageNumber := int32(1)
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageTextBoxFields(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageTextBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetTextBoxField(t *testing.T) {

	name := "FormDataTextBox.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	fieldName := "Petitioner"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetTextBoxField(name, fieldName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetTextBoxField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostTextBoxFields(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	textBox := TextBoxField{
		PageIndex:   1,
		IsGroup:     false,
		Color:       &Color{A: 255, R: 255, G: 0, B: 0},
		Multiline:   true,
		MaxLen:      100,
		Rect:        &Rectangle{LLX: 100, LLY: 100, URX: 500, URY: 200},
		Value:       "Page 1\nValue",
		PartialName: "testField",
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostTextBoxFields(name, []TextBoxField{textBox}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostTextBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutTextBoxField(t *testing.T) {

	name := "FormDataTextBox.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	fieldName := "Petitioner"

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	textBox := TextBoxField{
		PageIndex:   1,
		IsGroup:     false,
		Color:       &Color{A: 255, R: 255, G: 0, B: 0},
		Multiline:   true,
		MaxLen:      100,
		Rect:        &Rectangle{LLX: 100, LLY: 100, URX: 500, URY: 200},
		Value:       "Page 1\nValue",
		PartialName: "testField",
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutTextBoxField(name, fieldName, textBox, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutTextBoxField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// CheckBoxField

func TestGetDocumentCheckBoxFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentCheckBoxFields(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentCheckBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageCheckBoxFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	pageNumber := int32(1)
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageCheckBoxFields(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageCheckBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetCheckBoxField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	fieldName := "checkboxField"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetCheckBoxField(name, fieldName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetCheckBoxField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostCheckBoxFields(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	field := CheckBoxField{
		PageIndex:   1,
		IsGroup:     false,
		Color:       &Color{A: 255, R: 255, G: 0, B: 0},
		Checked:     true,
		Rect:        &Rectangle{LLX: 100, LLY: 100, URX: 500, URY: 200},
		ExportValue: "true",
		PartialName: "testField",
		Style:       BoxStyleCross,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostCheckBoxFields(name, []CheckBoxField{field}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostCheckBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutCheckBoxField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	fieldName := "checkboxField"

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	field := CheckBoxField{
		PageIndex:   1,
		IsGroup:     false,
		Color:       &Color{A: 255, R: 255, G: 0, B: 0},
		Checked:     true,
		Rect:        &Rectangle{LLX: 100, LLY: 100, URX: 500, URY: 200},
		ExportValue: "true",
		PartialName: "testField",
		Style:       BoxStyleCross,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutCheckBoxField(name, fieldName, field, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutCheckBoxField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// RadioButtonField

func TestGetDocumentRadioButtonFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentRadioButtonFields(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentRadioButtonFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageRadioButtonFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	pageNumber := int32(1)
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageRadioButtonFields(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageRadioButtonFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetRadioButtonField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	fieldName := "radiobuttonField"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetRadioButtonField(name, fieldName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetRadioButtonField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostRadioButtonFields(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	option1 := RadioButtonOptionField{
		PageIndex:  1,
		IsGroup:    false,
		OptionName: "1",
		Rect:       &Rectangle{LLX: 100, LLY: 130, URX: 160, URY: 140},
		Style:      BoxStyleCross,
	}

	option2 := RadioButtonOptionField{
		PageIndex:  1,
		IsGroup:    false,
		OptionName: "2",
		Rect:       &Rectangle{LLX: 150, LLY: 120, URX: 160, URY: 130},
		Style:      BoxStyleCross,
	}

	field := RadioButtonField{
		PageIndex:               1,
		IsGroup:                 false,
		Color:                   &Color{A: 255, R: 255, G: 0, B: 0},
		Selected:                1,
		Rect:                    &Rectangle{LLX: 100, LLY: 100, URX: 160, URY: 140},
		PartialName:             "testField",
		Style:                   BoxStyleCross,
		RadioButtonOptionsField: []RadioButtonOptionField{option1, option2},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostRadioButtonFields(name, []RadioButtonField{field}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostRadioButtonFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutRadioButtonField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	fieldName := "radiobuttonField"

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	option1 := RadioButtonOptionField{
		PageIndex:  1,
		IsGroup:    false,
		OptionName: "1",
		Rect:       &Rectangle{LLX: 100, LLY: 130, URX: 160, URY: 140},
		Style:      BoxStyleCross,
	}

	option2 := RadioButtonOptionField{
		PageIndex:  1,
		IsGroup:    false,
		OptionName: "2",
		Rect:       &Rectangle{LLX: 150, LLY: 120, URX: 160, URY: 130},
		Style:      BoxStyleCross,
	}

	field := RadioButtonField{
		PageIndex:               1,
		IsGroup:                 false,
		Color:                   &Color{A: 255, R: 255, G: 0, B: 0},
		Selected:                1,
		Rect:                    &Rectangle{LLX: 100, LLY: 100, URX: 160, URY: 140},
		PartialName:             "testField",
		Style:                   BoxStyleCross,
		RadioButtonOptionsField: []RadioButtonOptionField{option1, option2},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutRadioButtonField(name, fieldName, field, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutRadioButtonField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// ComboBoxField

func TestGetDocumentComboBoxFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentComboBoxFields(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentComboBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageComboBoxFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	pageNumber := int32(1)
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageComboBoxFields(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageComboBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetComboBoxField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	fieldName := "comboboxField"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetComboBoxField(name, fieldName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetComboBoxField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostComboBoxFields(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	option1 := Option{
		Name:  "one",
		Value: "one",
	}

	option2 := Option{
		Name:  "two",
		Value: "two",
	}

	field := ComboBoxField{
		PageIndex:   1,
		IsGroup:     false,
		Color:       &Color{A: 255, R: 255, G: 0, B: 0},
		Selected:    1,
		Rect:        &Rectangle{LLX: 100, LLY: 100, URX: 160, URY: 140},
		PartialName: "testField",
		Options:     []Option{option1, option2},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostComboBoxFields(name, []ComboBoxField{field}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostComboBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutComboBoxField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	fieldName := "comboboxField"

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	option1 := Option{
		Name:  "one",
		Value: "one",
	}

	option2 := Option{
		Name:  "two",
		Value: "two",
	}

	field := ComboBoxField{
		PageIndex:   1,
		IsGroup:     false,
		Color:       &Color{A: 255, R: 255, G: 0, B: 0},
		Selected:    1,
		Rect:        &Rectangle{LLX: 100, LLY: 100, URX: 160, URY: 140},
		PartialName: "testField",
		Options:     []Option{option1, option2},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutComboBoxField(name, fieldName, field, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutComboBoxField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// ListBoxField

func TestGetDocumentListBoxFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetDocumentListBoxFields(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDocumentListBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageListBoxFields(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	pageNumber := int32(1)
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageListBoxFields(name, pageNumber, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageListBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetListBoxField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	fieldName := "listboxField"
	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetListBoxField(name, fieldName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetListBoxField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostListBoxFields(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	option1 := Option{
		Name:  "one",
		Value: "one",
	}

	option2 := Option{
		Name:  "two",
		Value: "two",
	}

	option3 := Option{
		Name:  "three",
		Value: "three",
	}

	option4 := Option{
		Name:  "four",
		Value: "four",
	}

	field := ListBoxField{
		PageIndex:     1,
		MultiSelect:   true,
		Color:         &Color{A: 255, R: 255, G: 0, B: 0},
		SelectedItems: []int32{1, 4},
		Rect:          &Rectangle{LLX: 100, LLY: 100, URX: 160, URY: 140},
		PartialName:   "testField",
		Options:       []Option{option1, option2, option3, option4},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostListBoxFields(name, []ListBoxField{field}, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostListBoxFields - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutListBoxField(t *testing.T) {

	name := "PdfWithAcroForm.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	fieldName := "ListBoxField"

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	option1 := Option{
		Name:  "one",
		Value: "one",
	}

	option2 := Option{
		Name:  "two",
		Value: "two",
	}

	option3 := Option{
		Name:  "three",
		Value: "three",
	}

	option4 := Option{
		Name:  "four",
		Value: "four",
	}

	field := ListBoxField{
		PageIndex:     1,
		MultiSelect:   true,
		Color:         &Color{A: 255, R: 255, G: 0, B: 0},
		SelectedItems: []int32{1, 4},
		Rect:          &Rectangle{LLX: 100, LLY: 100, URX: 160, URY: 140},
		PartialName:   "testField",
		Options:       []Option{option1, option2, option3, option4},
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutListBoxField(name, fieldName, field, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutListBoxField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
