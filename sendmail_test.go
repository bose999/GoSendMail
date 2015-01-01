package gosendmail

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

// Test makeSendString 
func TestMakeSendString(t *testing.T) {
	Describe(t, "SendMail.makeSendString", func() {
		Context("A発注の件を変更する場合", func() {
			sendmail := NewSendMail("from", []string{"to1","to2"}, []string{"cc1","cc2"}, "subject", "body")
			retrunString := string(sendmail.makeSendString())
			It("期待値が返ってくる", func() {
				Expect(retrunString).To(Equal, "From:from\r\nTo:to1\r\nTo:to2\r\nCc:cc1\r\nCc:cc2\r\nSubject: =?utf-8?B?c3ViamVjdA==?=\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"utf-8\"\r\nContent-Transfer-Encoding: base64\r\n\r\nYm9keQ==")
			})
		})
			})
}

// Test encodeBase64Body
func TestEncodeBase64Body(t *testing.T) {
	Describe(t, "SendMail.encodeMIMESubject", func() {
		Context("sendmailインスタンスから実行する", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnStrings := sendmail.encodeBase64Body("お世話になっております。bose999です。\r\n以上、よろしくお願いします。\r\n")
			It("各フィールドの値から生成される期待値が返ってくる", func() {
				Expect(returnStrings).To(Equal, "44GK5LiW6Kmx44Gr44Gq44Gj44Gm44GK44KK44G+44GZ44CCYm9zZTk5OeOBp+OBmeOAgg0K5Lul\r\n5LiK44CB44KI44KN44GX44GP44GK6aGY44GE44GX44G+44GZ44CCDQo=")
			})
		})
			})
}

// Test encodeMIMESubject
func TestEncodeMIMESubject(t *testing.T) {
	Describe(t, "SendMail.encodeMIMESubject", func() {
		Context("A発注の件を変更する場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnStrings := sendmail.encodeMIMESubject("A発注の件を変更する場合")
			It("期待値が返ってくる", func() {
				Expect(returnStrings).To(Equal, "Subject: =?utf-8?B?QeeZuuazqOOBruS7tuOCkuWkieabtOOBmeOCi+WgtOWQiA==?=\r\n")
			})
		})
		Context("株式会社ＸＸＸＸＸＸＸＸＸ発注の件を変更する場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnStrings := sendmail.encodeMIMESubject("株式会社ＸＸＸＸＸＸＸＸＸ発注の件")
			It("期待値が返ってくる", func() {
				Expect(returnStrings).To(Equal, "Subject: =?utf-8?B?5qCq5byP5Lya56S+77y477y477y477y477y477y477y477y477y4?=\r\n =?utf-8?B?55m65rOo44Gu5Lu2?=\r\n")
			})
		})
		Context("株式会社ＸＸbＸＸＸaＸＸＸＸ発注の件", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnStrings := sendmail.encodeMIMESubject("株式会社ＸＸbＸＸＸaＸＸＸＸ発注の件")
			It("期待値が返ってくる", func() {
				Expect(returnStrings).To(Equal, "Subject: =?utf-8?B?5qCq5byP5Lya56S+77y477y4Yu+8uO+8uO+8uGHvvLjvvLg=?=\r\n =?utf-8?B?77y477y455m65rOo44Gu5Lu2?=\r\n")
			})
		})
	})
}

