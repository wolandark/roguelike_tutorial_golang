package main

type Inventory struct {
	Capacity int
	Items    []*Entity
}

func (inv *Inventory) Remove(item *Entity) {
	for i, it := range inv.Items {
		if it == item {
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return
		}
	}
}
