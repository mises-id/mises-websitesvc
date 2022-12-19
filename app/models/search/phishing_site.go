package search

import (
	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/lib/db/odm"
	"github.com/mises-id/mises-websitesvc/lib/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PhishingSiteSearch struct {
		ID         primitive.ObjectID
		Type       enum.PhishingSiteType
		DomainName string
		//sort
		SortBy     string `json:"sort_by" query:"sort_by"`
		CheckState enum.PhishingSiteCheckState
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

func (params *PhishingSiteSearch) BuildAdminSearch(chain *odm.DB) *odm.DB {
	//base
	//where
	if params.DomainName != "" {
		chain = chain.Where(bson.M{"domain_name": params.DomainName})
	}
	if params.Type > 0 {
		chain = chain.Where(bson.M{"type": params.Type})
	}
	if params.CheckState > 0 {
		chain = chain.Where(bson.M{"check_state": params.CheckState})
	}
	//sort
	chain = chain.Sort(bson.M{"_id": -1})

	//limit
	if (params.PageNum <= 0 || params.PageSize <= 0) && params.ListNum > 0 {
		chain = chain.Limit(params.ListNum)
	}
	return chain
}

func (params *PhishingSiteSearch) GetPageParams() *pagination.TraditionalParams {
	page := pagination.DefaultTraditionalParams()
	if params.PageNum > 0 {
		page.PageNum = params.PageNum
	}
	if params.PageSize > 0 {
		page.PageSize = params.PageSize
	}
	return page
}
func (params *PhishingSiteSearch) GetQuickPageParams() *pagination.PageQuickParams {
	res := pagination.DefaultQuickParams()
	if params.ListNum > 0 {
		res.Limit = params.Limit
	}
	if params.NextID != "" {
		res.NextID = params.NextID
	}
	return params.PageParams
}
