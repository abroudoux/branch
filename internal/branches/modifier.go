package branches

import "fmt"

func AddSymbolsToBranches(branches []Branch, head Branch) []BranchWithSymbol {
	branchesWithSymbols := []BranchWithSymbol{}
	for _, branch := range branches {
		if branch.Hash() == head.Hash() {
			branchesWithSymbols = append(branchesWithSymbols, BranchWithSymbol{
				Name:   fmt.Sprint("* ", branch.Name()),
				IsHead: true,
				Branch: branch,
			})
		} else {
			branchesWithSymbols = append(branchesWithSymbols, BranchWithSymbol{
				Name:   fmt.Sprint("  ", branch.Name()),
				IsHead: false,
				Branch: branch,
			})
		}
	}

	return branchesWithSymbols
}
