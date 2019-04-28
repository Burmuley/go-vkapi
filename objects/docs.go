package objects

// DocsDoc represents `docs_doc` API object
type DocsDoc struct {
	AccessKey string         `json:"access_key"`
	Date      int            `json:"date"`
	Extension string         `json:"ext"`
	ID        int            `json:"id"`
	Licensed  BaseBoolInt    `json:"is_licensed"`
	OwnerID   int            `json:"owner_id"`
	Preview   DocsDocPreview `json:"preview"`
	Size      int            `json:"size"`
	Title     string         `json:"title"`
	Type      int            `json:"type"`
	Url       string         `json:"url"`
}

// DocsDocPreview represents `docs_doc_preview` API object
type DocsDocPreview struct {
	Photo DocsDocPreviewPhoto `json:"photo"`
	Video DocsDocPreviewVideo `json:"video"`
}

// DocsDocPreviewPhoto represents `docs_doc_preview_photo` API object
type DocsDocPreviewPhoto struct {
	Sizes []PhotosPhotoSizes `json:"sizes"`
}

// DocsDocPreviewVideo represents `docs_doc_preview_video` API object
type DocsDocPreviewVideo struct {
	Size   int    `json:"file_size"`
	Source string `json:"src"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// DocsDocTypes represents `docs_doc_types` API object
type DocsDocTypes struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Count int    `json:"count"`
}

// DocsDocUploadResponse represents `docs_doc_upload_response` API object
type DocsDocUploadResponse struct {
	File string `json:"file"`
}
