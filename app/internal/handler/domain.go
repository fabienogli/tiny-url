package handler

type response struct {
	Url     string `json:"url"`
	TinyURL string `json:"slug"`
}

type createQuery struct {
	Url string `json:"url"`
}
