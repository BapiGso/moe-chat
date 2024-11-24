package part

import "mime/multipart"

type Message struct {
	Role    string          `form:"role"`
	Content string          `form:"content"`
	Files   *multipart.Form `form:"files"`
}
