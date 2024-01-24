package playwright

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const fileSizeLimitInBytes = 50 * 1024 * 1024

var ErrInputFilesSizeExceeded = errors.New("Cannot set buffer larger than 50Mb, please write it to a file and pass its path instead.")

type inputFiles struct {
	Selector   *string             `json:"selector,omitempty"`
	Streams    []*channel          `json:"streams,omitempty"` // writableStream
	LocalPaths []string            `json:"localPaths,omitempty"`
	Payloads   []map[string]string `json:"payloads,omitempty"`
}

// convertInputFiles converts files to proper format for Playwright
//
//   - files should be one of: string, []string, InputFile, []InputFile,
//     string: local file path
func convertInputFiles(files interface{}, context *browserContextImpl) (*inputFiles, error) {
	converted := &inputFiles{}
	switch items := files.(type) {
	case InputFile:
		if sizeOfInputFiles([]InputFile{items}) > fileSizeLimitInBytes {
			return nil, ErrInputFilesSizeExceeded
		}
		converted.Payloads = normalizeFilePayloads([]InputFile{items})
	case []InputFile:
		if sizeOfInputFiles(items) > fileSizeLimitInBytes {
			return nil, ErrInputFilesSizeExceeded
		}
		converted.Payloads = normalizeFilePayloads(items)
	case string: // local file path
		converted.LocalPaths = []string{items}
	case []string:
		converted.LocalPaths = items
	default:
		return nil, errors.New("files should be one of: string, []string, InputFile, []InputFile")
	}
	if len(converted.LocalPaths) > 0 && context.connection.isRemote {
		converted.Streams = make([]*channel, 0)
		for _, file := range converted.LocalPaths {
			lastModifiedMs, err := getFileLastModifiedMs(file)
			if err != nil {
				return nil, fmt.Errorf("failed to get last modified time of %s: %w", file, err)
			}
			result, err := context.connection.WrapAPICall(func() (interface{}, error) {
				return context.channel.Send("createTempFile", map[string]interface{}{
					"name":           filepath.Base(file),
					"lastModifiedMs": lastModifiedMs,
				})
			}, true)
			if err != nil {
				return nil, err
			}
			stream := fromChannel(result).(*writableStream)
			if err := stream.Copy(file); err != nil {
				return nil, err
			}
			converted.Streams = append(converted.Streams, stream.channel)
		}
		converted.LocalPaths = nil
	}

	return converted, nil
}

func getFileLastModifiedMs(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	if info.IsDir() {
		return 0, fmt.Errorf("%s is a directory", path)
	}
	return info.ModTime().UnixMilli(), nil
}

func sizeOfInputFiles(files []InputFile) int {
	size := 0
	for _, file := range files {
		size += len(file.Buffer)
	}
	return size
}

func normalizeFilePayloads(files []InputFile) []map[string]string {
	out := make([]map[string]string, 0)
	for _, file := range files {
		out = append(out, map[string]string{
			"name":     file.Name,
			"mimeType": file.MimeType,
			"buffer":   base64.StdEncoding.EncodeToString(file.Buffer),
		})
	}
	return out
}
