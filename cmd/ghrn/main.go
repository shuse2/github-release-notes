package main

import (
	"fmt"
	"github.com/shuse2/github-release-notes/githubber"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghrn"
	app.Usage = "Generate Release notes"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "project",
			Action: getByProject,
			Flags: []cli.Flag{
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
			},
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
	query := githubber.IssueSearchQuery{
		Organization: user,
		Repo:         repo,
		Project:      project,
	}
	items, err := client.GetIssuesAndPRs(query)
	if err != nil {
		return err
	}
	if err := githubber.SaveChangeLog(version, items); err != nil {
		return err
	}
	return nil
}
