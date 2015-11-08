package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// TopologicalSort does topological sort(ordering) with DFS.
// It returns true if the graph is a DAG (no cycle, with a topological sort).
// False if the graph is not a DAG (cycle, with no topological sort).
//
//	 0. TopologicalSort(G)
//	 1.
//	 2. 	L = Empty list that will contain the sorted nodes
//	 3. 	isDAG = true
//	 4.
//	 5. 	for each vertex v in G:
//	 6.
//	 7. 		if v.color == "white":
//	 8.
//	 9. 			topologicalSortVisit(v, L, isDAG)
//	10.
//	11.
//	12.
//	13.
//	14. topologicalSortVisit(v, L, isDAG)
//	15.
//	16. 	if v.color == "gray":
//	17. 		isDAG = false
//	18. 		return
//	19.
//	20. 	if v.color == "white":
//	21.
//	22. 		v.color = "gray":
//	23.
//	24.			for each child vertex w of v:
//	25. 			topologicalSortVisit(w, L, isDAG)
//	26.
//	27. 		v.color = "black"
//	28.			L.push_front(v)
//
func TopologicalSort(g Graph) ([]string, bool) {

	// L = Empty list that will contain the sorted nodes
	L := []string{}
	isDAG := true
	color := make(map[string]string)
	for v := range g.GetVertices() {
		color[v] = "white"
	}

	// for each vertex v in G:
	for v := range g.GetVertices() {
		// if v.color == "white":
		if color[v] == "white" {
			// topologicalSortVisit(v, L, isDAG)
			topologicalSortVisit(g, v, &L, &isDAG, &color)
		}
	}

	return L, isDAG
}

func topologicalSortVisit(
	g Graph,
	vtx string,
	L *[]string,
	isDAG *bool,
	color *map[string]string,
) {

	// if v.color == "gray":
	if (*color)[vtx] == "gray" {
		// isDAG = false
		*isDAG = false
		return
	}

	// if v.color == "white":
	if (*color)[vtx] == "white" {
		// v.color = "gray":
		(*color)[vtx] = "gray"

		// for each child vertex w of v:
		cmap, err := g.GetChildren(vtx)
		if err != nil {
			panic(err)
		}
		for w := range cmap {
			// topologicalSortVisit(w, L, isDAG)
			topologicalSortVisit(g, w, L, isDAG, color)
		}

		// v.color = "black"
		(*color)[vtx] = "black"

		// L.push_front(v)
		temp := make([]string, len(*L)+1)
		temp[0] = vtx
		copy(temp[1:], *L)
		*L = temp
	}
}

