# clean weibo

#### 介绍
一键清空个人名下所有微博

#### 使用说明
1. cd app;cp env.example .env;
2. 编辑.env
```ini
UID=xxx
COOKIE=login_sid_t=xxxx
X_XSRF_TOKEN=xxx
```
微博网页版获取到的网络参数
3. go run clean-wb.go
   或者编译为对应平台的binary，在app目录下跑./clean-wb

#### 已知问题
1. 脚本删除较多微博时会被微博拦截请求。可以停止脚本，等待几分钟后重新跑。

