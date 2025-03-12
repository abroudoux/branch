package branches

func GetBranches(repository Repository) ([]Branch, error) {
	branchIter, err := repository.Branches()
	if err != nil {
		return nil, err
	}

	var branches []Branch
	err = branchIter.ForEach(func(ref Branch) error {
		branches = append(branches, ref)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return branches, nil
}

func GetHead(repo Repository) (Branch, error) {
	head, err := repo.Head()
	if err != nil {
		return nil, err
	}

	return head, nil
}
