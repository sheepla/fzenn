package client

import (
    "time"
)

type Trending struct {
	DailyBooks []struct {
		CoverImageSmallURL  string `json:"coverImageSmallUrl"`
		CreatedAt           string `json:"createdAt"`
		ID                  int64  `json:"id"`
		IsSuspendingPrivate bool   `json:"isSuspendingPrivate"`
		LikedCount          int64  `json:"likedCount"`
		PostType            string `json:"postType"`
		Price               int64  `json:"price"`
		Published           bool   `json:"published"`
		PublishedAt         string `json:"publishedAt"`
		Slug                string `json:"slug"`
		SourceRepoUpdatedAt string `json:"sourceRepoUpdatedAt"`
		Title               string `json:"title"`
		User                struct {
			AvatarSmallURL string `json:"avatarSmallUrl"`
			ID             int64  `json:"id"`
			Name           string `json:"name"`
			Username       string `json:"username"`
		} `json:"user"`
	} `json:"dailyBooks"`
	DailyIdeaArticles []struct {
		ArticleType         string      `json:"articleType"`
		BodyLettersCount    int64       `json:"bodyLettersCount"`
		BodyUpdatedAt       string      `json:"bodyUpdatedAt"`
		CommentsCount       int64       `json:"commentsCount"`
		CreatedAt           string      `json:"createdAt"`
		Emoji               string      `json:"emoji"`
		ID                  int64       `json:"id"`
		IsSuspendingPrivate bool        `json:"isSuspendingPrivate"`
		LikedCount          int64       `json:"likedCount"`
		PostType            string      `json:"postType"`
		Publication         interface{} `json:"publication"`
		Published           bool        `json:"published"`
		PublishedAt         string      `json:"publishedAt"`
		ReadingTime         int64       `json:"readingTime"`
		Slug                string      `json:"slug"`
		SourceRepoUpdatedAt string      `json:"sourceRepoUpdatedAt"`
		Title               string      `json:"title"`
		Topics              []struct {
			DisplayName   string `json:"displayName"`
			ID            int64  `json:"id"`
			ImageURL      string `json:"imageUrl"`
			Name          string `json:"name"`
			TaggingsCount int64  `json:"taggingsCount"`
		} `json:"topics"`
		UpdatedAt string `json:"updatedAt"`
		User      struct {
			AvatarSmallURL string `json:"avatarSmallUrl"`
			ID             int64  `json:"id"`
			Name           string `json:"name"`
			Username       string `json:"username"`
		} `json:"user"`
	} `json:"dailyIdeaArticles"`
	DailyTechArticles []struct {
		ArticleType         string      `json:"articleType"`
		BodyLettersCount    int64       `json:"bodyLettersCount"`
		BodyUpdatedAt       string      `json:"bodyUpdatedAt"`
		CommentsCount       int64       `json:"commentsCount"`
		CreatedAt           string      `json:"createdAt"`
		Emoji               string      `json:"emoji"`
		ID                  int64       `json:"id"`
		IsSuspendingPrivate bool        `json:"isSuspendingPrivate"`
		LikedCount          int64       `json:"likedCount"`
		PostType            string      `json:"postType"`
		Publication         interface{} `json:"publication"`
		Published           bool        `json:"published"`
		PublishedAt         string      `json:"publishedAt"`
		ReadingTime         int64       `json:"readingTime"`
		Slug                string      `json:"slug"`
		SourceRepoUpdatedAt string      `json:"sourceRepoUpdatedAt"`
		Title               string      `json:"title"`
		Topics              []struct {
			DisplayName   string `json:"displayName"`
			ID            int64  `json:"id"`
			ImageURL      string `json:"imageUrl"`
			Name          string `json:"name"`
			TaggingsCount int64  `json:"taggingsCount"`
		} `json:"topics"`
		UpdatedAt string `json:"updatedAt"`
		User      struct {
			AvatarSmallURL string `json:"avatarSmallUrl"`
			ID             int64  `json:"id"`
			Name           string `json:"name"`
			Username       string `json:"username"`
		} `json:"user"`
	} `json:"dailyTechArticles"`
	FeaturedArticles []struct {
		ArticleType         string      `json:"articleType"`
		BodyLettersCount    int64       `json:"bodyLettersCount"`
		BodyUpdatedAt       string      `json:"bodyUpdatedAt"`
		CommentsCount       int64       `json:"commentsCount"`
		CreatedAt           string      `json:"createdAt"`
		Emoji               string      `json:"emoji"`
		ID                  int64       `json:"id"`
		IsSuspendingPrivate bool        `json:"isSuspendingPrivate"`
		LikedCount          int64       `json:"likedCount"`
		PostType            string      `json:"postType"`
		Publication         interface{} `json:"publication"`
		Published           bool        `json:"published"`
		PublishedAt         string      `json:"publishedAt"`
		ReadingTime         int64       `json:"readingTime"`
		Slug                string      `json:"slug"`
		SourceRepoUpdatedAt string      `json:"sourceRepoUpdatedAt"`
		Title               string      `json:"title"`
		Topics              []struct {
			DisplayName   string `json:"displayName"`
			ID            int64  `json:"id"`
			ImageURL      string `json:"imageUrl"`
			Name          string `json:"name"`
			TaggingsCount int64  `json:"taggingsCount"`
		} `json:"topics"`
		UpdatedAt string `json:"updatedAt"`
		User      struct {
			AvatarSmallURL string `json:"avatarSmallUrl"`
			ID             int64  `json:"id"`
			Name           string `json:"name"`
			Username       string `json:"username"`
		} `json:"user"`
	} `json:"featuredArticles"`
	Query struct{} `json:"query"`
}

