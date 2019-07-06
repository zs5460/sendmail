package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zs5460/mail"
	"github.com/zs5460/my"
)

func main() {
	var cfg mail.Config
	my.MustLoadConfig("config.json", &cfg)

	body := fmt.Sprint(strings.Join(os.Args[1:], "<br>\n"))

	err := mail.SendMail(
		cfg.MailSender,
		cfg.MailSenderPwd,
		cfg.MailServer,
		cfg.MailReciver,
		cfg.MailSubject,
		body,
	)
	if err != nil {
		log.Println(err)
	}

}
