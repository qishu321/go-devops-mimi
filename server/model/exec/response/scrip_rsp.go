package response

import (
	"go-devops-mimi/server/model/exec"
)

type ScriptListRsp struct {
	Total   int64         `json:"total"`
	Scripts []exec.Script `json:"script_s"`
}
