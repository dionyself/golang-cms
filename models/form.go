package models

// RegisterForm ...
type RegisterForm struct {
	Name          string            `form:"name" valid:"Required;"`
	Email         string            `form:"email" valid:"Required;"`
	Username      string            `form:"username" valid:"Required;AlphaNumeric;MinSize(4);MaxSize(300)"`
	Password      string            `form:"password" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe    string            `form:"passwordre" valid:"Required;MinSize(4);MaxSize(30)"`
	Gender        bool              `form:"gender" valid:"Required"`
	InvalidFields map[string]string `form:"-"`
}

// ArticleForm ...
type ArticleForm struct {
	Id            int               `form:"-"`
	Title         string            `form:"title" valid:"Required;MinSize(4);MaxSize(300)"`
	Category      int               `form:"category"`
	Content       string            `form:"content" valid:"Required;MinSize(50);MaxSize(2000)"`
	TopicTags     string            `form:"topic-tags" valid:"MinSize(4);MaxSize(300)"`
	TaggedUsers   string            `form:"tagged-users" valid:"MinSize(4);MaxSize(300)"`
	AllowReviews  bool              `form:"allow-reviews" valid:"Required"`
	AllowComments bool              `form:"allow-comments" valid:"Required"`
	InvalidFields map[string]string `form:"-"`
}
