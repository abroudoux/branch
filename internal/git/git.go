package git

import (
	"fmt"
	"os/exec"
	"strings"

	ui "github.com/abroudoux/branch/internal/ui"
	utils "github.com/abroudoux/branch/internal/utils"
)

func RenameBranch(branch string) error {
	newBranchName, err := utils.AskInput("Enter a name for the new branch: ")
	if err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	if strings.Contains(newBranchName, " ") {
		return fmt.Errorf("error: the branch name must not contain spaces")
	}

	cmd := exec.Command("git", "branch", "-m", branch, newBranchName)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error renaming branch: %v", err)
	}

	fmt.Printf("Branch %s renamed to %s\n", ui.RenderBranch(branch), ui.RenderBranch(newBranchName))
	return nil
}

func CheckoutBranch(branch string) error {
	cmd := exec.Command("git", "checkout", branch)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error checking out branch: %v", err)
	}

	fmt.Printf("Branch %s checked out\n", ui.RenderBranch(branch))
	return nil
}

func CopyName(branch string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(branch)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error copying branch name: %v", err)
	}

	fmt.Printf("Branch name '%s' copied to clipboard\n", ui.RenderBranch(branch))
	return nil
}

func CreateBranch(branch string) error {
	newBranchName, err := utils.AskInput("Enter the name of the new branch: ")
	if err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	// branches := getBranches()
	// for _, branch := range branches {
	// 	if branch == newBranchName {
	// 		return fmt.Errorf("branch '%s' already exists", ui.RenderBranch(newBranchName))
	// 	}
	// }

	// defaultBranch := getDefaultBranch()
	// if branchSelected != defaultBranch {
	// 	cmd := exec.Command("git", "checkout", branchSelected)
	// 	err := cmd.Run()
	// 	if err != nil {
	// 		return fmt.Errorf("error checking out default branch: %v", err)
	// 	}
	// }

	// if (utils.AskConfirmation("Do you want to checkout on the new branch?")) {
	// 	cmd := exec.Command("git", "checkout", "-b", newBranchName)
	// 	err := cmd.Run()
	// 	if err != nil {
	// 		return fmt.Errorf("error creating branch: %v", err)
	// 	}
	// } else {
	// 	cmd := exec.Command("git", "branch", newBranchName)
	// 	err := cmd.Run()
	// 	if err != nil {
	// 		return fmt.Errorf("error creating branch: %v", err)
	// 	}
	// }

	fmt.Printf("Branch '%s' based on '%s' created\n", ui.RenderBranch(newBranchName), ui.RenderBranch(branch))
	return nil
}

func MergeBranch(branch string) error {
	cmd := exec.Command("git", "merge", branch)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error merging branch", err)
		return fmt.Errorf("error merging branch: %v", err)
	}

	println("Branch '%s' merged", branch)

	shouldDeleteBranch := utils.AskConfirmation("Do you want to delete the merged branch?")
	if shouldDeleteBranch {
		err := DeleteBranch(branch)
		if err != nil {
			return fmt.Errorf("error deleting branch: %v", err)
		}
	}

	println("Branch '%s' deleted", ui.RenderBranch(branch))

	return nil
}

func DeleteBranch(branch string) error {
	if !utils.AskConfirmation(fmt.Sprintf("Are you sure you want to delete '%s'?", ui.RenderBranch(branch))) {
		return fmt.Errorf("branch deletion cancelled")
	}

	cmd := exec.Command("git", "branch", "-D", branch)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error deleting branch: %v", err)
	}

	println("Branch '%s' deleted", ui.RenderBranch(branch))

	if hasRemoteBranch(branch) && utils.AskConfirmation(fmt.Sprintf("Do you want to delete '%s' remotly?", ui.RenderBranch(branch))) {
		err := deleteRemoteBranch(branch)
		if err != nil {
			return fmt.Errorf("error deleting remote branch: %v", err)
		}
	}
	return nil
}

func deleteRemoteBranch(branch string) error {
	if !utils.AskConfirmation(fmt.Sprintf("Are you sure you want to delete '%s' remotly?", ui.RenderBranch(branch))) {
		return fmt.Errorf("branch deletion cancelled")
	}

	cmd := exec.Command("git", "push", "origin", "--delete", branch)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error deleting remote branch: %v", err)
	}

	return nil
}

func hasRemoteBranch(branchName string) bool {
	cmd := exec.Command("git", "ls-remote", "--heads", "origin", branchName)
	err := cmd.Run()
	return err == nil
}