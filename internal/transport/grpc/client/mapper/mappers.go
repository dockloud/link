package mapper

import (
	docv1 "github.com/dockloud/contract/gen/document/v1"
	linkv1 "github.com/dockloud/contract/gen/link/v1"
)

func ToGetDocumentByLinkResponse(out *docv1.GetUnsafeDocumentResponse) *linkv1.GetDocumentByLinkResponse {
	return &linkv1.GetDocumentByLinkResponse{
		DocumentId: out.DocumentId,
		OwnerUid:   out.OwnerUid,
		Name:       out.Name,
		Content:    out.Content,
		IsPrivate:  out.IsPrivate,
		CreatedAt:  out.CreatedAt,
		EditedAt:   out.EditedAt,
	}
}

func ToUpdateDocumentByLinkResponse(out *docv1.UpdateUnsafeDocumentResponse) *linkv1.UpdateDocumentByLinkResponse {
	return &linkv1.UpdateDocumentByLinkResponse{
		DocumentId: out.DocumentId,
		Name:       out.Name,
		Content:    out.Content,
		IsPrivate:  out.IsPrivate,
		EditedAt:   out.EditedAt,
	}
}

func ToDeleteDocumentByLinkResponse(out *docv1.DeleteUnsafeDocumentResponse) *linkv1.DeleteDocumentByLinkResponse {
	return &linkv1.DeleteDocumentByLinkResponse{}
}
