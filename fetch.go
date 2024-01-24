package playwright

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type apiRequestImpl struct {
	*Playwright
}

func (r *apiRequestImpl) NewContext(options ...APIRequestNewContextOptions) (APIRequestContext, error) {
	overrides := map[string]interface{}{}
	if len(options) == 1 {
		if options[0].ExtraHttpHeaders != nil {
			overrides["extraHTTPHeaders"] = serializeMapToNameAndValue(options[0].ExtraHttpHeaders)
			options[0].ExtraHttpHeaders = nil
		}
		if options[0].StorageStatePath != nil {
			var storageState *StorageState
			storageString, err := os.ReadFile(*options[0].StorageStatePath)
			if err != nil {
				return nil, fmt.Errorf("could not read storage state file: %w", err)
			}
			err = json.Unmarshal(storageString, &storageState)
			if err != nil {
				return nil, fmt.Errorf("could not parse storage state file: %w", err)
			}
			options[0].StorageState = storageState
			options[0].StorageStatePath = nil
		}
	}

	channel, err := r.channel.Send("newRequest", options, overrides)
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*apiRequestContextImpl), nil
}

func newApiRequestImpl(pw *Playwright) *apiRequestImpl {
	return &apiRequestImpl{pw}
}

type apiRequestContextImpl struct {
	channelOwner
	tracing *tracingImpl
}

func (r *apiRequestContextImpl) Dispose() error {
	_, err := r.channel.Send("dispose")
	return err
}

func (r *apiRequestContextImpl) Delete(url string, options ...APIRequestContextDeleteOptions) (APIResponse, error) {
	opts := APIRequestContextFetchOptions{
		Method: String("DELETE"),
	}
	if len(options) == 1 {
		err := assignStructFields(&opts, options[0], false)
		if err != nil {
			return nil, err
		}
	}

	return r.Fetch(url, opts)
}

func (r *apiRequestContextImpl) Fetch(urlOrRequest interface{}, options ...APIRequestContextFetchOptions) (APIResponse, error) {
	switch v := urlOrRequest.(type) {
	case string:
		return r.innerFetch(v, nil, options...)
	case Request:
		return r.innerFetch("", v, options...)
	default:
		return nil, fmt.Errorf("urlOrRequest has unsupported type: %T", urlOrRequest)
	}
}

