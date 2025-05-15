package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	asposepdfcloud "github.com/aspose-pdf-cloud/aspose-pdf-cloud-go/v25"
)

// Set file names and folders
const (
	REMOTE_FOLDER = "Your_Temp_Pdf_Cloud"
	LOCAL_FOLDER  = "c:\\Samples"
	PDF_DOCUMENT  = "sample.pdf"

	JPEG_FILE = "sample.jpg"
	BMP_FILE  = "sample.bmp"
	PNG_FILE  = "sample.png"
	GIF_FILE  = "sample.gif"

	PDF_BMP_OUTPUT  = "output_bmp.pdf"
	PDF_JPEG_OUTPUT = "output_jpeg.pdf"
	PDF_PNG_OUTPUT  = "output_png.pdf"
	PDF_GIF_OUTPUT  = "output_gif.pdf"

	HTML_FILE_NAME  = "sample.html"
	HTML_ZIP_FILE   = "sample_html.zip"
	PDF_HTML_OUTPUT = "output_html.pdf"

	MHTML_FILE_NAME  = "sample.mht"
	PDF_MHTML_OUTPUT = "output_mht.pdf"

	MD_FILE_NAME  = "sample.md"
	PDF_MD_OUTPUT = "output_md.pdf"

	EPUB_FILE_NAME  = "sample.epub"
	PDF_EPUB_OUTPUT = "output_epub.pdf"

	PCL_FILE_NAME  = "sample.pcl"
	PDF_PCL_OUTPUT = "output_pcl.pdf"

	PS_FILE_NAME  = "sample.ps"
	PDF_PS_OUTPUT = "output_ps.pdf"

	PDF_OUTPUT      = "output_sample.pdf"
	DOC_OUTPUT      = "output_sample.doc"
	DOCX_OUTPUT     = "output_sample.docx"
	EPUB_OUTPUT     = "output_sample.epub"
	HTML_ZIP_OUTPUT = "output_sample_html.zip"
	TIFF_OUTPUT     = "output_sample.tiff"
	SVG_ZIP_OUTPUT  = "output_sample_svg.zip"
	XLS_OUTPUT      = "output_sample.xls"
	XLSX_OUTPUT     = "output_sample.xlsx"
	PPTX_OUTPUT     = "output_sample.pptx"
	XML_OUTPUT      = "output_sample.xml"
	MXML_OUTPUT     = "output_sample.mobi"
	TEX_OUTPUT      = "output_sample.tex"
	XPS_OUTPUT      = "output_sample.xps"
)

// Initialize credentials and Rest API service
func initPdfApi() *asposepdfcloud.PdfApiService {
	AppSID := "********"
	AppKey := "********"

	pdfApi := asposepdfcloud.NewPdfApiService(AppSID, AppKey, "")
	return pdfApi
}

// Upload local file to the remote folder with check errors
func uploadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	file, err := os.Open(filepath.Join(LOCAL_FOLDER, name))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		result, httpResponse, err := pdf_api.UploadFile(path.Join(REMOTE_FOLDER, name), file, args)
		if err != nil {
			fmt.Println(err.Error())
		} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
			fmt.Println("Unexpected error!")
		} else {
			fmt.Println(result)
		}
	}
}

// Download file from remote folder and save it locally with check errors
func downloadFile(pdf_api *asposepdfcloud.PdfApiService, name string) {
	args := map[string]interface{}{
		"folder": REMOTE_FOLDER,
	}
	result_data, httpResponse, err := pdf_api.DownloadFile(name, args)
	if err != nil {
		fmt.Println(err.Error())
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Println("Unexpected error!")
	} else {
		fmt.Println("File '" + name + "'successfully downloaded.")
		simpleFileName := filepath.Base(name)
		fileName := path.Join(LOCAL_FOLDER, simpleFileName)
		f, _ := os.Create(fileName)
		_, _ = f.Write(result_data)
	}
}

func main() {
	pdfApi := initPdfApi()

	uploadFile(pdfApi, MD_FILE_NAME)
	convertMarkdownToPdf(pdfApi, MD_FILE_NAME, PDF_MD_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_MD_OUTPUT)

	uploadFile(pdfApi, EPUB_FILE_NAME)
	convertEpubToPdf(pdfApi, EPUB_FILE_NAME, PDF_EPUB_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_EPUB_OUTPUT)

	uploadFile(pdfApi, HTML_FILE_NAME)
	convertHtmlToPdf(pdfApi, HTML_FILE_NAME, HTML_ZIP_FILE, PDF_HTML_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_HTML_OUTPUT)

	uploadFile(pdfApi, MHTML_FILE_NAME)
	convertMhtmlToPdf(pdfApi, MHTML_FILE_NAME, PDF_MHTML_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_MHTML_OUTPUT)

	uploadFile(pdfApi, PCL_FILE_NAME)
	convertPclToPdf(pdfApi, PCL_FILE_NAME, PDF_PCL_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_PCL_OUTPUT)

	images := []string{BMP_FILE, JPEG_FILE, PNG_FILE}
	converImagesToPdf(pdfApi, images, PDF_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_OUTPUT)

	converBmpToPdf(pdfApi, BMP_FILE, PDF_BMP_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_BMP_OUTPUT)

	converJpegToPdf(pdfApi, JPEG_FILE, PDF_JPEG_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_JPEG_OUTPUT)

	converPngToPdf(pdfApi, PNG_FILE, PDF_PNG_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_PNG_OUTPUT)

	uploadFile(pdfApi, PS_FILE_NAME)
	convertPsToPdf(pdfApi, PS_FILE_NAME, PDF_PS_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_PS_OUTPUT)

	converGifToPdf(pdfApi, GIF_FILE, PDF_GIF_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, PDF_GIF_OUTPUT)

	uploadFile(pdfApi, PDF_DOCUMENT)
	convertPdfToDoc(pdfApi, PDF_DOCUMENT, DOC_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, DOC_OUTPUT))

	convertPdfToWord(pdfApi, PDF_DOCUMENT, DOCX_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, DOCX_OUTPUT))

	convertPdfToEpub(pdfApi, PDF_DOCUMENT, EPUB_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, EPUB_OUTPUT))

	convertPdfToHtml(pdfApi, PDF_DOCUMENT, HTML_ZIP_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, HTML_ZIP_OUTPUT))

	convertPdfToTiff(pdfApi, PDF_DOCUMENT, TIFF_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, TIFF_OUTPUT))

	convertPdfToSvg(pdfApi, PDF_DOCUMENT, SVG_ZIP_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, SVG_ZIP_OUTPUT))

	convertPdfToXls(pdfApi, PDF_DOCUMENT, XLS_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, XLS_OUTPUT))

	convertPdfToExcel(pdfApi, PDF_DOCUMENT, XLSX_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, XLSX_OUTPUT))

	convertPdfToPowerpoint(pdfApi, PDF_DOCUMENT, PPTX_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, PPTX_OUTPUT))

	convertPdfToMobi(pdfApi, PDF_DOCUMENT, MXML_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, MXML_OUTPUT))

	convertPdfToTex(pdfApi, PDF_DOCUMENT, TEX_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, TEX_OUTPUT))

	convertPdfToXps(pdfApi, PDF_DOCUMENT, XPS_OUTPUT, REMOTE_FOLDER)
	downloadFile(pdfApi, path.Join(REMOTE_FOLDER, XPS_OUTPUT))
}
