GoSendMail
======

GoSendMail is simple smtp sender.
encoding utf8 only.


use sample:

smtpServer := gosendmail.NewSmtpSever("xxx.xxx.xxx.xxx", 25, "yyy", "zzz")
sendMail := gosendmail.NewSendMail("abc@xxx.xxx", []string{"abc@xxx.xxx"},
[]string{"xyz@gxxx.xxx"}, "UTF8のSubject", "UTF8の本文")
err := gosendmail.SendSmtp(smtpServer, sendMail)
if err != nil {
    log.Fatal(err)
     os.Exit(1)
}

