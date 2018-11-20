// --------------------------------------------------------------------------------------------------------------------
// <copyright company="Aspose" file="OptimizeOptions.cs">
//   Copyright (c) 2018 Aspose.PDF Cloud
// </copyright>
// <summary>
//   Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
// 
//  The above copyright notice and this permission notice shall be included in all
//  copies or substantial portions of the Software.
// 
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//  SOFTWARE.
// </summary>
// --------------------------------------------------------------------------------------------------------------------

package asposepdfcloud

// Represents Pdf optimize options.
type OptimizeOptions struct {
	// If true page contents will be reused when document is optimized for equal pages.
	AllowReusePageContent bool `json:"AllowReusePageContent"`
	// If this flag is set to true images will be compressed in the document. compression level is specfied with ImageQuality property.
	CompressImages bool `json:"CompressImages"`
	// Specifie slevel of image compression when CompressIamges flag is used.
	ImageQuality int32 `json:"ImageQuality"`
	// If this flag is set to true, Resource streams will be analyzed. If duplicate streams are found (i.e. if stream contents is equal), then thes streams will be stored as one object.  This allows to decrease document size in some cases (for example, when same document was concatenedted multiple times).
	LinkDuplcateStreams bool `json:"LinkDuplcateStreams"`
	// If this flag is set to true, all document objects will be checked and unused objects (i.e. objects which does not have any reference) are removed from document.
	RemoveUnusedObjects bool `json:"RemoveUnusedObjects"`
	// If this flag set to true, every resource is checked on it's usage. If resource is never used, then resources is removed. This may decrease document size for example when pages were extracted from document. 
	RemoveUnusedStreams bool `json:"RemoveUnusedStreams"`
	// Make fonts not embedded if set to true. 
	UnembedFonts bool `json:"UnembedFonts"`
}
