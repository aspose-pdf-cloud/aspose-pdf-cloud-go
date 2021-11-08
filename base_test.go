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
	"os"
)

var BaseTestInstance *BaseTest

type BaseTest struct {
	remoteFolder        string
	localTestDataFolder string
	PdfAPI              *PdfApiService
	TestNumber          int
}

func (bt *BaseTest) UploadFile(name string) (err error) {
	args := make(map[string]interface{})

	file, err := os.Open(bt.localTestDataFolder + "/" + name)
	if err != nil {
		return err
	}

	_, _, err = GetBaseTest().PdfAPI.UploadFile(GetBaseTest().remoteFolder+"/"+name, file, args)
	return err
}

func (bt *BaseTest) GetTestNumber() int {
	bt.TestNumber++
	return bt.TestNumber
}

func NewBaseTest() *BaseTest {
	bt := &BaseTest{
		remoteFolder:        "TempPdfCloud",
		localTestDataFolder: "test_data",
		TestNumber:          0,
		// Get App key and App SID from https://aspose.cloud
		PdfAPI: NewPdfApiService("AppSID", "AppKey", "https://api.aspose.cloud/v3.0"),
	}
	return bt
}

func GetBaseTest() *BaseTest {
	if BaseTestInstance == nil {
		BaseTestInstance = NewBaseTest()
	}
	return BaseTestInstance
}
