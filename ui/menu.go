package ui

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

type MenuItem struct {
	Caption     string
	Description string
	Action func()
}


func DisplayMenu(menuItems []MenuItem) {
	fmt.Println("")
	collect_captions := func(menuItems []MenuItem) []string {
		var captions []string
		for _, MenuItem := range menuItems {
			captions = append(captions, MenuItem.Caption)
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
	
	// run the appropriate action
	menuItems[i].Action()
	// switch i {
	// case 0: // Generate prompt
	// 	generate_files_for_docker()
	// case 1:
	// 	generate_files_for_shell_prompt()
	// case 2:
	// 	displayMessage(messageItem{
	// 		message: "Hello" + color.Ize(color.Green, " World"),
	// 	})
	// case 3:
	// 	os.Exit(0)
	// }
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
}
