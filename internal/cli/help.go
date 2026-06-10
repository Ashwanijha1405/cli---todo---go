package cli

import (
	"cli-todo/internal/logger"
	"fmt"
)

func PrintHelp() {

	logger.Log.Info("displaying help menu")


	fmt.Println("Available Commands:")
	fmt.Println("add <title>     Add a new todo")
	fmt.Println("list            List all todos")
	fmt.Println("done <id>       Mark todo as completed")
	fmt.Println("delete <id>     Delete todo")
	fmt.Println("help            Show help menu")
}
