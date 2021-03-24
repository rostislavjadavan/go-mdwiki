package api

type ErrorResponse struct {
	Message string `json:"message"`
}

type RedirectResponse struct {
	Redirect string `json:"redirect"`
}

type FilenameRequest struct {
	Filename string `json:"filename"`
}

type UpdatePageRequest struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}
