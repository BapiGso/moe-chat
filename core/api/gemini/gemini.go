package gemini

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io"
	"log/slog"
	"moechat/core/api/part"
	"moechat/core/database"
	"strings"
)

type Client struct {
	resStream *genai.GenerateContentResponseIterator
}

func (c *Client) Read(p []byte) (n int, err error) {
	resp, err := c.resStream.Next()
	if errors.Is(err, iterator.Done) {
		return 0, io.EOF
	}
	if err != nil {
		return 0, err
	}
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				n = copy(p, fmt.Sprintf("%s", part))
				return n, nil
			}
		}
	}
	return 0, nil
}

func (c *Client) CreateResStream(ctx echo.Context, baseModel string, msgs []part.Message) error {
	var dbModel database.Model
	var err error
	if err := database.DB.Get(&dbModel, `SELECT * from model WHERE provider = 'Gemini' AND active = 1`); err != nil {
		return err
	}
	client, err := genai.NewClient(ctx.Request().Context(), option.WithAPIKey(dbModel.APIKey), option.WithLogger(slog.Default()))
	if err != nil {
		return err
	}
	//file, err := client.UploadFileFromPath(nil, "Cajun_instruments.jpg", nil)
	//if err != nil {
	//}
	//uploadFile, err := client.UploadFile(file)
	//if err != nil {
	//	return err
	//}
	//openFile, err := os.OpenFile("test.jpg", os.O_RDWR|os.O_CREATE, 0666)
	//if err != nil {
	//	return err
	//}
	model := client.GenerativeModel(baseModel)
	model.SetTemperature(0.9)
	model.SetTopP(0.5)
	model.SetTopK(20)
	model.SetMaxOutputTokens(1000)
	cs := model.StartChat()
	history, lastMessage, err := transformToProviderMessages(ctx, msgs)
	if err != nil {
		return err
	}
	cs.History = history
	//model.SystemInstruction = genai.NewUserContent(genai.Text("You are Yoda from Star Wars."))
	//model.ResponseMIMEType = "application/json"
	c.resStream = cs.SendMessageStream(ctx.Request().Context(), genai.Text(lastMessage.Content))

	return err
}

func (c *Client) Ping() {

}

func (c *Client) GetModelList() []string {
	var dbModel database.Model
	var err error
	if err := database.DB.Get(&dbModel, `SELECT * from model WHERE provider = 'Gemini' AND active = 1`); err != nil {
		return nil
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(dbModel.APIKey))
	if err != nil {
		return nil
	}
	defer client.Close()
	var ids []string

	iter := client.ListModels(ctx)
	for {
		m, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil
		}
		parts := strings.Split(m.Name, "/")
		if len(parts) > 1 {
			ids = append(ids, parts[1])
		}

	}
	return ids
}
