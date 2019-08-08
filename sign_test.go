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

func TestPostSignDocument(t *testing.T) {
 
	name := "BlankWithSignature.pdf"
	signatureName := "test1234.pfx"
 
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	if err := GetBaseTest().UploadFile(signatureName); err != nil {
		t.Error(err)
	}

	signature := Signature {
		Authority: "Sergey Smal",
		Contact: "test@mail.ru",
		Date: "08/01/2012 12:15:00.000 PM",
		FormFieldName: "Signature1",
		Location: "Ukraine",
		Password: "test1234",
		Rectangle: &Rectangle{ LLX: 100, LLY: 100, URX: 0, URY: 0},
		SignaturePath: GetBaseTest().remoteFolder + "/" + signatureName,
		SignatureType: SignatureTypePKCS7,
		Visible: true,
		ShowProperties: false, 
		}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostSignDocument(name, signature, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostSignDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostSignPage(t *testing.T) {
 
	name := "BlankWithSignature.pdf"
	signatureName := "test1234.pfx"
	pageNumber := int32(1)
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	if err := GetBaseTest().UploadFile(signatureName); err != nil {
		t.Error(err)
	}

	signature := Signature {
		Authority: "Sergey Smal",
		Contact: "test@mail.ru",
		Date: "08/01/2012 12:15:00.000 PM",
		FormFieldName: "Signature1",
		Location: "Ukraine",
		Password: "test1234",
		Rectangle: &Rectangle{ LLX: 100, LLY: 100, URX: 0, URY: 0},
		SignaturePath: GetBaseTest().remoteFolder + "/" + signatureName,
		SignatureType: SignatureTypePKCS7,
		Visible: true,
		ShowProperties: false, 
		}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostSignPage(name, pageNumber, signature, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostSignPage - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetVerifySignature(t *testing.T) {
 
	name := "BlankWithSignature.pdf"
	signatureName := "test1234.pfx"
 
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	if err := GetBaseTest().UploadFile(signatureName); err != nil {
		t.Error(err)
	}

	signature := Signature {
		Authority: "Sergey Smal",
		Contact: "test@mail.ru",
		Date: "08/01/2012 12:15:00.000 PM",
		FormFieldName: "Signature1",
		Location: "Ukraine",
		Password: "test1234",
		Rectangle: &Rectangle{ LLX: 100, LLY: 100, URX: 0, URY: 0},
		SignaturePath: GetBaseTest().remoteFolder + "/" + signatureName,
		SignatureType: SignatureTypePKCS7,
		Visible: true,
		ShowProperties: false, 
		}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	_, _, err := GetBaseTest().PdfAPI.PostSignDocument(name, signature, args)
	if err != nil {
		t.Error(err)
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetVerifySignature(name, signature.FormFieldName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetVerifySignature - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}