package main

import (
	_ "embed"
	"errors"
	"fmt"
	"projectGenerator/project_generator"
	"runtime"
	"time"

	"charm.land/huh/v2"
	"charm.land/huh/v2/spinner"
	"charm.land/lipgloss/v2"
)

func main() {

	var projectName string
	var selectedProject string
	var selectedDatabase string
	var cliFramework string
	var allowProjectType bool
	var allowDatabaseFramework bool
	var allowProjectName bool
	var allowTestCases bool

	asciColor := lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	cancelledText := lipgloss.NewStyle().Foreground(lipgloss.Red)
	successText := lipgloss.NewStyle().Foreground(lipgloss.BrightGreen)
	infoText := lipgloss.NewStyle().Foreground(lipgloss.Color("85"))

	userOs := runtime.GOOS
	fmt.Println(infoText.Render("\ncurrent operating system:", userOs))

	banner := `
    ____             _           __          ______
   / __ \_________  (_)__  _____/ /_        / ____/___  ____
  / /_/ / ___/ __ \/ / _ \/ ___/ __/_______/ / __/ __ \/ __ \
 / ____/ /  / /_/ / /  __/ /__/ /_/_______/ /_/ /  __/ / / /
/_/   /_/   \____/ /\___/\___/\__/        \____/\___/_/ /_/
              /___/
`

	fmt.Println("\n", asciColor.Render(banner))
	// basic form
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter project name\n").
				Validate(func(s string) error {
					if s == "" {
						return errors.New("project name is invalid. Please enter a new name")
					}
					return nil
				}).
				Placeholder("EX: funApi").
				Value(&projectName),
			huh.NewConfirm().
				Title("Confirm project name?").
				Affirmative("Yes").
				Negative("No").
				Value(&allowProjectName).
				Validate(func(b bool) error {
					if !allowProjectName {
						return errors.New("please enter a new name. Press shift+tab")
					}
					return nil
				}),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a project\n").
				Options(
					huh.NewOption("Http-Backend", "http_backend"),
					huh.NewOption("Cli-Application", "cli_project"),
					huh.NewOption("Empty-Project", "empty_project"),
				).Value(&selectedProject),

			huh.NewConfirm().
				Title("Confirm project type?").
				Affirmative("Yes").
				Negative("No").
				Value(&allowProjectType).
				Validate(func(b bool) error {
					if !allowProjectType {
						return errors.New("please select a new project type. Please press shift+tab")
					}
					return nil
				}),
		),
		// this will be hidden if they don't choose the http-backend
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a database framework\n").
				Options(
					huh.NewOption("PostgreSQL(PGX driver)", "postgres"),
					huh.NewOption("Mysql", "mysql"),
					huh.NewOption("No database", "none"),
				).Value(&selectedDatabase),
			huh.NewConfirm().
				Title("Confirm database?").
				Affirmative("Yes").
				Negative("No").
				Value(&allowDatabaseFramework).
				Validate(func(b bool) error {
					if !allowDatabaseFramework {
						return errors.New("please select a new database framework. Please press shift+tab")
					}
					return nil
				}),
		).WithHideFunc(func() bool {
			return selectedProject != "http_backend"
		}),

		huh.NewGroup(
			huh.NewConfirm().
				Title("Include test cases?").
				Affirmative("Yes").
				Negative("No").
				Value(&allowTestCases),
		).WithHideFunc(func() bool {
			return selectedProject != "http_backend"
		}),

		// this will only show with the command line selected
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What kind of CLI application?\n").
				Options(
					huh.NewOption("Basic Input Form (Huh)", "cli-huh"),
					huh.NewOption("Interactive TUI (Bubble Tea)", "cli-bubbletea"),
					huh.NewOption("Standard Flags Only (Cobra)", "cli-cobra"),
				).
				Value(&cliFramework),
		).WithHideFunc(func() bool {
			return selectedProject != "cli_project"
		}),
	)

	if err := form.Run(); err != nil {
		// catch user cancellations and print a clean exit message
		if errors.Is(err, huh.ErrUserAborted) {
			fmt.Println(cancelledText.Render("Scaffold cancelled."))
			return
		}

		// catch terminal or rendering errors
		fmt.Println("Error:", err)
		return
	}

	var generationErr error
	var projectDir string
	spinnerErr := spinner.New().
		Type(spinner.Dots).
		Title(" Generating project...").
		Action(func() {
			time.Sleep(1 * time.Second)
			projectDir, generationErr = project_generator.GenerateProject(projectName, selectedProject, selectedDatabase, allowTestCases)
		}).
		Run()

	if spinnerErr != nil {
		fmt.Println("error create spinner.")
		return
	}

	if generationErr != nil {
		fmt.Println(cancelledText.Render("Error making project: " + generationErr.Error()))
		return
	}

	fmt.Println(successText.Render("\nProject scaffolded successfully!"))
	fmt.Println(infoText.Render("\nPlease cd into your project. Project location: " + projectDir))

}
