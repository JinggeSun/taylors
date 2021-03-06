// 自动生成模板Stock
package model

type StockMonitor struct {
	Id          int64   `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Code        string  `gorm:"column:code" json:"code"`                 //编码
	MonitorHigh float64 `gorm:"column:monitor_high" json:"monitor_high"` //高位预警
	MonitorLow  float64 `gorm:"column:monitor_low" json:"monitor_low"`   //低位预警
	UserId      uint    `gorm:"column:user_id" json:"user_id"`           //所属用户
	IsDay       bool    `gorm:"column:is_day" json:"is_day"`             //日监控
	Day         int64   `gorm:"column:day" json:"day"`                   //日期（日监控）
	DelStatus   int64   `gorm:"column:del_status" json:"del_status"`     //删除状态
	CreateTime  int64   `gorm:"column:create_time" json:"create_time"`   //数据创建时间
}

func (StockMonitor) TableName() string {
	return "stock_monitor"
}
