package mapper

import (
	pb "github.com/dockloud/contract/gen/link/v1"
	"github.com/dockloud/link/internal/usecase/link"
)

func ToGetUserLinksInput(_ *pb.GetLinksRequest, uid string) *link.GetUserLinksInput {
	return &link.GetUserLinksInput{
		UID: uid,
	}
}

func ToGetDocumentInput(req *pb.GetDocumentByLinkRequest) *link.GetDocumentInput {
	return &link.GetDocumentInput{
		Token: req.Token,
	}
}

func ToCreateLinkInput(req *pb.CreateLinkRequest, uid string) *link.CreateLinkInput {
	var hoursToExpire *int
	if req.HoursToExpire != nil {
		hoursAsInt := int(*req.HoursToExpire)
		hoursToExpire = &hoursAsInt
	}

	return &link.CreateLinkInput{
		UID:           uid,
		DocumentID:    int(req.DocumentId),
		ReadAllowed:   req.ReadAllowed,
		WriteAllowed:  req.WriteAllowed,
		DeleteAllowed: req.DeleteAllowed,
		HoursToExpire: hoursToExpire,
	}
}
