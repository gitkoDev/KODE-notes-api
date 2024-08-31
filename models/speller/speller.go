package models

type SpellerRequest struct {
	Lang   string `json:"lang"`
	Format string `json:"format"`
	Text   string `json:"text"`
}

func NewRequest() *SpellerRequest {
	r := &SpellerRequest{}
	r.Lang = "ru"
	r.Format = "plain"
	return r
}

type SpellerResponse []Misspelling

type Misspelling struct {
	Code        int      `json:"code"`
	Pos         int      `json:"pos"`
	Row         int      `json:"row"`
	Col         int      `json:"col"`
	Len         int      `json:"len"`
	Word        string   `json:"word"`
	Suggestions []string `json:"s"`
}
