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
 
func TestPostDocumentTextReplace(t *testing.T) {

	name := "marketing.pdf"	

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	textReplaceListRequest := TextReplaceListRequest{
		TextReplaces: []TextReplace{
			TextReplace{
				OldValue: "market",
				NewValue: "m_a_r_k_e_t",
				Rect: &Rectangle{ LLX: 100, LLY: 100, URX: 200, URY: 200},
				Regex: false,
			},
		},
		StartIndex: 0,
		CountReplace: 0,
	}


	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostDocumentTextReplace(name, textReplaceListRequest, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostDocumentTextReplace - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPostPageTextReplace(t *testing.T) {

	name := "marketing.pdf"	

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	pageNumber := int32(1)
	textReplaceListRequest := TextReplaceListRequest{
		TextReplaces: []TextReplace{
			TextReplace{
				OldValue: "market",
				NewValue: "m_a_r_k_e_t",
				Rect: &Rectangle{ LLX: 100, LLY: 100, URX: 200, URY: 200},
				Regex: false,
			},
		},
		StartIndex: 0,
		CountReplace: 0,
	}


	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PostPageTextReplace(name, pageNumber, textReplaceListRequest, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPostPageTextReplace - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}