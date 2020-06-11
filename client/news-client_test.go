package client

import (
	"News-API-go/constants"
	"News-API-go/models"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestGetTopHeadlines_StandardRequest_ReturnsSuccessResponse(t *testing.T) {

	//Arrange
	expectedCorrectResponse := `{"status": "ok","totalResults": 1,"articles": [{"source": {"id": "the-wall-street-journal","name": "The Wall Street Journal"},"author": "Alison Sider","title": "Airlines Got $25 Billion in Stimulus; Industry Still Expected to Shrink - The Wall Street Journal","description": "U.S. carriers are planning to operate smaller companies with fewer flights and employees","url": "https://www.wsj.com/articles/airlines-got-25-billion-in-stimulus-industry-still-expected-to-shrink-11591527600","urlToImage": "https://images.wsj.net/im-194495/social","publishedAt": "2020-06-07T14:11:39Z","content": "Federal stimulus money for airlines is keeping them afloat through the coronavirus pandemic, but its not proving to be enough to sustain the industry at its pre-pandemic size.Carriers say they will… [+320 chars]"}]}`
	expectedCorrectRequest := `/v2/top-headlines?category=business&language=en&country=us`
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		request := req.URL.Path+"?"+req.URL.RawQuery

		if request == expectedCorrectRequest {
			io.WriteString(w, expectedCorrectResponse)
		} else {
			io.WriteString(w, "Bad request")
		}
	}))

	defer ts.Close()

	expectedStatus := "ok"
	expectedTotalResults := 1
	expectedFirstArticleSourceId := "the-wall-street-journal"
	expectedFirstArticleSourceName := "The Wall Street Journal"
	expectedFirstArticleAuthor := "Alison Sider"
	expectedFirstArticleTitle := "Airlines Got $25 Billion in Stimulus; Industry Still Expected to Shrink - The Wall Street Journal"
	expectedFirstArticleDescription := "U.S. carriers are planning to operate smaller companies with fewer flights and employees"
	expectedFirstArticleUrl := "https://www.wsj.com/articles/airlines-got-25-billion-in-stimulus-industry-still-expected-to-shrink-11591527600"
	expectedFirstArticleUrlToImage := "https://images.wsj.net/im-194495/social"
	expectedPublishedAt := "2020-06-07 14:11:39 +0000 UTC"
	
	
	topHeadlinesReq := models.TopHeadlinesRequest{
		Category: constants.Business,
		Country:  constants.US,
		Language: "EN",
	}

	newsClient := NewClient(ts.URL, "testKey")

	// Act
	response, err := newsClient.GetTopHeadlines(topHeadlinesReq)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Assert
	if response.Status != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Status, expectedStatus)
	}


	if response.TotalResults != expectedTotalResults {
		t.Errorf("handler returned unexpected total results: got %v want %v",
		response.TotalResults, expectedTotalResults)
	}

	if response.Articles[0].SourceName.Id != expectedFirstArticleSourceId {
		t.Errorf("handler returned unexpected firt source id: got %v want %v",
		response.Articles[0].SourceName.Id, expectedFirstArticleSourceId)
	}

	if response.Articles[0].SourceName.Name != expectedFirstArticleSourceName {
		t.Errorf("handler returned unexpected firt source name: got %v want %v",
		response.Articles[0].SourceName.Name, expectedFirstArticleSourceName)
	}

	if response.Articles[0].Author != expectedFirstArticleAuthor {
		t.Errorf("handler returned unexpected firt article author: got %v want %v",
		response.Articles[0].Author, expectedFirstArticleAuthor)
	}

	if response.Articles[0].Title != expectedFirstArticleTitle {
		t.Errorf("handler returned unexpected firt article title: got %v want %v",
		response.Articles[0].Title, expectedFirstArticleTitle)
	}

	if response.Articles[0].Description != expectedFirstArticleDescription {
		t.Errorf("handler returned unexpected firt article description: got %v want %v",
		response.Articles[0].Description, expectedFirstArticleDescription)
	}

	if response.Articles[0].Url != expectedFirstArticleUrl {
		t.Errorf("handler returned unexpected firt article url: got %v want %v",
		response.Articles[0].Url, expectedFirstArticleUrl)
	}

	if response.Articles[0].UrlToImage != expectedFirstArticleUrlToImage {
		t.Errorf("handler returned unexpected firt article urlToImage: got %v want %v",
		response.Articles[0].UrlToImage, expectedFirstArticleUrlToImage)
	}

	if response.Articles[0].PublishedAt.String() != expectedPublishedAt {
		t.Errorf("handler returned unexpected firt article publishedAt: got %v want %v",
		response.Articles[0].PublishedAt, expectedPublishedAt)
	}

	if response.Articles[0].Content == "" {
		t.Errorf("handler returned empty content")
	}

	if response.Error.Status != "ok"{
		t.Errorf("there is a error status, but it shouldn't be there")
	}

	articleLength := len(response.Articles)

	if response.TotalResults != articleLength {
		t.Errorf("handler returned unexpected number of articles: got %v want %v",
		response.TotalResults, articleLength)
	}	
}


