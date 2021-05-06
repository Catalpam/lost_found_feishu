package config

import (
	"context"
	"lost_found/core"
	"lost_found/core/constants"
	"lost_found/core/log"
	"lost_found/core/store"
)

var ctxKeyConfig = "-----ctxKeyConfig"

type Config struct {
	domain      constants.Domain
	appSettings *AppSettings
	store       store.Store // store
	logger      log.Logger  // logger
}

func NewTestConfig(domain constants.Domain, appSettings *AppSettings) *Config {
	return NewConfigWithDefaultStore(domain, appSettings, log.NewDefaultLogger(), log.LevelDebug)
}

func NewConfigWithDefaultStore(domain constants.Domain, appSettings *AppSettings, logger log.Logger, logLevel log.Level) *Config {
	loggerProxy := log.NewLoggerProxy(logLevel, logger)
	conf := &Config{
		domain:      domain,
		appSettings: appSettings,
		store:       store.NewDefaultStoreWithLog(loggerProxy),
		logger:      loggerProxy,
	}
	return conf
}

func NewConfig(domain constants.Domain, appSettings *AppSettings, logger log.Logger, logLevel log.Level, store store.Store) *Config {
	loggerProxy := log.NewLoggerProxy(logLevel, logger)
	conf := &Config{
		domain:      domain,
		appSettings: appSettings,
		store:       store,
		logger:      loggerProxy,
	}
	return conf
}

func (c *Config) GetDomain() string {
	return string(c.domain)
}

func (c *Config) GetAppSettings() *AppSettings {
	return c.appSettings
}

func (c *Config) GetLogger() log.Logger {
	return c.logger
}

func (c *Config) GetStore() store.Store {
	return c.store
}

func (c *Config) WithContext(ctx *core.Context) {
	ctx.Set(ctxKeyConfig, c)
}

func ByCtx(ctx context.Context) *Config {
	c := ctx.Value(ctxKeyConfig)
	return c.(*Config)
}
