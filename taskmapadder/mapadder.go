package main

import (
	"fmt"

	task "../taskdto"
	"github.com/hazelcast/hazelcast-go-client"
)

func main() {
	config := task.GetConfig()
	client, _ := hazelcast.NewClientWithConfig(config)

	mp, _ := client.GetReplicatedMap("taskmap")

	counter := 0
	for counter < 100 {

		var taskData task.TaskData
		taskData.ID = int32(counter)
		taskData.Name = "test"
		mp.Put(counter, task.Encode(taskData))
		fmt.Println("Added taskData to Map")
		counter++
	}

}
