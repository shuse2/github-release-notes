package main

import (
	"fmt"
	"github.com/shuse2/github-release-notes/githubber"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghrn"
	app.Usage = "Generate Release notes"
	app.Version = "0.1.0"
	commonFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Usage: "User/Organization of the repository",
		},
		cli.StringFlag{
			Name:  "repo, r",
			Usage: "Repository name",
		},
		cli.StringFlag{
			Name:  "token, t",
			Usage: "Token to use",
		},
		cli.StringFlag{
			Name:  "tag",
			Usage: "Tag of release note creating for",
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "project",
			Action: getByProject,
			Flags:  commonFlags[:],
		},
		cli.Command{
			Name:   "branch",
			Action: getByBranch,
			Flags:  commonFlags[:],
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func getByProject(c *cli.Context) error {
	token := c.String("token")
	client := githubber.NewFetcher(token)
	user := c.String("user")
	repo := c.String("repo")
	version := c.String("tag")
	projectStr := c.Args().First()
	if projectStr == "" {
		return fmt.Errorf("Project board nuber must be specified")
	}
	project, err := strconv.Atoi(projectStr)
	if err != nil {
		return err
	}
	query := &githubber.ProjectQuery{
		Organization: user,
		Repo:         repo,
		Project:      project,
	}
	items, err := client.Search(query)
	if err != nil {
		return err
	}
	if err := githubber.SaveChangeLog(version, items); err != nil {
		return err
	}
	return nil
}

func getByBranch(c *cli.Context) error {
	token := c.String("token")
	client := githubber.NewFetcher(token)
	user := c.String("user")
	repo := c.String("repo")
	version := c.String("tag")
	branch := c.Args().First()
	query := &githubber.BranchQuery{
		Organization: user,
		Repo:         repo,
		Branch:       branch,
	}
	items, err := client.Search(query)
	if err != nil {
		return err
	}
	issueNumbers := githubber.GithubItems(items).GetRelatedIssueNumber()
	ids := make([]string, len(issueNumbers))
	for i, iNum := range issueNumbers {
		ids[i] = strings.TrimLeft(iNum, "#")
	}
	issues, err := client.GetIssues(user, repo, ids)
	if err != nil {
		return err
	}
	if len(issues) > 0 {
		items = append(items, issues...)
	}
	if err := githubber.SaveChangeLog(version, items); err != nil {
		return err
	}
	return nil
}
