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
	programs := r.Group("/api/programs")
	{
		// 添加多个服务资源
		programs.POST("/", programs_post)

		// 修改多个服务的信息
		programs.PUT("/", programs_put)

		// 修改单个服务的信息
		programs.PUT("/:name", programs_put)

		// 获取多个服务信息
		programs.GET("/", programs_get)

		// 根据服务名称获取单个服务信息
		programs.GET("/:name", programs_get)

		// 删除多个服务资源
		programs.DELETE("/", programs_delete)

		// 删除单个服务资源
		programs.DELETE("/:name", programs_delete)

	}


	// 服务操作
	sessions := r.Group("/api/sessions")
	{
		// 创建启动服务的session
		sessions.POST("/programs/:name", m_session_post)

		// 获得服务的运行结果
		sessions.GET("/programs", m_session_get)

		// 获得服务的运行结果
		sessions.GET("/programs/:id", session_get)

		// 暂停服务
		sessions.DELETE("/programs", m_session_delete)

		// 暂停服务
		sessions.DELETE("/programs/:id", session_delete)


		// 创建启动服务的session
		sessions.POST("/scripts/:name", m_session_post)

		// 获得服务的运行结果
		sessions.GET("/scripts", m_session_get)

		// 获得服务的运行结果
		sessions.GET("/scripts/:id", session_get)

		// 暂停服务
		sessions.DELETE("/scripts", m_session_delete)

		// 暂停服务
		sessions.DELETE("/scripts/:id", session_delete)
	}


	// 服务状态
	status := r.Group("/api/status")
	{
		// 获得服务状态
		status.GET("/programs", m_status_get)

		// 获得服务状态
		status.GET("/programs/:id", status_get)

		// 获得服务状态
		status.GET("/scripts", m_status_get)

		// 获得服务状态
		status.GET("/scripts/:id", status_get)

		// 获取该服务的日志存储状况
		status.GET("/logs", log_status)
	}


	// 服务日志操作
	logs := r.Group("/api/logs")
	{
		// 获取日志信息
		logs.GET("/", log_get)

		// 根据服务名称获取日志信息
		logs.GET("/:name", log_get_by_name)

		// 获得具体信息
		logs.GET("/:name/:id", log_get_by_id)

		// 根据天删除之前的日志信息
		logs.DELETE("/", log_delete)

		// 根据天删除之前的日志信息
		logs.DELETE("/:date", log_delete_by_date)
	}



	// 脚本管理管理
	scripts := r.Group("/api/scripts")
	{
		// 添加shell信息
		scripts.POST("/", programs_stop)

		// 修改脚本信息
		scripts.PUT("/", programs_stop)

		// 修改脚本信息
		scripts.PUT("/:name", programs_stop)

		// 获取脚本信息
		scripts.GET("/", programs_stop)

		// 获取脚本信息
		scripts.GET("/:name", programs_stop)

		// 删除脚本
		scripts.DELETE("/", programs_stop)

		// 删除脚本
		scripts.DELETE("/:name", programs_stop)

	}

	r.GET("/socket.io/", func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	})
}
