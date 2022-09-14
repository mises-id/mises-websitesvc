package website_category

import (
	"context"

	"github.com/mises-id/mises-websitesvc/app/models"
	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/app/models/search"
)

type (
	WebsiteCategoryListInput struct {
	}
)

func ListNumWebsiteCategory(ctx context.Context, in *WebsiteCategoryListInput) ([]*models.WebsiteCategory, error) {

	params := &search.WebsiteCategorySearch{
		Type: enum.Web3,
	}
	list, err := models.ListWebsiteCategory(ctx, params)
	if err != nil {
		return nil, err
	}
	return list, nil
}
