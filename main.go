package main

import (
	"fmt"
	"github.com/TwinProduction/go-color"
	"github.com/manifoldco/promptui"
	"io"
	"os"
	"path/filepath"
)

func main() {
	for {
		displayMenu()
		// displayMessage("Welcome")
	}
}

type messageItem struct {
	message string
	purpose MessagePurpose
}

type MessagePurpose int

const (
	None MessagePurpose = iota
	Info
	Warning
	Error
	Success
	ActionRequired
	TroubleshootingTip
)

func messagePrefix(purpose MessagePurpose) string {
	switch purpose {
	case 0:
		return ""
	case 1:
		return color.Ize(color.White, "Info: ")
	case 2:
		return color.Ize(color.Yellow, "Warning: ")
	case 3:
		return color.Ize(color.Red, "Error: ")
	case 4:
		return color.Ize(color.Green, "✓ ")
	case 5:
		return color.Ize(color.Yellow, "Action required: ")
	case 6:
		return color.Ize(color.Cyan, "Troubleshooting Tip: ")
	}
	panic(fmt.Sprintf("Unsupported purpose (%v)", purpose))
}

type menuItem struct {
	caption     string
	description string
}

var menuItems = []menuItem{
	menuItem{
		caption:     "Generate default prompt files",
		description: "This will generate the files that create the default prompt for OPO docker projects, in docker/",
		// action:      generate_files_for_shell_prompt,
	},
	menuItem{
		caption:     "Hello World Message",
		description: "This will display 'Hello World' in the message area (footer)",
		// action:      func() { displayMessage("Hello World") },
	},
	menuItem{
		caption:     "Quit",
		description: "Press to exit",
		// action: func() { os.Exit(0) },
	},
}

func displayMessage(message messageItem) {
	fmt.Println(messagePrefix(message.purpose) + message.message)
}

func displayMessages(messages ...messageItem) {
	for _, message := range messages {
		displayMessage(message)
	}
}

func displayMenu() {
	fmt.Println("")
	collect_captions := func(menuItems []menuItem) []string {
		var captions []string
		for _, menuItem := range menuItems {
			captions = append(captions, menuItem.caption)
		}
		return captions
	}
	prompt := promptui.Select{
		Label: "Select Action",
		Items: collect_captions(menuItems),
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// fmt.Printf("You choose %q\n", result)
	switch i {
	case 0: // Generate prompt
		generate_files_for_shell_prompt()
	case 1:
		displayMessage(messageItem{
			message: "Hello" + color.Ize(color.Green, " World"),
		})
	case 2: //quit
		os.Exit(0)
	}
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
}

func generate_files_for_shell_prompt() {
	templateDir := "templates/home"
	currentDir, err := os.Getwd()
	check(err)
	destinationDir := filepath.Join(currentDir, "docker", "home")
	displayMessage(messageItem{ message: "Generating files in '" + destinationDir + "..." })
	// mkdir, including missing dirs along path
	os.RemoveAll(destinationDir)
	os.MkdirAll(destinationDir, os.ModePerm)

	promptFiles := []string{"profile", "_shell_prompt.sh", "_shell_colors.sh", "git-prompt.sh"}
	for _, sourceFile := range promptFiles {
		copyFile(filepath.Join(templateDir, sourceFile), filepath.Join(destinationDir, sourceFile))
	}

	displayMessages(
		messageItem{purpose: Success, message: "Done."},
		messageItem{purpose: ActionRequired, message: "Add: 'source $HOME/_shell_prompt.sh' to your .bashrc file."},
		messageItem{purpose: ActionRequired, message: "Restart `bin/shell` to utilize the new prompt."},
		messageItem{purpose: TroubleshootingTip, message: "Ensure your Dockerfile copies the files: 'COPY docker/home/*.sh /root/`"},
	)
}

// Copy source file to Destination
// Derived from https://stackoverflow.com/a/35353594
func copyFile(sourceFileNameAndPath string, DestinationFileNameAndPath string) {
	srcFile, err := os.Open(sourceFileNameAndPath)
	check(err)
	defer srcFile.Close()

	destFile, err := os.Create(DestinationFileNameAndPath) // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	check(err)

	err = destFile.Sync()
	check(err)
}

// Raises error if exists
// Derived from https://stackoverflow.com/a/35353594
func check(err error) {
	if err != nil {
		fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}
