package repositories

import "github.com/jinzhu/gorm"

type CpuUsage struct {
	Connection *gorm.DB
}


func (c *CpuUsage) GetSelected() {
	
}
