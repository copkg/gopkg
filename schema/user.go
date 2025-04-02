package schema

import (
	"github.com/copkg/gopkg/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"strings"
	"time"
)

type User struct {
	ID             uint       `json:"uid"` // 企微用户id/公众号openid
	Mobile         string     `json:"mobile,omitempty"`
	UserID         string     `json:"user_id,omitempty"`          // 企微用户id/公众号openid
	AppID          string     `json:"app_id,omitempty"`           // 所属应用
	ExternalUserID string     `json:"external_user_id,omitempty"` // 外部/内部联系人的userid
	Name           string     `json:"name"`                       // 名称
	Avatar         string     `json:"avatar,omitempty"`           // 头像
	Type           int8       `json:"-"`                          // 人的类型，1外部微信用户，2企业微信内部用户
	Gender         int8       `json:"gender,omitempty"`           // 性别 0-未知 1-男性 2-女性
	Remark         string     `json:"remark,omitempty"`           // 备注
	Description    string     `json:"description,omitempty"`      // 描述
	Email          string     `json:"email,omitempty"`            // email
	BizMail        string     `json:"biz_mail,omitempty"`         // 企业邮箱
	Address        string     `json:"address,omitempty"`          // address
	StaffNo        string     `json:"staff_no,omitempty"`         // 内部员工工号
	LastLoginAt    *time.Time `json:"last_login_at,omitempty"`    // 最后登录时间
	LastLoginIP    string     `json:"last_login_ip,omitempty"`    // 最后登录ip
	QRCode         string     `json:"qr_code,omitempty"`          // 员工个人二维码（扫描可添加为外部联系人），仅在用户同意snsapi_privateinfo授权时返回
}
type UserListRequest struct {
}

type UserListResponse struct {
	Users []*User `json:"users"`
	*errors.Error
}

type UserRequest struct {
	UID int `json:"uid"`
}

type UserResponse struct {
	User *User `json:"user"`
	*errors.Error
}
type AccessTokenRequest struct {
	Token string `json:"token"`
}

func (a AccessTokenRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Token, validation.Required.Error("token不能为空"), validation.By(checkJWTToken)),
	)
}
func (a UserRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.UID, validation.Required.Error("uid不能为空")),
	)
}

func checkJWTToken(value interface{}) error {
	token, _ := value.(string)
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errors.New(http.StatusBadRequest, "Token invalid", nil)
	}
	return nil
}

type SnsLoginRequest struct {
	AID    int    `json:"aid"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"` // 手机号码
}
type SnsLoginResponse struct {
	*errors.Error
	Token string `json:"token,omitempty"`
	Exp   int64  `json:"exp,omitempty"`
	User  *User  `json:"user,omitempty"`
	Code  int    `json:"-"`
}

func (a SnsLoginRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AID, validation.Required.Error("应用id不能为空")),
		validation.Field(&a.Code, validation.When(a.Name == "", validation.Required.Error("code不能为空"))),
		validation.Field(&a.Name, validation.When(a.Code == "", validation.Required.Error("姓名不能为空"))),
		validation.Field(&a.Mobile, validation.When(a.Code == "", validation.Required.Error("手机号不能为空"))),
	)
}

type UserUpdateRequest struct {
	UID         int64  `json:"uid"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`      // 头像
	Gender      int8   `json:"gender"`      // 性别 0-未知 1-男性 2-女性
	Remark      string `json:"remark"`      // 备注
	Description string `json:"description"` // 描述
	Mobile      string `json:"mobile"`      // 手机号码
	Email       string `json:"email"`       // email
	BizMail     string `json:"biz_mail"`    // 企业邮箱
	Address     string `json:"address"`     // address
	StaffNo     string `json:"staff_no"`    // 内部员工工号
}

type UserUpdateResponse struct {
	*errors.Error
}

func (a UserUpdateRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.UID, validation.Required),
	)
}

type UpdatePhoneRequest struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
	Uid    int    `json:"-"`
}

func (a UpdatePhoneRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Mobile, validation.Required),
		validation.Field(&a.Code, validation.Required),
	)
}

type UpdatePhoneResponse struct {
	*errors.Error
}
