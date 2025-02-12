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
			name := "* " + branchNameTrimed
			branchesWithSymbols = append(branchesWithSymbols, BranchWithSymbol{
				Name:   name,
				IsHead: true,
				Branch: branch,
			})
		} else {
			name := "  " + branchNameTrimed
			branchesWithSymbols = append(branchesWithSymbols, BranchWithSymbol{
				Name:   name,
				IsHead: false,
				Branch: branch,
			})
		}
	}

	return branchesWithSymbols
}
