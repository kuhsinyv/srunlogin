package tools

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	_absToRoot = `../../`
)

// GetRootAbsPath 兼容 go run 和直接执行可执行文件的获取项目根目录
func GetRootAbsPath() (string, error) {
	rootAbsolutePath, err := getRootAbsPathByExecutable()
	if err != nil {
		return "", err
	}

	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}

	if strings.Contains(rootAbsolutePath, tmpDir) {
		return getRootAbsPathByCaller(), nil
	}

	return rootAbsolutePath, nil
}

func getRootAbsPathByExecutable() (string, error) {
	rootAbsolutePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.EvalSymlinks(filepath.Dir(rootAbsolutePath))
}

func getRootAbsPathByCaller() string {
	_, rootAbsolutePath, _, _ := runtime.Caller(0)
	return filepath.Join(rootAbsolutePath, _absToRoot)
}
