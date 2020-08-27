 /*
 *
 *   Copyright (c) 2020 Aspose.PDF Cloud
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

// Represents a cell of the table's row.
type Cell struct {
	// Gets or sets the cell have border.
	IsNoBorder bool `json:"IsNoBorder,omitempty"`
	// Gets or sets the padding.
	Margin *MarginInfo `json:"Margin,omitempty"`
	// Gets or sets the border.
	Border *BorderInfo `json:"Border,omitempty"`
	// Gets or sets the background color.
	BackgroundColor *Color `json:"BackgroundColor,omitempty"`
	// Gets or sets the background image file.
	BackgroundImageFile string `json:"BackgroundImageFile,omitempty"`
	// Gets or sets path of the background image file from storage.
	BackgroundImageStorageFile string `json:"BackgroundImageStorageFile,omitempty"`
	// Gets or sets the alignment.
	Alignment HorizontalAlignment `json:"Alignment,omitempty"`
	// Gets or sets the default cell text state.
	DefaultCellTextState *TextState `json:"DefaultCellTextState,omitempty"`
	// Gets or sets the cell's formatted text.
	Paragraphs []TextRect `json:"Paragraphs,omitempty"`
	// Gets or sets the cell's text word wrapped.
	IsWordWrapped bool `json:"IsWordWrapped,omitempty"`
	// Gets or sets the vertical alignment.
	VerticalAlignment VerticalAlignment `json:"VerticalAlignment,omitempty"`
	// Gets or sets the column span.
	ColSpan int32 `json:"ColSpan,omitempty"`
	// Gets or sets the row span.
	RowSpan int32 `json:"RowSpan,omitempty"`
	// Gets or sets the column width.
	Width float64 `json:"Width,omitempty"`
	// Gets or sets Html fragment.
	HtmlFragment string `json:"HtmlFragment,omitempty"`
	// Gets or sets ImageFragment list.
	Images []ImageFragment `json:"Images,omitempty"`
}
