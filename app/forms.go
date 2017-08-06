package app

import (
	"time"
)

// programs form
type ProgramsForm struct {
	Name       string `json:"name" binding:"required" validate:"required"`    // 服务的名称,全剧唯一
	NumRetry   int    `json:"num_retry" binding:"required" validate:"isbn"`   // 重启次数
	Instances  int    `json:"instances" binding:"required" validate:"isbn"`   // 实例数量
	CurrentDir string `json:"cur_dir" binding:"required" validate:"required"` // 当前目录
	Command    string `json:"cmd" binding:"required" validate:"required"`     // 运行命令
	AutoStart  bool   `json:"auto_start" binding:"required"`                  // 自动启动
	CallBack   string `json:"callback" binding:"required"`                    // 回调接口,用于返回实时日志等,多个URL逗号隔开
}

// sessions数据
type SessionsForm struct {
	ID         string `json:"id" binding:"required"`           // ID
	MsgType    string `json:"msg_type" binding:"required"`     // session的类型
	ParentName string `json:"parent_name" binding:"required"`  // 从什么数据创造出来的
	Pid        int    `json:"parent_name" binding:"required"`  //  进程号
	CreatedAt  time.Time `json:"create_at" binding:"required"` // 创建时间
}

// status数据
// 根据进程号获取该进程的状态信息
type StatusForm struct {
	SessionID string `json:"session_id" binding:"required"`   // 状态ID
	IsAlive   bool   `json:"is_alive" binding:"required"`     // 存活状态
	CreatedAt time.Time `json:"create_at" binding:"required"` // 创建时间
}

// logs数据
// 线上环境的log会实时放到库里面的
type LogsForm struct {
	ID        int `json:"id" binding:"required"`              // 数据ID
	SessionID string `json:"session_id" binding:"required"`   // 状态ID
	CreatedAt time.Time `json:"create_at" binding:"required"` // 创建时间
	Data      string `json:"data" binding:"required"`         // session的日志输出结果
}

// scripts数据
type ScriptsForm struct {
	ID        int `json:"id" binding:"required"` // 数据ID
	Name      string `storm:"index,unique"`      // 脚本名称,全剧唯一
	Scripts   string                             // 需要运行的脚本
	CallBack  string                             // 回调接口,用于返回实时日志等,多个URL逗号隔开
	CreatedAt time.Time `storm:"index"`          // 创建时间
	Crontab   string                             // 定时执行脚本
}
