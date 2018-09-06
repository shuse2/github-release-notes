package fetcher

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
}

func (q IssueSearchQuery) Query() string {
	query := "?q="
	repo := q.Organization + "/" + q.Repo
	query += "repo:" + repo
	if q.Project != 0 {
		query += "+project:" + repo + "/" + strconv.Itoa(q.Project)
	}
	return query
}

func (f *Fetcher) GetIssuesAndPRs(query IssueSearchQuery) ([]GithubItem, error) {
	url := APIBaseUrl + APISearchIssueUrl + query.Query()
	fmt.Println(url)
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
	return body.Items, nil
}
