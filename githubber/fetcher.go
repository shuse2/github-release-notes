package githubber

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"strconv"
)

const (
	APIBaseUrl         = "https://api.github.com"
	APISearchIssueUrl  = "/search/issues"
	HeaderKeyAccept    = "Accept"
	HeaderValuePreview = "application/vnd.github.mercy-preview+json"
	MaxQueryCount      = 30
)

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

type IssueSearchQuery struct {
	Organization string
	Repo         string
	Project      int
	Page         int
}

func (q IssueSearchQuery) Query() string {
	query := "?q="
	repo := q.Organization + "/" + q.Repo
	query += "repo:" + repo
	if q.Project != 0 {
		query += "+project:" + repo + "/" + strconv.Itoa(q.Project)
	}
	if q.Page != 0 {
		query += "&page=" + strconv.Itoa(q.Page)
	}
	return query
}

func (f *Fetcher) GetIssuesAndPRs(query IssueSearchQuery) ([]GithubItem, error) {
	url := APIBaseUrl + APISearchIssueUrl + query.Query()
	resp, err := f.fetch(url)
	if err != nil {
		return nil, err
	}
	if resp.TotalCount <= MaxQueryCount {
		return resp.Items, nil
	}
	items := resp.Items
	lastPage := resp.TotalCount/MaxQueryCount + 1
	for p := 2; p <= lastPage; p++ {
		query.Page = p
		next := APIBaseUrl + APISearchIssueUrl + query.Query()
		remaining, err := f.fetch(next)
		if err != nil {
			return nil, err
		}
		items = append(items, remaining.Items...)
	}
	return items, nil
}

func (f *Fetcher) fetch(url string) (*GithubIssueSearchResponse, error) {
	fmt.Printf("Fetching %s \n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(HeaderKeyAccept, HeaderValuePreview)
	res, err := f.Do(req)
	if err != nil {
		return nil, err
	}
	if res.Body == nil {
		return nil, fmt.Errorf("Response body from %s was nil", url)
	}
	defer res.Body.Close()
	body := &GithubIssueSearchResponse{}
	if err := json.NewDecoder(res.Body).Decode(body); err != nil {
		return nil, err
	}
	return body, nil
}