func main() {
	f, err := os.Open("graph.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g, err := NewDefaultGraphFromJSON(f, "graph_06")
	if err != nil {
		panic(err)
	}
	L, isDAG := TopologicalSort(g)
	if isDAG == false {
		log.Fatalf("there is no directed cycle in the graph so isDAG should be true but %+v %+v", L, isDAG)
	}
	fmt.Println("graph_06:", L)
	// graph_06: [C B D F A H E G]
}

// Graph describes the methods of graph operations.
// It assumes that the identifier of a Vertex is string and unique.
// And weight values is float64.
type Graph interface {
	// Init initializes a Graph.
	Init()

	// GetVertices returns a map from vertex ID to
	// empty struct value. Graph does not allow duplicate
	// vertex ID.
	GetVertices() map[string]struct{}

	// FindVertex returns true if the vertex already
	// exists in the graph.
	FindVertex(vtx string) bool

	// AddVertex adds a vertex to a graph, and returns false
	// if the vertex already existed in the graph.
	AddVertex(vtx string) bool

	// DeleteVertex deletes a vertex from a graph.
	// It returns true if it got deleted.
	// And false if it didn't get deleted.
	DeleteVertex(vtx string) bool

	// AddEdge adds an edge from vtx1 to vtx2 with the weight.
	AddEdge(vtx1, vtx2 string, weight float64) error

	// ReplaceEdge replaces an edge from vtx1 to vtx2 with the weight.
	ReplaceEdge(vtx1, vtx2 string, weight float64) error

	// DeleteEdge deletes an edge from vtx1 to vtx2.
	DeleteEdge(vtx1, vtx2 string) error

	// GetWeight returns the weight from vtx1 to vtx2.
	GetWeight(vtx1, vtx2 string) (float64, error)

	// GetParents returns the map of parent vertices.
	// (Vertices that come towards the argument vertex.)
	GetParents(vtx string) (map[string]struct{}, error)

	// GetChildren returns the map of child vertices.
	// (Vertices that go out of the argument vertex.)
	GetChildren(vtx string) (map[string]struct{}, error)

	// String describes the Graph.
	String() string
}

// defaultGraph is an internal default graph type that
// implements all methods in Graph interface.
type defaultGraph struct {
	mu sync.Mutex // guards the following

	// Vertices stores all vertices.
	Vertices map[string]struct{}

	// VertexToChildren maps a Vertex identifer to children with edge weights.
	VertexToChildren map[string]map[string]float64

	// VertexToParents maps a Vertex identifer to parents with edge weights.
	VertexToParents map[string]map[string]float64
}

// newDefaultGraph returns a new defaultGraph.
func newDefaultGraph() *defaultGraph {
	return &defaultGraph{
		Vertices:         make(map[string]struct{}),
		VertexToChildren: make(map[string]map[string]float64),
		VertexToParents:  make(map[string]map[string]float64),
		//
		// without this
		// panic: assignment to entry in nil map
	}
}

// NewDefaultGraph returns a new defaultGraph.
func NewDefaultGraph() Graph {
	return newDefaultGraph()
}

// newDefaultGraphFromJSON creates a graph Data from JSON.
// Here's the sample JSON data:
//
//	{
//	    "graph_00": {
//	        "S": {
//	            "A": 100,
//	            "B": 14,
//	            "C": 200
//	        },
//	        "A": {
//	            "S": 15,
//	            "B": 5,
//	            "D": 20,
//	            "T": 44
//	        },
//	        "B": {
//	            "S": 14,
//	            "A": 5,
//	            "D": 30,
//	            "E": 18
//	        },
//	        "C": {
//	            "S": 9,
//	            "E": 24
//	        },
//	        "D": {
//	            "A": 20,
//	            "B": 30,
//	            "E": 2,
//	            "F": 11,
//	            "T": 16
//	        },
//	        "E": {
//	            "B": 18,
//	            "C": 24,
//	            "D": 2,
//	            "F": 6,
//	            "T": 19
//	        },
//	        "F": {
//	            "D": 11,
//	            "E": 6,
//	            "T": 6
//	        },
//	        "T": {
//	            "A": 44,
//	            "D": 16,
//	            "F": 6,
//	            "E": 19
//	        }
//	    },
//	}
//
func newDefaultGraphFromJSON(rd io.Reader, graphID string) (*defaultGraph, error) {
	js := make(map[string]map[string]map[string]float64)
	dec := json.NewDecoder(rd)
	for {
		if err := dec.Decode(&js); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	if _, ok := js[graphID]; !ok {
		return nil, fmt.Errorf("%s does not exist", graphID)
	}
	gmap := js[graphID]
	g := newDefaultGraph()
	for vtx1, mm := range gmap {
		if !g.FindVertex(vtx1) {
			g.AddVertex(vtx1)
		}
		for vtx2, weight := range mm {
			if !g.FindVertex(vtx2) {
				g.AddVertex(vtx2)
			}
			g.ReplaceEdge(vtx1, vtx2, weight)
		}
	}
	return g, nil
}

// NewDefaultGraphFromJSON returns a new defaultGraph from a JSON file.
func NewDefaultGraphFromJSON(rd io.Reader, graphID string) (Graph, error) {
	return newDefaultGraphFromJSON(rd, graphID)
}

func (g *defaultGraph) Init() {
	// (X) g = newDefaultGraph()
	// this only updates the pointer
	//
	*g = *newDefaultGraph()
}

func (g *defaultGraph) GetVertices() map[string]struct{} {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.Vertices
}

func (g *defaultGraph) FindVertex(vtx string) bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Vertices[vtx]; !ok {
		return false
	}
	return true
}

func (g *defaultGraph) AddVertex(vtx string) bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Vertices[vtx]; !ok {
		g.Vertices[vtx] = struct{}{}
		return true
	}
	return false
}

func (g *defaultGraph) DeleteVertex(vtx string) bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Vertices[vtx]; !ok {
		return false
	} else {
		delete(g.Vertices, vtx)
	}
	if _, ok := g.VertexToChildren[vtx]; ok {
		delete(g.VertexToChildren, vtx)
	}
	for _, smap := range g.VertexToChildren {
		if _, ok := smap[vtx]; ok {
			delete(smap, vtx)
		}
	}
	if _, ok := g.VertexToParents[vtx]; ok {
		delete(g.VertexToParents, vtx)
	}
	for _, smap := range g.VertexToParents {
		if _, ok := smap[vtx]; ok {
			delete(smap, vtx)
		}
	}
	return true
}

