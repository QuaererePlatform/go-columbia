package index

import (
	"encoding/json"
	"net/url"
	"strings"
)

type (
	CollInfo struct {
		ID string `json:"id,string" validate:"required"`
		Name string `json:"name,string" validate:"required"`
		Timegate url.URL `json:"timegate,string" validate:"required"`
		CDXAPI url.URL `json:"cdx-api,string" validate:"required"`
	}
)

const collInfoURL = "https://index.commoncrawl.org/collinfo.json"

func (c *CollInfo) UnmarshalJSON(j []byte) error {
	var rawStrings map[string]string

	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		switch strings.ToLower(k) {
		case "id":
			c.ID = v
		case "name":
			c.Name = v
		case "timegate":
			u, err := url.Parse(v)
			if err != nil {
				return err
			}
			c.Timegate = *u
		case "cdx-api":
			u, err := url.Parse(v)
			if err != nil {
				return err
			}
			c.CDXAPI = *u
		}
	}
	return nil
}
