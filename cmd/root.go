package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"git.cryptic.systems/volker.raschek/dcmerge/pkg/domain/dockerCompose"
	"git.cryptic.systems/volker.raschek/dcmerge/pkg/fetcher"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func Execute(version string) error {
	completionCmd := &cobra.Command{
		Use:                   "completion [bash|zsh|fish|powershell]",
		Short:                 "Generate completion script",
		Long:                  "To load completions",
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			}
		},
	}

	rootCmd := &cobra.Command{
		Use:   "dcmerge",
		Args:  cobra.MinimumNArgs(1),
		Short: "Merge docker-compose files from multiple resources",
		Example: `dcmerge docker-compose.yml ./integration-test/docker-compose.yml
dcmerge docker-compose.yml https://git.example.local/user/repo/docker-compose.yml`,
		RunE:    run,
		Version: version,
	}
	rootCmd.Flags().BoolP("existing-win", "f", false, "Protect existing attributes")
	rootCmd.Flags().BoolP("last-win", "l", false, "Overwrite existing attributes")
	rootCmd.Flags().StringP("output-file", "o", "", "Write instead on stdout into a file")
	rootCmd.AddCommand(completionCmd)

	return rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) error {
	mergeExisting, err := cmd.Flags().GetBool("existing-win")
	if err != nil {
		return fmt.Errorf("Failed to parse flag last-win: %s", err)
	}

	mergeLastWin, err := cmd.Flags().GetBool("last-win")
	if err != nil {
		return fmt.Errorf("Failed to parse flag last-win: %s", err)
	}

	outputFile, err := cmd.Flags().GetString("output-file")
	if err != nil {
		return fmt.Errorf("Failed to parse flag output-file: %s", err)
	}

	dockerComposeConfig := dockerCompose.NewConfig()

	dockerComposeConfigs, err := fetcher.Fetch(args...)
	if err != nil {
		return err
	}

	for _, config := range dockerComposeConfigs {
		switch {
		case mergeExisting && mergeLastWin:
			return fmt.Errorf("Neither --first-win or --last-win can be specified - not booth.")
		case mergeExisting && !mergeLastWin:
			dockerComposeConfig.MergeExistingWin(config)
		case !mergeExisting && mergeLastWin:
			dockerComposeConfig.MergeLastWin(config)
		default:
			dockerComposeConfig.Merge(config)
		}
	}

	switch {
	case len(outputFile) > 0:
		err = os.MkdirAll(filepath.Dir(outputFile), 0755)
		if err != nil {
			return err
		}

		f, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer f.Close()

		yamlEncoder := yaml.NewEncoder(f)
		yamlEncoder.SetIndent(2)
		return yamlEncoder.Encode(dockerComposeConfig)

	default:
		yamlEncoder := yaml.NewEncoder(os.Stdout)
		yamlEncoder.SetIndent(2)
		return yamlEncoder.Encode(dockerComposeConfig)
	}

}
