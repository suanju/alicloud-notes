package backend

import (
	_const "alicloud-notes/backend/const"
	"alicloud-notes/backend/types"
	"alicloud-notes/backend/utilis/env"
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/zalando/go-keyring"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) Init() {
	//保持登录状态
	refreshToken, err := keyring.Get(_const.KeyringRefreshTokenKey, _const.KeyringUser)
	if err != nil {
		return
	}
	if refreshToken != "" {
		_, err = a.CreatInstance(&types.CreatInstanceReq{RefreshToken: refreshToken})
		if err != nil {
			//登录过期
			go func() {
				exitChan := make(chan bool)
				runtime.EventsOnce(a.ctx, "logout-success", func(optionalData ...interface{}) {
					fmt.Println("退出登录成功")
					exitChan <- true
				})
				for {
					select {
					case <-exitChan:
						fmt.Println("Exiting loop")
						return
					case <-time.After(1 * time.Second):
						runtime.EventsEmit(a.ctx, "login_expired")
					}
				}
			}()
			fmt.Println(err)
			return
		}
	}
	fmt.Println(PanInstance)
	fmt.Println(PanUserInfo)
	//检查环境
	env.CheckEnvironment(PanInstance, PanUserInfo)
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.Init()
}

func (a *App) Shutdown(ctx context.Context) {

}

func (a *App) DomReady(ctx context.Context) {

}
