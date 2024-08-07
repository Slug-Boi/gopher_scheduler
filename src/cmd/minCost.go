package cmd

import (
	"fmt"
	"sort"

	"github.com/Slug-Boi/aion-cli/src/config"
	"github.com/Slug-Boi/aion-cli/src/forms"
	"github.com/Slug-Boi/aion-cli/src/solvers/graph"
	"github.com/facette/natsort"
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
		config.CheckConfig()

		if id, _ := cmd.Flags().GetBool("saveID"); id {
			config.CheckConfig()
			fmt.Println("\nSaving form ID to config file...")
			EditFormID(args[0])
			fmt.Println()
		}

		sink, users, cost, paths, nodeToTimeslot, groupTimeslotCost := graph.SolveMin_Cost(args)

		printSolutionMinCost(sink, users, cost, paths, nodeToTimeslot, groupTimeslotCost)
	},
}

func init() {
	solveCmd.AddCommand(minCostCmd)
	minCostCmd.Flags().Bool("saveID", false, "Save the formID to the config file")
}

func printSolutionMinCost(sink int, users map[int]forms.Form, cost float64, paths [][]int, nodeToTimeslot map[int]string, groupTimeslotCost map[string]float64) {

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
	// finalPaths keys sort on group number
	// https://www.geeksforgeeks.org/how-to-sort-golang-map-by-keys-or-values/
	keys := make([]int, 0, len(finalPaths))
	for k := range finalPaths {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return natsort.Compare(users[keys[i]].GroupNumber, users[keys[j]].GroupNumber)
	})

	//TODO: Would be nice if this was sorted on the group number so it always comes in the same order
	// Could be done using a byte array then join printing but check if its easier to sort on the HTML side

	fmt.Println("Min Cost Flow")
	fmt.Println("Min Cost: ", cost)
	fmt.Println("Sink:", sink)
	fmt.Println("Paths used:")
	for _, user := range keys {
		fmt.Println("Path:\n", sink, finalPaths[user], user, 0)
		fmt.Println("User:", users[user].GroupNumber, "->", nodeToTimeslot[finalPaths[user]], "Wish Level:", users[user].Votes[nodeToTimeslot[finalPaths[user]]], "|", "Cost:", groupTimeslotCost[users[user].GroupNumber+nodeToTimeslot[finalPaths[user]]])
	}
}
