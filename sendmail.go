package gosendmail

import (
	"bytes"
	"encoding/base64"
	"strings"
)

// mail value and build method
type SendMail struct {
	from      string
	to        []string
	cc        []string
	receivers []string
	subject   string
	body      string
}

// when send, make Strings
func (sendSendMail SendMail) makeSendString() []byte {
	bufferSring := bytes.NewBufferString("From:")
	bufferSring.WriteString(sendSendMail.from)
	bufferSring.WriteString("\r\n")
	for _, address := range sendSendMail.to {
		bufferSring.WriteString("To:")
		bufferSring.WriteString(address)
		bufferSring.WriteString("\r\n")
	}
	for _, address := range sendSendMail.cc {
		bufferSring.WriteString("Cc:")
		bufferSring.WriteString(address)
		bufferSring.WriteString("\r\n")
	}
	bufferSring.WriteString(sendSendMail.subject)
	bufferSring.WriteString("MIME-Version: 1.0\r\n")
	bufferSring.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	bufferSring.WriteString("Content-Transfer-Encoding: base64\r\n")
	bufferSring.WriteString("\r\n")
	bufferSring.WriteString(sendSendMail.body)
	return bufferSring.Bytes()
}

// Body Base64 Encode
func (sendSendMail SendMail) encodeBase64Body(body string) string {
	bytesBody := []byte(body)
	return sendSendMail.insertCrlf76(base64.StdEncoding.EncodeToString(bytesBody))
}

// Subject MIME Encode
func (sendSendMail SendMail) encodeMIMESubject(subject string) string {
	buffer := bytes.NewBufferString("Subject:")
	for _, splitString := range sendSendMail.splitUtf8String(subject, 13) {
		buffer.WriteString(" =?utf-8?B?")
		buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(splitString)))
		buffer.WriteString("?=\r\n")
	}
	return buffer.String()
}

// split UTF8 String
func (sendSendMail SendMail) splitUtf8String(utf8string string, length int) []string {
	resultString := []string{}
	buffer := bytes.NewBufferString("")
	for i, char := range strings.Split(utf8string, "") {
		buffer.WriteString(char)
		if (i+1)%length == 0 {
			resultString = append(resultString, buffer.String())
			buffer.Reset()
		}
	}
	if buffer.Len() > 0 {
		resultString = append(resultString, buffer.String())
	}
	return resultString
}

//  CRLF every 76 byte.
func (sendSendMail SendMail) insertCrlf76(msg string) string {
	buffer := bytes.NewBufferString("")
	for i, char := range strings.Split(msg, "") {
		buffer.WriteString(char)
		if (i+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
	return buffer.String()
}

// make receivers value
func (sendSendMail SendMail) makeReceivers(to []string, cc []string) []string {
	var receivers []string
	for _, address := range to {
		receivers = append(receivers, address)
	}
	for _, address := range cc {
		receivers = append(receivers, address)
	}
	return receivers
}
