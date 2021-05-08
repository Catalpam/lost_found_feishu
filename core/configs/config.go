package configs

import (
	"lost_found/core/config"
	"lost_found/core/constants"
	"lost_found/core/log"
)

// for Cutome APP（企业自建应用）
var appSettings = config.GetInternalAppSettingsByEnv()

func FeishuConfig(domain constants.Domain) *config.Config {
	return config.NewConfigWithDefaultStore(domain, appSettings, log.NewDefaultLogger(), log.LevelDebug)
}
