package service

type File struct {
	Filepath   string `json:"filepath"`
	ModifyTime string `json:"modify_time"`
	CreateTime string `json:"create_time"`
}

type Directory struct {
	Filepath   string      `json:"filepath"`
	Directory  []Directory `json:"directory"`
	File       []File      `json:"file"`
	ModifyTime string      `json:"modify_time"`
	CreateTime string      `json:"create_time"`
}

type GetDirectoryStructureRequest struct {
	DirectoryPath     string `json:"directory_path"`
	AuthorityPassword string `json:"authority_password"`
}

type GetDirectoryStructureResponse struct {
	Code int64     `json:"code"`
	Msg  string    `json:"msg"`
	Root Directory `json:"root"`
}

type GetMarkdownContentRequest struct {
	Filepath          string `json:"filepath"`
	AuthorityPassword string `json:"authority_password"`
}

type GetMarkdownContentResponse struct {
	Code    int64  `json:"code"`
	Msg     string `json:"msg"`
	Content string `json:"content"`
}

type Authority struct {
	ShowPost   bool `json:"show_post"`
	ShowSubdir bool `json:"show_subdir"`
}
