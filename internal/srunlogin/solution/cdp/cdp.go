package cdp

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// Options for the CDP Solution
type Options struct {
	LoginURL string
	ISP      string
	Username string
	Password string
	Timeout  time.Duration
	Flags    map[string]interface{}
}

// Solution using Chrome DP
type Solution struct {
	opts *Options
}

var selectISPExpr = `
	if ($("#domain").length > 0 && $("#domain")[0].options) {
		Array(...$("#domain")[0].options).find((option) => option.value === "%s").selected = true;
	}
`

// NewSolution build a new Solution with Options
func NewSolution(opts *Options) *Solution {
	return &Solution{opts}
}

// Run the solution
func (sln *Solution) Run() error {
	log.Println("开始执行 CDP 方案。")
	return sln.run()
}

func (sln *Solution) run() error {
	parent, cancel := context.WithTimeout(context.TODO(), sln.opts.Timeout)
	defer cancel()

	var execAllocatorOptions []chromedp.ExecAllocatorOption

	for k, v := range sln.opts.Flags {
		execAllocatorOptions = append(execAllocatorOptions, chromedp.Flag(k, v))
	}

	execAllocator, cancel := chromedp.NewExecAllocator(parent, execAllocatorOptions...)
	defer cancel()

	var contextOptions []chromedp.ContextOption

	ctx, cancel := chromedp.NewContext(execAllocator, contextOptions...)
	defer cancel()

	tasks := sln.getTasks()

	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}

	return nil
}

func (sln *Solution) getTasks() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(sln.opts.LoginURL),
		inputTask("username", sln.opts.Username),
		inputTask("password", sln.opts.Password),
		selectISPTask(sln.opts.ISP),
		loginTask(),
	}
}

func inputTask(id string, value string) chromedp.Tasks {
	sel := fmt.Sprintf(`//input[@id="%s"]`, id)

	return chromedp.Tasks{
		chromedp.WaitVisible(sel, chromedp.BySearch),
		chromedp.SendKeys(sel, value, chromedp.BySearch),
	}
}

func selectISPTask(isp string) chromedp.Tasks {
	expr := fmt.Sprintf(selectISPExpr, fmt.Sprintf("@%s", isp))

	return chromedp.Tasks{
		chromedp.Evaluate(expr, nil),
	}
}

func loginTask() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(`//button[@id="login"]`, chromedp.BySearch),
		chromedp.Click(`//button[@id="login"]`, chromedp.BySearch),
		chromedp.WaitVisible(`//button[@id="logout"]`, chromedp.BySearch),
	}
}
