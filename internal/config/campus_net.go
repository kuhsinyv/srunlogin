package config

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

// CampusNet 校园网类型
type CampusNet string

const (
	// NCUPublic NCUWLAN
	NCUPublic CampusNet = `http://222.204.3.221`
	// NCUPrivate NCU-2.4G / NCU-5G
	NCUPrivate CampusNet = `http://222.204.3.154`
)

var (
	// ErrInvalidCampusNet 非法的校园网类型
	ErrInvalidCampusNet = errors.New("invalid campus net")
	// ErrParseCampusNet 解析校园网类型失败
	ErrParseCampusNet = errors.New("failed to parse campus net")
)

// ParseCampusNet 获取校园网类型
func ParseCampusNet() (CampusNet, error) {
	resp, err := http.Get(Config.APP.TestURL)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	url := resp.Request.URL.String()

	if strings.Contains(url, string(NCUPublic)) {
		log.Println("即将登录 NCUWLAN")
		return NCUPublic, nil
	}

	if strings.Contains(url, string(NCUPrivate)) {
		log.Println("即将登录 NCU-2.4G / NCU-5G")
		return NCUPrivate, nil
	}

	if strings.Contains(url, Config.APP.TestURL) {
		log.Println("已经连接到互联网")
		return "", nil
	}

	log.Println("没有连接到正确的网络")

	return "", ErrParseCampusNet
}
