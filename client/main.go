package main

import (
	"context"
	"log"

	"github.com/barbierodb1/workmodexp/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:55051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewTodoServiceClient(conn)

	runSomething(client)
}

func runSomething(client proto.TodoServiceClient) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Print("running a Create and then a Get")

	todo := &proto.Todo{Task: "This is a task", CustomField: "custom field from v0.0.2"}
	res, err := client.Create(ctx, todo)
	if err != nil {
		log.Print("ERR Creating:", err)
		return
	}
	if res.Task != todo.Task {
		log.Printf("WARN: res.Task == %s but todo.Task == %s", res.Task, todo.Task)
	}

	gotten, err := client.Get(ctx, &proto.Todo{Id: res.GetId()})
	if err != nil {
		log.Print("ERR Getting: ", err)
		return
	}
	if gotten.Task != todo.Task {
		log.Printf("WARN: gotten.Task == %s but todo.Task == %s", gotten.Task, todo.Task)
	}

	log.Print("Finished running Create and then Get")
}
