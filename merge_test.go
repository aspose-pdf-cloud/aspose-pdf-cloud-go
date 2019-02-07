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

 func TestPutMergeDocuments(t *testing.T) {

	names := []string { "4pages.pdf", "PdfWithImages2.pdf", "marketing.pdf" }
	resultName := "MergingResult.pdf"

	mergeDocuments := MergeDocuments{
		List: []string{},
	}

	for _, name := range names {
		if err := GetBaseTest().UploadFile(name); err != nil {
			t.Error(err)
		}
		mergeDocuments.List = append(mergeDocuments.List, GetBaseTest().remoteFolder + "/" + name)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"mergeDocuments": mergeDocuments,
	}

	response, httpResponse, err := GetBaseTest().PdfAPI.PutMergeDocuments(resultName, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutMergeDocuments - %db\n", GetBaseTest().GetTestNumber(), len(response))
	}
}