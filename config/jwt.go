package config

type JWT struct {
	SigningKey  string `yaml:"signing-key"`
	ExpiresTime string `yaml:"expires-time"`
	BufferTime  string `yaml:"buffer-time"`
	Issuer      string `yaml:"issuer"`
}
