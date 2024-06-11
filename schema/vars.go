package schema

import (
	"database/sql"
)

var ErrNotFound = sql.ErrNoRows

type Error struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Err  error  `json:"err,omitempty"`
}

func (e Error) Error() string {
	return e.Msg
}
