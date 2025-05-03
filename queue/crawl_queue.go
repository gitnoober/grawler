package queue

import (
	"fmt"
	"sync"

	"github.com/gitnoober/grawler/models"
	"github.com/gitnoober/grawler/repository"
)

var CrawlQueue chan *models.Task
var CrawlQueueMutex sync.Mutex
var taskRepo repository.TaskRepository

func InitQueue(bufferSize int, numWorkers int) {
	CrawlQueue = make(chan *models.Task, bufferSize)
	startWorkerPool(numWorkers)
}

func startWorkerPool(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go func(){
			for task := range CrawlQueue {
				crawlUrl(task)
			}
		}()
	}
}

func crawlUrl(task *models.Task) {
	err := taskRepo.CreateTask(task)
	if err != nil {
		fmt.Println("Error creating task: ", err)
		return
	}
	// TODO: Implement the logic to crawl the URL
}
