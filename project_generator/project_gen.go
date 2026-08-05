package project_generator

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

//go:embed templates
var templates embed.FS

// templateFile maps a template's path inside the embedded FS to where it
// should land in the generated project, relative to parentDir.
type templateFile struct {
	src  string
	dest string
}

type TemplateData struct {
	ProjectName string
	GoVersion   string
}

// GenerateProject takes in fields from form, and inserts directories and files based of these.
func GenerateProject(projectName, projectType, selectedDatabase string, allowTestCases bool) (string, error) {
	userHomeDir, _ := os.UserHomeDir()
	parentDir := filepath.Join(userHomeDir, "generated_go_projects", projectName)

	// check if they have go installed first
	_, err := exec.LookPath("go")
	if err != nil {
		return "", fmt.Errorf("golang not installed. Please install golang first")
	}

	// get their go version and strip the word "go" from the front
	userGoVer := strings.ReplaceAll(runtime.Version(), "go", "")

	if err := os.MkdirAll(parentDir, 0750); err != nil {
		return "", fmt.Errorf("error making directory. error: %s", err)
	}

	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = parentDir

	// bubble error up
	// for now I do not think I need the output of this....
	_, err = cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error during go mod init, please check if project exists already. error: %s", err)
	}

	dirPaths := []string{
		filepath.Join(parentDir, ".github", "workflows"),
		filepath.Join(parentDir, "cmd"),
	}

	// initial files,
	// I might change this to include more base ones
	templateFiles := []templateFile{
		{
			"templates/template.gitignore.txt",
			".gitignore",
		},
		{
			"templates/template.github.workflow.ci.txt",
			filepath.Join(".github", "workflows", "ci.yml"),
		},
	}

	httpBackStdLibTemplates := "templates/http_backend_templates/standard_library/"

	switch projectType {
	case "http_backend":
		dirPaths = append(dirPaths,
			filepath.Join(parentDir, "cmd", "api"),
			filepath.Join(parentDir, "internal"),
			filepath.Join(parentDir, "internal", "domain"),
			filepath.Join(parentDir, "internal", "service"),
			filepath.Join(parentDir, "internal", "handler"))

		templateFiles = append(templateFiles,
			templateFile{
				"templates/http_backend_templates/template.http_backend.makefile.txt",
				"makefile",
			},
			templateFile{
				"templates/http_backend_templates/template.http_backend.env.txt",
				".env",
			},
		)

		switch selectedDatabase {

		case "postgres":
			postGressTemplate := httpBackStdLibTemplates + "with_database/postgressql/"

			templateFiles = append(templateFiles,
				templateFile{
					postGressTemplate + "template.http_backend.main.go.txt",
					filepath.Join("cmd", "api", "main.go")},

				templateFile{
					postGressTemplate + "template.http_backend.run.go.txt",
					filepath.Join("cmd", "api", "run.go"),
				},

				templateFile{
					postGressTemplate + "template.http_backend.setup_routes.go.txt",
					filepath.Join("cmd", "api", "setupRoutes.go"),
				},

				templateFile{
					postGressTemplate + "template.http_backend.handler.user.go.txt",
					filepath.Join("internal", "handler", "user.go"),
				},

				templateFile{
					postGressTemplate + "template.http_backend.service.user.go.txt",
					filepath.Join("internal", "service", "user.go"),
				},

				templateFile{
					postGressTemplate + "template.http_backend.domain.user.go.txt",
					filepath.Join("internal", "domain", "user.go"),
				},
			)

		// either none or DB that does not exist form template.
		default:

		}

	}

	for i := range dirPaths {
		// bugfix: for making .GitHub files,
		// I had a bug where they could not be made, so changing os.Mkdir to: os.MkdirAll fixed it.
		err := os.MkdirAll(dirPaths[i], 0750)
		if err != nil {
			return "", fmt.Errorf("error during internal file path creation. error: %s", err)
		}
	}

	// this will substitute the "{{ .ProjectName }}" in the text files.
	// it solves the internal scaffolding import problem I faced.
	templateData := TemplateData{ProjectName: projectName, GoVersion: userGoVer}

	for _, t := range templateFiles {
		rawData, err := templates.ReadFile(t.src)
		if err != nil {
			return "", fmt.Errorf("error during file template reading. error: %s", err)
		}

		tmpl, err := template.New(t.src).Parse(string(rawData))
		if err != nil {
			return "", fmt.Errorf("error parsing template %s. error: %s", t.src, err)
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, templateData); err != nil {
			return "", fmt.Errorf("error executing template %s. error: %s", t.src, err)
		}

		if err := os.WriteFile(filepath.Join(parentDir, t.dest), buf.Bytes(), 0660); err != nil {
			return "", fmt.Errorf("error during template file writing. error: %s", err)
		}
	}

	return parentDir, nil
}
