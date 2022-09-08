package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Website struct {
		ID         primitive.ObjectID `bson:"_id,omitempty"`
		Type       int                `bson:"type"` //1 web3   2 extension
		CategoryID primitive.ObjectID
		Title      string //标题
		Desc       string //功能描述
		Url        string
		Logo       string
		Keywords   []string //搜索关键词
		IsRec      bool     //是否推荐，优先展示推荐站点
		SortNum    uint32   //排序号
		ClickNum   uint32   //点击数
		Status     int      //1 启用， 2 未启用 3 删除
		Remark     string
		UpdatedAt  time.Time `bson:"updated_at"`
		CreatedAt  time.Time `bson:"created_at"`
	}
)
