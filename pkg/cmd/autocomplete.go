package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ZupIT/ritchie-cli/pkg/autocomplete"
)

const (
	zsh        autocomplete.ShellName = "zsh"
	bash       autocomplete.ShellName = "bash"
	fish       autocomplete.ShellName = "fish"
	powerShell autocomplete.ShellName = "powershell"
)

var supportedShell = []string{zsh.String(), bash.String(), fish.String(), powerShell.String()}

// autocompleteCmd type for set autocomplete command
type autocompleteCmd struct {
	autocomplete.Generator
}

// NewAutocompleteCmd creates a new cmd instance
func NewAutocompleteCmd() *cobra.Command {
	shells := strings.Join(supportedShell, ", ")
	return &cobra.Command{
		Use:     "completion SUBCOMMAND",
		Short:   "Add autocomplete for terminal (" + shells + ")",
		Long:    `Add autocomplete for terminal, Available for (` + shells + `).`,
		Example: "rit completion zsh",
	}
}

// NewAutocompleteZsh creates a new cmd instance zsh
func NewAutocompleteZsh(g autocomplete.Generator) *cobra.Command {
	a := &autocompleteCmd{g}

	return &cobra.Command{
		Use:     zsh.String(),
		Short:   "Add zsh autocomplete for terminal",
		Long:    "Add zsh autocomplete for terminal",
		Example: "rit completion zsh",
		RunE:    a.runFunc(),
	}
}

// NewAutocompleteBash creates a new cmd instance zsh
func NewAutocompleteBash(g autocomplete.Generator) *cobra.Command {
	a := &autocompleteCmd{g}

	return &cobra.Command{
		Use:     bash.String(),
		Short:   "Add bash autocomplete for terminal",
		Long:    "Add bash autocomplete for terminal",
		Example: "rit completion bash",
		RunE:    a.runFunc(),
	}
}

// NewAutocompleteFish creates a new cmd instance fish
func NewAutocompleteFish(g autocomplete.Generator) *cobra.Command {
	a := &autocompleteCmd{g}

	return &cobra.Command{
		Use:   fish.String(),
		Short: "Add fish autocomplete for terminal --help to know how to use",
		Long: `
Add fish autocomplete for terminal
Only fish >= version 3.X is supported (fish 2.X is not supported)

To test run: 
 $ rit completion fish | source

To install run: 
 $ echo "rit completion fish | source" >> ~/.config/fish/config.fish

`,
		Example: "rit completion fish | source",
		RunE:    a.runFunc(),
	}
}

// NewAutocompletePowerShell creates a new cmd instance PowerShell
func NewAutocompletePowerShell(g autocomplete.Generator) *cobra.Command {
	a := &autocompleteCmd{g}

	return &cobra.Command{
		Use:   powerShell.String(),
		Short: "Add powerShell autocomplete for terminal --help to know how to use",
		Long: `
Add powerShell autocomplete for terminal
Only powerShell >= version 5.X is supported

To install run and after restart powerShell:
	rit completion powershell >> $PROFILE
`,
		Example: "rit completion powershell >> $PROFILE",
		RunE:    a.runFunc(),
	}
}

func (a autocompleteCmd) runFunc() CommandRunnerFunc {
	return func(cmd *cobra.Command, args []string) error {
		s := autocomplete.ShellName(cmd.Use)
		c, err := a.Generate(s, cmd)
		if err != nil {
			return err
		}

		fmt.Println(c)
		return nil
	}

}
