package models

import "time"

type CpuUsage struct {
	Ts time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Host string `gorm:"type:varchar"`
	Usage float32 `gorm:"type:float8"`
}

func (CpuUsage) TableName() string {
	return "cpu_usage"
}
