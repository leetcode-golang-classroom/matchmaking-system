package match

type Habbits map[string]*Habbit

func NewHabbits(habbits []string) Habbits {
	habbitsResult := make(map[string]*Habbit)
	for _, v := range habbits {
		habbitsResult[v] = NewHabbit(v)
	}
	return habbitsResult
}

func (h Habbits) String() string {
	result := ""
	size := len(h)
	count := 0
	for _, v := range h {
		result += v.String()
		if count != size-1 {
			result += ","
		}
		count++
	}
	return result
}

func (h Habbits) HasHabbit(habbit string) bool {
	_, ok := h[habbit]
	return ok
}

func (h Habbits) GetHabbits() Habbits {
	return h
}
