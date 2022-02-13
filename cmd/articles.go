package cmd

import (
	"fmt"

	"github.com/sheepla/fzenn/client"
	"github.com/spf13/cobra"
)

var articlesCmd = &cobra.Command{
	Use:   "articles",
	Short: "Search articles",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
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

	for _, v := range a.Articles {
		fmt.Printf("%s %s\n [%s] by %s (@%s)\n ♥ %d, %d minutes to read\n\n",
			v.Emoji,
			v.Title,
			v.ArticleType,
			v.User.Name,
			v.User.Username,
			v.ReadingTime,
			v.LikedCount,
		)
	}
	return nil
}
