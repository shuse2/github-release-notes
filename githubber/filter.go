package githubber

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
