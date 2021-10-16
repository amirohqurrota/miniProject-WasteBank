package newsApi

type Domain struct {
	Source      string
	Author      string
	Title       string
	Description string
	Content     string
	Url         string
}

type Repository interface {
	GetNews() (Domain, error)
}

// type Service interface {
// 	GetNews(ctx context.Context) (Domain, error)
// }
