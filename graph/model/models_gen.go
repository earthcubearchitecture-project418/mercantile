// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Do struct {
	Subject      string        `json:"subject"`
	Score        string        `json:"score"`
	Type         string        `json:"type"`
	Name         string        `json:"name"`
	Relto        string        `json:"relto"`
	Addtype      string        `json:"addtype"`
	URL          string        `json:"url"`
	Description  string        `json:"description"`
	Distribution *Distribution `json:"distribution"`
}

type Distribution struct {
	Type           string `json:"type"`
	ContentURL     string `json:"contentUrl"`
	EncodingFormat string `json:"encodingFormat"`
}

type NewDo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}
