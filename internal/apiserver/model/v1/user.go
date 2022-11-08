package v1

import "iamgeek/internal/apiserver/meta/v1"

type User struct {
}

type UserList struct {
	v1.ListMeta
	Items []*User `json:"items"`
}
