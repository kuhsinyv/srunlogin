package command

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:     "srunlogin",
	Short:   "SrunLogin 是一个用于深澜软件管理的校园网的自动登录器。",
	Version: "0.2.0",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := probeCmd.RunE(cmd, args); err != nil {
			return err
		}

		return loginCmd.RunE(cmd, args)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configFile, "file", "f", "",
		"配置文件路径（如果不设置，将尝试在一些目录中寻找）")
}

// Execute 入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln("获取程序所在路径失败")
		return
	}

	if configFile != "" {
		viper.SetConfigFile(filepath.Join(pwd, configFile))
	} else {
		viper.AddConfigPath(pwd)
		viper.AddConfigPath(filepath.Join(os.Getenv("HOME"), ".config/srunlogin"))
		viper.AddConfigPath(filepath.Join(os.Getenv("HOME"), ".srunlogin"))
		viper.AddConfigPath(filepath.Join(os.Getenv("XDG_CONFIG_HOME"), ".config/srunlogin"))
		viper.AddConfigPath(filepath.Join(os.Getenv("XDG_CONFIG_HOME"), ".srunlogin"))
		viper.AddConfigPath(filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"), ".config/srunlogin"))
		viper.AddConfigPath(filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"), ".srunlogin"))
		viper.AddConfigPath(filepath.Join(os.Getenv("USERPROFILE"), ".config/srunlogin"))
		viper.AddConfigPath(filepath.Join(os.Getenv("USERPROFILE"), ".srunlogin"))
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("SRUN")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("获取配置文件失败：", err)
		return
	}
}
