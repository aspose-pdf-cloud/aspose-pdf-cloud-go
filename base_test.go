/**
 *
 * Copyright (c) 2025 Aspose.PDF Cloud
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

type Creds struct {
	ProductUri   string `json:"api_url"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	SelfHost     bool   `json:"self_host"`
}

func getCreds() Creds {
	bbCreds, err := os.ReadFile("settings/credentials.json")
	if err != nil {
		panic(err)
	}
	var creds Creds
	if err := json.Unmarshal(bbCreds, &creds); err != nil {
		panic(err)
	}
	if creds.SelfHost {
		if len(creds.ProductUri) == 0 {
			panic("no ProductUri")
		}
	} else {
		if len(creds.ClientId) == 0 {
			panic("no ClientId")
		}
		if len(creds.ClientSecret) == 0 {
			panic("no ClientSecret")
		}
	}
	return creds
}

func NewBaseTest() *BaseTest {
	creds := getCreds()
	var pdfApiService *PdfApiService
	if creds.SelfHost {
		pdfApiService = NewSelfHostPdfApiService(creds.ProductUri)
	} else {
		pdfApiService = NewPdfApiService(creds.ClientId, creds.ClientSecret, creds.ProductUri)
	}
	baseTest := BaseTest{
		remoteFolder:        "TempPdfCloud",
		localTestDataFolder: "test_data",
		TestNumber:          0,
		PdfAPI:              pdfApiService,
	}
	return &baseTest
}

func GetBaseTest() *BaseTest {
	if BaseTestInstance == nil {
		BaseTestInstance = NewBaseTest()
	}
	return BaseTestInstance
}
