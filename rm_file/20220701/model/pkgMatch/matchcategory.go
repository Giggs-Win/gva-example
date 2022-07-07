// 自动生成模板MatchCategory
package pkgMatch

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// MatchCategory 结构体
type MatchCategory struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:分类名称;size:50;"`
      Parent_id  *int `json:"parent_id" form:"parent_id" gorm:"column:parent_id;comment:上级分类;"`
      Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:分类描述;size:255;"`
}


// TableName MatchCategory 表名
func (MatchCategory) TableName() string {
  return "ut_match_category"
}

