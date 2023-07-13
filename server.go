package main

import (
	"context"
	"fmt"
	_ "fmt"
	gen "main/gen/go"
)

type server struct {
	gen.UnimplementedBizServer
}

func (s *server) GetUsers(ctx context.Context, input *gen.GetUserInput1) (*gen.GetUserOutput, error) {
	fmt.Println("Received GetUsers request", input)
	res := GetUsers(fmt.Sprint(input.UserId), int(input.MessageId))
	return &gen.GetUserOutput{
		User:      res,
		MessageId: 3,
	}, nil
}

func (s *server) GetUsersWithSqlInjection(ctx context.Context, input *gen.GetUserInput2) (*gen.GetUserOutput, error) {
	fmt.Println("Received GetUsersWithSqlInjection request", input)
	res := GetUsers(input.UserId, int(input.MessageId))
	return &gen.GetUserOutput{
		User:      res,
		MessageId: 3,
	}, nil
}
