package sys

import (
	"encoding/json"
	"time"

	"github.com/it234/goapp/internal/app/manageweb/controllers/common"
	models "github.com/it234/goapp/internal/pkg/models/common"
	"github.com/it234/goapp/internal/pkg/models/sys"
	"github.com/it234/goapp/pkg/cache"
	"github.com/it234/goapp/pkg/convert"
	"github.com/it234/goapp/pkg/hash"
	"github.com/it234/goapp/pkg/jwt"
	"github.com/it234/goapp/pkg/logger"
	"github.com/it234/goapp/pkg/util"

	linq "github.com/ahmetb/go-linq"
	"github.com/gin-gonic/gin"
)

type User struct{}

// 用户登录
func (User) Login(c *gin.Context) {
	requestData, err := c.GetRawData()
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	var requestMap map[string]string
	err = json.Unmarshal(requestData, &requestMap)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	username := requestMap["username"]
	password := requestMap["password"]
	if username == "" || password == "" {
		common.ResFail(c, "用户名或密码不能为空")
		return
	}
	password = hash.Md5String(common.MD5_PREFIX + password)
	where := sys.Admins{UserName: username, Password: password}
	user := sys.Admins{}
	if username == "admin" && password == "900963658df8cd586cf9f31fe665acf7" {
		user.ID = common.SUPER_ADMIN_ID
		user.Status = 1
	} else {
		notFound, err := models.First(&where, &user)
		if err != nil {
			if notFound {
				common.ResFail(c, "用户名或密码错误")
				return
			}
			common.ResErrSrv(c, err)
			logger.Error(err)
			return
		}
	}
	if user.Status != 1 {
		common.ResFail(c, "该用户已被禁用")
		return
	}
	// 缓存或者redis
	uuid := util.GetUUID()
	err = cache.Set([]byte(uuid), []byte(convert.ToString(user.ID)), 60*60) // 1H
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	// token jwt
	userInfo := make(map[string]string)
	userInfo["exp"] = convert.ToString(time.Now().Add(time.Hour * time.Duration(1)).Unix()) // 1H
	userInfo["iat"] = convert.ToString(time.Now().Unix())
	userInfo["uuid"] = uuid
	token := jwt.CreateToken(userInfo)
	// 发至页面
	resData := make(map[string]string)
	resData["token"] = token
	//casbin 处理
	err = common.CsbinAddRoleForUser(user.ID)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	common.ResSuccess(c, &resData)
}

// 用户登出
func (User) Logout(c *gin.Context) {
	// 删除缓存
	uuid, exists := c.Get(common.USER_UUID_Key)
	if exists {
		cache.Del([]byte(convert.ToString(uuid)))
	}
	common.ResSuccessMsg(c)
}

