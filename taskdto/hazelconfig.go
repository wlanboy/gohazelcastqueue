package taskdto

import (
	"time"

	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/config"
)

/*GetConfig to generate hazelcast configuration*/
func GetConfig() *config.Config {
	config := hazelcast.NewConfig()
	//config.GroupConfig().SetName("GOWORKER")
	networkConfig := config.NetworkConfig()
	networkConfig.AddAddress("nuc:5701")
	networkConfig.SetSmartRouting(true)
	networkConfig.SetRedoOperation(true)
	networkConfig.SetConnectionTimeout(10 * time.Second)
	networkConfig.SetConnectionAttemptPeriod(5 * time.Second)
	networkConfig.SetConnectionAttemptLimit(5)
	return config
}
