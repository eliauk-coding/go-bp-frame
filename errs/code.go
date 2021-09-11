package errs

type ErrCode int

const (
	Success ErrCode = 0x00

	// common errors
	ErrResNotFound    ErrCode = 0x1000 // resource not found
	ErrAlreadyExist   ErrCode = 0x1001 // resource already exist
	ErrOpNotAllowed   ErrCode = 0x1002 // operation not allowed
	ErrResNotComplete ErrCode = 0x1003 // resource not complete

	// auth errors
	ErrAuthNoLic             ErrCode = 0x1010 // no valid license found
	ErrAuthReject            ErrCode = 0x1011 // action rejected, mostly caused by authorization
	ErrAuthLicLimit          ErrCode = 0x1012 // max online client reached
	ErrAuthTokenInvalid      ErrCode = 0x1013 // token is invalid
	ErrAuthTokenExpired      ErrCode = 0x1014 // token is out of time
	ErrAuthTokenCreateFailed ErrCode = 0x1015 // token create failed
	ErrAuthNoPermissionEdit  ErrCode = 0x1016 // no permission to edit

	// user errors
	ErrUserName         ErrCode = 0x1100 // username is incorrect
	ErrUserPassword     ErrCode = 0x1101 // password is incorrect
	ErrUserNameOrPasswd ErrCode = 0x1102 // username or password is not match

	// request errors
	ErrReqParamMissing ErrCode = 0x1200 // required parameter(s) is(are) missing
	ErrReqParamInvalid ErrCode = 0x1201 // parameter(s) is(are) invalid

	// data errors
	ErrIndexOutOfRange ErrCode = 0x1300 // index out of range
	ErrTypeMotMatch    ErrCode = 0x1301 // type does not match
	ErrDataNotExist    ErrCode = 0x1302 // data does not exist

	// http request errors
	ErrNewRequest   ErrCode = 0x1400 // new request error
	ErrHttpResp     ErrCode = 0x1401 // HTTP response error
	ErrHttpRespBody ErrCode = 0x1402 // HTTP response body error

	// file errors
	ErrUploadFile ErrCode = 0x1500 // upload file error
	ErrStatFile   ErrCode = 0x1501 // stat file error
	ErrReadFile   ErrCode = 0x1502 // read file error
	ErrWriteFile  ErrCode = 0x1503 // write file error
	ErrNewFolder  ErrCode = 0x1504 // new folder error
	ErrNewFile    ErrCode = 0x1505 // new file error
	ErrFileType   ErrCode = 0x1506 // file type error
	ErrSaveFile   ErrCode = 0x1507 // save file error
	ErrZipFile    ErrCode = 0x1508 // zip file error
	ErrUnZipFile  ErrCode = 0x1509 // unzip file error
	ErrBrdFile    ErrCode = 0x1510 // brd file error

	// server errors
	ErrSvrInternal      ErrCode = 0x1F00 // server internal error
	ErrSvrSqlExecFailed ErrCode = 0x1F01 // database operation error
)

var errMsgMap = map[ErrCode]string{
	Success: "成功",

	// common errors
	ErrResNotFound:    "资源不存在",
	ErrAlreadyExist:   "资源已经存在",
	ErrOpNotAllowed:   "操作不允许",
	ErrResNotComplete: "资源数据不完整",

	// auth errors
	ErrAuthNoLic:             "无效许可证",
	ErrAuthReject:            "未授权",
	ErrAuthLicLimit:          "已达到最大在线客户端数量",
	ErrAuthTokenInvalid:      "令牌无效",
	ErrAuthTokenExpired:      "令牌过期",
	ErrAuthTokenCreateFailed: "创建令牌失败",
	ErrAuthNoPermissionEdit:  "没有编辑权限",

	// user errors
	ErrUserName:         "用户名不正确",
	ErrUserPassword:     "密码不正确",
	ErrUserNameOrPasswd: "用户名或密码不匹配",

	// data errors
	ErrIndexOutOfRange: "索引超出范围",
	ErrTypeMotMatch:    "类型不匹配",
	ErrDataNotExist:    "数据不存在",

	// request errors
	ErrReqParamMissing: "缺少必需参数",
	ErrReqParamInvalid: "参数无效",

	// new request errors
	ErrNewRequest:   "创建HTTP请求错误",
	ErrHttpResp:     "http响应错误",
	ErrHttpRespBody: "http响应正文错误",

	// file errors
	ErrUploadFile: "文件上传错误",
	ErrStatFile:   "文件打开错误",
	ErrReadFile:   "读取文件错误",
	ErrWriteFile:  "写入文件错误",
	ErrNewFolder:  "新建文件夹错误",
	ErrNewFile:    "新建文件错误",
	ErrFileType:   "文件类型错误",
	ErrSaveFile:   "保存文件错误",
	ErrZipFile:    "压缩文件错误",
	ErrUnZipFile:  "解压缩文件错误",
	ErrBrdFile:    "操作brd文件错误",

	// server errors
	ErrSvrInternal:      "服务器内部错误",
	ErrSvrSqlExecFailed: "数据库操作错误",
}
