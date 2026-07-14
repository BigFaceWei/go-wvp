package gb28181

import (
	"strconv"

	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
)

// GetRegionTree 获取行政区域树
func GetRegionTree(c *gin.Context) {
	var regions []system.Region
	if err := global.GVA_DB.Order("level asc, sort asc, code asc").Find(&regions).Error; err != nil {
		response.Fail(c, response.ERROR, nil)
		return
	}
	tree := buildRegionTree(regions, "")
	response.Success(c, gin.H{"list": tree})
}

func buildRegionTree(regions []system.Region, parentID string) []map[string]interface{} {
	var tree []map[string]interface{}
	for _, r := range regions {
		if r.ParentID == parentID {
			node := map[string]interface{}{
				"code":      r.Code,
				"name":      r.Name,
				"parent_id": r.ParentID,
				"level":     r.Level,
				"sort":      r.Sort,
				"remark":    r.Remark,
			}
			children := buildRegionTree(regions, r.Code)
			if len(children) > 0 {
				node["children"] = children
			}
			tree = append(tree, node)
		}
	}
	return tree
}

// GetRegionList 获取行政区域列表
func GetRegionList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	parentID := c.Query("parent_id")

	var regions []system.Region
	var total int64
	db := global.GVA_DB.Model(&system.Region{})
	if parentID != "" {
		db = db.Where("parent_id = ?", parentID)
	}
	db.Count(&total)
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Order("sort asc").Find(&regions).Error; err != nil {
		response.Fail(c, response.ERROR, nil)
		return
	}
	response.Success(c, gin.H{"list": regions, "total": total})
}

// CreateRegion 创建行政区域
func CreateRegion(c *gin.Context) {
	var region system.Region
	if err := c.ShouldBindJSON(&region); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := global.GVA_DB.Create(&region).Error; err != nil {
		response.Fail(c, response.ERROR, "create failed")
		return
	}
	response.Success(c, region)
}

// UpdateRegion 更新行政区域
func UpdateRegion(c *gin.Context) {
	code := c.Param("code")
	var region system.Region
	if err := global.GVA_DB.Where("code = ?", code).First(&region).Error; err != nil {
		response.Fail(c, response.ERROR, "region not found")
		return
	}
	if err := c.ShouldBindJSON(&region); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	global.GVA_DB.Save(&region)
	response.Success(c, nil)
}

// DeleteRegion 删除行政区域
func DeleteRegion(c *gin.Context) {
	code := c.Param("code")
	global.GVA_DB.Delete(&system.Region{}, "code = ?", code)
	response.Success(c, nil)
}
