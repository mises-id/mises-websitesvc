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
	params := &search.WebsiteCategorySearch{}
	params.Type = enum.Web3
	if in.Type != "" {
		wtype, err := enum.WebsiteTypeFromString(in.Type)
		if err != nil {
			return nil, err
		}
		params.Type = wtype
	}
	data, err := website_category.ListNumWebsiteCategory(ctx, &website_category.WebsiteCategoryListInput{WebsiteCategorySearch: params})
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
		params.PageNum = int64(in.Paginator.PageNum)
		params.PageSize = int64(in.Paginator.PageSize)
	}
	params.Type = enum.Web3
	if in.Type != "" {
		wtype, err := enum.WebsiteTypeFromString(in.Type)
		if err != nil {
			return nil, err
		}
		params.Type = wtype
	}
	if in.WebsiteCategoryId != "" {
		website_category_id, err := primitive.ObjectIDFromHex(in.WebsiteCategoryId)
		if err == nil {
			params.WebsiteCategoryID = website_category_id
		}
	}
	if in.SubcategoryId != "" {
		subcategory_id, err := primitive.ObjectIDFromHex(in.SubcategoryId)
		if err == nil {
			params.SubcategoryID = subcategory_id
		}
	}
	params.Keywords = in.Keywords
	data, page, err := website.PageWebsite(ctx, &website.WebsiteInput{WebsiteSearch: params})
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	resp.Data = factory.NewWebsiteSlice(data)
	tradpage := page.BuildJSONResult().(*pagination.TraditionalPagination)
	resp.Paginator = &pb.Page{
		PageNum:      uint64(tradpage.PageNum),
		PageSize:     uint64(tradpage.PageSize),
		TotalPage:    uint64(tradpage.TotalPages),
		TotalRecords: uint64(tradpage.TotalRecords),
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
