package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _GoWechatMgr struct {
	*_BaseMgr
}

// GoWechatMgr open func
func GoWechatMgr(db *gorm.DB) *_GoWechatMgr {
	if db == nil {
		panic(fmt.Errorf("GoWechatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoWechatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_wechat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoWechatMgr) Debug() *_GoWechatMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoWechatMgr) GetTableName() string {
	return "go_wechat"
}

// Reset 重置gorm会话
func (obj *_GoWechatMgr) Reset() *_GoWechatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoWechatMgr) Get() (result GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoWechatMgr) Gets() (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoWechatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoWechatMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBelongWx belong_wx获取 归属微信
func (obj *_GoWechatMgr) WithBelongWx(belongWx string) Option {
	return optionFunc(func(o *options) { o.query["belong_wx"] = belongWx })
}

// WithWxid wxid获取 微信ID
func (obj *_GoWechatMgr) WithWxid(wxid string) Option {
	return optionFunc(func(o *options) { o.query["wxid"] = wxid })
}

// WithAccount account获取 微信账号
func (obj *_GoWechatMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// WithSex sex获取 1男,2女,0未知
func (obj *_GoWechatMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithAvatar avatar获取 微信头像
func (obj *_GoWechatMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithCity city获取 城市
func (obj *_GoWechatMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithCountry country获取 国家
func (obj *_GoWechatMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithLabelidList labelid_list获取 标签ID
func (obj *_GoWechatMgr) WithLabelidList(labelidList string) Option {
	return optionFunc(func(o *options) { o.query["labelid_list"] = labelidList })
}

// WithNickname nickname获取 昵称
func (obj *_GoWechatMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithProvince province获取 省份
func (obj *_GoWechatMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithRemark remark获取 备注
func (obj *_GoWechatMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTs create_ts获取
func (obj *_GoWechatMgr) WithCreateTs(createTs int64) Option {
	return optionFunc(func(o *options) { o.query["create_ts"] = createTs })
}

// WithUpdateTs update_ts获取
func (obj *_GoWechatMgr) WithUpdateTs(updateTs int64) Option {
	return optionFunc(func(o *options) { o.query["update_ts"] = updateTs })
}

// WithDeleteTs delete_ts获取
func (obj *_GoWechatMgr) WithDeleteTs(deleteTs int64) Option {
	return optionFunc(func(o *options) { o.query["delete_ts"] = deleteTs })
}

// GetByOption 功能选项模式获取
func (obj *_GoWechatMgr) GetByOption(opts ...Option) (result GoWechat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoWechatMgr) GetByOptions(opts ...Option) (results []*GoWechat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoWechatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoWechat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_GoWechatMgr) GetFromID(id int) (result GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoWechatMgr) GetBatchFromID(ids []int) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBelongWx 通过belong_wx获取内容 归属微信
func (obj *_GoWechatMgr) GetFromBelongWx(belongWx string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`belong_wx` = ?", belongWx).Find(&results).Error

	return
}

// GetBatchFromBelongWx 批量查找 归属微信
func (obj *_GoWechatMgr) GetBatchFromBelongWx(belongWxs []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`belong_wx` IN (?)", belongWxs).Find(&results).Error

	return
}

// GetFromWxid 通过wxid获取内容 微信ID
func (obj *_GoWechatMgr) GetFromWxid(wxid string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`wxid` = ?", wxid).Find(&results).Error

	return
}

// GetBatchFromWxid 批量查找 微信ID
func (obj *_GoWechatMgr) GetBatchFromWxid(wxids []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`wxid` IN (?)", wxids).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容 微信账号
func (obj *_GoWechatMgr) GetFromAccount(account string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找 微信账号
func (obj *_GoWechatMgr) GetBatchFromAccount(accounts []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`account` IN (?)", accounts).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 1男,2女,0未知
func (obj *_GoWechatMgr) GetFromSex(sex int) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找 1男,2女,0未知
func (obj *_GoWechatMgr) GetBatchFromSex(sexs []int) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 微信头像
func (obj *_GoWechatMgr) GetFromAvatar(avatar string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 微信头像
func (obj *_GoWechatMgr) GetBatchFromAvatar(avatars []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 城市
func (obj *_GoWechatMgr) GetFromCity(city string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 城市
func (obj *_GoWechatMgr) GetBatchFromCity(citys []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 国家
func (obj *_GoWechatMgr) GetFromCountry(country string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`country` = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量查找 国家
func (obj *_GoWechatMgr) GetBatchFromCountry(countrys []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`country` IN (?)", countrys).Find(&results).Error

	return
}

// GetFromLabelidList 通过labelid_list获取内容 标签ID
func (obj *_GoWechatMgr) GetFromLabelidList(labelidList string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`labelid_list` = ?", labelidList).Find(&results).Error

	return
}

// GetBatchFromLabelidList 批量查找 标签ID
func (obj *_GoWechatMgr) GetBatchFromLabelidList(labelidLists []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`labelid_list` IN (?)", labelidLists).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 昵称
func (obj *_GoWechatMgr) GetFromNickname(nickname string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找 昵称
func (obj *_GoWechatMgr) GetBatchFromNickname(nicknames []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 省份
func (obj *_GoWechatMgr) GetFromProvince(province string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 省份
func (obj *_GoWechatMgr) GetBatchFromProvince(provinces []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_GoWechatMgr) GetFromRemark(remark string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_GoWechatMgr) GetBatchFromRemark(remarks []string) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTs 通过create_ts获取内容
func (obj *_GoWechatMgr) GetFromCreateTs(createTs int64) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`create_ts` = ?", createTs).Find(&results).Error

	return
}

// GetBatchFromCreateTs 批量查找
func (obj *_GoWechatMgr) GetBatchFromCreateTs(createTss []int64) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`create_ts` IN (?)", createTss).Find(&results).Error

	return
}

// GetFromUpdateTs 通过update_ts获取内容
func (obj *_GoWechatMgr) GetFromUpdateTs(updateTs int64) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`update_ts` = ?", updateTs).Find(&results).Error

	return
}

// GetBatchFromUpdateTs 批量查找
func (obj *_GoWechatMgr) GetBatchFromUpdateTs(updateTss []int64) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`update_ts` IN (?)", updateTss).Find(&results).Error

	return
}

// GetFromDeleteTs 通过delete_ts获取内容
func (obj *_GoWechatMgr) GetFromDeleteTs(deleteTs int64) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`delete_ts` = ?", deleteTs).Find(&results).Error

	return
}

// GetBatchFromDeleteTs 批量查找
func (obj *_GoWechatMgr) GetBatchFromDeleteTs(deleteTss []int64) (results []*GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`delete_ts` IN (?)", deleteTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoWechatMgr) FetchByPrimaryKey(id int) (result GoWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoWechat{}).Where("`id` = ?", id).First(&result).Error

	return
}
