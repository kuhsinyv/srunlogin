package retry

// RunFunc 需要重试的函数
type RunFunc func() error

// Retry 执行 RunFunc，当函数成功执行或者重试次数达到上限 times 时停止。
func Retry(times uint64, run RunFunc) error {
	var err error
	for ; times > 0; times-- {
		if err = run(); err == nil {
			return nil
		}
	}

	return err
}
