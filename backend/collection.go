package backend

import (
	_const "alicloud-notes/backend/const"
	"alicloud-notes/backend/utilis/directory"
	"fmt"
	"os"
	"path"
)

type File struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Children []*File `json:"children,omitempty"`
}

// GetCollectionDirectoryFramework 获取目录框架 所有文件
func (a *App) GetCollectionDirectoryFramework() (*directory.DirectoryStructure, error) {
	wd, _ := os.Getwd()
	ph := path.Join(wd, _const.WorkPathByLocal)
	if _, err := os.Stat(ph); err != nil {
		return nil, fmt.Errorf("工作目录不存在 %s err : %s", ph, err)
	}
	return directory.BuildDirectoryStructure(ph, true)
}

// GetCollectionDirectoryFolders 获取目录框架 所有文件夹
func (a *App) GetCollectionDirectoryFolders() (*directory.DirectoryStructure, error) {
	wd, _ := os.Getwd()
	ph := path.Join(wd, _const.WorkPathByLocal)
	if _, err := os.Stat(ph); err != nil {
		return nil, fmt.Errorf("工作目录不存在 %s err : %s", ph, err)
	}
	return directory.BuildDirectoryStructure(ph, false)
}

// GetDirectoryFrameworkByPath 获取指定目录所有文件
func (a *App) GetDirectoryFrameworkByPath(ph string) (*directory.DirectoryStructure, error) {
	if _, err := os.Stat(ph); err != nil {
		return nil, fmt.Errorf("工作目录不存在 %s err : %s", ph, err)
	}
	return directory.BuildDirectoryStructure(ph, true)
}
