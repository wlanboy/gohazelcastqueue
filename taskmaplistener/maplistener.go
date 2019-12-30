package main

import (
	"fmt"

	task "../taskdto"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/core"
)

type entryListener struct {
}

func (l *entryListener) EntryAdded(event core.EntryEvent) {
	fmt.Println("New Map Entry: ", event.Key(), " ", event.Value())
}

func main() {
	done := make(chan bool)
	go forever()
	<-done
}

func forever() {
	config := task.GetConfig()
	client, _ := hazelcast.NewClientWithConfig(config)

	mp, _ := client.GetReplicatedMap("taskmap")
	mp.AddEntryListener(&entryListener{})
}
