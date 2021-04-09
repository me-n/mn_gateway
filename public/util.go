package public

import (
	"crypto/sha256"
	"fmt"
)

func PasswordAddSalt(salt,password string) string {
	hs1 := sha256.New()
	hs1.Write([]byte(password))
	str := fmt.Sprintf("%x", hs1.Sum(nil))
	hs2 := sha256.New()
	hs2.Write([]byte(str+salt))
	return fmt.Sprintf("%x",hs2.Sum(nil))
}