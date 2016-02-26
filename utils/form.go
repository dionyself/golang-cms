package utils

type ArticleForm struct {
    Id    int         `form:"-"`
    Title  string `form:"title"`
    Content   string         `form:"content"`
    TopicTags string    `form:"topic-tags"`
    TaggedUsers string   `form:"tagged-users"`
    AllowReviews bool    `form:"allow-reviews"`
    AllowComments bool  `form:"allow-comments"`
    Errors map[string]string
}

func (this *ArticleForm) Validate() error{
	return nil
}