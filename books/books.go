package books

type CategoryType string

const (
	CategoryNovel      CategoryType = "Novel"
	CategoryShortStory CategoryType = "ShortStory"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) Category() CategoryType {
	if b.Pages > 300 {
		return CategoryNovel
	}
	return CategoryShortStory
}
