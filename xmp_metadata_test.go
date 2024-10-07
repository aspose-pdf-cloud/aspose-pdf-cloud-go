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
	"testing"
)

func TestGetXmpMetadataJson(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetXmpMetadataJson(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetXmpMetadataJson - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}

	if response.Properties == nil {
		t.Error("XmpMetadata.Properties is null")
	}
	if len(response.Properties) != 9 {
		t.Error("XmpMetadata.Properties != 9")
	}
}

func TestGetXmpMetadataXml(t *testing.T) {

	name := "4pages.pdf"
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder": GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.GetXmpMetadataXml(name, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetXmpMetadataXml - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}

	xml := string(response)
	if !(len(xml) > 0) {
		t.Fail()
	}
}

func TestPostXmpMetadata(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	test_date := "2024-10-27T09:59:52+02:00"

	xmpMetadata := XmpMetadata{
		Properties: []XmpMetadataProperty{
			// Modify Default property without prefix
			{
				Key:   "ModifyDate",
				Value: test_date,
			},
			// Modify Default property with prefix
			{
				Key:   "xmp:CreateDate",
				Value: test_date,
			},
			//Remove Default property
			{
				Key: "CreatorTool",
			},
			//Add Default property
			{
				Key:   "BaseURL",
				Value: "http://www.somename.com/path",
			},
			//Remove User defined property
			{
				Key: "dc:title",
			},
			// Update user defined property
			{
				Key:          "pdf:Producer",
				Value:        "Aspose.PDF Cloud",
				NamespaceUri: "http://ns.adobe.com/pdf/1.3/",
			},
			// Add user defined property
			{
				Key:          "pdf:Prop",
				Value:        "PropValue",
				NamespaceUri: "http://ns.adobe.com/pdf/1.3/",
			},
		},
	}

	{
		args := map[string]interface{}{
			"folder": GetBaseTest().remoteFolder,
		}
		response, httpResponse, err := GetBaseTest().PdfAPI.PostXmpMetadata(name, xmpMetadata, args)
		if err != nil {
			t.Error(err)
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			t.Fail()
		} else {
			fmt.Printf("%d\tTestPostXmpMetadata - %d\n", GetBaseTest().GetTestNumber(), response.Code)
		}
	}

	{
		args := map[string]interface{}{
			"folder": GetBaseTest().remoteFolder,
		}
		response, httpResponse, err := GetBaseTest().PdfAPI.GetXmpMetadataJson(name, args)
		if err != nil {
			t.Error(err)
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			t.Fail()
		} else {
			fmt.Printf("%d\tTestPostXmpMetadata - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
		}

		if response.Properties == nil {
			t.Error("XmpMetadata.Properties is null")
		}
		if len(response.Properties) != 9 {
			t.Error("XmpMetadata.Properties != 9")
		}

	}
}
