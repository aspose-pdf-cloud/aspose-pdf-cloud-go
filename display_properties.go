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

// Provides link to DisplayProperties.
type DisplayProperties struct {
	// Link to the document.
	Links []Link `json:"Links,omitempty"`
	// Gets or sets flag specifying whether position of the document's window will be centerd on the screen.
	CenterWindow bool `json:"CenterWindow,omitempty"`
	// Gets or sets reading order of text: L2R (left to right) or R2L (right to left).
	Direction Direction `json:"Direction,omitempty"`
	// Gets or sets flag specifying whether document's window title bar should display document title.
	DisplayDocTitle bool `json:"DisplayDocTitle,omitempty"`
	// Gets or sets flag specifying whether menu bar should be hidden when document is active.
	HideMenuBar bool `json:"HideMenuBar,omitempty"`
	// Gets or sets flag specifying whether toolbar should be hidden when document is active.
	HideToolBar bool `json:"HideToolBar,omitempty"`
	// Gets or sets flag specifying whether user interface elements should be hidden when document is active.
	HideWindowUI bool `json:"HideWindowUI,omitempty"`
	// Gets or sets page mode, specifying how to display the document on exiting full-screen mode.
	NonFullScreenPageMode PageMode `json:"NonFullScreenPageMode,omitempty"`
	// Gets or sets page layout which shall be used when the document is opened.
	PageLayout PageLayout `json:"PageLayout,omitempty"`
	// Gets or sets page mode, specifying how document should be displayed when opened.
	PageMode PageMode `json:"PageMode,omitempty"`
}
