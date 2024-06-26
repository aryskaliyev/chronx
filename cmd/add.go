package cmd

import (
	"chronx/pkg"
	"github.com/spf13/cobra"
)

var (
	description string
	colorId     string
	start       string
	end         string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new event to user's calendar",
	Long: `Chronx add command is used to create new events in google calendar
	by passing arguments like Title, description, start time, end time, event links etc`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		pkg.AddEvent(title, description, colorId, start, end)
	},
	Example: `chronx add "Visit Park" -s "2024-05-16 18:00:00"\
	-e "2024-05-16 19:00:00" -i "2" -d "Visit the park by 6pm"`,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringVarP(&description, "desc", "d", "", "add description to event")
	addCmd.PersistentFlags().StringVarP(&colorId, "id", "i", "", "add color id to event")
	addCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "add start time to event")
	addCmd.PersistentFlags().StringVarP(&end, "end", "e", "", "add end time to event")
}
