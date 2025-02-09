package config

type Config struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	MySQL MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