func (r *apiRequestContextImpl) innerFetch(url string, request Request, options ...APIRequestContextFetchOptions) (APIResponse, error) {
	overrides := map[string]interface{}{}
	if url != "" {
		overrides["url"] = url
	} else if request != nil {
		overrides["url"] = request.URL()
	}

	if len(options) == 1 {
		if options[0].MaxRedirects != nil && *options[0].MaxRedirects < 0 {
			return nil, errors.New("maxRedirects must be non-negative")
		}
		// only one of them can be specified
		if countNonNil(options[0].Data, options[0].Form, options[0].Multipart) > 1 {
			return nil, errors.New("only one of 'data', 'form' or 'multipart' can be specified")
		}
		if options[0].Method == nil {
			if request != nil {
				options[0].Method = String(request.Method())
			} else {
				options[0].Method = String("GET")
			}
		}
		if options[0].Headers == nil {
			if request != nil {
				overrides["headers"] = serializeMapToNameAndValue(request.Headers())
			}
		} else {
			overrides["headers"] = serializeMapToNameAndValue(options[0].Headers)
			options[0].Headers = nil
		}
		if options[0].Data != nil {
			switch v := options[0].Data.(type) {
			case string:
				headersArray, ok := overrides["headers"].([]map[string]string)
				if ok && isJsonContentType(headersArray) {
					if json.Valid([]byte(v)) {
						overrides["jsonData"] = v
					} else {
						data, err := json.Marshal(v)
						if err != nil {
							return nil, fmt.Errorf("could not marshal data: %w", err)
						}
						overrides["jsonData"] = string(data)
					}
				} else {
					overrides["postData"] = base64.StdEncoding.EncodeToString([]byte(v))
				}
			case []byte:
				overrides["postData"] = base64.StdEncoding.EncodeToString(v)
			case interface{}:
				data, err := json.Marshal(v)
				if err != nil {
					return nil, fmt.Errorf("could not marshal data: %w", err)
				}
				overrides["jsonData"] = string(data)
			default:
				return nil, errors.New("data must be a string, []byte, or interface{} that can marshal to json")
			}
			options[0].Data = nil
		} else if options[0].Form != nil {
			form, ok := options[0].Form.(map[string]interface{})
			if !ok {
				return nil, errors.New("form must be a map")
			}
			overrides["formData"] = serializeMapToNameValue(form)
			options[0].Form = nil
		} else if options[0].Multipart != nil {
			_, ok := options[0].Multipart.(map[string]interface{})
			if !ok {
				return nil, errors.New("multipart must be a map")
			}
			multipartData := []map[string]interface{}{}
			for name, value := range options[0].Multipart.(map[string]interface{}) {
				switch v := value.(type) {
				case InputFile:
					multipartData = append(multipartData, map[string]interface{}{
						"name": name,
						"file": map[string]string{
							"name":     v.Name,
							"mimeType": v.MimeType,
							"buffer":   base64.StdEncoding.EncodeToString(v.Buffer),
						},
					})
				default:
					multipartData = append(multipartData, map[string]interface{}{
						"name":  name,
						"value": String(fmt.Sprintf("%v", v)),
					})
				}
			}
			overrides["multipartData"] = multipartData
			options[0].Multipart = nil
		} else if request != nil {
			postDataBuf, err := request.PostDataBuffer()
			if err == nil {
				overrides["postData"] = base64.StdEncoding.EncodeToString(postDataBuf)
			}
		}
		if options[0].Params != nil {
			overrides["params"] = serializeMapToNameValue(options[0].Params)
			options[0].Params = nil
		}
	}

	response, err := r.channel.Send("fetch", options, overrides)
	if err != nil {
		return nil, err
	}

	return newAPIResponse(r, response.(map[string]interface{})), nil
}

func (r *apiRequestContextImpl) Get(url string, options ...APIRequestContextGetOptions) (APIResponse, error) {
	opts := APIRequestContextFetchOptions{
		Method: String("GET"),
	}
	if len(options) == 1 {
		err := assignStructFields(&opts, options[0], false)
		if err != nil {
			return nil, err
		}
	}

	return r.Fetch(url, opts)
}

func (r *apiRequestContextImpl) Head(url string, options ...APIRequestContextHeadOptions) (APIResponse, error) {
	opts := APIRequestContextFetchOptions{
		Method: String("HEAD"),
	}
	if len(options) == 1 {
		err := assignStructFields(&opts, options[0], false)
		if err != nil {
			return nil, err
		}
	}

	return r.Fetch(url, opts)
}

func (r *apiRequestContextImpl) Patch(url string, options ...APIRequestContextPatchOptions) (APIResponse, error) {
	opts := APIRequestContextFetchOptions{
		Method: String("PATCH"),
	}
	if len(options) == 1 {
		err := assignStructFields(&opts, options[0], false)
		if err != nil {
			return nil, err
		}
	}

	return r.Fetch(url, opts)
}

func (r *apiRequestContextImpl) Put(url string, options ...APIRequestContextPutOptions) (APIResponse, error) {
	opts := APIRequestContextFetchOptions{
		Method: String("PUT"),
	}
	if len(options) == 1 {
		err := assignStructFields(&opts, options[0], false)
		if err != nil {
			return nil, err
		}
	}

	return r.Fetch(url, opts)
}

func (r *apiRequestContextImpl) Post(url string, options ...APIRequestContextPostOptions) (APIResponse, error) {
	opts := APIRequestContextFetchOptions{
		Method: String("POST"),
	}
	if len(options) == 1 {
		err := assignStructFields(&opts, options[0], false)
		if err != nil {
			return nil, err
		}
	}

	return r.Fetch(url, opts)
}

