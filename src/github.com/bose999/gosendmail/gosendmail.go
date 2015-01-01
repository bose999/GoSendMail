// Package gosendmail can send utf8 mail. 
package gosendmail

import (
	"net/smtp"
	"strconv"
)

// creat new SmtpServer Instance
func NewSmtpSever(serverName string, port int, userName string, password string) SmtpServer {
	smtpServer := new(SmtpServer)
	smtpServer.serverName = serverName
	smtpServer.port = port
	smtpServer.connectServer = smtpServer.serverName + ":" + strconv.Itoa(smtpServer.port)
	smtpServer.userName = userName
	smtpServer.password = password
	return *smtpServer
}

// create new SendMail Instance
func NewSendMail(from string, to []string, cc []string, subject string, body string) SendMail {
	sendmail := new(SendMail)
	sendmail.from = from
	sendmail.to = to
	sendmail.cc = cc
	sendmail.receivers = sendmail.makeReceivers(to,cc)
	sendmail.subject = sendmail.encodeMIMESubject(subject)
	sendmail.body = sendmail.encodeBase64Body(body)
	return *sendmail
}

// send utf8 smtp mail
func SendSmtp(smtpServer SmtpServer, sendmail SendMail) error {
	auth := smtp.PlainAuth("", smtpServer.userName, smtpServer.password, smtpServer.serverName)
	err := smtp.SendMail(smtpServer.connectServer, auth, sendmail.from, sendmail.receivers, sendmail.makeSendString())
	return err
}
