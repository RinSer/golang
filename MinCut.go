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
    "math/rand"
    "math"
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
    // Create the edges copy
    edges := make([][2]int, len(graph.edges))
    copy(edges, graph.edges)
    // Map to store the removed vertices
    removed_vertices := make(map[int]bool)
    // Join the vertices removing randomly an edge until two are left
    for len(graph.vertices)-len(removed_vertices) > 2 {
        // Uniformly choose an edge
        edge_index := rand.Intn(len(edges))
        // Remove the edge
        edge := edges[edge_index]
        edges = append(edges[:edge_index], edges[edge_index+1:]...)
        // Add the removed vertex to the corresponding array
        removed_vertices[edge[1]] = true
        // Add all the edges corresponding to the removed vertex to the joined one
        for i := 0; i < len(edges); i++ {
            if edges[i][0] == edge[1] {
                edges[i][0] = edge[0]
            }
            if edges[i][1] == edge[1] {
                if edges[i][0] > edge[0] {
                    edges[i][1] = edges[i][0]
                    edges[i][0] = edge[0]
                } else {
                    edges[i][1] = edge[0]
                }
            }
            // Remove the loops
            if edges[i][0] == edges[i][1] {
                edges = append(edges[:i], edges[i+1:]...)
                i--
            }
        }
    }
    // Create the graph with minimum cut
    cut_graph := NewGraph()
    cut_graph.edges = edges
    for vdx := range graph.vertices {
        if !removed_vertices[graph.vertices[vdx]] {
            cut_graph.vertices = append(cut_graph.vertices, graph.vertices[vdx])
        }
    }

    return cut_graph
}


/**
 * Function to repeat MinCut algorithm with different seeds.
 * Returns the minimum cut size infimum.
 */
func IterateMinCut(times int, graph Graph) int {
    infimum := len(graph.edges)
    for i := 0; i < times; i++ {
        rand.Seed(int64(i))
        cut_graph := MinCut(graph)
        //fmt.Println(cut_graph)
        current_cut := len(cut_graph.edges)
        if current_cut < infimum {
            infimum = current_cut
        }
    }

    return infimum
}


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
    //fmt.Println(len(graph.vertices), len(graph.edges))
    //fmt.Println(MinCut(graph))
    // Find the Minimum Cut
    n_squared := int(math.Pow(float64(len(graph.vertices)), 2))
    fmt.Println(IterateMinCut(n_squared, graph))

    /* Testing
    test_graph := Graph{vertices: []int{1, 2, 3, 4}, 
    edges: [][2]int{[2]int{1, 2}, [2]int{1, 3}, [2]int{2, 3}, [2]int{2, 4}, [2]int{3, 4}}}
    //rand.Seed(5)
    //fmt.Println(MinCut(test_graph))
    n_cubed := int(math.Pow(float64(len(test_graph.vertices)), 3))
    fmt.Println(IterateMinCut(n_cubed, test_graph))*/
}
