package errmsg

// code 200 500 xxx
const (
	SUCC  = 200
	ERROR = 500

	// code = 1xxx表示用户模块的错误 2xxx表示文章模块的错误 3xxx表示分类模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PWD_WRONG        = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_EXPIRE     = 1005
	ERRPR_TOKEN_INVALID    = 1006
	ERRPR_TOKEN_TYPE_WRONG = 1007
)

var codeMsg = map[int]string{
	SUCC:  "OK",
	ERROR: "FAIL",

	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PWD_WRONG:        "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_EXPIRE:     "TOKEN已过期",
	ERRPR_TOKEN_INVALID:    "TOKEN不正确",
	ERRPR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

func GetErrmsg(code int) string {

	return codeMsg[code]
}
