package main

import (
	"math"
	"strings"
	"time"
)

type node struct {
	Id    string `json:"id"`
	Left  string `json:"left"`
	Right string `json:"right"`
	Value int    `json:"value"`
}

type Request struct {
	Tree struct {
		Nodes []node `json:"nodes"`
		Root  string `json:"root"`
	} `json:"tree"`
	Response
}

type Response struct {
	PathSum         int   `json:"path_sum"`
	NodeCount       int   `json:"node_count"`
	ExecuteDuration int64 `json:"execute_duration"`
}

func (n node) hasChild() bool {
	return (strings.TrimSpace(n.Left) != "") || (strings.TrimSpace(n.Right) != "")
}

func (r *Request) getNodeById(Id string) *node {
	for _, n := range r.Tree.Nodes {
		if n.Id == Id {
			if n.hasChild() {
				r.Response.PathSum += n.Value

				left := r.getNodeById(n.Left)
				right := r.getNodeById(n.Right)

				max := 0

				if !left.hasChild() && !right.hasChild() {
					max = int(math.Max(float64(left.Value), float64(right.Value)))
				}

				r.Response.PathSum += max
			}

			return &n
		}
	}

	return &node{}
}

func (r *Request) Init() {
	r.Response.NodeCount = len(r.Tree.Nodes)

	timeStart := time.Now().UnixNano()

	r.getNodeById(r.Tree.Root)

	timeEnd := time.Now().UnixNano()

	r.Response.ExecuteDuration = timeEnd - timeStart
}
