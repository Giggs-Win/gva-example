package pkgMatch

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgMatchReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch/request"
)

type MatchCategoryService struct {
}

// CreateMatchCategory 创建MatchCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchCategoryService *MatchCategoryService) CreateMatchCategory(matchCategory pkgMatch.MatchCategory) (err error) {
	err = global.GVA_DB.Create(&matchCategory).Error
	return err
}

// DeleteMatchCategory 删除MatchCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchCategoryService *MatchCategoryService)DeleteMatchCategory(matchCategory pkgMatch.MatchCategory) (err error) {
	err = global.GVA_DB.Delete(&matchCategory).Error
	return err
}

// DeleteMatchCategoryByIds 批量删除MatchCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchCategoryService *MatchCategoryService)DeleteMatchCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pkgMatch.MatchCategory{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMatchCategory 更新MatchCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchCategoryService *MatchCategoryService)UpdateMatchCategory(matchCategory pkgMatch.MatchCategory) (err error) {
	err = global.GVA_DB.Save(&matchCategory).Error
	return err
}

// GetMatchCategory 根据id获取MatchCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchCategoryService *MatchCategoryService)GetMatchCategory(id uint) (matchCategory pkgMatch.MatchCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&matchCategory).Error
	return
}

// GetMatchCategoryInfoList 分页获取MatchCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (matchCategoryService *MatchCategoryService)GetMatchCategoryInfoList(info pkgMatchReq.MatchCategorySearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgMatch.MatchCategory{})
    var matchCategorys []pkgMatch.MatchCategory
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&matchCategorys).Error
	return  matchCategorys, total, err
}