func (r *apiRequestContextImpl) StorageState(path ...string) (*StorageState, error) {
	result, err := r.channel.SendReturnAsDict("storageState")
	if err != nil {
		return nil, err
	}
	if len(path) == 1 {
		file, err := os.Create(path[0])
		if err != nil {
			return nil, err
		}
		if err := json.NewEncoder(file).Encode(result); err != nil {
			return nil, err
		}
		if err := file.Close(); err != nil {
			return nil, err
		}
	}
	var storageState StorageState
	remapMapToStruct(result, &storageState)
	return &storageState, nil
}

func newAPIRequestContext(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *apiRequestContextImpl {
	rc := &apiRequestContextImpl{}
	rc.createChannelOwner(rc, parent, objectType, guid, initializer)
	rc.tracing = fromChannel(initializer["tracing"]).(*tracingImpl)
	return rc
}

type apiResponseImpl struct {
	request     *apiRequestContextImpl
	initializer map[string]interface{}
	headers     *rawHeaders
}

func (r *apiResponseImpl) Body() ([]byte, error) {
	result, err := r.request.channel.SendReturnAsDict("fetchResponseBody", []map[string]interface{}{
		{
			"fetchUid": r.fetchUid(),
		},
	})
	if err != nil {
		if errors.Is(err, ErrTargetClosed) {
			return nil, errors.New("response has been disposed")
		}
		return nil, err
	}
	body := result.(map[string]interface{})["binary"]
	if body == nil {
		return nil, errors.New("response has been disposed")
	}
	return base64.StdEncoding.DecodeString(body.(string))
}

func (r *apiResponseImpl) Dispose() error {
	_, err := r.request.channel.Send("disposeAPIResponse", []map[string]interface{}{
		{
			"fetchUid": r.fetchUid(),
		},
	})
	return err
}

func (r *apiResponseImpl) Headers() map[string]string {
	return r.headers.Headers()
}

func (r *apiResponseImpl) HeadersArray() []NameValue {
	return r.headers.HeadersArray()
}

func (r *apiResponseImpl) JSON(v interface{}) error {
	body, err := r.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &v)
}

func (r *apiResponseImpl) Ok() bool {
	return r.Status() == 0 || (r.Status() >= 200 && r.Status() <= 299)
}

func (r *apiResponseImpl) Status() int {
	return int(r.initializer["status"].(float64))
}

func (r *apiResponseImpl) StatusText() string {
	return r.initializer["statusText"].(string)
}

func (r *apiResponseImpl) Text() (string, error) {
	body, err := r.Body()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (r *apiResponseImpl) URL() string {
	return r.initializer["url"].(string)
}

func (r *apiResponseImpl) fetchUid() string {
	return r.initializer["fetchUid"].(string)
}

func (r *apiResponseImpl) fetchLog() ([]string, error) {
	ret, err := r.request.channel.Send("fetchLog", map[string]interface{}{
		"fetchUid": r.fetchUid(),
	})
	if err != nil {
		return nil, err
	}
	result := make([]string, len(ret.([]interface{})))
	for i, v := range ret.([]interface{}) {
		result[i] = v.(string)
	}
	return result, nil
}

func newAPIResponse(context *apiRequestContextImpl, initializer map[string]interface{}) *apiResponseImpl {
	return &apiResponseImpl{
		request:     context,
		initializer: initializer,
		headers:     newRawHeaders(initializer["headers"]),
	}
}

func countNonNil(args ...interface{}) int {
	count := 0
	for _, v := range args {
		if v != nil {
			count++
		}
	}
	return count
}

func isJsonContentType(headers []map[string]string) bool {
	if len(headers) > 0 {
		for _, v := range headers {
			if strings.ToLower(v["name"]) == "content-type" {
				if v["value"] == "application/json" {
					return true
				}
			}
		}
	}
	return false
}

func serializeMapToNameValue(data map[string]interface{}) []map[string]string {
	serialized := make([]map[string]string, 0, len(data))
	for k, v := range data {
		serialized = append(serialized, map[string]string{
			"name":  k,
			"value": fmt.Sprintf("%v", v),
		})
	}
	return serialized
}
