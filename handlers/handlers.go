package handlers

import (
	"context"

	"github.com/mises-id/mises-websitesvc/app/factory"
	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/app/models/search"
	"github.com/mises-id/mises-websitesvc/app/services/phishing_origin"
	"github.com/mises-id/mises-websitesvc/app/services/phishing_site"
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

func (s websitesvcService) UpdateMetaMaskPhishing(ctx context.Context, in *pb.UpdateMetaMaskPhishingRequest) (*pb.UpdateMetaMaskPhishingResponse, error) {
	var resp pb.UpdateMetaMaskPhishingResponse
	err := phishing_site.UpdateMetaMaskPhishingConfig(ctx)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	return &resp, nil
}

func (s websitesvcService) UpdatePhishingOriginByWebSite(ctx context.Context, in *pb.UpdatePhishingOriginByWebSiteRequest) (*pb.UpdatePhishingOriginByWebSiteResponse, error) {
	var resp pb.UpdatePhishingOriginByWebSiteResponse
	err := phishing_origin.UpdatePhishingOriginByWebSite(ctx)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	return &resp, nil
}

func (s websitesvcService) UpdatePhishingSiteBlackOrigin(ctx context.Context, in *pb.UpdatePhishingSiteBlackOriginRequest) (*pb.UpdatePhishingSiteBlackOriginResponse, error) {
	var resp pb.UpdatePhishingSiteBlackOriginResponse
	err := phishing_site.UpdatePhishingSiteBlackOrigin(ctx)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	return &resp, nil
}

func (s websitesvcService) UpdatePhishingSiteByWebsite(ctx context.Context, in *pb.UpdatePhishingSiteByWebsiteRequest) (*pb.UpdatePhishingSiteByWebsiteResponse, error) {
	var resp pb.UpdatePhishingSiteByWebsiteResponse
	err := phishing_site.UpdatePhishingSiteByWebSite(ctx)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	return &resp, nil
}

func (s websitesvcService) PhishingCheck(ctx context.Context, in *pb.PhishingCheckRequest) (*pb.PhishingCheckResponse, error) {
	var resp pb.PhishingCheckResponse
	params := &phishing_site.PhishingCheckInput{
		DomainName: in.DomainName,
	}
	res, err := phishing_site.DomainNamePhishingCheck(ctx, params)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	resp.Type = uint64(res.Type)
	resp.DomainName = res.DomainName
	resp.Origin = res.Origin
	resp.TypeString = res.TypeString
	return &resp, nil
}

func (s websitesvcService) WebsiteSearch(ctx context.Context, in *pb.WebsiteSearchRequest) (*pb.WebsiteSearchResponse, error) {
	var resp pb.WebsiteSearchResponse
	params := &website.WebsiteSearchInput{Keywords: in.Keywords}
	data, err := website.SearchWebsite(ctx, params)
	if err != nil {
		return nil, err
	}
	resp.Code = 0
	resp.Data = factory.NewWebsiteSlice(data)
	return &resp, nil
}
