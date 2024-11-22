package claude

import "github.com/liushuangls/go-anthropic/v2"

type Client struct {
	resStream anthropic.MessagesResponse
}

func (c *Client) Ping() {

}

func (c *Client) GetModelList() {

}

//func (c *Client) TransformToProviderMessages(msgs []types.UnifiedMessage) (interface{}, error) {
//	anthropicMsgs := make([]map[string]interface{}, len(msgs))
//	for i, msg := range msgs {
//		anthropicMsg := map[string]interface{}{
//			"role": msg.Role,
//		}
//
//		// 处理多模态内容
//		if len(msg.MultiContent) > 0 {
//			content := make([]map[string]interface{}, len(msg.MultiContent))
//			for j, part := range msg.MultiContent {
//				content[j] = map[string]interface{}{
//					"type": part.Type,
//					"text": part.Text,
//				}
//				if part.ImageURL != "" {
//					content[j]["image_url"] = part.ImageURL
//				}
//			}
//			anthropicMsg["content"] = content
//		} else {
//			anthropicMsg["content"] = msg.Content
//		}
//
//		anthropicMsgs[i] = anthropicMsg
//	}
//	return anthropicMsgs, nil
//}
