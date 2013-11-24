package lib

//公共方法
import "crypto/md5"
import "encoding/hex"
import "net"
import "runtime"
import "strings"
import "github.com/robfig/revel"

//返回MD5加密
func EncryptPassword(s string) string {
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

//获取系统信息
func GetSysInfo() map[string]string {

	sysinfo := make(map[string]string)
	sysinfo["GO_OS"] = runtime.GOOS
	sysinfo["SERVER_SOFTWARE"] = runtime.Version()
	sysinfo["MYSQL_VERSION"] = "MySQL 5.7"

	return sysinfo
}
