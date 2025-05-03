package summarizer

import (
	"fmt"
	"strings"
	"time"

	"github.com/gitnoober/grawler/models"
	"github.com/gitnoober/grawler/repository"
	"golang.org/x/net/html"
)

const (
	SUMMARIZER_POLL_INTERVAL = 10 * time.Second
	PROMPT = `
	You are a helpful assistant that summarizes the content of a website.
	Please summarize the content of the website in a few sentences and highlight the most important points.
	`
)

func StartSummarizer(taskResponseRepo repository.TaskResponseRepository, urlSummaryRepo repository.UrlSummaryRepository) {
	fmt.Println("Starting summarizer")
	for {
		responses, err := taskResponseRepo.GetUnprocessedResponses(10)
		if err != nil {
			fmt.Println("Error getting unprocessed responses: ", err)
			continue
		}
		for _, response := range responses {
			summarizeAndSave(response, urlSummaryRepo)
			response.SummaryGenerated = true
			err = taskResponseRepo.UpdateTaskResponseSummary(response)
			if err != nil {
				fmt.Println("Error updating task response: ", err)
			}
		}
		time.Sleep(SUMMARIZER_POLL_INTERVAL)
	}
}

func ExtractTextFromHTML(htmlContent string) string {
    doc, err := html.Parse(strings.NewReader(htmlContent))
    if err != nil {
        return htmlContent
    }
    var textContent strings.Builder
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.TextNode {
            textContent.WriteString(n.Data + " ")
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
    return textContent.String()
}

func summarizeAndSave(response *models.TaskResponse, urlSummaryRepo repository.UrlSummaryRepository) {
	cleanedBody := ExtractTextFromHTML(response.Body)
	prompt := fmt.Sprintf("%s\n\n%s", PROMPT, cleanedBody)
	summary, err := GetSummary(prompt)
	if err != nil {
		fmt.Println("Error getting summary: ", err)
		return
	}
	urlSummary := &models.UrlSummary{
		UrlID: response.UrlID,
		TaskID: response.TaskID,
		Summary: summary,
	}
	urlSummaryRepo.CreateUrlSummary(urlSummary)
}
