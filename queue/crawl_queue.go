package queue

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gitnoober/grawler/models"
	"github.com/gitnoober/grawler/repository"
)

var CrawlQueue chan *models.Task
var CrawlQueueMutex sync.Mutex

func InitQueue(bufferSize int, numWorkers int, taskRepo repository.TaskRepository, urlRepo repository.UrlRepository, taskResponseRepo repository.TaskResponseRepository) {
	CrawlQueue = make(chan *models.Task, bufferSize)
	startWorkerPool(numWorkers, taskRepo, urlRepo, taskResponseRepo)
}

func startWorkerPool(numWorkers int, taskRepo repository.TaskRepository, urlRepo repository.UrlRepository, taskResponseRepo repository.TaskResponseRepository) {
	for i := 0; i < numWorkers; i++ {
		go func(){
			for task := range CrawlQueue {
				crawlUrl(task, taskRepo, urlRepo, taskResponseRepo)
			}
		}()
	}
}

func crawlUrl(task *models.Task, taskRepo repository.TaskRepository, urlRepo repository.UrlRepository, taskResponseRepo repository.TaskResponseRepository) {
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
	urlMetadata, err := urlRepo.GetUrlMetadataByID(task.UrlID)
	if err != nil {
		fmt.Println("Error getting URL metadata: ", err)
		task.Status = models.TaskStatusFailed
		err = taskRepo.UpdateTask(task)
		if err != nil {
			fmt.Println("Error updating task: ", err)
		}
		return
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	fmt.Println("URL metadata: ", urlMetadata)
	resp, err := client.Get(urlMetadata.Url)
	if err != nil {
		fmt.Println("Error getting URL: ", err)
		task.Status = models.TaskStatusFailed
		err = taskRepo.UpdateTask(task)
		if err != nil {
			fmt.Println("Error updating task: ", err)
		}
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Crawled URL: %s → Status Code: %d\n", urlMetadata.Url, resp.StatusCode)


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading URL: ", err)
		task.Status = models.TaskStatusFailed
		err = taskRepo.UpdateTask(task)
		if err != nil {
			fmt.Println("Error updating task: ", err)
		}
		return
	}
	
	taskResponse := &models.TaskResponse{
		Body: string(body),
		TaskID: task.ID,
		UrlID: urlMetadata.ID,
	}
	err = taskResponseRepo.CreateTaskResponse(taskResponse)
	if err != nil {
		fmt.Println("Error creating task response: ", err)
		task.Status = models.TaskStatusFailed
		err = taskRepo.UpdateTask(task)
		if err != nil {
			fmt.Println("Error updating task: ", err)
		}
		return
	}

	task.Status = models.TaskStatusCompleted
	err = taskRepo.UpdateTask(task)
	if err != nil {
		fmt.Println("Error updating task: ", err)
	}
}