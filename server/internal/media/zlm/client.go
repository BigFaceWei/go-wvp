package zlm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Client struct {
	config  *ZLMConfig
	client  *http.Client
	running bool
	mu      sync.RWMutex
	logger  *zap.Logger
}

type ZLMConfig struct {
	ID        string
	IP        string
	HTTPPort  int
	Secret    string
	Default   bool
	API       string
}

func NewClient(config *ZLMConfig, logger *zap.Logger) *Client {
	return &Client{
		config: config,
		client: &http.Client{Timeout: 10 * time.Second},
		logger: logger,
	}
}

func (c *Client) GetAPI() string {
	return fmt.Sprintf("http://%s:%d", c.config.IP, c.config.HTTPPort)
}

func (c *Client) GetID() string {
	return c.config.ID
}

func (c *Client) IsDefault() bool {
	return c.config.Default
}

func (c *Client) callAPI(apiPath string, params map[string]string) (map[string]interface{}, error) {
	apiURL := fmt.Sprintf("%s%s", c.GetAPI(), apiPath)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	q := req.URL.Query()
	q.Set("secret", c.config.Secret)
	for k, v := range params {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call API failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response failed: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	return result, nil
}

func (c *Client) GetMediaList(schema, vhost, app, stream string) ([]map[string]interface{}, error) {
	params := map[string]string{
		"schema": schema,
		"vhost":  vhost,
		"app":    app,
		"stream": stream,
	}

	result, err := c.callAPI("/index/api/getMediaList", params)
	if err != nil {
		return nil, err
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return nil, fmt.Errorf("ZLM API error: %v", result["msg"])
	}

	if list, ok := result["data"].([]interface{}); ok {
		var mediaList []map[string]interface{}
		for _, item := range list {
			if mediaMap, ok := item.(map[string]interface{}); ok {
				mediaList = append(mediaList, mediaMap)
			}
		}
		return mediaList, nil
	}

	return nil, nil
}

func (c *Client) CloseStream(schema, vhost, app, stream string) error {
	params := map[string]string{
		"schema": schema,
		"vhost":  vhost,
		"app":    app,
		"stream": stream,
	}

	result, err := c.callAPI("/index/api/close_stream", params)
	if err != nil {
		return err
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return fmt.Errorf("close stream failed: %v", result["msg"])
	}

	return nil
}

func (c *Client) GetServerConfig() (map[string]interface{}, error) {
	return c.callAPI("/index/api/getServerConfig", nil)
}

func (c *Client) GetRTPPort() (int, error) {
	params := url.Values{}
	params.Set("secret", c.config.Secret)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/index/api/openRtpServer?%s", c.GetAPI(), params.Encode()), nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if code, ok := result["code"].(float64); ok && code != 0 {
		return 0, fmt.Errorf("open RTP server failed: %v", result["msg"])
	}

	if port, ok := result["port"].(float64); ok {
		return int(port), nil
	}

	return 0, nil
}

type StreamInfo struct {
	Schema    string `json:"schema"`
	VHost     string `json:"vhost"`
	App       string `json:"app"`
	Stream    string `json:"stream"`
	VirtualPath string `json:"virtualPath"`
	CreateTime  int64  `json:"createStamp"`
	Status    string `json:"status"`
	BytesSpeed int    `json:"bytesSpeed"`
}

