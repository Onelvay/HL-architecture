package hasher

import (
	"crypto/sha1"
	"fmt"
	"github.com/Onelvay/HL-architecture/internal/entity"
)

const (
	hashSalt = "alenova"
)

func Password(user *entity.User) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(hashSalt))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))
}
