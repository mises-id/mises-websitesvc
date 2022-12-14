package search

import (
	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/lib/db/odm"
	"github.com/mises-id/mises-websitesvc/lib/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	WebsiteSearch struct {
		ID                primitive.ObjectID
		Type              enum.WebsiteType
		WebsiteCategoryID primitive.ObjectID
		SubcategoryID     primitive.ObjectID
		Statuses          []enum.StatusType
		Keywords          string
		RecState          int
		HotState          int
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

func (params *WebsiteSearch) BuildAdminSearch(chain *odm.DB) *odm.DB {
	//base
	//where

	if params.Type > 0 {
		chain = chain.Where(bson.M{"type": params.Type})
	}
	if params.WebsiteCategoryID != primitive.NilObjectID {
		chain = chain.Where(bson.M{"category_id": params.WebsiteCategoryID})
	}
	if params.SubcategoryID != primitive.NilObjectID {
		chain = chain.Where(bson.M{"subcategory_id": params.SubcategoryID})
	}
	if len(params.Statuses) > 0 {
		chain = chain.Where(bson.M{"status": bson.M{"$in": params.Statuses}})
	} else {
		chain = chain.Where(bson.M{"status": enum.StatusOpen})
	}
	if params.HotState == 1 {
		chain = chain.Where(bson.M{"is_hot": true})
	}
	if params.HotState == 2 {
		chain = chain.Where(bson.M{"is_hot": false})
	}
	if params.RecState == 1 {
		chain = chain.Where(bson.M{"is_rec": true})
	}
	if params.RecState == 2 {
		chain = chain.Where(bson.M{"is_rec": false})
	}
	if params.Type == enum.Extensions && params.SortBy == "" {
		params.SortBy = "category"
	}
	//sort
	switch params.SortBy {
	default:
		chain = chain.Sort(bson.D{
			bson.E{"is_hot", -1},
			bson.E{"sort_num", -1},
			bson.E{"title", 1},
		})
	case "category":
		chain = chain.Sort(bson.D{
			bson.E{"category_sort", -1},
			bson.E{"category_id", 1},
			bson.E{"subcategory_sort", -1},
			bson.E{"subcategory_id", 1},
			bson.E{"is_hot", -1},
			bson.E{"sort_num", -1},
			bson.E{"title", 1},
		})
	case "rec":
		chain = chain.Sort(bson.D{
			bson.E{"is_rec", -1},
			bson.E{"rec_sort_num", -1},
			bson.E{"title", 1},
		})
	case "id_desc":
		chain = chain.Sort(bson.M{"_id": -1})
	}

	//limit
	if (params.PageNum <= 0 || params.PageSize <= 0) && params.ListNum > 0 {
		chain = chain.Limit(params.ListNum)
	}
	return chain
}

func (params *WebsiteSearch) GetPageParams() *pagination.TraditionalParams {
	page := pagination.DefaultTraditionalParams()
	if params.PageNum > 0 {
		page.PageNum = params.PageNum
	}
	if params.PageSize > 0 {
		page.PageSize = params.PageSize
	}
	return page
}
func (params *WebsiteSearch) GetQuickPageParams() *pagination.PageQuickParams {
	res := pagination.DefaultQuickParams()
	if params.ListNum > 0 {
		res.Limit = params.Limit
	}
	if params.NextID != "" {
		res.NextID = params.NextID
	}
	return params.PageParams
}
