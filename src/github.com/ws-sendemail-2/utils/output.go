package utils

import (
	"github.com/gin-gonic/gin/render"
	"github.com/prometheus/common/log"
)

// 对于错误结构体
//对于String的处理
type String struct {
	Format string
	Data   []interface{}
}

//对于String处理封装多的一层
func (c *Context) String(code int, format string, values ...interface{}) {
	c.Render(code, render.String{Format: format, Data: values})
}

//对于json的处理
type JSON struct {
	Data interface{}
}

//对于json的处理封装多的一层
func (c *Context) JSON(code int, obj interface{}) {
	c.Render(code, render.JSON{Data: obj})
}

//核心的一致的处理
func (c *Context) Render(code int, r render.Render) {
	c.Status(code)

	if !bodyAllowedForStatus(code) {
		r.WriteContentType(c.Writer)
		c.Writer.WriteHeaderNow()
		return
	}
	if err := r.Render(c.Writer); err != nil {
		log.Errorln(err)
	}
}
