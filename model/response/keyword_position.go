package response

type KeywordPositionRes struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    KeywordPositionData `json:"data"`
}

type KeywordPositionData struct {
	KeywordPositions []KeywordPosition `json:"keywordPositions"`
}

type KeywordPosition struct {
	Keyword      string     `json:"keyword"`
	SearchResult bool       `json:"searchResult"`
	Positions    []Position `json:"positions"`
}

type Position struct {
	PageNum     int          `json:"pageNum"`
	Coordinates []Coordinate `json:"coordinates"`
}

type Coordinate struct {
	PositionX float64 `json:"positionX"`
	PositionY float64 `json:"positionY"`
}
