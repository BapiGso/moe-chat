package part

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Files   []struct {
		MimeType string `json:"MimeType"`
		Hash     string `json:"Hash"`
	} `json:"files"`
}
