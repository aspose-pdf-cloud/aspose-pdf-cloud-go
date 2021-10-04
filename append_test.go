/**
 *
 * Copyright (c) 2021 Aspose.PDF Cloud
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

func TestPostAppendDocument(t *testing.T) {

	name := "PdfWithImages2.pdf"
	appendFile := "4pages.pdf"
	appendFilePath := GetBaseTest().remoteFolder + "/" + appendFile
	startPage := int32(2)
	endPage := int32(4)

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	if err := GetBaseTest().UploadFile(appendFile); err != nil {
		t.Error(err)
	}

	args := map[string]interface{}{
		"folder":    GetBaseTest().remoteFolder,
		"startPage": startPage,
		"endPage":   endPage,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostAppendDocument(name, appendFilePath, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostAppendDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
