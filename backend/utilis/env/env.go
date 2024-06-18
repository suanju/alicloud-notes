package env

import (
	_const "alicloud-notes/backend/const"
	"fmt"
	"os"
	"path"

	"github.com/tickstep/aliyunpan-api/aliyunpan"
	"github.com/tickstep/aliyunpan-api/aliyunpan_web"
)

func CheckEnvironment(instance *aliyunpan_web.WebPanClient, userInfo *aliyunpan.UserInfo) {
	//检测工作目录
	isWorkingDirectoryExistsByLocal()
	isWorkingDirectoryExistsByCloud(instance, userInfo)
}

func isWorkingDirectoryExistsByLocal() {
	currentDir, _ := os.Getwd()
	notesDir := path.Join(currentDir, _const.WorkPathByLocal)
	// 检查目录是否存在
	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		err = os.MkdirAll(notesDir, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录 %s err: %v", notesDir, err)
		}
	}
}

func isWorkingDirectoryExistsByCloud(instance *aliyunpan_web.WebPanClient, userInfo *aliyunpan.UserInfo) {
	workPathByCloud := path.Join("/", _const.WorkPathByCloud)
	_, apierror := instance.FileInfoByPath(userInfo.FileDriveId, workPathByCloud)
	if apierror != nil {
		fmt.Printf("检测目录 ‘%s’ err : %s \n", workPathByCloud, apierror)
		_, apierr := instance.Mkdir(userInfo.FileDriveId, "root", workPathByCloud)
		if apierr != nil {
			fmt.Printf("创建目录 ‘%s’ err : %s \n", workPathByCloud, apierror)
		}
		return
	}
}
