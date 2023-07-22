package domain_test

import (
	"cosmic/domain"
	"testing"
)

func TestNewPost(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		images      []string
		expectedErr error
	}{
		{
			name:        "Empty content and empty images",
			content:     "",
			images:      []string{},
			expectedErr: domain.ErrImageAndContentMustHaveOneThatIsNotNull,
		},
		{
			name:        "Null content and non-empty images",
			content:     "",
			images:      []string{"image1", "image2"},
			expectedErr: nil,
		},
		{
			name:        "Non-empty content and null images",
			content:     "Hello",
			images:      []string{},
			expectedErr: nil,
		},
		{
			name:        "Non-empty content and non-empty images",
			content:     "Hello",
			images:      []string{"image1"},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := domain.NewPost(tt.content, tt.images)

			if err != tt.expectedErr {
				t.Errorf("Got error: %v, Expected: %v", err, tt.expectedErr)
			}
		})
	}
}
