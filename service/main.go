package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/barbierodb1/workmodexp/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type todoService struct {
	proto.UnimplementedTodoServiceServer

	todos map[uuid.UUID]*proto.Todo
}

func (t *todoService) Create(ctx context.Context, todo *proto.Todo) (*proto.Todo, error) {
	if t.todos == nil {
		t.todos = map[uuid.UUID]*proto.Todo{}
	}

	if todo.GetId() != "" {
		return nil, status.Error(codes.InvalidArgument, "ID must not be set for creation")
	}

	id := uuid.New()
	newTodo := &proto.Todo{
		Id:           id.String(),
		Task:         todo.GetTask(),
		CreatedAt:    timestamppb.Now(),
		CustomField:  todo.GetCustomField(),
		AnotherField: todo.GetAnotherField(),
		NewField:     todo.GetNewField(),
	}
	t.todos[id] = newTodo

	return newTodo, status.FromContextError(ctx.Err()).Err()
}

func (t *todoService) Get(ctx context.Context, todo *proto.Todo) (*proto.Todo, error) {
	if t.todos == nil {
		t.todos = map[uuid.UUID]*proto.Todo{}
	}

	if todo.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "ID must be set for Get")
	}

	parsedID, err := uuid.Parse(todo.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid uuid. err: %s", err.Error())
	}

	foundTodo, ok := t.todos[parsedID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "id %s not found", err.Error())
	}

	return foundTodo, status.FromContextError(ctx.Err()).Err()
}

func main() {
	lis, err := net.Listen("tcp", "localhost:55051")
	if err != nil {
		log.Fatal(err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	server := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)
	proto.RegisterTodoServiceServer(server, &todoService{})
	defer server.Stop()

	go func(lis net.Listener, server *grpc.Server) {
		log.Printf("Running server at address %s...", lis.Addr())
		if err := server.Serve(lis); err != nil {
			log.Print(err)
		}
	}(lis, server)

	<-sigs
	log.Print("Closing server...")
	server.GracefulStop()
}
