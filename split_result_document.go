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

// Represents split result document,
type SplitResultDocument struct {
	// The \"href\" attribute contains the link's IRI. atom:link elements MUST have an href attribute, whose value MUST be a IRI reference
	Href string `json:"Href,omitempty"`
	// atom:link elements MAY have a \"rel\" attribute that indicates the link relation type.  If the \"rel\" attribute is not present, the link element MUST be interpreted as if the link relation type is \"alternate\".
	Rel string `json:"Rel,omitempty"`
	// On the link element, the \"type\" attribute's value is an advisory media type: it is a hint about the type of the representation that is expected to be returned when the value of the href attribute is dereferenced.  Note that the type attribute does not override the actual media type returned with the representation.
	Type_ string `json:"Type,omitempty"`
	// The \"title\" attribute conveys human-readable information about the link.  The content of the \"title\" attribute is Language-Sensitive.
	Title string `json:"Title,omitempty"`
	// Gets or sets the page number.
	Id int32 `json:"Id,omitempty"`
}