func TestGetTopHeadlines_StandardRequestWithPageInfo_ReturnsSuccessResponse(t *testing.T) {

	//Arrange
	expectedCorrectResponse := `{"status": "ok","totalResults": 1,"articles": [{"source": {"id": "the-wall-street-journal","name": "The Wall Street Journal"},"author": "Alison Sider","title": "Airlines Got $25 Billion in Stimulus; Industry Still Expected to Shrink - The Wall Street Journal","description": "U.S. carriers are planning to operate smaller companies with fewer flights and employees","url": "https://www.wsj.com/articles/airlines-got-25-billion-in-stimulus-industry-still-expected-to-shrink-11591527600","urlToImage": "https://images.wsj.net/im-194495/social","publishedAt": "2020-06-07T14:11:39Z","content": "Federal stimulus money for airlines is keeping them afloat through the coronavirus pandemic, but its not proving to be enough to sustain the industry at its pre-pandemic size.Carriers say they will… [+320 chars]"}]}`
	expectedCorrectRequest := `/v2/top-headlines?category=business&language=en&country=us&page=1&pageSize=1`
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		request := req.URL.Path+"?"+req.URL.RawQuery

		if request == expectedCorrectRequest {
			io.WriteString(w, expectedCorrectResponse)
		} else {
			io.WriteString(w, "Bad request")
		}
	}))

	defer ts.Close()

	expectedStatus := "ok"
	expectedTotalResults := 1
	expectedFirstArticleSourceId := "the-wall-street-journal"
	expectedFirstArticleSourceName := "The Wall Street Journal"
	expectedFirstArticleAuthor := "Alison Sider"
	expectedFirstArticleTitle := "Airlines Got $25 Billion in Stimulus; Industry Still Expected to Shrink - The Wall Street Journal"
	expectedFirstArticleDescription := "U.S. carriers are planning to operate smaller companies with fewer flights and employees"
	expectedFirstArticleUrl := "https://www.wsj.com/articles/airlines-got-25-billion-in-stimulus-industry-still-expected-to-shrink-11591527600"
	expectedFirstArticleUrlToImage := "https://images.wsj.net/im-194495/social"
	expectedPublishedAt := "2020-06-07 14:11:39 +0000 UTC"
	
	
	topHeadlinesReq := models.TopHeadlinesRequest{
		Category: constants.Business,
		Country:  constants.US,
		Language: "EN",
		Page:     1,
		PageSize: 1,
	}

	newsClient := NewClient(ts.URL, "testKey")

	// Act
	response, err := newsClient.GetTopHeadlines(topHeadlinesReq)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Assert
	if response.Status != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Status, expectedStatus)
	}


	if response.TotalResults != expectedTotalResults {
		t.Errorf("handler returned unexpected total results: got %v want %v",
		response.TotalResults, expectedTotalResults)
	}

	if response.Articles[0].SourceName.Id != expectedFirstArticleSourceId {
		t.Errorf("handler returned unexpected firt source id: got %v want %v",
		response.Articles[0].SourceName.Id, expectedFirstArticleSourceId)
	}

	if response.Articles[0].SourceName.Name != expectedFirstArticleSourceName {
		t.Errorf("handler returned unexpected firt source name: got %v want %v",
		response.Articles[0].SourceName.Name, expectedFirstArticleSourceName)
	}

	if response.Articles[0].Author != expectedFirstArticleAuthor {
		t.Errorf("handler returned unexpected firt article author: got %v want %v",
		response.Articles[0].Author, expectedFirstArticleAuthor)
	}

	if response.Articles[0].Title != expectedFirstArticleTitle {
		t.Errorf("handler returned unexpected firt article title: got %v want %v",
		response.Articles[0].Title, expectedFirstArticleTitle)
	}

	if response.Articles[0].Description != expectedFirstArticleDescription {
		t.Errorf("handler returned unexpected firt article description: got %v want %v",
		response.Articles[0].Description, expectedFirstArticleDescription)
	}

	if response.Articles[0].Url != expectedFirstArticleUrl {
		t.Errorf("handler returned unexpected firt article url: got %v want %v",
		response.Articles[0].Url, expectedFirstArticleUrl)
	}

	if response.Articles[0].UrlToImage != expectedFirstArticleUrlToImage {
		t.Errorf("handler returned unexpected firt article urlToImage: got %v want %v",
		response.Articles[0].UrlToImage, expectedFirstArticleUrlToImage)
	}

	if response.Articles[0].PublishedAt.String() != expectedPublishedAt {
		t.Errorf("handler returned unexpected firt article publishedAt: got %v want %v",
		response.Articles[0].PublishedAt, expectedPublishedAt)
	}

	if response.Articles[0].Content == "" {
		t.Errorf("handler returned empty content")
	}

	articleLength := len(response.Articles)

	if response.TotalResults != articleLength {
		t.Errorf("handler returned unexpected number of articles: got %v want %v",
		response.TotalResults, articleLength)
	}	
}

