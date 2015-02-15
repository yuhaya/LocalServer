package tool

import (
	"encoding/hex"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"regexp"
)

const PHONE_RGX = `^1([38][0-9]|14[57]|5[^4])\d{8}$`

func Uuid() string {
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	uid := hex.EncodeToString(u[:])
	return uid
}

func IsPhone(phone string) bool {
	rgx := regexp.MustCompile(PHONE_RGX)
	return rgx.MatchString(phone)
}
