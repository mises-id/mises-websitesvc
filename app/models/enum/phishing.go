package enum

import "github.com/mises-id/mises-websitesvc/lib/codes"

type (
	PhishingSiteType       uint8
	PhishingSiteCheckState uint8
	PhishingSiteAddSource  uint8
)

const (
	//PhishingSiteType
	PhishingSiteNil   PhishingSiteType = 0
	PhishingSiteWhite PhishingSiteType = 1
	PhishingSiteFuzzy PhishingSiteType = 2
	PhishingSiteBlack PhishingSiteType = 3
	//PhishingSiteCheckState
	PhishingSiteCheckDefault PhishingSiteCheckState = 0
	PhishingSiteCheckWait    PhishingSiteCheckState = 1
	PhishingSiteCheckSuccess PhishingSiteCheckState = 2
	//PhishingSiteAddSource
	PhishingSiteAddByDefault PhishingSiteAddSource = 0
	PhishingSiteAddByCheck   PhishingSiteAddSource = 1
)

var (
	phishingSiteTypeMap = map[PhishingSiteType]string{
		PhishingSiteNil:   "",
		PhishingSiteWhite: "white",
		PhishingSiteFuzzy: "fuzzy",
		PhishingSiteBlack: "black",
	}
	phishingSiteTypeStringMap = map[string]PhishingSiteType{}
)

func init() {
	for key, val := range phishingSiteTypeMap {
		phishingSiteTypeStringMap[val] = key
	}
}

func (tp PhishingSiteType) String() string {
	return phishingSiteTypeMap[tp]
}

func PhishingSiteTypeFromString(tp string) (PhishingSiteType, error) {
	v, ok := phishingSiteTypeStringMap[tp]
	if !ok {
		return PhishingSiteNil, codes.ErrInvalidArgument.Newf("invalid type: %s", tp)
	}
	return v, nil
}
