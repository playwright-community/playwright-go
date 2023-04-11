package playwright

import "encoding/base64"

type localUtilsImpl struct {
	channelOwner
}

type (
	harLookupOptions struct {
		HarId               *string           `json:"harId"`
		Url                 *string           `json:"url,"`
		Method              *string           `json:"method"`
		Headers             map[string]string `json:"headers"`
		IsNavigationRequest *bool             `json:"isNavigationRequest"`
		PostData            []byte            `json:"postData,omitempty"`
	}
	harLookupResult struct {
		Action      *string
		Message     *string
		RedirectUrl *string
		Status      *string
		Headers     []map[string]string
		Body        *string
	}
)

func (l *localUtilsImpl) zip(zipFile string, entries map[string]string) error {
	_, err := l.channel.Send("zip", map[string]interface{}{
		"zipFile": zipFile,
		"entries": serializeMapToNameAndValue(entries),
	})
	return err
}

func (l *localUtilsImpl) harOpen(file string) error {
	_, err := l.channel.Send("harOpen", map[string]interface{}{
		"file": file,
	})
	return err
}

func (l *localUtilsImpl) harClose(harId string) error {
	_, err := l.channel.Send("harClose", map[string]interface{}{
		"harId": harId,
	})
	return err
}

func (l *localUtilsImpl) harUnzip(zipFile string, harFile string) error {
	_, err := l.channel.Send("harUnzip", map[string]interface{}{
		"zipFile": zipFile,
		"harFile": harFile,
	})
	return err
}

func (l *localUtilsImpl) harLookup(option harLookupOptions) (*harLookupResult, error) {
	overrides := make(map[string]interface{})
	if option.Url != nil {
		overrides["url"] = option.Url
	}
	if option.Method != nil {
		overrides["method"] = option.Method
	}
	if option.Headers != nil {
		overrides["headers"] = serializeMapToNameAndValue(option.Headers)
	}
	if option.PostData != nil {
		overrides["postData"] = base64.StdEncoding.EncodeToString(option.PostData)
	}
	result, err := l.channel.SendReturnAsDict("harLookup", overrides)
	if err != nil {
		return nil, err
	}
	var harResult harLookupResult
	remapMapToStruct(result, harResult)
	return &harResult, nil
}

// func (l *localUtilsImpl) connect(url string, Options ...interface{}) *jsonPipe {
// 	// TODO
// 	return nil
// }

func newLocalUtils(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *localUtilsImpl {
	l := &localUtilsImpl{}
	l.createChannelOwner(l, parent, objectType, guid, initializer)
	return l
}
