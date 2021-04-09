package dao

import (
	"errors"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"mn_gateway/dto"
	"mn_gateway/public"
	"time"
)

type Admin struct {
	Id       int       `json:"id" gorm:"primary_key" description:"自增ID主键"`
	UserName string    `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt     string    `json:"salt" gorm:"column:salt" description:"盐"`
	Password string    `json:"password" gorm:"column:password" description:"密码"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	IsDelete int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (*Admin) TableName() string {
	return "admin"
}

//核查密码
func (adn *Admin) AdminLoginCheckPwd(c *gin.Context, db *gorm.DB, adminLoginInput *dto.AdminLoginInput) (*Admin, error) {
	admin, err := adn.AdminFindFromDb(c, db, &Admin{UserName: adminLoginInput.UserName, IsDelete: 0})
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	passwordAddSalt := public.PasswordAddSalt(admin.Salt, adminLoginInput.Password)
	if passwordAddSalt != admin.Password {
		return admin, errors.New("密码不正确")
	}
	return admin, nil
}

//从数据库查找管理员信息
func (adn *Admin) AdminFindFromDb(c *gin.Context, db *gorm.DB, searchData *Admin) (*Admin, error) {
	out := &Admin{}
	err := db.SetCtx(public.FromGinTraceContext(c)).Where(searchData).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

//保存管理员信息到数据库
func (adn *Admin) AdminSaveToDb(c *gin.Context, db *gorm.DB) error {
	return db.SetCtx(public.FromGinTraceContext(c)).Save(adn).Error
}
