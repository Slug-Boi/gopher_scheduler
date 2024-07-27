/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Slug-Boi/aion-cli/forms"
	"github.com/Slug-Boi/aion-cli/graph"
	"github.com/spf13/cobra"
)

// minCostCmd represents the minCost command
var minCostCmd = &cobra.Command{
	Use:   "minCost [formID]",
	Short: "This command will run the minCost graph solver and print the solution to the terminal",
	Long: `The min_cost command will run the min cost flow graph solver to solve for minimum cost scheduling.
	The solver is based on SPFA (Shortest Path Faster Algorithm) uses negative cycles to reduce cost.
	It will print the solution to the terminal (this is mostly for debugging 
	and testing purposes use the generate or root command to actually run the program).
	
	`,
	Run: func(cmd *cobra.Command, args []string) {

	sink, users, cost, paths, nodeToTimeslot := SolveMin_Cost(args)

	printSolutionMinCost(sink, users, cost, paths, nodeToTimeslot)
	},
}

func init() {
	solveCmd.AddCommand(minCostCmd)

}

func SolveMin_Cost(args []string) (int, map[int]forms.Form, float64, [][]int, map[int]string) {

	// Get the config file
	conf := SetupConfig(args)

	//TODO: Make this a hidden form id see if there is a way to make it display when clicked
	fmt.Println("Form is being processed with the following Form ID:", conf.FormID)

	form := forms.GetForm(conf)

	// Create a graph
	g, sink, users, nodeToTimeslot := graph.Translate(form)

	groups := len(form)

	cost, paths := graph.MinCostPath(len(g), groups, 0, sink, g)


	return sink, users, cost, paths, nodeToTimeslot
}

func printSolutionMinCost(sink int, users map[int]forms.Form, cost float64, paths [][]int, nodeToTimeslot map[int]string) {

		fmt.Println("Sink:", sink)
		fmt.Println("Paths used:")

		finalPaths := map[int]int{}

		for _, path := range paths {
			i := sink
			timeslotNode := -1
			for i != 0 {
				if _, ok := users[i]; !ok {
					timeslotNode = i
				} else {
					if timeslotNode != -1 {
					finalPaths[i] = timeslotNode
					}
					timeslotNode = -1
				}
				i = path[i]
			}
		}
		//TODO: Would be nice if this was sorted on the group number so it always comes in the same order
		// Could be done using a byte array then join printing but check if its easier to sort on the HTML side
		for user, timeslot := range finalPaths {
			fmt.Println("Path:\n",sink, timeslot, user, 0)
			fmt.Println("User:", users[user].GroupNumber, "Timeslot:", nodeToTimeslot[timeslot])
		}

		println("Min cost:", int(cost), "≈", cost)
}