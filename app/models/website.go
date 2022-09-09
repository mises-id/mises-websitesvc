package models

import (
	"context"
	"time"

	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/lib/db"
	"github.com/mises-id/mises-websitesvc/lib/pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Website struct {
		ID         primitive.ObjectID `bson:"_id,omitempty"`
		Type       enum.WebsiteType   `bson:"type"` //1 web3   2 extension
		CategoryID primitive.ObjectID
		Title      string //标题
		Desc       string //功能描述
		Url        string
		Logo       string
		Keywords   []string        //搜索关键词
		IsRec      bool            //是否推荐，优先展示推荐站点
		SortNum    uint32          //排序号
		ClickNum   uint32          //点击数
		Status     enum.StatusType //1 启用， 2 未启用 3 删除
		Remark     string
		UpdatedAt  time.Time `bson:"updated_at"`
		CreatedAt  time.Time `bson:"created_at"`
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
	return out, page, nil
}
