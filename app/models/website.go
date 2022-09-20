package models

import (
	"context"
	"time"

	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/lib/db"
	"github.com/mises-id/mises-websitesvc/lib/pagination"
	"github.com/mises-id/mises-websitesvc/lib/storage/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Website struct {
		ID              primitive.ObjectID `bson:"_id,omitempty"`
		Type            enum.WebsiteType   `bson:"type"` //1 web3   2 extension
		CategoryID      primitive.ObjectID `bson:"category_id"`
		Title           string             `bson:"title"` //标题
		Desc            string             `bson:"desc"`  //功能描述
		Url             string             `bson:"url"`
		Logo            string             `bson:"logo"`
		Keywords        []string           `bson:"keywords"` //搜索关键词
		IsRec           bool               `bson:"is_rec"`   //是否推荐，优先展示推荐站点
		IsHot           bool               `bson:"is_hot"`
		SortNum         uint32             `bson:"sort_num"`  //排序号
		ClickNum        uint32             `bson:"click_num"` //点击数
		Status          enum.StatusType    `bson:"status"`    //1 启用， 2 未启用 3 删除
		Remark          string             `bson:"remark"`
		UpdatedAt       time.Time          `bson:"updated_at"`
		CreatedAt       time.Time          `bson:"created_at"`
		WebsiteCategory *WebsiteCategory   `bson:"-" json:"website_category"`
	}
)

func QuickPageWebsite(ctx context.Context, params IAdminQuickPageParams) ([]*Website, pagination.Pagination, error) {
	out := make([]*Website, 0)
	chain := params.BuildAdminSearch(db.ODM(ctx))
	pageParams := params.GetQuickPageParams()
	paginator := pagination.NewQuickPaginator(pageParams.Limit, pageParams.NextID, chain)
	page, err := paginator.Paginate(&out)
	if err != nil {
		return nil, nil, err
	}
	return out, page, preloadWebsite(ctx, out...)
}
func PageWebsite(ctx context.Context, params IAdminPageParams) ([]*Website, pagination.Pagination, error) {
	out := make([]*Website, 0)
	chain := params.BuildAdminSearch(db.ODM(ctx))
	pageParams := params.GetPageParams()
	paginator := pagination.NewTraditionalPaginator(pageParams.PageNum, pageParams.PageSize, chain)
	page, err := paginator.Paginate(&out)
	if err != nil {
		return nil, nil, err
	}
	return out, page, preloadWebsite(ctx, out...)
}
func preloadWebsite(ctx context.Context, lists ...*Website) error {
	cagegoryIDs := make([]primitive.ObjectID, 0)
	for _, v := range lists {
		cagegoryIDs = append(cagegoryIDs, v.CategoryID)
	}
	categoryList, err := FindWebsiteCategoryByIDs(ctx, cagegoryIDs...)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	dataMap := make(map[primitive.ObjectID]*WebsiteCategory)
	for _, v := range categoryList {
		dataMap[v.ID] = v
	}
	for _, v := range lists {
		v.WebsiteCategory = dataMap[v.CategoryID]
		vlogo, err := view.ImageClient.GetFileUrlOne(ctx, v.Logo)
		if err == nil {
			v.Logo = vlogo
		}
	}
	return nil
}
func CreateWebsiteMany(ctx context.Context, data []*Website) error {
	if len(data) == 0 {
		return nil
	}
	var in []interface{}
	for _, v := range data {
		in = append(in, v)
	}
	_, err := db.DB().Collection("websites").InsertMany(ctx, in)

	return err
}

func ListWebsite(ctx context.Context, params IAdminParams) ([]*Website, error) {
	res := make([]*Website, 0)
	chain := params.BuildAdminSearch(db.ODM(ctx))
	err := chain.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
