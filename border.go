/*
*
* Copyright (c) 2022 Aspose.PDF Cloud
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

// Class representing characteristics of annotation border.
type Border struct {
	// Gets or sets border width.
	Width int32 `json:"Width,omitempty"`
	// Gets or sets effect intencity. Valid range of value is [0..2].
	EffectIntensity int32 `json:"EffectIntensity,omitempty"`
	// Gets or sets border style.
	Style BorderStyle `json:"Style,omitempty"`
	// Gets or sets border effect.
	Effect BorderEffect `json:"Effect,omitempty"`
	// Gets or sets dash pattern.
	Dash *Dash `json:"Dash,omitempty"`
	// Gets or sets border color.
	Color *Color `json:"Color,omitempty"`
}
