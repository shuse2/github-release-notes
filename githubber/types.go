package githubber

import (
	"time"
)

type GithubIssueSearchResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []GithubItem `json:"items"`
}

type GithubItems []GithubItem

type GithubItem struct {
	URL               string             `json:"url"`
	RepositoryURL     string             `json:"repository_url"`
	LabelsURL         string             `json:"labels_url"`
	CommentsURL       string             `json:"comments_url"`
	EventsURL         string             `json:"events_url"`
	HTMLURL           string             `json:"html_url"`
	ID                int                `json:"id"`
	NodeID            string             `json:"node_id"`
	Number            int                `json:"number"`
	Title             string             `json:"title"`
	User              GithubUser         `json:"user"`
	Labels            []GithubLabel      `json:"labels"`
	State             string             `json:"state"`
	Locked            bool               `json:"locked"`
	Assignee          GithubUser         `json:"assignee"`
	Assignees         []GithubUser       `json:"assignees"`
	Milestone         *GithubMilestone   `json:"milestone"`
	Comments          int                `json:"comments"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	ClosedAt          time.Time          `json:"closed_at"`
	AuthorAssociation string             `json:"author_association"`
	PullRequest       *GithubPullRequest `json:"pull_request"`
	Body              string             `json:"body"`
	Score             float64            `json:"score"`
}

type GithubUser struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type GithubLabel struct {
	ID      int    `json:"id"`
	NodeID  string `json:"node_id"`
	URL     string `json:"url"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Default bool   `json:"default"`
}

type GithubPullRequest struct {
	URL      string `json:"url"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
}

type GithubMilestone struct {
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	LabelsURL   string `json:"labels_url"`
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	Number      int    `json:"number"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"creator"`
}
