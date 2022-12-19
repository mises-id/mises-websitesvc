package models

import (
	"context"
	"errors"
	"time"

	"github.com/mises-id/mises-websitesvc/lib/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	CheckMeta struct {
		Length    int            `json:"length" bson:"length"`
		SymbolNum int            `json:"symbol_num" bson:"symbol_num"`
		SymbolMap map[string]int `json:"symbol_map" bson:"symbol_map"`
	}
	PhishingOrigin struct {
		ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		Title         string             `json:"title" bson:"title"`
		DomainName    string             `json:"domain_name" bson:"domain_name"`
		DomainKeyword string             `json:"domain_keyword" bson:"domain_keyword"`
		CheckMeta     *CheckMeta         `json:"check_meta" bson:"check_meta"`
		CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	}
)

func SavePhishingOrigin(ctx context.Context, data *PhishingOrigin) (*PhishingOrigin, error) {
	if data == nil {
		return nil, errors.New("create data is nil")
	}
	if data.DomainName == "" {
		return nil, errors.New("domain_name is require")
	}
	data.CreatedAt = time.Now()
	opt := &options.FindOneAndUpdateOptions{}
	opt.SetUpsert(true)
	opt.SetReturnDocument(1)
	result := db.DB().Collection("phishingorigins").FindOneAndUpdate(ctx, bson.M{"domain_name": data.DomainName}, bson.D{{Key: "$set", Value: data}}, opt)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return data, result.Decode(data)
}

func ListPhishingOrigin(ctx context.Context, params IAdminParams) ([]*PhishingOrigin, error) {
	res := make([]*PhishingOrigin, 0)
	chain := params.BuildAdminSearch(db.ODM(ctx))
	err := chain.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
