package phishing_origin

import (
	"context"
	"fmt"
	"time"

	"github.com/mises-id/mises-websitesvc/app/models"
	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/app/models/search"
	"github.com/mises-id/mises-websitesvc/lib/utils"
)

type ()

//UpdatePhishingOriginByWebSite
func UpdatePhishingOriginByWebSite(ctx context.Context) error {
	fmt.Printf("[%s] UpdatePhishingOriginByWebSite Start\n", time.Now().Local().String())
	err := runUpdatePhishingOriginByWebSite(ctx)
	fmt.Printf("[%s] UpdatePhishingOriginByWebSite End\n", time.Now().Local().String())
	return err
}

//runUpdatePhishingOriginByWebSite
func runUpdatePhishingOriginByWebSite(ctx context.Context) error {
	//web3site list
	params := &search.WebsiteSearch{
		Type: enum.Web3,
	}
	lists, err := models.ListWebsite(ctx, params)
	if err != nil {
		return err
	}
	for _, v := range lists {
		domain_name := utils.UrlToDomainName(v.Url)
		domain_keyword := utils.DomainNameToKeyword(domain_name)
		add := &models.PhishingOrigin{
			DomainName:    domain_name,
			Title:         v.Title,
			DomainKeyword: domain_keyword,
		}
		if _, err := models.SavePhishingOrigin(ctx, add); err != nil {
			fmt.Printf("[%s] DomainName[%s] SavePhishingOrigin Error: %s\n", time.Now().Local().String(), domain_name, err.Error())
			continue
		}
		fmt.Printf("[%s] DomainName[%s] SavePhishingOrigin Success\n", time.Now().Local().String(), domain_name)
	}
	return nil
}
