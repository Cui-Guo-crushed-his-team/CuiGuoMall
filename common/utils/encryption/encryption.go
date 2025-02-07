package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(data string) string {
	//data: user_id+user_ip
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
