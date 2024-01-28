package service

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"server/config"
	"strings"
)

func ReadFileContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func GetMarkDownContent(request Request, config config.Config) Response {
	get_markdown_content_request := &GetMarkdownContentRequest{}
	json.Unmarshal([]byte(request.Body), get_markdown_content_request)
	get_markdown_content_response := &GetMarkdownContentResponse{}

	if strings.Contains(get_markdown_content_request.Filepath, "..") {
		return Response{
			Body: "",
		}
	}
	if strings.HasSuffix(get_markdown_content_request.Filepath, ".md") {
		root := config.ServerConf.System.Chroot
		markdown_filepath := filepath.Join(root, get_markdown_content_request.Filepath)

		get_markdown_content_response.Content, _ = ReadFileContent(markdown_filepath)
		get_markdown_content_response.Content = base64.StdEncoding.EncodeToString([]byte(get_markdown_content_response.Content))
	} else {
		return Response{
			Body: "",
		}
	}

	markdown_content, _ := json.Marshal(get_markdown_content_response)

	response := Response{
		Body: string(markdown_content),
	}
	return response
}
