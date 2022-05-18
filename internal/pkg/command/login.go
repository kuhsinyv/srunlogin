package command

import (
	"errors"
	"log"

	"github.com/kuhsinyv/srunlogin/internal/srunlogin/solution"
	"github.com/kuhsinyv/srunlogin/internal/srunlogin/solution/cdp"
	"github.com/kuhsinyv/srunlogin/pkg/retry"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var (
	errSolutionUndefined = errors.New("solution is undefined")
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登录网络",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return probeCmd.RunE(cmd, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		retryTimes := viper.GetUint64("app.retry")
		if retryTimes > 1 {
			log.Println("重试功能开启，最大重试次数：", retryTimes)
		}
		slnName := viper.GetString("app.solution")
		log.Println("使用方案：", slnName)
		sln, err := getSolution(slnName)
		if err != nil {
			return err
		}

		if err := retry.Retry(viper.GetUint64("app.retry"), sln.Run); err == nil {
			log.Println("登录成功！")
		}

		return err
	},
}

func getSolution(solution string) (solution.Solution, error) {
	loginURL := viper.GetString("app.login-url")
	timeout := viper.GetDuration("app.timeout")
	username := viper.GetString("account.username")
	password := viper.GetString("account.password")
	isp := viper.GetString("account.isp")

	switch solution {
	case "cdp":
		cdpFlags := make(map[string]interface{})

		for k, v := range viper.GetStringMap("cdp.flags") {
			cdpFlags[k] = v
		}

		return cdp.NewSolution(&cdp.Options{
			LoginURL: loginURL,
			ISP:      isp,
			Username: username,
			Password: password,
			Timeout:  timeout,
			Flags:    cdpFlags,
		}), nil
	default:
		return nil, errSolutionUndefined
	}
}
