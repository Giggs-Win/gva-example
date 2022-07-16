package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
)

type ExaFileUploadAndDownloadSearch struct {
	example.ExaFileUploadAndDownload
	CategoryId int
	Name       string
	Tag        string
	request.PageInfo
}
