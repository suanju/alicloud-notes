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

func (a *App) GetCollectionDirectoryFramework() (*directory.DirectoryStructure, error) {
	wd, _ := os.Getwd()
	ph := path.Join(wd, _const.WorkPathByLocal)
	if _, err := os.Stat(ph); err != nil {
		return nil, fmt.Errorf("工作目录不存在 %s err : %s", ph, err)
	}
	return directory.BuildDirectoryStructure(ph)
}
