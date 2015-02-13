package tool

import (
	"github.com/nu7hatch/gouuid"
	"encoding/hex"
	"fmt"
)

func Uuid() string{
	u, err := uuid.NewV4()
	if err != nil {
	    fmt.Println("Error:", err)
	    return ""
	}
	uid := hex.EncodeToString(u[:])
	return uid
}