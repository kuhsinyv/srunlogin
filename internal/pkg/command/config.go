package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置信息",
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := yaml.Marshal(viper.AllSettings())
		if err == nil {
			fmt.Println(string(out))
		}

		return err
	},
}
