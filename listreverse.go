package piscine

func ListReverse(l *List) {
	y := l.Head
	l.Tail = l.Head
	var x, z *NodeL
	for y != nil {
		z = y.Next
		y.Next = x
		x = y
		y = z
	}
	l.Head = x
}
