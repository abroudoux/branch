package git

import (
	"github.com/go-git/go-git/v5"
)

func GetRepositoryCurrentDir() (*git.Repository, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}

	return repo, nil
}
