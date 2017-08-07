package bussiness

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/ksuv/app"
	"github.com/sirupsen/logrus"
	"fmt"
)


/// 添加服务资源信息
func Programs_post(d []byte) (app.Returns, error) {
	var (
		err error
	)

	app.STATUS.ErrInMaintain

	pfs := []app.ProgramsForm{}
	if err = jsoniter.Unmarshal(d, &pfs); err != nil {
		logrus.Error(err.Error())
		return app.Returns{
			Code:app.STATUS.ErrInMaintain,
		}, err
	}

	for i := 0; i <= len(pfs); i++ {
		pf := pfs[i]
		err = app.DB.SavePrograms(pf.Name, pf.CurrentDir, pf.Command, pf.CallBack, pf.AutoStart, pf.NumRetry, pf.Instances)
		if err != nil {
			logrus.Error(err.Error())
			return app.Returns{
				Code:app.STATUS.ErrInMaintain,
			}, err

		}
	}
	return app.Returns{
		Code:app.STATUS.Ok,
	}, err
}

func programs_delete(d []byte) {
	fmt.Println(d)
}

func programs_put(d []byte) {
	fmt.Println(d)
}

func programs_get(d []byte) {
	fmt.Println(d)
}

func programs_stop(d []byte) {
	fmt.Println(d)
}