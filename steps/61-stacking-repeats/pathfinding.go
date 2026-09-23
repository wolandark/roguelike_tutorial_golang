package main

import "container/heap"

func FindPath(m *GameMap, x1, y1, x2, y2 int) [][2]int {
	const blockedCost = 10

	// cost[y][x] is the cheapest known cost to reach a cell (-1: not reached),
	// prev[y][x] is the cell we came from on that cheapest route.
	cost := make([][]int, m.Height)
	prev := make([][][2]int, m.Height)
	for y := range cost {
		cost[y] = make([]int, m.Width)
		prev[y] = make([][2]int, m.Width)
		for x := range cost[y] {
			cost[y][x] = -1
		}
	}
	cost[y1][x1] = 0
	pq := &pathQueue{{x: x1, y: y1, cost: 0}}

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(pathNode)
		if cur.x == x2 && cur.y == y2 {
			break
		}
		if cur.cost > cost[cur.y][cur.x] {
			continue // a cheaper route to this tile was found already
		}
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				nx, ny := cur.x+dx, cur.y+dy
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
				if cost[ny][nx] == -1 || cur.cost+step < cost[ny][nx] {
					cost[ny][nx] = cur.cost + step
					prev[ny][nx] = [2]int{cur.x, cur.y}
					heap.Push(pq, pathNode{x: nx, y: ny, cost: cost[ny][nx]})
				}
			}
		}
	}

	if cost[y2][x2] == -1 {
		return nil
	}
	// walk back from the goal to the start, then reverse
	var path [][2]int
	for p := [2]int{x2, y2}; p != [2]int{x1, y1}; p = prev[p[1]][p[0]] {
		path = append(path, p)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

type pathNode struct{ x, y, cost int }

type pathQueue []pathNode

func (q pathQueue) Len() int           { return len(q) }
func (q pathQueue) Less(i, j int) bool { return q[i].cost < q[j].cost }
func (q pathQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *pathQueue) Push(x any)        { *q = append(*q, x.(pathNode)) }
func (q *pathQueue) Pop() any          { old := *q; n := len(old); x := old[n-1]; *q = old[:n-1]; return x }
