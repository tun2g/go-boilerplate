package tasks

type Task struct {
	TaskQueue map[string]interface{}
}

func NewTask() *Task {
	return &Task{
		TaskQueue: make(map[string]interface{}),
	}
}

func (task *Task) Tasks() map[string]interface{} {
	return task.TaskQueue
}
