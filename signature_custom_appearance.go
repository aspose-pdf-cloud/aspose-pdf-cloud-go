 /**
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

// An abstract class which represents signature custon appearance object.
type SignatureCustomAppearance struct {
	// Gets/sets font family name. It should be existed in the document. Default value: Arial.
	FontFamilyName string `json:"FontFamilyName,omitempty"`
	// Gets/sets font size. Default value: 10.
	FontSize float64 `json:"FontSize"`
	// Gets/sets contact info visibility. Default value: true.
	ShowContactInfo bool `json:"ShowContactInfo"`
	// Gets/sets reason visibility. Default value: true.
	ShowReason bool `json:"ShowReason"`
	// Gets/sets location visibility. Default value: true.
	ShowLocation bool `json:"ShowLocation"`
	// Gets/sets contact info label. Default value: \"Contact\".
	ContactInfoLabel string `json:"ContactInfoLabel,omitempty"`
	// Gets/sets reason label. Default value: \"Reason\".
	ReasonLabel string `json:"ReasonLabel,omitempty"`
	// Gets/sets location label. Default value: \"Location\".
	LocationLabel string `json:"LocationLabel,omitempty"`
	// Gets/sets digital signed label. Default value: \"Digitally signed by\".
	DigitalSignedLabel string `json:"DigitalSignedLabel,omitempty"`
	// Gets/sets date signed label. Default value: \"Date\".
	DateSignedAtLabel string `json:"DateSignedAtLabel,omitempty"`
	// Gets/sets datetime local format. Default value: \"yyyy.MM.dd HH:mm:ss zzz\".
	DateTimeLocalFormat string `json:"DateTimeLocalFormat,omitempty"`
	// Gets/sets datetime format. Default value: \"yyyy.MM.dd HH:mm:ss\".
	DateTimeFormat string `json:"DateTimeFormat,omitempty"`
}
