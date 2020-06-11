package client

import (
	"News-API-go/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"strconv"
	"log"
)

// Client to aggregate basic information regarding http calls
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	ApiKey     string
	HttpClient *http.Client
}

func NewClient(myUrl string, apiKey string) *Client{

	parsedUrl, _ := url.Parse(myUrl)

	return &Client{
		BaseURL: parsedUrl,
		HttpClient: &http.Client{},
		ApiKey: apiKey,
	}
}

// GetTopHeadlines to request top headlines
func (c *Client) GetTopHeadlines(req models.TopHeadlinesRequest) (*models.ArticlesResult, error) {

	relativePath := &url.URL{Path: "/v2/top-headlines"}

	queryParams := []string{}

	if req.Keyword != "" {
		queryParams = append(queryParams, "q="+req.Keyword)
	}

	if len(req.Sources) > 0 {
		queryParams = append(queryParams, "sources="+strings.Join(req.Sources[:], ","))
	}

	if string(req.Category) != "" {
		queryParams = append(queryParams, "category="+strings.ToLower(string(req.Category)))
	}

	if string(req.Language) != "" {
		queryParams = append(queryParams, "language="+strings.ToLower(string(req.Language)))
	}

	if string(req.Country) != "" {
		queryParams = append(queryParams, "country="+strings.ToLower(string(req.Country)))
	}

	//page information

	if req.Page > 0 {
		queryParams = append(queryParams, "page="+strconv.Itoa(req.Page))
	}

	if req.PageSize > 0 {
		queryParams = append(queryParams, "pageSize="+strconv.Itoa(req.PageSize))
	}

	queryString := strings.Join(queryParams[:], "&")

	if len(queryParams) > 0{
		relativePath.Path += "?"
	}

	relativePath.Path += queryString

	parsedPath, err := url.Parse(relativePath.Path)
	
	if err != nil{
		log.Fatal(err)
	}

	urlAbsoluteReference := c.BaseURL.ResolveReference(parsedPath)

	return c.makeRequest(urlAbsoluteReference, queryString)

}

func (c *Client) GetEverything(req models.EverythingRequest) (*models.ArticlesResult, error) {

	relativePath := &url.URL{Path: "/v2/everything"}

	queryParams := []string{}

	if req.Keyword != "" {
		queryParams = append(queryParams, "q="+req.Keyword)
	}

	if len(req.Sources) > 0 {
		queryParams = append(queryParams, "sources="+strings.Join(req.Sources[:], ","))
	}

	if len(req.Domains) > 0 {
		queryParams = append(queryParams, "domains="+strings.Join(req.Domains[:], ","))
	}

	if req.From != nil {
		queryParams = append(queryParams, "from="+req.From.Format("2006-01-02"))
	}

	if req.To != nil {
		queryParams = append(queryParams, "to="+req.To.Format("2006-01-02"))
	}

	if string(req.Language) != "" {
		queryParams = append(queryParams, "language="+strings.ToLower(string(req.Language)))
	}

	if req.SortBy != "" {
		queryParams = append(queryParams, "sortBy="+strings.ToLower(string(req.SortBy)))
	}

	//page information

	if req.Page > 0 {
		queryParams = append(queryParams, "page="+strconv.Itoa(req.Page))
	}

	if req.PageSize > 0 {
		queryParams = append(queryParams, "pageSize="+strconv.Itoa(req.PageSize))
	}

	queryString := strings.Join(queryParams[:], "&")

	if len(queryParams) > 0{
		relativePath.Path += "?"
	}

	relativePath.Path += queryString

	parsedPath, err := url.Parse(relativePath.Path)
	
	if err != nil{
		log.Fatal(err)
	}

	urlAbsoluteReference := c.BaseURL.ResolveReference(parsedPath)

	return c.makeRequest(urlAbsoluteReference, queryString)

}


func (c *Client) makeRequest(url *url.URL, queryString string) (*models.ArticlesResult, error) {

	customReq, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return nil, err
	}

	customReq.Header.Set("accept", "application/json")
	customReq.Header.Set("x-api-key", c.ApiKey)

	httpResponse, err := c.HttpClient.Do(customReq)

	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	var articleResults models.ArticlesResult

	if httpResponse != nil {

		myResponse, err := ioutil.ReadAll(httpResponse.Body)

		if err != nil {
			return nil, err
		}

		httpResponse.Body.Close()

		err = json.Unmarshal(myResponse, &articleResults)

		if articleResults.Status != string(http.StatusOK) {

			err = json.Unmarshal(myResponse, &articleResults.Error)

			return &articleResults, err
		}

	} else {

			return &articleResults, err
	}

	return &articleResults, err
}