type ArticleList struct {
	Articles []struct {
		ArticleType         string      `json:"articleType"`
		BodyLettersCount    int64       `json:"bodyLettersCount"`
		BodyUpdatedAt       time.Time      `json:"bodyUpdatedAt"`
		CommentsCount       int64       `json:"commentsCount"`
		Emoji               string      `json:"emoji"`
		ID                  int64       `json:"id"`
		IsSuspendingPrivate bool        `json:"isSuspendingPrivate"`
		LikedCount          int64       `json:"likedCount"`
		PostType            string      `json:"postType"`
		Publication         interface{} `json:"publication"`
		Published           bool        `json:"published"`
		PublishedAt         time.Time      `json:"publishedAt"`
		ReadingTime         int64       `json:"readingTime"`
		Slug                string      `json:"slug"`
		SourceRepoUpdatedAt time.Time      `json:"sourceRepoUpdatedAt"`
		Title               string      `json:"title"`
		Topics              []struct {
			DisplayName   string `json:"displayName"`
			ID            int64  `json:"id"`
			ImageURL      string `json:"imageUrl"`
			Name          string `json:"name"`
			TaggingsCount int64  `json:"taggingsCount"`
		} `json:"topics"`
		User struct {
			AvatarSmallURL string `json:"avatarSmallUrl"`
			ID             int64  `json:"id"`
			Name           string `json:"name"`
			Username       string `json:"username"`
		} `json:"user"`
	} `json:"articles"`
	CurrentPage int64    `json:"currentPage"`
	NextPage    int64    `json:"nextPage"`
	Query       struct{} `json:"query"`
}

