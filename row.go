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

// Represents a row of the table.
type Row struct {
	// Gets or sets the background color.
	BackgroundColor *Color `json:"BackgroundColor,omitempty"`
	// Gets or sets the border.
	Border *BorderInfo `json:"Border,omitempty"`
	// Sets the cells of the row.
	Cells []Cell `json:"Cells"`
	// Gets default cell border;
	DefaultCellBorder *BorderInfo `json:"DefaultCellBorder,omitempty"`
	// Gets height for row;
	MinRowHeight float64 `json:"MinRowHeight,omitempty"`
	// Gets fixed row height - row may have fixed height;
	FixedRowHeight float64 `json:"FixedRowHeight,omitempty"`
	// Gets fixed row is in new page - page with this property should be printed to next page Default false;
	IsInNewPage bool `json:"IsInNewPage,omitempty"`
	// Gets is row can be broken between two pages
	IsRowBroken bool `json:"IsRowBroken,omitempty"`
	// Gets or sets default text state for row cells
	DefaultCellTextState *TextState `json:"DefaultCellTextState,omitempty"`
	// Gets or sets default margin for row cells
	DefaultCellPadding *MarginInfo `json:"DefaultCellPadding,omitempty"`
	// Gets or sets the vertical alignment.
	VerticalAlignment VerticalAlignment `json:"VerticalAlignment,omitempty"`
}
