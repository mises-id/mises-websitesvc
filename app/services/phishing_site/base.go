package phishing_site

import (
	"context"
	"strings"

	"github.com/mises-id/mises-websitesvc/app/models"
	"github.com/mises-id/mises-websitesvc/app/models/search"
	"github.com/mises-id/mises-websitesvc/lib/utils"
)

type ()

//DomainNameToOriginCheck
func DomainNameToOriginCheck(ctx context.Context, domain_name string) (*models.DomainNameCheckResult, error) {
	//origin_list
	lists, err := models.ListPhishingOrigin(ctx, &search.PhishingOriginSearch{})
	if err != nil {
		return nil, err
	}
	domain_length := len(domain_name)
	var origin_num, maxLen int
	var maxPercentage float32
	var origin_list []*models.OriginResult
	var origin_max *models.OriginResult

	for _, v := range lists {
		origin_keyword := utils.DomainNameToKeyword(v.DomainName)
		origin_keyword_length := len(origin_keyword)
		lcs := LCS(domain_name, origin_keyword)
		lcs_length := len(lcs)
		percentage := float32(lcs_length) / float32(origin_keyword_length) * float32(100)
		domain_percentage := float32(lcs_length) / float32(domain_length) * float32(100)
		edit_distance := editDistance(domain_name, origin_keyword)
		if percentage < 50 {
			continue
		}
		origin := &models.OriginResult{
			Origin:              v.DomainName,
			DomainKeyword:       origin_keyword,
			Lcs:                 lcs,
			LcsLength:           lcs_length,
			DomainKeywordLength: origin_keyword_length,
			OriginPercentage:    percentage,
			DomainLength:        domain_length,
			DomainPercentage:    domain_percentage,
			EditDistance:        edit_distance,
		}
		if origin_list == nil || len(origin_list) < 5 {
			origin_list = append(origin_list, origin)
		} else {
			var minIndex int
			for i, v := range origin_list {
				if i == 0 {
					continue
				}
				if v.OriginPercentage < origin_list[minIndex].OriginPercentage {
					minIndex = i
				}
			}
			origin_list[minIndex] = origin
		}
		if percentage > maxPercentage || (percentage == maxPercentage && lcs_length > maxLen) {
			maxPercentage = percentage
			maxLen = lcs_length
			origin_max = origin
		}
		origin_num++
	}
	out := &models.DomainNameCheckResult{
		DomainName:   domain_name,
		OriginList:   origin_list,
		OriginMax:    origin_max,
		OriginNum:    origin_num,
		DomainLength: domain_length,
	}
	return out, nil
}

//lcs
func LCS(str1 string, str2 string) string {
	l := 0
	r := 1
	var exist bool
	var maxLen int
	var maxL int
	var maxR int
	for ; r < len(str1)+1; r++ {
		exist = strings.Contains(str2, str1[l:r])
		if exist == false {
			//l = r - 1
			l++
			continue
		}
		if r-l >= maxLen {
			maxR = r
			maxL = l
			maxLen = r - l
		}
	}
	return str1[maxL:maxR]
}

//lcs
/* func LCS(s1 string, s2 string) string {
	// write code here
	if len(s2) == 0 {
		return "-1"
	}
	dp := make([]string, len(s2))
	var t, t1 string
	for i := 0; i < len(s1); i++ {
		t1 = ""
		for j := 0; j < len(s2); j++ {
			t = dp[j]
			if s2[j] == s1[i] {
				dp[j] = t1 + string(s1[i])
			} else {
				if j > 0 && len(dp[j]) < len(dp[j-1]) {
					dp[j] = dp[j-1]
				}
			}
			t1 = t
		}
	}
	if dp[len(s2)-1] == "" {
		return "-1"
	}
	return dp[len(s2)-1]
} */

//编辑距离
func editDistance(a, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = utils.IntMin(dp[i-1][j-1], utils.IntMin(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}