type BookList struct {
	Books []struct {
		CoverImageSmallURL  string `json:"coverImageSmallUrl"`
		CreatedAt           string `json:"createdAt"`
		ID                  int64  `json:"id"`
		IsSuspendingPrivate bool   `json:"isSuspendingPrivate"`
		LikedCount          int64  `json:"likedCount"`
		PostType            string `json:"postType"`
		Price               int64  `json:"price"`
		Published           bool   `json:"published"`
		PublishedAt         string `json:"publishedAt"`
		Slug                string `json:"slug"`
		SourceRepoUpdatedAt string `json:"sourceRepoUpdatedAt"`
		Title               string `json:"title"`
		User                struct {
			AvatarSmallURL string `json:"avatarSmallUrl"`
			ID             int64  `json:"id"`
			Name           string `json:"name"`
			Username       string `json:"username"`
		} `json:"user"`
	} `json:"books"`
	CurrentPage int64    `json:"currentPage"`
	NextPage    int64    `json:"nextPage"`
	Query       struct{} `json:"query"`
}

type User struct {
	ActiveItemType string `json:"activeItemType"`
	Articles       []struct {
		ArticleType         string      `json:"articleType"`
		BodyLettersCount    int64       `json:"bodyLettersCount"`
		BodyUpdatedAt       string      `json:"bodyUpdatedAt"`
		CommentsCount       int64       `json:"commentsCount"`
		CreatedAt           string      `json:"createdAt"`
		Emoji               string      `json:"emoji"`
		ID                  int64       `json:"id"`
		IsSuspendingPrivate bool        `json:"isSuspendingPrivate"`
		LikedCount          int64       `json:"likedCount"`
		PostType            string      `json:"postType"`
		Publication         interface{} `json:"publication"`
		Published           bool        `json:"published"`
		PublishedAt         string      `json:"publishedAt"`
		ReadingTime         int64       `json:"readingTime"`
		Slug                string      `json:"slug"`
		SourceRepoUpdatedAt interface{} `json:"sourceRepoUpdatedAt"`
		Title               string      `json:"title"`
		Topics              []struct {
			DisplayName   string `json:"displayName"`
			ID            int64  `json:"id"`
			ImageURL      string `json:"imageUrl"`
			Name          string `json:"name"`
			TaggingsCount int64  `json:"taggingsCount"`
		} `json:"topics"`
		UpdatedAt string `json:"updatedAt"`
		User      struct {
			AvatarSmallURL string `json:"avatarSmallUrl"`
			ID             int64  `json:"id"`
			Name           string `json:"name"`
			Username       string `json:"username"`
		} `json:"user"`
	} `json:"articles"`
	Query struct {
		Username string `json:"username"`
	} `json:"query"`
	ResUser struct {
		ArticlesCount   int64       `json:"articlesCount"`
		AutolinkedBio   string      `json:"autolinkedBio"`
		AvatarSmallURL  string      `json:"avatarSmallUrl"`
		AvatarURL       string      `json:"avatarUrl"`
		Bio             string      `json:"bio"`
		BooksCount      int64       `json:"booksCount"`
		FollowerCount   int64       `json:"followerCount"`
		FollowingCount  int64       `json:"followingCount"`
		GaTrackingID    interface{} `json:"gaTrackingId"`
		GithubUsername  string      `json:"githubUsername"`
		ID              int64       `json:"id"`
		IsSupportOpen   bool        `json:"isSupportOpen"`
		Name            string      `json:"name"`
		ScrapsCount     int64       `json:"scrapsCount"`
		TokusyoContact  interface{} `json:"tokusyoContact"`
		TokusyoName     interface{} `json:"tokusyoName"`
		TotalLikedCount int64       `json:"totalLikedCount"`
		TwitterUsername string      `json:"twitterUsername"`
		Username        string      `json:"username"`
		WebsiteDomain   string      `json:"websiteDomain"`
		WebsiteURL      string      `json:"websiteUrl"`
	} `json:"resUser"`
	User struct {
		ArticlesCount   int64       `json:"articlesCount"`
		AutolinkedBio   string      `json:"autolinkedBio"`
		AvatarSmallURL  string      `json:"avatarSmallUrl"`
		AvatarURL       string      `json:"avatarUrl"`
		Bio             string      `json:"bio"`
		BooksCount      int64       `json:"booksCount"`
		FollowerCount   int64       `json:"followerCount"`
		FollowingCount  int64       `json:"followingCount"`
		GaTrackingID    interface{} `json:"gaTrackingId"`
		GithubUsername  string      `json:"githubUsername"`
		ID              int64       `json:"id"`
		IsSupportOpen   bool        `json:"isSupportOpen"`
		Name            string      `json:"name"`
		ScrapsCount     int64       `json:"scrapsCount"`
		TokusyoContact  interface{} `json:"tokusyoContact"`
		TokusyoName     interface{} `json:"tokusyoName"`
		TotalLikedCount int64       `json:"totalLikedCount"`
		TwitterUsername string      `json:"twitterUsername"`
		Username        string      `json:"username"`
		WebsiteDomain   string      `json:"websiteDomain"`
		WebsiteURL      string      `json:"websiteUrl"`
	} `json:"user"`
}

