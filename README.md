# Aspose.PDF Cloud
- API version: 3.0
- Package version: 20.6.0

[Aspose.PDF Cloud](https://products.aspose.cloud/pdf) is a true REST API that enables you to perform a wide range of document processing operations including creation, manipulation, conversion and rendering of PDF documents in the cloud.

Our Cloud SDKs are wrappers around REST API in various programming languages, allowing you to process documents in language of your choice quickly and easily, gaining all benefits of strong types and IDE highlights. This repository contains new generation SDKs for Aspose.PDF Cloud and examples.

These SDKs are now fully supported. If you have any questions, see any bugs or have enhancement request, feel free to reach out to us at [Free Support Forums](https://forum.aspose.cloud/c/pdf).

Extract Text & Images of a PDF document online https://products.aspose.app/pdf/parser.

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

All URIs are relative to *https://api.aspose.cloud/v3.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PdfApi* | [**CopyFile**](docs/PdfApi.md#copyfile) | **Put** /pdf/storage/file/copy/{srcPath} | Copy file
*PdfApi* | [**CopyFolder**](docs/PdfApi.md#copyfolder) | **Put** /pdf/storage/folder/copy/{srcPath} | Copy folder
*PdfApi* | [**CreateFolder**](docs/PdfApi.md#createfolder) | **Put** /pdf/storage/folder/{path} | Create the folder
*PdfApi* | [**DeleteAnnotation**](docs/PdfApi.md#deleteannotation) | **Delete** /pdf/{name}/annotations/{annotationId} | Delete document annotation by ID
*PdfApi* | [**DeleteBookmark**](docs/PdfApi.md#deletebookmark) | **Delete** /pdf/{name}/bookmarks/bookmark/{bookmarkPath} | Delete document bookmark by ID.
*PdfApi* | [**DeleteDocumentAnnotations**](docs/PdfApi.md#deletedocumentannotations) | **Delete** /pdf/{name}/annotations | Delete all annotations from the document
*PdfApi* | [**DeleteDocumentBookmarks**](docs/PdfApi.md#deletedocumentbookmarks) | **Delete** /pdf/{name}/bookmarks/tree | Delete all document bookmarks.
*PdfApi* | [**DeleteDocumentLinkAnnotations**](docs/PdfApi.md#deletedocumentlinkannotations) | **Delete** /pdf/{name}/links | Delete all link annotations from the document
*PdfApi* | [**DeleteDocumentStamps**](docs/PdfApi.md#deletedocumentstamps) | **Delete** /pdf/{name}/stamps | Delete all stamps from the document
*PdfApi* | [**DeleteDocumentTables**](docs/PdfApi.md#deletedocumenttables) | **Delete** /pdf/{name}/tables | Delete all tables from the document
*PdfApi* | [**DeleteField**](docs/PdfApi.md#deletefield) | **Delete** /pdf/{name}/fields/{fieldName} | Delete document field by name.
*PdfApi* | [**DeleteFile**](docs/PdfApi.md#deletefile) | **Delete** /pdf/storage/file/{path} | Delete file
*PdfApi* | [**DeleteFolder**](docs/PdfApi.md#deletefolder) | **Delete** /pdf/storage/folder/{path} | Delete folder
*PdfApi* | [**DeleteImage**](docs/PdfApi.md#deleteimage) | **Delete** /pdf/{name}/images/{imageId} | Delete image from document page.
*PdfApi* | [**DeleteLinkAnnotation**](docs/PdfApi.md#deletelinkannotation) | **Delete** /pdf/{name}/links/{linkId} | Delete document page link annotation by ID
*PdfApi* | [**DeletePage**](docs/PdfApi.md#deletepage) | **Delete** /pdf/{name}/pages/{pageNumber} | Delete document page by its number.
*PdfApi* | [**DeletePageAnnotations**](docs/PdfApi.md#deletepageannotations) | **Delete** /pdf/{name}/pages/{pageNumber}/annotations | Delete all annotations from the page
*PdfApi* | [**DeletePageLinkAnnotations**](docs/PdfApi.md#deletepagelinkannotations) | **Delete** /pdf/{name}/pages/{pageNumber}/links | Delete all link annotations from the page
*PdfApi* | [**DeletePageStamps**](docs/PdfApi.md#deletepagestamps) | **Delete** /pdf/{name}/pages/{pageNumber}/stamps | Delete all stamps from the page
*PdfApi* | [**DeletePageTables**](docs/PdfApi.md#deletepagetables) | **Delete** /pdf/{name}/pages/{pageNumber}/tables | Delete all tables from the page
*PdfApi* | [**DeleteProperties**](docs/PdfApi.md#deleteproperties) | **Delete** /pdf/{name}/documentproperties | Delete custom document properties.
*PdfApi* | [**DeleteProperty**](docs/PdfApi.md#deleteproperty) | **Delete** /pdf/{name}/documentproperties/{propertyName} | Delete document property.
*PdfApi* | [**DeleteStamp**](docs/PdfApi.md#deletestamp) | **Delete** /pdf/{name}/stamps/{stampId} | Delete document stamp by ID
*PdfApi* | [**DeleteTable**](docs/PdfApi.md#deletetable) | **Delete** /pdf/{name}/tables/{tableId} | Delete document table by ID
*PdfApi* | [**DownloadFile**](docs/PdfApi.md#downloadfile) | **Get** /pdf/storage/file/{path} | Download file
*PdfApi* | [**GetBookmark**](docs/PdfApi.md#getbookmark) | **Get** /pdf/{name}/bookmarks/bookmark/{bookmarkPath} | Read document bookmark.
*PdfApi* | [**GetBookmarks**](docs/PdfApi.md#getbookmarks) | **Get** /pdf/{name}/bookmarks/list/{bookmarkPath} | Read document bookmarks node list.
*PdfApi* | [**GetCaretAnnotation**](docs/PdfApi.md#getcaretannotation) | **Get** /pdf/{name}/annotations/caret/{annotationId} | Read document page caret annotation by ID.
*PdfApi* | [**GetCheckBoxField**](docs/PdfApi.md#getcheckboxfield) | **Get** /pdf/{name}/fields/checkbox/{fieldName} | Read document checkbox field by name.
*PdfApi* | [**GetCircleAnnotation**](docs/PdfApi.md#getcircleannotation) | **Get** /pdf/{name}/annotations/circle/{annotationId} | Read document page circle annotation by ID.
*PdfApi* | [**GetComboBoxField**](docs/PdfApi.md#getcomboboxfield) | **Get** /pdf/{name}/fields/combobox/{fieldName} | Read document combobox field by name.
*PdfApi* | [**GetDiscUsage**](docs/PdfApi.md#getdiscusage) | **Get** /pdf/storage/disc | Get disc usage
*PdfApi* | [**GetDocument**](docs/PdfApi.md#getdocument) | **Get** /pdf/{name} | Read common document info.
*PdfApi* | [**GetDocumentAnnotations**](docs/PdfApi.md#getdocumentannotations) | **Get** /pdf/{name}/annotations | Read documant page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
*PdfApi* | [**GetDocumentAttachmentByIndex**](docs/PdfApi.md#getdocumentattachmentbyindex) | **Get** /pdf/{name}/attachments/{attachmentIndex} | Read document attachment info by its index.
*PdfApi* | [**GetDocumentAttachments**](docs/PdfApi.md#getdocumentattachments) | **Get** /pdf/{name}/attachments | Read document attachments info.
*PdfApi* | [**GetDocumentBookmarks**](docs/PdfApi.md#getdocumentbookmarks) | **Get** /pdf/{name}/bookmarks/tree | Read document bookmarks tree.
*PdfApi* | [**GetDocumentCaretAnnotations**](docs/PdfApi.md#getdocumentcaretannotations) | **Get** /pdf/{name}/annotations/caret | Read document caret annotations.
*PdfApi* | [**GetDocumentCheckBoxFields**](docs/PdfApi.md#getdocumentcheckboxfields) | **Get** /pdf/{name}/fields/checkbox | Read document checkbox fields.
*PdfApi* | [**GetDocumentCircleAnnotations**](docs/PdfApi.md#getdocumentcircleannotations) | **Get** /pdf/{name}/annotations/circle | Read document circle annotations.
*PdfApi* | [**GetDocumentComboBoxFields**](docs/PdfApi.md#getdocumentcomboboxfields) | **Get** /pdf/{name}/fields/combobox | Read document combobox fields.
*PdfApi* | [**GetDocumentDisplayProperties**](docs/PdfApi.md#getdocumentdisplayproperties) | **Get** /pdf/{name}/displayproperties | Read document display properties.
*PdfApi* | [**GetDocumentFileAttachmentAnnotations**](docs/PdfApi.md#getdocumentfileattachmentannotations) | **Get** /pdf/{name}/annotations/fileattachment | Read document FileAttachment annotations.
*PdfApi* | [**GetDocumentFreeTextAnnotations**](docs/PdfApi.md#getdocumentfreetextannotations) | **Get** /pdf/{name}/annotations/freetext | Read document free text annotations.
*PdfApi* | [**GetDocumentHighlightAnnotations**](docs/PdfApi.md#getdocumenthighlightannotations) | **Get** /pdf/{name}/annotations/highlight | Read document highlight annotations.
*PdfApi* | [**GetDocumentInkAnnotations**](docs/PdfApi.md#getdocumentinkannotations) | **Get** /pdf/{name}/annotations/ink | Read document ink annotations.
*PdfApi* | [**GetDocumentLineAnnotations**](docs/PdfApi.md#getdocumentlineannotations) | **Get** /pdf/{name}/annotations/line | Read document line annotations.
*PdfApi* | [**GetDocumentListBoxFields**](docs/PdfApi.md#getdocumentlistboxfields) | **Get** /pdf/{name}/fields/listbox | Read document listbox fields.
*PdfApi* | [**GetDocumentMovieAnnotations**](docs/PdfApi.md#getdocumentmovieannotations) | **Get** /pdf/{name}/annotations/movie | Read document movie annotations.
*PdfApi* | [**GetDocumentPolyLineAnnotations**](docs/PdfApi.md#getdocumentpolylineannotations) | **Get** /pdf/{name}/annotations/polyline | Read document polyline annotations.
*PdfApi* | [**GetDocumentPolygonAnnotations**](docs/PdfApi.md#getdocumentpolygonannotations) | **Get** /pdf/{name}/annotations/polygon | Read document polygon annotations.
*PdfApi* | [**GetDocumentPopupAnnotations**](docs/PdfApi.md#getdocumentpopupannotations) | **Get** /pdf/{name}/annotations/popup | Read document popup annotations.
*PdfApi* | [**GetDocumentPopupAnnotationsByParent**](docs/PdfApi.md#getdocumentpopupannotationsbyparent) | **Get** /pdf/{name}/annotations/{annotationId}/popup | Read document popup annotations by parent id.
*PdfApi* | [**GetDocumentProperties**](docs/PdfApi.md#getdocumentproperties) | **Get** /pdf/{name}/documentproperties | Read document properties.
*PdfApi* | [**GetDocumentProperty**](docs/PdfApi.md#getdocumentproperty) | **Get** /pdf/{name}/documentproperties/{propertyName} | Read document property by name.
*PdfApi* | [**GetDocumentRadioButtonFields**](docs/PdfApi.md#getdocumentradiobuttonfields) | **Get** /pdf/{name}/fields/radiobutton | Read document radiobutton fields.
*PdfApi* | [**GetDocumentRedactionAnnotations**](docs/PdfApi.md#getdocumentredactionannotations) | **Get** /pdf/{name}/annotations/redaction | Read document redaction annotations.
*PdfApi* | [**GetDocumentScreenAnnotations**](docs/PdfApi.md#getdocumentscreenannotations) | **Get** /pdf/{name}/annotations/screen | Read document screen annotations.
*PdfApi* | [**GetDocumentSignatureFields**](docs/PdfApi.md#getdocumentsignaturefields) | **Get** /pdf/{name}/fields/signature | Read document signature fields.
*PdfApi* | [**GetDocumentSoundAnnotations**](docs/PdfApi.md#getdocumentsoundannotations) | **Get** /pdf/{name}/annotations/sound | Read document sound annotations.
*PdfApi* | [**GetDocumentSquareAnnotations**](docs/PdfApi.md#getdocumentsquareannotations) | **Get** /pdf/{name}/annotations/square | Read document square annotations.
*PdfApi* | [**GetDocumentSquigglyAnnotations**](docs/PdfApi.md#getdocumentsquigglyannotations) | **Get** /pdf/{name}/annotations/squiggly | Read document squiggly annotations.
*PdfApi* | [**GetDocumentStampAnnotations**](docs/PdfApi.md#getdocumentstampannotations) | **Get** /pdf/{name}/annotations/stamp | Read document stamp annotations.
*PdfApi* | [**GetDocumentStamps**](docs/PdfApi.md#getdocumentstamps) | **Get** /pdf/{name}/stamps | Read document stamps.
*PdfApi* | [**GetDocumentStrikeOutAnnotations**](docs/PdfApi.md#getdocumentstrikeoutannotations) | **Get** /pdf/{name}/annotations/strikeout | Read document StrikeOut annotations.
*PdfApi* | [**GetDocumentTables**](docs/PdfApi.md#getdocumenttables) | **Get** /pdf/{name}/tables | Read document tables.
*PdfApi* | [**GetDocumentTextAnnotations**](docs/PdfApi.md#getdocumenttextannotations) | **Get** /pdf/{name}/annotations/text | Read document text annotations.
*PdfApi* | [**GetDocumentTextBoxFields**](docs/PdfApi.md#getdocumenttextboxfields) | **Get** /pdf/{name}/fields/textbox | Read document text box fields.
*PdfApi* | [**GetDocumentUnderlineAnnotations**](docs/PdfApi.md#getdocumentunderlineannotations) | **Get** /pdf/{name}/annotations/underline | Read document underline annotations.
*PdfApi* | [**GetDownloadDocumentAttachmentByIndex**](docs/PdfApi.md#getdownloaddocumentattachmentbyindex) | **Get** /pdf/{name}/attachments/{attachmentIndex}/download | Download document attachment content by its index.
*PdfApi* | [**GetEpubInStorageToPdf**](docs/PdfApi.md#getepubinstoragetopdf) | **Get** /pdf/create/epub | Convert EPUB file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetExportFieldsFromPdfToFdfInStorage**](docs/PdfApi.md#getexportfieldsfrompdftofdfinstorage) | **Get** /pdf/{name}/export/fdf | Export fields from from PDF in storage to FDF file.
*PdfApi* | [**GetExportFieldsFromPdfToXfdfInStorage**](docs/PdfApi.md#getexportfieldsfrompdftoxfdfinstorage) | **Get** /pdf/{name}/export/xfdf | Export fields from from PDF in storage to XFDF file.
*PdfApi* | [**GetExportFieldsFromPdfToXmlInStorage**](docs/PdfApi.md#getexportfieldsfrompdftoxmlinstorage) | **Get** /pdf/{name}/export/xml | Export fields from from PDF in storage to XML file.
*PdfApi* | [**GetField**](docs/PdfApi.md#getfield) | **Get** /pdf/{name}/fields/{fieldName} | Get document field by name.
*PdfApi* | [**GetFields**](docs/PdfApi.md#getfields) | **Get** /pdf/{name}/fields | Get document fields.
*PdfApi* | [**GetFileAttachmentAnnotation**](docs/PdfApi.md#getfileattachmentannotation) | **Get** /pdf/{name}/annotations/fileattachment/{annotationId} | Read document page FileAttachment annotation by ID.
*PdfApi* | [**GetFileAttachmentAnnotationData**](docs/PdfApi.md#getfileattachmentannotationdata) | **Get** /pdf/{name}/annotations/fileattachment/{annotationId}/data | Read document page FileAttachment annotation by ID.
*PdfApi* | [**GetFileVersions**](docs/PdfApi.md#getfileversions) | **Get** /pdf/storage/version/{path} | Get file versions
*PdfApi* | [**GetFilesList**](docs/PdfApi.md#getfileslist) | **Get** /pdf/storage/folder/{path} | Get all files and folders within a folder
*PdfApi* | [**GetFreeTextAnnotation**](docs/PdfApi.md#getfreetextannotation) | **Get** /pdf/{name}/annotations/freetext/{annotationId} | Read document page free text annotation by ID.
*PdfApi* | [**GetHighlightAnnotation**](docs/PdfApi.md#gethighlightannotation) | **Get** /pdf/{name}/annotations/highlight/{annotationId} | Read document page highlight annotation by ID.
*PdfApi* | [**GetHtmlInStorageToPdf**](docs/PdfApi.md#gethtmlinstoragetopdf) | **Get** /pdf/create/html | Convert HTML file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetImage**](docs/PdfApi.md#getimage) | **Get** /pdf/{name}/images/{imageId} | Read document image by ID.
*PdfApi* | [**GetImageExtractAsGif**](docs/PdfApi.md#getimageextractasgif) | **Get** /pdf/{name}/images/{imageId}/extract/gif | Extract document image in GIF format
*PdfApi* | [**GetImageExtractAsJpeg**](docs/PdfApi.md#getimageextractasjpeg) | **Get** /pdf/{name}/images/{imageId}/extract/jpeg | Extract document image in JPEG format
*PdfApi* | [**GetImageExtractAsPng**](docs/PdfApi.md#getimageextractaspng) | **Get** /pdf/{name}/images/{imageId}/extract/png | Extract document image in PNG format
*PdfApi* | [**GetImageExtractAsTiff**](docs/PdfApi.md#getimageextractastiff) | **Get** /pdf/{name}/images/{imageId}/extract/tiff | Extract document image in TIFF format
*PdfApi* | [**GetImages**](docs/PdfApi.md#getimages) | **Get** /pdf/{name}/pages/{pageNumber}/images | Read document images.
*PdfApi* | [**GetImportFieldsFromFdfInStorage**](docs/PdfApi.md#getimportfieldsfromfdfinstorage) | **Get** /pdf/{name}/import/fdf | Update fields from FDF file in storage.
*PdfApi* | [**GetImportFieldsFromXfdfInStorage**](docs/PdfApi.md#getimportfieldsfromxfdfinstorage) | **Get** /pdf/{name}/import/xfdf | Update fields from XFDF file in storage.
*PdfApi* | [**GetImportFieldsFromXmlInStorage**](docs/PdfApi.md#getimportfieldsfromxmlinstorage) | **Get** /pdf/{name}/import/xml | Import from XML file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetInkAnnotation**](docs/PdfApi.md#getinkannotation) | **Get** /pdf/{name}/annotations/ink/{annotationId} | Read document page ink annotation by ID.
*PdfApi* | [**GetLaTeXInStorageToPdf**](docs/PdfApi.md#getlatexinstoragetopdf) | **Get** /pdf/create/latex | Convert TeX file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetLineAnnotation**](docs/PdfApi.md#getlineannotation) | **Get** /pdf/{name}/annotations/line/{annotationId} | Read document page line annotation by ID.
*PdfApi* | [**GetLinkAnnotation**](docs/PdfApi.md#getlinkannotation) | **Get** /pdf/{name}/links/{linkId} | Read document link annotation by ID.
*PdfApi* | [**GetListBoxField**](docs/PdfApi.md#getlistboxfield) | **Get** /pdf/{name}/fields/listbox/{fieldName} | Read document listbox field by name.
*PdfApi* | [**GetMarkdownInStorageToPdf**](docs/PdfApi.md#getmarkdowninstoragetopdf) | **Get** /pdf/create/markdown | Convert MD file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetMhtInStorageToPdf**](docs/PdfApi.md#getmhtinstoragetopdf) | **Get** /pdf/create/mht | Convert MHT file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetMovieAnnotation**](docs/PdfApi.md#getmovieannotation) | **Get** /pdf/{name}/annotations/movie/{annotationId} | Read document page movie annotation by ID.
*PdfApi* | [**GetPage**](docs/PdfApi.md#getpage) | **Get** /pdf/{name}/pages/{pageNumber} | Read document page info.
*PdfApi* | [**GetPageAnnotations**](docs/PdfApi.md#getpageannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations | Read document page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
*PdfApi* | [**GetPageCaretAnnotations**](docs/PdfApi.md#getpagecaretannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/caret | Read document page caret annotations.
*PdfApi* | [**GetPageCheckBoxFields**](docs/PdfApi.md#getpagecheckboxfields) | **Get** /pdf/{name}/page/{pageNumber}/fields/checkbox | Read document page checkbox fields.
*PdfApi* | [**GetPageCircleAnnotations**](docs/PdfApi.md#getpagecircleannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/circle | Read document page circle annotations.
*PdfApi* | [**GetPageComboBoxFields**](docs/PdfApi.md#getpagecomboboxfields) | **Get** /pdf/{name}/page/{pageNumber}/fields/combobox | Read document page combobox fields.
*PdfApi* | [**GetPageConvertToBmp**](docs/PdfApi.md#getpageconverttobmp) | **Get** /pdf/{name}/pages/{pageNumber}/convert/bmp | Convert document page to Bmp image and return resulting file in response.
*PdfApi* | [**GetPageConvertToEmf**](docs/PdfApi.md#getpageconverttoemf) | **Get** /pdf/{name}/pages/{pageNumber}/convert/emf | Convert document page to Emf image and return resulting file in response.
*PdfApi* | [**GetPageConvertToGif**](docs/PdfApi.md#getpageconverttogif) | **Get** /pdf/{name}/pages/{pageNumber}/convert/gif | Convert document page to Gif image and return resulting file in response.
*PdfApi* | [**GetPageConvertToJpeg**](docs/PdfApi.md#getpageconverttojpeg) | **Get** /pdf/{name}/pages/{pageNumber}/convert/jpeg | Convert document page to Jpeg image and return resulting file in response.
*PdfApi* | [**GetPageConvertToPng**](docs/PdfApi.md#getpageconverttopng) | **Get** /pdf/{name}/pages/{pageNumber}/convert/png | Convert document page to Png image and return resulting file in response.
*PdfApi* | [**GetPageConvertToTiff**](docs/PdfApi.md#getpageconverttotiff) | **Get** /pdf/{name}/pages/{pageNumber}/convert/tiff | Convert document page to Tiff image  and return resulting file in response.
*PdfApi* | [**GetPageFileAttachmentAnnotations**](docs/PdfApi.md#getpagefileattachmentannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/fileattachment | Read document page FileAttachment annotations.
*PdfApi* | [**GetPageFreeTextAnnotations**](docs/PdfApi.md#getpagefreetextannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/freetext | Read document page free text annotations.
*PdfApi* | [**GetPageHighlightAnnotations**](docs/PdfApi.md#getpagehighlightannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/highlight | Read document page highlight annotations.
*PdfApi* | [**GetPageInkAnnotations**](docs/PdfApi.md#getpageinkannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/ink | Read document page ink annotations.
*PdfApi* | [**GetPageLineAnnotations**](docs/PdfApi.md#getpagelineannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/line | Read document page line annotations.
*PdfApi* | [**GetPageLinkAnnotation**](docs/PdfApi.md#getpagelinkannotation) | **Get** /pdf/{name}/pages/{pageNumber}/links/{linkId} | Read document page link annotation by ID.
*PdfApi* | [**GetPageLinkAnnotations**](docs/PdfApi.md#getpagelinkannotations) | **Get** /pdf/{name}/pages/{pageNumber}/links | Read document page link annotations.
*PdfApi* | [**GetPageListBoxFields**](docs/PdfApi.md#getpagelistboxfields) | **Get** /pdf/{name}/page/{pageNumber}/fields/listbox | Read document page listbox fields.
*PdfApi* | [**GetPageMovieAnnotations**](docs/PdfApi.md#getpagemovieannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/movie | Read document page movie annotations.
*PdfApi* | [**GetPagePolyLineAnnotations**](docs/PdfApi.md#getpagepolylineannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/polyline | Read document page polyline annotations.
*PdfApi* | [**GetPagePolygonAnnotations**](docs/PdfApi.md#getpagepolygonannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/polygon | Read document page polygon annotations.
*PdfApi* | [**GetPagePopupAnnotations**](docs/PdfApi.md#getpagepopupannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/popup | Read document page popup annotations.
*PdfApi* | [**GetPageRadioButtonFields**](docs/PdfApi.md#getpageradiobuttonfields) | **Get** /pdf/{name}/page/{pageNumber}/fields/radiobutton | Read document page radiobutton fields.
*PdfApi* | [**GetPageRedactionAnnotations**](docs/PdfApi.md#getpageredactionannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/redaction | Read document page redaction annotations.
*PdfApi* | [**GetPageScreenAnnotations**](docs/PdfApi.md#getpagescreenannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/screen | Read document page screen annotations.
*PdfApi* | [**GetPageSignatureFields**](docs/PdfApi.md#getpagesignaturefields) | **Get** /pdf/{name}/page/{pageNumber}/fields/signature | Read document page signature fields.
*PdfApi* | [**GetPageSoundAnnotations**](docs/PdfApi.md#getpagesoundannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/sound | Read document page sound annotations.
*PdfApi* | [**GetPageSquareAnnotations**](docs/PdfApi.md#getpagesquareannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/square | Read document page square annotations.
*PdfApi* | [**GetPageSquigglyAnnotations**](docs/PdfApi.md#getpagesquigglyannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/squiggly | Read document page squiggly annotations.
*PdfApi* | [**GetPageStampAnnotations**](docs/PdfApi.md#getpagestampannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/stamp | Read document page stamp annotations.
*PdfApi* | [**GetPageStamps**](docs/PdfApi.md#getpagestamps) | **Get** /pdf/{name}/pages/{pageNumber}/stamps | Read page document stamps.
*PdfApi* | [**GetPageStrikeOutAnnotations**](docs/PdfApi.md#getpagestrikeoutannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/strikeout | Read document page StrikeOut annotations.
*PdfApi* | [**GetPageTables**](docs/PdfApi.md#getpagetables) | **Get** /pdf/{name}/pages/{pageNumber}/tables | Read document page tables.
*PdfApi* | [**GetPageText**](docs/PdfApi.md#getpagetext) | **Get** /pdf/{name}/pages/{pageNumber}/text | Read page text items.
*PdfApi* | [**GetPageTextAnnotations**](docs/PdfApi.md#getpagetextannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/text | Read document page text annotations.
*PdfApi* | [**GetPageTextBoxFields**](docs/PdfApi.md#getpagetextboxfields) | **Get** /pdf/{name}/page/{pageNumber}/fields/textbox | Read document page text box fields.
*PdfApi* | [**GetPageUnderlineAnnotations**](docs/PdfApi.md#getpageunderlineannotations) | **Get** /pdf/{name}/pages/{pageNumber}/annotations/underline | Read document page underline annotations.
*PdfApi* | [**GetPages**](docs/PdfApi.md#getpages) | **Get** /pdf/{name}/pages | Read document pages info.
*PdfApi* | [**GetPclInStorageToPdf**](docs/PdfApi.md#getpclinstoragetopdf) | **Get** /pdf/create/pcl | Convert PCL file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetPdfAInStorageToPdf**](docs/PdfApi.md#getpdfainstoragetopdf) | **Get** /pdf/create/pdfa | Convert PDFA file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetPdfInStorageToDoc**](docs/PdfApi.md#getpdfinstoragetodoc) | **Get** /pdf/{name}/convert/doc | Converts PDF document (located on storage) to DOC format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToEpub**](docs/PdfApi.md#getpdfinstoragetoepub) | **Get** /pdf/{name}/convert/epub | Converts PDF document (located on storage) to EPUB format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToHtml**](docs/PdfApi.md#getpdfinstoragetohtml) | **Get** /pdf/{name}/convert/html | Converts PDF document (located on storage) to Html format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToLaTeX**](docs/PdfApi.md#getpdfinstoragetolatex) | **Get** /pdf/{name}/convert/latex | Converts PDF document (located on storage) to TeX format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToMobiXml**](docs/PdfApi.md#getpdfinstoragetomobixml) | **Get** /pdf/{name}/convert/mobixml | Converts PDF document (located on storage) to MOBIXML format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToPdfA**](docs/PdfApi.md#getpdfinstoragetopdfa) | **Get** /pdf/{name}/convert/pdfa | Converts PDF document (located on storage) to PdfA format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToPptx**](docs/PdfApi.md#getpdfinstoragetopptx) | **Get** /pdf/{name}/convert/pptx | Converts PDF document (located on storage) to PPTX format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToSvg**](docs/PdfApi.md#getpdfinstoragetosvg) | **Get** /pdf/{name}/convert/svg | Converts PDF document (located on storage) to SVG format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToTeX**](docs/PdfApi.md#getpdfinstoragetotex) | **Get** /pdf/{name}/convert/tex | Converts PDF document (located on storage) to TeX format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToTiff**](docs/PdfApi.md#getpdfinstoragetotiff) | **Get** /pdf/{name}/convert/tiff | Converts PDF document (located on storage) to TIFF format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToXls**](docs/PdfApi.md#getpdfinstoragetoxls) | **Get** /pdf/{name}/convert/xls | Converts PDF document (located on storage) to XLS format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToXlsx**](docs/PdfApi.md#getpdfinstoragetoxlsx) | **Get** /pdf/{name}/convert/xlsx | Converts PDF document (located on storage) to XLSX format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToXml**](docs/PdfApi.md#getpdfinstoragetoxml) | **Get** /pdf/{name}/convert/xml | Converts PDF document (located on storage) to XML format and returns resulting file in response content
*PdfApi* | [**GetPdfInStorageToXps**](docs/PdfApi.md#getpdfinstoragetoxps) | **Get** /pdf/{name}/convert/xps | Converts PDF document (located on storage) to XPS format and returns resulting file in response content
*PdfApi* | [**GetPolyLineAnnotation**](docs/PdfApi.md#getpolylineannotation) | **Get** /pdf/{name}/annotations/polyline/{annotationId} | Read document page polyline annotation by ID.
*PdfApi* | [**GetPolygonAnnotation**](docs/PdfApi.md#getpolygonannotation) | **Get** /pdf/{name}/annotations/polygon/{annotationId} | Read document page polygon annotation by ID.
*PdfApi* | [**GetPopupAnnotation**](docs/PdfApi.md#getpopupannotation) | **Get** /pdf/{name}/annotations/popup/{annotationId} | Read document page popup annotation by ID.
*PdfApi* | [**GetPsInStorageToPdf**](docs/PdfApi.md#getpsinstoragetopdf) | **Get** /pdf/create/ps | Convert PS file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetRadioButtonField**](docs/PdfApi.md#getradiobuttonfield) | **Get** /pdf/{name}/fields/radiobutton/{fieldName} | Read document RadioButton field by name.
*PdfApi* | [**GetRedactionAnnotation**](docs/PdfApi.md#getredactionannotation) | **Get** /pdf/{name}/annotations/redaction/{annotationId} | Read document page redaction annotation by ID.
*PdfApi* | [**GetScreenAnnotation**](docs/PdfApi.md#getscreenannotation) | **Get** /pdf/{name}/annotations/screen/{annotationId} | Read document page screen annotation by ID.
*PdfApi* | [**GetScreenAnnotationData**](docs/PdfApi.md#getscreenannotationdata) | **Get** /pdf/{name}/annotations/screen/{annotationId}/data | Read document page screen annotation by ID.
*PdfApi* | [**GetSignatureField**](docs/PdfApi.md#getsignaturefield) | **Get** /pdf/{name}/fields/signature/{fieldName} | Read document signature field by name.
*PdfApi* | [**GetSoundAnnotation**](docs/PdfApi.md#getsoundannotation) | **Get** /pdf/{name}/annotations/sound/{annotationId} | Read document page sound annotation by ID.
*PdfApi* | [**GetSoundAnnotationData**](docs/PdfApi.md#getsoundannotationdata) | **Get** /pdf/{name}/annotations/sound/{annotationId}/data | Read document page sound annotation by ID.
*PdfApi* | [**GetSquareAnnotation**](docs/PdfApi.md#getsquareannotation) | **Get** /pdf/{name}/annotations/square/{annotationId} | Read document page square annotation by ID.
*PdfApi* | [**GetSquigglyAnnotation**](docs/PdfApi.md#getsquigglyannotation) | **Get** /pdf/{name}/annotations/squiggly/{annotationId} | Read document page squiggly annotation by ID.
*PdfApi* | [**GetStampAnnotation**](docs/PdfApi.md#getstampannotation) | **Get** /pdf/{name}/annotations/stamp/{annotationId} | Read document page stamp annotation by ID.
*PdfApi* | [**GetStampAnnotationData**](docs/PdfApi.md#getstampannotationdata) | **Get** /pdf/{name}/annotations/stamp/{annotationId}/data | Read document page stamp annotation by ID.
*PdfApi* | [**GetStrikeOutAnnotation**](docs/PdfApi.md#getstrikeoutannotation) | **Get** /pdf/{name}/annotations/strikeout/{annotationId} | Read document page StrikeOut annotation by ID.
*PdfApi* | [**GetSvgInStorageToPdf**](docs/PdfApi.md#getsvginstoragetopdf) | **Get** /pdf/create/svg | Convert SVG file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetTable**](docs/PdfApi.md#gettable) | **Get** /pdf/{name}/tables/{tableId} | Read document page table by ID.
*PdfApi* | [**GetTeXInStorageToPdf**](docs/PdfApi.md#gettexinstoragetopdf) | **Get** /pdf/create/tex | Convert TeX file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetText**](docs/PdfApi.md#gettext) | **Get** /pdf/{name}/text | Read document text.
*PdfApi* | [**GetTextAnnotation**](docs/PdfApi.md#gettextannotation) | **Get** /pdf/{name}/annotations/text/{annotationId} | Read document page text annotation by ID.
*PdfApi* | [**GetTextBoxField**](docs/PdfApi.md#gettextboxfield) | **Get** /pdf/{name}/fields/textbox/{fieldName} | Read document text box field by name.
*PdfApi* | [**GetUnderlineAnnotation**](docs/PdfApi.md#getunderlineannotation) | **Get** /pdf/{name}/annotations/underline/{annotationId} | Read document page underline annotation by ID.
*PdfApi* | [**GetVerifySignature**](docs/PdfApi.md#getverifysignature) | **Get** /pdf/{name}/verifySignature | Verify signature document.
*PdfApi* | [**GetWebInStorageToPdf**](docs/PdfApi.md#getwebinstoragetopdf) | **Get** /pdf/create/web | Convert web page to PDF format and return resulting file in response. 
*PdfApi* | [**GetWordsPerPage**](docs/PdfApi.md#getwordsperpage) | **Get** /pdf/{name}/pages/wordCount | Get number of words per document page.
*PdfApi* | [**GetXfaPdfInStorageToAcroForm**](docs/PdfApi.md#getxfapdfinstoragetoacroform) | **Get** /pdf/{name}/convert/xfatoacroform | Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and returns resulting file response content
*PdfApi* | [**GetXmlInStorageToPdf**](docs/PdfApi.md#getxmlinstoragetopdf) | **Get** /pdf/create/xml | Convert XML file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetXpsInStorageToPdf**](docs/PdfApi.md#getxpsinstoragetopdf) | **Get** /pdf/create/xps | Convert XPS file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**GetXslFoInStorageToPdf**](docs/PdfApi.md#getxslfoinstoragetopdf) | **Get** /pdf/create/xslfo | Convert XslFo file (located on storage) to PDF format and return resulting file in response. 
*PdfApi* | [**MoveFile**](docs/PdfApi.md#movefile) | **Put** /pdf/storage/file/move/{srcPath} | Move file
*PdfApi* | [**MoveFolder**](docs/PdfApi.md#movefolder) | **Put** /pdf/storage/folder/move/{srcPath} | Move folder
*PdfApi* | [**ObjectExists**](docs/PdfApi.md#objectexists) | **Get** /pdf/storage/exist/{path} | Check if file or folder exists
*PdfApi* | [**PostAppendDocument**](docs/PdfApi.md#postappenddocument) | **Post** /pdf/{name}/appendDocument | Append document to existing one.
*PdfApi* | [**PostBookmark**](docs/PdfApi.md#postbookmark) | **Post** /pdf/{name}/bookmarks/bookmark/{bookmarkPath} | Add document bookmarks.
*PdfApi* | [**PostChangePasswordDocumentInStorage**](docs/PdfApi.md#postchangepassworddocumentinstorage) | **Post** /pdf/{name}/changepassword | Change document password in storage.
*PdfApi* | [**PostCheckBoxFields**](docs/PdfApi.md#postcheckboxfields) | **Post** /pdf/{name}/fields/checkbox | Add document checkbox fields.
*PdfApi* | [**PostComboBoxFields**](docs/PdfApi.md#postcomboboxfields) | **Post** /pdf/{name}/fields/combobox | Add document combobox fields.
*PdfApi* | [**PostCreateDocument**](docs/PdfApi.md#postcreatedocument) | **Post** /pdf/{name} | Create empty document.
*PdfApi* | [**PostCreateField**](docs/PdfApi.md#postcreatefield) | **Post** /pdf/{name}/fields | Create field.
*PdfApi* | [**PostDecryptDocumentInStorage**](docs/PdfApi.md#postdecryptdocumentinstorage) | **Post** /pdf/{name}/decrypt | Decrypt document in storage.
*PdfApi* | [**PostDocumentImageFooter**](docs/PdfApi.md#postdocumentimagefooter) | **Post** /pdf/{name}/footer/image | Add document image footer.
*PdfApi* | [**PostDocumentImageHeader**](docs/PdfApi.md#postdocumentimageheader) | **Post** /pdf/{name}/header/image | Add document image header.
*PdfApi* | [**PostDocumentPageNumberStamps**](docs/PdfApi.md#postdocumentpagenumberstamps) | **Post** /pdf/{name}/stamps/pagenumber | Add document page number stamps.
*PdfApi* | [**PostDocumentTextFooter**](docs/PdfApi.md#postdocumenttextfooter) | **Post** /pdf/{name}/footer/text | Add document text footer.
*PdfApi* | [**PostDocumentTextHeader**](docs/PdfApi.md#postdocumenttextheader) | **Post** /pdf/{name}/header/text | Add document text header.
*PdfApi* | [**PostDocumentTextReplace**](docs/PdfApi.md#postdocumenttextreplace) | **Post** /pdf/{name}/text/replace | Document&#39;s replace text method.
*PdfApi* | [**PostEncryptDocumentInStorage**](docs/PdfApi.md#postencryptdocumentinstorage) | **Post** /pdf/{name}/encrypt | Encrypt document in storage.
*PdfApi* | [**PostFlattenDocument**](docs/PdfApi.md#postflattendocument) | **Post** /pdf/{name}/flatten | Flatten the document.
*PdfApi* | [**PostImportFieldsFromFdf**](docs/PdfApi.md#postimportfieldsfromfdf) | **Post** /pdf/{name}/import/fdf | Update fields from FDF file in request.
*PdfApi* | [**PostImportFieldsFromXfdf**](docs/PdfApi.md#postimportfieldsfromxfdf) | **Post** /pdf/{name}/import/xfdf | Update fields from XFDF file in request.
*PdfApi* | [**PostImportFieldsFromXml**](docs/PdfApi.md#postimportfieldsfromxml) | **Post** /pdf/{name}/import/xml | Update fields from XML file in request.
*PdfApi* | [**PostInsertImage**](docs/PdfApi.md#postinsertimage) | **Post** /pdf/{name}/pages/{pageNumber}/images | Insert image to document page.
*PdfApi* | [**PostListBoxFields**](docs/PdfApi.md#postlistboxfields) | **Post** /pdf/{name}/fields/listbox | Add document listbox fields.
*PdfApi* | [**PostMovePage**](docs/PdfApi.md#postmovepage) | **Post** /pdf/{name}/pages/{pageNumber}/movePage | Move page to new position.
*PdfApi* | [**PostOptimizeDocument**](docs/PdfApi.md#postoptimizedocument) | **Post** /pdf/{name}/optimize | Optimize document.
*PdfApi* | [**PostPageCaretAnnotations**](docs/PdfApi.md#postpagecaretannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/caret | Add document page caret annotations.
*PdfApi* | [**PostPageCertify**](docs/PdfApi.md#postpagecertify) | **Post** /pdf/{name}/pages/{pageNumber}/certify | Certify document page.
*PdfApi* | [**PostPageCircleAnnotations**](docs/PdfApi.md#postpagecircleannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/circle | Add document page circle annotations.
*PdfApi* | [**PostPageFileAttachmentAnnotations**](docs/PdfApi.md#postpagefileattachmentannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/fileattachment | Add document page FileAttachment annotations.
*PdfApi* | [**PostPageFreeTextAnnotations**](docs/PdfApi.md#postpagefreetextannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/freetext | Add document page free text annotations.
*PdfApi* | [**PostPageHighlightAnnotations**](docs/PdfApi.md#postpagehighlightannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/highlight | Add document page highlight annotations.
*PdfApi* | [**PostPageImageStamps**](docs/PdfApi.md#postpageimagestamps) | **Post** /pdf/{name}/pages/{pageNumber}/stamps/image | Add document page image stamps.
*PdfApi* | [**PostPageInkAnnotations**](docs/PdfApi.md#postpageinkannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/ink | Add document page ink annotations.
*PdfApi* | [**PostPageLineAnnotations**](docs/PdfApi.md#postpagelineannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/line | Add document page line annotations.
*PdfApi* | [**PostPageLinkAnnotations**](docs/PdfApi.md#postpagelinkannotations) | **Post** /pdf/{name}/pages/{pageNumber}/links | Add document page link annotations.
*PdfApi* | [**PostPageMovieAnnotations**](docs/PdfApi.md#postpagemovieannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/movie | Add document page movie annotations.
*PdfApi* | [**PostPagePdfPageStamps**](docs/PdfApi.md#postpagepdfpagestamps) | **Post** /pdf/{name}/pages/{pageNumber}/stamps/pdfpage | Add document pdf page stamps.
*PdfApi* | [**PostPagePolyLineAnnotations**](docs/PdfApi.md#postpagepolylineannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/polyline | Add document page polyline annotations.
*PdfApi* | [**PostPagePolygonAnnotations**](docs/PdfApi.md#postpagepolygonannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/polygon | Add document page polygon annotations.
*PdfApi* | [**PostPageRedactionAnnotations**](docs/PdfApi.md#postpageredactionannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/redaction | Add document page redaction annotations.
*PdfApi* | [**PostPageScreenAnnotations**](docs/PdfApi.md#postpagescreenannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/screen | Add document page screen annotations.
*PdfApi* | [**PostPageSoundAnnotations**](docs/PdfApi.md#postpagesoundannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/sound | Add document page sound annotations.
*PdfApi* | [**PostPageSquareAnnotations**](docs/PdfApi.md#postpagesquareannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/square | Add document page square annotations.
*PdfApi* | [**PostPageSquigglyAnnotations**](docs/PdfApi.md#postpagesquigglyannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/squiggly | Add document page squiggly annotations.
*PdfApi* | [**PostPageStampAnnotations**](docs/PdfApi.md#postpagestampannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/stamp | Add document page stamp annotations.
*PdfApi* | [**PostPageStrikeOutAnnotations**](docs/PdfApi.md#postpagestrikeoutannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/strikeout | Add document page StrikeOut annotations.
*PdfApi* | [**PostPageTables**](docs/PdfApi.md#postpagetables) | **Post** /pdf/{name}/pages/{pageNumber}/tables | Add document page tables.
*PdfApi* | [**PostPageTextAnnotations**](docs/PdfApi.md#postpagetextannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/text | Add document page text annotations.
*PdfApi* | [**PostPageTextReplace**](docs/PdfApi.md#postpagetextreplace) | **Post** /pdf/{name}/pages/{pageNumber}/text/replace | Page&#39;s replace text method.
*PdfApi* | [**PostPageTextStamps**](docs/PdfApi.md#postpagetextstamps) | **Post** /pdf/{name}/pages/{pageNumber}/stamps/text | Add document page text stamps.
*PdfApi* | [**PostPageUnderlineAnnotations**](docs/PdfApi.md#postpageunderlineannotations) | **Post** /pdf/{name}/pages/{pageNumber}/annotations/underline | Add document page underline annotations.
*PdfApi* | [**PostPopupAnnotation**](docs/PdfApi.md#postpopupannotation) | **Post** /pdf/{name}/annotations/{annotationId}/popup | Add document popup annotations.
*PdfApi* | [**PostRadioButtonFields**](docs/PdfApi.md#postradiobuttonfields) | **Post** /pdf/{name}/fields/radiobutton | Add document RadioButton fields.
*PdfApi* | [**PostSignDocument**](docs/PdfApi.md#postsigndocument) | **Post** /pdf/{name}/sign | Sign document.
*PdfApi* | [**PostSignPage**](docs/PdfApi.md#postsignpage) | **Post** /pdf/{name}/pages/{pageNumber}/sign | Sign page.
*PdfApi* | [**PostSignatureField**](docs/PdfApi.md#postsignaturefield) | **Post** /pdf/{name}/fields/signature | Add document signature field.
*PdfApi* | [**PostSplitDocument**](docs/PdfApi.md#postsplitdocument) | **Post** /pdf/{name}/split | Split document to parts.
*PdfApi* | [**PostTextBoxFields**](docs/PdfApi.md#posttextboxfields) | **Post** /pdf/{name}/fields/textbox | Add document text box fields.
*PdfApi* | [**PutAddNewPage**](docs/PdfApi.md#putaddnewpage) | **Put** /pdf/{name}/pages | Add new page to end of the document.
*PdfApi* | [**PutAddText**](docs/PdfApi.md#putaddtext) | **Put** /pdf/{name}/pages/{pageNumber}/text | Add text to PDF document page.
*PdfApi* | [**PutAnnotationsFlatten**](docs/PdfApi.md#putannotationsflatten) | **Put** /pdf/{name}/annotations/flatten | Flattens the annotations of the specified types
*PdfApi* | [**PutBookmark**](docs/PdfApi.md#putbookmark) | **Put** /pdf/{name}/bookmarks/bookmark/{bookmarkPath} | Update document bookmark.
*PdfApi* | [**PutCaretAnnotation**](docs/PdfApi.md#putcaretannotation) | **Put** /pdf/{name}/annotations/caret/{annotationId} | Replace document caret annotation
*PdfApi* | [**PutChangePasswordDocument**](docs/PdfApi.md#putchangepassworddocument) | **Put** /pdf/changepassword | Change document password from content.
*PdfApi* | [**PutCheckBoxField**](docs/PdfApi.md#putcheckboxfield) | **Put** /pdf/{name}/fields/checkbox/{fieldName} | Replace document checkbox field
*PdfApi* | [**PutCircleAnnotation**](docs/PdfApi.md#putcircleannotation) | **Put** /pdf/{name}/annotations/circle/{annotationId} | Replace document circle annotation
*PdfApi* | [**PutComboBoxField**](docs/PdfApi.md#putcomboboxfield) | **Put** /pdf/{name}/fields/combobox/{fieldName} | Replace document combobox field
*PdfApi* | [**PutCreateDocument**](docs/PdfApi.md#putcreatedocument) | **Put** /pdf/{name} | Create empty document.
*PdfApi* | [**PutDecryptDocument**](docs/PdfApi.md#putdecryptdocument) | **Put** /pdf/decrypt | Decrypt document from content.
*PdfApi* | [**PutDocumentDisplayProperties**](docs/PdfApi.md#putdocumentdisplayproperties) | **Put** /pdf/{name}/displayproperties | Update document display properties.
*PdfApi* | [**PutEncryptDocument**](docs/PdfApi.md#putencryptdocument) | **Put** /pdf/encrypt | Encrypt document from content.
*PdfApi* | [**PutEpubInStorageToPdf**](docs/PdfApi.md#putepubinstoragetopdf) | **Put** /pdf/{name}/create/epub | Convert EPUB file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutExportFieldsFromPdfToFdfInStorage**](docs/PdfApi.md#putexportfieldsfrompdftofdfinstorage) | **Put** /pdf/{name}/export/fdf | Export fields from from PDF in storage to FDF file in storage.
*PdfApi* | [**PutExportFieldsFromPdfToXfdfInStorage**](docs/PdfApi.md#putexportfieldsfrompdftoxfdfinstorage) | **Put** /pdf/{name}/export/xfdf | Export fields from from PDF in storage to XFDF file in storage.
*PdfApi* | [**PutExportFieldsFromPdfToXmlInStorage**](docs/PdfApi.md#putexportfieldsfrompdftoxmlinstorage) | **Put** /pdf/{name}/export/xml | Export fields from from PDF in storage to XML file in storage.
*PdfApi* | [**PutFieldsFlatten**](docs/PdfApi.md#putfieldsflatten) | **Put** /pdf/{name}/fields/flatten | Flatten form fields in document.
*PdfApi* | [**PutFileAttachmentAnnotation**](docs/PdfApi.md#putfileattachmentannotation) | **Put** /pdf/{name}/annotations/fileattachment/{annotationId} | Replace document FileAttachment annotation
*PdfApi* | [**PutFileAttachmentAnnotationDataExtract**](docs/PdfApi.md#putfileattachmentannotationdataextract) | **Put** /pdf/{name}/annotations/fileattachment/{annotationId}/data/extract | Extract document FileAttachment annotation content to storage
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
*PdfApi* | [**PutImportFieldsFromFdfInStorage**](docs/PdfApi.md#putimportfieldsfromfdfinstorage) | **Put** /pdf/{name}/import/fdf | Update fields from FDF file in storage.
*PdfApi* | [**PutImportFieldsFromXfdfInStorage**](docs/PdfApi.md#putimportfieldsfromxfdfinstorage) | **Put** /pdf/{name}/import/xfdf | Update fields from XFDF file in storage.
*PdfApi* | [**PutImportFieldsFromXmlInStorage**](docs/PdfApi.md#putimportfieldsfromxmlinstorage) | **Put** /pdf/{name}/import/xml | Update fields from XML file in storage.
*PdfApi* | [**PutInkAnnotation**](docs/PdfApi.md#putinkannotation) | **Put** /pdf/{name}/annotations/ink/{annotationId} | Replace document ink annotation
*PdfApi* | [**PutLaTeXInStorageToPdf**](docs/PdfApi.md#putlatexinstoragetopdf) | **Put** /pdf/{name}/create/latex | Convert TeX file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutLineAnnotation**](docs/PdfApi.md#putlineannotation) | **Put** /pdf/{name}/annotations/line/{annotationId} | Replace document line annotation
*PdfApi* | [**PutLinkAnnotation**](docs/PdfApi.md#putlinkannotation) | **Put** /pdf/{name}/links/{linkId} | Replace document page link annotations
*PdfApi* | [**PutListBoxField**](docs/PdfApi.md#putlistboxfield) | **Put** /pdf/{name}/fields/listbox/{fieldName} | Replace document listbox field
*PdfApi* | [**PutMarkdownInStorageToPdf**](docs/PdfApi.md#putmarkdowninstoragetopdf) | **Put** /pdf/{name}/create/markdown | Convert MD file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutMergeDocuments**](docs/PdfApi.md#putmergedocuments) | **Put** /pdf/{name}/merge | Merge a list of documents.
*PdfApi* | [**PutMhtInStorageToPdf**](docs/PdfApi.md#putmhtinstoragetopdf) | **Put** /pdf/{name}/create/mht | Convert MHT file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutMovieAnnotation**](docs/PdfApi.md#putmovieannotation) | **Put** /pdf/{name}/annotations/movie/{annotationId} | Replace document movie annotation
*PdfApi* | [**PutPageAddStamp**](docs/PdfApi.md#putpageaddstamp) | **Put** /pdf/{name}/pages/{pageNumber}/stamp | Add page stamp.
*PdfApi* | [**PutPageConvertToBmp**](docs/PdfApi.md#putpageconverttobmp) | **Put** /pdf/{name}/pages/{pageNumber}/convert/bmp | Convert document page to bmp image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToEmf**](docs/PdfApi.md#putpageconverttoemf) | **Put** /pdf/{name}/pages/{pageNumber}/convert/emf | Convert document page to emf image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToGif**](docs/PdfApi.md#putpageconverttogif) | **Put** /pdf/{name}/pages/{pageNumber}/convert/gif | Convert document page to gif image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToJpeg**](docs/PdfApi.md#putpageconverttojpeg) | **Put** /pdf/{name}/pages/{pageNumber}/convert/jpeg | Convert document page to Jpeg image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToPng**](docs/PdfApi.md#putpageconverttopng) | **Put** /pdf/{name}/pages/{pageNumber}/convert/png | Convert document page to png image and upload resulting file to storage.
*PdfApi* | [**PutPageConvertToTiff**](docs/PdfApi.md#putpageconverttotiff) | **Put** /pdf/{name}/pages/{pageNumber}/convert/tiff | Convert document page to Tiff image and upload resulting file to storage.
*PdfApi* | [**PutPclInStorageToPdf**](docs/PdfApi.md#putpclinstoragetopdf) | **Put** /pdf/{name}/create/pcl | Convert PCL file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutPdfAInStorageToPdf**](docs/PdfApi.md#putpdfainstoragetopdf) | **Put** /pdf/{name}/create/pdfa | Convert PDFA file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutPdfInRequestToDoc**](docs/PdfApi.md#putpdfinrequesttodoc) | **Put** /pdf/convert/doc | Converts PDF document (in request content) to DOC format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToEpub**](docs/PdfApi.md#putpdfinrequesttoepub) | **Put** /pdf/convert/epub | Converts PDF document (in request content) to EPUB format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToHtml**](docs/PdfApi.md#putpdfinrequesttohtml) | **Put** /pdf/convert/html | Converts PDF document (in request content) to Html format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToLaTeX**](docs/PdfApi.md#putpdfinrequesttolatex) | **Put** /pdf/convert/latex | Converts PDF document (in request content) to TeX format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToMobiXml**](docs/PdfApi.md#putpdfinrequesttomobixml) | **Put** /pdf/convert/mobixml | Converts PDF document (in request content) to MOBIXML format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToPdfA**](docs/PdfApi.md#putpdfinrequesttopdfa) | **Put** /pdf/convert/pdfa | Converts PDF document (in request content) to PdfA format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToPptx**](docs/PdfApi.md#putpdfinrequesttopptx) | **Put** /pdf/convert/pptx | Converts PDF document (in request content) to PPTX format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToSvg**](docs/PdfApi.md#putpdfinrequesttosvg) | **Put** /pdf/convert/svg | Converts PDF document (in request content) to SVG format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToTeX**](docs/PdfApi.md#putpdfinrequesttotex) | **Put** /pdf/convert/tex | Converts PDF document (in request content) to TeX format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToTiff**](docs/PdfApi.md#putpdfinrequesttotiff) | **Put** /pdf/convert/tiff | Converts PDF document (in request content) to TIFF format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToXls**](docs/PdfApi.md#putpdfinrequesttoxls) | **Put** /pdf/convert/xls | Converts PDF document (in request content) to XLS format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToXlsx**](docs/PdfApi.md#putpdfinrequesttoxlsx) | **Put** /pdf/convert/xlsx | Converts PDF document (in request content) to XLSX format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToXml**](docs/PdfApi.md#putpdfinrequesttoxml) | **Put** /pdf/convert/xml | Converts PDF document (in request content) to XML format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInRequestToXps**](docs/PdfApi.md#putpdfinrequesttoxps) | **Put** /pdf/convert/xps | Converts PDF document (in request content) to XPS format and uploads resulting file to storage.
*PdfApi* | [**PutPdfInStorageToDoc**](docs/PdfApi.md#putpdfinstoragetodoc) | **Put** /pdf/{name}/convert/doc | Converts PDF document (located on storage) to DOC format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToEpub**](docs/PdfApi.md#putpdfinstoragetoepub) | **Put** /pdf/{name}/convert/epub | Converts PDF document (located on storage) to EPUB format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToHtml**](docs/PdfApi.md#putpdfinstoragetohtml) | **Put** /pdf/{name}/convert/html | Converts PDF document (located on storage) to Html format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToLaTeX**](docs/PdfApi.md#putpdfinstoragetolatex) | **Put** /pdf/{name}/convert/latex | Converts PDF document (located on storage) to TeX format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToMobiXml**](docs/PdfApi.md#putpdfinstoragetomobixml) | **Put** /pdf/{name}/convert/mobixml | Converts PDF document (located on storage) to MOBIXML format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToPdfA**](docs/PdfApi.md#putpdfinstoragetopdfa) | **Put** /pdf/{name}/convert/pdfa | Converts PDF document (located on storage) to PdfA format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToPptx**](docs/PdfApi.md#putpdfinstoragetopptx) | **Put** /pdf/{name}/convert/pptx | Converts PDF document (located on storage) to PPTX format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToSvg**](docs/PdfApi.md#putpdfinstoragetosvg) | **Put** /pdf/{name}/convert/svg | Converts PDF document (located on storage) to SVG format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToTeX**](docs/PdfApi.md#putpdfinstoragetotex) | **Put** /pdf/{name}/convert/tex | Converts PDF document (located on storage) to TeX format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToTiff**](docs/PdfApi.md#putpdfinstoragetotiff) | **Put** /pdf/{name}/convert/tiff | Converts PDF document (located on storage) to TIFF format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToXls**](docs/PdfApi.md#putpdfinstoragetoxls) | **Put** /pdf/{name}/convert/xls | Converts PDF document (located on storage) to XLS format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToXlsx**](docs/PdfApi.md#putpdfinstoragetoxlsx) | **Put** /pdf/{name}/convert/xlsx | Converts PDF document (located on storage) to XLSX format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToXml**](docs/PdfApi.md#putpdfinstoragetoxml) | **Put** /pdf/{name}/convert/xml | Converts PDF document (located on storage) to XML format and uploads resulting file to storage
*PdfApi* | [**PutPdfInStorageToXps**](docs/PdfApi.md#putpdfinstoragetoxps) | **Put** /pdf/{name}/convert/xps | Converts PDF document (located on storage) to XPS format and uploads resulting file to storage
*PdfApi* | [**PutPolyLineAnnotation**](docs/PdfApi.md#putpolylineannotation) | **Put** /pdf/{name}/annotations/polyline/{annotationId} | Replace document polyline annotation
*PdfApi* | [**PutPolygonAnnotation**](docs/PdfApi.md#putpolygonannotation) | **Put** /pdf/{name}/annotations/polygon/{annotationId} | Replace document polygon annotation
*PdfApi* | [**PutPopupAnnotation**](docs/PdfApi.md#putpopupannotation) | **Put** /pdf/{name}/annotations/popup/{annotationId} | Replace document popup annotation
*PdfApi* | [**PutPrivileges**](docs/PdfApi.md#putprivileges) | **Put** /pdf/{name}/privileges | Update privilege document.
*PdfApi* | [**PutPsInStorageToPdf**](docs/PdfApi.md#putpsinstoragetopdf) | **Put** /pdf/{name}/create/ps | Convert PS file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutRadioButtonField**](docs/PdfApi.md#putradiobuttonfield) | **Put** /pdf/{name}/fields/radiobutton/{fieldName} | Replace document RadioButton field
*PdfApi* | [**PutRedactionAnnotation**](docs/PdfApi.md#putredactionannotation) | **Put** /pdf/{name}/annotations/redaction/{annotationId} | Replace document redaction annotation
*PdfApi* | [**PutReplaceImage**](docs/PdfApi.md#putreplaceimage) | **Put** /pdf/{name}/images/{imageId} | Replace document image.
*PdfApi* | [**PutScreenAnnotation**](docs/PdfApi.md#putscreenannotation) | **Put** /pdf/{name}/annotations/screen/{annotationId} | Replace document screen annotation
*PdfApi* | [**PutScreenAnnotationDataExtract**](docs/PdfApi.md#putscreenannotationdataextract) | **Put** /pdf/{name}/annotations/screen/{annotationId}/data/extract | Extract document screen annotation content to storage
*PdfApi* | [**PutSearchableDocument**](docs/PdfApi.md#putsearchabledocument) | **Put** /pdf/{name}/ocr | Create searchable PDF document. Generate OCR layer for images in input PDF document.
*PdfApi* | [**PutSetProperty**](docs/PdfApi.md#putsetproperty) | **Put** /pdf/{name}/documentproperties/{propertyName} | Add/update document property.
*PdfApi* | [**PutSignatureField**](docs/PdfApi.md#putsignaturefield) | **Put** /pdf/{name}/fields/signature/{fieldName} | Replace document signature field.
*PdfApi* | [**PutSoundAnnotation**](docs/PdfApi.md#putsoundannotation) | **Put** /pdf/{name}/annotations/sound/{annotationId} | Replace document sound annotation
*PdfApi* | [**PutSoundAnnotationDataExtract**](docs/PdfApi.md#putsoundannotationdataextract) | **Put** /pdf/{name}/annotations/sound/{annotationId}/data/extract | Extract document sound annotation content to storage
*PdfApi* | [**PutSquareAnnotation**](docs/PdfApi.md#putsquareannotation) | **Put** /pdf/{name}/annotations/square/{annotationId} | Replace document square annotation
*PdfApi* | [**PutSquigglyAnnotation**](docs/PdfApi.md#putsquigglyannotation) | **Put** /pdf/{name}/annotations/squiggly/{annotationId} | Replace document squiggly annotation
*PdfApi* | [**PutStampAnnotation**](docs/PdfApi.md#putstampannotation) | **Put** /pdf/{name}/annotations/stamp/{annotationId} | Replace document stamp annotation
*PdfApi* | [**PutStampAnnotationDataExtract**](docs/PdfApi.md#putstampannotationdataextract) | **Put** /pdf/{name}/annotations/stamp/{annotationId}/data/extract | Extract document stamp annotation content to storage
*PdfApi* | [**PutStrikeOutAnnotation**](docs/PdfApi.md#putstrikeoutannotation) | **Put** /pdf/{name}/annotations/strikeout/{annotationId} | Replace document StrikeOut annotation
*PdfApi* | [**PutSvgInStorageToPdf**](docs/PdfApi.md#putsvginstoragetopdf) | **Put** /pdf/{name}/create/svg | Convert SVG file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutTable**](docs/PdfApi.md#puttable) | **Put** /pdf/{name}/tables/{tableId} | Replace document page table.
*PdfApi* | [**PutTeXInStorageToPdf**](docs/PdfApi.md#puttexinstoragetopdf) | **Put** /pdf/{name}/create/tex | Convert TeX file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutTextAnnotation**](docs/PdfApi.md#puttextannotation) | **Put** /pdf/{name}/annotations/text/{annotationId} | Replace document text annotation
*PdfApi* | [**PutTextBoxField**](docs/PdfApi.md#puttextboxfield) | **Put** /pdf/{name}/fields/textbox/{fieldName} | Replace document text box field
*PdfApi* | [**PutUnderlineAnnotation**](docs/PdfApi.md#putunderlineannotation) | **Put** /pdf/{name}/annotations/underline/{annotationId} | Replace document underline annotation
*PdfApi* | [**PutUpdateField**](docs/PdfApi.md#putupdatefield) | **Put** /pdf/{name}/fields/{fieldName} | Update field.
*PdfApi* | [**PutUpdateFields**](docs/PdfApi.md#putupdatefields) | **Put** /pdf/{name}/fields | Update fields.
*PdfApi* | [**PutWebInStorageToPdf**](docs/PdfApi.md#putwebinstoragetopdf) | **Put** /pdf/{name}/create/web | Convert web page to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutXfaPdfInRequestToAcroForm**](docs/PdfApi.md#putxfapdfinrequesttoacroform) | **Put** /pdf/convert/xfatoacroform | Converts PDF document which contatins XFA form (in request content) to PDF with AcroForm and uploads resulting file to storage.
*PdfApi* | [**PutXfaPdfInStorageToAcroForm**](docs/PdfApi.md#putxfapdfinstoragetoacroform) | **Put** /pdf/{name}/convert/xfatoacroform | Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and uploads resulting file to storage
*PdfApi* | [**PutXmlInStorageToPdf**](docs/PdfApi.md#putxmlinstoragetopdf) | **Put** /pdf/{name}/create/xml | Convert XML file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutXpsInStorageToPdf**](docs/PdfApi.md#putxpsinstoragetopdf) | **Put** /pdf/{name}/create/xps | Convert XPS file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**PutXslFoInStorageToPdf**](docs/PdfApi.md#putxslfoinstoragetopdf) | **Put** /pdf/{name}/create/xslfo | Convert XslFo file (located on storage) to PDF format and upload resulting file to storage. 
*PdfApi* | [**StorageExists**](docs/PdfApi.md#storageexists) | **Get** /pdf/storage/{storageName}/exist | Check if storage exists
*PdfApi* | [**UploadFile**](docs/PdfApi.md#uploadfile) | **Put** /pdf/storage/file/{path} | Upload file


## Documentation For Models

 - [AnnotationFlags](docs/AnnotationFlags.md)
 - [AnnotationState](docs/AnnotationState.md)
 - [AnnotationType](docs/AnnotationType.md)
 - [AntialiasingProcessingType](docs/AntialiasingProcessingType.md)
 - [AsposeResponse](docs/AsposeResponse.md)
 - [Border](docs/Border.md)
 - [BorderCornerStyle](docs/BorderCornerStyle.md)
 - [BorderEffect](docs/BorderEffect.md)
 - [BorderInfo](docs/BorderInfo.md)
 - [BorderStyle](docs/BorderStyle.md)
 - [BoxStyle](docs/BoxStyle.md)
 - [CapStyle](docs/CapStyle.md)
 - [CaptionPosition](docs/CaptionPosition.md)
 - [CaretSymbol](docs/CaretSymbol.md)
 - [Cell](docs/Cell.md)
 - [CellRecognized](docs/CellRecognized.md)
 - [Color](docs/Color.md)
 - [ColorDepth](docs/ColorDepth.md)
 - [ColumnAdjustment](docs/ColumnAdjustment.md)
 - [CompressionType](docs/CompressionType.md)
 - [CryptoAlgorithm](docs/CryptoAlgorithm.md)
 - [Dash](docs/Dash.md)
 - [DefaultPageConfig](docs/DefaultPageConfig.md)
 - [Direction](docs/Direction.md)
 - [DiscUsage](docs/DiscUsage.md)
 - [DocFormat](docs/DocFormat.md)
 - [DocMdpAccessPermissionType](docs/DocMdpAccessPermissionType.md)
 - [DocRecognitionMode](docs/DocRecognitionMode.md)
 - [DocumentConfig](docs/DocumentConfig.md)
 - [DocumentPrivilege](docs/DocumentPrivilege.md)
 - [EpubRecognitionMode](docs/EpubRecognitionMode.md)
 - [ErrorDetails](docs/ErrorDetails.md)
 - [FieldType](docs/FieldType.md)
 - [FileIcon](docs/FileIcon.md)
 - [FileVersions](docs/FileVersions.md)
 - [FilesList](docs/FilesList.md)
 - [FilesUploadResult](docs/FilesUploadResult.md)
 - [FontEncodingRules](docs/FontEncodingRules.md)
 - [FontSavingModes](docs/FontSavingModes.md)
 - [FontStyles](docs/FontStyles.md)
 - [FreeTextIntent](docs/FreeTextIntent.md)
 - [GraphInfo](docs/GraphInfo.md)
 - [HorizontalAlignment](docs/HorizontalAlignment.md)
 - [HtmlDocumentType](docs/HtmlDocumentType.md)
 - [HtmlMarkupGenerationModes](docs/HtmlMarkupGenerationModes.md)
 - [ImageFragment](docs/ImageFragment.md)
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
 - [ModelError](docs/ModelError.md)
 - [ObjectExist](docs/ObjectExist.md)
 - [OptimizeOptions](docs/OptimizeOptions.md)
 - [Option](docs/Option.md)
 - [OutputFormat](docs/OutputFormat.md)
 - [PageLayout](docs/PageLayout.md)
 - [PageMode](docs/PageMode.md)
 - [PageWordCount](docs/PageWordCount.md)
 - [Paragraph](docs/Paragraph.md)
 - [PartsEmbeddingModes](docs/PartsEmbeddingModes.md)
 - [PdfAType](docs/PdfAType.md)
 - [PermissionsFlags](docs/PermissionsFlags.md)
 - [Point](docs/Point.md)
 - [PolyIntent](docs/PolyIntent.md)
 - [Position](docs/Position.md)
 - [RasterImagesSavingModes](docs/RasterImagesSavingModes.md)
 - [Rectangle](docs/Rectangle.md)
 - [Rotation](docs/Rotation.md)
 - [Row](docs/Row.md)
 - [RowRecognized](docs/RowRecognized.md)
 - [Segment](docs/Segment.md)
 - [ShapeType](docs/ShapeType.md)
 - [Signature](docs/Signature.md)
 - [SignatureCustomAppearance](docs/SignatureCustomAppearance.md)
 - [SignatureType](docs/SignatureType.md)
 - [SoundEncoding](docs/SoundEncoding.md)
 - [SoundIcon](docs/SoundIcon.md)
 - [SplitResult](docs/SplitResult.md)
 - [Stamp](docs/Stamp.md)
 - [StampIcon](docs/StampIcon.md)
 - [StampType](docs/StampType.md)
 - [StorageExist](docs/StorageExist.md)
 - [StorageFile](docs/StorageFile.md)
 - [TableBroken](docs/TableBroken.md)
 - [TextHorizontalAlignment](docs/TextHorizontalAlignment.md)
 - [TextIcon](docs/TextIcon.md)
 - [TextLine](docs/TextLine.md)
 - [TextRect](docs/TextRect.md)
 - [TextRects](docs/TextRects.md)
 - [TextReplace](docs/TextReplace.md)
 - [TextReplaceListRequest](docs/TextReplaceListRequest.md)
 - [TextState](docs/TextState.md)
 - [TextStyle](docs/TextStyle.md)
 - [TimestampSettings](docs/TimestampSettings.md)
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
 - [Bookmark](docs/Bookmark.md)
 - [BookmarkResponse](docs/BookmarkResponse.md)
 - [Bookmarks](docs/Bookmarks.md)
 - [BookmarksResponse](docs/BookmarksResponse.md)
 - [CaretAnnotationResponse](docs/CaretAnnotationResponse.md)
 - [CaretAnnotations](docs/CaretAnnotations.md)
 - [CaretAnnotationsResponse](docs/CaretAnnotationsResponse.md)
 - [CheckBoxFieldResponse](docs/CheckBoxFieldResponse.md)
 - [CheckBoxFields](docs/CheckBoxFields.md)
 - [CheckBoxFieldsResponse](docs/CheckBoxFieldsResponse.md)
 - [CircleAnnotationResponse](docs/CircleAnnotationResponse.md)
 - [CircleAnnotations](docs/CircleAnnotations.md)
 - [CircleAnnotationsResponse](docs/CircleAnnotationsResponse.md)
 - [ComboBoxFieldResponse](docs/ComboBoxFieldResponse.md)
 - [ComboBoxFields](docs/ComboBoxFields.md)
 - [ComboBoxFieldsResponse](docs/ComboBoxFieldsResponse.md)
 - [DisplayProperties](docs/DisplayProperties.md)
 - [DisplayPropertiesResponse](docs/DisplayPropertiesResponse.md)
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
 - [FileAttachmentAnnotationResponse](docs/FileAttachmentAnnotationResponse.md)
 - [FileAttachmentAnnotations](docs/FileAttachmentAnnotations.md)
 - [FileAttachmentAnnotationsResponse](docs/FileAttachmentAnnotationsResponse.md)
 - [FileVersion](docs/FileVersion.md)
 - [FormField](docs/FormField.md)
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
 - [ListBoxFieldResponse](docs/ListBoxFieldResponse.md)
 - [ListBoxFields](docs/ListBoxFields.md)
 - [ListBoxFieldsResponse](docs/ListBoxFieldsResponse.md)
 - [MovieAnnotationResponse](docs/MovieAnnotationResponse.md)
 - [MovieAnnotations](docs/MovieAnnotations.md)
 - [MovieAnnotationsResponse](docs/MovieAnnotationsResponse.md)
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
 - [RadioButtonFieldResponse](docs/RadioButtonFieldResponse.md)
 - [RadioButtonFields](docs/RadioButtonFields.md)
 - [RadioButtonFieldsResponse](docs/RadioButtonFieldsResponse.md)
 - [RedactionAnnotationResponse](docs/RedactionAnnotationResponse.md)
 - [RedactionAnnotations](docs/RedactionAnnotations.md)
 - [RedactionAnnotationsResponse](docs/RedactionAnnotationsResponse.md)
 - [ScreenAnnotationResponse](docs/ScreenAnnotationResponse.md)
 - [ScreenAnnotations](docs/ScreenAnnotations.md)
 - [ScreenAnnotationsResponse](docs/ScreenAnnotationsResponse.md)
 - [SignatureFieldResponse](docs/SignatureFieldResponse.md)
 - [SignatureFields](docs/SignatureFields.md)
 - [SignatureFieldsResponse](docs/SignatureFieldsResponse.md)
 - [SignatureVerifyResponse](docs/SignatureVerifyResponse.md)
 - [SoundAnnotationResponse](docs/SoundAnnotationResponse.md)
 - [SoundAnnotations](docs/SoundAnnotations.md)
 - [SoundAnnotationsResponse](docs/SoundAnnotationsResponse.md)
 - [SplitResultDocument](docs/SplitResultDocument.md)
 - [SplitResultResponse](docs/SplitResultResponse.md)
 - [SquareAnnotationResponse](docs/SquareAnnotationResponse.md)
 - [SquareAnnotations](docs/SquareAnnotations.md)
 - [SquareAnnotationsResponse](docs/SquareAnnotationsResponse.md)
 - [SquigglyAnnotationResponse](docs/SquigglyAnnotationResponse.md)
 - [SquigglyAnnotations](docs/SquigglyAnnotations.md)
 - [SquigglyAnnotationsResponse](docs/SquigglyAnnotationsResponse.md)
 - [StampAnnotationResponse](docs/StampAnnotationResponse.md)
 - [StampAnnotations](docs/StampAnnotations.md)
 - [StampAnnotationsResponse](docs/StampAnnotationsResponse.md)
 - [StampBase](docs/StampBase.md)
 - [StampInfo](docs/StampInfo.md)
 - [StampsInfo](docs/StampsInfo.md)
 - [StampsInfoResponse](docs/StampsInfoResponse.md)
 - [StrikeOutAnnotationResponse](docs/StrikeOutAnnotationResponse.md)
 - [StrikeOutAnnotations](docs/StrikeOutAnnotations.md)
 - [StrikeOutAnnotationsResponse](docs/StrikeOutAnnotationsResponse.md)
 - [Table](docs/Table.md)
 - [TableRecognized](docs/TableRecognized.md)
 - [TableRecognizedResponse](docs/TableRecognizedResponse.md)
 - [TablesRecognized](docs/TablesRecognized.md)
 - [TablesRecognizedResponse](docs/TablesRecognizedResponse.md)
 - [TextAnnotationResponse](docs/TextAnnotationResponse.md)
 - [TextAnnotations](docs/TextAnnotations.md)
 - [TextAnnotationsResponse](docs/TextAnnotationsResponse.md)
 - [TextBoxFieldResponse](docs/TextBoxFieldResponse.md)
 - [TextBoxFields](docs/TextBoxFields.md)
 - [TextBoxFieldsResponse](docs/TextBoxFieldsResponse.md)
 - [TextRectsResponse](docs/TextRectsResponse.md)
 - [TextReplaceResponse](docs/TextReplaceResponse.md)
 - [UnderlineAnnotationResponse](docs/UnderlineAnnotationResponse.md)
 - [UnderlineAnnotations](docs/UnderlineAnnotations.md)
 - [UnderlineAnnotationsResponse](docs/UnderlineAnnotationsResponse.md)
 - [WordCountResponse](docs/WordCountResponse.md)
 - [AnnotationInfo](docs/AnnotationInfo.md)
 - [CheckBoxField](docs/CheckBoxField.md)
 - [ChoiceField](docs/ChoiceField.md)
 - [ImageFooter](docs/ImageFooter.md)
 - [ImageHeader](docs/ImageHeader.md)
 - [ImageStamp](docs/ImageStamp.md)
 - [MarkupAnnotation](docs/MarkupAnnotation.md)
 - [MovieAnnotation](docs/MovieAnnotation.md)
 - [PageNumberStamp](docs/PageNumberStamp.md)
 - [PdfPageStamp](docs/PdfPageStamp.md)
 - [PopupAnnotation](docs/PopupAnnotation.md)
 - [RadioButtonOptionField](docs/RadioButtonOptionField.md)
 - [RedactionAnnotation](docs/RedactionAnnotation.md)
 - [ScreenAnnotation](docs/ScreenAnnotation.md)
 - [SignatureField](docs/SignatureField.md)
 - [TextBoxField](docs/TextBoxField.md)
 - [TextFooter](docs/TextFooter.md)
 - [TextHeader](docs/TextHeader.md)
 - [TextStamp](docs/TextStamp.md)
 - [CaretAnnotation](docs/CaretAnnotation.md)
 - [ComboBoxField](docs/ComboBoxField.md)
 - [CommonFigureAnnotation](docs/CommonFigureAnnotation.md)
 - [FileAttachmentAnnotation](docs/FileAttachmentAnnotation.md)
 - [FreeTextAnnotation](docs/FreeTextAnnotation.md)
 - [HighlightAnnotation](docs/HighlightAnnotation.md)
 - [InkAnnotation](docs/InkAnnotation.md)
 - [LineAnnotation](docs/LineAnnotation.md)
 - [ListBoxField](docs/ListBoxField.md)
 - [PolyAnnotation](docs/PolyAnnotation.md)
 - [PopupAnnotationWithParent](docs/PopupAnnotationWithParent.md)
 - [RadioButtonField](docs/RadioButtonField.md)
 - [SoundAnnotation](docs/SoundAnnotation.md)
 - [SquigglyAnnotation](docs/SquigglyAnnotation.md)
 - [StampAnnotation](docs/StampAnnotation.md)
 - [StrikeOutAnnotation](docs/StrikeOutAnnotation.md)
 - [TextAnnotation](docs/TextAnnotation.md)
 - [UnderlineAnnotation](docs/UnderlineAnnotation.md)
 - [CircleAnnotation](docs/CircleAnnotation.md)
 - [PolyLineAnnotation](docs/PolyLineAnnotation.md)
 - [PolygonAnnotation](docs/PolygonAnnotation.md)
 - [SquareAnnotation](docs/SquareAnnotation.md)

