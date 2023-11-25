package match

import "log"

type Intro struct {
	intro string
}

func NewIntro(intro string) *Intro {
	lenOfIntro := len(intro)
	if lenOfIntro == 0 || lenOfIntro > 18 {
		log.Fatal("intro length must in range 1 to 18")
	}
	return &Intro{
		intro: intro,
	}
}

func (i *Intro) String() string {
	return i.intro
}