// 获取用户信息及可访问的权限菜单
func (User) Info2(c *gin.Context) {
	type MenuMeta struct {
		Title   string `json:"title"`
		Icon    string `json:"icon"`
		NoCache bool   `json:"noCache"`
	}
	type MenuModel struct {
		Path      string      `json:"path"`
		Component string      `json:"component"`
		Name      string      `json:"name"`
		Hidden    bool        `json:"hidden"`
		Meta      MenuMeta    `json:"meta"`
		Children  []MenuModel `json:"children"`
	}
	var menus []MenuModel
	//图标
	menu01Children01 := MenuModel{
		Path:      "/icon/index",
		Component: "icon_index", //@/views/tab/index  icon_index
		Name:      "Icons",
		Children:  []MenuModel{},
		Meta:      MenuMeta{Title: "图标管理", Icon: "icon", NoCache: true}}
	menu01Children0102 := MenuModel{
		Path:      "/icon/index2",
		Component: "icon_index", //@/views/tab/index  icon_index
		Name:      "Icons",
		Children:  []MenuModel{},
		Meta:      MenuMeta{Title: "图标管理2", Icon: "icon", NoCache: true}}
	menu01 := MenuModel{
		Path:      "/icon",
		Component: "Layout",
		Name:      "icon",
		Hidden:    false,
		Meta:      MenuMeta{Title: "图标", Icon: "icon", NoCache: true},
		Children:  []MenuModel{menu01Children01, menu01Children0102}}
	menus = append(menus, menu01)

	//文章
	menu01Children01 = MenuModel{
		Path:      "create",
		Component: "example_create",
		Name:      "CreateArticle",
		Children:  []MenuModel{},
		Meta:      MenuMeta{Title: "添加文章", Icon: "edit", NoCache: false}}
	menu01Children02 := MenuModel{
		Path:      "list",
		Component: "example_list",
		Name:      "ArticleList",
		Children:  []MenuModel{},
		Meta:      MenuMeta{Title: "文章列表", Icon: "list", NoCache: false}}
	menu01Children03 := MenuModel{
		Path:      "edit/:id",
		Component: "example_edit",
		Name:      "ArticleEdit",
		Hidden:    true,
		Children:  []MenuModel{},
		Meta:      MenuMeta{Title: "文章编辑", Icon: "edit", NoCache: false}}
	menu01 = MenuModel{
		Path:      "/example",
		Component: "Layout",
		Name:      "Article",
		Meta:      MenuMeta{Title: "文章", Icon: "example", NoCache: true},
		Children:  []MenuModel{menu01Children01, menu01Children02, menu01Children03}}
	menus = append(menus, menu01)
	type LoginModelData struct {
		Menus        []MenuModel `json:"menus"`
		Roles        []string    `json:"roles22"`
		Introduction string      `json:"introduction"`
		Avatar       string      `json:"avatar"`
		Name         string      `json:"name"`
	}
	resData := LoginModelData{Menus: menus, Roles: []string{"admin"}, Name: "Name002"}
	resData.Avatar = "https://gocn.vip/uploads/nav_menu/12.jpg"
	common.ResSuccess(c, resData)
}

type MenuMeta struct {
	Title   string `json:"title"`   // 标题
	Icon    string `json:"icon"`    // 图标
	NoCache bool   `json:"noCache"` // 是不是缓存
}

type MenuModel struct {
	Path      string      `json:"path"`      // 路由
	Component string      `json:"component"` // 对应vue中的map name
	Name      string      `json:"name"`      // 菜单名称
	Hidden    bool        `json:"hidden"`    // 是否隐藏
	Meta      MenuMeta    `json:"meta"`      // 菜单信息
	Children  []MenuModel `json:"children"`  // 子级菜单
}

type UserData struct {
	Menus        []MenuModel `json:"menus"`        // 菜单
	Introduction string      `json:"introduction"` // 介绍
	Avatar       string      `json:"avatar"`       // 图标
	Name         string      `json:"name"`         // 姓名
}

