package utils

import (
	"fmt"
	"io"
	"os"
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

func UrlToDomainName(url string) string {
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
	hostName := UrlToDomainName(url)
	if hostName != "" {
		logo = fmt.Sprintf("%s%s", KiwiLogoApi, hostName)
	}
	return logo
}

func WirteLogAppend(path string, content string) error {

	arr := strings.Split(path, "/")
	filePath := strings.Join(arr[:len(arr)-1], "/")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	fileObj, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer fileObj.Close()
	if _, err := io.WriteString(fileObj, content); err == nil {
		return err
	}
	return nil
}

func DomainNameToKeyword(dn string) string {
	if dn == "" {
		return dn
	}
	arr := strings.Split(dn, ".")
	length := len(arr)
	if length == 1 {
		return arr[0]
	}
	k1 := arr[length-2]
	k2 := arr[length-1]
	if len(k1) <= 5 {
		k1 = fmt.Sprintf("%s%s", k1, k2)
	}
	return k1
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
