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

// DFS does depth-first search, and returns the list of vertices.
// (https://en.wikipedia.org/wiki/Depth-first_search)
//
//	 0. DFS(G, v):
//	 1.
//	 2. 	let S be a stack
//	 3. 	S.push(v)
//	 4.
//	 5. 	while S is not empty:
//	 6.
//	 7. 		u = S.pop()
//	 8.
//	 9. 		if u is not visited yet:
//	10.
//	11. 			label u as visited
//	12.
//	13. 			for each vertex w adjacent to u:
//	14.
//	15. 				if w is not visited yet:
//	16. 					S.push(w)
//
func DFS(g Graph, vtx string) []string {

	if !g.FindVertex(vtx) {
		return nil
	}

	s := []string{vtx}
	visited := make(map[string]bool)
	rs := []string{}

	// while S is not empty:
	for len(s) != 0 {

		u := s[len(s)-1]
		s = s[:len(s)-1 : len(s)-1]

		// if u is not visited yet:
		if _, ok := visited[u]; !ok {
			// label u as visited
			visited[u] = true

			rs = append(rs, u)

			// for each vertex w adjacent to u:
			cmap, _ := g.GetChildren(u)
			for w := range cmap {
				// if w is not visited yet:
				if _, ok := visited[w]; !ok {
					s = append(s, w) // S.push(w)
				}
			}
			pmap, _ := g.GetParents(u)
			for w := range pmap {
				// if w is not visited yet:
				if _, ok := visited[w]; !ok {
					s = append(s, w) // S.push(w)
				}
			}
		}
	}

	return rs
}

func main() {
	f, err := os.Open("graph.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g, err := NewDefaultGraphFromJSON(f, "graph_00")
	if err != nil {
		panic(err)
	}
	rs := DFS(g, "S")
	fmt.Println("DFS:", rs) // [S A B C D T E F]
	if len(rs) != 8 {
		log.Fatalf("should be 8 vertices but %s", g)
	}
}

// Graph describes the methods of graph operations.
// It assumes that the identifier of a Vertex is string and unique.
// And weight values is float64.
type Graph interface {
	// GetVertices returns a map of all vertices.
	GetVertices() map[string]bool

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
	// (Vertices that comes to the argument vertex.)
	GetParents(vtx string) (map[string]bool, error)

	// GetChildren returns the map of child vertices.
	// (Vertices that goes out of the argument vertex.)
	GetChildren(vtx string) (map[string]bool, error)
}

// DefaultGraph type implements all methods in Graph interface.
type DefaultGraph struct {
	sync.Mutex

	// Vertices stores all vertices.
	Vertices map[string]bool

	// VertexToChildren maps a Vertex identifer to children with edge weights.
	VertexToChildren map[string]map[string]float64

	// VertexToParents maps a Vertex identifer to parents with edge weights.
	VertexToParents map[string]map[string]float64
}

// NewDefaultGraph returns a new DefaultGraph.
func NewDefaultGraph() *DefaultGraph {
	return &DefaultGraph{
		Vertices:         make(map[string]bool),
		VertexToChildren: make(map[string]map[string]float64),
		VertexToParents:  make(map[string]map[string]float64),
		//
		// without this
		// panic: assignment to entry in nil map
	}
}

func (g *DefaultGraph) Init() {
	// (X) g = NewDefaultGraph()
	// this only updates the pointer
	//
	*g = *NewDefaultGraph()
}

func (g DefaultGraph) GetVertices() map[string]bool {
	g.Lock()
	defer g.Unlock()
	return g.Vertices
}

func (g DefaultGraph) FindVertex(vtx string) bool {
	g.Lock()
	defer g.Unlock()
	if _, ok := g.Vertices[vtx]; !ok {
		return false
	}
	return true
}

func (g *DefaultGraph) AddVertex(vtx string) bool {
	g.Lock()
	defer g.Unlock()
	if _, ok := g.Vertices[vtx]; !ok {
		g.Vertices[vtx] = true
		return true
	}
	return false
}

func (g *DefaultGraph) DeleteVertex(vtx string) bool {
	g.Lock()
	defer g.Unlock()
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

func (g *DefaultGraph) AddEdge(vtx1, vtx2 string, weight float64) error {
	g.Lock()
	defer g.Unlock()
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

func (g *DefaultGraph) ReplaceEdge(vtx1, vtx2 string, weight float64) error {
	g.Lock()
	defer g.Unlock()
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

func (g *DefaultGraph) DeleteEdge(vtx1, vtx2 string) error {
	g.Lock()
	defer g.Unlock()
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

func (g *DefaultGraph) GetWeight(vtx1, vtx2 string) (float64, error) {
	g.Lock()
	defer g.Unlock()
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

func (g *DefaultGraph) GetParents(vtx string) (map[string]bool, error) {
	g.Lock()
	defer g.Unlock()
	if _, ok := g.Vertices[vtx]; !ok {
		return nil, fmt.Errorf("%s does not exist in the graph.", vtx)
	}
	rs := make(map[string]bool)
	if _, ok := g.VertexToParents[vtx]; ok {
		for k := range g.VertexToParents[vtx] {
			rs[k] = true
		}
	}
	return rs, nil
}

func (g *DefaultGraph) GetChildren(vtx string) (map[string]bool, error) {
	g.Lock()
	defer g.Unlock()
	if _, ok := g.Vertices[vtx]; !ok {
		return nil, fmt.Errorf("%s does not exist in the graph.", vtx)
	}
	rs := make(map[string]bool)
	if _, ok := g.VertexToChildren[vtx]; ok {
		for k := range g.VertexToChildren[vtx] {
			rs[k] = true
		}
	}
	return rs, nil
}

// FromJSON creates a graph Data from JSON. Here's the sample JSON data:
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
func NewDefaultGraphFromJSON(rd io.Reader, graphID string) (*DefaultGraph, error) {
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
	g := NewDefaultGraph()
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

func (g DefaultGraph) String() string {
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
