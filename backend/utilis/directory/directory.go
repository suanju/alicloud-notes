package directory

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type DirectoryStructure struct {
	Genre    string                `json:"genre"`
	Name     string                `json:"name"`
	Size     int64                 `json:"size"`
	DirPath  string                `json:"dir_path"`
	Path     string                `json:"path,omitempty"`
	Children []*DirectoryStructure `json:"children,omitempty"`
}

const (
	Directory = "directory"
	File      = "file"
)

// buildDirectoryStructure 递归构建目录结构
func BuildDirectoryStructure(dirPath string, includeFiles bool) (*DirectoryStructure, error) {
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		return nil, fmt.Errorf("目录不存在 %s: %v", dirPath, err)
	}
	dir := &DirectoryStructure{
		Name:    dirInfo.Name(),
		Genre:   Directory,
		DirPath: dirPath,
		Path:    filepath.Join(dirPath, dirInfo.Name()),
		Size:    dirInfo.Size(),
	}
	// 读取目录内容
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}
	// 处理子文件和子目录
	var subFiles []*DirectoryStructure
	for _, entry := range entries {
		if entry.IsDir() {
			subDirPath := filepath.Join(dirPath, entry.Name())
			subDir, _ := BuildDirectoryStructure(subDirPath, includeFiles)
			subFiles = append(subFiles, subDir)
		} else if includeFiles {
			info, _ := entry.Info()
			subFiles = append(subFiles, &DirectoryStructure{
				Name:    entry.Name(),
				Genre:   File,
				DirPath: dirPath,
				Path:    filepath.Join(dirPath, entry.Name()),
				Size:    info.Size(),
			})
		}
	}
	// 将文件夹排在文件前面排序
	sort.Slice(subFiles, func(i, j int) bool {
		return subFiles[i].Genre == Directory && subFiles[j].Genre == File
	})

	dir.Children = subFiles
	return dir, nil
}
