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
	fmt.Println("Crawling URL: ", task.UrlID)
	fmt.Println("Task: ", task)
	// TODO: Implement the logic to crawl the URL
}