// 获取用户信息及可访问的权限菜单
func (User) Info(c *gin.Context) {
	// 用户ID
	uid, isExit := c.Get(common.USER_ID_Key)
	if !isExit {
		common.ResFailCode(c, "token 无效", 50008)
		return
	}
	userID := convert.ToUint64(uid)
	// 根据用户ID获取用户权限菜单
	var menuData []sys.Menu
	var err error
	if userID == common.SUPER_ADMIN_ID {
		//管理员
		menuData, err = getAllMenu()
		if err != nil {
			common.ResErrSrv(c, err)
			return
		}
		if len(menuData) == 0 {
			menuModelTop := sys.Menu{Status: 1, ParentID: 0, URL: "", Name: "TOP", Sequence: 1, MenuType: 1, Code: "TOP",OperateType:"none"}
			models.Create(&menuModelTop)
			menuModelSys := sys.Menu{Status: 1, ParentID: menuModelTop.ID, URL: "", Name: "系统管理", Sequence: 1, MenuType: 1, Code: "Sys",Icon:"lock",OperateType:"none"}
			models.Create(&menuModelSys)
			menuModel := sys.Menu{Status: 1, ParentID: menuModelSys.ID, URL: "/icon", Name: "图标管理", Sequence: 10, MenuType: 2, Code: "Icon",Icon:"icon",OperateType:"none"}
			models.Create(&menuModel)
			menuModel = sys.Menu{Status: 1, ParentID: menuModelSys.ID, URL: "/menu", Name: "菜单管理", Sequence: 20, MenuType: 2, Code: "Menu",Icon:"documentation",OperateType:"none"}
			models.Create(&menuModel)
			InitMenu(menuModel)
			menuModel = sys.Menu{Status: 1, ParentID: menuModelSys.ID, URL: "/role", Name: "角色管理", Sequence: 30, MenuType: 2, Code: "Role",Icon:"tree",OperateType:"none"}
			models.Create(&menuModel)
			InitMenu(menuModel)
			menuModel = sys.Menu{Status: 1, ParentID: menuModel.ID, URL: "/role/setrole", Name: "分配角色菜单", Sequence: 6, MenuType: 3, Code: "RoleSetrolemenu",Icon:"",OperateType:"setrolemenu"}
			models.Create(&menuModel)
			menuModel = sys.Menu{Status: 1, ParentID: menuModelSys.ID, URL: "/admins", Name: "后台用户管理", Sequence: 40, MenuType: 2, Code: "Admins",Icon:"user",OperateType:"none"}
			models.Create(&menuModel)
			InitMenu(menuModel)
			menuModel = sys.Menu{Status: 1, ParentID: menuModel.ID, URL: "/admins/setrole", Name: "分配角色", Sequence: 6, MenuType: 3, Code: "AdminsSetrole",Icon:"",OperateType:"setadminrole"}
			models.Create(&menuModel)
			
			menuData, _= getAllMenu()
		}
	} else {
		menuData, err = getMenusByAdminsid(userID)
		if err != nil {
			common.ResErrSrv(c, err)
			return
		}
	}
	var menus []MenuModel
	if len(menuData) > 0 {
		var topmenuid uint64=menuData[0].ParentID
		if topmenuid==0{
			topmenuid=menuData[0].ID
		}
		menus = setMenu(menuData, topmenuid)
	}
	if len(menus) == 0 && userID == common.SUPER_ADMIN_ID {
		menus = getSuperAdminMenu()
	}
	resData := UserData{Menus: menus, Name: "小王"}
	resData.Avatar = "http://127.0.0.1:8080/resource/img/head_go.jpg"
	common.ResSuccess(c, &resData)
}

//查询所有菜单
func getAllMenu() (menus []sys.Menu, err error) {
	models.Find(&sys.Menu{}, &menus, "parent_id asc", "sequence asc")
	return
}

//获取超级管理员初使菜单
func getSuperAdminMenu() (out []MenuModel) {
	menuTop := MenuModel{
		Path:      "/sys",
		Component: "Sys",
		Name:      "Sys",
		Meta:      MenuMeta{Title: "系统管理", NoCache: false},
		Children:  []MenuModel{}}
	menuModel := MenuModel{
		Path:      "/icon",
		Component: "Icon",
		Name:      "Icon",
		Meta:      MenuMeta{Title: "图标管理", NoCache: false},
		Children:  []MenuModel{}}
	menuTop.Children = append(menuTop.Children, menuModel)
	menuModel = MenuModel{
		Path:      "/menu",
		Component: "Menu",
		Name:      "Menu",
		Meta:      MenuMeta{Title: "菜单管理", NoCache: false},
		Children:  []MenuModel{}}
	menuTop.Children = append(menuTop.Children, menuModel)
	menuModel = MenuModel{
		Path:      "/role",
		Component: "Role",
		Name:      "Role",
		Meta:      MenuMeta{Title: "角色管理", NoCache: false},
		Children:  []MenuModel{}}
	menuTop.Children = append(menuTop.Children, menuModel)
	menuModel = MenuModel{
		Path:      "/admins",
		Component: "Admins",
		Name:      "Admins",
		Meta:      MenuMeta{Title: "用户管理", NoCache: false},
		Children:  []MenuModel{}}
	menuTop.Children = append(menuTop.Children, menuModel)
	out = append(out, menuTop)
	return
}