type TopicsList struct {
	Query  struct{} `json:"query"`
	Topics []struct {
		DisplayName   string `json:"displayName"`
		ID            int64  `json:"id"`
		ImageURL      string `json:"imageUrl"`
		Name          string `json:"name"`
		TaggingsCount int64  `json:"taggingsCount"`
	} `json:"topics"`
}

type Topics struct {
	ActiveItemType string `json:"activeItemType"`
	Articles       []struct {
		ArticleType         string      `json:"articleType"`
		BodyLettersCount    int64       `json:"bodyLettersCount"`
		BodyUpdatedAt       string      `json:"bodyUpdatedAt"`
		CommentsCount       int64       `json:"commentsCount"`
		CreatedAt           string      `json:"createdAt"`
		Emoji               string      `json:"emoji"`
		ID                  int64       `json:"id"`
		IsSuspendingPrivate bool        `json:"isSuspendingPrivate"`
		LikedCount          int64       `json:"likedCount"`
		PostType            string      `json:"postType"`
		Publication         interface{} `json:"publication"`
		Published           bool        `json:"published"`
		PublishedAt         string      `json:"publishedAt"`
		ReadingTime         int64       `json:"readingTime"`
		Slug                string      `json:"slug"`
		SourceRepoUpdatedAt string      `json:"sourceRepoUpdatedAt"`
		Title               string      `json:"title"`
		Topics              []struct {
			DisplayName   string `json:"displayName"`
			ID            int64  `json:"id"`
			ImageURL      string `json:"imageUrl"`
			Name          string `json:"name"`
			TaggingsCount int64  `json:"taggingsCount"`
		} `json:"topics"`
		UpdatedAt string `json:"updatedAt"`
		User      struct {
			AvatarSmallURL string `json:"avatarSmallUrl"`
			ID             int64  `json:"id"`
			Name           string `json:"name"`
			Username       string `json:"username"`
		} `json:"user"`
	} `json:"articles"`
	CurrentPage int64 `json:"currentPage"`
	NextPage    int64 `json:"nextPage"`
	Query       struct {
		Name string `json:"name"`
	} `json:"query"`
	ResTopic struct {
		ArticlesCount int64  `json:"articlesCount"`
		BooksCount    int64  `json:"booksCount"`
		DisplayName   string `json:"displayName"`
		ID            int64  `json:"id"`
		ImageURL      string `json:"imageUrl"`
		Name          string `json:"name"`
		ScrapsCount   int64  `json:"scrapsCount"`
		TaggingsCount int64  `json:"taggingsCount"`
	} `json:"resTopic"`
	Topic struct {
		ArticlesCount int64  `json:"articlesCount"`
		BooksCount    int64  `json:"booksCount"`
		DisplayName   string `json:"displayName"`
		ID            int64  `json:"id"`
		ImageURL      string `json:"imageUrl"`
		Name          string `json:"name"`
		ScrapsCount   int64  `json:"scrapsCount"`
		TaggingsCount int64  `json:"taggingsCount"`
	} `json:"topic"`
}
