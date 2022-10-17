package factory

import (
	"github.com/mises-id/mises-websitesvc/app/models"
	pb "github.com/mises-id/mises-websitesvc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func docID(id primitive.ObjectID) string {
	if id.IsZero() {
		return ""
	}
	return id.Hex()
}
func NewWebsiteCategorySlice(data []*models.WebsiteCategory) []*pb.WebsiteCategory {
	if data == nil {
		return nil
	}
	result := make([]*pb.WebsiteCategory, len(data))
	for i, v := range data {
		result[i] = NewWebsiteCategory(v)
	}
	return result
}
func NewWebsiteCategory(data *models.WebsiteCategory) *pb.WebsiteCategory {
	if data == nil {
		return nil
	}
	resp := pb.WebsiteCategory{
		Id:          docID(data.ID),
		ParentId:    docID(data.ParentID),
		Name:        data.Name,
		ShorterName: data.ShorterName,
		Desc:        data.Desc,
		TypeString:  data.Type.String(),
	}
	resp.ChildrenCategory = NewWebsiteCategorySlice(data.ChildrenCategory)
	return &resp
}
func NewWebsiteSlice(data []*models.Website) []*pb.Website {
	result := make([]*pb.Website, len(data))
	for i, v := range data {
		result[i] = NewWebsite(v)
	}
	return result
}
func NewWebsite(data *models.Website) *pb.Website {
	if data == nil {
		return nil
	}
	resp := pb.Website{
		Id:                docID(data.ID),
		WebsiteCategoryId: docID(data.CategoryID),
		SubcategoryId:     docID(data.SubcategoryID),
		Title:             data.Title,
		Logo:              data.Logo,
		Url:               data.Url,
		Desc:              data.Desc,
	}
	if data.WebsiteCategory != nil {
		resp.WebsiteCategory = NewWebsiteCategory(data.WebsiteCategory)
	}
	if data.Subcategory != nil {
		resp.Subcategory = NewWebsiteCategory(data.Subcategory)
	}
	return &resp
}
