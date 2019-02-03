package utilities

// RESPONSE
type Response struct {
	Copyright string `json:copyright`
	DataMarvel Data `json:"data"`
}

type Data struct {
	Hero []Results `json:"results"`
}

// INFO GENERAL DEL PERSONAJE
type Results struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Modified string `json:"modified"`
	Comics Comic `json:"comics"`
	Stories Storie `json:"stories"`
	Events Event `json:"events"`
}

// COMICS
type Comic struct {
	Items []ItemsComic `json:"items"`
}
type ItemsComic struct {
	NameComic string `json:"name"`
}

// STORIES
type Storie struct {
	Items []ItemsStorie `json:"items"`
}
type ItemsStorie struct {
	NameStorie string `json:"name"`	
}

// EVENTS
type Event struct {
	Items []ItemsEvent `json:"items"`
}
type ItemsEvent struct {
	NameEvent string `json:"name"`	
}