func TestGetTopHeadlines_StandardRequestWithWithSource_ReturnsParametersIncompatible(t *testing.T) {

	//Arrange
	expectedResponse := `{
		"status": "error",
		"code": "parametersIncompatible",
		"message": "You cannot mix the sources parameter with the country or category parameters."
	}`
	expectedRequest := `/v2/top-headlines?sources=techcrunch,cnn&language=en&country=us&page=1&pageSize=1`
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		request := req.URL.Path+"?"+req.URL.RawQuery

		if request == expectedRequest {
			io.WriteString(w, expectedResponse)
		} else {
			io.WriteString(w, "Bad request")
		}
	}))

	defer ts.Close()

	expectedStatus := "error"
	expectedCode := constants.ParametersIncompatible
	expectedMessage := "You cannot mix the sources parameter with the country or category parameters."
	
	topHeadlinesReq := models.TopHeadlinesRequest{
		Country:  constants.US,
		Language: "EN",
		Page:     1,
		PageSize: 1,
		Sources:  []string{"techcrunch", "cnn"},
	}

	newsClient := NewClient(ts.URL, "testKey")

	// Act
	response, err := newsClient.GetTopHeadlines(topHeadlinesReq)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Assert
	
	if response.Status != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Error.Status, expectedStatus)
	}

	if response.Error.Code != string(expectedCode) {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Error.Code, expectedCode)
	}

	if response.Error.Message != expectedMessage {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Error.Message, expectedMessage)
	}
}


func TestGetTopHeadlines_StandardRequestWithWithSource_ReturnsApiKeyMissing(t *testing.T) {

	//Arrange
	expectedResponse := `{
		"status": "error",
		"code": "apiKeyMissing",
		"message": "Your API key is missing. Append this to the URL with the apiKey param, or use the x-api-key HTTP header."
	}`
	expectedRequest := `/v2/top-headlines?sources=techcrunch,cnn&language=en&country=us&page=1&pageSize=1`
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		request := req.URL.Path+"?"+req.URL.RawQuery

		if request == expectedRequest {
			io.WriteString(w, expectedResponse)
		} else {
			io.WriteString(w, "Bad request")
		}
	}))

	defer ts.Close()

	expectedStatus := "error"
	expectedCode := constants.ApiKeyMissing
	expectedMessage := "Your API key is missing. Append this to the URL with the apiKey param, or use the x-api-key HTTP header."
	
	topHeadlinesReq := models.TopHeadlinesRequest{
		Country:  constants.US,
		Language: "EN",
		Page:     1,
		PageSize: 1,
		Sources:  []string{"techcrunch", "cnn"},
	}

	newsClient := NewClient(ts.URL, "testKey")

	// Act
	response, err := newsClient.GetTopHeadlines(topHeadlinesReq)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Assert
	
	if response.Status != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Error.Status, expectedStatus)
	}

	if response.Error.Code != string(expectedCode) {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Error.Code, expectedCode)
	}

	if response.Error.Message != expectedMessage {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Error.Message, expectedMessage)
	}
}

// GET EVERYTHING REQUEST TESTS

