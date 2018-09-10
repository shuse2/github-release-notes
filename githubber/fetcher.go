package githubber

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"sync"
)

const (
	APIBaseUrl         = "https://api.github.com"
	APISearchIssueUrl  = "/search/issues"
	HeaderKeyAccept    = "Accept"
	HeaderValuePreview = "application/vnd.github.mercy-preview+json"
	MaxQueryCount      = 30
)

type Querier interface {
	Query() string
	UpdatePage(int)
}

type Fetcher struct {
	http.Client
}

func NewFetcher(token string) *Fetcher {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	client := oauth2.NewClient(ctx, ts)
	return &Fetcher{*client}
}

func (f *Fetcher) Search(query Querier) ([]GithubItem, error) {
	url := APIBaseUrl + APISearchIssueUrl + query.Query()
	resp := &GithubIssueSearchResponse{}
	if err := f.fetch(url, resp); err != nil {
		return nil, err
	}
	if resp.TotalCount <= MaxQueryCount {
		return resp.Items, nil
	}
	items := resp.Items
	lastPage := resp.TotalCount/MaxQueryCount + 1
	for p := 2; p <= lastPage; p++ {
		query.UpdatePage(p)
		next := APIBaseUrl + APISearchIssueUrl + query.Query()
		remaining := &GithubIssueSearchResponse{}
		if err := f.fetch(next, remaining); err != nil {
			return nil, err
		}
		items = append(items, remaining.Items...)
	}
	return items, nil
}

func (f *Fetcher) GetIssues(owner, repo string, ids []string) ([]GithubItem, error) {
	wg := sync.WaitGroup{}
	url := APIBaseUrl + "/repos/" + owner + "/" + repo + "/issues/"
	resp := make(chan GithubItem)
	errChan := make(chan error)
	wg.Add(len(ids))
	for _, id := range ids {
		go func(id string) {
			defer wg.Done()
			fmt.Printf("Fetching %s \n", url+id)
			req, err := http.NewRequest("GET", url+id, nil)
			if err != nil {
				errChan <- err
				return
			}
			req.Header.Set(HeaderKeyAccept, HeaderValuePreview)
			res, err := f.Do(req)
			if err != nil {
				errChan <- err
				return
			}
			defer res.Body.Close()
			body := GithubItem{}
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				errChan <- err
				return
			}
			resp <- body
		}(id)
	}
	items := []GithubItem{}
	go func() {
		for res := range resp {
			items = append(items, res)
		}
		for err := range errChan {
			fmt.Println(err)
		}
	}()
	wg.Wait()
	return items, nil
}

func (f *Fetcher) fetch(url string, body interface{}) error {
	fmt.Printf("Fetching %s \n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set(HeaderKeyAccept, HeaderValuePreview)
	res, err := f.Do(req)
	if err != nil {
		return err
	}
	if res.Body == nil {
		return fmt.Errorf("Response body from %s was nil", url)
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(body); err != nil {
		return err
	}
	return nil
}
