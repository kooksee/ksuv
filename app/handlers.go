package app

import (
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"fmt"
	"net/http"
	"github.com/json-iterator/go"
)

var app = GetApp()


// 检测服务存活
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// 添加服务资源信息
func programs_post(c *gin.Context) {
	log := app.Log

	d, err := c.GetRawData()
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(string(d))

	pfs := []ProgramsForm{}
	if err = jsoniter.Unmarshal(d, &pfs); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "",
		})
	}

	for i := 0; i <= len(pfs); i++ {
		pf := pfs[i]
		err = app.DB.SavePrograms(pf.Name, pf.CurrentDir, pf.Command, pf.CallBack, pf.AutoStart, pf.NumRetry, pf.Instances)
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "",
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}

func programs_delete(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func programs_put(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func programs_get(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func programs_stop(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func m_session_post(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func m_session_get(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func m_session_ping(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func session_get(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func m_session_delete(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func status_get(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func m_status_get(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func session_ping(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func log_get_by_id(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func session_delete(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func log_get_by_name(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func log_get(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}
func log_delete_by_date(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}
func log_delete(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}
func log_status(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func socketio_conn(so socketio.Socket) {
	log := app.Log
	log.Info("on connection")
	so.Join("chat")
	so.On("chat message", func(msg string) {
		m := make(map[string]interface{})
		m["a"] = "你好"
		e := so.Emit("cn1111", m)
		//这个没有问题
		fmt.Println("\n\n")

		b := make(map[string]string)
		b["u-a"] = "中文内容" //这个不能是中文
		m["b-c"] = b
		e = so.Emit("cn2222", m)
		log.Info(e.Error())

		log.Info("emit:", so.Emit("chat message", msg))
		so.BroadcastTo("chat", "chat message", msg)
	})
	// Socket.io acknowledgement example
	// The return type may vary depending on whether you will return
	// For this example it is "string" type
	so.On("chat message with ack", func(msg string) string {
		return msg
	})
	so.On("disconnection", func() {
		log.Info("on disconnect")
	})
}

func socketio_error(so socketio.Socket, err error) {
	log := app.Log
	log.Info("error:", err)
}