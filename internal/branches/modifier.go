package branches

import (
	"strings"
)

func CreateBranchesWithDetails(branches []Branch, head Branch) []BranchDetails {
	var branchesDetails []BranchDetails
	var headBranch BranchDetails

	for _, branch := range branches {
		branchName := string(branch.Name())
		branchNameTrimmed := strings.TrimPrefix(branchName, "refs/heads/")

		branchDetail := BranchDetails{
			Name:   branchNameTrimmed,
			Head:   branch.Name() == head.Name(),
			Branch: branch,
		}

		if branchDetail.Head {
			headBranch = branchDetail
		} else {
			branchesDetails = append(branchesDetails, branchDetail)
		}
	}

	if (headBranch != BranchDetails{}) {
		branchesDetails = append(branchesDetails, headBranch)
	}

	return branchesDetails
}
