package main

import (
	gen "main/gen/go"
)

func GetUsers(userId string, messageId int) []*gen.User {
	if userId == "" {
		return get_top_100_users()
	} else {
		return get_user_by_id(userId)
	}
}
