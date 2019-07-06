# sendmail

[![Build Status](https://www.travis-ci.org/zs5460/sendmail.svg?branch=master)](https://www.travis-ci.org/zs5460/sendmail)

sendmail是一个发送邮件的命令行小工具。

## 下载

[最新版](https://github.com/zs5460/sendmail/releases/latest)

## 配置

将config.json.sample改名为config.json，填入相关信息。

```javascript
{
    "MailSubject": "alarm mail", // 邮件主题，如报警邮件
    "MailServer": "smtp.exmail.qq.com:465", // 发送邮件服务器地址及端口，支持25|465|587
    "MailSender": "alarm@test.cn", // 发件人邮箱帐号
    "MailSenderPwd": "123456", // 发件人邮箱密码
    "MailReciver":"tech@test.cn;yunwei@test.cn" // 收件人，多个收件人用;分隔
}
```

## 使用

```shell
sendmail hello
```

## Licence

Released under MIT license, see [LICENSE](LICENSE) for details.
