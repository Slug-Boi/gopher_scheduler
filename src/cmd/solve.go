package cmd

import (
	"strconv"

	"github.com/Slug-Boi/aion-cli/graph"
	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "This command will run the selected solver and print the solution to the terminal",
	Long: `This is mostly a debugging tool to see the output of the solver.
	The two available solvers are the min_cost graph (min) and gurobi (gur) solver.
	....
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Create a graph
		g := debugGraphBuilder()

		groups, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		cost, paths := graph.MinCostPath(9, 3, 0, 8, g)

		println("Paths used")

		for j, path := range paths {
			println("Path:", j)
			i := 8
			println(i)
			for i != 0 {
				println(path[i])
				i = path[i]
			}
		}

		cost = ((cost - groups) - len(paths))

		println("Min cost: ", cost)

	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

}

// This is a debugging graph
func debugGraphBuilder() []graph.Edge {
	// Edge values
	// From, To, Capacity, Cost
	g := []graph.Edge{}

	// 0 source
	// 8 sink
	// Groups -> 1, 2, 3
	// Timeslots -> 4, 5, 6, 7
	// Add edges to the graph
	g = append(g, graph.Edge{From: 0, To: 1, Capacity: 100, Cost: 1})
	g = append(g, graph.Edge{From: 0, To: 2, Capacity: 100, Cost: 1})
	g = append(g, graph.Edge{From: 0, To: 3, Capacity: 100, Cost: 1})

	g = append(g, graph.Edge{From: 1, To: 4, Capacity: 1, Cost: 1})
	g = append(g, graph.Edge{From: 2, To: 4, Capacity: 1, Cost: 1})
	g = append(g, graph.Edge{From: 2, To: 5, Capacity: 2, Cost: 2})
	g = append(g, graph.Edge{From: 3, To: 6, Capacity: 3, Cost: 3})
	g = append(g, graph.Edge{From: 3, To: 7, Capacity: 1, Cost: 1})

	g = append(g, graph.Edge{From: 4, To: 8, Capacity: 1, Cost: 1})
	g = append(g, graph.Edge{From: 5, To: 8, Capacity: 1, Cost: 1})
	g = append(g, graph.Edge{From: 6, To: 8, Capacity: 1, Cost: 1})
	g = append(g, graph.Edge{From: 7, To: 8, Capacity: 1, Cost: 1})

	return g
}