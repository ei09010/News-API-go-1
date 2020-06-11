# News Api client library for Go
News API is a simple HTTP REST API for searching and retrieving live articles from all over the web. It can help you answer questions like:

- What top stories is the NY Times running right now?
- What new articles were published about the next iPhone today?
- Has my company or product been mentioned or reviewed by any blogs recently?

You can search for articles with any combination of the following criteria:

- Keyword or phrase. Eg: find all articles containing the word 'Microsoft'.
- Date published. Eg: find all articles published yesterday.
- Source name. Eg: find all articles by 'TechCrunch'.
- Source domain name. Eg: find all articles published on nytimes.com.
- Language. Eg: find all articles written in English.

You can sort the results in the following orders:

- Date published
- Relevancy to search keyword
- Popularity of source

You need an API key to use the API - this is a unique key that identifies your requests. They're free for development, open-source, and non-commercial use. You can get one here: [https://newsapi.org](https://newsapi.org).

## Installation
The News API client library is available on GitHub, just need to run the following command:
```shell
go get https://github.com/ei09010/News-API-go
```

## Usage example
```
package main

import (
	"News-API-go/models"
	"News-API-go/constants"
	"News-API-go/client"
	"fmt"
)

// Example request
var newsClient = client.NewClient("https://newsapi.org", "XXXXXXXXXXXXXXXXXXXXXXXXXXXX")

func main(){

	topHeadlineRequest := models.TopHeadlinesRequest{
		Category: constants.Business,
		Country:  constants.US,
		Language: "EN",
	}

	response, err := newsClient.GetTopHeadlines(topHeadlineRequest)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("we had the following result amount: ", response.TotalResults)

	for _, article := range response.Articles{
	
		fmt.Println(article.SourceName)

		fmt.Println(article.Author)

		fmt.Println(article.Title)

		fmt.Println(article.Description)

		fmt.Println(article.Url)

		fmt.Println(article.UrlToImage)

		fmt.Println(article.PublishedAt)

		fmt.Println(`
		
		*******************
		
		`)
	}
		
```