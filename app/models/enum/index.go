package enum

import "github.com/mises-id/mises-websitesvc/lib/codes"

type (
	WebsiteType uint8
	StatusType  uint8
)

const (
	WebsiteNil WebsiteType = 0
	Web3       WebsiteType = 1
	Extension  WebsiteType = 2
	//StatusType
	StatusOpen   StatusType = 1
	StatusClose  StatusType = 2
	StatusDelete StatusType = 3
)

var (
	websiteTypeMap = map[WebsiteType]string{
		WebsiteNil: "",
		Web3:       "web3",
		Extension:  "Extension",
	}
	websiteTypeStringMap = map[string]WebsiteType{}
)

func init() {
	for key, val := range websiteTypeMap {
		websiteTypeStringMap[val] = key
	}
}

func (tp WebsiteType) String() string {
	return websiteTypeMap[tp]
}

func WebsiteTypeFromString(tp string) (WebsiteType, error) {
	v, ok := websiteTypeStringMap[tp]
	if !ok {
		return WebsiteNil, codes.ErrInvalidArgument.Newf("invalid type: %s", tp)
	}
	return v, nil
}
