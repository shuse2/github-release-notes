package githubber

import (
	"os"
	"text/template"
	"time"
)

const FileLogName = "CHANGELOG.md"

const ChangeLogTemplate = `
# Change Log

## [{{.Version}}](https://github.com/LiskHQ/lisk-commander/tree/{{.Version}}) ({{.Date}})
[Full Changelog](https://github.com/LiskHQ/lisk-commander/compare/{{.Version}}...HEAD)

**Closed issues:**
{{range .ClosedIssues}}- {{.Title}} [#{{.Number}}]({{.HTMLURL}})
{{end}}

**Merged pull requests:**
{{range .ClosedPRs}}- {{.Title}} [#{{.Number}}]({{.PullRequest.HTMLURL}}) ([{{.User.Login}}]({{.User.HTMLURL}}))
{{end}}
`

type ChangeLogInfo struct {
	Version      string
	Date         string
	ClosedIssues []GithubItem
	ClosedPRs    []GithubItem
}

func SaveChangeLog(version string, items []GithubItem) error {
	filename := "./" + FileLogName
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	prs := GithubItems(items).GetPRs()
	issues := GithubItems(items).GetIssues()
	now := time.Now().Format("2006-01-02")
	changeLog := &ChangeLogInfo{
		Version:      version,
		ClosedIssues: issues,
		ClosedPRs:    prs,
		Date:         now,
	}
	t := template.Must(template.New("changeLog").Parse(ChangeLogTemplate))
	t.Execute(f, changeLog)
	return nil
}
