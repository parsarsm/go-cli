package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "get sum of numbers",
	Long: `calculating sum of numbers
		use -f(-float) for float numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		floatFlag, _ := cmd.Flags().GetBool("float")
		if floatFlag {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

func addInt(args []string) {
	var sum int

	for _, val := range args {
		temp, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Addition of numbers %s is %d", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, val := range args {
		temp, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Sum of floating numbers %s is %f", args, sum)
}
