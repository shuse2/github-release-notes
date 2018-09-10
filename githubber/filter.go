package githubber

import (
	"regexp"
)

func (items GithubItems) GetIssues() []GithubItem {
	issues := []GithubItem{}
	for _, item := range items {
		if item.PullRequest == nil {
			issues = append(issues, item)
		}
	}
	return issues
}

func (items GithubItems) GetPRs() []GithubItem {
	prs := []GithubItem{}
	for _, item := range items {
		if item.PullRequest != nil {
			prs = append(prs, item)
		}
	}
	return prs
}

func (items GithubItems) GetRelatedIssueNumber() []string {
	keys := map[string]bool{}
	re := regexp.MustCompile("#[0-9]*")
	for _, item := range items {
		searchStr := item.Title + item.Body
		issueNums := re.FindAllString(searchStr, -1)
		for _, num := range issueNums {
			keys[num] = true
		}
	}
	result := []string{}
	for key, _ := range keys {
		if key != "" && key != "#" {
			result = append(result, key)
		}
	}
	return result
}
