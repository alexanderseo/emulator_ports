package dto

type Answer struct {
	Number int `json:"number"`
	Value  int `json:"value"`
}

type AnswerOut struct {
	Number int
	Value  int
}
