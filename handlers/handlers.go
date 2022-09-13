package handlers

import (
	"context"

	"github.com/mises-id/mises-websitesvc/app/factory"
	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/app/models/search"
	"github.com/mises-id/mises-websitesvc/app/services/website"
	"github.com/mises-id/mises-websitesvc/app/services/website_category"
	"github.com/mises-id/mises-websitesvc/lib/pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"

	pb "github.com/mises-id/mises-websitesvc/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.WebsitesvcServer {
	return websitesvcService{}
}

type websitesvcService struct{}

func (s websitesvcService) WebsiteCategoryList(ctx context.Context, in *pb.WebsiteCategoryListRequest) (*pb.WebsiteCategoryListResponse, error) {
	var resp pb.WebsiteCategoryListResponse
	params := &website_category.WebsiteCategoryListInput{}
	data, err := website_category.ListNumWebsiteCategory(ctx, params)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	resp.Data = factory.NewWebsiteCategorySlice(data)
	return &resp, nil
}

func (s websitesvcService) WebsitePage(ctx context.Context, in *pb.WebsitePageRequest) (*pb.WebsitePageResponse, error) {
	var resp pb.WebsitePageResponse
	params := &search.WebsiteSearch{}
	if in.Paginator != nil {
		params.PageParams = &pagination.PageQuickParams{
			Limit:  int64(in.Paginator.Limit),
			NextID: in.Paginator.NextId,
		}
	}
	params.Type = enum.Web3
	if in.Type != "" {
		wtype, err := enum.WebsiteTypeFromString(in.Type)
		if err == nil {
			params.Type = wtype
		}
	}
	if in.WebsiteCategoryId != "" {
		website_category_id, err := primitive.ObjectIDFromHex(in.WebsiteCategoryId)
		if err == nil {
			params.WebsiteCategoryID = website_category_id
		}
	}
	params.Keywords = in.Keywords
	data, page, err := website.PageWebsite(ctx, &website.WebsiteInput{WebsiteSearch: params})
	if err != nil {
		return nil, err
	}
	quickpage := page.BuildJSONResult().(*pagination.QuickPagination)
	resp.Code = 0
	resp.Data = factory.NewWebsiteSlice(data)
	resp.Paginator = &pb.PageQuick{
		Limit:  uint64(quickpage.Limit),
		NextId: quickpage.NextID,
	}
	return &resp, nil
}

func (s websitesvcService) WebsiteRecommend(ctx context.Context, in *pb.WebsiteRecommendRequest) (*pb.WebsiteRecommendResponse, error) {
	var resp pb.WebsiteRecommendResponse
	params := &website.WebsiteRecommendInput{Num: uint(in.ListNum)}
	data, err := website.RecommendWebsite(ctx, params)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	resp.Data = factory.NewWebsiteSlice(data)
	return &resp, nil
}

func (s websitesvcService) WebsiteImport(ctx context.Context, in *pb.WebsiteImportRequest) (*pb.WebsiteImportResponse, error) {
	var resp pb.WebsiteImportResponse
	err := website.ImportWebsite(ctx, in.FilePath)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	return &resp, nil
}
