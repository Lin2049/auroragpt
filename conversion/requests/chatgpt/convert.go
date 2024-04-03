package chatgpt

import (
	chatgpt_types "aurora/typings/chatgpt"
	official_types "aurora/typings/official"
	"strings"
)

func ConvertAPIRequest(api_request official_types.APIRequest, puid string, requireArk bool, proxy string) chatgpt_types.ChatGPTRequest {
	chatgpt_request := chatgpt_types.NewChatGPTRequest()
	if strings.HasPrefix(api_request.Model, "gpt-3.5") {
		chatgpt_request.Model = "text-davinci-002-render-sha"
	} else if strings.HasPrefix(api_request.Model, "gpt-4") {
		chatgpt_request.Model = api_request.Model
		// Cover some models like gpt-4-32k
		if len(api_request.Model) >= 7 && api_request.Model[6] >= 48 && api_request.Model[6] <= 57 {
			chatgpt_request.Model = "gpt-4"
		}
	}
	if api_request.PluginIDs != nil {
		chatgpt_request.PluginIDs = api_request.PluginIDs
		chatgpt_request.Model = "gpt-4-plugins"
	}
	for _, api_message := range api_request.Messages {
		if api_message.Role == "system" {
			api_message.Role = "critic"
		}
		chatgpt_request.AddMessage(api_message.Role, api_message.Content)
	}
	return chatgpt_request
}
