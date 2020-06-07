package config

import "github.com/kelseyhightower/envconfig"

type TokenConfig struct {
	Time int `json:"-" envconfig:"TOKEN_TIME" default:"15"`
	Key string `json:"-" envconfig:"TOKEN_KEY" default:"estoesunallave"`
	Kid string `json:"-" envconfig:"TOKEN_KID" default:"tkone"`
}
func NewTokenConfig() (*TokenConfig, error) {
	tokenConfig := &TokenConfig{}
	if err := envconfig.Process("", tokenConfig); err != nil {
		return nil, err
	}
	return tokenConfig, nil
}
