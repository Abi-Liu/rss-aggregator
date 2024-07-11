package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Abi-Liu/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Atom    string   `xml:"atom,attr"`
	Version string   `xml:"version,attr"`
	Link    []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Rel  string `xml:"rel,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"link"`
	Style []struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
		Type string `xml:"type,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"style"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description   string `xml:"description"`
		Generator     string `xml:"generator"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
			Guid        string `xml:"guid"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func startScraping(db *database.Queries, n int, sleepTime time.Duration) {
	ticker := time.NewTicker(sleepTime)
	for range ticker.C {
		feedsToFetch, err := db.GetNextFeedsToFetch(context.Background(), int32(n))
		if err != nil {
			log.Printf("failed to get next feeds: %s", err)
			continue
		}

		log.Printf("Found %v feeds to fetch", len(feedsToFetch))

		waitGroup := &sync.WaitGroup{}
		for _, feed := range feedsToFetch {
			waitGroup.Add(1)
			go fetchXML(feed, waitGroup, db)
		}

		waitGroup.Wait()
	}
}

func fetchXML(feed database.Feed, wg *sync.WaitGroup, db *database.Queries) {
	defer wg.Done()
	res, err := http.Get(feed.Url)

	if err != nil {
		log.Printf("Failed GET: %s", feed.Url)
		return
	}

	rss, err := parseXML(res)

	if err != nil {
		log.Printf("Failed to parse xml: %v", err)
		return
	}

	for _, post := range rss.Channel.Item {
		nullTime := sql.NullTime{}
		time, err := time.Parse(time.RFC1123Z, post.PubDate)
		if err == nil {
			nullTime.Time = time
			nullTime.Valid = true
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:    uuid.New(),
			Title: post.Title,
			Url:   post.Link,
			Description: sql.NullString{
				String: post.Description,
				Valid:  true,
			},
			PublishedAt: nullTime,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "violates unique constraint") {
				continue
			}
			log.Printf("Failed to create post: %v", err)
			continue
		}
	}

	_, err = db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Failed to mark feed as fetched. ID: %v", feed.ID)
		return
	}
}

func parseXML(r *http.Response) (*Rss, error) {
	data, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return &Rss{}, err
	}

	rss := &Rss{}
	err = xml.Unmarshal(data, rss)

	if err != nil {
		return &Rss{}, err
	}

	return rss, nil
}
