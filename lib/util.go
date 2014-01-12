package lib

//公共方法
import "crypto/md5"
import "encoding/hex"
import "net"
import "strings"
import "github.com/robfig/revel"

//返回MD5加密
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//获取客户端IP
func GetClientIP() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		revel.WARN.Println(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

//元素是否包含在数组中
func In_Array(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
