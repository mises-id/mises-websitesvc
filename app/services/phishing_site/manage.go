package phishing_site

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/mises-id/mises-websitesvc/app/models"
	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/app/models/search"
	"github.com/mises-id/mises-websitesvc/lib/utils"
)

type (
	MetaMaskPhishingConfigJson struct {
		Version   string   `json:"version" bson:"version"`
		Tolerance string   `json:"tolerance" bson:"tolerance"`
		Fuzzylist []string `json:"fuzzylist" bson:"fuzzylist"`
		Whitelist []string `json:"whitelist" bson:"whitelist"`
		Blacklist []string `json:"blacklist" bson:"blacklist"`
	}
)

//UpdateMetaMaskPhishingConfig
func UpdateMetaMaskPhishingConfig(ctx context.Context) error {
	fmt.Printf("[%s] UpdateMetaMaskPhishingConfig Start\n", time.Now().Local().String())
	err := runUpdateMetaMaskPhishingConfig(ctx)
	fmt.Printf("[%s] UpdateMetaMaskPhishingConfig End\n", time.Now().Local().String())
	return err
}

//runUpdateMetaMaskPhishingConfig
func runUpdateMetaMaskPhishingConfig(ctx context.Context) error {
	//GetMetaMaskPhishingConfigData
	phishing_data, err := GetMetaMaskPhishingConfigData()
	if err != nil {
		fmt.Printf("[%s] GetMetaMaskPhishingConfigData Error: %s\n", time.Now().Local().String(), err.Error())
		return err
	}
	//updatePhishingSiteByMetaMaskPhishingConfig
	return updatePhishingSiteByMetaMaskPhishingConfig(ctx, phishing_data)
}

//updatePhishingSiteByMetaMaskPhishingConfig
func updatePhishingSiteByMetaMaskPhishingConfig(ctx context.Context, data *MetaMaskPhishingConfigJson) error {
	if data == nil {
		return nil
	}
	//Whitelist
	if err := updatePhishingSiteByDomainType(ctx, enum.PhishingSiteWhite, data.Whitelist); err != nil {
		fmt.Printf("[%s] updatePhishingSiteByDomainType Whitelist Error: %s\n", time.Now().Local().String(), err.Error())
	}
	//Blacklist
	if err := updatePhishingSiteByDomainType(ctx, enum.PhishingSiteBlack, data.Blacklist); err != nil {
		fmt.Printf("[%s] updatePhishingSiteByDomainType Blacklist Error: %s\n", time.Now().Local().String(), err.Error())
	}
	//Fuzzylist
	if err := updatePhishingSiteByDomainType(ctx, enum.PhishingSiteFuzzy, data.Fuzzylist); err != nil {
		fmt.Printf("[%s] updatePhishingSiteByDomainType Fuzzylist Error: %s\n", time.Now().Local().String(), err.Error())
	}
	return nil
}

//updatePhishingSiteByDomainType
func updatePhishingSiteByDomainType(ctx context.Context, ptype enum.PhishingSiteType, lists []string) error {
	if ptype == enum.PhishingSiteNil {
		return errors.New("Invalid phishing type")
	}
	for _, v := range lists {
		add := &models.PhishingSite{
			Type:       ptype,
			DomainName: utils.UrlToDomainName(v),
		}
		if _, err := models.SavePhishingSite(ctx, add); err != nil {
			fmt.Printf("[%s] PhishingType[%s] DomainName[%s] SavePhishingSite Error: %s\n", time.Now().Local().String(), ptype.String(), v, err.Error())
			continue
		}
		fmt.Printf("[%s] DomainName[%s] SavePhishingSite Success\n", time.Now().Local().String(), v)
	}
	return nil
}

//GetMetaMaskPhishingConfigData
func GetMetaMaskPhishingConfigData() (*MetaMaskPhishingConfigJson, error) {

	//local json
	localfile := "./webh.json"
	jsonFile, err := os.Open(localfile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	out := &MetaMaskPhishingConfigJson{}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteValue, out)
	return out, nil
}

