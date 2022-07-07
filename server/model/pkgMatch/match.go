// 自动生成模板Match
package pkgMatch

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Match 结构体
type Match struct {
      global.GVA_MODEL
      CategoryId  *int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类;size:19;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:比赛名称;size:50;"`
      Cover  string `json:"cover" form:"cover" gorm:"column:cover;comment:封面;size:128;"`
      Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:desc;size:255;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:作者id;size:19;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:19;"`
      Pv  *int `json:"pv" form:"pv" gorm:"column:pv;comment:page view;size:19;"`
      VoteNum  *int `json:"voteNum" form:"voteNum" gorm:"column:vote_num;comment:点赞数;size:19;"`
}


// TableName Match 表名
func (Match) TableName() string {
  return "ut_match"
}

