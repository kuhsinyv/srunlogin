package cdp

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"github.com/kuhsinyv/srunlogin/internal/config"
)

// Options chromedp 方式的配置
type Options struct {
	CampusNet  config.CampusNet
	Domain     config.Domain
	Username   string
	Password   string
	Display    bool
	Timeout    time.Duration
	LogEnabled bool
}

// Solution 使用 chromedp
type Solution struct {
	opts Options
}

// NewSolution 构建 Options
func NewSolution(opts Options) *Solution {
	return &Solution{opts: opts}
}

// Run 执行 Options
func (sln *Solution) Run() error {
	return sln.run()
}

func (sln *Solution) parseDomain() {
	var keyArrowDownCount int

	switch sln.opts.Domain {
	case config.CMCC:
		keyArrowDownCount = 1
	case config.NCU:
		keyArrowDownCount = 2
	case config.NDCARD:
		keyArrowDownCount = 3
	case config.UNICOM:
		keyArrowDownCount = 4
	}

	domainSelectActions := make([]string, keyArrowDownCount+1)

	for i := 0; i < keyArrowDownCount; i++ {
		domainSelectActions[i] = kb.ArrowDown
	}

	domainSelectActions[len(domainSelectActions)-1] = kb.Enter

	sln.opts.Domain = config.Domain(strings.Join(domainSelectActions, ""))
}

func (sln *Solution) getActions() chromedp.Tasks {
	sln.parseDomain()

	switch sln.opts.CampusNet {
	case config.NCUPublic:
		return chromedp.Tasks{
			chromedp.Navigate(string(sln.opts.CampusNet)),

			chromedp.WaitVisible(`//input[@id="username"]`, chromedp.BySearch),
			chromedp.SendKeys(`//input[@id="username"]`, sln.opts.Username, chromedp.BySearch),

			chromedp.WaitVisible(`//input[@id="password"]`, chromedp.BySearch),
			chromedp.SendKeys(`//input[@id="password"]`, sln.opts.Password, chromedp.BySearch),

			chromedp.WaitVisible(`//button[@id="login"]`, chromedp.BySearch),
			chromedp.Click(`//button[@id="login"]`, chromedp.BySearch),
			chromedp.WaitVisible(`//button[@id="logout"]`, chromedp.BySearch),
		}
	case config.NCUPrivate:
		return chromedp.Tasks{
			chromedp.Navigate(string(sln.opts.CampusNet)),

			chromedp.WaitVisible(`//input[@id="username"]`, chromedp.BySearch),
			chromedp.SendKeys(`//input[@id="username"]`, sln.opts.Username, chromedp.BySearch),

			chromedp.WaitVisible(`//input[@id="password"]`, chromedp.BySearch),
			chromedp.SendKeys(`//input[@id="password"]`, sln.opts.Password, chromedp.BySearch),

			chromedp.WaitVisible(`//select[@id="domain"]`, chromedp.BySearch),
			chromedp.Click(`//select[@id="domain"]`, chromedp.BySearch),
			chromedp.SendKeys(`//select[@id="domain"]`, string(sln.opts.Domain), chromedp.BySearch),

			chromedp.WaitVisible(`//button[@id="login"]`, chromedp.BySearch),
			chromedp.Click(`//button[@id="login"]`, chromedp.BySearch),
			chromedp.WaitVisible(`//button[@id="logout"]`, chromedp.BySearch),
		}
	}

	return nil
}

func (sln *Solution) run() error {
	parent, cancel := context.WithTimeout(context.TODO(), sln.opts.Timeout)
	defer cancel()

	var execAllocatorOptions []chromedp.ExecAllocatorOption

	if sln.opts.Display {
		execAllocatorOptions = append(
			execAllocatorOptions,
			chromedp.Flag("headless", false),
		)
	}

	execAllocator, cancel := chromedp.NewExecAllocator(parent, execAllocatorOptions...)
	defer cancel()

	var contextOptions []chromedp.ContextOption

	if sln.opts.LogEnabled {
		contextOptions = append(
			contextOptions,
			chromedp.WithDebugf(log.Printf),
		)
	}

	ctx, cancel := chromedp.NewContext(execAllocator, contextOptions...)
	defer cancel()

	actions := sln.getActions()
	if actions == nil {
		return config.ErrInvalidCampusNet
	}

	if err := chromedp.Run(ctx, actions); err != nil {
		log.Println("chromedp 执行任务失败：", err)
		return err
	}

	log.Println("chromedp 执行任务成功")

	return nil
}
