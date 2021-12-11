# srunlogin
一个简单的南昌大学校园网自动登录实现
目前是通过 Chrome DP 操作的，方便维护。后续增加通过 net/http 方式。

## 使用说明

### Windows
- 下载 `release` 中相应的压缩包（如 `srunlogin_windows_amd64.zip`），解压后，根据说明编辑 `configs` 目录下的 `config.ini` 文件，双击 `srunlogin.exe` 即可运行（需要安装 Google Chrome 浏览器）。
- 添加到开机启动项
  - 为 `srunlogin.exe` 生成快捷方式；
  - 打开 `文件资源管理器`，在地址栏输入 `shell:startup`，回车进入；
  - 将第一步生成的快捷方式放到该目录下即可。

## 配置文件说明
```ini
[account]
# 账号
username = 1234567890

# 密码
password = 123456

# 运营商
# 校园网：@ncu；移动：@cmcc；联通：@unicom；电信：@ndcard。
domain = @unicom

[app]
# 用于判断连接的网络是 NCUWLAN，还是 NCU-2.4G / NCU-5G
# 必须设置为非强制使用 https 的网站
# 使用 http://baidu.com/ 亦可。
test_url = http://example.com/

# 运行方式
# 使用 chromedp，需要安装 Google Chrome 浏览器。
solution = cdp

# 当运行方式为 cdp 时，是否显示窗口
# true：显示，false：不显示
display = true

# 延后执行时间
# 示例：500ms，10s，1m30s
delay = 10s

# 重试次数
# 设置为 0 时，只执行 1 次。
retry = 3

# 每次执行的超时限制
# 最小值为 15s
timeout = 15s

# 是否显示 CDP 执行任务日志
# true：显示，false：显示
log_enabled = true

# 删除可能被自动创建的快捷方式
# 留空则不删除
delete_lnk_path = C:\Users\xxx\Desktop\Google Chrome.lnk
```
