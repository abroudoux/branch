package branches

import "fmt"

func AddSymbolsToBranches(branches []Branch, head Branch) []BranchWithSymbol {
	branchesWithSymbols := []BranchWithSymbol{}
	for _, branch := range branches {
		if branch.Hash() == head.Hash() {
			branchesWithSymbols = append(branchesWithSymbols, BranchWithSymbol{
				Name:   fmt.Sprint("* ", branch.Name()),
				Branch: branch,
			})
		} else {
			branchesWithSymbols = append(branchesWithSymbols, BranchWithSymbol{
				Name:   fmt.Sprint("  ", branch.Name()),
				Branch: branch,
			})
		}
	}

	return branchesWithSymbols
}
