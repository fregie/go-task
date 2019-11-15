package task

import (
	"log"
	"os"
	"time"
)

var TaskLog *log.Logger

type TaskExecutor struct {
	taskMap map[string]Task
}

type Task struct {
	Exec interface{}
	// para interface{}
	Interval    string
	Immediately bool
}

func init() {
	TaskLog = log.New(os.Stdout,
		"[TASK] ",
		log.Ldate|log.Ltime)
}

func NewTaskExecutor(taskMap map[string]Task) *TaskExecutor {
	taskExecutor := new(TaskExecutor)
	taskExecutor.taskMap = taskMap

	return taskExecutor
}

func registerTask(task Task) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("task register panic：", err) // 这里的err其实就是panic传入的内容
	// 	}
	// }()
	D, err := time.ParseDuration(task.Interval)
	if err != nil {
		TaskLog.Printf("registerTask error:%s", err)
	}
	ticker := time.NewTicker(D)
	if task.Immediately {
		task.Exec.(func())()
	}
	for range ticker.C {
		task.Exec.(func())()
	}
}

func (taskExecutor *TaskExecutor) Run() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("task panic：", err) // 这里的err其实就是panic传入的内容
	// 	}
	// }()
	for name, task := range taskExecutor.taskMap {
		TaskLog.Printf("register task [%s]", name)
		go registerTask(task)
	}
}

func StartFromMap(taskMap map[string]Task) {
	TE := NewTaskExecutor(taskMap)
	TE.Run()
}
