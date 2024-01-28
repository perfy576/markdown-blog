package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	filepath_module "path/filepath"
	"server/config"
	"server/utils"
	"strconv"
	"strings"
)

func TraverseDirectory(path string, dir *Directory, curDepth, maxDepth int, matchPwd bool) bool {
	if curDepth > maxDepth {
		return false
	}

	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}

	files := []File{}
	directorys := []Directory{}
	authority := Authority{
		ShowPost:   true,
		ShowSubdir: true,
	}

	for _, fileInfo := range fileInfos {
		filepath := filepath_module.Join(path, fileInfo.Name())

		if fileInfo.IsDir() {
			subDir := Directory{
				Filepath:   utils.ConvertPathToLinux(filepath),
				ModifyTime: fileInfo.ModTime().String(),
				CreateTime: fileInfo.ModTime().String(),
				Directory:  []Directory{},
				File:       []File{},
			}

			directorys = append(directorys, subDir)

		} else {

			if strings.HasSuffix(filepath, ".md") {
				file := File{
					Filepath:   utils.ConvertPathToLinux(filepath),
					ModifyTime: strconv.FormatInt(fileInfo.ModTime().Unix(), 10),
					CreateTime: strconv.FormatInt(fileInfo.ModTime().Unix(), 10),
				}

				files = append(files, file)
			} else if strings.HasSuffix(filepath, "authority.json") {
				file, err := os.Open(filepath)
				if err != nil {
					fmt.Println("无法打开文件:", err)
					return false
				}
				err = json.NewDecoder(file).Decode(&authority)
				if err != nil {
					fmt.Println("反序列化失败:", err)
					return false
				}
				defer file.Close()

			}

		}
	}

	if authority.ShowPost || matchPwd {
		dir.File = files
	} else {
		dir.File = []File{}
	}

	if authority.ShowSubdir || matchPwd {
		for i := range directorys {
			addDirecotry := TraverseDirectory(directorys[i].Filepath, &directorys[i], curDepth+1, maxDepth, matchPwd)
			if addDirecotry {
				dir.Directory = append(dir.Directory, directorys[i])
			}
		}
	} else {
		dir.Directory = []Directory{}
	}

	return authority.ShowPost || authority.ShowSubdir || matchPwd
}

func removePrefix(str string, n int) string {
	ret := str
	if n >= len(str) {
		ret = str
	} else {
		ret = str[n:]
	}

	return ret
}

func ProcessFilePaths(dir *Directory, n int) {
	dir.Filepath = removePrefix(dir.Filepath, n)

	for i := range dir.File {

		dir.File[i].Filepath = removePrefix(dir.File[i].Filepath, n)
	}

	for i := range dir.Directory {
		ProcessFilePaths(&dir.Directory[i], n)
	}
}

func GetDirectoryStructure(request Request, config config.Config) Response {
	get_directory_structure_request := &GetDirectoryStructureRequest{}
	json.Unmarshal([]byte(request.Body), get_directory_structure_request)

	if strings.Contains(get_directory_structure_request.DirectoryPath, "..") {
		return Response{
			Body: "",
		}
	}

	maxDepth := 3

	root := config.ServerConf.System.Chroot
	rootPath := filepath.Join(root, get_directory_structure_request.DirectoryPath)

	rootDir := Directory{
		Filepath:  rootPath,
		Directory: []Directory{},
	}

	TraverseDirectory(rootPath, &rootDir, 0, maxDepth, get_directory_structure_request.AuthorityPassword == config.ServerConf.Authority.Password)

	rootDir.Filepath = get_directory_structure_request.DirectoryPath

	ProcessFilePaths(&rootDir, len(utils.ConverPath(root)))
	get_directory_structure_response := &GetDirectoryStructureResponse{}
	get_directory_structure_response.Root = rootDir
	get_directory_structure_response.Msg = "ok"
	markdown_content, _ := json.Marshal(get_directory_structure_response)

	response := Response{
		Body: string(markdown_content),
	}
	return response
}
