package app

import "time"


// programs数据
type Programs struct {
	ID         int `storm:"id,increment"`    // 数据ID
	Name       string `storm:"index,unique"` // 服务的名称,全剧唯一
	NumRetry   int                           // 重启次数
	Instances  int                           // 实例数量
	CurrentDir string                        // 当前目录
	Command    string                        // 运行命令
	AutoStart  bool                          // 自动启动
	CallBack   string                        // 回调接口,用于返回实时日志等,多个URL逗号隔开
	CreatedAt  time.Time `storm:"index"`     // 创建时间
}

// sessions数据
type Sessions struct {
	ID         string `storm:"id,index,unique"` // ID
	MsgType    string                           // session的类型
	ParentName string                           // 从什么数据创造出来的
	Pid        int `storm:"index"`              //  进程号
	CreatedAt  time.Time `storm:"index"`        // 创建时间
}

// status数据
// 根据进程号获取该进程的状态信息
type Status struct {
	SessionID string `storm:"index"`    // 状态ID
	IsAlive   bool                      // 存活状态
	CreatedAt time.Time `storm:"index"` // 创建时间
}

// logs数据
// 线上环境的log会实时放到库里面的
type Logs struct {
	ID        int `storm:"id,increment"` // 数据ID
	SessionID string `storm:"index"`     // 状态ID
	CreatedAt time.Time `storm:"index"`  // 创建时间
	Data      string                     // session的日志输出结果
}

// scripts数据
type scripts struct {
	ID        int `storm:"id,increment"`    // 数据ID
	Name      string `storm:"index,unique"` // 脚本名称,全剧唯一
	scripts   string                        // 需要运行的脚本
	CallBack  string                        // 回调接口,用于返回实时日志等,多个URL逗号隔开
	CreatedAt time.Time `storm:"index"`     // 创建时间
	Crontab   string                        // 定时执行脚本
}