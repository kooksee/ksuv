package app

import "time"

type User struct {
	ID    int                     // primary key
	Group string `storm:"index"`  // this field will be indexed
	Email string `storm:"unique"` // this field will be indexed with a unique constraint
	Name  string                  // this field will not be indexed
	Age   int `storm:"index"`
}

type Base struct {
	Ident string `storm:"id"`
}

type User struct {
	Base      `storm:"inline"`
	Group     string `storm:"index"`
	Email     string `storm:"unique"`
	Name      string
	CreatedAt time.Time `storm:"index"`
}

type Product struct {
	Pk                  int `storm:"id,increment"`            // primary key with auto increment
	Name                string
	IntegerField        uint64 `storm:"increment"`
	IndexedIntegerField uint32 `storm:"index,increment"`
	UniqueIntegerField  int16  `storm:"unique,increment=100"` // the starting value can be set
}
////////////////////////////////////////////////////////////////////////////////

// 服务数据
type Programs struct {
	Id         int `storm:"id,increment"`    // 数据ID
	Name       string `storm:"index,unique"` // 服务的名称,全剧唯一
	NumRetry   int                           // 重启次数
	Instances  int                           // 实例数量
	CurrentDir string                        // 当前目录
	Cammand    string                        // 运行命令
	AutoStart  bool                          // 自动启动
	CallBack   string                        // 回调接口,用于返回实时日志等
}

// sessions数据
type Sessions struct {
	Id         string `storm:"id,index,unique"` // ID
	MsgType    string                           // session的类型
	ParentName string                           // 从什么数据创造出来的
	Pid        int `storm:"index"`              //  进程号
}

// status数据
// 根据进程号获取该进程的状态信息
type Status struct {
	SessionId string `storm:"id,index,unique"` // 状态ID
	IsAlive   bool                             // 存活状态
}

// logs数据
// 线上环境的log会实时放到库里面的
type Logs struct {
	SessionId string `storm:"id,index,unique"` // 状态ID
	CreateTime string `storm:"index"`           // 创建时间
}

// scripts数据
type scripts struct {
	Id         int `storm:"id,increment"`    // 数据ID
	Name       string `storm:"index,unique"` // 脚本名称,全剧唯一
	CurrentDir string                        // 当前目录
	scripts    string                        // 需要运行的脚本
	CallBack   string                        // 回调接口,用于返回实时日志等
}