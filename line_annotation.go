 /**
 *
 *   Copyright (c) 2018 Aspose.PDF Cloud
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

// Provides LineAnnotation.
type LineAnnotation struct {
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
	Rect *Rectangle `json:"Rect,omitempty"`
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
	// Gets or sets starting point of line.
	Starting *Point `json:"Starting,omitempty"`
	// Gets or sets line ending style for line starting point.
	StartingStyle LineEnding `json:"StartingStyle,omitempty"`
	// Gets or sets ending point of line.
	Ending *Point `json:"Ending,omitempty"`
	// Gets or sets ending style for end point of line.
	EndingStyle LineEnding `json:"EndingStyle,omitempty"`
	// Gets or sets interior color of the annotation.
	InteriorColor *Color `json:"InteriorColor,omitempty"`
	// Gets or sets leader line length.
	LeaderLine float64 `json:"LeaderLine,omitempty"`
	// Gets or sets length of leader line extension.
	LeaderLineExtension float64 `json:"LeaderLineExtension,omitempty"`
	// Gets or sets leader line offset.
	LeaderLineOffset float64 `json:"LeaderLineOffset,omitempty"`
	// Gets or sets boolean flag which determinies is contents must be shown as caption.
	ShowCaption bool `json:"ShowCaption,omitempty"`
	// Gets or sets caption text offset from its normal position.
	CaptionOffset *Point `json:"CaptionOffset,omitempty"`
	// Gets or sets annotation caption position.
	CaptionPosition CaptionPosition `json:"CaptionPosition,omitempty"`
	// Gets or sets the intent of the line annotation.
	Intent LineIntent `json:"Intent,omitempty"`
}