//UpdatePhishingSiteBlackOrigin
func UpdatePhishingSiteBlackOrigin(ctx context.Context) error {
	fmt.Printf("[%s] UpdatePhishingSiteBlackOrigin Start\n", time.Now().Local().String())
	err := runUpdatePhishingSiteBlackOrigin(ctx)
	fmt.Printf("[%s] UpdatePhishingSiteBlackOrigin End\n", time.Now().Local().String())
	return err
}

//runUpdatePhishingSiteBlackOrigin
func runUpdatePhishingSiteBlackOrigin(ctx context.Context) error {

	//black_list
	lists, err := models.ListPhishingSite(ctx, &search.PhishingSiteSearch{
		Type:       enum.PhishingSiteBlack,
		ListNum:    20,
		CheckState: enum.PhishingSiteCheckWait,
	})
	if err != nil {
		return err
	}
	num := len(lists)
	for _, v := range lists {
		fmt.Println(num)
		num--
		if v.Origin != "" {
			continue
		}
		domain_name := v.DomainName
		res, err := DomainNameToOriginCheck(ctx, domain_name)
		if err != nil {
			fmt.Printf("[%s] DomainName[%s] DomainNameToOriginCheck Error: %s\n", time.Now().Local().String(), domain_name, err.Error())
			continue
		}
		v.CheckState = enum.PhishingSiteCheckSuccess
		v.DomainNameCheckResult = res
		if err := models.UpdatePhishingSiteCheck(ctx, v); err != nil {
			fmt.Printf("[%s] DomainName[%s] DomainNameToOriginCheck UpdatePhishingSiteCheck Error: %s\n", time.Now().Local().String(), domain_name, err.Error())
			continue
		}
		if res.OriginNum == 0 {
			fmt.Printf("[%s] DomainNameToOriginCheck not find. DomainName[%s] \n", time.Now().Local().String(), domain_name)
			continue
		}
		fmt.Printf("[%s] DomainNameToOriginCheck Success. domain_name:%s,origin:%s,keywords:%s,lsc:%s,origin_percentage:%f，domain_percentage：%f \n", time.Now().Local().String(), res.DomainName, res.OriginMax.Origin, res.OriginMax.DomainKeyword, res.OriginMax.Lcs, res.OriginMax.OriginPercentage, res.OriginMax.DomainPercentage)
		con := fmt.Sprintf("domain_name：%s，origin: %s，keywords：%s，lsc：%s，origin_percentage：%f，domain_percentage：%f,edit_distance:%d\n", res.DomainName, res.OriginMax.Origin, res.OriginMax.DomainKeyword, res.OriginMax.Lcs, res.OriginMax.OriginPercentage, res.OriginMax.DomainPercentage, res.OriginMax.EditDistance)
		utils.WirteLogAppend("./phishing_check.txt", con)
	}
	return nil
}

//UpdatePhishingSiteByWebSite
func UpdatePhishingSiteByWebSite(ctx context.Context) error {
	fmt.Printf("[%s] UpdatePhishingSiteByWebSite Start\n", time.Now().Local().String())
	err := runpdatePhishingSiteByWebSite(ctx)
	fmt.Printf("[%s] UpdatePhishingSiteByWebSite End\n", time.Now().Local().String())
	return err
}

//runpdatePhishingSiteByWebSite
func runpdatePhishingSiteByWebSite(ctx context.Context) error {
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
		add := &models.PhishingSite{
			Type:       enum.PhishingSiteWhite,
			DomainName: domain_name,
		}
		if _, err := models.SavePhishingSite(ctx, add); err != nil {
			fmt.Printf("[%s] RunpdatePhishingSiteByWebSite DomainName[%s] SavePhishingSite Error: %s\n", time.Now().Local().String(), domain_name, err.Error())
			continue
		}
		fmt.Printf("[%s] SavePhishingSite Success. DomainName[%s]\n", time.Now().Local().String(), domain_name)
	}
	return nil
}
