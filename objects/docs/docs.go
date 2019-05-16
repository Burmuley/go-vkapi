package docs

import (
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/photos"
)

/////////////////////////////////////////////////////////////
// Documents related API objects	                       //
/////////////////////////////////////////////////////////////

// Doc represents `docs_doc` API object
type Doc struct {
	AccessKey string       `json:"access_key"`
	Date      int          `json:"date"`
	Extension string       `json:"ext"`
	ID        int          `json:"id"`
	Licensed  base.BoolInt `json:"is_licensed"`
	OwnerID   int          `json:"owner_id"`
	Preview   DocPreview   `json:"preview"`
	Size      int          `json:"size"`
	Title     string       `json:"title"`
	Type      int          `json:"type"`
	Url       string       `json:"url"`
}

// DocPreview represents `docs_doc_preview` API object
type DocPreview struct {
	Photo DocPreviewPhoto `json:"photo"`
	Video DocPreviewVideo `json:"video"`
}

// DocPreviewPhoto represents `docs_doc_preview_photo` API object
type DocPreviewPhoto struct {
	Sizes []photos.PhotoSizes `json:"sizes"`
}

// DocPreviewVideo represents `docs_doc_preview_video` API object
type DocPreviewVideo struct {
	Size   int    `json:"file_size"`
	Source string `json:"src"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// DocTypes represents `docs_doc_types` API object
type DocTypes struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Count int    `json:"count"`
}

// DocUploadResponse represents `docs_doc_upload_response` API object
type DocUploadResponse struct {
	File string `json:"file"`
}
