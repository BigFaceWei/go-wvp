package config

// WVP WVP-GB28181配置
type WVP struct {
	// SIP服务器配置
	SIP SIPConfig `mapstructure:"sip" json:"sip" yaml:"sip"`
	
	// ZLMediaKit配置
	ZLMediaKit ZLMediaKitConfig `mapstructure:"zlmediakit" json:"zlmediakit" yaml:"zlmediakit"`
	
	// 录像配置
	Record RecordConfig `mapstructure:"record" json:"record" yaml:"record"`
	
	// 用户设置
	UserSettings UserSettingsConfig `mapstructure:"user_settings" json:"user_settings" yaml:"user_settings"`
}

// SIPConfig SIP服务器配置
type SIPConfig struct {
	// 监听IP
	ListenIP string `mapstructure:"listen_ip" json:"listen_ip" yaml:"listen_ip"`
	
	// 监听端口
	ListenPort int `mapstructure:"listen_port" json:"listen_port" yaml:"listen_port"`
	
	// 传输协议（udp/tcp）
	Transport string `mapstructure:"transport" json:"transport" yaml:"transport"`
	
	// SIP域
	Domain string `mapstructure:"domain" json:"domain" yaml:"domain"`
	
	// 服务器ID
	ServerID string `mapstructure:"server_id" json:"server_id" yaml:"server_id"`
	
	// 公共密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	
	// 注册有效期（秒）
	Expires int `mapstructure:"expires" json:"expires" yaml:"expires"`
	
	// 最大重试次数
	MaxRetry int `mapstructure:"max_retry" json:"max_retry" yaml:"max_retry"`
	
	// 是否开启SIP信令日志
	SIPLog bool `mapstructure:"sip_log" json:"sip_log" yaml:"sip_log"`
}

// ZLMediaKitConfig ZLMediaKit配置
type ZLMediaKitConfig struct {
	// 流媒体服务器列表
	Servers []ZLMServerConfig `mapstructure:"servers" json:"servers" yaml:"servers"`
	
	// 流配置
	Stream StreamConfig `mapstructure:"stream" json:"stream" yaml:"stream"`
	
	// RTP配置
	RTP RTPConfig `mapstructure:"rtp" json:"rtp" yaml:"rtp"`
}

// ZLMServerConfig ZLMediaKit服务器配置
type ZLMServerConfig struct {
	// 服务器ID
	ID string `mapstructure:"id" json:"id" yaml:"id"`
	
	// 服务器IP
	IP string `mapstructure:"ip" json:"ip" yaml:"ip"`
	
	// HTTP端口
	HTTPPort int `mapstructure:"http_port" json:"http_port" yaml:"http_port"`
	
	// 密钥
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	
	// 是否默认服务器
	Default bool `mapstructure:"default" json:"default" yaml:"default"`
}

// StreamConfig 流配置
type StreamConfig struct {
	// 应用名
	App string `mapstructure:"app" json:"app" yaml:"app"`
	
	// 虚拟主机
	Vhost string `mapstructure:"vhost" json:"vhost" yaml:"vhost"`
}

// RTPConfig RTP配置
type RTPConfig struct {
	// 是否启用RTP代理
	Enable bool `mapstructure:"enable" json:"enable" yaml:"enable"`
	
	// RTP端口范围
	PortRange string `mapstructure:"port_range" json:"port_range" yaml:"port_range"`
	
	// 发送RTP端口范围
	SendPortRange string `mapstructure:"send_port_range" json:"send_port_range" yaml:"send_port_range"`
}

// RecordConfig 录像配置
type RecordConfig struct {
	// 是否启用录像
	Enable bool `mapstructure:"enable" json:"enable" yaml:"enable"`
	
	// 存储配置
	Storage StorageConfig `mapstructure:"storage" json:"storage" yaml:"storage"`
	
	// FFmpeg配置
	FFmpeg FFmpegConfig `mapstructure:"ffmpeg" json:"ffmpeg" yaml:"ffmpeg"`
}

// StorageConfig 存储配置
type StorageConfig struct {
	// 存储类型（local/minio/s3）
	Type string `mapstructure:"type" json:"type" yaml:"type"`
	
	// 存储路径
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}

// FFmpegConfig FFmpeg配置
type FFmpegConfig struct {
	// FFmpeg路径
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	
	// 线程数
	Threads int `mapstructure:"threads" json:"threads" yaml:"threads"`
}

// UserSettingsConfig 用户设置配置
type UserSettingsConfig struct {
	// 播放超时时间（毫秒）
	PlayTimeout int `mapstructure:"play_timeout" json:"play_timeout" yaml:"play_timeout"`
	
	// SSRC随机模式
	SSRCRandom bool `mapstructure:"ssrc_random" json:"ssrc_random" yaml:"ssrc_random"`
	
	// 自动点播
	AutoApplyPlay bool `mapstructure:"auto_apply_play" json:"auto_apply_play" yaml:"auto_apply_play"`
	
	// 流按需拉取
	StreamOnDemand bool `mapstructure:"stream_on_demand" json:"stream_on_demand" yaml:"stream_on_demand"`
}