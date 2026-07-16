package zlm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
	ID           string
	IP           string
	HTTPPort     int  // ZLM HTTP port (stream serving, e.g. 10080)
	APIPort      int  // ZLM API port (default 8898, may be same as HTTPPort if configured)
	Secret       string
	Default      bool
	HookIP       string // IP of this WVP server for hooks (ZLM calls back)
	HookPort     int    // Port of this WVP server for hooks
	MediaHTTPPort int   // Media HTTP port for stream URLs (may differ from HTTPPort)
}

func NewClient(config *ZLMConfig, logger *zap.Logger) *Client {
	if config.APIPort == 0 {
		config.APIPort = config.HTTPPort // fallback: use HTTP port for API
	}
	if config.MediaHTTPPort == 0 {
		config.MediaHTTPPort = config.HTTPPort
	}
	return &Client{
		config: config,
		client: &http.Client{Timeout: 30 * time.Second},
		logger: logger,
	}
}

func (c *Client) GetAPI() string {
	return fmt.Sprintf("http://%s:%d", c.config.IP, c.config.APIPort)
}

func (c *Client) GetMediaBase() string {
	return fmt.Sprintf("http://%s:%d", c.config.IP, c.config.MediaHTTPPort)
}

func (c *Client) GetID() string {
	return c.config.ID
}

func (c *Client) IsDefault() bool {
	return c.config.Default
}

func (c *Client) GetHookAddr() string {
	return fmt.Sprintf("%s:%d", c.config.HookIP, c.config.HookPort)
}

