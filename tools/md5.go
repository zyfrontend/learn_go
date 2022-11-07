package tools

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"
)

func Md5(address string) string {
	timeStr := time.Now().Format("20060102")
	md5Data := address + "." + timeStr
	w := md5.New()
	io.WriteString(w, md5Data)
	bw := w.Sum(nil)
	verifySign := hex.EncodeToString(bw)[0:10]
	return verifySign
}
