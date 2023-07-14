package main

import (
	gen "main/gen/go"
)

//type User struct {
//	Name      string `json:"name"`
//	Family    string `json:"family"`
//	Id        int32  `json:"id"`
//	Age       int32  `json:"age"`
//	Sex       string `json:"sex"`
//	CreatedAt int64  `json:"createdAt"`
//}

func GetUsers(userId string, messageId int) []*gen.User {
	if userId == "" {
		return get_top_100_users()
	} else {
		return get_user_by_id(userId)
	}
}
