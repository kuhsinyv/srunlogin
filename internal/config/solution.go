package config

import (
	"errors"
)

// SolutionName 运行方式名称
type SolutionName string

const (
	// CDP 使用 chromedp 执行
	CDP SolutionName = "cdp"
)

var (
	// ErrInvalidSolutionName 非法的运行方式
	ErrInvalidSolutionName = errors.New("invalid domain")
)

// CheckSolutionName 判断配置中的 domain 是否正确
func CheckSolutionName(domain string) error {
	switch SolutionName(domain) {
	case CDP:
		return nil
	default:
		return ErrInvalidSolutionName
	}
}
