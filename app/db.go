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
}

// sessions数据
type Sessions struct {
	Id         string `storm:"id,index,unique"` // ID
	MsgType    string                           // session的类型
	ParentName string                           // 从什么数据创造出来的
	CallBack   string                           // 回调接口,用于返回实时日志等
}

// status数据
type Status struct {
}

// logs数据
type Logs struct {
}

// scripts数据
type scripts struct {
}