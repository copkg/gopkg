package errors

type ErrorCode int

const (
	ErrNotFound ErrorCode = iota + 110001
	ErrAlreadyExist
	ErrReachMaxCount
	ErrSecretNotFound
	ErrSuccess
	ErrUnknown
	ErrBind
	ErrValidation
	ErrTokenInvalid
	ErrDatabase
	ErrEncrypt
	ErrSignatureInvalid
	ErrExpired
	ErrInvalidAuthHeader
	ErrMissingHeader
	ErrPasswordIncorrect
	ErrPermissionDenied
	ErrEncodingFailed
	ErrDecodingFailed
	ErrInvalidJSON
	ErrEncodingJSON
	ErrDecodingJSON
	ErrInvalidYaml
	ErrEncodingYaml
	ErrDecodingYaml
	ErrInitWeComSDK
	ErrThirdAPI
	ErrAPPDisabled
	ErrOAuthScope
	NotSupportedOAuth
	ErrFileFailed
)

// Error 描述错误结构体
type Error struct {
	Code ErrorCode `json:"code,omitempty"`
	Msg  string    `json:"msg,omitempty"`
	Err  error     `json:"-"`
}

// Error 实现 error 接口
func (e Error) Error() string {
	return e.Msg
}
func (e Error) Wrap() error {
	return e.Err
}
func (e Error) HttpCode() ErrorCode {
	return e.Code
}

// New 创建一个新的错误实例
func New(code ErrorCode, message string, err error) *Error {
	return &Error{
		Code: code,
		Msg:  message,
		Err:  err,
	}
}

// 定义常见的错误
//var (
//	ErrNotFoundError          = New(ErrNotFound, "Resource not found")
//	ErrAlreadyExistError      = New(ErrAlreadyExist, "already exist")
//	ErrReachMaxCountError     = New(ErrReachMaxCount, "Secret reach the max count")
//	ErrSecretNotFoundError    = New(ErrSecretNotFound, "Secret not found")
//	ErrSuccessError           = New(ErrSuccess, "OK")
//	ErrUnknownError           = New(ErrUnknown, "Internal server error")
//	ErrBindError              = New(ErrBind, "Error occurred while binding the request body to the struct")
//	ErrValidationError        = New(ErrValidation, "Validation failed")
//	ErrTokenInvalidError      = New(ErrTokenInvalid, "Token invalid")
//	ErrDatabaseError          = New(ErrDatabase, "Database error")
//	ErrEncryptError           = New(ErrEncrypt, "Error occurred while encrypting the user password")
//	ErrSignatureInvalidError  = New(ErrSignatureInvalid, "Signature is invalid")
//	ErrExpiredError           = New(ErrExpired, "Token expired")
//	ErrInvalidAuthHeaderError = New(ErrInvalidAuthHeader, "Invalid authorization header")
//	ErrMissingHeaderError     = New(ErrMissingHeader, "The `Authorization` header was empty")
//	ErrPasswordIncorrectError = New(ErrPasswordIncorrect, "Password was incorrect")
//	ErrPermissionDeniedError  = New(ErrPermissionDenied, "SDKermission denied")
//	ErrEncodingFailedError    = New(ErrEncodingFailed, "Encoding failed due to an error with the data")
//	ErrDecodingFailedError    = New(ErrDecodingFailed, "Decoding failed due to an error with the data")
//	ErrInvalidJSONError       = New(ErrInvalidJSON, "Data is not valid JSON")
//	ErrEncodingJSONError      = New(ErrEncodingJSON, "JSON data could not be encoded")
//	ErrDecodingJSONError      = New(ErrDecodingJSON, "JSON data could not be decoded")
//	ErrInvalidYamlError       = New(ErrInvalidYaml, "Data is not valid Yaml")
//	ErrEncodingYamlError      = New(ErrEncodingYaml, "Yaml data could not be encoded")
//	ErrDecodingYamlError      = New(ErrDecodingYaml, "Yaml data could not be decoded")
//	ErrInitWeComSDKError      = New(ErrInitWeComSDK, "init wecom SDK fail")
//	ErrAPPDisabledError       = New(ErrAPPDisabled, "the applicatin disabled")
//	ErrOAuthScopeError        = New(ErrAPPDisabled, "oauth2.0 scope error,only supporte snsapi_privateinfo")
//	ErrNotSupportedOAuthError = New(NotSupportedOAuth, "not supported oauth2.0 scope error,only supporte official_account，wecom")
//	ErrFileError              = New(ErrFileFailed, "read file error")
//)
