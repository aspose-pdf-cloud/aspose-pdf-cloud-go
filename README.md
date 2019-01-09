# Aspose.PDF Cloud
- API version: 2.0
- Package version: 18.12.0

[Aspose.PDF Cloud](https://products.aspose.cloud/pdf) is a true REST API that enables you to perform a wide range of document processing operations including creation, manipulation, conversion and rendering of PDF documents in the cloud.

Our Cloud SDKs are wrappers around REST API in various programming languages, allowing you to process documents in language of your choice quickly and easily, gaining all benefits of strong types and IDE highlights. This repository contains new generation SDKs for Aspose.PDF Cloud and examples.

These SDKs are now fully supported. If you have any questions, see any bugs or have enhancement request, feel free to reach out to us at [Free Support Forums](https://forum.aspose.cloud/c/pdf).

## Installation
Put the package under your project folder and add the following in import:
```
    "./asposepdfcloud"
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Java code:

```go
func GetDocumentCircleAnnotations() (CircleAnnotationsResponse, *http.Response, error) {
    pdfAPI := NewPdfApiService("AppSid", "AppKey", "")
	name := "PdfWithAnnotations.pdf"	

	args := map[string]interface{} {
		"folder": "path/to/remote/folder",
	}

	return pdfAPI.GetDocumentCircleAnnotations(name, args)
}
```

## Unit Tests
Aspose PDF Cloud SDK includes a suite of unit tests within the "test" subdirectory. These Unit Tests also serves as examples of how to use the Aspose PDF Cloud SDK.

## Licensing
All Aspose.PDF Cloud SDKs are licensed under [MIT License](LICENSE).

## Resources
+ [**Website**](https://www.aspose.cloud)
+ [**Product Home**](https://products.aspose.cloud/pdf/cloud)
+ [**Documentation**](https://docs.aspose.cloud/display/pdfcloud/Home)
+ [**Free Support Forum**](https://forum.aspose.cloud/c/pdf)
+ [**Paid Support Helpdesk**](https://helpdesk.aspose.cloud/)
+ [**Blog**](https://blog.aspose.cloud/category/aspose-products/aspose-pdf-product-family/)

## Documentation for API Endpoints

All URIs are relative to *https://api.aspose.cloud/v2.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PdfApi* | [**DeleteAnnotation**](docs/PdfApi.md#deleteannotation) | **Delete** /pdf/{name}/annotations/{annotationId} | Delete document annotation by ID
*PdfApi* | [**DeleteDocumentAnnotations**](docs/PdfApi.md#deletedocumentannotations) | **Delete** /pdf/{name}/annotations | Delete all annotations from the document
*PdfApi* | [**DeleteDocumentLinkAnnotations**](docs/PdfApi.md#deletedocumentlinkannotations) | **Delete** /pdf/{name}/links | Delete all link annotations from the document
*PdfApi* | [**DeleteField**](docs/PdfApi.md#deletefield) | **Delete** /pdf/{name}/fields/{fieldName} | Delete document field by name.
*PdfApi* | [**DeleteFile**](docs/PdfApi.md#deletefile) | **Delete** /storage/file | Remove a specific file 
*PdfApi* | [**DeleteFolder**](docs/PdfApi.md#deletefolder) | **Delete** /storage/folder | Remove a specific folder 
*PdfApi* | [**DeleteImage**](docs/PdfApi.md#deleteimage) | **Delete** /pdf/{name}/images/{imageId} | Delete image from document page.
*PdfApi* | [**DeleteLinkAnnotation**](docs/PdfApi.md#deletelinkannotation) | **Delete** /pdf/{name}/links/{linkId} | Delete document page link annotation by ID
*PdfApi* | [**DeletePage**](docs/PdfApi.md#deletepage) | **Delete** /pdf/{name}/pages/{pageNumber} | Delete document page by its number.
*PdfApi* | [**DeletePageAnnotations**](docs/PdfApi.md#deletepageannotations) | **Delete** /pdf/{name}/pages/{pageNumber}/annotations | Delete all annotations from the page
*PdfApi* | [**DeletePageLinkAnnotations**](docs/PdfApi.md#deletepagelinkannotations) | **Delete** /pdf/{name}/pages/{pageNumber}/links | Delete all link annotations from the page
*PdfApi* | [**DeleteProperties**](docs/PdfApi.md#deleteproperties) | **Delete** /pdf/{name}/documentproperties | Delete custom document properties.
*PdfApi* | [**DeleteProperty**](docs/PdfApi.md#deleteproperty) | **Delete** /pdf/{name}/documentproperties/{propertyName} | Delete document property.
*PdfApi* | [**GetCaretAnnotation**](docs/PdfApi.md#getcaretannotation) | **Get** /pdf/{name}/annotations/caret/{annotationId} | Read document page caret annotation by ID.
*PdfApi* | [**GetCircleAnnotation**](docs/PdfApi.md#getcircleannotation) | **Get** /pdf/{name}/annotations/circle/{annotationId} | Read document page circle annotation by ID.
*PdfApi* | [**GetDiscUsage**](docs/PdfApi.md#getdiscusage) | **Get** /storage/disc | Check the disk usage of the current account 
*PdfApi* | [**GetDocument**](docs/PdfApi.md#getdocument) | **Get** /pdf/{name} | Read common document info.
*PdfApi* | [**GetDocumentAnnotations**](docs/PdfApi.md#getdocumentannotations) | **Get** /pdf/{name}/annotations | Read documant page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
*PdfApi* | [**GetDocumentAttachmentByIndex**](docs/PdfApi.md#getdocumentattachmentbyindex) | **Get** /pdf/{name}/attachments/{attachmentIndex} | Read document attachment info by its index.
*PdfApi* | [**GetDocumentAttachments**](docs/PdfApi.md#getdocumentattachments) | **Get** /pdf/{name}/attachments | Read document attachments info.
*PdfApi* | [**GetDocumentBookmarks**](docs/PdfApi.md#getdocumentbookmarks) | **Get** /pdf/{name}/bookmarks | Read document bookmark/bookmarks (including children).
*PdfApi* | [**GetDocumentCaretAnnotations**](docs/PdfApi.md#getdocumentcaretannotations) | **Get** /pdf/{name}/annotations/caret | Read document caret annotations.
*PdfApi* | [**GetDocumentCircleAnnotations**](docs/PdfApi.md#getdocumentcircleannotations) | **Get** /pdf/{name}/annotations/circle | Read document circle annotations.
*PdfApi* | [**GetDocumentFreeTextAnnotations**](docs/PdfApi.md#getdocumentfreetextannotations) | **Get** /pdf/{name}/annotations/freetext | Read document free text annotations.
*PdfApi* | [**GetDocumentHighlightAnnotations**](docs/PdfApi.md#getdocumenthighlightannotations) | **Get** /pdf/{name}/annotations/highlight | Read document highlight annotations.
*PdfApi* | [**GetDocumentInkAnnotations**](docs/PdfApi.md#getdocumentinkannotations) | **Get** /pdf/{name}/annotations/ink | Read document ink annotations.
*PdfApi* | [**GetDocumentLineAnnotations**](docs/PdfApi.md#getdocumentlineannotations) | **Get** /pdf/{name}/annotations/line | Read document line annotations.
*PdfApi* | [**GetDocumentPolyLineAnnotations**](docs/PdfApi.md#getdocumentpolylineannotations) | **Get** /pdf/{name}/annotations/polyline | Read document polyline annotations.
*PdfApi* | [**GetDocumentPolygonAnnotations**](docs/PdfApi.md#getdocumentpolygonannotations) | **Get** /pdf/{name}/annotations/polygon | Read document polygon annotations.
*PdfApi* | [**GetDocumentPopupAnnotations**](docs/PdfApi.md#getdocumentpopupannotations) | **Get** /pdf/{name}/annotations/popup | Read document popup annotations.
*PdfApi* | [**GetDocumentPopupAnnotationsByParent**](docs/PdfApi.md#getdocumentpopupannotationsbyparent) | **Get** /pdf/{name}/annotations/{annotationId}/popup | Read document popup annotations by parent id.
*PdfApi* | [**GetDocumentProperties**](docs/PdfApi.md#getdocumentproperties) | **Get** /pdf/{name}/documentproperties | Read document properties.
*PdfApi* | [**GetDocumentProperty**](docs/PdfApi.md#getdocumentproperty) | **Get** /pdf/{name}/documentproperties/{propertyName} | Read document property by name.
*PdfApi* | [**GetDocumentSquareAnnotations**](docs/PdfApi.md#getdocumentsquareannotations) | **Get** /pdf/{name}/annotations/square | Read document square annotations.
*PdfApi* | [**GetDocumentSquigglyAnnotations**](docs/PdfApi.md#getdocumentsquigglyannotations) | **Get** /pdf/{name}/annotations/squiggly | Read document squiggly annotations.
*PdfApi* | [**GetDocumentStrikeOutAnnotations**](docs/PdfApi.md#getdocumentstrikeoutannotations) | **Get** /pdf/{name}/annotations/strikeout | Read document StrikeOut annotations.
*PdfApi* | [**GetDocumentTextAnnotations**](docs/PdfApi.md#getdocumenttextannotations) | **Get** /pdf/{name}/annotations/text | Read document text annotations.
*PdfApi* | [**GetDocumentUnderlineAnnotations**](docs/PdfApi.md#getdocumentunderlineannotations) | **Get** /pdf/{name}/annotations/underline | Read document underline annotations.
*PdfApi* | [**GetDownload**](docs/PdfApi.md#getdownload) | **Get** /storage/file | Download a specific file 
*PdfApi* | [**GetDownloadDocumentAttachmentByIndex**](docs/PdfApi.md#getdownloaddocumentattachmentbyindex) | **Get** /pdf/{name}/attachments/{attachmentIndex}/download | Download document attachment content by its index.
*PdfApi* | [**GetEpubInStorageToPdf**](docs/PdfApi.md#getepubinstoragetopdf) | **Get** /pdf/create/epub | Convert EPUB file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetField**](docs/PdfApi.md#getfield) | **Get** /pdf/{name}/fields/{fieldName} | Get document field by name.
*PdfApi* | [**GetFields**](docs/PdfApi.md#getfields) | **Get** /pdf/{name}/fields | Get document fields.
*PdfApi* | [**GetFreeTextAnnotation**](docs/PdfApi.md#getfreetextannotation) | **Get** /pdf/{name}/annotations/freetext/{annotationId} | Read document page free text annotation by ID.
*PdfApi* | [**GetHighlightAnnotation**](docs/PdfApi.md#gethighlightannotation) | **Get** /pdf/{name}/annotations/highlight/{annotationId} | Read document page highlight annotation by ID.
*PdfApi* | [**GetHtmlInStorageToPdf**](docs/PdfApi.md#gethtmlinstoragetopdf) | **Get** /pdf/create/html | Convert HTML file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetImage**](docs/PdfApi.md#getimage) | **Get** /pdf/{name}/images/{imageId} | Read document image by ID.
*PdfApi* | [**GetImageExtractAsGif**](docs/PdfApi.md#getimageextractasgif) | **Get** /pdf/{name}/images/{imageId}/extract/gif | Extract document image in GIF format
*PdfApi* | [**GetImageExtractAsJpeg**](docs/PdfApi.md#getimageextractasjpeg) | **Get** /pdf/{name}/images/{imageId}/extract/jpeg | Extract document image in JPEG format
*PdfApi* | [**GetImageExtractAsPng**](docs/PdfApi.md#getimageextractaspng) | **Get** /pdf/{name}/images/{imageId}/extract/png | Extract document image in PNG format
*PdfApi* | [**GetImageExtractAsTiff**](docs/PdfApi.md#getimageextractastiff) | **Get** /pdf/{name}/images/{imageId}/extract/tiff | Extract document image in TIFF format
*PdfApi* | [**GetImages**](docs/PdfApi.md#getimages) | **Get** /pdf/{name}/pages/{pageNumber}/images | Read document images.
*PdfApi* | [**GetInkAnnotation**](docs/PdfApi.md#getinkannotation) | **Get** /pdf/{name}/annotations/ink/{annotationId} | Read document page ink annotation by ID.
*PdfApi* | [**GetIsExist**](docs/PdfApi.md#getisexist) | **Get** /storage/exist | Check if a specific file or folder exists
*PdfApi* | [**GetIsStorageExist**](docs/PdfApi.md#getisstorageexist) | **Get** /storage/{name}/exist | Check if storage exists 
*PdfApi* | [**GetLaTeXInStorageToPdf**](docs/PdfApi.md#getlatexinstoragetopdf) | **Get** /pdf/create/latex | Convert LaTeX file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetLineAnnotation**](docs/PdfApi.md#getlineannotation) | **Get** /pdf/{name}/annotations/line/{annotationId} | Read document page line annotation by ID.
*PdfApi* | [**GetLinkAnnotation**](docs/PdfApi.md#getlinkannotation) | **Get** /pdf/{name}/links/{linkId} | Read document link annotation by ID.
*PdfApi* | [**GetListFileVersions**](docs/PdfApi.md#getlistfileversions) | **Get** /storage/version | Get the file&#39;s versions list 
*PdfApi* | [**GetListFiles**](docs/PdfApi.md#getlistfiles) | **Get** /storage/folder | Get the file listing of a specific folder 
*PdfApi* | [**GetMhtInStorageToPdf**](docs/PdfApi.md#getmhtinstoragetopdf) | **Get** /pdf/create/mht | Convert MHT file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetPage**](docs/PdfApi.md#getpage) | **Get** /pdf/{name}/pages/{pageNumber} | Read document page info.
*PdfApi* | [**GetPageAnnotations**](docs/PdfApi.md#getpageannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations | Read documant page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
*PdfApi* | [**GetPageCaretAnnotations**](docs/PdfApi.md#getpagecaretannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/caret | Read document page caret annotations.
*PdfApi* | [**GetPageCircleAnnotations**](docs/PdfApi.md#getpagecircleannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/circle | Read document page circle annotations.
*PdfApi* | [**GetPageConvertToBmp**](docs/PdfApi.md#getpageconverttobmp) | **Get** /pdf/{name}/pages/{pageNumber}/convert/bmp | Convert document page to Bmp image and return resulting file in response.
*PdfApi* | [**GetPageConvertToEmf**](docs/PdfApi.md#getpageconverttoemf) | **Get** /pdf/{name}/pages/{pageNumber}/convert/emf | Convert document page to Emf image and return resulting file in response.
*PdfApi* | [**GetPageConvertToGif**](docs/PdfApi.md#getpageconverttogif) | **Get** /pdf/{name}/pages/{pageNumber}/convert/gif | Convert document page to Gif image and return resulting file in response.
*PdfApi* | [**GetPageConvertToJpeg**](docs/PdfApi.md#getpageconverttojpeg) | **Get** /pdf/{name}/pages/{pageNumber}/convert/jpeg | Convert document page to Jpeg image and return resulting file in response.
*PdfApi* | [**GetPageConvertToPng**](docs/PdfApi.md#getpageconverttopng) | **Get** /pdf/{name}/pages/{pageNumber}/convert/png | Convert document page to Png image and return resulting file in response.
*PdfApi* | [**GetPageConvertToTiff**](docs/PdfApi.md#getpageconverttotiff) | **Get** /pdf/{name}/pages/{pageNumber}/convert/tiff | Convert document page to Tiff image  and return resulting file in response.
*PdfApi* | [**GetPageFreeTextAnnotations**](docs/PdfApi.md#getpagefreetextannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/freetext | Read document page free text annotations.
*PdfApi* | [**GetPageHighlightAnnotations**](docs/PdfApi.md#getpagehighlightannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/highlight | Read document page highlight annotations.
*PdfApi* | [**GetPageInkAnnotations**](docs/PdfApi.md#getpageinkannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/ink | Read document page ink annotations.
*PdfApi* | [**GetPageLineAnnotations**](docs/PdfApi.md#getpagelineannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/line | Read document page line annotations.
*PdfApi* | [**GetPageLinkAnnotation**](docs/PdfApi.md#getpagelinkannotation) | **Get** /pdf/{name}/pages/{pageNumber}/links/{linkId} | Read document page link annotation by ID.
*PdfApi* | [**GetPageLinkAnnotations**](docs/PdfApi.md#getpagelinkannotations) | **Get** /pdf/{name}/pages/{pageNumber}/links | Read document page link annotations.
*PdfApi* | [**GetPagePolyLineAnnotations**](docs/PdfApi.md#getpagepolylineannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/polyline | Read document page polyline annotations.
*PdfApi* | [**GetPagePolygonAnnotations**](docs/PdfApi.md#getpagepolygonannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/polygon | Read document page polygon annotations.
*PdfApi* | [**GetPagePopupAnnotations**](docs/PdfApi.md#getpagepopupannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/popup | Read document page popup annotations.
*PdfApi* | [**GetPageSquareAnnotations**](docs/PdfApi.md#getpagesquareannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/square | Read document page square annotations.
*PdfApi* | [**GetPageSquigglyAnnotations**](docs/PdfApi.md#getpagesquigglyannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/squiggly | Read document page squiggly annotations.
*PdfApi* | [**GetPageStrikeOutAnnotations**](docs/PdfApi.md#getpagestrikeoutannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/strikeout | Read document page StrikeOut annotations.
*PdfApi* | [**GetPageText**](docs/PdfApi.md#getpagetext) | **Get** /pdf/{name}/pages/{pageNumber}/text | Read page text items.
*PdfApi* | [**GetPageTextAnnotations**](docs/PdfApi.md#getpagetextannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/text | Read document page text annotations.
*PdfApi* | [**GetPageUnderlineAnnotations**](docs/PdfApi.md#getpageunderlineannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/underline | Read document page underline annotations.
*PdfApi* | [**GetPages**](docs/PdfApi.md#getpages) | **Get** /pdf/{name}/pages | Read document pages info.
*PdfApi* | [**GetPclInStorageToPdf**](docs/PdfApi.md#getpclinstoragetopdf) | **Get** /pdf/create/pcl | Convert PCL file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetPdfInStorageToDoc**](docs/PdfApi.md#getpdfinstoragetodoc) | **Get** /pdf/{name}/convert/doc | Converts PDF document (located on storage) to DOC format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToEpub**](docs/PdfApi.md#getpdfinstoragetoepub) | **Get** /pdf/{name}/convert/epub | Converts PDF document (located on storage) to EPUB format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToHtml**](docs/PdfApi.md#getpdfinstoragetohtml) | **Get** /pdf/{name}/convert/html | Converts PDF document (located on storage) to Html format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToLaTeX**](docs/PdfApi.md#getpdfinstoragetolatex) | **Get** /pdf/{name}/convert/latex | Converts PDF document (located on storage) to LaTeX format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToMobiXml**](docs/PdfApi.md#getpdfinstoragetomobixml) | **Get** /pdf/{name}/convert/mobixml | Converts PDF document (located on storage) to MOBIXML format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToPdfA**](docs/PdfApi.md#getpdfinstoragetopdfa) | **Get** /pdf/{name}/convert/pdfa | Converts PDF document (located on storage) to PdfA format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToPptx**](docs/PdfApi.md#getpdfinstoragetopptx) | **Get** /pdf/{name}/convert/pptx | Converts PDF document (located on storage) to PPTX format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToSvg**](docs/PdfApi.md#getpdfinstoragetosvg) | **Get** /pdf/{name}/convert/svg | Converts PDF document (located on storage) to SVG format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToTiff**](docs/PdfApi.md#getpdfinstoragetotiff) | **Get** /pdf/{name}/convert/tiff | Converts PDF document (located on storage) to TIFF format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToXls**](docs/PdfApi.md#getpdfinstoragetoxls) | **Get** /pdf/{name}/convert/xls | Converts PDF document (located on storage) to XLS format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToXml**](docs/PdfApi.md#getpdfinstoragetoxml) | **Get** /pdf/{name}/convert/xml | Converts PDF document (located on storage) to XML format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToXps**](docs/PdfApi.md#getpdfinstoragetoxps) | **Get** /pdf/{name}/convert/xps | Converts PDF document (located on storage) to XPS format and returns resulting file in response content
*PdfApi* | [**GetPolyLineAnnotation**](docs/PdfApi.md#getpolylineannotation) | **Get** /pdf/{name}/annotations/polyline/{annotationId} | Read document page polyline annotation by ID.
*PdfApi* | [**GetPolygonAnnotation**](docs/PdfApi.md#getpolygonannotation) | **Get** /pdf/{name}/annotations/polygon/{annotationId} | Read document page polygon annotation by ID.
*PdfApi* | [**GetPopupAnnotation**](docs/PdfApi.md#getpopupannotation) | **Get** /pdf/{name}/annotations/popup/{annotationId} | Read document page popup annotation by ID.
*PdfApi* | [**GetPsInStorageToPdf**](docs/PdfApi.md#getpsinstoragetopdf) | **Get** /pdf/create/ps | Convert PS file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetSquareAnnotation**](docs/PdfApi.md#getsquareannotation) | **Get** /pdf/{name}/annotations/square/{annotationId} | Read document page square annotation by ID.
*PdfApi* | [**GetSquigglyAnnotation**](docs/PdfApi.md#getsquigglyannotation) | **Get** /pdf/{name}/annotations/squiggly/{annotationId} | Read document page squiggly annotation by ID.
*PdfApi* | [**GetStrikeOutAnnotation**](docs/PdfApi.md#getstrikeoutannotation) | **Get** /pdf/{name}/annotations/strikeout/{annotationId} | Read document page StrikeOut annotation by ID.
*PdfApi* | [**GetSvgInStorageToPdf**](docs/PdfApi.md#getsvginstoragetopdf) | **Get** /pdf/create/svg | Convert SVG file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetText**](docs/PdfApi.md#gettext) | **Get** /pdf/{name}/text | Read document text.
*PdfApi* | [**GetTextAnnotation**](docs/PdfApi.md#gettextannotation) | **Get** /pdf/{name}/annotations/text/{annotationId} | Read document page text annotation by ID.
*PdfApi* | [**GetUnderlineAnnotation**](docs/PdfApi.md#getunderlineannotation) | **Get** /pdf/{name}/annotations/underline/{annotationId} | Read document page underline annotation by ID.
*PdfApi* | [**GetVerifySignature**](docs/PdfApi.md#getverifysignature) | **Get** /pdf/{name}/verifySignature | Verify signature document.
*PdfApi* | [**GetWebInStorageToPdf**](docs/PdfApi.md#getwebinstoragetopdf) | **Get** /pdf/create/web | Convert web page to PDF format and return resulting file in response. 
*PdfApi* | [**GetWordsPerPage**](docs/PdfApi.md#getwordsperpage) | **Get** /pdf/{name}/pages/wordCount | Get number of words per document page.
*PdfApi* | [**GetXfaPdfInStorageToAcroForm**](docs/PdfApi.md#getxfapdfinstoragetoacroform) | **Get** /pdf/{name}/convert/xfatoacroform | Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and returns resulting file response content
*PdfApi* | [**GetXmlInStorageToPdf**](docs/PdfApi.md#getxmlinstoragetopdf) | **Get** /pdf/create/xml | Convert XML file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetXpsInStorageToPdf**](docs/PdfApi.md#getxpsinstoragetopdf) | **Get** /pdf/create/xps | Convert XPS file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetXslFoInStorageToPdf**](docs/PdfApi.md#getxslfoinstoragetopdf) | **Get** /pdf/create/xslfo | Convert XslFo file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**PostAppendDocument**](docs/PdfApi.md#postappenddocument) | **Post** /pdf/{name}/appendDocument | Append document to existing one.
*PdfApi* | [**PostCreateField**](docs/PdfApi.md#postcreatefield) | **Post** /pdf/{name}/fields | Create field.
*PdfApi* | [**PostDocumentTextReplace**](docs/PdfApi.md#postdocumenttextreplace) | **Post** /pdf/{name}/text/replace | Document&#39;s replace text method.
*PdfApi* | [**PostFlattenDocument**](docs/PdfApi.md#postflattendocument) | **Post** /pdf/{name}/flatten | Removes all fields from the document and place their values instead.
*PdfApi* | [**PostInsertImage**](docs/PdfApi.md#postinsertimage) | **Post** /pdf/{name}/pages/{pageNumber}/images | Insert image to document page.
*PdfApi* | [**PostMoveFile**](docs/PdfApi.md#postmovefile) | **Post** /storage/file | Move a specific file
*PdfApi* | [**PostMoveFolder**](docs/PdfApi.md#postmovefolder) | **Post** /storage/folder | Move a specific folder 
*PdfApi* | [**PostMovePage**](docs/PdfApi.md#postmovepage) | **Post** /pdf/{name}/pages/{pageNumber}/movePage | Move page to new position.
*PdfApi* | [**PostOptimizeDocument**](docs/PdfApi.md#postoptimizedocument) | **Post** /pdf/{name}/optimize | Optimize document.
*PdfApi* | [**PostPageCaretAnnotations**](docs/PdfApi.md#postpagecaretannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/caret | Add document page caret annotations.
*PdfApi* | [**PostPageCircleAnnotations**](docs/PdfApi.md#postpagecircleannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/circle | Add document page circle annotations.
*PdfApi* | [**PostPageFreeTextAnnotations**](docs/PdfApi.md#postpagefreetextannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/freetext | Add document page free text annotations.
*PdfApi* | [**PostPageHighlightAnnotations**](docs/PdfApi.md#postpagehighlightannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/highlight | Add document page highlight annotations.
*PdfApi* | [**PostPageInkAnnotations**](docs/PdfApi.md#postpageinkannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/ink | Add document page ink annotations.
*PdfApi* | [**PostPageLineAnnotations**](docs/PdfApi.md#postpagelineannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/line | Add document page line annotations.
*PdfApi* | [**PostPageLinkAnnotations**](docs/PdfApi.md#postpagelinkannotations) | **Post** /pdf/{name}/pages/{pageNumber}/links | Add document page link annotations.
*PdfApi* | [**PostPagePolyLineAnnotations**](docs/PdfApi.md#postpagepolylineannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/polyline | Add document page polyline annotations.
*PdfApi* | [**PostPagePolygonAnnotations**](docs/PdfApi.md#postpagepolygonannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/polygon | Add document page polygon annotations.
*PdfApi* | [**PostPageSquareAnnotations**](docs/PdfApi.md#postpagesquareannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/square | Add document page square annotations.
*PdfApi* | [**PostPageSquigglyAnnotations**](docs/PdfApi.md#postpagesquigglyannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/squiggly | Add document page squiggly annotations.
*PdfApi* | [**PostPageStrikeOutAnnotations**](docs/PdfApi.md#postpagestrikeoutannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/strikeout | Add document page StrikeOut annotations.
*PdfApi* | [**PostPageTextAnnotations**](docs/PdfApi.md#postpagetextannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/text | Add document page text annotations.
*PdfApi* | [**PostPageTextReplace**](docs/PdfApi.md#postpagetextreplace) | **Post** /pdf/{name}/pages/{pageNumber}/text/replace | Page&#39;s replace text method.
*PdfApi* | [**PostPageUnderlineAnnotations**](docs/PdfApi.md#postpageunderlineannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/underline | Add document page underline annotations.
*PdfApi* | [**PostPopupAnnotation**](docs/PdfApi.md#postpopupannotation) | **Post** /pdf/{name}/annotations/{annotationId}/popup | Add document popup annotations.
*PdfApi* | [**PostSignDocument**](docs/PdfApi.md#postsigndocument) | **Post** /pdf/{name}/sign | Sign document.
*PdfApi* | [**PostSignPage**](docs/PdfApi.md#postsignpage) | **Post** /pdf/{name}/pages/{pageNumber}/sign | Sign page.
*PdfApi* | [**PostSplitDocument**](docs/PdfApi.md#postsplitdocument) | **Post** /pdf/{name}/split | Split document to parts.
*PdfApi* | [**PutAddNewPage**](docs/PdfApi.md#putaddnewpage) | **Put** /pdf/{name}/pages | Add new page to end of the document.
*PdfApi* | [**PutAddText**](docs/PdfApi.md#putaddtext) | **Put** /pdf/{name}/pages/{pageNumber}/text | Add text to PDF document page.
*PdfApi* | [**PutCaretAnnotation**](docs/PdfApi.md#putcaretannotation) | **Put** /pdf/{name}/annotations/caret/{annotationId} | Replace document caret annotation
*PdfApi* | [**PutCircleAnnotation**](docs/PdfApi.md#putcircleannotation) | **Put** /pdf/{name}/annotations/circle/{annotationId} | Replace document circle annotation
*PdfApi* | [**PutCreate**](docs/PdfApi.md#putcreate) | **Put** /storage/file | Upload a specific file 
*PdfApi* | [**PutCreateDocument**](docs/PdfApi.md#putcreatedocument) | **Put** /pdf/{name} | Create empty document.
*PdfApi* | [**PutCreateFolder**](docs/PdfApi.md#putcreatefolder) | **Put** /storage/folder | Create the folder 
*PdfApi* | [**PutEpubInStorageToPdf**](docs/PdfApi.md#putepubinstoragetopdf) | **Put** /pdf/{name}/create/epub | Convert EPUB file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutFieldsFlatten**](docs/PdfApi.md#putfieldsflatten) | **Put** /pdf/{name}/fields/flatten | Flatten form fields in document.
*PdfApi* | [**PutFreeTextAnnotation**](docs/PdfApi.md#putfreetextannotation) | **Put** /pdf/{name}/annotations/freetext/{annotationId} | Replace document free text annotation
*PdfApi* | [**PutHighlightAnnotation**](docs/PdfApi.md#puthighlightannotation) | **Put** /pdf/{name}/annotations/highlight/{annotationId} | Replace document highlight annotation
*PdfApi* | [**PutHtmlInStorageToPdf**](docs/PdfApi.md#puthtmlinstoragetopdf) | **Put** /pdf/{name}/create/html | Convert HTML file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutImageExtractAsGif**](docs/PdfApi.md#putimageextractasgif) | **Put** /pdf/{name}/images/{imageId}/extract/gif | Extract document image in GIF format to folder
*PdfApi* | [**PutImageExtractAsJpeg**](docs/PdfApi.md#putimageextractasjpeg) | **Put** /pdf/{name}/images/{imageId}/extract/jpeg | Extract document image in JPEG format to folder
*PdfApi* | [**PutImageExtractAsPng**](docs/PdfApi.md#putimageextractaspng) | **Put** /pdf/{name}/images/{imageId}/extract/png | Extract document image in PNG format to folder
*PdfApi* | [**PutImageExtractAsTiff**](docs/PdfApi.md#putimageextractastiff) | **Put** /pdf/{name}/images/{imageId}/extract/tiff | Extract document image in TIFF format to folder
*PdfApi* | [**PutImageInStorageToPdf**](docs/PdfApi.md#putimageinstoragetopdf) | **Put** /pdf/{name}/create/images | Convert image file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutImagesExtractAsGif**](docs/PdfApi.md#putimagesextractasgif) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/gif | Extract document images in GIF format to folder.
*PdfApi* | [**PutImagesExtractAsJpeg**](docs/PdfApi.md#putimagesextractasjpeg) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/jpeg | Extract document images in JPEG format to folder.
*PdfApi* | [**PutImagesExtractAsPng**](docs/PdfApi.md#putimagesextractaspng) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/png | Extract document images in PNG format to folder.
*PdfApi* | [**PutImagesExtractAsTiff**](docs/PdfApi.md#putimagesextractastiff) | **Put** /pdf/{name}/pages/{pageNumber}/images/extract/tiff | Extract document images in TIFF format to folder.
*PdfApi* | [**PutInkAnnotation**](docs/PdfApi.md#putinkannotation) | **Put** /pdf/{name}/annotations/ink/{annotationId} | Replace document ink annotation
*PdfApi* | [**PutLaTeXInStorageToPdf**](docs/PdfApi.md#putlatexinstoragetopdf) | **Put** /pdf/{name}/create/latex | Convert LaTeX file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutLineAnnotation**](docs/PdfApi.md#putlineannotation) | **Put** /pdf/{name}/annotations/line/{annotationId} | Replace document line annotation
*PdfApi* | [**PutLinkAnnotation**](docs/PdfApi.md#putlinkannotation) | **Put** /pdf/{name}/links/{linkId} | Replace document page link annotations
*PdfApi* | [**PutMergeDocuments**](docs/PdfApi.md#putmergedocuments) | **Put** /pdf/{name}/merge | Merge a list of documents.
*PdfApi* | [**PutMhtInStorageToPdf**](docs/PdfApi.md#putmhtinstoragetopdf) | **Put** /pdf/{name}/create/mht | Convert MHT file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutPageAddStamp**](docs/PdfApi.md#putpageaddstamp) | **Put** /pdf/{name}/pages/{pageNumber}/stamp | Add page stamp.
*PdfApi* | [**PutPageConvertToBmp**](docs/PdfApi.md#putpageconverttobmp) | **Put** /pdf/{name}/pages/{pageNumber}/convert/bmp | Convert document page to bmp image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToEmf**](docs/PdfApi.md#putpageconverttoemf) | **Put** /pdf/{name}/pages/{pageNumber}/convert/emf | Convert document page to emf image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToGif**](docs/PdfApi.md#putpageconverttogif) | **Put** /pdf/{name}/pages/{pageNumber}/convert/gif | Convert document page to gif image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToJpeg**](docs/PdfApi.md#putpageconverttojpeg) | **Put** /pdf/{name}/pages/{pageNumber}/convert/jpeg | Convert document page to Jpeg image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToPng**](docs/PdfApi.md#putpageconverttopng) | **Put** /pdf/{name}/pages/{pageNumber}/convert/png | Convert document page to png image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToTiff**](docs/PdfApi.md#putpageconverttotiff) | **Put** /pdf/{name}/pages/{pageNumber}/convert/tiff | Convert document page to Tiff image and upload resulting file to storage.
*PdfApi* | [**PutPclInStorageToPdf**](docs/PdfApi.md#putpclinstoragetopdf) | **Put** /pdf/{name}/create/pcl | Convert PCL file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutPdfInRequestToDoc**](docs/PdfApi.md#putpdfinrequesttodoc) | **Put** /pdf/convert/doc | Converts PDF document (in request content) to DOC format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToEpub**](docs/PdfApi.md#putpdfinrequesttoepub) | **Put** /pdf/convert/epub | Converts PDF document (in request content) to EPUB format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToHtml**](docs/PdfApi.md#putpdfinrequesttohtml) | **Put** /pdf/convert/html | Converts PDF document (in request content) to Html format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToLaTeX**](docs/PdfApi.md#putpdfinrequesttolatex) | **Put** /pdf/convert/latex | Converts PDF document (in request content) to LaTeX format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToMobiXml**](docs/PdfApi.md#putpdfinrequesttomobixml) | **Put** /pdf/convert/mobixml | Converts PDF document (in request content) to MOBIXML format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToPdfA**](docs/PdfApi.md#putpdfinrequesttopdfa) | **Put** /pdf/convert/pdfa | Converts PDF document (in request content) to PdfA format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToPptx**](docs/PdfApi.md#putpdfinrequesttopptx) | **Put** /pdf/convert/pptx | Converts PDF document (in request content) to PPTX format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToSvg**](docs/PdfApi.md#putpdfinrequesttosvg) | **Put** /pdf/convert/svg | Converts PDF document (in request content) to SVG format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToTiff**](docs/PdfApi.md#putpdfinrequesttotiff) | **Put** /pdf/convert/tiff | Converts PDF document (in request content) to TIFF format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToXls**](docs/PdfApi.md#putpdfinrequesttoxls) | **Put** /pdf/convert/xls | Converts PDF document (in request content) to XLS format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToXml**](docs/PdfApi.md#putpdfinrequesttoxml) | **Put** /pdf/convert/xml | Converts PDF document (in request content) to XML format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToXps**](docs/PdfApi.md#putpdfinrequesttoxps) | **Put** /pdf/convert/xps | Converts PDF document (in request content) to XPS format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInStorageToDoc**](docs/PdfApi.md#putpdfinstoragetodoc) | **Put** /pdf/{name}/convert/doc | Converts PDF document (located on storage) to DOC format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToEpub**](docs/PdfApi.md#putpdfinstoragetoepub) | **Put** /pdf/{name}/convert/epub | Converts PDF document (located on storage) to EPUB format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToHtml**](docs/PdfApi.md#putpdfinstoragetohtml) | **Put** /pdf/{name}/convert/html | Converts PDF document (located on storage) to Html format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToLaTeX**](docs/PdfApi.md#putpdfinstoragetolatex) | **Put** /pdf/{name}/convert/latex | Converts PDF document (located on storage) to LaTeX format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToMobiXml**](docs/PdfApi.md#putpdfinstoragetomobixml) | **Put** /pdf/{name}/convert/mobixml | Converts PDF document (located on storage) to MOBIXML format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToPdfA**](docs/PdfApi.md#putpdfinstoragetopdfa) | **Put** /pdf/{name}/convert/pdfa | Converts PDF document (located on storage) to PdfA format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToPptx**](docs/PdfApi.md#putpdfinstoragetopptx) | **Put** /pdf/{name}/convert/pptx | Converts PDF document (located on storage) to PPTX format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToSvg**](docs/PdfApi.md#putpdfinstoragetosvg) | **Put** /pdf/{name}/convert/svg | Converts PDF document (located on storage) to SVG format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToTiff**](docs/PdfApi.md#putpdfinstoragetotiff) | **Put** /pdf/{name}/convert/tiff | Converts PDF document (located on storage) to TIFF format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToXls**](docs/PdfApi.md#putpdfinstoragetoxls) | **Put** /pdf/{name}/convert/xls | Converts PDF document (located on storage) to XLS format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToXml**](docs/PdfApi.md#putpdfinstoragetoxml) | **Put** /pdf/{name}/convert/xml | Converts PDF document (located on storage) to XML format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToXps**](docs/PdfApi.md#putpdfinstoragetoxps) | **Put** /pdf/{name}/convert/xps | Converts PDF document (located on storage) to XPS format and uploads resulting file to storage
*PdfApi* | [**PutPolyLineAnnotation**](docs/PdfApi.md#putpolylineannotation) | **Put** /pdf/{name}/annotations/polyline/{annotationId} | Replace document polyline annotation
*PdfApi* | [**PutPolygonAnnotation**](docs/PdfApi.md#putpolygonannotation) | **Put** /pdf/{name}/annotations/polygon/{annotationId} | Replace document polygon annotation
*PdfApi* | [**PutPopupAnnotation**](docs/PdfApi.md#putpopupannotation) | **Put** /pdf/{name}/annotations/popup/{annotationId} | Replace document popup annotation
*PdfApi* | [**PutPrivileges**](docs/PdfApi.md#putprivileges) | **Put** /pdf/{name}/privileges | Update privilege document.
*PdfApi* | [**PutPsInStorageToPdf**](docs/PdfApi.md#putpsinstoragetopdf) | **Put** /pdf/{name}/create/ps | Convert PS file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutReplaceImage**](docs/PdfApi.md#putreplaceimage) | **Put** /pdf/{name}/images/{imageId} | Replace document image.
*PdfApi* | [**PutSearchableDocument**](docs/PdfApi.md#putsearchabledocument) | **Put** /pdf/{name}/ocr | Create searchable PDF document. Generate OCR layer for images in input PDF document.
*PdfApi* | [**PutSetProperty**](docs/PdfApi.md#putsetproperty) | **Put** /pdf/{name}/documentproperties/{propertyName} | Add/update document property.
*PdfApi* | [**PutSquareAnnotation**](docs/PdfApi.md#putsquareannotation) | **Put** /pdf/{name}/annotations/square/{annotationId} | Replace document square annotation
*PdfApi* | [**PutSquigglyAnnotation**](docs/PdfApi.md#putsquigglyannotation) | **Put** /pdf/{name}/annotations/squiggly/{annotationId} | Replace document squiggly annotation
*PdfApi* | [**PutStrikeOutAnnotation**](docs/PdfApi.md#putstrikeoutannotation) | **Put** /pdf/{name}/annotations/strikeout/{annotationId} | Replace document StrikeOut annotation
*PdfApi* | [**PutSvgInStorageToPdf**](docs/PdfApi.md#putsvginstoragetopdf) | **Put** /pdf/{name}/create/svg | Convert SVG file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutTextAnnotation**](docs/PdfApi.md#puttextannotation) | **Put** /pdf/{name}/annotations/text/{annotationId} | Replace document text annotation
*PdfApi* | [**PutUnderlineAnnotation**](docs/PdfApi.md#putunderlineannotation) | **Put** /pdf/{name}/annotations/underline/{annotationId} | Replace document underline annotation
*PdfApi* | [**PutUpdateField**](docs/PdfApi.md#putupdatefield) | **Put** /pdf/{name}/fields/{fieldName} | Update field.
*PdfApi* | [**PutUpdateFields**](docs/PdfApi.md#putupdatefields) | **Put** /pdf/{name}/fields | Update fields.
*PdfApi* | [**PutWebInStorageToPdf**](docs/PdfApi.md#putwebinstoragetopdf) | **Put** /pdf/{name}/create/web | Convert web page to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutXfaPdfInRequestToAcroForm**](docs/PdfApi.md#putxfapdfinrequesttoacroform) | **Put** /pdf/convert/xfatoacroform | Converts PDF document which contatins XFA form (in request content) to PDF with AcroForm and uploads resulting file to storage.
*PdfApi* | [**PutXfaPdfInStorageToAcroForm**](docs/PdfApi.md#putxfapdfinstoragetoacroform) | **Put** /pdf/{name}/convert/xfatoacroform | Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and uploads resulting file to storage
*PdfApi* | [**PutXmlInStorageToPdf**](docs/PdfApi.md#putxmlinstoragetopdf) | **Put** /pdf/{name}/create/xml | Convert XML file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutXpsInStorageToPdf**](docs/PdfApi.md#putxpsinstoragetopdf) | **Put** /pdf/{name}/create/xps | Convert XPS file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutXslFoInStorageToPdf**](docs/PdfApi.md#putxslfoinstoragetopdf) | **Put** /pdf/{name}/create/xslfo | Convert XslFo file (located on storage) to PDF format and upload resulting file to storage. 


## Documentation For Models

 - [AnnotationFlags](docs/AnnotationFlags.md)
 - [AnnotationState](docs/AnnotationState.md)
 - [AnnotationType](docs/AnnotationType.md)
 - [AntialiasingProcessingType](docs/AntialiasingProcessingType.md)
 - [AppendDocument](docs/AppendDocument.md)
 - [AsposeResponse](docs/AsposeResponse.md)
 - [CapStyle](docs/CapStyle.md)
 - [CaptionPosition](docs/CaptionPosition.md)
 - [CaretSymbol](docs/CaretSymbol.md)
 - [Color](docs/Color.md)
 - [ColorDepth](docs/ColorDepth.md)
 - [CompressionType](docs/CompressionType.md)
 - [DiscUsage](docs/DiscUsage.md)
 - [DocFormat](docs/DocFormat.md)
 - [DocRecognitionMode](docs/DocRecognitionMode.md)
 - [DocumentPrivilege](docs/DocumentPrivilege.md)
 - [EpubRecognitionMode](docs/EpubRecognitionMode.md)
 - [FieldType](docs/FieldType.md)
 - [File](docs/File.md)
 - [FileExist](docs/FileExist.md)
 - [FontEncodingRules](docs/FontEncodingRules.md)
 - [FontSavingModes](docs/FontSavingModes.md)
 - [FontStyles](docs/FontStyles.md)
 - [FreeTextIntent](docs/FreeTextIntent.md)
 - [HorizontalAlignment](docs/HorizontalAlignment.md)
 - [HtmlDocumentType](docs/HtmlDocumentType.md)
 - [HtmlMarkupGenerationModes](docs/HtmlMarkupGenerationModes.md)
 - [ImageSrcType](docs/ImageSrcType.md)
 - [ImageTemplate](docs/ImageTemplate.md)
 - [ImageTemplatesRequest](docs/ImageTemplatesRequest.md)
 - [Justification](docs/Justification.md)
 - [LettersPositioningMethods](docs/LettersPositioningMethods.md)
 - [LineEnding](docs/LineEnding.md)
 - [LineIntent](docs/LineIntent.md)
 - [LineSpacing](docs/LineSpacing.md)
 - [Link](docs/Link.md)
 - [LinkActionType](docs/LinkActionType.md)
 - [LinkElement](docs/LinkElement.md)
 - [LinkHighlightingMode](docs/LinkHighlightingMode.md)
 - [MarginInfo](docs/MarginInfo.md)
 - [MergeDocuments](docs/MergeDocuments.md)
 - [OptimizeOptions](docs/OptimizeOptions.md)
 - [PageWordCount](docs/PageWordCount.md)
 - [Paragraph](docs/Paragraph.md)
 - [PartsEmbeddingModes](docs/PartsEmbeddingModes.md)
 - [PdfAType](docs/PdfAType.md)
 - [Point](docs/Point.md)
 - [PolyIntent](docs/PolyIntent.md)
 - [RasterImagesSavingModes](docs/RasterImagesSavingModes.md)
 - [Rectangle](docs/Rectangle.md)
 - [Rotation](docs/Rotation.md)
 - [Segment](docs/Segment.md)
 - [ShapeType](docs/ShapeType.md)
 - [Signature](docs/Signature.md)
 - [SignatureType](docs/SignatureType.md)
 - [SplitResult](docs/SplitResult.md)
 - [Stamp](docs/Stamp.md)
 - [StampType](docs/StampType.md)
 - [TextHorizontalAlignment](docs/TextHorizontalAlignment.md)
 - [TextIcon](docs/TextIcon.md)
 - [TextLine](docs/TextLine.md)
 - [TextRect](docs/TextRect.md)
 - [TextRects](docs/TextRects.md)
 - [TextReplace](docs/TextReplace.md)
 - [TextReplaceListRequest](docs/TextReplaceListRequest.md)
 - [TextState](docs/TextState.md)
 - [TextStyle](docs/TextStyle.md)
 - [VerticalAlignment](docs/VerticalAlignment.md)
 - [WordCount](docs/WordCount.md)
 - [WrapMode](docs/WrapMode.md)
 - [Annotation](docs/Annotation.md)
 - [AnnotationsInfo](docs/AnnotationsInfo.md)
 - [AnnotationsInfoResponse](docs/AnnotationsInfoResponse.md)
 - [Attachment](docs/Attachment.md)
 - [AttachmentResponse](docs/AttachmentResponse.md)
 - [Attachments](docs/Attachments.md)
 - [AttachmentsResponse](docs/AttachmentsResponse.md)
 - [CaretAnnotationResponse](docs/CaretAnnotationResponse.md)
 - [CaretAnnotations](docs/CaretAnnotations.md)
 - [CaretAnnotationsResponse](docs/CaretAnnotationsResponse.md)
 - [CircleAnnotationResponse](docs/CircleAnnotationResponse.md)
 - [CircleAnnotations](docs/CircleAnnotations.md)
 - [CircleAnnotationsResponse](docs/CircleAnnotationsResponse.md)
 - [DiscUsageResponse](docs/DiscUsageResponse.md)
 - [Document](docs/Document.md)
 - [DocumentPageResponse](docs/DocumentPageResponse.md)
 - [DocumentPagesResponse](docs/DocumentPagesResponse.md)
 - [DocumentProperties](docs/DocumentProperties.md)
 - [DocumentPropertiesResponse](docs/DocumentPropertiesResponse.md)
 - [DocumentProperty](docs/DocumentProperty.md)
 - [DocumentPropertyResponse](docs/DocumentPropertyResponse.md)
 - [DocumentResponse](docs/DocumentResponse.md)
 - [Field](docs/Field.md)
 - [FieldResponse](docs/FieldResponse.md)
 - [Fields](docs/Fields.md)
 - [FieldsResponse](docs/FieldsResponse.md)
 - [FileExistResponse](docs/FileExistResponse.md)
 - [FileVersion](docs/FileVersion.md)
 - [FileVersionsResponse](docs/FileVersionsResponse.md)
 - [FilesResponse](docs/FilesResponse.md)
 - [FreeTextAnnotationResponse](docs/FreeTextAnnotationResponse.md)
 - [FreeTextAnnotations](docs/FreeTextAnnotations.md)
 - [FreeTextAnnotationsResponse](docs/FreeTextAnnotationsResponse.md)
 - [HighlightAnnotationResponse](docs/HighlightAnnotationResponse.md)
 - [HighlightAnnotations](docs/HighlightAnnotations.md)
 - [HighlightAnnotationsResponse](docs/HighlightAnnotationsResponse.md)
 - [Image](docs/Image.md)
 - [ImageResponse](docs/ImageResponse.md)
 - [Images](docs/Images.md)
 - [ImagesResponse](docs/ImagesResponse.md)
 - [InkAnnotationResponse](docs/InkAnnotationResponse.md)
 - [InkAnnotations](docs/InkAnnotations.md)
 - [InkAnnotationsResponse](docs/InkAnnotationsResponse.md)
 - [LineAnnotationResponse](docs/LineAnnotationResponse.md)
 - [LineAnnotations](docs/LineAnnotations.md)
 - [LineAnnotationsResponse](docs/LineAnnotationsResponse.md)
 - [LinkAnnotation](docs/LinkAnnotation.md)
 - [LinkAnnotationResponse](docs/LinkAnnotationResponse.md)
 - [LinkAnnotations](docs/LinkAnnotations.md)
 - [LinkAnnotationsResponse](docs/LinkAnnotationsResponse.md)
 - [Page](docs/Page.md)
 - [Pages](docs/Pages.md)
 - [PolyLineAnnotationResponse](docs/PolyLineAnnotationResponse.md)
 - [PolyLineAnnotations](docs/PolyLineAnnotations.md)
 - [PolyLineAnnotationsResponse](docs/PolyLineAnnotationsResponse.md)
 - [PolygonAnnotationResponse](docs/PolygonAnnotationResponse.md)
 - [PolygonAnnotations](docs/PolygonAnnotations.md)
 - [PolygonAnnotationsResponse](docs/PolygonAnnotationsResponse.md)
 - [PopupAnnotationResponse](docs/PopupAnnotationResponse.md)
 - [PopupAnnotations](docs/PopupAnnotations.md)
 - [PopupAnnotationsResponse](docs/PopupAnnotationsResponse.md)
 - [SignatureVerifyResponse](docs/SignatureVerifyResponse.md)
 - [SplitResultDocument](docs/SplitResultDocument.md)
 - [SplitResultResponse](docs/SplitResultResponse.md)
 - [SquareAnnotationResponse](docs/SquareAnnotationResponse.md)
 - [SquareAnnotations](docs/SquareAnnotations.md)
 - [SquareAnnotationsResponse](docs/SquareAnnotationsResponse.md)
 - [SquigglyAnnotationResponse](docs/SquigglyAnnotationResponse.md)
 - [SquigglyAnnotations](docs/SquigglyAnnotations.md)
 - [SquigglyAnnotationsResponse](docs/SquigglyAnnotationsResponse.md)
 - [StorageExistResponse](docs/StorageExistResponse.md)
 - [StrikeOutAnnotationResponse](docs/StrikeOutAnnotationResponse.md)
 - [StrikeOutAnnotations](docs/StrikeOutAnnotations.md)
 - [StrikeOutAnnotationsResponse](docs/StrikeOutAnnotationsResponse.md)
 - [TextAnnotationResponse](docs/TextAnnotationResponse.md)
 - [TextAnnotations](docs/TextAnnotations.md)
 - [TextAnnotationsResponse](docs/TextAnnotationsResponse.md)
 - [TextRectsResponse](docs/TextRectsResponse.md)
 - [TextReplaceResponse](docs/TextReplaceResponse.md)
 - [UnderlineAnnotationResponse](docs/UnderlineAnnotationResponse.md)
 - [UnderlineAnnotations](docs/UnderlineAnnotations.md)
 - [UnderlineAnnotationsResponse](docs/UnderlineAnnotationsResponse.md)
 - [WordCountResponse](docs/WordCountResponse.md)
 - [AnnotationInfo](docs/AnnotationInfo.md)
 - [MarkupAnnotation](docs/MarkupAnnotation.md)
 - [PopupAnnotation](docs/PopupAnnotation.md)
 - [CaretAnnotation](docs/CaretAnnotation.md)
 - [CommonFigureAnnotation](docs/CommonFigureAnnotation.md)
 - [FreeTextAnnotation](docs/FreeTextAnnotation.md)
 - [HighlightAnnotation](docs/HighlightAnnotation.md)
 - [InkAnnotation](docs/InkAnnotation.md)
 - [LineAnnotation](docs/LineAnnotation.md)
 - [PolyAnnotation](docs/PolyAnnotation.md)
 - [PopupAnnotationWithParent](docs/PopupAnnotationWithParent.md)
 - [SquigglyAnnotation](docs/SquigglyAnnotation.md)
 - [StrikeOutAnnotation](docs/StrikeOutAnnotation.md)
 - [TextAnnotation](docs/TextAnnotation.md)
 - [UnderlineAnnotation](docs/UnderlineAnnotation.md)
 - [CircleAnnotation](docs/CircleAnnotation.md)
 - [PolyLineAnnotation](docs/PolyLineAnnotation.md)
 - [PolygonAnnotation](docs/PolygonAnnotation.md)
 - [SquareAnnotation](docs/SquareAnnotation.md)

