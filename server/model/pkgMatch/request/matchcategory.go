package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MatchCategorySearch struct{
    pkgMatch.MatchCategory
    request.PageInfo
}
