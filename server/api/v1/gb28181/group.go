package gb28181

import (
	"strconv"

	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
)

// GetGroupTree 获取分组树
func GetGroupTree(c *gin.Context) {
	var groups []system.BusinessGroup
	if err := global.GVA_DB.Order("sort asc, id asc").Find(&groups).Error; err != nil {
		response.Fail(c, response.ERROR, nil)
		return
	}

	// Build tree structure
	tree := buildGroupTree(groups, "")
	response.Success(c, gin.H{
		"list": tree,
	})
}

func buildGroupTree(groups []system.BusinessGroup, parentID string) []map[string]interface{} {
	var tree []map[string]interface{}
	for _, g := range groups {
		if g.ParentID == parentID {
			node := map[string]interface{}{
				"id":        g.ID,
				"name":      g.Name,
				"parent_id": g.ParentID,
				"sort":      g.Sort,
				"remark":    g.Remark,
			}
			children := buildGroupTree(groups, strconv.FormatUint(uint64(g.ID), 10))
			if len(children) > 0 {
				node["children"] = children
			}
			tree = append(tree, node)
		}
	}
	return tree
}

// GetGroupList 获取分组列表
func GetGroupList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	var groups []system.BusinessGroup
	var total int64
	global.GVA_DB.Model(&system.BusinessGroup{}).Count(&total)
	if err := global.GVA_DB.Offset((page - 1) * pageSize).Limit(pageSize).Order("sort asc").Find(&groups).Error; err != nil {
		response.Fail(c, response.ERROR, nil)
		return
	}
	response.Success(c, gin.H{"list": groups, "total": total})
}

// CreateGroup 创建分组
func CreateGroup(c *gin.Context) {
	var group system.BusinessGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := global.GVA_DB.Create(&group).Error; err != nil {
		response.Fail(c, response.ERROR, "create failed")
		return
	}
	response.Success(c, group)
}

// UpdateGroup 更新分组
func UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	var group system.BusinessGroup
	if err := global.GVA_DB.Where("id = ?", id).First(&group).Error; err != nil {
		response.Fail(c, response.ERROR, "group not found")
		return
	}
	if err := c.ShouldBindJSON(&group); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	global.GVA_DB.Save(&group)
	response.Success(c, nil)
}

// DeleteGroup 删除分组
func DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	global.GVA_DB.Delete(&system.BusinessGroup{}, "id = ?", id)
	response.Success(c, nil)
}
