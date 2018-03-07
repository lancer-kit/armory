package log

type Config struct {
	AppName  string `json:"app_name" yaml:"app_name"`
	Level    string `json:"level" yaml:"level"`
	Sentry   string `json:"sentry" yaml:"sentry"`
	AddTrace bool   `json:"add_trace" yaml:"add_trace"`
	JSON     bool   `json:"json" yaml:"json"`
}

func (cfg Config) Validate() error {
	return nil
}
