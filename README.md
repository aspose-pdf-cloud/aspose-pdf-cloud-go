![](https://img.shields.io/badge/api-v3.0-lightgrey) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/aspose-pdf-cloud/aspose-pdf-cloud-go)   [![GitHub license](https://img.shields.io/github/license/aspose-pdf-cloud/aspose-pdf-cloud-go)](https://github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/blob/master/LICENSE)

# Go SDK for Aspose.PDF Cloud REST API
[Aspose.PDF Cloud](https://products.aspose.cloud/pdf) is a true REST API that enables you to perform a wide range of document processing operations including creation, manipulation, conversion and rendering of PDF documents in the cloud.

Our Cloud SDKs are wrappers around REST API in various programming languages, allowing you to process documents in language of your choice quickly and easily, gaining all benefits of strong types and IDE highlights. This repository contains new generation SDKs for Aspose.PDF Cloud and examples.

These SDKs are now fully supported. If you have any questions, see any bugs or have enhancement request, feel free to reach out to us at [Free Support Forums](https://forum.aspose.cloud/c/pdf).

Extract Text & Images of a PDF document online https://products.aspose.app/pdf/parser.

## Enhancements in Version 24.9
- Develop DeleteDocumentLayer Method.
- Develop GetDocumentLayers Method.
- How to extract PDF layer elements and create a new PDF
- A new version of Aspose.PDF Cloud was prepared using the latest version of Aspose.PDF for .NET.

## Installation
```
    go get -u github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v24
```

## Getting Started
Please follow the [installation](#installation) instruction and execute the following Go code:

## Get PDF Page Circle Annotations
```
        import "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v24"

        ...

	// Get your ClientId and ClientSecret from https://dashboard.aspose.cloud (free registration required).
	pdfApi := asposepdfcloud.NewPdfApiService("MY_CLIENT_ID", "MY_CLIENT_SECRET", "")
	args := map[string]interface{} {
		"folder": "path/to/remote/folder",
	}
	return pdfApi.GetDocumentCircleAnnotations("PdfWithAnnotations.pdf", args)
```

## SelfHost Aspose.PDF Cloud
Instead of **NewPdfApiService** use **NewSelfHostPdfApiService** function to create PdfApiService instance:
```
	pdfApi := asposepdfcloud.NewSelfHostPdfApiService(<base URL of SelfHost Aspose.PDF Cloud>)
```

## Unit Tests
Aspose PDF Cloud SDK includes a suite of unit tests within the "test" subdirectory. These Unit Tests also serves as examples of how to use the Aspose PDF Cloud SDK.

## Licensing
All Aspose.PDF Cloud SDKs are licensed under [MIT License](LICENSE).
