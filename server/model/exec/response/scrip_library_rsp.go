package response

import (
	"go-devops-mimi/server/model/exec"
)

type ScriptLibraryListRsp struct {
	Total          int64                `json:"total"`
	ScriptLibrarys []exec.ScriptLibrary `json:"script_library_s"`
}