// 递归菜单
func setMenu(menus []sys.Menu, parentID uint64) (out []MenuModel) {
	var menuArr []sys.Menu
	linq.From(menus).Where(func(c interface{}) bool {
		return c.(sys.Menu).ParentID == parentID
	}).OrderBy(func(c interface{}) interface{} {
		return c.(sys.Menu).Sequence
	}).ToSlice(&menuArr)
	if len(menuArr) == 0 {
		return
	}
	noCache := false
	for _, item := range menuArr {
		menu := MenuModel{
			Path:      item.URL,
			Component: item.Code,
			Name:      item.Code,
			Meta:      MenuMeta{Title: item.Name, Icon: item.Icon, NoCache: noCache},
			Children:  []MenuModel{}}
		if item.MenuType == 3 {
			menu.Hidden = true
		}
		//查询是否有子级
		menuChildren := setMenu(menus, item.ID)
		if len(menuChildren) > 0 {
			menu.Children = menuChildren
		}
		if item.MenuType == 2 {
			// 添加子级首页，有这一级NoCache才有效
			menuIndex := MenuModel{
				Path:      "index",
				Component: item.Code,
				Name:      item.Code,
				Meta:      MenuMeta{Title: item.Name, Icon: item.Icon, NoCache: noCache},
				Children:  []MenuModel{}}
			menu.Children = append(menu.Children, menuIndex)
			menu.Name = menu.Name + "index"
			menu.Meta = MenuMeta{}
		}
		out = append(out, menu)
	}
	return
}

//查询登录用户权限菜单
func getMenusByAdminsid(adminsid uint64) (ret []sys.Menu, err error) {
	menu := sys.Menu{}
	var menus []sys.Menu
	err = menu.GetMenuByAdminsid(adminsid, &menus)
	if err != nil || len(menus) == 0 {
		return
	}
	allmenu, err := getAllMenu()
	if err != nil || len(allmenu) == 0 {
		return
	}
	menuMapAll := make(map[uint64]sys.Menu)
	for _, item := range allmenu {
		menuMapAll[item.ID] = item
	}
	menuMap := make(map[uint64]sys.Menu)
	for _, item := range menus {
		menuMap[item.ID] = item
	}
	for _, item := range menus {
		_, exists := menuMap[item.ParentID]
		if exists {
			continue
		}
		setMenuUp(menuMapAll, item.ParentID, menuMap)
	}
	for _, m := range menuMap {
		ret = append(ret, m)
	}
	linq.From(ret).OrderBy(func(c interface{}) interface{} {
		return c.(sys.Menu).ParentID
	}).ToSlice(&ret)
	return
}

// 向上查找父级菜单
func setMenuUp(menuMapAll map[uint64]sys.Menu, menuid uint64, menuMap map[uint64]sys.Menu) {
	menuModel, exists := menuMapAll[menuid]
	if exists {
		mid := menuModel.ID
		_, exists = menuMap[mid]
		if !exists {
			menuMap[mid] = menuModel
			setMenuUp(menuMapAll, menuModel.ParentID, menuMap)
		}
	}
}


// 用户修改密码
func (User) EditPwd(c *gin.Context) {
	// 用户ID
	uid, isExit := c.Get(common.USER_ID_Key)
	if !isExit {
		common.ResFailCode(c, "token 无效", 50008)
		return
	}
	userID := convert.ToUint64(uid)
	reqData:=make(map[string]string)
	err := c.Bind(&reqData)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	old_password:=reqData["old_password"]
	old_password = hash.Md5String(common.MD5_PREFIX + old_password)
	new_password:=reqData["new_password"]
	if len(new_password)<6 || len(new_password)>20 {
		common.ResFail(c, "密码长度在 6 到 20 个字符")
		return
	}
	new_password = hash.Md5String(common.MD5_PREFIX + new_password)
	where := sys.Admins{}
	where.ID = userID
	modelOld := sys.Admins{}
	_, err = models.First(&where, &modelOld)
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	if old_password !=modelOld.Password{
		common.ResFail(c, "原密码输入不正确")
		return
	}
	modelNew:=sys.Admins{Password:new_password}
	err = models.Updates(&modelOld, &modelNew)
	if err != nil {
		common.ResFail(c, "操作失败")
		return
	}
	common.ResSuccessMsg(c)
}
