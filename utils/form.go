package utils
import (
	"github.com/dionyself/golang-cms/models"
)

type ArticleForm struct {
    Id    int         `form:"-"`
    Title  string `form:"title"`
    Category int  `form:"category"`
    Content   string         `form:"content"`
    TopicTags string    `form:"topic-tags"`
    TaggedUsers string   `form:"tagged-users"`
    AllowReviews bool    `form:"allow-reviews"`
    AllowComments bool  `form:"allow-comments"`
    Errors map[string]string
}

func (this *ArticleForm) Validate(Art **models.Article) error{
	return nil
}