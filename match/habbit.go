package match

import "log"

type Habbit struct {
	habbit string
}

func NewHabbit(habbit string) *Habbit {
	lengthOfHabbit := len(habbit)
	if lengthOfHabbit == 0 || lengthOfHabbit > 10 {
		log.Fatal("length Of habbit must in range 1 to 10")
	}
	return &Habbit{
		habbit: habbit,
	}
}
func (h *Habbit) String() string {
	return h.habbit
}
