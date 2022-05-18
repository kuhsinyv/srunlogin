package command

import (
	"errors"
	"log"

	"github.com/kuhsinyv/srunlogin/internal/srunlogin/probe"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(probeCmd)
}

var probeCmd = &cobra.Command{
	Use:   "probe",
	Short: "探测登录页面",
	RunE: func(cmd *cobra.Command, args []string) error {
		url, err := probe.Probe(viper.GetString("app.test-url"))
		if err != nil {
			if errors.Is(err, probe.ErrNetworkLoggedIn) {
				log.Println("你已经登录了。")
				cmd.SilenceErrors = true
				cmd.SilenceUsage = true
			}
			return err
		}

		log.Println("登录页面：", url)
		viper.Set("app.login-url", url)

		return nil
	},
}
