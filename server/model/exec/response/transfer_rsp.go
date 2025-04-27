package response

import (
	"go-devops-mimi/server/model/exec"
)

type TransferRsp struct {
	Total     int64           `json:"total"`
	Transfers []exec.Transfer `json:"transfer_s"`
}
