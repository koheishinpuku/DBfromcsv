package Models

type Table struct {
	Table_name    string `json:table_name`
	Table_name_ja string `gorm:"column:table_name_ja"`
	Update_info   string `json:update_info`
	Describe      string `json:describe`
}

func (m *Table) TableName() string {
	return "search_table"
}
