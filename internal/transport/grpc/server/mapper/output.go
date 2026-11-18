package mapper

import (
	linkv1 "github.com/dockloud/contract/gen/link/v1"
	"github.com/dockloud/link/internal/domain/link"
	usecase "github.com/dockloud/link/internal/usecase/link"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func domainLinkToPBLink(dom link.Link) *linkv1.Link {
	return &linkv1.Link{
		Token:         dom.Token,
		DocumentId:    int64(dom.DocumentID),
		ReadAllowed:   dom.ReadAllowed,
		WriteAllowed:  dom.WriteAllowed,
		DeleteAllowed: dom.DeleteAllowed,
		CreatedAt:     timestamppb.New(dom.CreatedAt),
		ExpiresAt:     nil,
	}
}

func ToGetUserLinksResponse(out *usecase.GetUserLinksOutput) *linkv1.GetLinksResponse {
	var links []*linkv1.Link
	for _, l := range out.Links {
		links = append(links, domainLinkToPBLink(l))
	}

	return &linkv1.GetLinksResponse{
		Links: links,
	}
}

func ToCreateLinkResponse(out *usecase.CreateLinkOutput) *linkv1.CreateLinkResponse {
	l := out.Link

	var mappedTimestamp *timestamppb.Timestamp
	if l.ExpiresAt != nil {
		mappedTimestamp = timestamppb.New(*l.ExpiresAt)
	} else {
		mappedTimestamp = nil
	}

	return &linkv1.CreateLinkResponse{
		Link: &linkv1.Link{
			Token:         l.Token,
			DocumentId:    int64(l.DocumentID),
			ReadAllowed:   l.ReadAllowed,
			WriteAllowed:  l.WriteAllowed,
			DeleteAllowed: l.DeleteAllowed,
			CreatedAt:     timestamppb.New(l.CreatedAt),
			ExpiresAt:     mappedTimestamp,
		},
	}
}

func ToRevokeLinkResponse(out *usecase.RevokeLinkOutput) *linkv1.RevokeLinkResponse {
	return &linkv1.RevokeLinkResponse{}
}
