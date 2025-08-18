package fetch

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aenlemmea/garnet/back/internal/data"
)

type newsAPIResponse struct {
	Articles []gnewsTubeArticle `json:"articles"`
}

type gnewsTubeArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Source      struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"source"`
}

type NewsAPIFetcher struct {
	aggregatorStore data.AggregatorStore
	logger          *log.Logger
}

func CreateNewsAPIFetcher(aggregatorStore data.AggregatorStore, logger *log.Logger) *NewsAPIFetcher {
	return &NewsAPIFetcher{
		aggregatorStore: aggregatorStore,
		logger:          logger,
	}
}

type NewsFetcher interface {
	StartFetch() error
}

func (nf *NewsAPIFetcher) fetch() []*data.Aggregator {
	client := http.Client{Timeout: 120 * time.Second}
	// TODO Use params encoder to replace this URL style.
	url := "https://gnews.io/api/v4/top-headlines?q=agriculture&country=in&apikey=" + os.Getenv("GNEWS")
	resp, err := client.Get(url)
	if err != nil {
		nf.logger.Fatalf("📰 Fetch: Error %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		nf.logger.Fatalln("📰 Fetch: Status NOT Ok")
	}

	var apiResp newsAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		nf.logger.Fatalln("📰 Fetch: JSON Decode Failed")
	}

	var aggs []*data.Aggregator
	for _, article := range apiResp.Articles {
		agg := &data.Aggregator{
			Title:      article.Title,
			Blurb:      article.Description,
			Link:       article.URL,
			OriginName: article.Source.Name,
			Tags:       []string{"agriculture"}, //TODO
		}
		aggs = append(aggs, agg)
	}

	return aggs
}

func (nf *NewsAPIFetcher) StartFetch() error {
	aggs := nf.fetch()

	for _, agg := range aggs {
		_, err := nf.aggregatorStore.PopulateAggregator(agg)
		if err != nil {
			nf.logger.Printf("🛢️ Failed Populating Aggregator")
			return err
		}
	}
	return nil
}
