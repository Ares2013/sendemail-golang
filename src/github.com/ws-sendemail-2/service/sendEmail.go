package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/ws-sendemail-2/defs"
	"strings"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

/**
检查发送邮件的实体设置
@author Ares Peng(pengpp@wondershare.cn)
@Param email
**/
func CheckEmail(email *defs.Email) (Email *defs.Email, err error) {
	if !govalidator.IsEmail(email.From) {
		err := "From fromat is error"
		return email, &errorString{err}
	}
	if len(strings.TrimSpace(email.Subject)) == 0 {
		err := "Subject is empty"
		return email, &errorString{err}
	}
	if len(strings.TrimSpace(email.Body)) == 0 {
		err := "Body is empty"
		return email, &errorString{err}
	}
	if email.HasAttach && len(strings.TrimSpace(email.Attachment)) == 0 {
		err := "Set hasAttach ,Attchment is must set"
		return email, &errorString{err}
	}
	return email, nil
}
func SetMail(email *defs.Email) (sid int, msg error) {
	if _, err := CheckEmail(email); err == nil {
		return 1, nil
	} else {
		return 0, err
	}
	var err = "Not to do anything"
	return 0, &errorString{err}
}
