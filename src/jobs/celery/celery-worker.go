package worker

import (
	"time"

	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"

	"fist-app/src/config"
)

type CeleryWorker struct {
	worker *gocelery.CeleryClient
}

func StartCeleryWorker() *CeleryWorker {
	// create redis connection pool
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(config.AppConfiguration.RedisBrokerUrl)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	// initialize celery client
	cli, _ := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		1,
	)

	return &CeleryWorker{
		worker: cli,
	}
}

func (worker *CeleryWorker) CeleryAsyncTask(taskName string, args ...interface{}) (interface{}, error) {
	// run task
	asyncResult, err := worker.worker.Delay(taskName, args...)
	if err != nil {
		panic(err)
	}

	// get results from backend with timeout
	res, err := asyncResult.Get(10 * time.Second)
	if err != nil {
		return nil, err
	}

	return res, err
}
