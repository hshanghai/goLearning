package People

type China struct {
	Name string
}

func (p China) Eat(thing string) string {
	return "正在吃"+thing
}


