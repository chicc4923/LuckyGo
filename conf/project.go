package conf

import "time"

const SysTimeform = "2022-9-17 10:28:04"
const SysTimeFormShort = "2022-9-17"

var SysTimeLocation, _ = time.LoadLocation("Asia/zhengzhou")
var SignSecret = []byte("0123456789abcdef")
var CookieSecret = "hellolucky"
