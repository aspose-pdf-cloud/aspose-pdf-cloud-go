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

// Provides RadioButtonField.
type RadioButtonField struct {
	// Link to the document.
	Links []Link `json:"Links,omitempty"`
	// Field name.
	PartialName string `json:"PartialName,omitempty"`
	// Field rectangle.
	Rect *Rectangle `json:"Rect,omitempty"`
	// Field value.
	Value string `json:"Value,omitempty"`
	// Page index.
	PageIndex int32 `json:"PageIndex"`
	// Gets or sets height of the field.
	Height float64 `json:"Height,omitempty"`
	// Gets or sets width of the field.
	Width float64 `json:"Width,omitempty"`
	// Z index.
	ZIndex int32 `json:"ZIndex,omitempty"`
	// Is group.
	IsGroup bool `json:"IsGroup"`
	// Gets field parent.
	Parent *FormField `json:"Parent,omitempty"`
	// Property for Generator support. Used when field is added to header or footer. If true, this field will created once and it's appearance will be visible on all pages of the document. If false, separated field will be created for every document page.
	IsSharedField bool `json:"IsSharedField,omitempty"`
	// Gets Flags of the field.
	Flags []AnnotationFlags `json:"Flags,omitempty"`
	// Color of the annotation.
	Color *Color `json:"Color,omitempty"`
	// Get the field content.
	Contents string `json:"Contents,omitempty"`
	// Gets or sets a outer margin for paragraph (for pdf generation)
	Margin *MarginInfo `json:"Margin,omitempty"`
	// Field highlighting mode.
	Highlighting LinkHighlightingMode `json:"Highlighting,omitempty"`
	// Gets HorizontalAlignment of the field.
	HorizontalAlignment HorizontalAlignment `json:"HorizontalAlignment,omitempty"`
	// Gets VerticalAlignment of the field.
	VerticalAlignment VerticalAlignment `json:"VerticalAlignment,omitempty"`
	// Gets or sets multiselection flag.
	MultiSelect bool `json:"MultiSelect,omitempty"`
	// Gets collection of options of the radio button.
	Options []Option `json:"Options,omitempty"`
	// Gets collection of radio button options field.
	RadioButtonOptionsField []RadioButtonOptionField `json:"RadioButtonOptionsField,omitempty"`
	// Style of field box.
	Style BoxStyle `json:"Style,omitempty"`
	// Gets or sets index of selected item. Numbering of items is started from 1.
	Selected int32 `json:"Selected"`
}
