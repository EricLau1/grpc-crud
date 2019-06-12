package server

import (
	"context"
	"errors"
	"fmt"
	"grpc-crud/server/pb/messages"
	"grpc-crud/server/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	Users []*messages.User
}

func Run(port string) {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	messages.RegisterUserServiceServer(s, &Server{})
	fmt.Printf("rpc server on %s...\n", port)
	s.Serve(listen)
}

func (s *Server) Add(ctx context.Context, user *messages.User) (*messages.User, error) {
	if user.GetEmail() == "" || user.GetPassword() == "" {
		return nil, errors.New("Verifique os campos")
	}
	user.ObjectId = &messages.ObjectId{Uid: utils.NewID()}
	s.Users = append(s.Users, user)
	return user, nil
}

func (s *Server) Find(ctx context.Context, obj_id *messages.ObjectId) (*messages.User, error) {
	if s.Users != nil {
		for _, u := range s.Users {
			if u.GetObjectId().GetUid() == obj_id.GetUid() {
				return u, nil
			}
		}
	}
	return nil, errors.New("User not found")
}

func (s *Server) Update(ctx context.Context, user *messages.User) (*messages.User, error) {
	if s.Users != nil {
		for i, u := range s.Users {
			if u.GetObjectId().GetUid() == user.GetObjectId().GetUid() {
				s.Users[i].Email = user.GetEmail()
				s.Users[i].Password = user.GetPassword()
				return s.Users[i], nil
			}
		}
	}
	return nil, errors.New("User not found")
}

func (s *Server) Delete(ctx context.Context, obj_id *messages.ObjectId) (*messages.ObjectId, error) {
	if s.Users != nil {
		for i, u := range s.Users {
			if u.GetObjectId().GetUid() == obj_id.GetUid() {
				s.Users = append(s.Users[:i], s.Users[i+1:]...)
				return u.GetObjectId(), nil
			}
		}
	}
	return nil, errors.New("User not found")
}
