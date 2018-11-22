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

type Signature struct {
	// Gets or sets the signature path.
	SignaturePath string `json:"SignaturePath"`
	// Gets or sets the type of the signature.
	SignatureType *SignatureType `json:"SignatureType"`
	// Gets or sets the signature password.
	Password string `json:"Password,omitempty"`
	// Sets or gets a graphic appearance for the signature. Property value represents an image file name.
	Appearance string `json:"Appearance,omitempty"`
	// Gets or sets the reason of the signature.
	Reason string `json:"Reason,omitempty"`
	// Gets or sets the contact of the signature.
	Contact string `json:"Contact,omitempty"`
	// Gets or sets the location of the signature.
	Location string `json:"Location,omitempty"`
	// Gets or sets a value indicating whether this  is visible. Supports only when signing particular page.
	Visible bool `json:"Visible"`
	// Gets or sets the visible rectangle of the signature. Supports only when signing particular page.
	Rectangle *RectanglePdf `json:"Rectangle,omitempty"`
	// Gets or sets the name of the signature field. Supports only when signing document with particular form field.
	FormFieldName string `json:"FormFieldName,omitempty"`
	// Gets or sets the name of the person or authority signing the document..
	Authority string `json:"Authority,omitempty"`
	// Gets or sets the time of signing.
	Date string `json:"Date,omitempty"`
	// Gets or sets the showproperties in signature field
	ShowProperties bool `json:"ShowProperties"`
}
