package git

import (
	"fmt"
	"os/exec"
	"strings"

	ui "github.com/abroudoux/branch/internal/ui"
	utils "github.com/abroudoux/branch/internal/utils"
)

func DoAction(branchSelected string, action string) error {
	switch action {
	case "Exit":
		fmt.Println("Exiting...")
		return nil
	case "Delete":
		return deleteBranch(branchSelected)
	case "Merge":
		return mergeBranch(branchSelected)
	case "Branch":
		return createBranch(branchSelected)
	case "Rename":
		return renameBranch(branchSelected)
	case "Checkout":
		return checkoutBranch(branchSelected)
	case "Name":
		return copyName(branchSelected)
	default:
		fmt.Println("Exiting...")
		return nil
	}
}

func renameBranch(branch string) error {
	newBranchName, err := utils.AskInput("Enter a name for the new branch: ")
	if err != nil {
		return err
	}

	if strings.Contains(newBranchName, " ") {
		return err
	}

	cmd := exec.Command("git", "branch", "-m", branch, newBranchName)
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("Branch %s renamed to %s\n", ui.RenderBranch(branch), ui.RenderBranch(newBranchName))
	return nil
}

func checkoutBranch(branch string) error {
	cmd := exec.Command("git", "checkout", branch)
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("Branch %s checked out\n", ui.RenderBranch(branch))
	return nil
}

func copyName(branch string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(branch)
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("Branch name '%s' copied to clipboard\n", ui.RenderBranch(branch))
	return nil
}

func createBranch(branchSelected string) error {
	newBranchName, err := utils.AskInput("Enter the name of the new branch: ")
	if err != nil {
		return err
	}

	defaultBranch, err := getDefaultBranch()
	if err != nil {
		return err
	}

	if branchSelected != defaultBranch {
		cmd := exec.Command("git", "checkout", branchSelected)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	if (utils.AskConfirmation("Do you want to checkout on the new branch?")) {
		cmd := exec.Command("git", "checkout", "-b", newBranchName)
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		cmd := exec.Command("git", "branch", newBranchName)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	fmt.Printf("Branch '%s' based on '%s' created\n", ui.RenderBranch(newBranchName), ui.RenderBranch(branchSelected))
	return nil
}

func mergeBranch(branch string) error {
	cmd := exec.Command("git", "merge", branch)
	err := cmd.Run()
	if err != nil {
		return err
	}

	println(fmt.Sprintf("Branch '%s' merged", branch))

	shouldDeleteBranch := utils.AskConfirmation("Do you want to delete the merged branch?")
	if shouldDeleteBranch {
		err := deleteBranch(branch)
		if err != nil {
			return err
		}
	}

	println("Branch '%s' deleted", ui.RenderBranch(branch))
	return nil
}

func deleteBranch(branch string) error {
	if !utils.AskConfirmation(fmt.Sprintf("Are you sure you want to delete '%s'?", ui.RenderBranch(branch))) {
		return fmt.Errorf("branch deletion cancelled")
	}

	cmd := exec.Command("git", "branch", "-D", branch)
	err := cmd.Run()
	if err != nil {
		return err
	}

	println(fmt.Sprintf("Branch '%s' deleted", ui.RenderBranch(branch)))

	if utils.AskConfirmation(fmt.Sprintf("Do you want to delete '%s' remotly?", ui.RenderBranch(branch))) {
		err := deleteRemoteBranch(branch)
		if err != nil {
			return err
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
		return err
	}

	return nil
}

func getBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--format=%(refname:short)")
	branches, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}

	return strings.Fields(string(branches)), nil
}

func getDefaultBranch() (string, error) {
	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	defaultBranch, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(defaultBranch)), nil
}

func GetBranchesWithDefaultIndication() ([]string, error) {
	branches, err := getBranches()
	if err != nil {
		return []string{}, err
	}

	defaultBranch, err := getDefaultBranch()
	if err != nil {
		return []string{}, nil
	}

	branchesWithDefaultIndication := []string{}

	for _, branch := range branches {
		if branch == defaultBranch {
			branchesWithDefaultIndication = append(branchesWithDefaultIndication, "* "+branch)
		} else {
			branchesWithDefaultIndication = append(branchesWithDefaultIndication, "  "+branch)
		}
	}

	return branchesWithDefaultIndication, nil
}