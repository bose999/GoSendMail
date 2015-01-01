package gosendmail

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

// Test NewSmtpServer
func TestNewSmtpSever(t *testing.T) {
	Describe(t, "NewSmtpSever", func() {
		Context("インスタンスが正しい値を保持しているか確認する", func() {
			smtpServer := NewSmtpSever("serverName", 25, "userName", "password")
			It("期待値が返ってくる", func() {
				Expect(smtpServer.serverName).To(Equal, "serverName")
				Expect(smtpServer.port).To(Equal, 25)
				Expect(smtpServer.connectServer).To(Equal, "serverName:25")
				Expect(smtpServer.userName).To(Equal, "userName")
				Expect(smtpServer.password).To(Equal, "password")
			})
		})
	})
}

// Test NewSendMail
func TestNewSendMail(t *testing.T) {
	Describe(t, "NewSendMail", func() {
		Context("インスタンスが正しい値を保持しているか確認する ", func() {
			sendmail := NewSendMail("from@bose999.com", []string{"to1@bose999.com", "to2@bose999.com"}, []string{"cc1@bose999.com", "cc2@bose999.com"}, "subject", "body")
			It("期待値が返ってくる", func() {
				Expect(sendmail.from).To(Equal, "from@bose999.com")
				Expect(sendmail.to[0]).To(Equal, "to1@bose999.com")
				Expect(sendmail.to[1]).To(Equal, "to2@bose999.com")
				Expect(sendmail.cc[0]).To(Equal, "cc1@bose999.com")
				Expect(sendmail.cc[1]).To(Equal, "cc2@bose999.com")
				Expect(sendmail.receivers[0]).To(Equal, "to1@bose999.com")
				Expect(sendmail.receivers[1]).To(Equal, "to2@bose999.com")
				Expect(sendmail.receivers[2]).To(Equal, "cc1@bose999.com")
				Expect(sendmail.receivers[3]).To(Equal, "cc2@bose999.com")
				Expect(sendmail.subject).To(Equal, "Subject: =?utf-8?B?c3ViamVjdA==?=\r\n")
				Expect(sendmail.body).To(Equal, "Ym9keQ==")
			})
		})
	})
}

// Test SendSmtp
func TestSendSmtp(t *testing.T) {
	Describe(t, "SendSmtp", func() {
		Context("エラーを起こしてエラーが返ってくることを確認する ", func() {
			smtpServer := NewSmtpSever("serverName", 25, "userName", "password")
			sendmail := NewSendMail("from@bose999.com", []string{"to1@bose999.com", "to2@bose999.com"}, []string{"cc1@bose999.com", "cc2@bose999.com"}, "subject", "body")
			err := SendSmtp(smtpServer, sendmail)
			It("接続が出来ないのでエラーが起こる", func() {
				Expect(err).To(Exist)
			})
		})
	})
}
