package model

import (
	"crypto/md5"
	"encoding/hex"
)

// Encripitando a senha usando o MD5
// ud --> user domain
// função anônima
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))

}
