// mysql表数据初始化
// @author chen

package middleware

import (
	"errors"
	"fmt"
	"gin-admin-api/model"
	util "gin-admin-api/utils"
	"gorm.io/gorm"
	"time"
)

func parseTime(timeStr string) util.HTime {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	return util.HTime{Time: t}
}

func seedCategories(db *gorm.DB) error {
	categories := []model.BCategory{
		{ID: 1, CategoryName: "后端开发", Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 2, CategoryName: "前端开发", Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 3, CategoryName: "苹果开发", Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 4, CategoryName: "安卓开发", Sort: 4, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 6, CategoryName: "数据库", Sort: 5, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 7, CategoryName: "测试", Sort: 6, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 8, CategoryName: "运维", Sort: 7, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 9, CategoryName: "产品经理", Sort: 8, CreateTime: util.HTime{Time: time.Now()}},
	}
	for _, c := range categories {
		if err := db.FirstOrCreate(&c, "id = ?", c.ID).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedTags(db *gorm.DB) error {
	tags := []model.BTags{
		{ID: 1, TagName: "golang", Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 2, TagName: "java", Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 3, TagName: "前端", Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 4, TagName: "Mysql", Sort: 4, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 8, TagName: "Linux", Sort: 5, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 14, TagName: "ios", Sort: 6, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 20, TagName: "Android", Sort: 7, CreateTime: util.HTime{Time: time.Now()}},
	}
	for _, t := range tags {
		if err := db.FirstOrCreate(&t, "id = ?", t.ID).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedSysAdmins(db *gorm.DB) error {
	admins := []model.SysAdmin{
		{
			ID:         1,
			Username:   "admin",
			Password:   "96e79218965eb72c92a549dd5a330112",
			Nickname:   "admin",
			Status:     1,
			Icon:       "",
			Sex:        1,
			Email:      "admin001@qq.com",
			Phone:      "13889898989",
			Note:       "admin",
			CreateTime: util.HTime{Time: time.Now()},
		},
		{
			ID:         4,
			Username:   "user",
			Password:   "96e79218965eb72c92a549dd5a330112",
			Nickname:   "user",
			Status:     1,
			Icon:       "",
			Sex:        1,
			Email:      "user@qq.com",
			Phone:      "19898989891",
			Note:       "user",
			CreateTime: util.HTime{Time: time.Now()},
		},
	}
	for _, a := range admins {
		if err := db.FirstOrCreate(&a, "id = ?", a.ID).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedSysAdminRoles(db *gorm.DB) error {
	adminRoles := []model.SysAdminRole{
		{AdminId: 1, RoleId: 1},
		{AdminId: 4, RoleId: 2},
	}
	for _, ar := range adminRoles {
		var exist model.SysAdminRole
		if err := db.Where("admin_id = ? AND role_id = ?", ar.AdminId, ar.RoleId).First(&exist).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&ar).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}

func seedSysConfigs(db *gorm.DB) error {
	configs := []model.SysConfig{
		{
			ID:          1,
			Name:        "前台用户注册语",
			ConfigKey:   "registerMessage",
			ConfigValue: "Chen博客社区欢迎你的加入，尽情发表文章吧！",
			Sort:        1,
			Remark:      "前台用户注册语",
			CreateTime:  util.HTime{Time: time.Now()},
		},
		{
			ID:          2,
			Name:        "邮箱验证码主题",
			ConfigKey:   "emailSubject",
			ConfigValue: "Chen博客社区 验证码",
			Sort:        2,
			Remark:      "邮箱验证码主题",
			CreateTime:  util.HTime{Time: time.Now()},
		},
	}
	for _, c := range configs {
		if err := db.FirstOrCreate(&c, "id = ?", c.ID).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedSysMenus(db *gorm.DB) error {
	menus := []model.SysMenu{
		{ID: 1, ParentId: 0, MenuName: "系统管理", Icon: "Setting", Value: "", MenuType: 1, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 2, ParentId: 1, MenuName: "用户信息", Icon: "Avatar", Value: "", MenuType: 2, Url: "/system/sysAdmin", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 3, ParentId: 1, MenuName: "角色信息", Icon: "Operation", Value: "", MenuType: 2, Url: "/system/sysRole", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 4, ParentId: 1, MenuName: "菜单信息", Icon: "Menu", Value: "", MenuType: 2, Url: "/system/sysMenu", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 5, ParentId: 2, MenuName: "新增用户", Icon: "", Value: "system:sysAdmin:add", MenuType: 3, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 6, ParentId: 2, MenuName: "修改用户", Icon: "", Value: "system:sysAdmin:edit", MenuType: 3, Url: "", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 7, ParentId: 2, MenuName: "删除用户", Icon: "", Value: "system:sysAdmin:delete", MenuType: 3, Url: "", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 8, ParentId: 2, MenuName: "重置用户密码", Icon: "", Value: "system:sysAdmin:reset", MenuType: 3, Url: "", MenuStatus: 1, Sort: 4, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 9, ParentId: 3, MenuName: "新增角色", Icon: "", Value: "system:sysRole:add", MenuType: 3, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 10, ParentId: 3, MenuName: "修改角色", Icon: "", Value: "system:sysRole:edit", MenuType: 3, Url: "", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 11, ParentId: 3, MenuName: "删除角色", Icon: "", Value: "system:sysRole:delete", MenuType: 3, Url: "", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 12, ParentId: 3, MenuName: "分配权限", Icon: "", Value: "system:sysRole:assign", MenuType: 3, Url: "", MenuStatus: 1, Sort: 4, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 13, ParentId: 4, MenuName: "新增菜单", Icon: "", Value: "system:sysMenu:add", MenuType: 3, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 14, ParentId: 4, MenuName: "修改菜单", Icon: "", Value: "system:sysMenu:edit", MenuType: 3, Url: "", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 15, ParentId: 4, MenuName: "删除菜单", Icon: "", Value: "system:sysMenu:delete", MenuType: 3, Url: "", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 16, ParentId: 0, MenuName: "文章管理", Icon: "Connection", Value: "", MenuType: 1, Url: "", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 17, ParentId: 16, MenuName: "轮播图", Icon: "Clock", Value: "", MenuType: 2, Url: "/article/carousel", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 18, ParentId: 17, MenuName: "新增轮播图", Icon: "", Value: "article:carousel:add", MenuType: 3, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 19, ParentId: 17, MenuName: "修改轮播图", Icon: "", Value: "article:carousel:edit", MenuType: 3, Url: "", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 20, ParentId: 17, MenuName: "删除轮播图", Icon: "", Value: "article:carousel:delete", MenuType: 3, Url: "", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 21, ParentId: 16, MenuName: "文章标签", Icon: "Aim", Value: "", MenuType: 2, Url: "/article/tags", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 22, ParentId: 21, MenuName: "新增标签", Icon: "", Value: "article:tags:add", MenuType: 3, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 23, ParentId: 21, MenuName: "修改标签", Icon: "", Value: "article:tags:edit", MenuType: 3, Url: "", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 24, ParentId: 21, MenuName: "删除标签", Icon: "", Value: "article:tags:delete", MenuType: 3, Url: "", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 25, ParentId: 16, MenuName: "文章分类", Icon: "Star", Value: "", MenuType: 2, Url: "/article/category", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 26, ParentId: 25, MenuName: "新增分类", Icon: "", Value: "article:category:add", MenuType: 3, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 27, ParentId: 25, MenuName: "修改分类", Icon: "", Value: "article:category:edit", MenuType: 3, Url: "", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 28, ParentId: 25, MenuName: "删除分类", Icon: "", Value: "article:category:delete", MenuType: 3, Url: "", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 30, ParentId: 0, MenuName: "前台用户", Icon: "User", Value: "", MenuType: 1, Url: "", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 31, ParentId: 30, MenuName: "用户信息", Icon: "UserFilled", Value: "", MenuType: 2, Url: "/member/user", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 32, ParentId: 31, MenuName: "发送消息", Icon: "", Value: "member:user:message", MenuType: 3, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 33, ParentId: 16, MenuName: "文章信息", Icon: "Connection", Value: "", MenuType: 2, Url: "/article/article", MenuStatus: 1, Sort: 4, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 39, ParentId: 33, MenuName: "删除文章", Icon: "", Value: "article:article:delete", MenuType: 3, Url: "", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 40, ParentId: 33, MenuName: "查看评论", Icon: "", Value: "article:article:query", MenuType: 3, Url: "", MenuStatus: 1, Sort: 0, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 41, ParentId: 16, MenuName: "文章评论", Icon: "Notification", Value: "", MenuType: 2, Url: "/article/comment", MenuStatus: 1, Sort: 5, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 42, ParentId: 41, MenuName: "评论删除", Icon: "", Value: "article:comment:delete", MenuType: 3, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 48, ParentId: 1, MenuName: "参数信息", Icon: "Discount", Value: "", MenuType: 2, Url: "/system/sysConfig", MenuStatus: 1, Sort: 4, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 49, ParentId: 48, MenuName: "新增参数", Icon: "", Value: "system:sysConfig:add", MenuType: 3, Url: "", MenuStatus: 1, Sort: 1, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 50, ParentId: 48, MenuName: "修改参数", Icon: "", Value: "system:sysConfig:edit", MenuType: 3, Url: "", MenuStatus: 1, Sort: 2, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 51, ParentId: 48, MenuName: "删除配置", Icon: "", Value: "system:sysConfig:delete", MenuType: 3, Url: "", MenuStatus: 1, Sort: 3, CreateTime: util.HTime{Time: time.Now()}},
		{ID: 52, ParentId: 48, MenuName: "刷新配置", Icon: "", Value: "system:sysConfig:refresh", MenuType: 3, Url: "", MenuStatus: 1, Sort: 4, CreateTime: util.HTime{Time: time.Now()}},
	}

	for _, m := range menus {
		if err := db.FirstOrCreate(&m, "id = ?", m.ID).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedSysRoles(db *gorm.DB) error {
	roles := []model.SysRole{
		{
			ID:          1,
			RoleName:    "超级管理员",
			RoleKey:     "admin",
			Status:      1,
			Description: "全部权限",
			CreateTime:  parseTime("2024-03-22 19:21:30"),
		},
		{
			ID:          2,
			RoleName:    "管理员",
			RoleKey:     "user",
			Status:      1,
			Description: "部分权限",
			CreateTime:  parseTime("2024-03-22 19:21:44"),
		},
	}
	for _, r := range roles {
		if err := db.FirstOrCreate(&r, "id = ?", r.ID).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedSysRoleMenus(db *gorm.DB) error {
	roleMenus := []model.SysRoleMenu{
		// 角色ID=1的权限
		{RoleId: 1, MenuId: 1},
		{RoleId: 1, MenuId: 2},
		{RoleId: 1, MenuId: 5},
		{RoleId: 1, MenuId: 6},
		{RoleId: 1, MenuId: 7},
		{RoleId: 1, MenuId: 8},
		{RoleId: 1, MenuId: 3},
		{RoleId: 1, MenuId: 9},
		{RoleId: 1, MenuId: 10},
		{RoleId: 1, MenuId: 11},
		{RoleId: 1, MenuId: 12},
		{RoleId: 1, MenuId: 4},
		{RoleId: 1, MenuId: 13},
		{RoleId: 1, MenuId: 14},
		{RoleId: 1, MenuId: 15},
		{RoleId: 1, MenuId: 48},
		{RoleId: 1, MenuId: 49},
		{RoleId: 1, MenuId: 50},
		{RoleId: 1, MenuId: 51},
		{RoleId: 1, MenuId: 52},
		{RoleId: 1, MenuId: 16},
		{RoleId: 1, MenuId: 17},
		{RoleId: 1, MenuId: 18},
		{RoleId: 1, MenuId: 19},
		{RoleId: 1, MenuId: 20},
		{RoleId: 1, MenuId: 21},
		{RoleId: 1, MenuId: 22},
		{RoleId: 1, MenuId: 23},
		{RoleId: 1, MenuId: 24},
		{RoleId: 1, MenuId: 25},
		{RoleId: 1, MenuId: 26},
		{RoleId: 1, MenuId: 27},
		{RoleId: 1, MenuId: 28},
		{RoleId: 1, MenuId: 33},
		{RoleId: 1, MenuId: 40},
		{RoleId: 1, MenuId: 39},
		{RoleId: 1, MenuId: 41},
		{RoleId: 1, MenuId: 42},
		{RoleId: 1, MenuId: 30},
		{RoleId: 1, MenuId: 31},
		{RoleId: 1, MenuId: 32},

		// 角色ID=2的权限
		{RoleId: 2, MenuId: 8},
		{RoleId: 2, MenuId: 10},
		{RoleId: 2, MenuId: 4},
		{RoleId: 2, MenuId: 13},
		{RoleId: 2, MenuId: 14},
		{RoleId: 2, MenuId: 15},
		{RoleId: 2, MenuId: 1},
		{RoleId: 2, MenuId: 2},
		{RoleId: 2, MenuId: 3},
	}

	for _, rm := range roleMenus {
		var exist model.SysRoleMenu
		// 使用复合主键查询是否存在
		if err := db.Where("role_id = ? AND menu_id = ?", rm.RoleId, rm.MenuId).First(&exist).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 不存在则创建
				if err := db.Create(&rm).Error; err != nil {
					return fmt.Errorf("failed to create role menu: %v", err)
				}
			} else {
				return fmt.Errorf("failed to query role menu: %v", err)
			}
		}
		// 已存在则跳过
	}
	return nil
}

// SeedData 更新SeedData函数
func SeedData(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := seedCategories(tx); err != nil {
			return err
		}
		if err := seedTags(tx); err != nil {
			return err
		}
		if err := seedSysAdmins(tx); err != nil {
			return err
		}
		if err := seedSysAdminRoles(tx); err != nil {
			return err
		}
		if err := seedSysConfigs(tx); err != nil {
			return err
		}
		if err := seedSysMenus(tx); err != nil {
			return err
		}
		if err := seedSysRoles(tx); err != nil {
			return err
		}
		if err := seedSysRoleMenus(tx); err != nil {
			return err
		}
		return nil
	})
}
