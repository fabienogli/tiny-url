package handler

type response struct {
	Url  string `json:"url"`
	Slug string `json:"slug"`
}

type createQuery struct {
	Url string `json:"url"`
}