func TestGetEverything_StandardRequest_ReturnsSuccessResponse(t *testing.T) {

	//Arrange
	expectedCorrectResponse := `{"status": "ok","totalResults": 7904,"articles": [{"source": {"id": "techcrunch","name": "TechCrunch"},"author": "Jake Bright","title": "アフリカのテックニュースまとめ読み：UAVによる救命医療用品の配送など","description": "2020年5月におきた各種イベントは、アフリカがグローバルに応用できるテクノロジーを育むことができるという主張を支持するものとなった。\r\n\r\nアフリカ大陸でビジネスモデルを開発した2つのスタートアップ、MallforAfricaとZiplineが国際的な注目を浴びている。\r\n\r\nDHLは、ナイジェリアのデジタルリテール・スタートアップ、MallforAfrica.comから成長したターンキーのeコマース企業である<a target=\"_blank\" href=\"https://linkcommerce.com/\"…","url": "https://jp.techcrunch.com/2020/06/10/2020-06-01-africa-roundup-dhl-invests-in-mallforafrica-zipline-launches-in-us-novastar-raises-200m/","urlToImage": "https://techcrunchjp.files.wordpress.com/2020/06/202205hh1017-1.jpg?w=1024","publishedAt": "2020-06-09T23:00:29Z","content": "20205\r\n2MallforAfricaZipline\r\nDHLMallforAfrica.comeLink Commerce\r\nLink Commerce\r\nChris Folayan2011MallforAfrica\r\nMallforAfrica\r\ne25030\r\nLink CommerceMallforAfrica.comDHL\r\nZiplineNovant Health\r\nZiplin… [+551 chars]"}]}`
	expectedCorrectRequest := `/v2/everything?sources=techcrunch,cnn&sortBy=popularity&page=1&pageSize=1`
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		request := req.URL.Path+"?"+req.URL.RawQuery

		if request == expectedCorrectRequest {
			io.WriteString(w, expectedCorrectResponse)
		} else {
			io.WriteString(w, "Bad request")
		}
	}))

	defer ts.Close()

	expectedStatus := "ok"
	expectedTotalResults := 7904
	expectedFirstArticleSourceId := "techcrunch"
	expectedFirstArticleSourceName := "TechCrunch"
	expectedFirstArticleAuthor := "Jake Bright"
	expectedFirstArticleTitle := "アフリカのテックニュースまとめ読み：UAVによる救命医療用品の配送など"
	expectedFirstArticleDescription := "2020年5月におきた各種イベントは、アフリカがグローバルに応用できるテクノロジーを育むことができるという主張を支持するものとなった。\r\n\r\nアフリカ大陸でビジネスモデルを開発した2つのスタートアップ、MallforAfricaとZiplineが国際的な注目を浴びている。\r\n\r\nDHLは、ナイジェリアのデジタルリテール・スタートアップ、MallforAfrica.comから成長したターンキーのeコマース企業である<a target=\"_blank\" href=\"https://linkcommerce.com/\"…"

	everythingReq := models.EverythingRequest{
		Page: 1,
		PageSize:  1,
		Sources: []string{"techcrunch","cnn"},
		SortBy: "popularity",
	}

	newsClient := NewClient(ts.URL, "testKey")

	// Act
	response, err := newsClient.GetEverything(everythingReq)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Assert
	if response.Status != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Status, expectedStatus)
	}


	if response.TotalResults != expectedTotalResults {
		t.Errorf("handler returned unexpected total results: got %v want %v",
		response.TotalResults, expectedTotalResults)
	}

	if response.Articles[0].SourceName.Id != expectedFirstArticleSourceId {
		t.Errorf("handler returned unexpected firt source id: got %v want %v",
		response.Articles[0].SourceName.Id, expectedFirstArticleSourceId)
	}

	if response.Articles[0].SourceName.Name != expectedFirstArticleSourceName {
		t.Errorf("handler returned unexpected firt source name: got %v want %v",
		response.Articles[0].SourceName.Name, expectedFirstArticleSourceName)
	}

	if response.Articles[0].Author != expectedFirstArticleAuthor {
		t.Errorf("handler returned unexpected firt article author: got %v want %v",
		response.Articles[0].Author, expectedFirstArticleAuthor)
	}

	if response.Articles[0].Title != expectedFirstArticleTitle {
		t.Errorf("handler returned unexpected firt article title: got %v want %v",
		response.Articles[0].Title, expectedFirstArticleTitle)
	}

	if response.Articles[0].Description != expectedFirstArticleDescription {
		t.Errorf("handler returned unexpected firt article description: got %v want %v",
		response.Articles[0].Description, expectedFirstArticleDescription)
	}

	if response.Error.Status != "ok"{
		t.Errorf("there is an error status, but it shouldn't be there")
	}
}