func (c *Client) GetStreamList(app, stream string) ([]StreamInfo, error) {
	params := map[string]string{
		"app":    app,
		"stream": stream,
	}

	result, err := c.callAPI("/index/api/getMediaList", params)
	if err != nil {
		return nil, err
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return nil, fmt.Errorf("get stream list failed: %v", result["msg"])
	}

	var streams []StreamInfo
	if data, ok := result["data"].([]interface{}); ok {
		for _, item := range data {
			if itemMap, ok := item.(map[string]interface{}); ok {
				streamInfo := StreamInfo{
					Schema:   getString(itemMap, "schema"),
					VHost:    getString(itemMap, "vhost"),
					App:      getString(itemMap, "app"),
					Stream:   getString(itemMap, "stream"),
					CreateTime: getInt64(itemMap, "createStamp"),
				}
				streams = append(streams, streamInfo)
			}
		}
	}

	return streams, nil
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func getInt64(m map[string]interface{}, key string) int64 {
	if v, ok := m[key].(float64); ok {
		return int64(v)
	}
	return 0
}

// AddStreamProxy adds a stream proxy in ZLMediaKit to pull an external stream.
// Returns the stream key on success.
func (c *Client) AddStreamProxy(vhost, app, stream, url, srcURL, timeoutSec string) (string, error) {
	params := map[string]string{
		"vhost": vhost,
		"app":   app,
		"stream": stream,
		"url":   url,
		"src_url": srcURL,
		"timeout_sec": timeoutSec,
		"enable_rtsp": "1",
		"enable_rtmp": "1",
		"enable_hls":  "1",
		"enable_fmp4": "1",
		"enable_ts":   "1",
		"enable_audio": "1",
	}

	result, err := c.callAPI("/index/api/addStreamProxy", params)
	if err != nil {
		return "", fmt.Errorf("add stream proxy failed: %w", err)
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return "", fmt.Errorf("add stream proxy error: %v", result["msg"])
	}

	if key, ok := result["key"].(string); ok {
		return key, nil
	}
	return "", nil
}

// DelStreamProxy removes a stream proxy from ZLMediaKit.
func (c *Client) DelStreamProxy(vhost, app, stream string) error {
	params := map[string]string{
		"vhost":  vhost,
		"app":    app,
		"stream": stream,
	}

	result, err := c.callAPI("/index/api/delStreamProxy", params)
	if err != nil {
		return fmt.Errorf("del stream proxy failed: %w", err)
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return fmt.Errorf("del stream proxy error: %v", result["msg"])
	}

	return nil
}

// AddStreamPusherProxy adds a stream pusher in ZLMediaKit to push a stream to a target.
// Returns the push key on success.
func (c *Client) AddStreamPusherProxy(vhost, app, stream, dstURL string) (string, error) {
	params := map[string]string{
		"vhost":  vhost,
		"app":    app,
		"stream": stream,
		"dst_url": dstURL,
	}

	result, err := c.callAPI("/index/api/addStreamPusherProxy", params)
	if err != nil {
		return "", fmt.Errorf("add stream pusher failed: %w", err)
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return "", fmt.Errorf("add stream pusher error: %v", result["msg"])
	}

	if key, ok := result["key"].(string); ok {
		return key, nil
	}
	return "", nil
}

// DelStreamPusherProxy removes a stream pusher from ZLMediaKit.
func (c *Client) DelStreamPusherProxy(vhost, app, stream string) error {
	params := map[string]string{
		"vhost":  vhost,
		"app":    app,
		"stream": stream,
	}

	result, err := c.callAPI("/index/api/delStreamPusherProxy", params)
	if err != nil {
		return fmt.Errorf("del stream pusher failed: %w", err)
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return fmt.Errorf("del stream pusher error: %v", result["msg"])
	}

	return nil
}

// SetRecordSpeed sets recording status for a specific stream (app/stream).
func (c *Client) SetRecordSpeed(vhost, app, stream string, speed int) error {
	params := map[string]string{
		"vhost":  vhost,
		"app":    app,
		"stream": stream,
		"speed":  fmt.Sprintf("%d", speed),
	}
	_, err := c.callAPI("/index/api/setRecordSpeed", params)
	return err
}

// AddFFmpegSource adds an ffmpeg pull source.
func (c *Client) AddFFmpegSource(srcURL, dstURL, timeoutSec string) (string, error) {
	params := map[string]string{
		"src_url":     srcURL,
		"dst_url":     dstURL,
		"timeout_sec": timeoutSec,
	}

	result, err := c.callAPI("/index/api/addFFmpegSource", params)
	if err != nil {
		return "", fmt.Errorf("add ffmpeg source failed: %w", err)
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return "", fmt.Errorf("add ffmpeg source error: %v", result["msg"])
	}

	if key, ok := result["key"].(string); ok {
		return key, nil
	}
	return "", nil
}

// DelFFmpegSource removes an ffmpeg pull source.
func (c *Client) DelFFmpegSource(key string) error {
	params := map[string]string{
		"key": key,
	}
	_, err := c.callAPI("/index/api/delFFmpegSource", params)
	return err
}