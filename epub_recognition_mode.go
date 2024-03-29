/*
*
* Copyright (c) 2024 Aspose.PDF Cloud
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
// EpubRecognitionMode : When PDF file (that usually has fixed layout) is being converted,             the conversion engine tries to perform grouping and multi-level analysis to restore             the original document author's intent and produce result in flow layout.  This property tunes that conversion for this or that             desirable method of recognition of content.             
type EpubRecognitionMode string

// List of EpubRecognitionMode
const (
	EpubRecognitionModeFlow EpubRecognitionMode = "Flow"
	EpubRecognitionModePdfFlow EpubRecognitionMode = "PdfFlow"
	EpubRecognitionModeFixed EpubRecognitionMode = "Fixed"
)
