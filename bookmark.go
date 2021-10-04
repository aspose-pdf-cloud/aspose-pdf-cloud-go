 /*
 *
 * Copyright (c) 2021 Aspose.PDF Cloud
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

// Provides link to bookmark.
type Bookmark struct {
	// Link to the document.
	Links []Link `json:"Links,omitempty"`
	// Get the Title;
	Title string `json:"Title,omitempty"`
	// Is bookmark italic.
	Italic bool `json:"Italic,omitempty"`
	// Is bookmark bold.
	Bold bool `json:"Bold,omitempty"`
	// Get the color
	Color *Color `json:"Color,omitempty"`
	// Gets or sets the action bound with the bookmark. If PageNumber is presented the action can not be specified. The action type includes: \"GoTo\", \"GoToR\", \"Launch\", \"Named\".
	Action string `json:"Action,omitempty"`
	// Gets or sets bookmark's hierarchy level.
	Level int32 `json:"Level,omitempty"`
	// Gets or sets bookmark's destination page. Required if action is set as string.Empty.
	Destination string `json:"Destination,omitempty"`
	// Gets or sets the type of display bookmark's destination page.
	PageDisplay string `json:"PageDisplay,omitempty"`
	// Gets or sets the bottom coordinate of page display.
	PageDisplayBottom int32 `json:"PageDisplay_Bottom,omitempty"`
	// Gets or sets the left coordinate of page display.
	PageDisplayLeft int32 `json:"PageDisplay_Left,omitempty"`
	// Gets or sets the right coordinate of page display.
	PageDisplayRight int32 `json:"PageDisplay_Right,omitempty"`
	// Gets or sets the top coordinate of page display.
	PageDisplayTop int32 `json:"PageDisplay_Top,omitempty"`
	// Gets or sets the zoom factor of page display.
	PageDisplayZoom int32 `json:"PageDisplay_Zoom,omitempty"`
	// Gets or sets the number of bookmark's destination page. 
	PageNumber int32 `json:"PageNumber,omitempty"`
	// Gets or sets the file (path) which is required for \"GoToR\" action of bookmark.
	RemoteFile string `json:"RemoteFile,omitempty"`
	// The children bookmarks.
	Bookmarks *Bookmarks `json:"Bookmarks,omitempty"`
}
