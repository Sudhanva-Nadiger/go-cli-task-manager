package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/sudhanva-nadiger/task/db"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
				continue
			}
			ids = append(ids, id)
		}
		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}

			t := tasks[id-1]
			err := db.DeleteTask(t.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed, Error: %s\n", id, err)
			} else {
				fmt.Printf("Successfully marked \"%d\" as completed\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
