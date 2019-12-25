package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	longDescription = `
	Write shell completions for the given shell to stdout (bash/zsh).

	MacOS X:
		$ brew install bash-completion
		$ source $(brew --prefix)/etc/bash_completion
		$ kubectl view-secret completion bash > ~/.view-secret-completion  # for bash users
		$ kubectl view-secret completion zsh > ~/.view-secret-completion   # for zsh users
		$ source ~/.view-secret-completion
	Linux:
		$ apt-get install bash-completion
		$ source /etc/bash-completion
		$ source <(kubectl view-secret completion bash) # for bash users
		$ source <(kubectl view-secret completion zsh)  # for zsh users
`
)

// ErrTooManyArguments is thrown if too many arguments are specified while generating shell completions
var ErrTooManyArguments = errors.New("please provide only the desired shell, either 'bash' or 'zsh'")

func genCompletion(cmd *cobra.Command, args []string) error {
	shell := strings.ToLower(args[0])
	switch shell {
	case "bash":
		return rootCmd(cmd).GenBashCompletion(os.Stdout)
	case "zsh":
		return rootCmd(cmd).GenZshCompletion(os.Stdout)
	}

	return nil
}

// NewCmdCompletion returns the cobra command that outputs shell completion code
func NewCmdCompletion() *cobra.Command {
	return &cobra.Command{
		Use: "completion [shell]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ErrTooManyArguments
			}
			return cobra.OnlyValidArgs(cmd, args)
		},
		ValidArgs: []string{"bash", "zsh"},
		Short:     "Output shell completion for the given shell (bash or zsh)",
		Long:      longDescription,
		RunE:      genCompletion,
	}
}

func rootCmd(cmd *cobra.Command) *cobra.Command {
	root := cmd
	for root.HasParent() {
		root = root.Parent()
	}
	return root
}
