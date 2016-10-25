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
    //"math"
    "math/rand"
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
 */
func MinCut(graph Graph) Graph {
    // Join the vertices removing randomly an edge until two are left
    for len(graph.vertices) > 2 {
        // Uniformly choose an edge
        edge_index := rand.Intn(len(graph.edges))
        // Remove the edge
        edge := graph.edges[edge_index]
        graph.edges = append(graph.edges[:edge_index], graph.edges[edge_index+1:]...)
        // Remove the vertex with smaller index
        for i := 0; i < len(graph.vertices); i++ {
            if graph.vertices[i] == edge[1] {
                graph.vertices = append(graph.vertices[:i], graph.vertices[i+1:]...)
            }
        }
        // Add all the edges corresponding to the removed vertex to the joined one
        for i := 0; i < len(graph.edges); i++ {
            if graph.edges[i][0] == edge[1] {
                graph.edges[i][0] = edge[0]
            }
            if graph.edges[i][1] == edge[1] {
                if graph.edges[i][0] > edge[0] {
                    graph.edges[i][1] = graph.edges[i][0]
                    graph.edges[i][0] = edge[0]
                } else {
                    graph.edges[i][1] = edge[0]
                }
            }
            // Remove the loops
            if graph.edges[i][0] == graph.edges[i][1] {
                graph.edges = append(graph.edges[:i], graph.edges[i+1:]...)
            }
        }
    }
    return graph
}


/**
 * Function to repeat MinCut algorithm with different seeds.
 * Returns the minimum cut size infimum.
 *
func IterateMinCut(times int, graph Graph) int {
    infimum := math.Inf(1)
    for i := 0; i < times; i++ {
*/


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
    // Testing
    test_graph := Graph{vertices: []int{1, 2, 3, 4}, 
    edges: [][2]int{[2]int{1, 2}, [2]int{1, 3}, [2]int{2, 3}, [2]int{2, 4}, [2]int{3, 4}}}
    rand.Seed(5)
    fmt.Println(MinCut(test_graph))
}
