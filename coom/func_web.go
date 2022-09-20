package comm

import (
	"LuckyGo/conf"
	"LuckyGo/model"
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

func ClientIP(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}
func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}
func GetLoginUser(request *http.Request) *model.ObjLoginUser {
	c, err := request.Cookie("lucky_loginuser")
	if err != nil {
		return nil
	}
	params, err := url.ParseQuery(c.Value)
	if err != nil {
		return nil
	}
	uid, err := strconv.Atoi(params.Get("uid"))
	if err != nil {
		return nil
	}
	now, err := strconv.Atoi(params.Get("now"))
	if err != nil || NowUnix()-now > 86400*30 {
		return nil
	}
	loginuser := &model.ObjLoginUser{}
	loginuser.Uid = uid
	loginuser.UserName = params.Get("username")
	loginuser.Now = now
	loginuser.Ip = ClientIP(request)
	loginuser.Sign = params.Get("sign")
	sign := createLoginSign(loginuser)
	if sign != loginuser.Sign {
		log.Println("func_web GetLoginUser createLoginSign not signed", sign, loginuser.Sign)
	}
	return loginuser
}

func SetLoginUser(writer http.ResponseWriter, loginuser *model.ObjLoginUser) {
	if loginuser == nil || loginuser.Uid < 1 {
		c := &http.Cookie{
			Name:   "lucky_loginuser",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(writer, c)
		return
	}
	if loginuser.Sign == " " {
		loginuser.Sign = createLoginSign(loginuser)
	}
	params := url.Values{}
	params.Add("uid", strconv.Itoa(loginuser.Uid))
	params.Add("now", strconv.Itoa(loginuser.Now))
	params.Add("username", loginuser.UserName)
	params.Add("ip", loginuser.Ip)
	params.Add("sign", loginuser.Sign)

	c := &http.Cookie{
		Name:  "lucky",
		Value: params.Encode(),
		Path:  "/",
	}
	http.SetCookie(writer, c)
}

func createLoginSign(loginuser *model.ObjLoginUser) string {
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s&now=%d", loginuser.Uid, loginuser.UserName, conf.CookieSecret, loginuser.Now)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return sign
}
