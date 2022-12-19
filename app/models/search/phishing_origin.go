package search

import (
	"github.com/mises-id/mises-websitesvc/lib/db/odm"
	"github.com/mises-id/mises-websitesvc/lib/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PhishingOriginSearch struct {
		ID primitive.ObjectID
		//sort
		SortBy string `json:"sort_by" query:"sort_by"`
		//limit
		ListNum int64
		//page
		PageNum    int64  `json:"page_num" query:"page_num"`
		PageSize   int64  `json:"page_size" query:"page_size"`
		Limit      int64  `json:"limit" query:"limit"`
		NextID     string `json:"last_id" query:"last_id"`
		PageParams *pagination.PageQuickParams
	}
)

func (params *PhishingOriginSearch) BuildAdminSearch(chain *odm.DB) *odm.DB {
	//base
	//where
	//sort
	chain = chain.Sort(bson.M{"_id": -1})

	//limit
	if (params.PageNum <= 0 || params.PageSize <= 0) && params.ListNum > 0 {
		chain = chain.Limit(params.ListNum)
	}
	return chain
}

func (params *PhishingOriginSearch) GetPageParams() *pagination.TraditionalParams {
	page := pagination.DefaultTraditionalParams()
	if params.PageNum > 0 {
		page.PageNum = params.PageNum
	}
	if params.PageSize > 0 {
		page.PageSize = params.PageSize
	}
	return page
}
func (params *PhishingOriginSearch) GetQuickPageParams() *pagination.PageQuickParams {
	res := pagination.DefaultQuickParams()
	if params.ListNum > 0 {
		res.Limit = params.Limit
	}
	if params.NextID != "" {
		res.NextID = params.NextID
	}
	return params.PageParams
}
