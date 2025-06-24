package main

func main() {
	pdfApi := initPdfApi()

	getDocumentAttachments(pdfApi, PDF_DOCUMENT_WITH_ATTACH, REMOTE_FOLDER)
	getAttachmentByIndex(pdfApi, PDF_DOCUMENT_WITH_ATTACH, 1, REMOTE_FOLDER, LOCAL_FOLDER)

	appendAttachment(pdfApi, PDF_DOCUMENT, NEW_ATTACHMENT_FILE, NEW_ATTACHMENT_DECRIPTION, NEW_ATTACHMENT_MIME, REMOTE_FOLDER)
}
