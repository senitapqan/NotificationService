package dtos

type SendMessageRequest struct {
	Emails []string `json:"emails"`
	Title string `json:"title"`
}