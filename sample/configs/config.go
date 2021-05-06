package configs

import (
	"lost_found/core/config"
	"lost_found/core/constants"
	"lost_found/core/log"
)

// for Cutome APP（企业自建应用）
var appSettings = config.GetInternalAppSettingsByEnv()

func TestConfigWithLogrusAndRedisStore(domain constants.Domain) *config.Config {
	logger := Logrus{}
	store := NewRedisStore()
	return config.NewConfig(domain, appSettings, logger, log.LevelDebug, store)
}

func TestConfig(domain constants.Domain) *config.Config {
	return config.NewConfigWithDefaultStore(domain, appSettings, log.NewDefaultLogger(), log.LevelDebug)
}
