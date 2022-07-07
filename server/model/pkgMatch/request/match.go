package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch"

)


type MatchSearch struct {
	pkgMatch.Match
	CategoryName string
	UserName     string
	request.PageInfo
}
