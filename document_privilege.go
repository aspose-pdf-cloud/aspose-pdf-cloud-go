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

// Represents the privileges for accessing Pdf file./>.
type DocumentPrivilege struct {
	// Sets the permission which allow print or not.  true is allow and false or not set is forbidden.
	AllowPrint bool `json:"AllowPrint,omitempty"`
	// Sets the permission which allow degraded printing or not.  true is allow and false or not set is forbidden.
	AllowDegradedPrinting bool `json:"AllowDegradedPrinting,omitempty"`
	// Sets the permission which allow modify contents or not.  true is allow and false or not set is forbidden.
	AllowModifyContents bool `json:"AllowModifyContents,omitempty"`
	// Sets the permission which allow copy or not.  true is allow and false or not set is forbidden.
	AllowCopy bool `json:"AllowCopy,omitempty"`
	// Sets the permission which allow modify annotations or not.  true is allow and false or not set is forbidden.
	AllowModifyAnnotations bool `json:"AllowModifyAnnotations,omitempty"`
	// Sets the permission which allow fill in forms or not.  true is allow and false or not set is forbidden.
	AllowFillIn bool `json:"AllowFillIn,omitempty"`
	// Sets the permission which allow screen readers or not.  true is allow and false or not set is forbidden.
	AllowScreenReaders bool `json:"AllowScreenReaders,omitempty"`
	// Sets the permission which allow assembly or not.  true is allow and false or not set is forbidden.
	AllowAssembly bool `json:"AllowAssembly,omitempty"`
	// Sets the print level of  document's privilege. Just as the Adobe Professional's Printing Allowed settings. 0: None. 1: Low Resolution (150 dpi). 2: High Resolution.
	PrintAllowLevel int32 `json:"PrintAllowLevel,omitempty"`
	// Sets the change level of  document's privilege. Just as the Adobe Professional's Changes Allowed settings. 0: None. 1: Inserting, Deleting and Rotating pages. 2: Filling in form fields and signing existing signature fields. 3: Commenting, filling in form fields, and signing existing signature fields. 4: Any except extracting pages.
	ChangeAllowLevel int32 `json:"ChangeAllowLevel,omitempty"`
	// Sets the copy level of  document's privilege. Just as the Adobe Professional's permission settings. 0: None. 1: Enable text access for screen reader devices for the visually impaired. 2: Enable copying of text, images and other content.
	CopyAllowLevel int32 `json:"CopyAllowLevel,omitempty"`
}
