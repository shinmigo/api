package filter

import (
	"encoding/json"
	"errors"
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/shoppb"
)

type BannerAd struct {
	*gin.Context
}

func NewBannerAd(c *gin.Context) *BannerAd {
	return &BannerAd{Context: c}
}

func (m *BannerAd) Index() (*shoppb.ListBannerAdRes, error) {
	id := m.Query("id")
	eleType := m.Query("type")
	status := m.Query("status")
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")

	var idNum uint64
	idLen := len(id)
	statusLen := len(status)
	valid := validation.Validation{}
	valid.Match(eleType, regexp.MustCompile(`^0|1|2$`)).Message("数据类型格式错误！")
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if idLen > 0 {
		valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("轮播数据不正确")
	}
	if statusLen > 0 {
		valid.Match(status, regexp.MustCompile(`^1|2$`)).Message("状态格式错误！")
	}
	if valid.HasError() {
		return nil, valid.GetError()
	}

	if idLen > 0 {
		idNum, _ = strconv.ParseUint(id, 10, 64)
	}
	pageNum, _ := strconv.ParseUint(page, 10, 64)
	pageSizeNum, _ := strconv.ParseUint(pageSize, 10, 64)
	eleTypeNum, _ := strconv.ParseUint(eleType, 10, 64)
	req := &shoppb.ListBannerAdReq{
		Page:     pageNum,
		PageSize: pageSizeNum,
		Id:       idNum,
		EleType:  uint32(eleTypeNum),
		Status:   0,
	}
	if statusLen > 0 {
		var statusNum shoppb.BannerAdStatus
		if status == "1" {
			statusNum = shoppb.BannerAdStatus_Enabled
		} else {
			statusNum = shoppb.BannerAdStatus_Disabled
		}
		req.Status = statusNum
	}

	return service.NewBannerAd(m.Context).Index(req)
}

func (m *BannerAd) Add() error {
	eleType := m.PostForm("type")
	imageUrl := m.PostForm("image_url")
	redirectUrl := m.PostForm("redirect_url")
	sort := m.PostForm("sort")
	status := m.PostForm("status")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	var statusNum shoppb.BannerAdStatus
	valid := validation.Validation{}
	valid.Required(eleType).Message("请选择数据类型")
	valid.Match(eleType, regexp.MustCompile(`^1|2$`)).Message("数据类型格式错误！")
	valid.Required(imageUrl).Message("请提交轮播图地址")
	valid.Match(imageUrl, regexp.MustCompile(`^[a-zA-z0-9,\-\.]+$`)).Message("轮播图格式错误")
	valid.Required(redirectUrl).Message("请提交轮播图跳转地址")
	valid.Required(sort).Message("请填写排序信息！")
	valid.Match(sort, regexp.MustCompile(`^[0-9]*$`)).Message("排序格式错误！")
	if valid.HasError() {
		return valid.GetError()
	}

	sortNum, _ := strconv.ParseUint(sort, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	eleTypeNum, _ := strconv.ParseUint(eleType, 10, 64)
	if status == "1" {
		statusNum = shoppb.BannerAdStatus_Enabled
	} else {
		statusNum = shoppb.BannerAdStatus_Disabled
	}
	req := &shoppb.BannerAd{
		EleType:     uint32(eleTypeNum),
		ImageUrl:    imageUrl,
		RedirectUrl: redirectUrl,
		Sort:        uint32(sortNum),
		Status:      statusNum,
		AdminId:     adminIdNum,
	}

	return service.NewBannerAd(m.Context).Add(req)
}

func (m *BannerAd) Edit() error {
	id := m.PostForm("id")
	eleType := m.PostForm("type")
	imageUrl := m.PostForm("image_url")
	redirectUrl := m.PostForm("redirect_url")
	sort := m.PostForm("sort")
	status := m.PostForm("status")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	var statusNum shoppb.BannerAdStatus
	valid := validation.Validation{}
	valid.Required(id).Message("请选择要修改的数据")
	valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("轮播数据不正确")
	valid.Required(eleType).Message("请选择数据类型")
	valid.Match(eleType, regexp.MustCompile(`^1|2$`)).Message("数据类型格式错误！")
	valid.Required(imageUrl).Message("请提交轮播图地址")
	valid.Required(redirectUrl).Message("请提交轮播图跳转地址")
	valid.Required(sort).Message("请填写排序信息！")
	valid.Match(sort, regexp.MustCompile(`^[0-9]*$`)).Message("排序格式错误！")
	if valid.HasError() {
		return valid.GetError()
	}

	idNum, _ := strconv.ParseUint(id, 10, 64)
	sortNum, _ := strconv.ParseUint(sort, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	eleTypeNum, _ := strconv.ParseUint(eleType, 10, 64)
	if status == "1" {
		statusNum = shoppb.BannerAdStatus_Enabled
	} else {
		statusNum = shoppb.BannerAdStatus_Disabled
	}
	req := &shoppb.BannerAd{
		Id:          idNum,
		EleType:     uint32(eleTypeNum),
		ImageUrl:    imageUrl,
		RedirectUrl: redirectUrl,
		Sort:        uint32(sortNum),
		Status:      statusNum,
		AdminId:     adminIdNum,
	}

	return service.NewBannerAd(m.Context).Edit(req)
}

func (m *BannerAd) EditStatus() error {
	id := m.PostForm("id")
	status := m.PostForm("status")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	var statusNum shoppb.BannerAdStatus
	valid := validation.Validation{}
	valid.Required(id).Message("请选择要修改的数据")
	if valid.HasError() {
		return valid.GetError()
	}

	idNum := make([]uint64, 0, 8)
	err := json.Unmarshal([]byte(id), &idNum)
	if err != nil {
		return errors.New("轮播数据格式错误")
	}
	if status == "1" {
		statusNum = shoppb.BannerAdStatus_Enabled
	} else {
		statusNum = shoppb.BannerAdStatus_Disabled
	}
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	req := &shoppb.EditBannerAdStatusReq{
		Id:      idNum,
		Status:  statusNum,
		AdminId: adminIdNum,
	}

	return service.NewBannerAd(m.Context).EditStatus(req)
}

func (m *BannerAd) Delete() error {
	id := m.PostForm("id")

	valid := validation.Validation{}
	valid.Required(id).Message("请选择要删除的数据！")
	if valid.HasError() {
		return valid.GetError()
	}

	idNum := make([]uint64, 0, 8)
	err := json.Unmarshal([]byte(id), &idNum)
	if err != nil {
		return errors.New("物流公司数据格式错误")
	}
	req := &shoppb.DelBannerAdReq{
		Id: idNum,
	}
	return service.NewBannerAd(m.Context).Delete(req)
}