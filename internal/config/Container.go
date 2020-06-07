package config

type Container interface {
	BuildContainer() error
}
