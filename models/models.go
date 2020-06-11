package models

import (
	"News-API-go/constants"
	"time"
)

type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	SourceName  Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content string `json:"content"`
}

type ArticlesResult struct {
	Status       string `json:"status"`
	Error 		Error
	TotalResults int  `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type TopHeadlinesRequest struct {
	Keyword  string
	Sources  []string
	Category constants.Category
	Language constants.Language
	Country  constants.Country
	Page     int
	PageSize int
}

type EverythingRequest struct {
	Keyword  string
	Sources  []string
	Domains  []string
	From 	 *time.Time
	To 		 *time.Time
	Language constants.Language
	SortBy	 constants.SortBy
	Page     int
	PageSize int
}


type Error struct {
	Status  string `json:"status"`
	Code   string `json:"code"`
	Message string `json:"message"`
}
