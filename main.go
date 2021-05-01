package main

import (
	"dkman/ui"
	"embed"
	"fmt"
	"github.com/TwinProduction/go-color"
	"os"
	"path/filepath"
)

// Embed template files in executable
// WORAROUND: embed "templates/*" ignores _* files
//go:embed templates/* templates/home/_shell_colors.sh templates/home/_shell_prompt.sh
var templateFiles embed.FS

func main() {
	// Loop. After the user selects an action, re-display the menu
	for {
		ui.DisplayMenu(menuItems)
	}
}

var menuItems = []ui.MenuItem{
	ui.MenuItem{
		Caption:     "Generate all docker files",
		Description: "This will generate all the files that are copied into OPO docker projects, in docker/",
		Action:      generate_files_for_docker,
	},
	ui.MenuItem{
		Caption:     "Generate default prompt files",
		Description: "This will generate the files that create the default prompt for OPO docker projects, in docker/",
		Action:      generate_files_for_shell_prompt,
	},
	ui.MenuItem{
		Caption:     "Hello World Message",
		Description: "This will display 'Hello World' in the message area (footer)",
		Action:      func() { ui.DisplayMessage(ui.MessageItem{Message: "Hello" + color.Ize(color.Green, " World")}) },
	},
	ui.MenuItem{
		Caption:     "Quit",
		Description: "Press to exit",
		Action:      func() { os.Exit(0) },
	},
}

func generate_files_for_docker() {
	templateDir := "templates/"
	currentDir, err := os.Getwd()
	check(err)
	destinationDir := filepath.Join(currentDir, "docker")
	ui.DisplayMessage(ui.MessageItem{Message: "Generating files in '" + destinationDir + "..."})
	// mkdir, including missing dirs along path
	os.RemoveAll(destinationDir)
	os.MkdirAll(destinationDir, os.ModePerm)

	dockerFiles := []string{"readme.md", "/home/profile", "/home/_shell_colors.sh"}
	for _, sourceFile := range dockerFiles {
		copyEmbeddedFile(filepath.Join(templateDir, sourceFile), filepath.Join(destinationDir, sourceFile))
	}

	generate_files_for_shell_prompt()
}

func generate_files_for_shell_prompt() {
	templateDir := "templates/home"
	currentDir, err := os.Getwd()
	check(err)
	destinationDir := filepath.Join(currentDir, "docker", "home")
	ui.DisplayMessage(ui.MessageItem{Message: "Generating files in '" + destinationDir + "..."})
	// mkdir, including missing dirs along path
	// os.RemoveAll(destinationDir)
	os.MkdirAll(destinationDir, os.ModePerm)

	promptFiles := []string{"_shell_prompt.sh", "_shell_colors.sh", "git-prompt.sh"}
	for _, sourceFile := range promptFiles {
		copyEmbeddedFile(filepath.Join(templateDir, sourceFile), filepath.Join(destinationDir, sourceFile))
	}

	ui.DisplayMessages(
		ui.MessageItem{Purpose: ui.Success, Message: "Done."},
		ui.MessageItem{Purpose: ui.ActionRequired, Message: "Add: 'source $HOME/_shell_prompt.sh' to your .bashrc file."},
		ui.MessageItem{Purpose: ui.ActionRequired, Message: "Restart `bin/shell` to utilize the new prompt."},
		ui.MessageItem{Purpose: ui.TroubleshootingTip, Message: "Ensure your Dockerfile copies the files: 'COPY docker/home/*.sh /root/`"},
	)
}

// Copies contents of "embedded file" to OS file
// derived from: https://golang.org/pkg/embed/
func copyEmbeddedFile(embeddedFileName string, destinationFileNameAndPath string) {
	ui.DisplayMessage(ui.MessageItem{Message: "Copying '" + embeddedFileName + "..."})
	embeddedFileContentsAsBase64, err := templateFiles.ReadFile(embeddedFileName)
	check(err)
	embeddedFileContents := string(embeddedFileContentsAsBase64)
	_, err = createFile(embeddedFileContents, destinationFileNameAndPath)
	check(err)
}

// Creates destination file from contents, overwritting existing files
// Returns count_of_bytes_written, err
func createFile(contents string, destinationFileNameAndPath string) (int, error) {
	destinationDir := filepath.Dir(destinationFileNameAndPath)
	os.MkdirAll(destinationDir, os.ModePerm)
	ui.DisplayMessage(ui.MessageItem{Message: "Creating '" + destinationFileNameAndPath + "..."})
	destFile, err := os.Create(destinationFileNameAndPath) // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	return destFile.WriteString(contents)
}

// // Copy source file to Destination
// // Derived from https://stackoverflow.com/a/35353594
// func copyFile(sourceFileNameAndPath string, DestinationFileNameAndPath string) {
// 	srcFile, err := os.Open(sourceFileNameAndPath)
// 	check(err)
// 	defer srcFile.Close()

// 	destFile, err := os.Create(DestinationFileNameAndPath) // creates if file doesn't exist
// 	check(err)
// 	defer destFile.Close()

// 	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
// 	check(err)

// 	err = destFile.Sync()
// 	check(err)
// }

// Raises error if exists
// Derived from https://stackoverflow.com/a/35353594
func check(err error) {
	if err != nil {
		fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}
