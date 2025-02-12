package branches

import (
	"strings"
)

func AddSymbolsToBranches(branches []Branch, head Branch) []BranchWithSymbol {
	branchesWithSymbols := []BranchWithSymbol{}
	for _, branch := range branches {
		branchName := string(branch.Name())
		branchNameTrimed := strings.TrimPrefix(string(branchName), "refs/heads/")

		if branch.Name() == head.Name() {
			branchesWithSymbols = append(branchesWithSymbols, BranchWithSymbol{
				Name:   branchNameTrimed,
				IsHead: true,
				Branch: branch,
			})
		} else {
			branchesWithSymbols = append(branchesWithSymbols, BranchWithSymbol{
				Name:   branchNameTrimed,
				IsHead: false,
				Branch: branch,
			})
		}
	}

	return branchesWithSymbols
}
