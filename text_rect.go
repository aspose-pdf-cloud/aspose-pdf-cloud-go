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

// Represents text occurrence.
type TextRect struct {
	// Text of the occurrence.
	Text string `json:"Text,omitempty"`
	// Page on which the occurrence is found.
	Page int32 `json:"Page,omitempty"`
	// Rectangle of the occurrence.
	Rect *Rectangle `json:"Rect,omitempty"`
	// Gets or sets a horizontal alignment of text fragment. 
	HorizontalAlignment HorizontalAlignment `json:"HorizontalAlignment,omitempty"`
	// Gets or sets a vertical alignment of text fragment. 
	VerticalAlignment VerticalAlignment `json:"VerticalAlignment,omitempty"`
	// Gets or sets text position for text, represented with TextRect object.
	Position *Position `json:"Position,omitempty"`
	// Gets text position for text, represented with TextRect object. The YIndent of the Position structure represents baseline coordinate of the text fragment.
	BaselinePosition *Position `json:"BaselinePosition,omitempty"`
	// Gets or sets text state for the text that TextRect object represents.
	TextState *TextState `json:"TextState,omitempty"`
}
