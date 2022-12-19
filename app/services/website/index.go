package website

import (
	"context"
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/mises-id/mises-websitesvc/app/models"
	"github.com/mises-id/mises-websitesvc/app/models/enum"
	"github.com/mises-id/mises-websitesvc/app/models/search"
	"github.com/mises-id/mises-websitesvc/lib/codes"
	"github.com/mises-id/mises-websitesvc/lib/pagination"
	"github.com/mises-id/mises-websitesvc/lib/utils"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	WebsiteInput struct {
		*search.WebsiteSearch
	}
	WebsiteRecommendInput struct {
		Num uint
	}
	WebsiteSearchInput struct {
		Keywords string
	}
)

//PageWebsite
func PageWebsite(ctx context.Context, in *WebsiteInput) ([]*models.Website, pagination.Pagination, error) {
	params := in.WebsiteSearch
	return models.PageWebsite(ctx, params)
}

//SearchWebsite
func SearchWebsite(ctx context.Context, in *WebsiteSearchInput) ([]*models.Website, error) {
	if in == nil {
		return nil, codes.ErrInvalidArgument
	}
	return models.SearchWebsite(ctx, in.Keywords)
}

//RecommendWebsite
func RecommendWebsite(ctx context.Context, in *WebsiteRecommendInput) ([]*models.Website, error) {
	var list_num uint
	params := &search.WebsiteSearch{
		SortBy: "rec",
	}
	list_num = 8
	if in.Num > 0 && in.Num <= 200 {
		list_num = in.Num
	}
	params.ListNum = int64(list_num)
	params.RecState = 1
	data, err := models.ListWebsite(ctx, params)
	if err != nil {
		return nil, err
	}
	//json
	return data, nil
}

//ImportWebsite
func ImportWebsite(ctx context.Context, filePath string) error {
	ext := path.Ext(filePath)
	if ext != ".xlsx" {
		return errors.New("文件格式必须为xlsx")
	}
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("sheet1")
	if err != nil {
		return err
	}
	site_list := make([]*models.Website, 0)
	if len(rows) < 1 {
		return nil
	}
	header := rows[0]
	//处理header
	titleIndex := utils.FindStringArrayValueIndex(header, "title")
	urlIndex := utils.FindStringArrayValueIndex(header, "url")
	cateIndex := utils.FindStringArrayValueIndex(header, "category")
	descIndex := utils.FindStringArrayValueIndex(header, "desc")
	typeIndex := utils.FindStringArrayValueIndex(header, "type")
	logoIndex := utils.FindStringArrayValueIndex(header, "logo")
	remarkIndex := utils.FindStringArrayValueIndex(header, "remark")
	if titleIndex == -1 || urlIndex == -1 {
		return errors.New("文件表头格式错误！")
	}
	//处理分类
	params := &search.WebsiteCategorySearch{}
	categoryMap := make(map[string]primitive.ObjectID, 0)
	WebsiteCategoryList, err := models.ListWebsiteCategory(ctx, params)
	if err == nil {
		for _, v := range WebsiteCategoryList {
			typeName := fmt.Sprintf("%s-%s", v.Type.String(), v.Name)
			categoryMap[typeName] = v.ID
		}
	}
	nowTime := time.Now()
	sortNum := 50000
	sortSetp := 100
	sortMin := 100
	for i, row := range rows {
		maxIndex := len(row) - 1
		if i == 0 || titleIndex > maxIndex || urlIndex > maxIndex {
			continue
		}
		title := row[titleIndex]
		url := row[urlIndex]
		if title == "" {
			continue
		}
		if url == "" {
			continue
		}
		site := &models.Website{}
		if descIndex > -1 && descIndex <= maxIndex {
			site.Desc = row[descIndex]
		}
		if remarkIndex > -1 && remarkIndex < maxIndex {
			site.Remark = row[remarkIndex]
		}
		logo := ""
		if logoIndex > -1 && logoIndex <= maxIndex {
			logo = row[logoIndex]
		}
		wtype := enum.Web3
		if typeIndex > -1 && typeIndex <= maxIndex {
			stype, err := enum.WebsiteTypeFromString(row[typeIndex])
			if err == nil {
				wtype = stype
			}
		}
		if logo == "" {
			logo = utils.GetUrlLogoByKiwi(url)
		}
		if cateIndex > -1 && cateIndex <= maxIndex {
			cateMapKey := fmt.Sprintf("%s-%s", wtype.String(), row[cateIndex])
			cateID, ok := categoryMap[cateMapKey]
			if ok {
				site.CategoryID = cateID
			}
		}
		site.Type = wtype
		site.Title = title
		site.Logo = logo
		site.Url = url
		site.SortNum = uint32(sortNum)
		if sortNum > sortMin {
			sortNum -= sortSetp
		}
		site.Status = enum.StatusOpen
		site.UpdatedAt = nowTime
		site.CreatedAt = nowTime
		site_list = append(site_list, site)
	}
	if err = models.CreateWebsiteMany(ctx, site_list); err != nil {
		return err
	}
	fmt.Printf("成功导入%d \n", len(site_list))
	return nil
}
