// @program:     web-template
// @file:        table.go
// @author:      edte
// @create:      2020-12-19 20:22
// @description: 
package model

var (
	tables = []struct {
		name string
		data interface{}
	}{
		{
			name: "",
			data: nil,
		},
	}
)

// initTable 初始化 table 结构
func initTable() {
	for _, v := range tables {
		if db.HasTable(v.data) {
			db.AutoMigrate(v.data)
		} else {
			db.CreateTable(v.data)
		}
	}
}
