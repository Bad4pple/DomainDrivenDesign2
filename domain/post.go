package domain

import "errors"

var (
	ErrImageAndContentMustHaveOneThatIsNotNull = errors.New("creating post the image and content must have one that is not null")
)

type Post struct {
	Content string
	Images  []string
}

func NewPost(content string, images []string) (*Post, error) {
	if len(content) == 0 && len(images) == 0 {
		return nil, ErrImageAndContentMustHaveOneThatIsNotNull
	}
	return &Post{
		Content: content,
		Images:  images,
	}, nil
}
