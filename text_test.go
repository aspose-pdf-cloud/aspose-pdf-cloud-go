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
 
func TestGetText(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}

	llx := float64(0)
	lly := float64(0)
	urx := float64(0)
	ury := float64(0)

	response, httpResponse, err := GetBaseTest().PdfAPI.GetText(name, llx, lly, urx, ury, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetText - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestGetPageText(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	pageNumber := int32(1)
	format := []string {"First Page", "Second Page"}
	
	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
		"format": format,
	}

	llx := float64(0)
	lly := float64(0)
	urx := float64(0)
	ury := float64(0)

	response, httpResponse, err := GetBaseTest().PdfAPI.GetPageText(name, pageNumber, llx, lly, urx, ury, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetPageText - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestPutAddText(t *testing.T) {

	name := "4pages.pdf"

	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	pageNumber := int32(1)
	paragraph := Paragraph {
		Rectangle: &Rectangle{ LLX: 100, LLY: 100, URX: 200, URY: 200},
		LeftMargin: 10,
		RightMargin: 10,
		TopMargin: 20,
		BottomMargin: 20,
		HorizontalAlignment: TextHorizontalAlignmentFullJustify,
		LineSpacing: LineSpacingFontSize,
		Rotation: 10,
		SubsequentLinesIndent: 20,
		VerticalAlignment: VerticalAlignmentCenter,
		WrapMode: WrapModeByWords,

		Lines: []TextLine {
			TextLine {
				HorizontalAlignment: TextHorizontalAlignmentRight,
				Segments: []Segment {
					Segment {
						Value: "segment 1",
						TextState: &TextState {
							Font: "Arial",
							FontSize: 10,
							ForegroundColor: &Color {A: 0x00, R: 0x00, G: 0xFF, B: 0x00 },
							BackgroundColor: &Color {A: 0x00, R: 0xFF, G: 0x00, B: 0x00},
							FontStyle: FontStylesBold,
						},
					},
				},
			},
		},
	}
	
	args := map[string]interface{} {
		"folder":  GetBaseTest().remoteFolder,
	}


	response, httpResponse, err := GetBaseTest().PdfAPI.PutAddText(name, pageNumber, paragraph, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPutAddText - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}