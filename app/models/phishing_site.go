package models

import (
	"context"
	"errors"
	"time"

	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/lib/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	OriginResult struct {
		Origin              string  `json:"origin" bson:"origin"`
		DomainKeyword       string  `json:"domain_keyword" bson:"domain_keyword"`
		DomainKeywordLength int     `json:"domain_keyword_length" bson:"domain_keyword_length"`
		Lcs                 string  `json:"lcs" bson:"lcs"`
		LcsLength           int     `json:"lcs_length" bson:"lcs_length"`
		OriginPercentage    float32 `json:"origin_percentage" bson:"origin_percentage"`
		DomainPercentage    float32 `json:"domain_percentage" bson:"domain_percentage"`
		DomainLength        int     `json:"domain_length" bson:"domain_length"`
		EditDistance        int     `json:"edit_distance" bson:"edit_distance"`
	}
	DomainNameCheckResult struct {
		DomainName   string `json:"domain_name" bson:"domain_name"`
		DomainLength int
		OriginList   []*OriginResult `json:"origin_list" bson:"origin_list"`
		OriginMax    *OriginResult   `json:"origin_max" bson:"origin_max"`
		OriginNum    int             `json:"origin_num" bson:"origin_num"`
	}
	PhishingSite struct {
		ID                    primitive.ObjectID          `bson:"_id,omitempty" json:"id"`
		Type                  enum.PhishingSiteType       `json:"type" bson:"type"`
		DomainName            string                      `json:"domain_name" bson:"domain_name"`
		Origin                string                      `json:"origin" bson:"origin"`
		DomainNameCheckResult *DomainNameCheckResult      `json:"domain_name_check_result" bson:"domain_name_check_result"`
		CheckState            enum.PhishingSiteCheckState `json:"check_state" bson:"check_state"`
		AddSource             enum.PhishingSiteAddSource
		UpdatedAt             time.Time `bson:"updated_at" json:"updated_at"`
		CreatedAt             time.Time `bson:"created_at" json:"created_at"`
		TypeString            string    `bson:"-"`
	}
)

func SavePhishingSite(ctx context.Context, data *PhishingSite) (*PhishingSite, error) {
	if data == nil {
		return nil, errors.New("create data is nil")
	}
	if data.DomainName == "" {
		return nil, errors.New("domain_name is require")
	}
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	opt := &options.FindOneAndUpdateOptions{}
	opt.SetUpsert(true)
	opt.SetReturnDocument(1)
	result := db.DB().Collection("phishingsites").FindOneAndUpdate(ctx, bson.M{"domain_name": data.DomainName}, bson.D{{Key: "$set", Value: data}}, opt)
	if result.Err() != nil {
		return nil, result.Err()
	}
	if err := result.Decode(data); err != nil {
		return nil, err
	}
	return data, preloadPhishingSite(ctx, data)
}

func preloadPhishingSite(ctx context.Context, lists ...*PhishingSite) error {
	for _, v := range lists {
		v.TypeString = v.Type.String()
	}
	return nil
}

func ListPhishingSite(ctx context.Context, params IAdminParams) ([]*PhishingSite, error) {
	res := make([]*PhishingSite, 0)
	chain := params.BuildAdminSearch(db.ODM(ctx))
	err := chain.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdatePhishingSiteCheck(ctx context.Context, data *PhishingSite) error {

	update := bson.M{}
	update["check_state"] = data.CheckState
	update["domain_name_check_result"] = data.DomainNameCheckResult
	_, err := db.DB().Collection("phishingsites").UpdateByID(ctx, data.ID, bson.D{{Key: "$set", Value: update}})
	return err
}

func FindPhishingSiteByDomainName(ctx context.Context, domain_name string) (*PhishingSite, error) {
	res := &PhishingSite{}
	err := db.ODM(ctx).Last(&res, bson.M{"domain_name": domain_name}).Error
	if err != nil {
		return nil, err
	}
	return res, preloadPhishingSite(ctx, res)
}
