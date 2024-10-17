package errors

import "fmt"

type ErrorCode int

const (
	ErrNotFound ErrorCode = iota + 110001
	ErrUserAlreadyExist
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
	ErrorExpired
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
)

// Error 描述错误结构体
type Error struct {
	Code ErrorCode
	Msg  string
}

// Error 实现 error 接口
func (e Error) Error() string {
	return fmt.Sprintf(`{"code": %d, "msg": "%s"}`, e.Code, e.Msg)
}

// New 创建一个新的错误实例
func New(code ErrorCode, message string) Error {
	return Error{
		Code: code,
		Msg:  message,
	}
}

// 定义常见的错误
var (
	ErrNotFoundError          = New(ErrNotFound, "Resource not found")
	ErrUserAlreadyExistError  = New(ErrUserAlreadyExist, "User already exist")
	ErrReachMaxCountError     = New(ErrReachMaxCount, "Secret reach the max count")
	ErrSecretNotFoundError    = New(ErrSecretNotFound, "Secret not found")
	ErrSuccessError           = New(ErrSuccess, "OK")
	ErrUnknownError           = New(ErrUnknown, "Internal server error")
	ErrBindError              = New(ErrBind, "Error occurred while binding the request body to the struct")
	ErrValidationError        = New(ErrValidation, "Validation failed")
	ErrTokenInvalidError      = New(ErrTokenInvalid, "Token invalid")
	ErrDatabaseError          = New(ErrDatabase, "Database error")
	ErrEncryptError           = New(ErrEncrypt, "Error occurred while encrypting the user password")
	ErrSignatureInvalidError  = New(ErrSignatureInvalid, "Signature is invalid")
	ErrExpiredError           = New(ErrExpired, "Token expired")
	ErrInvalidAuthHeaderError = New(ErrInvalidAuthHeader, "Invalid authorization header")
	ErrMissingHeaderError     = New(ErrMissingHeader, "The `Authorization` header was empty")
	ErrPasswordIncorrectError = New(ErrPasswordIncorrect, "Password was incorrect")
	ErrPermissionDeniedError  = New(ErrPermissionDenied, "Permission denied")
	ErrEncodingFailedError    = New(ErrEncodingFailed, "Encoding failed due to an error with the data")
	ErrDecodingFailedError    = New(ErrDecodingFailed, "Decoding failed due to an error with the data")
	ErrInvalidJSONError       = New(ErrInvalidJSON, "Data is not valid JSON")
	ErrEncodingJSONError      = New(ErrEncodingJSON, "JSON data could not be encoded")
	ErrDecodingJSONError      = New(ErrDecodingJSON, "JSON data could not be decoded")
	ErrInvalidYamlError       = New(ErrInvalidYaml, "Data is not valid Yaml")
	ErrEncodingYamlError      = New(ErrEncodingYaml, "Yaml data could not be encoded")
	ErrDecodingYamlError      = New(ErrDecodingYaml, "Yaml data could not be decoded")
	ErrInitWeComSDKError      = New(ErrInitWeComSDK, "init wecom SDK fail")
)