func TestGetEverything_StandardRequestWithTimeStamps_ReturnsSuccessResponse(t *testing.T) {

	//Arrange
	expectedCorrectResponse := `{"status": "ok","totalResults": 7904,"articles": [{"source": {"id": "techcrunch","name": "TechCrunch"},"author": "Jake Bright","title": "アフリカのテックニュースまとめ読み：UAVによる救命医療用品の配送など","description": "2020年5月におきた各種イベントは、アフリカがグローバルに応用できるテクノロジーを育むことができるという主張を支持するものとなった。\r\n\r\nアフリカ大陸でビジネスモデルを開発した2つのスタートアップ、MallforAfricaとZiplineが国際的な注目を浴びている。\r\n\r\nDHLは、ナイジェリアのデジタルリテール・スタートアップ、MallforAfrica.comから成長したターンキーのeコマース企業である<a target=\"_blank\" href=\"https://linkcommerce.com/\"…","url": "https://jp.techcrunch.com/2020/06/10/2020-06-01-africa-roundup-dhl-invests-in-mallforafrica-zipline-launches-in-us-novastar-raises-200m/","urlToImage": "https://techcrunchjp.files.wordpress.com/2020/06/202205hh1017-1.jpg?w=1024","publishedAt": "2020-06-09T23:00:29Z","content": "20205\r\n2MallforAfricaZipline\r\nDHLMallforAfrica.comeLink Commerce\r\nLink Commerce\r\nChris Folayan2011MallforAfrica\r\nMallforAfrica\r\ne25030\r\nLink CommerceMallforAfrica.comDHL\r\nZiplineNovant Health\r\nZiplin… [+551 chars]"}]}`
	expectedCorrectRequest := `/v2/everything?sources=techcrunch,cnn&from=2020-06-08&to=2020-06-08&sortBy=popularity&page=1&pageSize=1`
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		request := req.URL.Path+"?"+req.URL.RawQuery

		if request == expectedCorrectRequest {
			io.WriteString(w, expectedCorrectResponse)
		} else {
			io.WriteString(w, "Bad request")
		}
	}))

	defer ts.Close()

	expectedStatus := "ok"
	expectedTotalResults := 7904
	expectedFirstArticleSourceId := "techcrunch"
	expectedFirstArticleSourceName := "TechCrunch"
	expectedFirstArticleAuthor := "Jake Bright"
	expectedFirstArticleTitle := "アフリカのテックニュースまとめ読み：UAVによる救命医療用品の配送など"
	expectedFirstArticleDescription := "2020年5月におきた各種イベントは、アフリカがグローバルに応用できるテクノロジーを育むことができるという主張を支持するものとなった。\r\n\r\nアフリカ大陸でビジネスモデルを開発した2つのスタートアップ、MallforAfricaとZiplineが国際的な注目を浴びている。\r\n\r\nDHLは、ナイジェリアのデジタルリテール・スタートアップ、MallforAfrica.comから成長したターンキーのeコマース企業である<a target=\"_blank\" href=\"https://linkcommerce.com/\"…"

	t1, _ := time.Parse("2006-01-02","2020-06-08")
	t2, _ := time.Parse("2006-01-02","2020-06-08")

	everythingReq := models.EverythingRequest{
		Page: 1,
		PageSize:  1,
		Sources: []string{"techcrunch","cnn"},
		SortBy: "popularity",
		From: &t1,
		To: &t2,
	}

	newsClient := NewClient(ts.URL, "testKey")

	// Act
	// execute client method and save result
	response, err := newsClient.GetEverything(everythingReq)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Assert
	
	if response.Status != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
		response.Status, expectedStatus)
	}


	if response.TotalResults != expectedTotalResults {
		t.Errorf("handler returned unexpected total results: got %v want %v",
		response.TotalResults, expectedTotalResults)
	}

	if response.Articles[0].SourceName.Id != expectedFirstArticleSourceId {
		t.Errorf("handler returned unexpected firt source id: got %v want %v",
		response.Articles[0].SourceName.Id, expectedFirstArticleSourceId)
	}

	if response.Articles[0].SourceName.Name != expectedFirstArticleSourceName {
		t.Errorf("handler returned unexpected firt source name: got %v want %v",
		response.Articles[0].SourceName.Name, expectedFirstArticleSourceName)
	}

	if response.Articles[0].Author != expectedFirstArticleAuthor {
		t.Errorf("handler returned unexpected firt article author: got %v want %v",
		response.Articles[0].Author, expectedFirstArticleAuthor)
	}

	if response.Articles[0].Title != expectedFirstArticleTitle {
		t.Errorf("handler returned unexpected firt article title: got %v want %v",
		response.Articles[0].Title, expectedFirstArticleTitle)
	}

	if response.Articles[0].Description != expectedFirstArticleDescription {
		t.Errorf("handler returned unexpected firt article description: got %v want %v",
		response.Articles[0].Description, expectedFirstArticleDescription)
	}

	if response.Error.Status != "ok"{
		t.Errorf("there is an error status, but it shouldn't be there")
	}
}