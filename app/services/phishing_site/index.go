package phishing_site

import (
	"context"
	"errors"
	"strings"

	"github.com/mises-id/mises-websitesvc/app/models"
	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/lib/codes"
	"github.com/mises-id/mises-websitesvc/lib/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PhishingCheckInput struct {
		DomainName string
	}
)

//check domain_name
func DomainNamePhishingCheck(ctx context.Context, in *PhishingCheckInput) (*models.PhishingSite, error) {
	if in == nil || in.DomainName == "" {
		return nil, codes.ErrInvalidArgument
	}
	domain_name := utils.UrlToDomainName(in.DomainName)
	//check domain name
	if err := isValidDomainName(domain_name); err != nil {
		return nil, codes.ErrInvalidArgument.New(err.Error())
	}
	//cache
	return phishingCheckByDomainName(ctx, domain_name)
}

func isValidDomainName(domain_name string) error {
	if !strings.Contains(domain_name, ".") {
		return errors.New("Invalid domain name")
	}
	return nil
}

func phishingCheckByDomainName(ctx context.Context, domain_name string) (*models.PhishingSite, error) {
	//find check record
	phishing_site, err := findCheckRecordByDomainName(ctx, domain_name)
	if err != nil {
		return nil, err
	}
	//find check record
	if phishing_site != nil {
		return phishing_site, nil
	}
	//check new domain name
	return checkNewDomainName(ctx, domain_name)
}

//findCheckRecordByDomainName
func findCheckRecordByDomainName(ctx context.Context, domain_name string) (*models.PhishingSite, error) {
	phishing_site, err := models.FindPhishingSiteByDomainName(ctx, domain_name)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	return phishing_site, nil
}

//checkNewDomainName
func checkNewDomainName(ctx context.Context, domain_name string) (*models.PhishingSite, error) {
	check_result, err := DomainNameToOriginCheck(ctx, domain_name)
	if err != nil {
		return nil, err
	}
	site_type := enum.PhishingSiteWhite
	var origin_domain_name string
	if check_result.OriginMax != nil {
		origin := check_result.OriginMax
		if origin.OriginPercentage > 80 && origin.DomainPercentage > 30 {
			site_type = enum.PhishingSiteFuzzy
			origin_domain_name = origin.Origin
		}
	}
	add := &models.PhishingSite{
		Type:                  site_type,
		DomainName:            domain_name,
		CheckState:            enum.PhishingSiteCheckSuccess,
		AddSource:             enum.PhishingSiteAddByCheck,
		DomainNameCheckResult: check_result,
		Origin:                origin_domain_name,
	}
	return models.SavePhishingSite(ctx, add)
}
