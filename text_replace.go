/*
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

// Single text replacement setting.
type TextReplace struct {
	// Original text.
	OldValue string `json:"OldValue"`
	// New text.
	NewValue string `json:"NewValue,omitempty"`
	// Gets or sets a value indicating whether search text is regular expression.
	Regex bool `json:"Regex"`
	// Text properties of a new text.
	TextState *TextState `json:"TextState,omitempty"`
	// Rectangle area where searched original text.
	Rect *Rectangle `json:"Rect,omitempty"`
	// The text after replacement is centered horizontally relative to the text being replaced.
	CenterTextHorizontally bool `json:"CenterTextHorizontally,omitempty"`
}
