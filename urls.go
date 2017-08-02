package main

import (
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"log"
)

func init_urls(r *gin.Engine) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", socketio_conn)
	server.On("error", socketio_error)

	r.GET("/", ping)
	r.GET("/ping", ping)

	// 服务资源操作
	// 添加多个服务资源
	r.POST("/api/programs", programs_post)

	// 修改多个服务的信息
	r.PUT("/api/programs", programs_put)

	// 修改单个服务的信息
	r.PUT("/api/programs/:name", programs_put)

	// 获取多个服务信息
	r.GET("/api/programs", programs_get)

	// 根据服务名称获取单个服务信息
	r.GET("/api/programs/:name", programs_get)

	// 删除多个服务资源
	r.DELETE("/api/programs", programs_delete)

	// 删除单个服务资源
	r.DELETE("/api/programs/:name", programs_delete)

	// 服务操作
	// 创建启动服务的session
	r.POST("/api/sessions/programs/:name", m_session_post)

	// 获得服务的运行结果
	r.GET("/api/sessions/programs", m_session_get)

	// 获得服务的运行结果
	r.GET("/api/sessions/programs/:id", session_get)

	// 暂停服务
	r.DELETE("/api/sessions/programs", m_session_delete)

	// 暂停服务
	r.DELETE("/api/sessions/programs/:id", session_delete)


	// 创建启动服务的session
	r.POST("/api/sessions/scripts/:name", m_session_post)

	// 获得服务的运行结果
	r.GET("/api/sessions/scripts", m_session_get)

	// 获得服务的运行结果
	r.GET("/api/sessions/scripts/:id", session_get)

	// 暂停服务
	r.DELETE("/api/sessions/scripts", m_session_delete)

	// 暂停服务
	r.DELETE("/api/sessions/scripts/:id", session_delete)


	// 服务状态
	// 获得服务状态
	r.GET("/api/status/programs", m_status_get)

	// 获得服务状态
	r.GET("/api/status/programs/:id", status_get)

	// 获得服务状态
	r.GET("/api/status/scripts", m_status_get)

	// 获得服务状态
	r.GET("/api/status/scripts/:id", status_get)

	// 获取该服务的日志存储状况
	r.GET("/api/status/logs", log_status)


	// 服务日志操作
	// 获取日志信息
	r.GET("/api/logs", log_get)

	// 根据服务名称获取日志信息
	r.GET("/api/logs/:name", log_get_by_name)

	// 获得具体信息
	r.GET("/api/logs/:name/:id", log_get_by_id)

	// 根据天删除之前的日志信息
	r.DELETE("/api/logs", log_delete)

	// 根据天删除之前的日志信息
	r.DELETE("/api/logs/:date", log_delete_by_date)


	// 脚本管理管理
	// 添加shell信息
	r.POST("/api/scripts", programs_stop)

	// 修改脚本信息
	r.PUT("/api/scripts", programs_stop)

	// 修改脚本信息
	r.PUT("/api/scripts/:name", programs_stop)

	// 获取脚本信息
	r.GET("/api/scripts", programs_stop)

	// 获取脚本信息
	r.GET("/api/scripts/:name", programs_stop)

	// 删除脚本
	r.DELETE("/api/scripts", programs_stop)

	// 删除脚本
	r.DELETE("/api/scripts/:name", programs_stop)

	r.GET("/socket.io/", func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	})
}
