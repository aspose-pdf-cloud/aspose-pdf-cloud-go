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
	"encoding/json"
	"os"
	"path/filepath"
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
	file, err := os.Open(filepath.Join(bt.localTestDataFolder, name))
	if err != nil {
		return err
	}
	_, _, err = GetBaseTest().PdfAPI.UploadFile(filepath.Join(GetBaseTest().remoteFolder, name), file, args)
	return err
}

func (bt *BaseTest) GetTestNumber() int {
	bt.TestNumber++
	return bt.TestNumber
}

func getServercredsJson() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(wd)
	for !(dir[len(dir)-1] == filepath.Separator || dir == ".") {
		servercreds_json := filepath.Join(dir, "Settings", "servercreds.json")
		if _, err := os.Stat(servercreds_json); err == nil {
			// fmt.Println(`Settings\servercreds.json found: ` + servercreds_json)
			return servercreds_json
		}
		dir = filepath.Dir(dir)
	}
	panic(`Settings\servercreds.json not found`)
}

type Creds struct {
	AppSID     string `json:"AppSID"`
	AppKey     string `json:"AppKey"`
	ProductUri string `json:"ProductUri"`
}

func getCreds() Creds {
	bbCreds, err := os.ReadFile(getServercredsJson())
	if err != nil {
		panic(err)
	}
	var creds Creds
	if err := json.Unmarshal(bbCreds, &creds); err != nil {
		panic(err)
	}
	if len(creds.AppSID) == 0 {
		panic("no AppSID")
	}
	if len(creds.AppKey) == 0 {
		panic("no AppKey")
	}
	if len(creds.ProductUri) == 0 {
		panic("no ProductUri")
	}
	return creds
}

func NewBaseTest() *BaseTest {
	creds := getCreds()
	bt := BaseTest{
		remoteFolder:        "TempPdfCloud",
		localTestDataFolder: "test_data",
		TestNumber:          0,
		PdfAPI:              NewPdfApiService(creds.AppSID, creds.AppKey, creds.ProductUri),
	}
	return &bt
}

func GetBaseTest() *BaseTest {
	if BaseTestInstance == nil {
		BaseTestInstance = NewBaseTest()
	}
	return BaseTestInstance
}
