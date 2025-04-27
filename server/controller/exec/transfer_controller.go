package exec

import (
	"fmt"
	"go-devops-mimi/server/config"
	Req "go-devops-mimi/server/model/exec/request"
	"go-devops-mimi/server/public/tools"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type TransferController struct{}

// 创建脚本库
func (m *TransferController) Add(c *gin.Context) {
	req := new(Req.TransferAddReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return TransferLogic.Add(c, req)
	})
}

// 显示脚本库列表
func (m *TransferController) List(c *gin.Context) {
	req := new(Req.TransferListReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return TransferLogic.List(c, req)
	})
}

// 显示指定脚本库
func (m *TransferController) Info(c *gin.Context) {
	req := new(Req.TransferInfoReq)
	tools.Run(c, req, func() (interface{}, interface{}) {
		return TransferLogic.Info(c, req)
	})
}

// UploadFile 上传文件处理函数
func (m *TransferController) UploadFile(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		tools.UploadObj(c, "获取文件失败："+err.Error())
		return
	}

	// 确保目录存在
	if _, err := os.Stat(config.Conf.System.UploadPath); os.IsNotExist(err) {
		err := os.MkdirAll(config.Conf.System.UploadPath, 0755)
		if err != nil {
			tools.UploadObj(c, "创建目录失败："+err.Error())
			return
		}
	}

	// 保存文件（防止重名，添加时间戳）
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	fullPath := filepath.Join(config.Conf.System.UploadPath, fileName)
	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		tools.UploadObj(c, "保存文件失败"+err.Error())
		return
	}
	tools.Success(c, config.Conf.System.UploadPath+"/"+fileName)
}
