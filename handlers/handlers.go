package handlers

import (
	"context"

	pb "github.com/mises-id/mises-websitesvc/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.WebsitesvcServer {
	return websitesvcService{}
}

type websitesvcService struct{}

func (s websitesvcService) WebsiteCategoryList(ctx context.Context, in *pb.WebsiteCategoryListRequest) (*pb.WebsiteCategoryListResponse, error) {
	var resp pb.WebsiteCategoryListResponse
	resp.Code = 0
	return &resp, nil
}
