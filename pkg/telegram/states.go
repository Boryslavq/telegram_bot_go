package telegram

type SignState int

const (
	StateText SignState = iota
	StateOver
	StateCourse
)

type Support struct {
	State SignState // 0 - start, over - 1
	Text  string
}
