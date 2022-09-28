package website_category

import (
	"context"

	"github.com/mises-id/mises-websitesvc/app/models"
	"github.com/mises-id/mises-websitesvc/app/models/search"
)

type (
	WebsiteCategoryListInput struct {
		*search.WebsiteCategorySearch
	}
)

func ListNumWebsiteCategory(ctx context.Context, in *WebsiteCategoryListInput) ([]*models.WebsiteCategory, error) {

	params := in.WebsiteCategorySearch
	list, err := models.ListWebsiteCategory(ctx, params)
	if err != nil {
		return nil, err
	}
	return list, nil
}
