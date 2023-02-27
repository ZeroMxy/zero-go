package cipher

import (
	"zero-go/framework/log"
	"golang.org/x/crypto/bcrypt"
)

type Cipher struct {}

// 加密
func Encrypt (param string) string {

	var param_byte = []byte(param)
	
	var param_hash, hash_err = bcrypt.GenerateFromPassword(param_byte, bcrypt.MinCost)

	if (hash_err != nil) {
		log.Error(hash_err)
	}
	
	return string(param_hash)

}

// 验证
func Verify (param_encrypted, param string) bool {

	var param_encrypted_byte = []byte(param_encrypted)
	var param_byte 			 = []byte(param)

	var hash_err = bcrypt.CompareHashAndPassword(param_encrypted_byte, param_byte)

	if (hash_err != nil) {
		return false
	}

	return true

}