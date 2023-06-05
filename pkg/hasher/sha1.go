package hasher

import (
	"crypto/sha1"
	"fmt"
	"github.com/Onelvay/HL-architecture/internal/domain/user"
)

const (
	hashSalt = "alenova"
)

func Password(user *user.User) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(hashSalt))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))
}
