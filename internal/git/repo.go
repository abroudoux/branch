package git

import (
	"github.com/go-git/go-git/v5"
)

func GetRepository(dir string) (Repository, error) {
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
