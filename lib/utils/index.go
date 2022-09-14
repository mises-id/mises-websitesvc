package utils

import (
	"fmt"
	"strings"
)

const (
	KiwiLogoApi = "https://logos.kiwibrowser.com/"
)

func FindStringArrayValueIndex(header []string, value string) int {
	index := -1
	for k, v := range header {
		if v == value {
			index = k
			break
		}
	}
	return index
}

func GetHostNameByUrl(url string) string {
	if url == "" {
		return url
	}
	host := ""
	arr := strings.Split(url, "?")
	host = strings.TrimPrefix(arr[0], "https://")
	host = strings.TrimPrefix(host, "http://")
	host = strings.Split(host, "/")[0]
	return host
}

func GetUrlLogoByKiwi(url string) (logo string) {
	hostName := GetHostNameByUrl(url)
	if hostName != "" {
		logo = fmt.Sprintf("%s%s", KiwiLogoApi, hostName)
	}
	return logo
}
