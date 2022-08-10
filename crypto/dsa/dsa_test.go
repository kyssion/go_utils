package dsa

import (
	"fmt"
	"testing"
)

func TestDsa(t *testing.T) {
	msg := "this is message"
	t.Run("私钥加密解密", func(t *testing.T) {
		// 私钥加密		t.Logf("sign : %s ,  error info : %v", string(ans), err)

		ans, err := sign([]byte(msg), "privkey.pem")
		// 公钥解密
		status, err := verify([]byte(msg), "pubkey.pem", ans)
		fmt.Printf("status : %v ,  err : %v", status, err)
	})
}
