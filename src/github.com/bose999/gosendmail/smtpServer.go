package gosendmail

import (

)

type SmtpServer struct {
	serverName string
	port       int
	connectServer string
	userName   string
	password   string
}