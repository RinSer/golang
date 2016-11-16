package main

/**
 * My implementation of Dijkstra shortest path algorithm
 *
 * created by RinSer
 */

import (
	"fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)


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
 * Implementation of Dijkstra algorithm that 
 * finds the minimum path for each vertex from
 * the start vertex.
 * Returns a list of minimum paths.
 */
func Dijkstra(graph []map[int]int, start int) []int {
    infinity := 1000000
    graph_length := len(graph)
    distance := make([]int, graph_length)
    path := make([][]int, graph_length)
    unvisited := make(map[int]bool)
    // Initialize paths and lengths
    for i := 0; i < graph_length; i++ {
        if i+1 == start {
            distance[i] = 0
            path[i] = make([]int, 1)
            path[i][0] = start
        } else {
            distance[i] = infinity
            path[i] = make([]int, 0)
        }
        unvisited[i] = true
    }
    // Main loop
    for len(unvisited) > 0 {
        // Extract the unvisited node with min distance and mark it visited
        var current_node int
        min_distance := infinity+1
        for u := range(unvisited) {
            if distance[u] < min_distance {
                min_distance = distance[u]
                current_node = u
            }
        }
        delete(unvisited, current_node)
        // Compute the min distances from the current node to all unvisited ones
        for v := range(graph[current_node]) {
            if unvisited[v-1] {
                alt_distance := distance[current_node]+graph[current_node][v]
                if alt_distance < distance[v-1] {
                    distance[v-1] = alt_distance
                    path[v-1] = append(path[current_node], v)
                }
            }
        }
    }
    //fmt.Println(path)
    return distance
}


func main() {
    // Open the graph file
    file_path := "dijkstraData.txt"
    file_pointer, err := os.Open(file_path)
    if err != nil {
        panic(err)
    }
    // Create the assignment graph
    AssignmentGraph := make([]map[int]int, 200)
    buffer := bufio.NewReader(file_pointer)
    line, err := buffer.ReadString('\n')
    for err == nil {
        row := strings.Fields(line)
        node := atoi(row[0])
        row = row[1:]
        arcs := make(map[int]int)
        for i := range(row) {
            arc := strings.SplitN(row[i], ",", 2)
            arc_head := atoi(arc[0])
            arc_len := atoi(arc[1])
            arcs[arc_head] = arc_len
        }
        AssignmentGraph[node-1] = arcs
        line, err = buffer.ReadString('\n')
    }
    file_pointer.Close()
    // Run the Dijkstra algorithm on the assignment graph
    distances := Dijkstra(AssignmentGraph, 1)
    //fmt.Println(distances)
    vertices := [10]int{7,37,59,82,99,115,133,165,188,197}
    answer := ""
    for v := range(vertices) {
        dist := distances[vertices[v]-1]
        sdist := strconv.Itoa(dist)
        answer = answer + "," + sdist
    }
    fmt.Println(answer)

    // Test graph
    //TestGraph := []map[int]int{map[int]int{2:1, 3:4}, map[int]int{1:1, 3:2, 4:6}, map[int]int{4:3}, map[int]int{}}
    //fmt.Println(Dijkstra(TestGraph, 1))
}

