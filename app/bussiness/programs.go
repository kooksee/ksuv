package bussiness

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/ksuv/app"
	"github.com/sirupsen/logrus"
)


// 添加服务资源信息
func programs_post(d string) (app.H, error) {
	var (
		err error
	)

	app.STATUS.ErrInMaintain

	pfs := []app.ProgramsForm{}
	if err = jsoniter.Unmarshal(d, &pfs); err != nil {
		logrus.Error(err.Error())
		return app.H{"status":"ok"}, err
	}

	for i := 0; i <= len(pfs); i++ {
		pf := pfs[i]
		err = app.DB.SavePrograms(pf.Name, pf.CurrentDir, pf.Command, pf.CallBack, pf.AutoStart, pf.NumRetry, pf.Instances)
		if err != nil {
			logrus.Error(err.Error())
			return app.H{"status":"ok"}, err

		}
	}
}
