/*
*
* Copyright (c) 2026 Aspose.PDF Cloud
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

// Provides link to attachment.
type Attachment struct {
	// Link to the document.
	Links []Link `json:"Links,omitempty"`
	// Gets text associated with the attachment. 
	Description string `json:"Description,omitempty"`
	// Gets subtype of the attachment file.
	MimeType string `json:"MimeType,omitempty"`
	// Gets the name of the attachment. 
	Name string `json:"Name,omitempty"`
	// The date and time when the embedded file was created.
	CreationDate string `json:"CreationDate,omitempty"`
	// The date and time when the embedded file was last modified.
	ModificationDate string `json:"ModificationDate,omitempty"`
	// The size of the uncompressed embedded file, in bytes.
	Size int32 `json:"Size"`
	// A 16-byte string that is the checksum of the bytes of the uncompressed embedded file.  The checksum is calculated by applying the standard MD5 message-digest algorithm  to the bytes of the embedded file stream.
	CheckSum string `json:"CheckSum,omitempty"`
}
