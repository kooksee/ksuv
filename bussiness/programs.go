package bussiness

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/ksuv/kapp"
	"github.com/sirupsen/logrus"
	"fmt"
)

var app = kapp.GetApp()

/// 添加服务资源信息
func Programs_post(d []byte) (kapp.Returns, error) {
	var (
		err error
	)

	pfs := []kapp.ProgramsForm{}
	if err = jsoniter.Unmarshal(d, &pfs); err != nil {
		logrus.Error(err.Error())
		return kapp.Returns{
			Code:kapp.STATUS.ErrInMaintain,
		}, err
	}

	for i := 0; i <= len(pfs)-1; i++ {
		pf := pfs[i]
		fmt.Println(pf)
		err = app.DB.SavePrograms(pf.Name, pf.CurrentDir, pf.Command, pf.CallBack, pf.AutoStart, pf.NumRetry, pf.Instances)
		if err != nil {
			logrus.Error(err.Error())
			return kapp.Returns{
				Code:kapp.STATUS.ErrInMaintain,
			}, err

		}
	}
	return kapp.Returns{
		Code:kapp.STATUS.Ok,
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