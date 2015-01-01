GoSendMail
======

GoSendMail is simple smtp sender.
encoding utf8 only.


use sample:

    // smtpServer info : IP Address, port, userName, Password
    smtpServer := gosendmail.NewSmtpSever("xxx.xxx.xxx.xxx", 25, "yyy", "zzz")
    
    // sendmail info : from address, to address, cc address, subject, body
    sendMail := gosendmail.NewSendMail("abc@xxx.xxx", []string{"abc@xxx.xxx"},
      []string{"xyz@gxxx.xxx"}, "UTF8のSubject", "UTF8の本文")

    // send smtp mail
    err := gosendmail.SendSmtp(smtpServer, sendMail)
    if err != nil {
        log.Fatal(err)
         os.Exit(1)
    }


install:
    % go get github.com/bose999/gosendmail