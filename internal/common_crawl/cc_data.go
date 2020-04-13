package common_crawl

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	CCData struct {
		URLKey       string    `json:"urlkey,string" validate:"required"`
		Timestamp    time.Time `json:"timestamp,string" validate:"required"`
		URL          url.URL   `json:"url,string" validate:"required"`
		Status       int       `json:"status,string" validate:"required"`
		Filename     string    `json:"filename,string" validate:"required"`
		Offset       int       `json:"offset,string,omitempty"`
		Digest       string    `json:"digest,string,omitempty"`
		Length       int       `json:"length,string,omitempty"`
		Mime         string    `json:"mime,string,omitempty"`
		MimeDetected string    `json:"mime-detected,string,omitempty"`
		CharSet      string    `json:"charset,string,omitempty"`
		Languages    string    `json:"languages,string,omitempty"`
	}


	CCDataI interface {
		GetURLKey() string
	}
)

const ccTimestamp = "20060102150405"

func (ccd *CCData) UnmarshalJSON(j []byte) error {
	var rawStrings map[string]string

	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		switch strings.ToLower(k) {
		case "urlkey":
			ccd.URLKey = v
		case "timestamp":
			t, err := time.Parse(ccTimestamp, v)
			if err != nil {
				return err
			}
			ccd.Timestamp = t
		case "url":
			u, err := url.Parse(v)
			if err != nil {
				return err
			}
			ccd.URL = *u
		case "status":
			ccd.Status, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		case "filename":
			ccd.Filename = v
		case "offset":
			ccd.Offset, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		case "digest":
			ccd.Digest = v
		case "length":
			ccd.Length, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		case "mime":
			ccd.Mime = v
		case "mime-detected":
			ccd.MimeDetected = v
		case "charset":
			ccd.CharSet = v
		case "languages":
			ccd.Languages = v
		}
	}
		return nil
}

func (ccd *CCData) GetURLKey() string {
	return ccd.URLKey
}

func (ccd *CCData) GetTimestamp() time.Time {
	return ccd.Timestamp
}

