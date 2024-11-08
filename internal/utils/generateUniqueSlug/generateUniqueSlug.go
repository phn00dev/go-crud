package generateuniqueslug

import (
	"fmt"

	"github.com/gosimple/slug"

	"github.com/phn00dev/go-crud/internal/domain/post/repository"

)

func GenerateUniqueSlug(title string, postRepo repository.PostRepository) (string, error) {
	slug := slug.Make(title)
	count := 1

	// Unikal slug döretmek üçin barlamak
	for postRepo.SlugExists(slug) {
		slug = fmt.Sprintf("%s-%d", slug, count)
		count++
	}
	return slug, nil
}
