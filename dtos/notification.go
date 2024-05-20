package dtos

type SendMessageRequest struct {
	Emails []string `json:"emails"`
}