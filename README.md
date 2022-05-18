# SrunLogin

一个简单的南昌大学校园网自动登录器的实现。

## 使用说明

### Windows

- 下载 `release` 中相应的压缩包（如 `srunlogin-windows-amd64.zip`）。

- 解压后，根据说明编辑 `config.yaml` 文件，双击 `srunlogin-windows-amd64.exe`
  即可运行。（如果使用 CDP，则需要安装 Google Chrome 浏览器。）

- 添加到开机启动项。

    - 为 `srunlogin-windows-amd64.exe` 生成快捷方式；

    - 打开 `文件资源管理器`，在地址栏输入 `shell:startup`，回车进入；

    - 将上面步骤中生成的快捷方式放到该目录下即可。

## 配置文件说明

```yaml
account:
  isp: <isp> # 运营商名称，可设置为 ncu、cmcc、unicom 或者 ndcard。具体参考“运营商名称说明”
  password: <password> # 帐号
  username: <username> # 密码
app:
  solution: cdp # 使用 Chrome DP
  test-url: http://www.baidu.com # 用于被校园网重定向到登录页面，请不要设置为能够在不登录的情况下还能解析的域名
  timeout: 60s # 每次尝试登录的超时时间
  retry: 3 # 最小有效值为 1
cdp:
  flags: # 设置参考 https://pkg.go.dev/github.com/chromedp/chromedp
    headless: true # 设置为 ture 则不打开浏览器窗口
    hide-scrollbars: true
    mute-audio: true
    no-default-browser-check: true # 可能需要设置 headless 为 true 才有效
```

### 运营商名称说明

| 运营商  | `account.isp` |
|:----:|:-------------:|
| 校园网  |     `ncu`     |
| 中国移动 |    `cmcc`     |
| 中国联通 |   `unicom`    |
| 中国电信 |   `ndcard`    |

### 使用环境变量覆盖配置

- 环境变量开头统一为 `SRUN`。

- 环境变量中节点间的分隔符为 `_`。

- 例子：

|        配置名         |          环境变量           |
|:------------------:|:-----------------------:|
|   `account.isp`    |   `SRUN_ACCOUNT_ISP`    |
| `account.username` | `SRUN_ACCOUNT_USERNAME` |
| `account.password` | `SRUN_ACCOUNT_PASSWORD` |

### 其他配置文件路径

优先级从高到底：

- `<程序所在路径>`

- `<用户主目录>/.config/srunlogin`

- `<用户主目录>/.srunlogin`
