package main

import (
	"fmt"
)

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(incommingCode int) error {
	if s.code != incommingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
