package pkgMatch

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch"
	pkgMatchReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch/request"
)

type MatchService struct {
}

// CreateMatch 创建Match记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchService *MatchService) CreateMatch(match pkgMatch.Match) (err error) {
	err = global.GVA_DB.Create(&match).Error
	return err
}

// DeleteMatch 删除Match记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchService *MatchService) DeleteMatch(match pkgMatch.Match) (err error) {
	err = global.GVA_DB.Delete(&match).Error
	return err
}

// DeleteMatchByIds 批量删除Match记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchService *MatchService) DeleteMatchByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pkgMatch.Match{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMatch 更新Match记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchService *MatchService) UpdateMatch(match pkgMatch.Match) (err error) {
	err = global.GVA_DB.Save(&match).Error
	return err
}

// GetMatch 根据id获取Match记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchService *MatchService) GetMatch(id uint) (match pkgMatch.Match, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&match).Error
	return
}

// GetMatchInfoList 分页获取Match记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchService *MatchService) GetMatchInfoList(info pkgMatchReq.MatchSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	//db := global.GVA_DB.Model(&pkgMatch.Match{})
	db := global.GVA_DB.Table("ut_match").Select("ut_match.id,ut_match.name as name,ut_match.category_id,ut_match.created_at,user_id,sort,pv,vote_num,ut_match_category.name as category_name,sys_users.username as user_name").Joins("left join ut_match_category on ut_match.category_id=ut_match_category.id").Joins("left join sys_users on ut_match.user_id=sys_users.id")
	//var matchs []pkgMatch.Match
	var matchs []pkgMatchReq.MatchSearch
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CategoryId != nil {
		db = db.Where("category_id = ?", info.CategoryId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&matchs).Error
	return err, matchs, total
}
