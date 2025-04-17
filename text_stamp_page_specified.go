/*
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

// Represents Pdf stamps.
type TextStampPageSpecified struct {
	// Link to the document.
	Links []Link `json:"Links,omitempty"`
	// Sets or gets a bool value that indicates the content is stamped as background. If the value is true, the stamp content is layed at the bottom. By defalt, the value is false, the stamp content is layed at the top.
	Background bool `json:"Background,omitempty"`
	// Gets or sets Horizontal alignment of stamp on the page. 
	HorizontalAlignment HorizontalAlignment `json:"HorizontalAlignment,omitempty"`
	// Gets or sets a value to indicate the stamp opacity. The value is from 0.0 to 1.0. By default the value is 1.0.
	Opacity float64 `json:"Opacity,omitempty"`
	// Sets or gets the rotation of stamp content according Rotation values. Note. This property is for set angles which are multiples of 90 degrees (0, 90, 180, 270 degrees). To set arbitrary angle use RotateAngle property.  If angle set by ArbitraryAngle is not multiple of 90 then Rotate property returns Rotation.None.
	Rotate Rotation `json:"Rotate,omitempty"`
	// Gets or sets rotate angle of stamp in degrees. This property allows to set arbitrary rotate angle. 
	RotateAngle float64 `json:"RotateAngle,omitempty"`
	// Horizontal stamp coordinate, starting from the left.
	XIndent float64 `json:"XIndent,omitempty"`
	// Vertical stamp coordinate, starting from the bottom.
	YIndent float64 `json:"YIndent,omitempty"`
	// Zooming factor of the stamp. Allows to scale stamp.
	Zoom float64 `json:"Zoom,omitempty"`
	// Alignment of the text inside the stamp.
	TextAlignment HorizontalAlignment `json:"TextAlignment,omitempty"`
	// Gets or sets string value which is used as stamp on the page.
	Value string `json:"Value,omitempty"`
	// Gets text properties of the stamp. See TextState for details.
	TextState *TextState `json:"TextState,omitempty"`
	// Gets or sets vertical alignment of stamp on page.
	VerticalAlignment VerticalAlignment `json:"VerticalAlignment,omitempty"`
	// Gets or sets bottom margin of stamp.
	BottomMargin float64 `json:"BottomMargin,omitempty"`
	// Gets or sets left margin of stamp.
	LeftMargin float64 `json:"LeftMargin,omitempty"`
	// Gets or sets top margin of stamp.
	TopMargin float64 `json:"TopMargin,omitempty"`
	// Gets or sets right margin of stamp.
	RightMargin float64 `json:"RightMargin,omitempty"`
	// Page number.
	PageNumber int32 `json:"PageNumber"`
}
