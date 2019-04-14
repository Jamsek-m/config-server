package config

type Config struct {
	Service    _Service    `yaml:"service,omitempty"`
	Server     _Server     `yaml:"server,omitempty"`
	Datasource _Datasource `yaml:"datasource,omitempty"`
}
type _Service struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
	Env     string `yaml:"env,omitempty"`
}

type _Server struct {
	Port            int    `yaml:"port,omitempty"`
	BaseUrl         string `yaml:"base-url,omitempty"`
	SessionDuration int    `yaml:"session-duration,omitempty"`
}

type _Datasource struct {
	Type     string `yaml:"type,omitempty"`
	Location string `yaml:"location,omitempty"`
}
