package filter

import (
	"github.com/shinmigo/pb/memberpb"
	"goshop/api/pkg/validation"
	"goshop/api/service"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Member struct {
	validation validation.Validation
	*gin.Context
}

func NewMember(c *gin.Context) *Member {
	return &Member{Context: c, validation: validation.Validation{}}
}

// 会员列表
func (m *Member) Index() (*memberpb.ListRes, error) {
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	m.validation.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	m.validation.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}


	pNumber, _ := strconv.ParseUint(page, 10, 32)
	pSize, _ := strconv.ParseUint(pageSize, 10, 32)
	list, err := service.NewMember(m.Context).Index(pNumber, pSize)
	if err != nil {
		return nil, err
	}
	
	return list, nil
}

// 添加会员
func (m *Member) Add() error {
	nickname := m.Query("nickname")
	mobile := m.Query("mobile")
	status := m.Query("status")
	gender := m.Query("gender")
	birthday := m.Query("birthday")
	memberLevelId := m.Query("member_level_id")
	password := m.Query("password")
	operator := m.Query("operator")
	m.validation.Required(nickname).Message("昵称不能为空！")
	m.validation.Required(mobile).Message("手机号不能为空！")
	m.validation.Required(status).Message("状态不能为空！")
	m.validation.Required(gender).Message("性别不能为空！")
	m.validation.Required(birthday).Message("生日不能为空！")
	m.validation.Required(memberLevelId).Message("会员等级不能为空！")
	m.validation.Required(password).Message("密码不能为空！")
	m.validation.Required(operator).Message("操作人不能为空！")

	if err:= service.NewMember(m.Context).Add(); err != nil {
		return err
	}

	return nil
}
