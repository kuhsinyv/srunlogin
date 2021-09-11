package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kuhsinyv/srunlogin/internal/config"
	"github.com/kuhsinyv/srunlogin/internal/solution"
	"github.com/kuhsinyv/srunlogin/internal/solution/cdp"
	"github.com/kuhsinyv/srunlogin/tools"
)

var (
	_rootAbsolutePath string
)

var (
	_configRelativePathFlag string
)

func getRootAbsolutePath() error {
	rootPath, err := tools.GetRootAbsPath()
	if err != nil {
		return err
	}

	_rootAbsolutePath = rootPath

	return nil
}

func parseFlags() {
	flag.StringVar(&_configRelativePathFlag, "config", config.DefaultConfigRelativePath, "relative path to executable program")

	flag.Parse()
}

func retry(times uint32, sln solution.Solution) (err error) {
	for i := 0; i < int(times); i++ {
		if err = sln.Run(); err == nil {
			return nil
		}
	}

	return err
}

func deleteLnk() {
	_ = os.Remove(config.Config.APP.DeleteLnkPath)
}

func run() error {
	if err := getRootAbsolutePath(); err != nil {
		return err
	}

	parseFlags()

	if err := config.InitConfig(filepath.Join(_rootAbsolutePath, _configRelativePathFlag)); err != nil {
		return err
	}

	defer deleteLnk()

	time.Sleep(config.Config.APP.Delay)

	campusNet, err := config.ParseCampusNet()
	if err != nil {
		return err
	}

	if campusNet == "" {
		return nil
	}

	config.Config.Account.CampusNet = campusNet

	var sln solution.Solution

	switch config.SolutionName(config.Config.APP.Solution) {
	case config.CDP:
		sln = cdp.NewSolution(cdp.Options{
			CampusNet:  config.Config.Account.CampusNet,
			Username:   config.Config.Account.Username,
			Password:   config.Config.Account.Password,
			Domain:     config.Domain(config.Config.Account.Domain),
			Display:    config.Config.APP.Display,
			Timeout:    config.Config.APP.Timeout,
			LogEnabled: config.Config.APP.LogEnabled,
		})
	}

	if err = retry(config.Config.APP.Retry, sln); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
