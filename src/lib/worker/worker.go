package worker

import (
	"strings"

	"github.com/RichardKnop/machinery/v1/config"

	envConfig "fist-app/src/config"
)

func configWorker() *config.Config {
	if strings.Split(envConfig.AppConfiguration.RedisBrokerUrl, ":")[0] == "amqp" {
		return &config.Config{
			Broker:          envConfig.AppConfiguration.RedisBrokerUrl,
			ResultBackend:   envConfig.AppConfiguration.RedisBrokerUrl,
			ResultsExpireIn: 3600,
			AMQP: &config.AMQPConfig{
				Exchange:      "machinery_exchange",
				ExchangeType:  "direct",
				BindingKey:    "test",
				PrefetchCount: 3,
			},
		}
	}

	return &config.Config{
		ResultsExpireIn: 3600,
		Broker:          envConfig.AppConfiguration.RedisBrokerUrl,
		ResultBackend:   envConfig.AppConfiguration.RedisBrokerUrl,
	}
}

var WorkerConfig = configWorker()