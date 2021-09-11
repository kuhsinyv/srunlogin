package config

import "errors"

// Domain 运营商
type Domain string

const (
	// CMCC 移动
	CMCC Domain = "@cmcc"
	// NCU 校园网
	NCU Domain = "@ncu"
	// NDCARD 电信
	NDCARD Domain = "@ndcard"
	// UNICOM 联通
	UNICOM Domain = "@unicom"
)

var (
	// ErrInvalidDomain 非法的运营商类型
	ErrInvalidDomain = errors.New("invalid domain")
)

// CheckDomain 判断配置中的 domain 是否正确
func CheckDomain(domain string) error {
	switch Domain(domain) {
	case CMCC:
		fallthrough
	case NCU:
		fallthrough
	case NDCARD:
		fallthrough
	case UNICOM:
		return nil
	default:
		return ErrInvalidDomain
	}
}
