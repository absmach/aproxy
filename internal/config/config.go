package config

import (
	"io"
	"os"
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/pelletier/go-toml/v2"
)

type MQTTAdapterConfig struct {
	MQTTPort              string   `toml:"PORT"              env:"APROXY_MQTT_ADAPTER_MQTT_PORT"                envDefault:"1883"`
	MQTTTargetHost        string   `toml:"TARGET_HOST"       env:"APROXY_MQTT_ADAPTER_MQTT_TARGET_HOST"         envDefault:"localhost"`
	MQTTTargetPort        string   `toml:"TARGET_PORT"       env:"APROXY_MQTT_ADAPTER_MQTT_TARGET_PORT"         envDefault:"1883"`
	MQTTForwarderTimeout  Duration `toml:"FORWARDER_TIMEOUT" env:"APROXY_MQTT_ADAPTER_FORWARDER_TIMEOUT"        envDefault:"30s"`
	MQTTTargetHealthCheck string   `toml:"HEALTH_CHECK"      env:"APROXY_MQTT_ADAPTER_MQTT_TARGET_HEALTH_CHECK" envDefault:""`
}

type HTTPAdapterConfig struct {
	HTTPPort       string `toml:"PORT"        env:"APROXY_MQTT_ADAPTER_WS_PORT"        envDefault:"8080"`
	HTTPTargetHost string `toml:"TARGET_HOST" env:"APROXY_MQTT_ADAPTER_WS_TARGET_HOST" envDefault:"localhost"`
	HTTPTargetPort string `toml:"TARGET_PORT" env:"APROXY_MQTT_ADAPTER_WS_TARGET_PORT" envDefault:"8080"`
	HTTPTargetPath string `toml:"TARGET_PATH" env:"APROXY_MQTT_ADAPTER_WS_TARGET_PATH" envDefault:"/mqtt"`
}

type GeneralConfig struct {
	LogLevel   string `toml:"LOG_LEVEL"   env:"APROXY_MQTT_ADAPTER_LOG_LEVEL"   envDefault:"info"`
	Instance   string `toml:"INSTANCE"    env:"APROXY_MQTT_ADAPTER_INSTANCE"    envDefault:""`
	JaegerURL  string `toml:"JAEGER_URL"  env:"APROXY_JAEGER_URL"               envDefault:"http://jaeger:14268/api/traces"`
	InstanceID string `toml:"INSTANCE_ID" env:"APROXY_MQTT_ADAPTER_INSTANCE_ID" envDefault:""`
}

type Config struct {
	MQTTAdapter MQTTAdapterConfig `toml:"MQTTAdapter"`
	HTTPAdapter HTTPAdapterConfig `toml:"HTTPAdapter"`
	General     GeneralConfig     `toml:"General"`
	ConfigFile  string            `toml:"-" env:"APROXY_MQTT_ADAPTER_CONFIG_FILE" envDefault:"config.toml"`
}

type Duration time.Duration

func (d *Duration) UnmarshalText(b []byte) error {
	x, err := time.ParseDuration(string(b))
	if err != nil {
		return err
	}
	*d = Duration(x)
	return nil
}

func parseConfigFile(cfg *Config) error {
	file, err := os.Open(cfg.ConfigFile)
	if err != nil {
		return err
	}
	fileData, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(fileData, cfg); err != nil {
		return err
	}

	return nil
}

func NewConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}
	if cfg.ConfigFile != "" {
		if err := parseConfigFile(&cfg); err != nil {
			return cfg, err
		}
	}
	return cfg, nil
}