// Test sendMailSplitUtf8String
func TestSendMailSplitUtf8String(t *testing.T) {
	Describe(t, "SendMail.splitUtf8String", func() {
		Context("空文字を3で分割する場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnStrings := sendmail.splitUtf8String("", 3)
			It("[]string{ }が返ってくるので配列数は0になる", func() {
				Expect(len(returnStrings)).To(Equal, 0)
			})
		})
		Context("文字列12を3で分割する場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnStrings := sendmail.splitUtf8String("12", 3)
			It("[]string{12}が返ってくる", func() {
				Expect(returnStrings).To(Equal, []string{"12"})
			})
		})
		Context("文字列123を3で分割する場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnStrings := sendmail.splitUtf8String("123", 3)
			It("[]string{123}が返ってくる", func() {
				Expect(returnStrings).To(Equal, []string{"123"})
			})
		})
		Context("文字列1234を3で分割する場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnStrings := sendmail.splitUtf8String("1234", 3)
			It("[]string{123,4}が返ってくる", func() {
				Expect(returnStrings).To(Equal, []string{"123", "4"})
			})
		})
		Context("文字列1234567890を3で分割する場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnStrings := sendmail.splitUtf8String("1234567890", 3)
			It("[]string{123,456,789,0}が返ってくる", func() {
				Expect(returnStrings).To(Equal, []string{"123", "456", "789", "0"})
			})
		})
	})
}

// Test sendMailInsertCrlf76
func TestSendMailInsertCrlf76(t *testing.T) {
	Describe(t, "SendMail.insertCflf76", func() {
		Context("文字列123456789012345678901234567890123456789012345678901234567890123456789012345を渡した場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnString := sendmail.insertCrlf76("123456789012345678901234567890123456789012345678901234567890123456789012345")
			It("返り値は123456789012345678901234567890123456789012345678901234567890123456789012345が返ってくる", func() {
				Expect(returnString).To(Equal, "123456789012345678901234567890123456789012345678901234567890123456789012345")
			})
		})
		Context("文字列1234567890123456789012345678901234567890123456789012345678901234567890123456を渡した場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnString := sendmail.insertCrlf76("1234567890123456789012345678901234567890123456789012345678901234567890123456")
			It("返り値は1234567890123456789012345678901234567890123456789012345678901234567890123456CFLFが返ってくる", func() {
				Expect(returnString).To(Equal, "1234567890123456789012345678901234567890123456789012345678901234567890123456\r\n")
			})
		})
		Context("文字列12345678901234567890123456789012345678901234567890123456789012345678901234567を渡した場合", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnString := sendmail.insertCrlf76("12345678901234567890123456789012345678901234567890123456789012345678901234567")
			It("返り値は1234567890123456789012345678901234567890123456789012345678901234567890123456CRLF7が返って来る", func() {
				Expect(returnString).To(Equal, "1234567890123456789012345678901234567890123456789012345678901234567890123456\r\n7")
			})
		})
		Context("長い文字列を渡した場合でも", func() {
			sendmail := NewSendMail("", []string{""}, []string{""}, "", "")
			returnString := sendmail.insertCrlf76("1234567890123456789012345678901234567890123456789012345678901234567890123456123456789012345678901234567890123456789012345678901234567890123456789012345612345678901234567890123456789012345678901234567890123456789012345678901234567")
			It("期待値が返ってくる", func() {
				Expect(returnString).To(Equal, "1234567890123456789012345678901234567890123456789012345678901234567890123456\r\n1234567890123456789012345678901234567890123456789012345678901234567890123456\r\n1234567890123456789012345678901234567890123456789012345678901234567890123456\r\n7")
			})
		})
	})
}

// Test makeReceivers
func TestMakeReceivers(t *testing.T) {
	Describe(t, "SendMail.makeReceivers", func() {
		Context("to []stringとcc []stringからrecives []stringを作成する", func() {
			sendmail := NewSendMail("", []string{"aaa", "bbb"}, []string{"ccc", "ddd"}, "", "")
			reciveres := sendmail.makeReceivers(sendmail.to, sendmail.cc)
			It("返り値はtoとccを連結した配列が返ってくる", func() {
				Expect(reciveres[0]).To(Equal, "aaa")
				Expect(reciveres[1]).To(Equal, "bbb")
				Expect(reciveres[2]).To(Equal, "ccc")
				Expect(reciveres[3]).To(Equal, "ddd")
			})
		})
	})
}
