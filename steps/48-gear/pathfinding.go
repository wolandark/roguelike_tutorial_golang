package main

import "container/heap"

func FindPath(m *GameMap, x1, y1, x2, y2 int) [][2]int {
	const blockedCost = 10

	cost := make([]int, m.Width*m.Height) // cheapest known cost to reach each tile
	prev := make([]int, m.Width*m.Height) // which tile we came from
	for i := range cost {
		cost[i] = -1
	}
	start, goal := y1*m.Width+x1, y2*m.Width+x2
	cost[start] = 0
	pq := &pathQueue{{idx: start, cost: 0}}

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(pathNode)
		if cur.idx == goal {
			break
		}
		if cur.cost > cost[cur.idx] {
			continue // a cheaper route to this tile was found already
		}
		cx, cy := cur.idx%m.Width, cur.idx/m.Width
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				nx, ny := cx+dx, cy+dy
				if !m.InBounds(nx, ny) || !m.TileAt(nx, ny).Walkable {
					continue
				}
				step := 2
				if dx != 0 && dy != 0 {
					step = 3
				}
				if m.GetBlockingEntityAt(nx, ny) != nil {
					step += blockedCost
				}
				ni := ny*m.Width + nx
				if cost[ni] == -1 || cur.cost+step < cost[ni] {
					cost[ni] = cur.cost + step
					prev[ni] = cur.idx
					heap.Push(pq, pathNode{idx: ni, cost: cost[ni]})
				}
			}
		}
	}

	if cost[goal] == -1 {
		return nil
	}
	// walk back from the goal to the start, then reverse
	var path [][2]int
	for i := goal; i != start; i = prev[i] {
		path = append(path, [2]int{i % m.Width, i / m.Width})
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

type pathNode struct{ idx, cost int }

type pathQueue []pathNode

func (q pathQueue) Len() int           { return len(q) }
func (q pathQueue) Less(i, j int) bool { return q[i].cost < q[j].cost }
func (q pathQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *pathQueue) Push(x any)        { *q = append(*q, x.(pathNode)) }
func (q *pathQueue) Pop() any          { old := *q; n := len(old); x := old[n-1]; *q = old[:n-1]; return x }
