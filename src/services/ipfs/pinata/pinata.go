package pinata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"storage/src/pkg/logger"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var apiKey string
var apiSecret string
var apiJwt string
var host string
var unpinPath string
var pinPath string

var once sync.Once

func setup() {
	once.Do(func() {
		apiKey = viper.GetString("pinata.api_key")
		apiSecret = viper.GetString("pinata.secret")
		apiJwt = viper.GetString("pinata.jwt")
		host = viper.GetString("pinata.host")
		unpinPath = viper.GetString("pinata.unpin_path")
		pinPath = viper.GetString("pinata.pin_path")
	})
}

func UnpinFile(ctx context.Context, cid string) (bool, error) {
	setup()
	url := fmt.Sprintf("%s/%s/%s", host, unpinPath, cid)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", "Bearer "+apiJwt)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	if (res.StatusCode != http.StatusOK) && (res.StatusCode != http.StatusAccepted) {
		return false, fmt.Errorf("unpin failed: %d", res.StatusCode)
	}
	return true, nil
}

func UploadFileToIPFS(ctx context.Context, file *multipart.FileHeader) (resp UploadFileResponse, err error) {
	setup()
	log := logger.GetContextLogger(ctx)
	url := fmt.Sprintf("%s/%s", host, pinPath)

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Write the file to the multipart writer
	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		log.Error("failed to create form file", "error", err)
		return resp, err
	}

	fileStream, err := file.Open()
	if err != nil {
		log.Error("failed to open file", "error", err)
		return resp, err
	}
	defer fileStream.Close()

	_, err = io.Copy(part, fileStream)
	if err != nil {
		log.Error("failed to copy file", "error", err)
		return resp, err
	}

	// Write the pinataMetadata to the multipart writer
	metadataPart, err := writer.CreateFormField("pinataMetadata")
	if err != nil {
		log.Error("failed to create metadata part", "error", err)
		return resp, err
	}

	metadata := fmt.Sprintf(`{"name": "%s"}`, file.Filename)
	_, err = metadataPart.Write([]byte(metadata))
	if err != nil {
		log.Error("failed to write metadata", "error", err)
		return resp, err
	}

	// Write the pinataOptions to the multipart writer
	optionsPart, err := writer.CreateFormField("pinataOptions")
	if err != nil {
		log.Error("failed to create options part", "error", err)
		return resp, err
	}

	options := `{"cidVersion": 1}`
	_, err = optionsPart.Write([]byte(options))
	if err != nil {
		log.Error("failed to write options", "error", err)
		return resp, err
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		log.Error("failed to close writer", "error", err)
		return resp, err
	}

	// Set the content type to multipart/form-data
	contentType := writer.FormDataContentType()

	// Send the request
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		log.Error("failed to create request", "error", err)
		return resp, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer "+apiJwt)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("failed to send request", "error", err)
		return resp, err
	}
	defer res.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("failed to read response body", "error", err)
		return resp, err
	}

	// Unmarshal the response

	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		log.Error("failed to unmarshal response", "error", err)
		return resp, err
	}
	resp.IpfsStorage = "pinata"

	if res.StatusCode != http.StatusOK {
		log.Error("failed to upload file", "response", resp)
		return resp, err
	}

	return resp, nil
}

type UploadFileResponse struct {
	IpfsHash    string    `json:"IpfsHash",omitempty"`
	PinSize     int64     `json:"PinSize",omitempty"`
	Timestamp   time.Time `json:"Timestamp,omitempty"`
	IsDuplicate bool      `json:"isDuplicate",omitempty"`
	Error       string    `json:"error",omitempty"`
	IpfsStorage string    `json:"ipfsStorage,omitempty"`
}
