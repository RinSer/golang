/**
 * The implementation of Kosaraju's algorithm 
 * that finds all the strongly connected components 
 * in a given graph.
 *
 * created by RinSer
 */

package main


import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)


func Visit(vertex int, graph_edges [][2]int, explored map[int]bool, nodes_order []int) []int {
	/*
	 * Helper subroutine for the main algorithm.
     * Visits recursively all the graph nodes and adds them 
     * to a list in order reversed to finding time.
     * Returns the current slice of nodes' ordering.
     */
	if !explored[vertex] {
		explored[vertex] = true
		for edge := range(graph_edges) {
			if graph_edges[edge][0] == vertex {
				nodes_order = Visit(graph_edges[edge][1], graph_edges, explored, nodes_order)
			}
		}
		nodes_order = append([]int{vertex}, nodes_order...)
	}
    
    return nodes_order
}


func Assign(node, root int, graph_edges [][2]int, roots map[int]int) {
	/*
	 * Helper subroutine for the main algorithm.
     * Traverses recursively the transposed graph in arbitrary order 
     * finding the roots of each node.
     * Returns nothing.
	 */
	if roots[node] == 0 {
		roots[node] = root
		for edge := range(graph_edges) {
			if graph_edges[edge][1] == node {
				Assign(graph_edges[edge][0], root, graph_edges, roots)
			}
		}
	}
}


func Kosaraju(graph_edges [][2]int, number_of_vertices int) map[int]int {
	/*
	 * Kosaraju algorithm that finds all the strongly connected components 
     * in a given graph.
     * Returns a map of nodes onto their roots 
     * (each root with its nodes corresponds to a strongly connected component).
	 */
	explored := make(map[int]bool)
    roots := make(map[int]int)
	// Mark all the nodes as unvisited and unrooted
	for i := 1; i < number_of_vertices+1; i++ {
		explored[i] = false
        roots[i] = 0
	}
	// First traversal of the graph to compute the nodes ordering
	nodes_order := make([]int, 0)
	for i := number_of_vertices; i > 0; i-- {
		nodes_order = Visit(i, graph_edges, explored, nodes_order)
	}
    // Free the explored map
    explored = nil
	// Second traversal of the graph to find the root of each node
	for n := range(nodes_order) {
		Assign(nodes_order[n], nodes_order[n], graph_edges, roots)
	}
    // Free the nodes_order slice
    nodes_order = nil

	return roots
}


func main() {
    /* Testing case
    TestEdges := [][2]int{[2]int{1, 7}, [2]int{7, 4}, [2]int{4, 1}, [2]int{7, 9},
			[2]int{9, 6}, [2]int{6, 3}, [2]int{3, 9}, [2]int{6, 8},
			[2]int{8, 2}, [2]int{2, 5}, [2]int{5, 8}} */
    // Assignment data
    // Open the file and populate the graph edges slice
    graph_edges := make([][2]int, 0)
    graph_file, err := os.Open("SCC.txt")
    if err != nil {
        panic(err)
    }
    buffer := bufio.NewReader(graph_file)
    line, err := buffer.ReadString('\n')
    for err == nil {
        edge_string := line[:len(line)-1]
        edges := strings.Split(edge_string, " ")
        tail, ter := strconv.Atoi(edges[0])
        if ter != nil {
            panic(ter)
        }
        head, her := strconv.Atoi(edges[1])
        if her != nil {
            panic(her)
        }
        graph_edges = append(graph_edges, [2]int{tail, head})
        line, err = buffer.ReadString('\n')
    }
    graph_file.Close()
    // Execute the algorithm
    graph_nodes_number := 875714
	scc := Kosaraju(graph_edges, graph_nodes_number)
    fmt.Println(scc)
    // Free the graph edges slice
    graph_edges = nil
    // Compute the size of each SCC
    scc_distribution := make(map[int]int)
    for i := range(scc) {
        scc_distribution[scc[i]]++
    }
    fmt.Println(scc_distribution)
}