func (g *defaultGraph) AddEdge(vtx1, vtx2 string, weight float64) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Vertices[vtx1]; !ok {
		return fmt.Errorf("%s does not exist in the graph.", vtx1)
	}
	if _, ok := g.Vertices[vtx2]; !ok {
		return fmt.Errorf("%s does not exist in the graph.", vtx2)
	}
	if _, ok := g.VertexToChildren[vtx1]; ok {
		if v, ok2 := g.VertexToChildren[vtx1][vtx2]; ok2 {
			g.VertexToChildren[vtx1][vtx2] = v + weight
		} else {
			g.VertexToChildren[vtx1][vtx2] = weight
		}
	} else {
		tmap := make(map[string]float64)
		tmap[vtx2] = weight
		g.VertexToChildren[vtx1] = tmap
	}
	if _, ok := g.VertexToParents[vtx2]; ok {
		if v, ok2 := g.VertexToParents[vtx2][vtx1]; ok2 {
			g.VertexToParents[vtx2][vtx1] = v + weight
		} else {
			g.VertexToParents[vtx2][vtx1] = weight
		}
	} else {
		tmap := make(map[string]float64)
		tmap[vtx1] = weight
		g.VertexToParents[vtx2] = tmap
	}
	return nil
}

func (g *defaultGraph) ReplaceEdge(vtx1, vtx2 string, weight float64) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Vertices[vtx1]; !ok {
		return fmt.Errorf("%s does not exist in the graph.", vtx1)
	}
	if _, ok := g.Vertices[vtx2]; !ok {
		return fmt.Errorf("%s does not exist in the graph.", vtx2)
	}
	if _, ok := g.VertexToChildren[vtx1]; ok {
		g.VertexToChildren[vtx1][vtx2] = weight
	} else {
		tmap := make(map[string]float64)
		tmap[vtx2] = weight
		g.VertexToChildren[vtx1] = tmap
	}
	if _, ok := g.VertexToParents[vtx2]; ok {
		g.VertexToParents[vtx2][vtx1] = weight
	} else {
		tmap := make(map[string]float64)
		tmap[vtx1] = weight
		g.VertexToParents[vtx2] = tmap
	}
	return nil
}

func (g *defaultGraph) DeleteEdge(vtx1, vtx2 string) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Vertices[vtx1]; !ok {
		return fmt.Errorf("%s does not exist in the graph.", vtx1)
	}
	if _, ok := g.Vertices[vtx2]; !ok {
		return fmt.Errorf("%s does not exist in the graph.", vtx2)
	}
	if _, ok := g.VertexToChildren[vtx1]; ok {
		if _, ok := g.VertexToChildren[vtx1][vtx2]; ok {
			delete(g.VertexToChildren[vtx1], vtx2)
		}
	}
	if _, ok := g.VertexToParents[vtx2]; ok {
		if _, ok := g.VertexToParents[vtx2][vtx1]; ok {
			delete(g.VertexToParents[vtx2], vtx1)
		}
	}
	return nil
}

func (g *defaultGraph) GetWeight(vtx1, vtx2 string) (float64, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Vertices[vtx1]; !ok {
		return 0.0, fmt.Errorf("%s does not exist in the graph.", vtx1)
	}
	if _, ok := g.Vertices[vtx2]; !ok {
		return 0.0, fmt.Errorf("%s does not exist in the graph.", vtx2)
	}
	if _, ok := g.VertexToChildren[vtx1]; ok {
		if v, ok := g.VertexToChildren[vtx1][vtx2]; ok {
			return v, nil
		}
	}
	return 0.0, fmt.Errorf("there is not edge from %s to %s", vtx1, vtx2)
}

func (g *defaultGraph) GetParents(vtx string) (map[string]struct{}, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Vertices[vtx]; !ok {
		return nil, fmt.Errorf("%s does not exist in the graph.", vtx)
	}
	rs := make(map[string]struct{})
	if _, ok := g.VertexToParents[vtx]; ok {
		for k := range g.VertexToParents[vtx] {
			rs[k] = struct{}{}
		}
	}
	return rs, nil
}

func (g *defaultGraph) GetChildren(vtx string) (map[string]struct{}, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.Vertices[vtx]; !ok {
		return nil, fmt.Errorf("%s does not exist in the graph.", vtx)
	}
	rs := make(map[string]struct{})
	if _, ok := g.VertexToChildren[vtx]; ok {
		for k := range g.VertexToChildren[vtx] {
			rs[k] = struct{}{}
		}
	}
	return rs, nil
}

func (g *defaultGraph) String() string {
	buf := new(bytes.Buffer)
	for vtx1 := range g.Vertices {
		cmap, _ := g.GetChildren(vtx1)
		for vtx2 := range cmap {
			weight, _ := g.GetWeight(vtx1, vtx2)
			fmt.Fprintf(buf, "%s -- %.3f --> %s\n", vtx1, weight, vtx2)
		}
	}
	return buf.String()
}
