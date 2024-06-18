package worker

import (
	"fmt"
	"reflect"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"

	tracers "github.com/RichardKnop/machinery/example/tracers"

	workerTasks "fist-app/src/jobs/tasks"
	workerConfig "fist-app/src/lib/worker"
	logger "fist-app/src/lib/logger"
)

var _logger = logger.NewLogger("worker")

var server *machinery.Server = startServerWorker(workerConfig.WorkerConfig)

func startServerWorker(config *config.Config) *machinery.Server {
	_server, err := machinery.NewServer(config)
	if err != nil {
		_logger.Fatalf(err.Error())
		return nil
	}

	workerTasks := workerTasks.NewTask()
	err = _server.RegisterTasks(workerTasks.Tasks())
	if err != nil {
		_logger.Fatalf(err.Error())
		return nil
	}

	return _server
}

func InitWorker(tag string, concurrency int) error {
	consumerTag := tag
	cleanup, err := tracers.SetupTracer(consumerTag)
	if err != nil {
		_logger.Fatalf("Unable to instantiate a tracer: %s\n", err)
	}
	defer cleanup()

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	machineryWorker := server.NewWorker(consumerTag, concurrency)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorHandler := func(err error) {
		_logger.Errorf("I am an error handler: %s\n", err)
	}

	preTaskHandler := func(signature *tasks.Signature) {
		_logger.Infof("I am a start of task handler for: %s\n", signature.Name)
	}

	postTaskHandler := func(signature *tasks.Signature) {
		_logger.Infof("I am an end of task handler for: %s\n", signature.Name)
	}

	machineryWorker.SetPostTaskHandler(postTaskHandler)
	machineryWorker.SetErrorHandler(errorHandler)
	machineryWorker.SetPreTaskHandler(preTaskHandler)

	return machineryWorker.Launch()
}

func AsyncTask(taskName string, args ...interface{}) error {
	var taskArgs []tasks.Arg
	for _, arg := range args {
		taskArgs = append(taskArgs, tasks.Arg{
			Type:  reflect.ValueOf(arg).Type().Name(),
			Value: reflect.ValueOf(arg).Interface(),
		})
	}
	task := &tasks.Signature{
		Name: taskName,
		Args: taskArgs,
	}

	asyncResult, err := server.SendTask(task)
	if err != nil {
		_logger.Error(err)
	}


	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return fmt.Errorf("Getting task result failed with error: %s", err)
	}

	_logger.Info(tasks.HumanReadableResults(results))

	return nil
}
