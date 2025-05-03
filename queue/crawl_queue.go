package queue

import (
	"fmt"
	"sync"

	"github.com/gitnoober/grawler/models"
	"github.com/gitnoober/grawler/repository"
)

var CrawlQueue chan *models.Task
var CrawlQueueMutex sync.Mutex

func InitQueue(bufferSize int, numWorkers int, taskRepo repository.TaskRepository) {
	CrawlQueue = make(chan *models.Task, bufferSize)
	startWorkerPool(numWorkers, taskRepo)
}

func startWorkerPool(numWorkers int, taskRepo repository.TaskRepository) {
	for i := 0; i < numWorkers; i++ {
		go func(){
			for task := range CrawlQueue {
				crawlUrl(task, taskRepo)
			}
		}()
	}
}

func crawlUrl(task *models.Task, taskRepo repository.TaskRepository) {
	err := taskRepo.CreateTask(task)
	if err != nil {
		fmt.Println("Error creating task: ", err)
		return
	}
	fmt.Println("Task created: ", task)
	task.Status = models.TaskStatusRunning
	err = taskRepo.UpdateTask(task)
	if err != nil {
		fmt.Println("Error updating task: ", err)
		return
	}
	fmt.Println("Crawling URL: ", task.ID)
	fmt.Println("Task: ", task)
	task.Status = models.TaskStatusCompleted
	err = taskRepo.UpdateTask(task)
	if err != nil {
		fmt.Println("Error updating task: ", err)
		return
	}
}
