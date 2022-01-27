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

// Provides FreeTextAnnotation.
type FreeTextAnnotation struct {
	// Link to the document.
	Links []Link `json:"Links,omitempty"`
	// Color of the annotation.
	Color *Color `json:"Color,omitempty"`
	// Get the annotation content.
	Contents string `json:"Contents,omitempty"`
	// The date and time when the annotation was last modified.
	Modified string `json:"Modified,omitempty"`
	// Gets ID of the annotation.
	Id string `json:"Id,omitempty"`
	// Gets Flags of the annotation.
	Flags []AnnotationFlags `json:"Flags,omitempty"`
	// Gets Name of the annotation.
	Name string `json:"Name,omitempty"`
	// Gets Rect of the annotation.
	Rect *Rectangle `json:"Rect"`
	// Gets PageIndex of the annotation.
	PageIndex int32 `json:"PageIndex,omitempty"`
	// Gets ZIndex of the annotation.
	ZIndex int32 `json:"ZIndex,omitempty"`
	// Gets HorizontalAlignment of the annotation.
	HorizontalAlignment HorizontalAlignment `json:"HorizontalAlignment,omitempty"`
	// Gets VerticalAlignment of the annotation.
	VerticalAlignment VerticalAlignment `json:"VerticalAlignment,omitempty"`
	// The date and time when the annotation was created.
	CreationDate string `json:"CreationDate,omitempty"`
	// Get the annotation subject.
	Subject string `json:"Subject,omitempty"`
	// Get the annotation title.
	Title string `json:"Title,omitempty"`
	// Get the annotation RichText.
	RichText string `json:"RichText,omitempty"`
	// Gets Justification of the annotation.
	Justification Justification `json:"Justification,omitempty"`
	// Gets or sets the intent of the free text annotation.
	Intent FreeTextIntent `json:"Intent,omitempty"`
	// Angle of annotation rotation.
	Rotate Rotation `json:"Rotate,omitempty"`
	// Text style of the annotation.
	TextStyle *TextStyle `json:"TextStyle"`
}
