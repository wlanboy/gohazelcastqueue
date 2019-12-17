package main

import (
	"fmt"

	task "../taskdto"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/core"
)

func main() {
	config := task.GetConfig()
	client, _ := hazelcast.NewClientWithConfig(config)

	queue, _ := client.GetQueue("taskqueue")

	for i := 0; i < 10; i++ {
		data := task.TaskData{}
		data.Name = "Test"
		data.ID = int32(i)
		jsonValue, _ := core.CreateHazelcastJSONValue(data)

		queue.Offer(jsonValue)
		fmt.Println("Offer: ", jsonValue)
	}
}
