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

// This class represents border for graphics elements.
type BorderInfo struct {
	// Gets or sets a object that indicates left of the border.
	Left *GraphInfo `json:"Left,omitempty"`
	// Gets or sets a object that indicates right of the border.
	Right *GraphInfo `json:"Right,omitempty"`
	// Gets or sets a object that indicates the top border.
	Top *GraphInfo `json:"Top,omitempty"`
	// Gets or sets a object that indicates bottom of the border.
	Bottom *GraphInfo `json:"Bottom,omitempty"`
	// Gets or sets a rouded border radius
	RoundedBorderRadius float64 `json:"RoundedBorderRadius,omitempty"`
}
