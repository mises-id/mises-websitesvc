package models

import (
	"context"
	"time"

	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/lib/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	WebsiteCategory struct {
		ID               primitive.ObjectID `bson:"_id,omitempty"`
		ParentID         primitive.ObjectID `bson:"parent_id"`
		Type             enum.WebsiteType   `bson:"type"`
		Name             string             `bson:"name"`
		ShorterName      string             `bson:"shorter_name"`
		Desc             string             `bson:"desc"`
		SortNum          uint32             `bson:"sort_num"`
		UpdatedAt        time.Time          `bson:"updated_at"`
		CreatedAt        time.Time          `bson:"created_at"`
		ChildrenCategory []*WebsiteCategory `bson:"-"`
	}
)

func (u *WebsiteCategory) BeforeCreate(ctx context.Context) error {
	u.CreatedAt = time.Now()
	return u.BeforeUpdate(ctx)
}

func (u *WebsiteCategory) BeforeUpdate(ctx context.Context) error {
	u.UpdatedAt = time.Now()
	return nil
}

func CreateWebsiteCategory(ctx context.Context, data *WebsiteCategory) (*WebsiteCategory, error) {

	if err := data.BeforeCreate(ctx); err != nil {
		return nil, err
	}
	res, err := db.DB().Collection("websitecategories").InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	data.ID = res.InsertedID.(primitive.ObjectID)
	return data, err
}

func ListWebsiteCategory(ctx context.Context, params IAdminParams) ([]*WebsiteCategory, error) {
	res := make([]*WebsiteCategory, 0)
	chain := params.BuildAdminSearch(db.ODM(ctx))
	err := chain.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, preloadWebsiteCategory(ctx, res...)
}
func FindWebsiteCategoryByIDs(ctx context.Context, ids ...primitive.ObjectID) ([]*WebsiteCategory, error) {
	res := make([]*WebsiteCategory, 0)
	err := db.ODM(ctx).Find(&res, bson.M{"_id": bson.M{"$in": ids}}).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindWebsiteCategoryByParentID(ctx context.Context, parent_id primitive.ObjectID) ([]*WebsiteCategory, error) {
	res := make([]*WebsiteCategory, 0)
	err := db.ODM(ctx).Find(&res, bson.M{"parent_id": parent_id}).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func preloadWebsiteCategory(ctx context.Context, lists ...*WebsiteCategory) error {
	for _, v := range lists {
		if v.ParentID.IsZero() {
			children_category, err := FindWebsiteCategoryByParentID(ctx, v.ID)
			if err == nil {
				v.ChildrenCategory = children_category
			}
		}
	}
	return nil
}
