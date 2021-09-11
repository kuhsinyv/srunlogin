package config

import (
	"time"

	"gopkg.in/ini.v1"
)

// account 校园网账户信息
type account struct {
	CampusNet CampusNet
	Username  string `ini:"username"`
	Password  string `ini:"password"`
	Domain    string `ini:"domain"`
}

// app 可选的配置
type app struct {
	TestURL       string        `ini:"test_url"`
	Solution      string        `ini:"solution"`
	Display       bool          `ini:"display"`
	Delay         time.Duration `ini:"delay"`
	Retry         uint32        `ini:"retry"`
	Timeout       time.Duration `ini:"timeout"`
	LogEnabled    bool          `ini:"log_enabled"`
	DeleteLnkPath string        `ini:"delete_lnk_path"`
}

// config 配置
type config struct {
	Account account `ini:"account"`
	APP     app     `ini:"app"`
}

const (
	// DefaultConfigRelativePath 相对运行程序所在位置的相对路径
	DefaultConfigRelativePath = `./configs/config.ini`
)

var (
	// Config 配置实例
	Config config
)

// InitConfig 根据传入的配置路径初始化配置
func InitConfig(configPath string) error {
	cfg, err := ini.Load(configPath)
	if err != nil {
		return err
	}

	if err = cfg.MapTo(&Config); err != nil {
		return err
	}

	if err = CheckDomain(Config.Account.Domain); err != nil {
		return err
	}

	if err = CheckSolutionName(Config.APP.Solution); err != nil {
		return err
	}

	return nil
}
