package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ws-sendemail-2/defs"
		"net/http"
	"time"
	"log"
)

/**
发送邮件的Handler
@author Ares Peng(pengpp@wondersahre.cn)
@desc 接收gin的上下文，返回具体操作值
@Method POST
@Param String from		 发件人
@Param String to		 收件人
@Param String cc		 抄送人
@Param String subject	 发送主题
@Param String Body		 发送邮件内容
@Param String isHtml	 是否html格式
@Param String hasAttach  是否含有附件
@Param String attachment 附件的具体地址
@return Json
@HttpStatus
	200	程序访问成功
	201 访问成功并处理成功
	400 bad request
	500 内部服务器错误
**/
func SendEmailsV1(c *gin.Context) {
	var Email = defs.Email{}
	if err := c.ShouldBind(&Email); err != nil {
		log.Println(c.ShouldBind(&Email))
		c.JSON(http.StatusBadRequest, gin.H{
			"code":      http.StatusBadRequest,
			"msg":       "error",
			"timestamp": time.Now(),
		})
		return
	}
	if _, err := CheckEmail(&Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":      http.StatusBadRequest,
			"msg":       err.Error(),
			"timestamp": time.Now(),
		})
		return
	}
	if sid, err := SetMail(&Email); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":      http.StatusOK,
			"msg":       "success",
			"sid":       sid,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":      http.StatusBadRequest,
			"msg":       "db error",
		})
		return
	}
}

/**
查看邮件列表Handler
@author Ares Peng(pengpp@wondersahre.cn)
@desc 接收gin的上下文，返回具体操作值
@Method POST
@Param String limit		 每页显示多少数据
@Param String page		 当前页
@Param String status	 查询状态
@return Json
@HttpStatus
	200	程序访问成功
	400 bad request
	500 内部服务器错误
**/
func EmailsListV1(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"msg":  "Welcome server 01",
		},
	)
}

/**
查看具体某个发件人邮件列表的Handler
@author Ares Peng(pengpp@wondersahre.cn)
@desc 接收gin的上下文，返回具体操作值
@Method POST
@Param String email		 发件人邮箱
@Param String limit		 每页显示多少数据
@Param String page		 当前页
@Param String status	 查询状态
@return Json
@HttpStatus
	200	程序访问成功
	400 bad request
	500 内部服务器错误
**/
func EmailInfoV1(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"msg":  "Welcome server 01",
		},
	)
}

/**
重新发送邮件 的Handler
@author Ares Peng(pengpp@wondersahre.cn)
@desc 接收gin的上下文，返回具体操作值
@Method POST
@Param String email		 发件人邮箱
@return Json
@HttpStatus
	200	程序访问成功
	201	程序访问成功并发送成功
	400 bad request
	403 校验失败
	500 内部服务器错误
**/
func ResendEmailsV1(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"msg":  "Welcome server 01",
		},
	)
}
