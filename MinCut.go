package main

/**
 * My implementation of Karger's Graph Minimum Cut 
 * search randomized algorithm.
 *
 * created by RinSer
 */

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)


// Data structure to store a graph
type Graph struct {
    vertices []int
    edges [][2]int
}
// Initialize an empty graph
func NewGraph() Graph {
    var graph Graph
    graph.vertices = make([]int, 0)
    graph.edges = make([][2]int, 0)

    return graph
}


/**
 * Helper function to convert a string to an integer
 */
func atoi(word string) int {
    number, err := strconv.Atoi(word)
    if err != nil {
        panic(err)
    }

    return number
}


/**
 * Function to compute the Minimum Cut in a graph.
 * Returns a graph reduced to two vertices.
 *
func MinCut(graph Graph) Graph {
    // Uniformly choose an edge
    
    return 0
}*/


func main() {
    // Open the graph file
    file_pointer, err := os.Open("kargerMinCut.txt")
    if err != nil {
        panic(err)
    }
    // Initialize the adjacency list
    buffer := bufio.NewReader(file_pointer)
    line, err := buffer.ReadString('\n')
    graph := NewGraph()
    added_edges := make(map[[2]int]bool)
    for err == nil {
        row := strings.Fields(line)
        vertex := atoi(row[0])
        graph.vertices = append(graph.vertices, vertex)
        row = row[1:]
        for i := range row {
            var edge [2]int
            edge_tail := atoi(row[i])
            if edge_tail < vertex {
                edge[0] = edge_tail
                edge[1] = vertex
            } else {
                edge[0] = vertex
                edge[1] = edge_tail
            }
            if !added_edges[edge] {
                graph.edges = append(graph.edges, edge)
                added_edges[edge] = true
            }
        }
        line, err = buffer.ReadString('\n')
    }
    fmt.Println(len(graph.vertices), len(graph.edges))
}