// callAPI sends a POST request to the ZLMediaKit API (matching wvp-GB28181-pro's approach).
// All ZLM API requests are logged for debugging.
func (c *Client) callAPI(apiPath string, params url.Values) (map[string]interface{}, error) {
	apiURL := fmt.Sprintf("%s%s", c.GetAPI(), apiPath)

	if params == nil {
		params = url.Values{}
	}
	params.Set("secret", c.config.Secret)
	body := params.Encode()

	c.logger.Info("ZLM API request",
		zap.String("method", "POST"),
		zap.String("url", apiURL),
		zap.String("body", body),
	)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(body))
	if err != nil {
		c.logger.Error("ZLM API create request failed", zap.Error(err))
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("ZLM API request failed", zap.String("url", apiURL), zap.Error(err))
		return nil, fmt.Errorf("call API failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("ZLM API read response failed", zap.Error(err))
		return nil, fmt.Errorf("read response failed: %w", err)
	}

	c.logger.Info("ZLM API response",
		zap.Int("status", resp.StatusCode),
		zap.Int("body_len", len(respBody)),
	)

	if len(respBody) < 2048 {
		c.logger.Debug("ZLM API response body",
			zap.String("body", string(respBody)),
		)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		c.logger.Error("ZLM API parse response failed", zap.Error(err))
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	// Check for ZLM API error codes
	if code, ok := result["code"].(float64); ok && code != 0 {
		msg := ""
		if m, ok := result["msg"].(string); ok {
			msg = m
		}
		c.logger.Error("ZLM API returned error",
			zap.Float64("code", code),
			zap.String("msg", msg),
		)
	}

	return result, nil
}

func (c *Client) GetMediaList(schema, vhost, app, stream string) ([]map[string]interface{}, error) {
	params := url.Values{}
	params.Set("schema", schema)
	params.Set("vhost", vhost)
	params.Set("app", app)
	params.Set("stream", stream)

	result, err := c.callAPI("/index/api/getMediaList", params)
	if err != nil {
		return nil, fmt.Errorf("getMediaList failed: %w", err)
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
	params := url.Values{}
	params.Set("schema", schema)
	params.Set("vhost", vhost)
	params.Set("app", app)
	params.Set("stream", stream)

	result, err := c.callAPI("/index/api/close_stream", params)
	if err != nil {
		return err
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return fmt.Errorf("close stream failed: %v", result["msg"])
	}

	return nil
}

// GetServerConfig fetches the current ZLMediaKit server config via POST (matches wvp-GB28181-pro).
func (c *Client) GetServerConfig() (map[string]interface{}, error) {
	result, err := c.callAPI("/index/api/getServerConfig", nil)
	if err != nil {
		return nil, fmt.Errorf("getServerConfig failed: %w", err)
	}
	resultData, ok := result["data"].([]interface{})
	if !ok || len(resultData) == 0 {
		return nil, fmt.Errorf("getServerConfig: no config data returned")
	}
	configMap, ok := resultData[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("getServerConfig: unexpected config format")
	}
	return configMap, nil
}

// SetServerConfig sets ZLMediaKit server configuration via POST (matches wvp-GB28181-pro).
func (c *Client) SetServerConfig(config map[string]string) error {
	params := url.Values{}
	for k, v := range config {
		params.Set(k, v)
	}

	result, err := c.callAPI("/index/api/setServerConfig", params)
	if err != nil {
		return fmt.Errorf("setServerConfig failed: %w", err)
	}

	changed := false
	if ch, ok := result["changed"].(float64); ok && ch > 0 {
		changed = true
	}
	c.logger.Info("ZLM setServerConfig result",
		zap.Bool("changed", changed),
		zap.Int("code", int(result["code"].(float64))),
	)
	return nil
}

func (c *Client) GetRTPPort() (int, error) {
	params := url.Values{}
	params.Set("port", "0")
	result, err := c.callAPI("/index/api/openRtpServer", params)
	if err != nil {
		return 0, err
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
	params := url.Values{}
	params.Set("app", app)
	params.Set("stream", stream)

	result, err := c.callAPI("/index/api/getMediaList", params)
	if err != nil {
		return nil, fmt.Errorf("get stream list failed: %w", err)
	}

	var streams []StreamInfo
	if data, ok := result["data"].([]interface{}); ok {
		for _, item := range data {
			if itemMap, ok := item.(map[string]interface{}); ok {
				streamInfo := StreamInfo{
					Schema:     getString(itemMap, "schema"),
					VHost:      getString(itemMap, "vhost"),
					App:        getString(itemMap, "app"),
					Stream:     getString(itemMap, "stream"),
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
func (c *Client) AddStreamProxy(vhost, app, stream, streamURL, srcURL, timeoutSec string) (string, error) {
	params := url.Values{}
	params.Set("vhost", vhost)
	params.Set("app", app)
	params.Set("stream", stream)
	params.Set("url", streamURL)
	params.Set("src_url", srcURL)
	params.Set("timeout_sec", timeoutSec)
	params.Set("enable_rtsp", "1")
	params.Set("enable_rtmp", "1")
	params.Set("enable_hls", "1")
	params.Set("enable_fmp4", "1")
	params.Set("enable_ts", "1")
	params.Set("enable_audio", "1")

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
	params := url.Values{}
	params.Set("vhost", vhost)
	params.Set("app", app)
	params.Set("stream", stream)

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
	params := url.Values{}
	params.Set("vhost", vhost)
	params.Set("app", app)
	params.Set("stream", stream)
	params.Set("dst_url", dstURL)

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
	params := url.Values{}
	params.Set("vhost", vhost)
	params.Set("app", app)
	params.Set("stream", stream)

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
	params := url.Values{}
	params.Set("vhost", vhost)
	params.Set("app", app)
	params.Set("stream", stream)
	params.Set("speed", fmt.Sprintf("%d", speed))
	_, err := c.callAPI("/index/api/setRecordSpeed", params)
	return err
}

// AddFFmpegSource adds an ffmpeg pull source.
func (c *Client) AddFFmpegSource(srcURL, dstURL, timeoutSec string) (string, error) {
	params := url.Values{}
	params.Set("src_url", srcURL)
	params.Set("dst_url", dstURL)
	params.Set("timeout_sec", timeoutSec)

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
	params := url.Values{}
	params.Set("key", key)
	_, err := c.callAPI("/index/api/delFFmpegSource", params)
	return err
}