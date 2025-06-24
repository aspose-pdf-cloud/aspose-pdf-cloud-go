package main

func main() {
	pdfApi := initPdfApi(PDF_API_SID, PDF_API_KEY)

	СomparePdf(pdfApi, PDF_DOCUMENT_1, PDF_DOCUMENT_2, PDF_OUTPUT, REMOTE_FOLDER)
}
