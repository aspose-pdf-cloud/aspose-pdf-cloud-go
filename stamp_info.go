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

// Provides stamp info.
type StampInfo struct {
	// Link to the document.
	Links []Link `json:"Links,omitempty"`
	// Gets ID of the stamp.
	Id string `json:"Id,omitempty"`
	// Gets index on page of the stamp.
	IndexOnPage int32 `json:"IndexOnPage,omitempty"`
	// Gets PageIndex of the annotation.
	PageIndex int32 `json:"PageIndex,omitempty"`
	// Gets Rect of the annotation.
	Rect *Rectangle `json:"Rect,omitempty"`
	// Get the text content.
	Text string `json:"Text,omitempty"`
	// Gets the stamp is visible.
	Visible bool `json:"Visible,omitempty"`
	// Gets stamp type.
	StampType StampType `json:"StampType,omitempty"`
}
