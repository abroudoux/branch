package actions

import (
	"fmt"

	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/logs"
	"github.com/abroudoux/branch/internal/ui"
	"github.com/atotto/clipboard"
)

func copyBranchName(branch branches.BranchDetails) error {
	if clipboard.Unsupported {
		return fmt.Errorf("Clipboard not supported on this plateform.")
	}

	err := clipboard.WriteAll(branch.Name)
	if err != nil {
		return err
	}

	logs.Info(fmt.Sprintf("Branch name '%s' copy to the clipboard.", ui.RenderElementSelected(branch.Name)))
	return nil
}
