# \PdfApi

All URIs are relative to *https://api.aspose.cloud/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnnotation**](PdfApi.md#DeleteAnnotation) | **Delete** /pdf/{name}/annotations/{annotationId} | Delete document annotation by ID
[**DeleteDocumentAnnotations**](PdfApi.md#DeleteDocumentAnnotations) | **Delete** /pdf/{name}/annotations | Delete all annotations from the document
[**DeleteDocumentLinkAnnotations**](PdfApi.md#DeleteDocumentLinkAnnotations) | **Delete** /pdf/{name}/links | Delete all link annotations from the document
[**DeleteField**](PdfApi.md#DeleteField) | **Delete** /pdf/{name}/fields/{fieldName} | Delete document field by name.
[**DeleteImage**](PdfApi.md#DeleteImage) | **Delete** /pdf/{name}/images/{imageId} | Delete image from document page.
[**DeleteLinkAnnotation**](PdfApi.md#DeleteLinkAnnotation) | **Delete** /pdf/{name}/links/{linkId} | Delete document page link annotation by ID
[**DeletePage**](PdfApi.md#DeletePage) | **Delete** /pdf/{name}/pages/{pageNumber} | Delete document page by its number.
[**DeletePageAnnotations**](PdfApi.md#DeletePageAnnotations) | **Delete** /pdf/{name}/pages/{pageNumber}/annotations | Delete all annotations from the page
[**DeletePageLinkAnnotations**](PdfApi.md#DeletePageLinkAnnotations) | **Delete** /pdf/{name}/pages/{pageNumber}/links | Delete all link annotations from the page
[**DeleteProperties**](PdfApi.md#DeleteProperties) | **Delete** /pdf/{name}/documentproperties | Delete custom document properties.
[**DeleteProperty**](PdfApi.md#DeleteProperty) | **Delete** /pdf/{name}/documentproperties/{propertyName} | Delete document property.
[**GetCaretAnnotation**](PdfApi.md#GetCaretAnnotation) | **Get** /pdf/{name}/annotations/caret/{annotationId} | Read document page caret annotation by ID.
[**GetCircleAnnotation**](PdfApi.md#GetCircleAnnotation) | **Get** /pdf/{name}/annotations/circle/{annotationId} | Read document page circle annotation by ID.
[**GetDocument**](PdfApi.md#GetDocument) | **Get** /pdf/{name} | Read common document info.
[**GetDocumentAnnotations**](PdfApi.md#GetDocumentAnnotations) | **Get** /pdf/{name}/annotations | Read documant page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
[**GetDocumentAttachmentByIndex**](PdfApi.md#GetDocumentAttachmentByIndex) | **Get** /pdf/{name}/attachments/{attachmentIndex} | Read document attachment info by its index.
[**GetDocumentAttachments**](PdfApi.md#GetDocumentAttachments) | **Get** /pdf/{name}/attachments | Read document attachments info.
[**GetDocumentBookmarks**](PdfApi.md#GetDocumentBookmarks) | **Get** /pdf/{name}/bookmarks | Read document bookmark/bookmarks (including children).
[**GetDocumentCaretAnnotations**](PdfApi.md#GetDocumentCaretAnnotations) | **Get** /pdf/{name}/annotations/caret | Read document caret annotations.
[**GetDocumentCircleAnnotations**](PdfApi.md#GetDocumentCircleAnnotations) | **Get** /pdf/{name}/annotations/circle | Read document circle annotations.
[**GetDocumentFreeTextAnnotations**](PdfApi.md#GetDocumentFreeTextAnnotations) | **Get** /pdf/{name}/annotations/freetext | Read document free text annotations.
[**GetDocumentHighlightAnnotations**](PdfApi.md#GetDocumentHighlightAnnotations) | **Get** /pdf/{name}/annotations/highlight | Read document highlight annotations.
[**GetDocumentInkAnnotations**](PdfApi.md#GetDocumentInkAnnotations) | **Get** /pdf/{name}/annotations/ink | Read document ink annotations.
[**GetDocumentLineAnnotations**](PdfApi.md#GetDocumentLineAnnotations) | **Get** /pdf/{name}/annotations/line | Read document line annotations.
[**GetDocumentPolyLineAnnotations**](PdfApi.md#GetDocumentPolyLineAnnotations) | **Get** /pdf/{name}/annotations/polyline | Read document polyline annotations.
[**GetDocumentPolygonAnnotations**](PdfApi.md#GetDocumentPolygonAnnotations) | **Get** /pdf/{name}/annotations/polygon | Read document polygon annotations.
[**GetDocumentPopupAnnotations**](PdfApi.md#GetDocumentPopupAnnotations) | **Get** /pdf/{name}/annotations/popup | Read document popup annotations.
[**GetDocumentPopupAnnotationsByParent**](PdfApi.md#GetDocumentPopupAnnotationsByParent) | **Get** /pdf/{name}/annotations/{annotationId}/popup | Read document popup annotations by parent id.
[**GetDocumentProperties**](PdfApi.md#GetDocumentProperties) | **Get** /pdf/{name}/documentproperties | Read document properties.
[**GetDocumentProperty**](PdfApi.md#GetDocumentProperty) | **Get** /pdf/{name}/documentproperties/{propertyName} | Read document property by name.
[**GetDocumentSquareAnnotations**](PdfApi.md#GetDocumentSquareAnnotations) | **Get** /pdf/{name}/annotations/square | Read document square annotations.
[**GetDocumentSquigglyAnnotations**](PdfApi.md#GetDocumentSquigglyAnnotations) | **Get** /pdf/{name}/annotations/squiggly | Read document squiggly annotations.
[**GetDocumentStrikeOutAnnotations**](PdfApi.md#GetDocumentStrikeOutAnnotations) | **Get** /pdf/{name}/annotations/strikeout | Read document StrikeOut annotations.
[**GetDocumentTextAnnotations**](PdfApi.md#GetDocumentTextAnnotations) | **Get** /pdf/{name}/annotations/text | Read document text annotations.
[**GetDocumentUnderlineAnnotations**](PdfApi.md#GetDocumentUnderlineAnnotations) | **Get** /pdf/{name}/annotations/underline | Read document underline annotations.
[**GetDownload**](PdfApi.md#GetDownload) | **Get** /storage/file | Download a specific file 
[**GetDownloadDocumentAttachmentByIndex**](PdfApi.md#GetDownloadDocumentAttachmentByIndex) | **Get** /pdf/{name}/attachments/{attachmentIndex}/download | Download document attachment content by its index.
[**GetEpubInStorageToPdf**](PdfApi.md#GetEpubInStorageToPdf) | **Get** /pdf/create/epub | Convert EPUB file (located on storage) to PDF format and return resulting file in response. 
[**GetField**](PdfApi.md#GetField) | **Get** /pdf/{name}/fields/{fieldName} | Get document field by name.
[**GetFields**](PdfApi.md#GetFields) | **Get** /pdf/{name}/fields | Get document fields.
[**GetFreeTextAnnotation**](PdfApi.md#GetFreeTextAnnotation) | **Get** /pdf/{name}/annotations/freetext/{annotationId} | Read document page free text annotation by ID.
[**GetHighlightAnnotation**](PdfApi.md#GetHighlightAnnotation) | **Get** /pdf/{name}/annotations/highlight/{annotationId} | Read document page highlight annotation by ID.
[**GetHtmlInStorageToPdf**](PdfApi.md#GetHtmlInStorageToPdf) | **Get** /pdf/create/html | Convert HTML file (located on storage) to PDF format and return resulting file in response. 
[**GetImage**](PdfApi.md#GetImage) | **Get** /pdf/{name}/images/{imageId} | Read document image by ID.
[**GetImageExtractAsGif**](PdfApi.md#GetImageExtractAsGif) | **Get** /pdf/{name}/images/{imageId}/extract/gif | Extract document image in GIF format
[**GetImageExtractAsJpeg**](PdfApi.md#GetImageExtractAsJpeg) | **Get** /pdf/{name}/images/{imageId}/extract/jpeg | Extract document image in JPEG format
[**GetImageExtractAsPng**](PdfApi.md#GetImageExtractAsPng) | **Get** /pdf/{name}/images/{imageId}/extract/png | Extract document image in PNG format
[**GetImageExtractAsTiff**](PdfApi.md#GetImageExtractAsTiff) | **Get** /pdf/{name}/images/{imageId}/extract/tiff | Extract document image in TIFF format
[**GetImages**](PdfApi.md#GetImages) | **Get** /pdf/{name}/pages/{pageNumber}/images | Read document images.
[**GetInkAnnotation**](PdfApi.md#GetInkAnnotation) | **Get** /pdf/{name}/annotations/ink/{annotationId} | Read document page ink annotation by ID.
[**GetLaTeXInStorageToPdf**](PdfApi.md#GetLaTeXInStorageToPdf) | **Get** /pdf/create/latex | Convert LaTeX file (located on storage) to PDF format and return resulting file in response. 
[**GetLineAnnotation**](PdfApi.md#GetLineAnnotation) | **Get** /pdf/{name}/annotations/line/{annotationId} | Read document page line annotation by ID.
[**GetLinkAnnotation**](PdfApi.md#GetLinkAnnotation) | **Get** /pdf/{name}/links/{linkId} | Read document link annotation by ID.
[**GetListFiles**](PdfApi.md#GetListFiles) | **Get** /storage/folder | Get the file listing of a specific folder 
[**GetMhtInStorageToPdf**](PdfApi.md#GetMhtInStorageToPdf) | **Get** /pdf/create/mht | Convert MHT file (located on storage) to PDF format and return resulting file in response. 
[**GetPage**](PdfApi.md#GetPage) | **Get** /pdf/{name}/pages/{pageNumber} | Read document page info.
[**GetPageAnnotations**](PdfApi.md#GetPageAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations | Read documant page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
[**GetPageCaretAnnotations**](PdfApi.md#GetPageCaretAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/caret | Read document page caret annotations.
[**GetPageCircleAnnotations**](PdfApi.md#GetPageCircleAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/circle | Read document page circle annotations.
[**GetPageConvertToBmp**](PdfApi.md#GetPageConvertToBmp) | **Get** /pdf/{name}/pages/{pageNumber}/convert/bmp | Convert document page to Bmp image and return resulting file in response.
[**GetPageConvertToEmf**](PdfApi.md#GetPageConvertToEmf) | **Get** /pdf/{name}/pages/{pageNumber}/convert/emf | Convert document page to Emf image and return resulting file in response.
[**GetPageConvertToGif**](PdfApi.md#GetPageConvertToGif) | **Get** /pdf/{name}/pages/{pageNumber}/convert/gif | Convert document page to Gif image and return resulting file in response.
[**GetPageConvertToJpeg**](PdfApi.md#GetPageConvertToJpeg) | **Get** /pdf/{name}/pages/{pageNumber}/convert/jpeg | Convert document page to Jpeg image and return resulting file in response.
[**GetPageConvertToPng**](PdfApi.md#GetPageConvertToPng) | **Get** /pdf/{name}/pages/{pageNumber}/convert/png | Convert document page to Png image and return resulting file in response.
[**GetPageConvertToTiff**](PdfApi.md#GetPageConvertToTiff) | **Get** /pdf/{name}/pages/{pageNumber}/convert/tiff | Convert document page to Tiff image  and return resulting file in response.
[**GetPageFreeTextAnnotations**](PdfApi.md#GetPageFreeTextAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/freetext | Read document page free text annotations.
[**GetPageHighlightAnnotations**](PdfApi.md#GetPageHighlightAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/highlight | Read document page highlight annotations.
[**GetPageInkAnnotations**](PdfApi.md#GetPageInkAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/ink | Read document page ink annotations.
[**GetPageLineAnnotations**](PdfApi.md#GetPageLineAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/line | Read document page line annotations.
[**GetPageLinkAnnotation**](PdfApi.md#GetPageLinkAnnotation) | **Get** /pdf/{name}/pages/{pageNumber}/links/{linkId} | Read document page link annotation by ID.
[**GetPageLinkAnnotations**](PdfApi.md#GetPageLinkAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/links | Read document page link annotations.
[**GetPagePolyLineAnnotations**](PdfApi.md#GetPagePolyLineAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/polyline | Read document page polyline annotations.
[**GetPagePolygonAnnotations**](PdfApi.md#GetPagePolygonAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/polygon | Read document page polygon annotations.
[**GetPagePopupAnnotations**](PdfApi.md#GetPagePopupAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/popup | Read document page popup annotations.
[**GetPageSquareAnnotations**](PdfApi.md#GetPageSquareAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/square | Read document page square annotations.
[**GetPageSquigglyAnnotations**](PdfApi.md#GetPageSquigglyAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/squiggly | Read document page squiggly annotations.
[**GetPageStrikeOutAnnotations**](PdfApi.md#GetPageStrikeOutAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/strikeout | Read document page StrikeOut annotations.
[**GetPageText**](PdfApi.md#GetPageText) | **Get** /pdf/{name}/pages/{pageNumber}/text | Read page text items.
[**GetPageTextAnnotations**](PdfApi.md#GetPageTextAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/text | Read document page text annotations.
[**GetPageUnderlineAnnotations**](PdfApi.md#GetPageUnderlineAnnotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/underline | Read document page underline annotations.
[**GetPages**](PdfApi.md#GetPages) | **Get** /pdf/{name}/pages | Read document pages info.
[**GetPclInStorageToPdf**](PdfApi.md#GetPclInStorageToPdf) | **Get** /pdf/create/pcl | Convert PCL file (located on storage) to PDF format and return resulting file in response. 
[**GetPdfInStorageToDoc**](PdfApi.md#GetPdfInStorageToDoc) | **Get** /pdf/{name}/convert/doc | Converts PDF document (located on storage) to DOC format and returns resulting file in response content
[**GetPdfInStorageToEpub**](PdfApi.md#GetPdfInStorageToEpub) | **Get** /pdf/{name}/convert/epub | Converts PDF document (located on storage) to EPUB format and returns resulting file in response content
[**GetPdfInStorageToHtml**](PdfApi.md#GetPdfInStorageToHtml) | **Get** /pdf/{name}/convert/html | Converts PDF document (located on storage) to Html format and returns resulting file in response content
[**GetPdfInStorageToLaTeX**](PdfApi.md#GetPdfInStorageToLaTeX) | **Get** /pdf/{name}/convert/latex | Converts PDF document (located on storage) to LaTeX format and returns resulting file in response content
[**GetPdfInStorageToMobiXml**](PdfApi.md#GetPdfInStorageToMobiXml) | **Get** /pdf/{name}/convert/mobixml | Converts PDF document (located on storage) to MOBIXML format and returns resulting file in response content
[**GetPdfInStorageToPdfA**](PdfApi.md#GetPdfInStorageToPdfA) | **Get** /pdf/{name}/convert/pdfa | Converts PDF document (located on storage) to PdfA format and returns resulting file in response content
[**GetPdfInStorageToPptx**](PdfApi.md#GetPdfInStorageToPptx) | **Get** /pdf/{name}/convert/pptx | Converts PDF document (located on storage) to PPTX format and returns resulting file in response content
[**GetPdfInStorageToSvg**](PdfApi.md#GetPdfInStorageToSvg) | **Get** /pdf/{name}/convert/svg | Converts PDF document (located on storage) to SVG format and returns resulting file in response content
[**GetPdfInStorageToTiff**](PdfApi.md#GetPdfInStorageToTiff) | **Get** /pdf/{name}/convert/tiff | Converts PDF document (located on storage) to TIFF format and returns resulting file in response content
[**GetPdfInStorageToXls**](PdfApi.md#GetPdfInStorageToXls) | **Get** /pdf/{name}/convert/xls | Converts PDF document (located on storage) to XLS format and returns resulting file in response content
[**GetPdfInStorageToXml**](PdfApi.md#GetPdfInStorageToXml) | **Get** /pdf/{name}/convert/xml | Converts PDF document (located on storage) to XML format and returns resulting file in response content
[**GetPdfInStorageToXps**](PdfApi.md#GetPdfInStorageToXps) | **Get** /pdf/{name}/convert/xps | Converts PDF document (located on storage) to XPS format and returns resulting file in response content
[**GetPolyLineAnnotation**](PdfApi.md#GetPolyLineAnnotation) | **Get** /pdf/{name}/annotations/polyline/{annotationId} | Read document page polyline annotation by ID.
[**GetPolygonAnnotation**](PdfApi.md#GetPolygonAnnotation) | **Get** /pdf/{name}/annotations/polygon/{annotationId} | Read document page polygon annotation by ID.
[**GetPopupAnnotation**](PdfApi.md#GetPopupAnnotation) | **Get** /pdf/{name}/annotations/popup/{annotationId} | Read document page popup annotation by ID.
[**GetPsInStorageToPdf**](PdfApi.md#GetPsInStorageToPdf) | **Get** /pdf/create/ps | Convert PS file (located on storage) to PDF format and return resulting file in response. 
[**GetSquareAnnotation**](PdfApi.md#GetSquareAnnotation) | **Get** /pdf/{name}/annotations/square/{annotationId} | Read document page square annotation by ID.
[**GetSquigglyAnnotation**](PdfApi.md#GetSquigglyAnnotation) | **Get** /pdf/{name}/annotations/squiggly/{annotationId} | Read document page squiggly annotation by ID.
[**GetStrikeOutAnnotation**](PdfApi.md#GetStrikeOutAnnotation) | **Get** /pdf/{name}/annotations/strikeout/{annotationId} | Read document page StrikeOut annotation by ID.
[**GetSvgInStorageToPdf**](PdfApi.md#GetSvgInStorageToPdf) | **Get** /pdf/create/svg | Convert SVG file (located on storage) to PDF format and return resulting file in response. 
[**GetText**](PdfApi.md#GetText) | **Get** /pdf/{name}/text | Read document text.
[**GetTextAnnotation**](PdfApi.md#GetTextAnnotation) | **Get** /pdf/{name}/annotations/text/{annotationId} | Read document page text annotation by ID.
[**GetUnderlineAnnotation**](PdfApi.md#GetUnderlineAnnotation) | **Get** /pdf/{name}/annotations/underline/{annotationId} | Read document page underline annotation by ID.
[**GetVerifySignature**](PdfApi.md#GetVerifySignature) | **Get** /pdf/{name}/verifySignature | Verify signature document.
[**GetWebInStorageToPdf**](PdfApi.md#GetWebInStorageToPdf) | **Get** /pdf/create/web | Convert web page to PDF format and return resulting file in response. 
[**GetWordsPerPage**](PdfApi.md#GetWordsPerPage) | **Get** /pdf/{name}/pages/wordCount | Get number of words per document page.
[**GetXfaPdfInStorageToAcroForm**](PdfApi.md#GetXfaPdfInStorageToAcroForm) | **Get** /pdf/{name}/convert/xfatoacroform | Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and returns resulting file response content
[**GetXmlInStorageToPdf**](PdfApi.md#GetXmlInStorageToPdf) | **Get** /pdf/create/xml | Convert XML file (located on storage) to PDF format and return resulting file in response. 
[**GetXpsInStorageToPdf**](PdfApi.md#GetXpsInStorageToPdf) | **Get** /pdf/create/xps | Convert XPS file (located on storage) to PDF format and return resulting file in response. 
[**GetXslFoInStorageToPdf**](PdfApi.md#GetXslFoInStorageToPdf) | **Get** /pdf/create/xslfo | Convert XslFo file (located on storage) to PDF format and return resulting file in response. 
[**PostAppendDocument**](PdfApi.md#PostAppendDocument) | **Post** /pdf/{name}/appendDocument | Append document to existing one.
[**PostCreateField**](PdfApi.md#PostCreateField) | **Post** /pdf/{name}/fields | Create field.
[**PostDocumentTextReplace**](PdfApi.md#PostDocumentTextReplace) | **Post** /pdf/{name}/text/replace | Document&#39;s replace text method.
[**PostInsertImage**](PdfApi.md#PostInsertImage) | **Post** /pdf/{name}/pages/{pageNumber}/images | Insert image to document page.
[**PostMovePage**](PdfApi.md#PostMovePage) | **Post** /pdf/{name}/pages/{pageNumber}/movePage | Move page to new position.
[**PostOptimizeDocument**](PdfApi.md#PostOptimizeDocument) | **Post** /pdf/{name}/optimize | Optimize document.
[**PostPageCaretAnnotations**](PdfApi.md#PostPageCaretAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/caret | Add document page caret annotations.
[**PostPageCircleAnnotations**](PdfApi.md#PostPageCircleAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/circle | Add document page circle annotations.
[**PostPageFreeTextAnnotations**](PdfApi.md#PostPageFreeTextAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/freetext | Add document page free text annotations.
[**PostPageHighlightAnnotations**](PdfApi.md#PostPageHighlightAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/highlight | Add document page highlight annotations.
[**PostPageInkAnnotations**](PdfApi.md#PostPageInkAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/ink | Add document page ink annotations.
[**PostPageLineAnnotations**](PdfApi.md#PostPageLineAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/line | Add document page line annotations.
[**PostPageLinkAnnotations**](PdfApi.md#PostPageLinkAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/links | Add document page link annotations.
[**PostPagePolyLineAnnotations**](PdfApi.md#PostPagePolyLineAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/polyline | Add document page polyline annotations.
[**PostPagePolygonAnnotations**](PdfApi.md#PostPagePolygonAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/polygon | Add document page polygon annotations.
[**PostPageSquareAnnotations**](PdfApi.md#PostPageSquareAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/square | Add document page square annotations.
[**PostPageSquigglyAnnotations**](PdfApi.md#PostPageSquigglyAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/squiggly | Add document page squiggly annotations.
[**PostPageStrikeOutAnnotations**](PdfApi.md#PostPageStrikeOutAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/strikeout | Add document page StrikeOut annotations.
[**PostPageTextAnnotations**](PdfApi.md#PostPageTextAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/text | Add document page text annotations.
[**PostPageTextReplace**](PdfApi.md#PostPageTextReplace) | **Post** /pdf/{name}/pages/{pageNumber}/text/replace | Page&#39;s replace text method.
[**PostPageUnderlineAnnotations**](PdfApi.md#PostPageUnderlineAnnotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/underline | Add document page underline annotations.
[**PostPopupAnnotation**](PdfApi.md#PostPopupAnnotation) | **Post** /pdf/{name}/annotations/{annotationId}/popup | Add document popup annotations.
[**PostSignDocument**](PdfApi.md#PostSignDocument) | **Post** /pdf/{name}/sign | Sign document.
[**PostSignPage**](PdfApi.md#PostSignPage) | **Post** /pdf/{name}/pages/{pageNumber}/sign | Sign page.
[**PostSplitDocument**](PdfApi.md#PostSplitDocument) | **Post** /pdf/{name}/split | Split document to parts.
[**PutAddNewPage**](PdfApi.md#PutAddNewPage) | **Put** /pdf/{name}/pages | Add new page to end of the document.
[**PutAddText**](PdfApi.md#PutAddText) | **Put** /pdf/{name}/pages/{pageNumber}/text | Add text to PDF document page.
[**PutCaretAnnotation**](PdfApi.md#PutCaretAnnotation) | **Put** /pdf/{name}/annotations/caret/{annotationId} | Replace document caret annotation
[**PutCircleAnnotation**](PdfApi.md#PutCircleAnnotation) | **Put** /pdf/{name}/annotations/circle/{annotationId} | Replace document circle annotation
[**PutCreate**](PdfApi.md#PutCreate) | **Put** /storage/file | Upload a specific file 
[**PutCreateDocument**](PdfApi.md#PutCreateDocument) | **Put** /pdf/{name} | Create empty document.
[**PutEpubInStorageToPdf**](PdfApi.md#PutEpubInStorageToPdf) | **Put** /pdf/{name}/create/epub | Convert EPUB file (located on storage) to PDF format and upload resulting file to storage. 
[**PutFieldsFlatten**](PdfApi.md#PutFieldsFlatten) | **Put** /pdf/{name}/fields/flatten | Flatten form fields in document.
[**PutFreeTextAnnotation**](PdfApi.md#PutFreeTextAnnotation) | **Put** /pdf/{name}/annotations/freetext/{annotationId} | Replace document free text annotation
[**PutHighlightAnnotation**](PdfApi.md#PutHighlightAnnotation) | **Put** /pdf/{name}/annotations/highlight/{annotationId} | Replace document highlight annotation
[**PutHtmlInStorageToPdf**](PdfApi.md#PutHtmlInStorageToPdf) | **Put** /pdf/{name}/create/html | Convert HTML file (located on storage) to PDF format and upload resulting file to storage. 
[**PutImageExtractAsGif**](PdfApi.md#PutImageExtractAsGif) | **Put** /pdf/{name}/images/{imageId}/extract/gif | Extract document image in GIF format to folder
[**PutImageExtractAsJpeg**](PdfApi.md#PutImageExtractAsJpeg) | **Put** /pdf/{name}/images/{imageId}/extract/jpeg | Extract document image in JPEG format to folder
[**PutImageExtractAsPng**](PdfApi.md#PutImageExtractAsPng) | **Put** /pdf/{name}/images/{imageId}/extract/png | Extract document image in PNG format to folder
[**PutImageExtractAsTiff**](PdfApi.md#PutImageExtractAsTiff) | **Put** /pdf/{name}/images/{imageId}/extract/tiff | Extract document image in TIFF format to folder
[**PutImageInStorageToPdf**](PdfApi.md#PutImageInStorageToPdf) | **Put** /pdf/{name}/create/images | Convert image file (located on storage) to PDF format and upload resulting file to storage. 
[**PutImagesExtractAsGif**](PdfApi.md#PutImagesExtractAsGif) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/gif | Extract document images in GIF format to folder.
[**PutImagesExtractAsJpeg**](PdfApi.md#PutImagesExtractAsJpeg) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/jpeg | Extract document images in JPEG format to folder.
[**PutImagesExtractAsPng**](PdfApi.md#PutImagesExtractAsPng) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/png | Extract document images in PNG format to folder.
[**PutImagesExtractAsTiff**](PdfApi.md#PutImagesExtractAsTiff) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/tiff | Extract document images in TIFF format to folder.
[**PutInkAnnotation**](PdfApi.md#PutInkAnnotation) | **Put** /pdf/{name}/annotations/ink/{annotationId} | Replace document ink annotation
[**PutLaTeXInStorageToPdf**](PdfApi.md#PutLaTeXInStorageToPdf) | **Put** /pdf/{name}/create/latex | Convert LaTeX file (located on storage) to PDF format and upload resulting file to storage. 
[**PutLineAnnotation**](PdfApi.md#PutLineAnnotation) | **Put** /pdf/{name}/annotations/line/{annotationId} | Replace document line annotation
[**PutLinkAnnotation**](PdfApi.md#PutLinkAnnotation) | **Put** /pdf/{name}/links/{linkId} | Replace document page link annotations
[**PutMergeDocuments**](PdfApi.md#PutMergeDocuments) | **Put** /pdf/{name}/merge | Merge a list of documents.
[**PutMhtInStorageToPdf**](PdfApi.md#PutMhtInStorageToPdf) | **Put** /pdf/{name}/create/mht | Convert MHT file (located on storage) to PDF format and upload resulting file to storage. 
[**PutPageAddStamp**](PdfApi.md#PutPageAddStamp) | **Put** /pdf/{name}/pages/{pageNumber}/stamp | Add page stamp.
[**PutPageConvertToBmp**](PdfApi.md#PutPageConvertToBmp) | **Put** /pdf/{name}/pages/{pageNumber}/convert/bmp | Convert document page to bmp image and upload resulting file to storage.
[**PutPageConvertToEmf**](PdfApi.md#PutPageConvertToEmf) | **Put** /pdf/{name}/pages/{pageNumber}/convert/emf | Convert document page to emf image and upload resulting file to storage.
[**PutPageConvertToGif**](PdfApi.md#PutPageConvertToGif) | **Put** /pdf/{name}/pages/{pageNumber}/convert/gif | Convert document page to gif image and upload resulting file to storage.
[**PutPageConvertToJpeg**](PdfApi.md#PutPageConvertToJpeg) | **Put** /pdf/{name}/pages/{pageNumber}/convert/jpeg | Convert document page to Jpeg image and upload resulting file to storage.
[**PutPageConvertToPng**](PdfApi.md#PutPageConvertToPng) | **Put** /pdf/{name}/pages/{pageNumber}/convert/png | Convert document page to png image and upload resulting file to storage.
[**PutPageConvertToTiff**](PdfApi.md#PutPageConvertToTiff) | **Put** /pdf/{name}/pages/{pageNumber}/convert/tiff | Convert document page to Tiff image and upload resulting file to storage.
[**PutPclInStorageToPdf**](PdfApi.md#PutPclInStorageToPdf) | **Put** /pdf/{name}/create/pcl | Convert PCL file (located on storage) to PDF format and upload resulting file to storage. 
[**PutPdfInRequestToDoc**](PdfApi.md#PutPdfInRequestToDoc) | **Put** /pdf/convert/doc | Converts PDF document (in request content) to DOC format and uploads resulting file to storage.
[**PutPdfInRequestToEpub**](PdfApi.md#PutPdfInRequestToEpub) | **Put** /pdf/convert/epub | Converts PDF document (in request content) to EPUB format and uploads resulting file to storage.
[**PutPdfInRequestToHtml**](PdfApi.md#PutPdfInRequestToHtml) | **Put** /pdf/convert/html | Converts PDF document (in request content) to Html format and uploads resulting file to storage.
[**PutPdfInRequestToLaTeX**](PdfApi.md#PutPdfInRequestToLaTeX) | **Put** /pdf/convert/latex | Converts PDF document (in request content) to LaTeX format and uploads resulting file to storage.
[**PutPdfInRequestToMobiXml**](PdfApi.md#PutPdfInRequestToMobiXml) | **Put** /pdf/convert/mobixml | Converts PDF document (in request content) to MOBIXML format and uploads resulting file to storage.
[**PutPdfInRequestToPdfA**](PdfApi.md#PutPdfInRequestToPdfA) | **Put** /pdf/convert/pdfa | Converts PDF document (in request content) to PdfA format and uploads resulting file to storage.
[**PutPdfInRequestToPptx**](PdfApi.md#PutPdfInRequestToPptx) | **Put** /pdf/convert/pptx | Converts PDF document (in request content) to PPTX format and uploads resulting file to storage.
[**PutPdfInRequestToSvg**](PdfApi.md#PutPdfInRequestToSvg) | **Put** /pdf/convert/svg | Converts PDF document (in request content) to SVG format and uploads resulting file to storage.
[**PutPdfInRequestToTiff**](PdfApi.md#PutPdfInRequestToTiff) | **Put** /pdf/convert/tiff | Converts PDF document (in request content) to TIFF format and uploads resulting file to storage.
[**PutPdfInRequestToXls**](PdfApi.md#PutPdfInRequestToXls) | **Put** /pdf/convert/xls | Converts PDF document (in request content) to XLS format and uploads resulting file to storage.
[**PutPdfInRequestToXml**](PdfApi.md#PutPdfInRequestToXml) | **Put** /pdf/convert/xml | Converts PDF document (in request content) to XML format and uploads resulting file to storage.
[**PutPdfInRequestToXps**](PdfApi.md#PutPdfInRequestToXps) | **Put** /pdf/convert/xps | Converts PDF document (in request content) to XPS format and uploads resulting file to storage.
[**PutPdfInStorageToDoc**](PdfApi.md#PutPdfInStorageToDoc) | **Put** /pdf/{name}/convert/doc | Converts PDF document (located on storage) to DOC format and uploads resulting file to storage
[**PutPdfInStorageToEpub**](PdfApi.md#PutPdfInStorageToEpub) | **Put** /pdf/{name}/convert/epub | Converts PDF document (located on storage) to EPUB format and uploads resulting file to storage
[**PutPdfInStorageToHtml**](PdfApi.md#PutPdfInStorageToHtml) | **Put** /pdf/{name}/convert/html | Converts PDF document (located on storage) to Html format and uploads resulting file to storage
[**PutPdfInStorageToLaTeX**](PdfApi.md#PutPdfInStorageToLaTeX) | **Put** /pdf/{name}/convert/latex | Converts PDF document (located on storage) to LaTeX format and uploads resulting file to storage
[**PutPdfInStorageToMobiXml**](PdfApi.md#PutPdfInStorageToMobiXml) | **Put** /pdf/{name}/convert/mobixml | Converts PDF document (located on storage) to MOBIXML format and uploads resulting file to storage
[**PutPdfInStorageToPdfA**](PdfApi.md#PutPdfInStorageToPdfA) | **Put** /pdf/{name}/convert/pdfa | Converts PDF document (located on storage) to PdfA format and uploads resulting file to storage
[**PutPdfInStorageToPptx**](PdfApi.md#PutPdfInStorageToPptx) | **Put** /pdf/{name}/convert/pptx | Converts PDF document (located on storage) to PPTX format and uploads resulting file to storage
[**PutPdfInStorageToSvg**](PdfApi.md#PutPdfInStorageToSvg) | **Put** /pdf/{name}/convert/svg | Converts PDF document (located on storage) to SVG format and uploads resulting file to storage
[**PutPdfInStorageToTiff**](PdfApi.md#PutPdfInStorageToTiff) | **Put** /pdf/{name}/convert/tiff | Converts PDF document (located on storage) to TIFF format and uploads resulting file to storage
[**PutPdfInStorageToXls**](PdfApi.md#PutPdfInStorageToXls) | **Put** /pdf/{name}/convert/xls | Converts PDF document (located on storage) to XLS format and uploads resulting file to storage
[**PutPdfInStorageToXml**](PdfApi.md#PutPdfInStorageToXml) | **Put** /pdf/{name}/convert/xml | Converts PDF document (located on storage) to XML format and uploads resulting file to storage
[**PutPdfInStorageToXps**](PdfApi.md#PutPdfInStorageToXps) | **Put** /pdf/{name}/convert/xps | Converts PDF document (located on storage) to XPS format and uploads resulting file to storage
[**PutPolyLineAnnotation**](PdfApi.md#PutPolyLineAnnotation) | **Put** /pdf/{name}/annotations/polyline/{annotationId} | Replace document polyline annotation
[**PutPolygonAnnotation**](PdfApi.md#PutPolygonAnnotation) | **Put** /pdf/{name}/annotations/polygon/{annotationId} | Replace document polygon annotation
[**PutPopupAnnotation**](PdfApi.md#PutPopupAnnotation) | **Put** /pdf/{name}/annotations/popup/{annotationId} | Replace document popup annotation
[**PutPrivileges**](PdfApi.md#PutPrivileges) | **Put** /pdf/{name}/privileges | Update privilege document.
[**PutPsInStorageToPdf**](PdfApi.md#PutPsInStorageToPdf) | **Put** /pdf/{name}/create/ps | Convert PS file (located on storage) to PDF format and upload resulting file to storage. 
[**PutReplaceImage**](PdfApi.md#PutReplaceImage) | **Put** /pdf/{name}/images/{imageId} | Replace document image.
[**PutSearchableDocument**](PdfApi.md#PutSearchableDocument) | **Put** /pdf/{name}/ocr | Create searchable PDF document. Generate OCR layer for images in input PDF document.
[**PutSetProperty**](PdfApi.md#PutSetProperty) | **Put** /pdf/{name}/documentproperties/{propertyName} | Add/update document property.
[**PutSquareAnnotation**](PdfApi.md#PutSquareAnnotation) | **Put** /pdf/{name}/annotations/square/{annotationId} | Replace document square annotation
[**PutSquigglyAnnotation**](PdfApi.md#PutSquigglyAnnotation) | **Put** /pdf/{name}/annotations/squiggly/{annotationId} | Replace document squiggly annotation
[**PutStrikeOutAnnotation**](PdfApi.md#PutStrikeOutAnnotation) | **Put** /pdf/{name}/annotations/strikeout/{annotationId} | Replace document StrikeOut annotation
[**PutSvgInStorageToPdf**](PdfApi.md#PutSvgInStorageToPdf) | **Put** /pdf/{name}/create/svg | Convert SVG file (located on storage) to PDF format and upload resulting file to storage. 
[**PutTextAnnotation**](PdfApi.md#PutTextAnnotation) | **Put** /pdf/{name}/annotations/text/{annotationId} | Replace document text annotation
[**PutUnderlineAnnotation**](PdfApi.md#PutUnderlineAnnotation) | **Put** /pdf/{name}/annotations/underline/{annotationId} | Replace document underline annotation
[**PutUpdateField**](PdfApi.md#PutUpdateField) | **Put** /pdf/{name}/fields/{fieldName} | Update field.
[**PutUpdateFields**](PdfApi.md#PutUpdateFields) | **Put** /pdf/{name}/fields | Update fields.
[**PutWebInStorageToPdf**](PdfApi.md#PutWebInStorageToPdf) | **Put** /pdf/{name}/create/web | Convert web page to PDF format and upload resulting file to storage. 
[**PutXfaPdfInRequestToAcroForm**](PdfApi.md#PutXfaPdfInRequestToAcroForm) | **Put** /pdf/convert/xfatoacroform | Converts PDF document which contatins XFA form (in request content) to PDF with AcroForm and uploads resulting file to storage.
[**PutXfaPdfInStorageToAcroForm**](PdfApi.md#PutXfaPdfInStorageToAcroForm) | **Put** /pdf/{name}/convert/xfatoacroform | Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and uploads resulting file to storage
[**PutXmlInStorageToPdf**](PdfApi.md#PutXmlInStorageToPdf) | **Put** /pdf/{name}/create/xml | Convert XML file (located on storage) to PDF format and upload resulting file to storage. 
[**PutXpsInStorageToPdf**](PdfApi.md#PutXpsInStorageToPdf) | **Put** /pdf/{name}/create/xps | Convert XPS file (located on storage) to PDF format and upload resulting file to storage. 
[**PutXslFoInStorageToPdf**](PdfApi.md#PutXslFoInStorageToPdf) | **Put** /pdf/{name}/create/xslfo | Convert XslFo file (located on storage) to PDF format and upload resulting file to storage. 


# **DeleteAnnotation**
> AsposeResponse DeleteAnnotation(ctx, name, annotationId, optional)
Delete document annotation by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocumentAnnotations**
> AsposeResponse DeleteDocumentAnnotations(ctx, name, optional)
Delete all annotations from the document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDocumentLinkAnnotations**
> AsposeResponse DeleteDocumentLinkAnnotations(ctx, name, optional)
Delete all link annotations from the document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteField**
> AsposeResponse DeleteField(ctx, name, fieldName, optional)
Delete document field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name/ | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name/ | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteImage**
> AsposeResponse DeleteImage(ctx, name, imageId, optional)
Delete image from document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLinkAnnotation**
> AsposeResponse DeleteLinkAnnotation(ctx, name, linkId, optional)
Delete document page link annotation by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **linkId** | **string**| The link ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **linkId** | **string**| The link ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePage**
> AsposeResponse DeletePage(ctx, name, pageNumber, optional)
Delete document page by its number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePageAnnotations**
> AsposeResponse DeletePageAnnotations(ctx, name, pageNumber, optional)
Delete all annotations from the page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePageLinkAnnotations**
> AsposeResponse DeletePageLinkAnnotations(ctx, name, pageNumber, optional)
Delete all link annotations from the page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProperties**
> AsposeResponse DeleteProperties(ctx, name, optional)
Delete custom document properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProperty**
> AsposeResponse DeleteProperty(ctx, name, propertyName, optional)
Delete document property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
  **propertyName** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **propertyName** | **string**|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCaretAnnotation**
> CaretAnnotationResponse GetCaretAnnotation(ctx, name, annotationId, optional)
Read document page caret annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CaretAnnotationResponse**](CaretAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCircleAnnotation**
> CircleAnnotationResponse GetCircleAnnotation(ctx, name, annotationId, optional)
Read document page circle annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CircleAnnotationResponse**](CircleAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocument**
> DocumentResponse GetDocument(ctx, name, optional)
Read common document info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentAnnotations**
> AnnotationsInfoResponse GetDocumentAnnotations(ctx, name, optional)
Read documant page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AnnotationsInfoResponse**](AnnotationsInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentAttachmentByIndex**
> AttachmentResponse GetDocumentAttachmentByIndex(ctx, name, attachmentIndex, optional)
Read document attachment info by its index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **attachmentIndex** | **int32**| The attachment index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **attachmentIndex** | **int32**| The attachment index. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AttachmentResponse**](AttachmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentAttachments**
> AttachmentsResponse GetDocumentAttachments(ctx, name, optional)
Read document attachments info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AttachmentsResponse**](AttachmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentBookmarks**
> *os.File GetDocumentBookmarks(ctx, name, optional)
Read document bookmark/bookmarks (including children).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **bookmarkPath** | **string**| The bookmark path. Leave it empty if you want to get all the bookmarks in the document. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentCaretAnnotations**
> CaretAnnotationsResponse GetDocumentCaretAnnotations(ctx, name, optional)
Read document caret annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CaretAnnotationsResponse**](CaretAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentCircleAnnotations**
> CircleAnnotationsResponse GetDocumentCircleAnnotations(ctx, name, optional)
Read document circle annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CircleAnnotationsResponse**](CircleAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentFreeTextAnnotations**
> FreeTextAnnotationsResponse GetDocumentFreeTextAnnotations(ctx, name, optional)
Read document free text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FreeTextAnnotationsResponse**](FreeTextAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentHighlightAnnotations**
> HighlightAnnotationsResponse GetDocumentHighlightAnnotations(ctx, name, optional)
Read document highlight annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**HighlightAnnotationsResponse**](HighlightAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentInkAnnotations**
> InkAnnotationsResponse GetDocumentInkAnnotations(ctx, name, optional)
Read document ink annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**InkAnnotationsResponse**](InkAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentLineAnnotations**
> LineAnnotationsResponse GetDocumentLineAnnotations(ctx, name, optional)
Read document line annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LineAnnotationsResponse**](LineAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentPolyLineAnnotations**
> PolyLineAnnotationsResponse GetDocumentPolyLineAnnotations(ctx, name, optional)
Read document polyline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolyLineAnnotationsResponse**](PolyLineAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentPolygonAnnotations**
> PolygonAnnotationsResponse GetDocumentPolygonAnnotations(ctx, name, optional)
Read document polygon annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolygonAnnotationsResponse**](PolygonAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentPopupAnnotations**
> PopupAnnotationsResponse GetDocumentPopupAnnotations(ctx, name, optional)
Read document popup annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationsResponse**](PopupAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentPopupAnnotationsByParent**
> PopupAnnotationsResponse GetDocumentPopupAnnotationsByParent(ctx, name, annotationId, optional)
Read document popup annotations by parent id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The parent annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The parent annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationsResponse**](PopupAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentProperties**
> DocumentPropertiesResponse GetDocumentProperties(ctx, name, optional)
Read document properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**DocumentPropertiesResponse**](DocumentPropertiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentProperty**
> DocumentPropertyResponse GetDocumentProperty(ctx, name, propertyName, optional)
Read document property by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
  **propertyName** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **propertyName** | **string**|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**DocumentPropertyResponse**](DocumentPropertyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentSquareAnnotations**
> SquareAnnotationsResponse GetDocumentSquareAnnotations(ctx, name, optional)
Read document square annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquareAnnotationsResponse**](SquareAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentSquigglyAnnotations**
> SquigglyAnnotationsResponse GetDocumentSquigglyAnnotations(ctx, name, optional)
Read document squiggly annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquigglyAnnotationsResponse**](SquigglyAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentStrikeOutAnnotations**
> StrikeOutAnnotationsResponse GetDocumentStrikeOutAnnotations(ctx, name, optional)
Read document StrikeOut annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StrikeOutAnnotationsResponse**](StrikeOutAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentTextAnnotations**
> TextAnnotationsResponse GetDocumentTextAnnotations(ctx, name, optional)
Read document text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextAnnotationsResponse**](TextAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDocumentUnderlineAnnotations**
> UnderlineAnnotationsResponse GetDocumentUnderlineAnnotations(ctx, name, optional)
Read document underline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**UnderlineAnnotationsResponse**](UnderlineAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDownload**
> *os.File GetDownload(ctx, path, optional)
Download a specific file 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Path of the file including the file name and extension e.g. /file.ext | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Path of the file including the file name and extension e.g. /file.ext | 
 **versionId** | **string**| File&#39;s version | 
 **storage** | **string**| User&#39;s storage name | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDownloadDocumentAttachmentByIndex**
> *os.File GetDownloadDocumentAttachmentByIndex(ctx, name, attachmentIndex, optional)
Download document attachment content by its index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **attachmentIndex** | **int32**| The attachment index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **attachmentIndex** | **int32**| The attachment index. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEpubInStorageToPdf**
> *os.File GetEpubInStorageToPdf(ctx, srcPath, optional)
Convert EPUB file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.epub) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.epub) | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetField**
> FieldResponse GetField(ctx, name, fieldName, optional)
Get document field by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **fieldName** | **string**| The field name/ | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The field name/ | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FieldResponse**](FieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFields**
> FieldsResponse GetFields(ctx, name, optional)
Get document fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FieldsResponse**](FieldsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFreeTextAnnotation**
> FreeTextAnnotationResponse GetFreeTextAnnotation(ctx, name, annotationId, optional)
Read document page free text annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FreeTextAnnotationResponse**](FreeTextAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHighlightAnnotation**
> HighlightAnnotationResponse GetHighlightAnnotation(ctx, name, annotationId, optional)
Read document page highlight annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**HighlightAnnotationResponse**](HighlightAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHtmlInStorageToPdf**
> *os.File GetHtmlInStorageToPdf(ctx, srcPath, optional)
Convert HTML file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.zip) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.zip) | 
 **htmlFileName** | **string**| Name of HTML file in ZIP. | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImage**
> ImageResponse GetImage(ctx, name, imageId, optional)
Read document image by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ImageResponse**](ImageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageExtractAsGif**
> *os.File GetImageExtractAsGif(ctx, name, imageId, optional)
Extract document image in GIF format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageExtractAsJpeg**
> *os.File GetImageExtractAsJpeg(ctx, name, imageId, optional)
Extract document image in JPEG format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageExtractAsPng**
> *os.File GetImageExtractAsPng(ctx, name, imageId, optional)
Extract document image in PNG format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImageExtractAsTiff**
> *os.File GetImageExtractAsTiff(ctx, name, imageId, optional)
Extract document image in TIFF format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImages**
> ImagesResponse GetImages(ctx, name, pageNumber, optional)
Read document images.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**ImagesResponse**](ImagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInkAnnotation**
> InkAnnotationResponse GetInkAnnotation(ctx, name, annotationId, optional)
Read document page ink annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**InkAnnotationResponse**](InkAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLaTeXInStorageToPdf**
> *os.File GetLaTeXInStorageToPdf(ctx, srcPath, optional)
Convert LaTeX file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.tex) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.tex) | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLineAnnotation**
> LineAnnotationResponse GetLineAnnotation(ctx, name, annotationId, optional)
Read document page line annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LineAnnotationResponse**](LineAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkAnnotation**
> LinkAnnotationResponse GetLinkAnnotation(ctx, name, linkId, optional)
Read document link annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **linkId** | **string**| The link ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **linkId** | **string**| The link ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LinkAnnotationResponse**](LinkAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListFiles**
> FilesResponse GetListFiles(ctx, optional)
Get the file listing of a specific folder 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Start with name of storage e.g. root folder &#39;/&#39;or some folder &#39;/folder1/..&#39; | [default to /]
 **storage** | **string**| User&#39;s storage name | 

### Return type

[**FilesResponse**](FilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMhtInStorageToPdf**
> *os.File GetMhtInStorageToPdf(ctx, srcPath, optional)
Convert MHT file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.mht) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.mht) | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPage**
> DocumentPageResponse GetPage(ctx, name, pageNumber, optional)
Read document page info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**DocumentPageResponse**](DocumentPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageAnnotations**
> AnnotationsInfoResponse GetPageAnnotations(ctx, name, pageNumber, optional)
Read documant page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AnnotationsInfoResponse**](AnnotationsInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageCaretAnnotations**
> CaretAnnotationsResponse GetPageCaretAnnotations(ctx, name, pageNumber, optional)
Read document page caret annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CaretAnnotationsResponse**](CaretAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageCircleAnnotations**
> CircleAnnotationsResponse GetPageCircleAnnotations(ctx, name, pageNumber, optional)
Read document page circle annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CircleAnnotationsResponse**](CircleAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToBmp**
> *os.File GetPageConvertToBmp(ctx, name, pageNumber, optional)
Convert document page to Bmp image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToEmf**
> *os.File GetPageConvertToEmf(ctx, name, pageNumber, optional)
Convert document page to Emf image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToGif**
> *os.File GetPageConvertToGif(ctx, name, pageNumber, optional)
Convert document page to Gif image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToJpeg**
> *os.File GetPageConvertToJpeg(ctx, name, pageNumber, optional)
Convert document page to Jpeg image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToPng**
> *os.File GetPageConvertToPng(ctx, name, pageNumber, optional)
Convert document page to Png image and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageConvertToTiff**
> *os.File GetPageConvertToTiff(ctx, name, pageNumber, optional)
Convert document page to Tiff image  and return resulting file in response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageFreeTextAnnotations**
> FreeTextAnnotationsResponse GetPageFreeTextAnnotations(ctx, name, pageNumber, optional)
Read document page free text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FreeTextAnnotationsResponse**](FreeTextAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageHighlightAnnotations**
> HighlightAnnotationsResponse GetPageHighlightAnnotations(ctx, name, pageNumber, optional)
Read document page highlight annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**HighlightAnnotationsResponse**](HighlightAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageInkAnnotations**
> InkAnnotationsResponse GetPageInkAnnotations(ctx, name, pageNumber, optional)
Read document page ink annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**InkAnnotationsResponse**](InkAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageLineAnnotations**
> LineAnnotationsResponse GetPageLineAnnotations(ctx, name, pageNumber, optional)
Read document page line annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LineAnnotationsResponse**](LineAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageLinkAnnotation**
> LinkAnnotationResponse GetPageLinkAnnotation(ctx, name, pageNumber, linkId, optional)
Read document page link annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **linkId** | **string**| The link ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **linkId** | **string**| The link ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LinkAnnotationResponse**](LinkAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageLinkAnnotations**
> LinkAnnotationsResponse GetPageLinkAnnotations(ctx, name, pageNumber, optional)
Read document page link annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LinkAnnotationsResponse**](LinkAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagePolyLineAnnotations**
> PolyLineAnnotationsResponse GetPagePolyLineAnnotations(ctx, name, pageNumber, optional)
Read document page polyline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolyLineAnnotationsResponse**](PolyLineAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagePolygonAnnotations**
> PolygonAnnotationsResponse GetPagePolygonAnnotations(ctx, name, pageNumber, optional)
Read document page polygon annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolygonAnnotationsResponse**](PolygonAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagePopupAnnotations**
> PopupAnnotationsResponse GetPagePopupAnnotations(ctx, name, pageNumber, optional)
Read document page popup annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationsResponse**](PopupAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageSquareAnnotations**
> SquareAnnotationsResponse GetPageSquareAnnotations(ctx, name, pageNumber, optional)
Read document page square annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquareAnnotationsResponse**](SquareAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageSquigglyAnnotations**
> SquigglyAnnotationsResponse GetPageSquigglyAnnotations(ctx, name, pageNumber, optional)
Read document page squiggly annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquigglyAnnotationsResponse**](SquigglyAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageStrikeOutAnnotations**
> StrikeOutAnnotationsResponse GetPageStrikeOutAnnotations(ctx, name, pageNumber, optional)
Read document page StrikeOut annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StrikeOutAnnotationsResponse**](StrikeOutAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageText**
> TextRectsResponse GetPageText(ctx, name, pageNumber, lLX, lLY, uRX, uRY, optional)
Read page text items.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| Number of page (starting from 1). | 
  **lLX** | **float64**|  | 
  **lLY** | **float64**|  | 
  **uRX** | **float64**|  | 
  **uRY** | **float64**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| Number of page (starting from 1). | 
 **lLX** | **float64**|  | 
 **lLY** | **float64**|  | 
 **uRX** | **float64**|  | 
 **uRY** | **float64**|  | 
 **format** | [**[]string**](string.md)| List of formats for search. | 
 **regex** | **string**| Formats are specified as a regular expression. | 
 **splitRects** | **bool**| Split result fragments (default is true). | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**TextRectsResponse**](TextRectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageTextAnnotations**
> TextAnnotationsResponse GetPageTextAnnotations(ctx, name, pageNumber, optional)
Read document page text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextAnnotationsResponse**](TextAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPageUnderlineAnnotations**
> UnderlineAnnotationsResponse GetPageUnderlineAnnotations(ctx, name, pageNumber, optional)
Read document page underline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**UnderlineAnnotationsResponse**](UnderlineAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPages**
> DocumentPagesResponse GetPages(ctx, name, optional)
Read document pages info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**DocumentPagesResponse**](DocumentPagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPclInStorageToPdf**
> *os.File GetPclInStorageToPdf(ctx, srcPath, optional)
Convert PCL file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pcl) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pcl) | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToDoc**
> *os.File GetPdfInStorageToDoc(ctx, name, optional)
Converts PDF document (located on storage) to DOC format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **addReturnToLineEnd** | **bool**| Add return to line end. | 
 **format** | **string**| Allows to specify .doc or .docx file format. | 
 **imageResolutionX** | **int32**| Image resolution X. | 
 **imageResolutionY** | **int32**| Image resolution Y. | 
 **maxDistanceBetweenTextLines** | **float64**| Max distance between text lines. | 
 **mode** | **string**| Allows to control how a PDF document is converted into a word processing document. | 
 **recognizeBullets** | **bool**| Recognize bullets. | 
 **relativeHorizontalProximity** | **float64**| Relative horizontal proximity. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToEpub**
> *os.File GetPdfInStorageToEpub(ctx, name, optional)
Converts PDF document (located on storage) to EPUB format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **contentRecognitionMode** | **string**| ?roperty tunes conversion for this or that desirable method of recognition of content. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToHtml**
> *os.File GetPdfInStorageToHtml(ctx, name, optional)
Converts PDF document (located on storage) to Html format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **additionalMarginWidthInPoints** | **int32**| Defines width of margin that will be forcibly left around that output HTML-areas. | 
 **compressSvgGraphicsIfAny** | **bool**| The flag that indicates whether found SVG graphics(if any) will be compressed(zipped) into SVGZ format during saving. | 
 **convertMarkedContentToLayers** | **bool**| If attribute ConvertMarkedContentToLayers set to true then an all elements inside a PDF marked content (layer) will be put into an HTML div with \&quot;data-pdflayer\&quot; attribute specifying a layer name. This layer name will be extracted from optional properties of PDF marked content. If this attribute is false (by default) then no any layers will be created from PDF marked content. | 
 **defaultFontName** | **string**| Specifies the name of an installed font which is used to substitute any document font that is not embedded and not installed in the system. If null then default substitution font is used. | 
 **documentType** | **string**| Result document type. | 
 **fixedLayout** | **bool**| The value indicating whether that HTML is created as fixed layout. | 
 **imageResolution** | **int32**| Resolution for image rendering. | 
 **minimalLineWidth** | **int32**| This attribute sets minimal width of graphic path line. If thickness of line is less than 1px Adobe Acrobat rounds it to this value. So this attribute can be used to emulate this behavior for HTML browsers. | 
 **preventGlyphsGrouping** | **bool**| This attribute switch on the mode when text glyphs will not be grouped into words and strings This mode allows to keep maximum precision during positioning of glyphs on the page and it can be used for conversion documents with music notes or glyphs that should be placed separately each other. This parameter will be applied to document only when the value of FixedLayout attribute is true. | 
 **splitCssIntoPages** | **bool**| When multipage-mode selected(i.e &#39;SplitIntoPages&#39; is &#39;true&#39;), then this attribute defines whether should be created separate CSS-file for each result HTML page. | 
 **splitIntoPages** | **bool**| The flag that indicates whether each page of source document will be converted into it&#39;s own target HTML document, i.e whether result HTML will be splitted into several HTML-pages. | 
 **useZOrder** | **bool**| If attribute UseZORder set to true, graphics and text are added to resultant HTML document accordingly Z-order in original PDF document. If this attribute is false all graphics is put as single layer which may cause some unnecessary effects for overlapped objects. | 
 **antialiasingProcessing** | **string**| The parameter defines required antialiasing measures during conversion of compound background images from PDF to HTML. | 
 **cssClassNamesPrefix** | **string**| When PDFtoHTML converter generates result CSSs, CSS class names (something like \&quot;.stl_01 {}\&quot; ... \&quot;.stl_NN {}) are generated and used in result CSS. This property allows forcibly set class name prefix. | 
 **explicitListOfSavedPages** | [**[]int32**](int32.md)| With this property You can explicitely define what pages of document should be converted. Pages in this list must have 1-based numbers. I.e. valid numbers of pages must be taken from range (1...[NumberOfPagesInConvertedDocument]) Order of appearing of pages in this list does not affect their order in result HTML page(s) - in result pages allways will go in order in which they are present in source PDF. | 
 **fontEncodingStrategy** | **string**| Defines encoding special rule to tune PDF decoding for current document. | 
 **fontSavingMode** | **string**| Defines font saving mode that will be used during saving of PDF to desirable format. | 
 **htmlMarkupGenerationMode** | **string**| Sometimes specific reqirments to generation of HTML markup are present. This parameter defines HTML preparing modes that can be used during conversion of PDF to HTML to match such specific requirments. | 
 **lettersPositioningMethod** | **string**| The mode of positioning of letters in words in result HTML. | 
 **pagesFlowTypeDependsOnViewersScreenSize** | **bool**| If attribute &#39;SplitOnPages&#x3D;false&#39;, than whole HTML representing all input PDF pages will be put into one big result HTML file. This flag defines whether result HTML will be generated in such way that flow of areas that represent PDF pages in result HTML will depend on screen resolution of viewer. | 
 **partsEmbeddingMode** | **string**| It defines whether referenced files (HTML, Fonts,Images, CSSes) will be embedded into main HTML file or will be generated as apart binary entities. | 
 **rasterImagesSavingMode** | **string**| Converted PDF can contain raster images This parameter defines how they should be handled during conversion of PDF to HTML. | 
 **removeEmptyAreasOnTopAndBottom** | **bool**| Defines whether in created HTML will be removed top and bottom empty area without any content (if any). | 
 **saveShadowedTextsAsTransparentTexts** | **bool**| Pdf can contain texts that are shadowed by another elements (f.e. by images) but can be selected to clipboard in Acrobat Reader (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML to mimic behaviour of Acrobat Reader (othervise such texts are usually saved as hidden, not available for copying to clipboard). | 
 **saveTransparentTexts** | **bool**| Pdf can contain transparent texts that can be selected to clipboard (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML. | 
 **specialFolderForAllImages** | **string**| The path to directory to which must be saved any images if they are encountered during saving of document as HTML. If parameter is empty or null then image files(if any) wil be saved together with other files linked to HTML It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **specialFolderForSvgImages** | **string**| The path to directory to which must be saved only SVG-images if they are encountered during saving of document as HTML. If parameter is empty or null then SVG files(if any) wil be saved together with other image-files (near to output file) or in special folder for images (if it specified in SpecialImagesFolderIfAny option). It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **trySaveTextUnderliningAndStrikeoutingInCss** | **bool**| PDF itself does not contain underlining markers for texts. It emulated with line situated under text. This option allows converter try guess that this or that line is a text&#39;s underlining and put this info into CSS instead of drawing of underlining graphically. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToLaTeX**
> *os.File GetPdfInStorageToLaTeX(ctx, name, optional)
Converts PDF document (located on storage) to LaTeX format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pagesCount** | **int32**| Pages count. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToMobiXml**
> *os.File GetPdfInStorageToMobiXml(ctx, name, optional)
Converts PDF document (located on storage) to MOBIXML format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToPdfA**
> *os.File GetPdfInStorageToPdfA(ctx, name, type_, optional)
Converts PDF document (located on storage) to PdfA format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **type_** | **string**| Type of PdfA format. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **type_** | **string**| Type of PdfA format. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToPptx**
> *os.File GetPdfInStorageToPptx(ctx, name, optional)
Converts PDF document (located on storage) to PPTX format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **separateImages** | **bool**| Separate images. | 
 **slidesAsImages** | **bool**| Slides as images. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToSvg**
> *os.File GetPdfInStorageToSvg(ctx, name, optional)
Converts PDF document (located on storage) to SVG format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **compressOutputToZipArchive** | **bool**| Specifies whether output will be created as one zip-archive. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToTiff**
> *os.File GetPdfInStorageToTiff(ctx, name, optional)
Converts PDF document (located on storage) to TIFF format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **brightness** | **float64**| Image brightness. | 
 **compression** | **string**| Tiff compression. Possible values are: LZW, CCITT4, CCITT3, RLE, None. | 
 **colorDepth** | **string**| Image color depth. Possible valuse are: Default, Format8bpp, Format4bpp, Format1bpp. | 
 **leftMargin** | **int32**| Left image margin. | 
 **rightMargin** | **int32**| Right image margin. | 
 **topMargin** | **int32**| Top image margin. | 
 **bottomMargin** | **int32**| Bottom image margin. | 
 **orientation** | **string**| Image orientation. Possible values are: None, Landscape, Portait. | 
 **skipBlankPages** | **bool**| Skip blank pages flag. | 
 **width** | **int32**| Image width. | 
 **height** | **int32**| Image height. | 
 **xResolution** | **int32**| Horizontal resolution. | 
 **yResolution** | **int32**| Vertical resolution. | 
 **pageIndex** | **int32**| Start page to export. | 
 **pageCount** | **int32**| Number of pages to export. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToXls**
> *os.File GetPdfInStorageToXls(ctx, name, optional)
Converts PDF document (located on storage) to XLS format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **scaleFactor** | **float64**| Scale factor | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToXml**
> *os.File GetPdfInStorageToXml(ctx, name, optional)
Converts PDF document (located on storage) to XML format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPdfInStorageToXps**
> *os.File GetPdfInStorageToXps(ctx, name, optional)
Converts PDF document (located on storage) to XPS format and returns resulting file in response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolyLineAnnotation**
> PolyLineAnnotationResponse GetPolyLineAnnotation(ctx, name, annotationId, optional)
Read document page polyline annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolyLineAnnotationResponse**](PolyLineAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolygonAnnotation**
> PolygonAnnotationResponse GetPolygonAnnotation(ctx, name, annotationId, optional)
Read document page polygon annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolygonAnnotationResponse**](PolygonAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPopupAnnotation**
> PopupAnnotationResponse GetPopupAnnotation(ctx, name, annotationId, optional)
Read document page popup annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationResponse**](PopupAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPsInStorageToPdf**
> *os.File GetPsInStorageToPdf(ctx, srcPath, optional)
Convert PS file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.ps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.ps) | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSquareAnnotation**
> SquareAnnotationResponse GetSquareAnnotation(ctx, name, annotationId, optional)
Read document page square annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquareAnnotationResponse**](SquareAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSquigglyAnnotation**
> SquigglyAnnotationResponse GetSquigglyAnnotation(ctx, name, annotationId, optional)
Read document page squiggly annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquigglyAnnotationResponse**](SquigglyAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStrikeOutAnnotation**
> StrikeOutAnnotationResponse GetStrikeOutAnnotation(ctx, name, annotationId, optional)
Read document page StrikeOut annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StrikeOutAnnotationResponse**](StrikeOutAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSvgInStorageToPdf**
> *os.File GetSvgInStorageToPdf(ctx, srcPath, optional)
Convert SVG file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.svg) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.svg) | 
 **adjustPageSize** | **bool**| Adjust page size | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetText**
> TextRectsResponse GetText(ctx, name, lLX, lLY, uRX, uRY, optional)
Read document text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **lLX** | **float64**|  | 
  **lLY** | **float64**|  | 
  **uRX** | **float64**|  | 
  **uRY** | **float64**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **lLX** | **float64**|  | 
 **lLY** | **float64**|  | 
 **uRX** | **float64**|  | 
 **uRY** | **float64**|  | 
 **format** | [**[]string**](string.md)| List of formats for search. | 
 **regex** | **string**| Formats are specified as a regular expression. | 
 **splitRects** | **bool**| Split result fragments (default is true). | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**TextRectsResponse**](TextRectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTextAnnotation**
> TextAnnotationResponse GetTextAnnotation(ctx, name, annotationId, optional)
Read document page text annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextAnnotationResponse**](TextAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnderlineAnnotation**
> UnderlineAnnotationResponse GetUnderlineAnnotation(ctx, name, annotationId, optional)
Read document page underline annotation by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**UnderlineAnnotationResponse**](UnderlineAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerifySignature**
> SignatureVerifyResponse GetVerifySignature(ctx, name, signName, optional)
Verify signature document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **signName** | **string**| Sign name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **signName** | **string**| Sign name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SignatureVerifyResponse**](SignatureVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebInStorageToPdf**
> *os.File GetWebInStorageToPdf(ctx, url, optional)
Convert web page to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **url** | **string**| Source url | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string**| Source url | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWordsPerPage**
> WordCountResponse GetWordsPerPage(ctx, name, optional)
Get number of words per document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**WordCountResponse**](WordCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXfaPdfInStorageToAcroForm**
> *os.File GetXfaPdfInStorageToAcroForm(ctx, name, optional)
Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and returns resulting file response content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXmlInStorageToPdf**
> *os.File GetXmlInStorageToPdf(ctx, srcPath, optional)
Convert XML file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **xslFilePath** | **string**| Full XSL source filename (ex. /folder1/folder2/template.xsl) | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXpsInStorageToPdf**
> *os.File GetXpsInStorageToPdf(ctx, srcPath, optional)
Convert XPS file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetXslFoInStorageToPdf**
> *os.File GetXslFoInStorageToPdf(ctx, srcPath, optional)
Convert XslFo file (located on storage) to PDF format and return resulting file in response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xslfo) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xslfo) | 
 **storage** | **string**| The document storage. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAppendDocument**
> DocumentResponse PostAppendDocument(ctx, name, optional)
Append document to existing one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The original document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The original document name. | 
 **appendDocument** | [**AppendDocument**](AppendDocument.md)| with the append document data. | 
 **appendFile** | **string**| Append file server path. | 
 **startPage** | **int32**| Appending start page. | [default to 0]
 **endPage** | **int32**| Appending end page. | [default to 0]
 **storage** | **string**| The documents storage. | 
 **folder** | **string**| The original document folder. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCreateField**
> AsposeResponse PostCreateField(ctx, name, page, optional)
Create field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **page** | **int32**| Document page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **page** | **int32**| Document page number. | 
 **field** | [**Field**](Field.md)| with the field data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDocumentTextReplace**
> TextReplaceResponse PostDocumentTextReplace(ctx, name, textReplace, optional)
Document's replace text method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
  **textReplace** | [**TextReplaceListRequest**](TextReplaceListRequest.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **textReplace** | [**TextReplaceListRequest**](TextReplaceListRequest.md)|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**TextReplaceResponse**](TextReplaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInsertImage**
> AsposeResponse PostInsertImage(ctx, name, pageNumber, llx, lly, urx, ury, optional)
Insert image to document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **llx** | **float64**| Coordinate lower left X. | 
  **lly** | **float64**| Coordinate lower left Y. | 
  **urx** | **float64**| Coordinate upper right X. | 
  **ury** | **float64**| Coordinate upper right Y. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **llx** | **float64**| Coordinate lower left X. | 
 **lly** | **float64**| Coordinate lower left Y. | 
 **urx** | **float64**| Coordinate upper right X. | 
 **ury** | **float64**| Coordinate upper right Y. | 
 **imageFilePath** | **string**| Path to image file if specified. Request content is used otherwise. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **image** | ***os.File**| Image file. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMovePage**
> AsposeResponse PostMovePage(ctx, name, pageNumber, newIndex, optional)
Move page to new position.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **newIndex** | **int32**| The new page position/index. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **newIndex** | **int32**| The new page position/index. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOptimizeDocument**
> AsposeResponse PostOptimizeDocument(ctx, name, optional)
Optimize document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **options** | [**OptimizeOptions**](OptimizeOptions.md)| The optimization options. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageCaretAnnotations**
> AsposeResponse PostPageCaretAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page caret annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]CaretAnnotation**](CaretAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]CaretAnnotation**](CaretAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageCircleAnnotations**
> AsposeResponse PostPageCircleAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page circle annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]CircleAnnotation**](CircleAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]CircleAnnotation**](CircleAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageFreeTextAnnotations**
> AsposeResponse PostPageFreeTextAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page free text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]FreeTextAnnotation**](FreeTextAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]FreeTextAnnotation**](FreeTextAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageHighlightAnnotations**
> AsposeResponse PostPageHighlightAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page highlight annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]HighlightAnnotation**](HighlightAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]HighlightAnnotation**](HighlightAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageInkAnnotations**
> AsposeResponse PostPageInkAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page ink annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]InkAnnotation**](InkAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]InkAnnotation**](InkAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageLineAnnotations**
> AsposeResponse PostPageLineAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page line annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]LineAnnotation**](LineAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]LineAnnotation**](LineAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageLinkAnnotations**
> AsposeResponse PostPageLinkAnnotations(ctx, name, pageNumber, links, optional)
Add document page link annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **links** | [**[]LinkAnnotation**](LinkAnnotation.md)| Array of link anotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **links** | [**[]LinkAnnotation**](LinkAnnotation.md)| Array of link anotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagePolyLineAnnotations**
> AsposeResponse PostPagePolyLineAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page polyline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]PolyLineAnnotation**](PolyLineAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]PolyLineAnnotation**](PolyLineAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagePolygonAnnotations**
> AsposeResponse PostPagePolygonAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page polygon annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]PolygonAnnotation**](PolygonAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]PolygonAnnotation**](PolygonAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageSquareAnnotations**
> AsposeResponse PostPageSquareAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page square annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]SquareAnnotation**](SquareAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]SquareAnnotation**](SquareAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageSquigglyAnnotations**
> AsposeResponse PostPageSquigglyAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page squiggly annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]SquigglyAnnotation**](SquigglyAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]SquigglyAnnotation**](SquigglyAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageStrikeOutAnnotations**
> AsposeResponse PostPageStrikeOutAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page StrikeOut annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]StrikeOutAnnotation**](StrikeOutAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]StrikeOutAnnotation**](StrikeOutAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageTextAnnotations**
> AsposeResponse PostPageTextAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page text annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]TextAnnotation**](TextAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]TextAnnotation**](TextAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageTextReplace**
> TextReplaceResponse PostPageTextReplace(ctx, name, pageNumber, textReplaceListRequest, optional)
Page's replace text method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
  **pageNumber** | **int32**|  | 
  **textReplaceListRequest** | [**TextReplaceListRequest**](TextReplaceListRequest.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **pageNumber** | **int32**|  | 
 **textReplaceListRequest** | [**TextReplaceListRequest**](TextReplaceListRequest.md)|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**TextReplaceResponse**](TextReplaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPageUnderlineAnnotations**
> AsposeResponse PostPageUnderlineAnnotations(ctx, name, pageNumber, annotations, optional)
Add document page underline annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **annotations** | [**[]UnderlineAnnotation**](UnderlineAnnotation.md)| The array of annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **annotations** | [**[]UnderlineAnnotation**](UnderlineAnnotation.md)| The array of annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPopupAnnotation**
> AsposeResponse PostPopupAnnotation(ctx, name, annotationId, annotation, optional)
Add document popup annotations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The parent annotation ID. | 
  **annotation** | [**PopupAnnotation**](PopupAnnotation.md)| The annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The parent annotation ID. | 
 **annotation** | [**PopupAnnotation**](PopupAnnotation.md)| The annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSignDocument**
> AsposeResponse PostSignDocument(ctx, name, optional)
Sign document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **signature** | [**Signature**](Signature.md)| Signature object containing signature data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSignPage**
> AsposeResponse PostSignPage(ctx, name, pageNumber, optional)
Sign page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **signature** | [**Signature**](Signature.md)| Signature object containing signature data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSplitDocument**
> SplitResultResponse PostSplitDocument(ctx, name, optional)
Split document to parts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| Document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Document name. | 
 **format** | **string**| Resulting documents format. | 
 **from** | **int32**| Start page if defined. | 
 **to** | **int32**| End page if defined. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SplitResultResponse**](SplitResultResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAddNewPage**
> DocumentPagesResponse PutAddNewPage(ctx, name, optional)
Add new page to end of the document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**DocumentPagesResponse**](DocumentPagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAddText**
> AsposeResponse PutAddText(ctx, name, pageNumber, optional)
Add text to PDF document page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| Number of page (starting from 1). | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| Number of page (starting from 1). | 
 **paragraph** | [**Paragraph**](Paragraph.md)| Paragraph data. | 
 **folder** | **string**| Document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCaretAnnotation**
> CaretAnnotationResponse PutCaretAnnotation(ctx, name, annotationId, annotation, optional)
Replace document caret annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**CaretAnnotation**](CaretAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**CaretAnnotation**](CaretAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CaretAnnotationResponse**](CaretAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCircleAnnotation**
> CircleAnnotationResponse PutCircleAnnotation(ctx, name, annotationId, annotation, optional)
Replace document circle annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**CircleAnnotation**](CircleAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**CircleAnnotation**](CircleAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**CircleAnnotationResponse**](CircleAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCreate**
> AsposeResponse PutCreate(ctx, path, file, optional)
Upload a specific file 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext | 
  **file** | ***os.File**| File to upload | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext | 
 **file** | ***os.File**| File to upload | 
 **versionId** | **string**| Source file&#39;s version | 
 **storage** | **string**| User&#39;s storage name | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCreateDocument**
> DocumentResponse PutCreateDocument(ctx, name, optional)
Create empty document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The new document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The new document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The new document folder. | 

### Return type

[**DocumentResponse**](DocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutEpubInStorageToPdf**
> AsposeResponse PutEpubInStorageToPdf(ctx, name, srcPath, optional)
Convert EPUB file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.epub) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.epub) | 
 **storage** | **string**| The document storage. | 
 **dstFolder** | **string**| The destination document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFieldsFlatten**
> AsposeResponse PutFieldsFlatten(ctx, name, optional)
Flatten form fields in document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFreeTextAnnotation**
> FreeTextAnnotationResponse PutFreeTextAnnotation(ctx, name, annotationId, annotation, optional)
Replace document free text annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**FreeTextAnnotation**](FreeTextAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**FreeTextAnnotation**](FreeTextAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FreeTextAnnotationResponse**](FreeTextAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutHighlightAnnotation**
> HighlightAnnotationResponse PutHighlightAnnotation(ctx, name, annotationId, annotation, optional)
Replace document highlight annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**HighlightAnnotation**](HighlightAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**HighlightAnnotation**](HighlightAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**HighlightAnnotationResponse**](HighlightAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutHtmlInStorageToPdf**
> AsposeResponse PutHtmlInStorageToPdf(ctx, name, srcPath, optional)
Convert HTML file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.zip) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.zip) | 
 **htmlFileName** | **string**| Name of HTML file in ZIP. | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageExtractAsGif**
> AsposeResponse PutImageExtractAsGif(ctx, name, imageId, optional)
Extract document image in GIF format to folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageExtractAsJpeg**
> AsposeResponse PutImageExtractAsJpeg(ctx, name, imageId, optional)
Extract document image in JPEG format to folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageExtractAsPng**
> AsposeResponse PutImageExtractAsPng(ctx, name, imageId, optional)
Extract document image in PNG format to folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageExtractAsTiff**
> AsposeResponse PutImageExtractAsTiff(ctx, name, imageId, optional)
Extract document image in TIFF format to folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| Image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| Image ID. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImageInStorageToPdf**
> AsposeResponse PutImageInStorageToPdf(ctx, name, imageTemplates, optional)
Convert image file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageTemplates** | [**ImageTemplatesRequest**](ImageTemplatesRequest.md)| Image templates | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageTemplates** | [**ImageTemplatesRequest**](ImageTemplatesRequest.md)| Image templates | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImagesExtractAsGif**
> AsposeResponse PutImagesExtractAsGif(ctx, name, pageNumber, optional)
Extract document images in GIF format to folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImagesExtractAsJpeg**
> AsposeResponse PutImagesExtractAsJpeg(ctx, name, pageNumber, optional)
Extract document images in JPEG format to folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**|  | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImagesExtractAsPng**
> AsposeResponse PutImagesExtractAsPng(ctx, name, pageNumber, optional)
Extract document images in PNG format to folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutImagesExtractAsTiff**
> AsposeResponse PutImagesExtractAsTiff(ctx, name, pageNumber, optional)
Extract document images in TIFF format to folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **destFolder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutInkAnnotation**
> InkAnnotationResponse PutInkAnnotation(ctx, name, annotationId, annotation, optional)
Replace document ink annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**InkAnnotation**](InkAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**InkAnnotation**](InkAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**InkAnnotationResponse**](InkAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutLaTeXInStorageToPdf**
> AsposeResponse PutLaTeXInStorageToPdf(ctx, name, srcPath, optional)
Convert LaTeX file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.tex) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.tex) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutLineAnnotation**
> LineAnnotationResponse PutLineAnnotation(ctx, name, annotationId, annotation, optional)
Replace document line annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**LineAnnotation**](LineAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**LineAnnotation**](LineAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LineAnnotationResponse**](LineAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutLinkAnnotation**
> LinkAnnotationResponse PutLinkAnnotation(ctx, name, linkId, link, optional)
Replace document page link annotations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **linkId** | **string**| The link ID. | 
  **link** | [**LinkAnnotation**](LinkAnnotation.md)| Link anotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **linkId** | **string**| The link ID. | 
 **link** | [**LinkAnnotation**](LinkAnnotation.md)| Link anotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**LinkAnnotationResponse**](LinkAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMergeDocuments**
> *os.File PutMergeDocuments(ctx, name, optional)
Merge a list of documents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| Resulting documen name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Resulting documen name. | 
 **mergeDocuments** | [**MergeDocuments**](MergeDocuments.md)| with a list of documents. | 
 **storage** | **string**| Resulting document storage. | 
 **folder** | **string**| Resulting document folder. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMhtInStorageToPdf**
> AsposeResponse PutMhtInStorageToPdf(ctx, name, srcPath, optional)
Convert MHT file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.mht) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.mht) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageAddStamp**
> AsposeResponse PutPageAddStamp(ctx, name, pageNumber, stamp, optional)
Add page stamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **stamp** | [**Stamp**](Stamp.md)| with data. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **stamp** | [**Stamp**](Stamp.md)| with data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToBmp**
> AsposeResponse PutPageConvertToBmp(ctx, name, pageNumber, outPath, optional)
Convert document page to bmp image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToEmf**
> AsposeResponse PutPageConvertToEmf(ctx, name, pageNumber, outPath, optional)
Convert document page to emf image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToGif**
> AsposeResponse PutPageConvertToGif(ctx, name, pageNumber, outPath, optional)
Convert document page to gif image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToJpeg**
> AsposeResponse PutPageConvertToJpeg(ctx, name, pageNumber, outPath, optional)
Convert document page to Jpeg image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToPng**
> AsposeResponse PutPageConvertToPng(ctx, name, pageNumber, outPath, optional)
Convert document page to png image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPageConvertToTiff**
> AsposeResponse PutPageConvertToTiff(ctx, name, pageNumber, outPath, optional)
Convert document page to Tiff image and upload resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **pageNumber** | **int32**| The page number. | 
  **outPath** | **string**| The out path of result image. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **pageNumber** | **int32**| The page number. | 
 **outPath** | **string**| The out path of result image. | 
 **width** | **int32**| The converted image width. | 
 **height** | **int32**| The converted image height. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPclInStorageToPdf**
> AsposeResponse PutPclInStorageToPdf(ctx, name, srcPath, optional)
Convert PCL file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pcl) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.pcl) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToDoc**
> AsposeResponse PutPdfInRequestToDoc(ctx, outPath, optional)
Converts PDF document (in request content) to DOC format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **addReturnToLineEnd** | **bool**| Add return to line end. | 
 **format** | **string**| Allows to specify .doc or .docx file format. | 
 **imageResolutionX** | **int32**| Image resolution X. | 
 **imageResolutionY** | **int32**| Image resolution Y. | 
 **maxDistanceBetweenTextLines** | **float64**| Max distance between text lines. | 
 **mode** | **string**| Allows to control how a PDF document is converted into a word processing document. | 
 **recognizeBullets** | **bool**| Recognize bullets. | 
 **relativeHorizontalProximity** | **float64**| Relative horizontal proximity. | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToEpub**
> AsposeResponse PutPdfInRequestToEpub(ctx, outPath, optional)
Converts PDF document (in request content) to EPUB format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.epub) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.epub) | 
 **contentRecognitionMode** | **string**| ?roperty tunes conversion for this or that desirable method of recognition of content. | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToHtml**
> AsposeResponse PutPdfInRequestToHtml(ctx, outPath, optional)
Converts PDF document (in request content) to Html format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.html) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.html) | 
 **additionalMarginWidthInPoints** | **int32**| Defines width of margin that will be forcibly left around that output HTML-areas. | 
 **compressSvgGraphicsIfAny** | **bool**| The flag that indicates whether found SVG graphics(if any) will be compressed(zipped) into SVGZ format during saving. | 
 **convertMarkedContentToLayers** | **bool**| If attribute ConvertMarkedContentToLayers set to true then an all elements inside a PDF marked content (layer) will be put into an HTML div with \&quot;data-pdflayer\&quot; attribute specifying a layer name. This layer name will be extracted from optional properties of PDF marked content. If this attribute is false (by default) then no any layers will be created from PDF marked content. | 
 **defaultFontName** | **string**| Specifies the name of an installed font which is used to substitute any document font that is not embedded and not installed in the system. If null then default substitution font is used. | 
 **documentType** | **string**| Result document type. | 
 **fixedLayout** | **bool**| The value indicating whether that HTML is created as fixed layout. | 
 **imageResolution** | **int32**| Resolution for image rendering. | 
 **minimalLineWidth** | **int32**| This attribute sets minimal width of graphic path line. If thickness of line is less than 1px Adobe Acrobat rounds it to this value. So this attribute can be used to emulate this behavior for HTML browsers. | 
 **preventGlyphsGrouping** | **bool**| This attribute switch on the mode when text glyphs will not be grouped into words and strings This mode allows to keep maximum precision during positioning of glyphs on the page and it can be used for conversion documents with music notes or glyphs that should be placed separately each other. This parameter will be applied to document only when the value of FixedLayout attribute is true. | 
 **splitCssIntoPages** | **bool**| When multipage-mode selected(i.e &#39;SplitIntoPages&#39; is &#39;true&#39;), then this attribute defines whether should be created separate CSS-file for each result HTML page. | 
 **splitIntoPages** | **bool**| The flag that indicates whether each page of source document will be converted into it&#39;s own target HTML document, i.e whether result HTML will be splitted into several HTML-pages. | 
 **useZOrder** | **bool**| If attribute UseZORder set to true, graphics and text are added to resultant HTML document accordingly Z-order in original PDF document. If this attribute is false all graphics is put as single layer which may cause some unnecessary effects for overlapped objects. | 
 **antialiasingProcessing** | **string**| The parameter defines required antialiasing measures during conversion of compound background images from PDF to HTML. | 
 **cssClassNamesPrefix** | **string**| When PDFtoHTML converter generates result CSSs, CSS class names (something like \&quot;.stl_01 {}\&quot; ... \&quot;.stl_NN {}) are generated and used in result CSS. This property allows forcibly set class name prefix. | 
 **explicitListOfSavedPages** | [**[]int32**](int32.md)| With this property You can explicitely define what pages of document should be converted. Pages in this list must have 1-based numbers. I.e. valid numbers of pages must be taken from range (1...[NumberOfPagesInConvertedDocument]) Order of appearing of pages in this list does not affect their order in result HTML page(s) - in result pages allways will go in order in which they are present in source PDF. | 
 **fontEncodingStrategy** | **string**| Defines encoding special rule to tune PDF decoding for current document. | 
 **fontSavingMode** | **string**| Defines font saving mode that will be used during saving of PDF to desirable format. | 
 **htmlMarkupGenerationMode** | **string**| Sometimes specific reqirments to generation of HTML markup are present. This parameter defines HTML preparing modes that can be used during conversion of PDF to HTML to match such specific requirments. | 
 **lettersPositioningMethod** | **string**| The mode of positioning of letters in words in result HTML. | 
 **pagesFlowTypeDependsOnViewersScreenSize** | **bool**| If attribute &#39;SplitOnPages&#x3D;false&#39;, than whole HTML representing all input PDF pages will be put into one big result HTML file. This flag defines whether result HTML will be generated in such way that flow of areas that represent PDF pages in result HTML will depend on screen resolution of viewer. | 
 **partsEmbeddingMode** | **string**| It defines whether referenced files (HTML, Fonts,Images, CSSes) will be embedded into main HTML file or will be generated as apart binary entities. | 
 **rasterImagesSavingMode** | **string**| Converted PDF can contain raster images This parameter defines how they should be handled during conversion of PDF to HTML. | 
 **removeEmptyAreasOnTopAndBottom** | **bool**| Defines whether in created HTML will be removed top and bottom empty area without any content (if any). | 
 **saveShadowedTextsAsTransparentTexts** | **bool**| Pdf can contain texts that are shadowed by another elements (f.e. by images) but can be selected to clipboard in Acrobat Reader (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML to mimic behaviour of Acrobat Reader (othervise such texts are usually saved as hidden, not available for copying to clipboard). | 
 **saveTransparentTexts** | **bool**| Pdf can contain transparent texts that can be selected to clipboard (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML. | 
 **specialFolderForAllImages** | **string**| The path to directory to which must be saved any images if they are encountered during saving of document as HTML. If parameter is empty or null then image files(if any) wil be saved together with other files linked to HTML It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **specialFolderForSvgImages** | **string**| The path to directory to which must be saved only SVG-images if they are encountered during saving of document as HTML. If parameter is empty or null then SVG files(if any) wil be saved together with other image-files (near to output file) or in special folder for images (if it specified in SpecialImagesFolderIfAny option). It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **trySaveTextUnderliningAndStrikeoutingInCss** | **bool**| PDF itself does not contain underlining markers for texts. It emulated with line situated under text. This option allows converter try guess that this or that line is a text&#39;s underlining and put this info into CSS instead of drawing of underlining graphically. | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToLaTeX**
> AsposeResponse PutPdfInRequestToLaTeX(ctx, outPath, optional)
Converts PDF document (in request content) to LaTeX format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tex) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tex) | 
 **pagesCount** | **int32**| Pages count. | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToMobiXml**
> AsposeResponse PutPdfInRequestToMobiXml(ctx, outPath, optional)
Converts PDF document (in request content) to MOBIXML format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.mobixml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.mobixml) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToPdfA**
> AsposeResponse PutPdfInRequestToPdfA(ctx, outPath, type_, optional)
Converts PDF document (in request content) to PdfA format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
  **type_** | **string**| Type of PdfA format. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **type_** | **string**| Type of PdfA format. | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToPptx**
> AsposeResponse PutPdfInRequestToPptx(ctx, outPath, optional)
Converts PDF document (in request content) to PPTX format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pptx) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pptx) | 
 **separateImages** | **bool**| Separate images. | 
 **slidesAsImages** | **bool**| Slides as images. | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToSvg**
> AsposeResponse PutPdfInRequestToSvg(ctx, outPath, optional)
Converts PDF document (in request content) to SVG format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.svg) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.svg) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToTiff**
> AsposeResponse PutPdfInRequestToTiff(ctx, outPath, optional)
Converts PDF document (in request content) to TIFF format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tiff) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tiff) | 
 **brightness** | **float64**| Image brightness. | 
 **compression** | **string**| Tiff compression. Possible values are: LZW, CCITT4, CCITT3, RLE, None. | 
 **colorDepth** | **string**| Image color depth. Possible valuse are: Default, Format8bpp, Format4bpp, Format1bpp. | 
 **leftMargin** | **int32**| Left image margin. | 
 **rightMargin** | **int32**| Right image margin. | 
 **topMargin** | **int32**| Top image margin. | 
 **bottomMargin** | **int32**| Bottom image margin. | 
 **orientation** | **string**| Image orientation. Possible values are: None, Landscape, Portait. | 
 **skipBlankPages** | **bool**| Skip blank pages flag. | 
 **width** | **int32**| Image width. | 
 **height** | **int32**| Image height. | 
 **xResolution** | **int32**| Horizontal resolution. | 
 **yResolution** | **int32**| Vertical resolution. | 
 **pageIndex** | **int32**| Start page to export. | 
 **pageCount** | **int32**| Number of pages to export. | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToXls**
> AsposeResponse PutPdfInRequestToXls(ctx, outPath, optional)
Converts PDF document (in request content) to XLS format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xls) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xls) | 
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **scaleFactor** | **float64**| Scale factor | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToXml**
> AsposeResponse PutPdfInRequestToXml(ctx, outPath, optional)
Converts PDF document (in request content) to XML format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xml) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInRequestToXps**
> AsposeResponse PutPdfInRequestToXps(ctx, outPath, optional)
Converts PDF document (in request content) to XPS format and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xps) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToDoc**
> AsposeResponse PutPdfInStorageToDoc(ctx, name, outPath, optional)
Converts PDF document (located on storage) to DOC format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.doc) | 
 **addReturnToLineEnd** | **bool**| Add return to line end. | 
 **format** | **string**| Allows to specify .doc or .docx file format. | 
 **imageResolutionX** | **int32**| Image resolution X. | 
 **imageResolutionY** | **int32**| Image resolution Y. | 
 **maxDistanceBetweenTextLines** | **float64**| Max distance between text lines. | 
 **mode** | **string**| Allows to control how a PDF document is converted into a word processing document. | 
 **recognizeBullets** | **bool**| Recognize bullets. | 
 **relativeHorizontalProximity** | **float64**| Relative horizontal proximity. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToEpub**
> AsposeResponse PutPdfInStorageToEpub(ctx, name, outPath, optional)
Converts PDF document (located on storage) to EPUB format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.epub) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.epub) | 
 **contentRecognitionMode** | **string**| ?roperty tunes conversion for this or that desirable method of recognition of content. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToHtml**
> AsposeResponse PutPdfInStorageToHtml(ctx, name, outPath, optional)
Converts PDF document (located on storage) to Html format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.html) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.html) | 
 **additionalMarginWidthInPoints** | **int32**| Defines width of margin that will be forcibly left around that output HTML-areas. | 
 **compressSvgGraphicsIfAny** | **bool**| The flag that indicates whether found SVG graphics(if any) will be compressed(zipped) into SVGZ format during saving. | 
 **convertMarkedContentToLayers** | **bool**| If attribute ConvertMarkedContentToLayers set to true then an all elements inside a PDF marked content (layer) will be put into an HTML div with \&quot;data-pdflayer\&quot; attribute specifying a layer name. This layer name will be extracted from optional properties of PDF marked content. If this attribute is false (by default) then no any layers will be created from PDF marked content. | 
 **defaultFontName** | **string**| Specifies the name of an installed font which is used to substitute any document font that is not embedded and not installed in the system. If null then default substitution font is used. | 
 **documentType** | **string**| Result document type. | 
 **fixedLayout** | **bool**| The value indicating whether that HTML is created as fixed layout. | 
 **imageResolution** | **int32**| Resolution for image rendering. | 
 **minimalLineWidth** | **int32**| This attribute sets minimal width of graphic path line. If thickness of line is less than 1px Adobe Acrobat rounds it to this value. So this attribute can be used to emulate this behavior for HTML browsers. | 
 **preventGlyphsGrouping** | **bool**| This attribute switch on the mode when text glyphs will not be grouped into words and strings This mode allows to keep maximum precision during positioning of glyphs on the page and it can be used for conversion documents with music notes or glyphs that should be placed separately each other. This parameter will be applied to document only when the value of FixedLayout attribute is true. | 
 **splitCssIntoPages** | **bool**| When multipage-mode selected(i.e &#39;SplitIntoPages&#39; is &#39;true&#39;), then this attribute defines whether should be created separate CSS-file for each result HTML page. | 
 **splitIntoPages** | **bool**| The flag that indicates whether each page of source document will be converted into it&#39;s own target HTML document, i.e whether result HTML will be splitted into several HTML-pages. | 
 **useZOrder** | **bool**| If attribute UseZORder set to true, graphics and text are added to resultant HTML document accordingly Z-order in original PDF document. If this attribute is false all graphics is put as single layer which may cause some unnecessary effects for overlapped objects. | 
 **antialiasingProcessing** | **string**| The parameter defines required antialiasing measures during conversion of compound background images from PDF to HTML. | 
 **cssClassNamesPrefix** | **string**| When PDFtoHTML converter generates result CSSs, CSS class names (something like \&quot;.stl_01 {}\&quot; ... \&quot;.stl_NN {}) are generated and used in result CSS. This property allows forcibly set class name prefix. | 
 **explicitListOfSavedPages** | [**[]int32**](int32.md)| With this property You can explicitely define what pages of document should be converted. Pages in this list must have 1-based numbers. I.e. valid numbers of pages must be taken from range (1...[NumberOfPagesInConvertedDocument]) Order of appearing of pages in this list does not affect their order in result HTML page(s) - in result pages allways will go in order in which they are present in source PDF. | 
 **fontEncodingStrategy** | **string**| Defines encoding special rule to tune PDF decoding for current document. | 
 **fontSavingMode** | **string**| Defines font saving mode that will be used during saving of PDF to desirable format. | 
 **htmlMarkupGenerationMode** | **string**| Sometimes specific reqirments to generation of HTML markup are present. This parameter defines HTML preparing modes that can be used during conversion of PDF to HTML to match such specific requirments. | 
 **lettersPositioningMethod** | **string**| The mode of positioning of letters in words in result HTML. | 
 **pagesFlowTypeDependsOnViewersScreenSize** | **bool**| If attribute &#39;SplitOnPages&#x3D;false&#39;, than whole HTML representing all input PDF pages will be put into one big result HTML file. This flag defines whether result HTML will be generated in such way that flow of areas that represent PDF pages in result HTML will depend on screen resolution of viewer. | 
 **partsEmbeddingMode** | **string**| It defines whether referenced files (HTML, Fonts,Images, CSSes) will be embedded into main HTML file or will be generated as apart binary entities. | 
 **rasterImagesSavingMode** | **string**| Converted PDF can contain raster images This parameter defines how they should be handled during conversion of PDF to HTML. | 
 **removeEmptyAreasOnTopAndBottom** | **bool**| Defines whether in created HTML will be removed top and bottom empty area without any content (if any). | 
 **saveShadowedTextsAsTransparentTexts** | **bool**| Pdf can contain texts that are shadowed by another elements (f.e. by images) but can be selected to clipboard in Acrobat Reader (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML to mimic behaviour of Acrobat Reader (othervise such texts are usually saved as hidden, not available for copying to clipboard). | 
 **saveTransparentTexts** | **bool**| Pdf can contain transparent texts that can be selected to clipboard (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML. | 
 **specialFolderForAllImages** | **string**| The path to directory to which must be saved any images if they are encountered during saving of document as HTML. If parameter is empty or null then image files(if any) wil be saved together with other files linked to HTML It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **specialFolderForSvgImages** | **string**| The path to directory to which must be saved only SVG-images if they are encountered during saving of document as HTML. If parameter is empty or null then SVG files(if any) wil be saved together with other image-files (near to output file) or in special folder for images (if it specified in SpecialImagesFolderIfAny option). It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file. | 
 **trySaveTextUnderliningAndStrikeoutingInCss** | **bool**| PDF itself does not contain underlining markers for texts. It emulated with line situated under text. This option allows converter try guess that this or that line is a text&#39;s underlining and put this info into CSS instead of drawing of underlining graphically. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToLaTeX**
> AsposeResponse PutPdfInStorageToLaTeX(ctx, name, outPath, optional)
Converts PDF document (located on storage) to LaTeX format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tex) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tex) | 
 **pagesCount** | **int32**| Pages count. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToMobiXml**
> AsposeResponse PutPdfInStorageToMobiXml(ctx, name, outPath, optional)
Converts PDF document (located on storage) to MOBIXML format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.mobixml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.mobixml) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToPdfA**
> AsposeResponse PutPdfInStorageToPdfA(ctx, name, outPath, type_, optional)
Converts PDF document (located on storage) to PdfA format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
  **type_** | **string**| Type of PdfA format. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **type_** | **string**| Type of PdfA format. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToPptx**
> AsposeResponse PutPdfInStorageToPptx(ctx, name, outPath, optional)
Converts PDF document (located on storage) to PPTX format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pptx) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pptx) | 
 **separateImages** | **bool**| Separate images. | 
 **slidesAsImages** | **bool**| Slides as images. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToSvg**
> AsposeResponse PutPdfInStorageToSvg(ctx, name, outPath, optional)
Converts PDF document (located on storage) to SVG format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.svg) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.svg) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToTiff**
> AsposeResponse PutPdfInStorageToTiff(ctx, name, outPath, optional)
Converts PDF document (located on storage) to TIFF format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tiff) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.tiff) | 
 **brightness** | **float64**| Image brightness. | 
 **compression** | **string**| Tiff compression. Possible values are: LZW, CCITT4, CCITT3, RLE, None. | 
 **colorDepth** | **string**| Image color depth. Possible valuse are: Default, Format8bpp, Format4bpp, Format1bpp. | 
 **leftMargin** | **int32**| Left image margin. | 
 **rightMargin** | **int32**| Right image margin. | 
 **topMargin** | **int32**| Top image margin. | 
 **bottomMargin** | **int32**| Bottom image margin. | 
 **orientation** | **string**| Image orientation. Possible values are: None, Landscape, Portait. | 
 **skipBlankPages** | **bool**| Skip blank pages flag. | 
 **width** | **int32**| Image width. | 
 **height** | **int32**| Image height. | 
 **xResolution** | **int32**| Horizontal resolution. | 
 **yResolution** | **int32**| Vertical resolution. | 
 **pageIndex** | **int32**| Start page to export. | 
 **pageCount** | **int32**| Number of pages to export. | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToXls**
> AsposeResponse PutPdfInStorageToXls(ctx, name, outPath, optional)
Converts PDF document (located on storage) to XLS format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xls) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xls) | 
 **insertBlankColumnAtFirst** | **bool**| Insert blank column at first | 
 **minimizeTheNumberOfWorksheets** | **bool**| Minimize the number of worksheets | 
 **scaleFactor** | **float64**| Scale factor | 
 **uniformWorksheets** | **bool**| Uniform worksheets | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToXml**
> AsposeResponse PutPdfInStorageToXml(ctx, name, outPath, optional)
Converts PDF document (located on storage) to XML format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xml) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPdfInStorageToXps**
> AsposeResponse PutPdfInStorageToXps(ctx, name, outPath, optional)
Converts PDF document (located on storage) to XPS format and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.xps) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPolyLineAnnotation**
> PolyLineAnnotationResponse PutPolyLineAnnotation(ctx, name, annotationId, annotation, optional)
Replace document polyline annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**PolyLineAnnotation**](PolyLineAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**PolyLineAnnotation**](PolyLineAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolyLineAnnotationResponse**](PolyLineAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPolygonAnnotation**
> PolygonAnnotationResponse PutPolygonAnnotation(ctx, name, annotationId, annotation, optional)
Replace document polygon annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**PolygonAnnotation**](PolygonAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**PolygonAnnotation**](PolygonAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PolygonAnnotationResponse**](PolygonAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPopupAnnotation**
> PopupAnnotationResponse PutPopupAnnotation(ctx, name, annotationId, annotation, optional)
Replace document popup annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**PopupAnnotation**](PopupAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**PopupAnnotation**](PopupAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**PopupAnnotationResponse**](PopupAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPrivileges**
> AsposeResponse PutPrivileges(ctx, name, optional)
Update privilege document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **privileges** | [**DocumentPrivilege**](DocumentPrivilege.md)| Document privileges.  | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPsInStorageToPdf**
> AsposeResponse PutPsInStorageToPdf(ctx, name, srcPath, optional)
Convert PS file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.ps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.ps) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutReplaceImage**
> ImageResponse PutReplaceImage(ctx, name, imageId, optional)
Replace document image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **imageId** | **string**| The image ID. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **imageId** | **string**| The image ID. | 
 **imageFilePath** | **string**| Path to image file if specified. Request content is used otherwise. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **image** | ***os.File**| Image file. | 

### Return type

[**ImageResponse**](ImageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSearchableDocument**
> AsposeResponse PutSearchableDocument(ctx, name, optional)
Create searchable PDF document. Generate OCR layer for images in input PDF document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 
 **lang** | **string**| language for OCR engine. Possible values: eng, ara, bel, ben, bul, ces, dan, deu, ell, fin, fra, heb, hin, ind, isl, ita, jpn, kor, nld, nor, pol, por, ron, rus, spa, swe, tha, tur, ukr, vie, chi_sim, chi_tra or thier combination e.g. eng,rus  | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSetProperty**
> DocumentPropertyResponse PutSetProperty(ctx, name, propertyName, value, optional)
Add/update document property.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
  **propertyName** | **string**|  | 
  **value** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **propertyName** | **string**|  | 
 **value** | **string**|  | 
 **storage** | **string**|  | 
 **folder** | **string**|  | 

### Return type

[**DocumentPropertyResponse**](DocumentPropertyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSquareAnnotation**
> SquareAnnotationResponse PutSquareAnnotation(ctx, name, annotationId, annotation, optional)
Replace document square annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**SquareAnnotation**](SquareAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**SquareAnnotation**](SquareAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquareAnnotationResponse**](SquareAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSquigglyAnnotation**
> SquigglyAnnotationResponse PutSquigglyAnnotation(ctx, name, annotationId, annotation, optional)
Replace document squiggly annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**SquigglyAnnotation**](SquigglyAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**SquigglyAnnotation**](SquigglyAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**SquigglyAnnotationResponse**](SquigglyAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutStrikeOutAnnotation**
> StrikeOutAnnotationResponse PutStrikeOutAnnotation(ctx, name, annotationId, annotation, optional)
Replace document StrikeOut annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**StrikeOutAnnotation**](StrikeOutAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**StrikeOutAnnotation**](StrikeOutAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**StrikeOutAnnotationResponse**](StrikeOutAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSvgInStorageToPdf**
> AsposeResponse PutSvgInStorageToPdf(ctx, name, srcPath, optional)
Convert SVG file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.svg) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.svg) | 
 **adjustPageSize** | **bool**| Adjust page size | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutTextAnnotation**
> TextAnnotationResponse PutTextAnnotation(ctx, name, annotationId, annotation, optional)
Replace document text annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**TextAnnotation**](TextAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**TextAnnotation**](TextAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**TextAnnotationResponse**](TextAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUnderlineAnnotation**
> UnderlineAnnotationResponse PutUnderlineAnnotation(ctx, name, annotationId, annotation, optional)
Replace document underline annotation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **annotationId** | **string**| The annotation ID. | 
  **annotation** | [**UnderlineAnnotation**](UnderlineAnnotation.md)| Annotation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **annotationId** | **string**| The annotation ID. | 
 **annotation** | [**UnderlineAnnotation**](UnderlineAnnotation.md)| Annotation. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**UnderlineAnnotationResponse**](UnderlineAnnotationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUpdateField**
> FieldResponse PutUpdateField(ctx, name, fieldName, optional)
Update field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **fieldName** | **string**| The name of a field to be updated. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fieldName** | **string**| The name of a field to be updated. | 
 **field** | [**Field**](Field.md)| with the field data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FieldResponse**](FieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUpdateFields**
> FieldsResponse PutUpdateFields(ctx, name, optional)
Update fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **fields** | [**Fields**](Fields.md)| with the fields data. | 
 **storage** | **string**| The document storage. | 
 **folder** | **string**| The document folder. | 

### Return type

[**FieldsResponse**](FieldsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutWebInStorageToPdf**
> AsposeResponse PutWebInStorageToPdf(ctx, name, url, optional)
Convert web page to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **url** | **string**| Source url | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **url** | **string**| Source url | 
 **height** | **float64**| Page height | 
 **width** | **float64**| Page width | 
 **isLandscape** | **bool**| Is page landscaped | 
 **marginLeft** | **float64**| Page margin left | 
 **marginBottom** | **float64**| Page margin bottom | 
 **marginRight** | **float64**| Page margin right | 
 **marginTop** | **float64**| Page margin top | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXfaPdfInRequestToAcroForm**
> AsposeResponse PutXfaPdfInRequestToAcroForm(ctx, outPath, optional)
Converts PDF document which contatins XFA form (in request content) to PDF with AcroForm and uploads resulting file to storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **storage** | **string**| The document storage. | 
 **file** | ***os.File**| A file to be converted. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXfaPdfInStorageToAcroForm**
> AsposeResponse PutXfaPdfInStorageToAcroForm(ctx, name, outPath, optional)
Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and uploads resulting file to storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **outPath** | **string**| Full resulting filename (ex. /folder1/folder2/result.pdf) | 
 **folder** | **string**| The document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXmlInStorageToPdf**
> AsposeResponse PutXmlInStorageToPdf(ctx, name, srcPath, optional)
Convert XML file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xml) | 
 **xslFilePath** | **string**| Full XSL source filename (ex. /folder1/folder2/template.xsl) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXpsInStorageToPdf**
> AsposeResponse PutXpsInStorageToPdf(ctx, name, srcPath, optional)
Convert XPS file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xps) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutXslFoInStorageToPdf**
> AsposeResponse PutXslFoInStorageToPdf(ctx, name, srcPath, optional)
Convert XslFo file (located on storage) to PDF format and upload resulting file to storage. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The document name. | 
  **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xpsfo) | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The document name. | 
 **srcPath** | **string**| Full source filename (ex. /folder1/folder2/template.xpsfo) | 
 **dstFolder** | **string**| The destination document folder. | 
 **storage** | **string**| The document storage. | 

### Return type

[**AsposeResponse**](AsposeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

