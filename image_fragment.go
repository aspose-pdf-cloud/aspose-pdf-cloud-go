/*
*
* Copyright (c) 2026 Aspose.PDF Cloud
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

// Represents Image Fragment DTO.
type ImageFragment struct {
	// Gets or sets full storage path of image.
	ImageFile string `json:"ImageFile"`
	// Gets or sets fix width of the image.
	FixWidth float64 `json:"FixWidth,omitempty"`
	// Gets or sets fix height of the image.
	FixHeight float64 `json:"FixHeight,omitempty"`
	// Gets or sets horizontal alignment of the image.
	HorizontalAlignment HorizontalAlignment `json:"HorizontalAlignment,omitempty"`
	// Gets or sets vertical alignment of the image.
	VerticalAlignment VerticalAlignment `json:"VerticalAlignment,omitempty"`
	// Gets or sets ImageScale of the image.
	ImageScale float64 `json:"ImageScale,omitempty"`
	// Gets or sets Margin of the image.
	Margin *MarginInfo `json:"Margin,omitempty"`
}
