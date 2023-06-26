package playwright

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type localUtilsImpl struct {
	channelOwner
}

type (
	localUtilsZipOptions struct {
		ZipFile string        `json:"zipFile"`
		Entries []interface{} `json:"entries"`
	}

	harLookupOptions struct {
		HarId               string            `json:"harId"`
		URL                 string            `json:"url"`
		Method              string            `json:"method"`
		Headers             map[string]string `json:"headers"`
		IsNavigationRequest bool              `json:"isNavigationRequest"`
		PostData            interface{}       `json:"postData,omitempty"`
	}

	harLookupResult struct {
		Action      string              `json:"action"`
		Message     *string             `json:"message,omitempty"`
		RedirectURL *string             `json:"redirectUrl,omitempty"`
		Status      *int                `json:"status,omitempty"`
		Headers     []map[string]string `json:"headers,omitempty"`
		Body        *string             `json:"body,omitempty"`
	}
)

func (l *localUtilsImpl) Zip(options localUtilsZipOptions) (interface{}, error) {
	return l.channel.Send("zip", options)
}

func (l *localUtilsImpl) HarOpen(file string) (string, error) {
	harId, err := l.channel.Send("harOpen", []map[string]interface{}{
		{
			"file": file,
		},
	})
	if harId == nil {
		return "", err
	}
	return harId.(string), err
}

func (l *localUtilsImpl) HarLookup(option harLookupOptions) (*harLookupResult, error) {
	overrides := make(map[string]interface{})
	overrides["harId"] = option.HarId
	overrides["url"] = option.URL
	overrides["method"] = option.Method
	if option.Headers != nil {
		overrides["headers"] = serializeMapToNameAndValue(option.Headers)
	}
	overrides["isNavigationRequest"] = option.IsNavigationRequest
	if option.PostData != nil {
		switch v := option.PostData.(type) {
		case string:
			overrides["postData"] = base64.StdEncoding.EncodeToString([]byte(v))
		case []byte:
			overrides["postData"] = base64.StdEncoding.EncodeToString(v)
		}
	}
	ret, err := l.channel.SendReturnAsDict("harLookup", overrides)
	if ret == nil {
		return nil, err
	}
	retMap, ok := ret.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected harLookupResult, got %T", retMap)
	}
	var result harLookupResult
	mJson, err := json.Marshal(ret)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(mJson, &result)
	if err != nil {
		return nil, err
	}
	if result.Body != nil {
		body, err := base64.StdEncoding.DecodeString(*result.Body)
		if err != nil {
			return nil, err
		}
		result.Body = String(string(body))
	}
	return &result, err
}

func (l *localUtilsImpl) HarClose(harId string) error {
	_, err := l.channel.Send("harClose", []map[string]interface{}{
		{
			"harId": harId,
		},
	})
	return err
}

func (l *localUtilsImpl) HarUnzip(zipFile, harFile string) error {
	_, err := l.channel.Send("harUnzip", []map[string]interface{}{
		{
			"zipFile": zipFile,
			"harFile": harFile,
		},
	})
	return err
}

func newLocalUtils(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *localUtilsImpl {
	l := &localUtilsImpl{}
	l.createChannelOwner(l, parent, objectType, guid, initializer)
	return l
}
