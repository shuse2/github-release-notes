package githubber

import (
	"strconv"
)

type ProjectQuery struct {
	Organization string
	Repo         string
	Project      int
	Page         int
}

func (q ProjectQuery) Query() string {
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

func (q *ProjectQuery) UpdatePage(page int) {
	q.Page = page
}

type BranchQuery struct {
	Organization string
	Repo         string
	Page         int
	Branch       string
}

func (q BranchQuery) Query() string {
	query := "?q="
	repo := q.Organization + "/" + q.Repo
	query += "repo:" + repo
	if q.Branch != "" {
		query += "+base:" + q.Branch
	}
	if q.Page != 0 {
		query += "&page=" + strconv.Itoa(q.Page)
	}
	return query
}

func (q *BranchQuery) UpdatePage(page int) {
	q.Page = page
}
