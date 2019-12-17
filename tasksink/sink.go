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

	var taskstring interface{}
	counter := 0
	for counter < 100 {
		taskstring, _ = queue.Poll()
		counter++

		if taskstring == nil {
			fmt.Println("Break, empty queue")
			break
		}
		fmt.Println("Polling: ", taskstring)

		value := taskstring.(*core.HazelcastJSONValue)

		var taskData task.TaskData
		value.Unmarshal(&taskData)
		fmt.Println("ID: ", taskData.ID)
	}

}
