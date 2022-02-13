package cmd

import (
	"fmt"
	"path"

	humanize "github.com/dustin/go-humanize"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/sheepla/fzenn/client"
	"github.com/spf13/cobra"
	"github.com/toqueteos/webbrowser"
)

const (
	baseURL = "https://zenn.dev"
)

var articlesCmd = &cobra.Command{
	Use:   "articles",
	Short: "Search articles",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("too many arguments")
		}
		return articlesMain()
	},
}

func init() {
	rootCmd.AddCommand(articlesCmd)
}

func articlesMain() error {
	a, err := client.FetchArticleList()
	if err != nil {
		return err
	}

	choises, err := find(a)
	if err != nil {
		return err
	}

	for _, idx := range choises {
		url := newPageURL(a.Articles[idx].User.Username, a.Articles[idx].Slug)
		fmt.Println(url)
		open(url)
	}

	return nil
}

func find(a *client.ArticleList) ([]int, error) {
	return fuzzyfinder.FindMulti(
		a.Articles,
		func(i int) string {
			if i == -1 {
				return ""
			}
			return a.Articles[i].Title
		},
		fuzzyfinder.WithPreviewWindow(
			func(i, w, h int) string {
				if i == -1 {
					return ""
				}
				return preview(i, w, h, a)
			},
		),
	)
}

func preview(i, w, h int, a *client.ArticleList) string {
	return fmt.Sprintf(
		"[%s/%s]\n\n%s %s\n\n  by %s (@%s)\n\n  ♥  %d 💬 %d\n  %d min to read\n\n  created at %s, updated at %s\n\n",
		a.Articles[i].ArticleType,
		a.Articles[i].PostType,
		a.Articles[i].Emoji,
		a.Articles[i].Title,
		a.Articles[i].User.Name,
		a.Articles[i].User.Username,
		a.Articles[i].LikedCount,
		a.Articles[i].CommentsCount,
		a.Articles[i].ReadingTime,
		humanize.Time(a.Articles[i].CreatedAt),
		humanize.Time(a.Articles[i].UpdatedAt),
	)
}

func newPageURL(username, slug string) string {
	return path.Join(baseURL, username, "articles", slug)
}

func open(url string) error {
	return webbrowser.Open(url)
}
