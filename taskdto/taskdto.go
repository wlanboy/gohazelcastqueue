package taskdto

import (
	"encoding/json"
	"log"
)

/*TaskData struct containing id and name*/
type TaskData struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

/*Decode jsonstring*/
func Decode(jsonstring string) TaskData {
	data := TaskData{}

	err := json.Unmarshal([]byte(jsonstring), &data)
	if err != nil {
		log.Println("Cannot parse JSON")
		log.Println(err)
	}

	return data
}

/*Encode jsonstring*/
func Encode(task TaskData) string {

	jsonsdata, err := json.Marshal(task)
	if err != nil {
		log.Println("Cannot parse JSON")
		log.Println(err)
	}
	jsonstring := string(jsonsdata)
	return jsonstring
}
