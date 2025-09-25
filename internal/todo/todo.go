package todo

type Item struct {
	Text string
	Done bool
}

func NewItem(text string) Item {
	return Item{
		Text: text,
		Done: false,
	}
}
