package models

type ChangelogAdminReq struct {
	TableName  string
	ChangeUser string
	RowData    string
}

type TOutput struct {
	CModifier      string `gorm:"column:c_modifier;type:varchar(255);comment:修改人" json:"c_modifier"`
	CStrategyID    string `gorm:"column:c_strategy_id;type:varchar(64);not null;comment:策略ID" json:"c_strategy_id"`
	CFamilyID      string `gorm:"column:c_family_id;type:varchar(64);comment:家族ID" json:"c_family_id"`
	CParentID      string `gorm:"column:c_parent_id;type:varchar(64);comment:父ID" json:"c_parent_id"`
	CDataType      string `gorm:"column:c_data_type;type:varchar(50);comment:数据类型" json:"c_data_type"`
	CDataTypeValue string `gorm:"column:c_data_type_value;type:varchar(50);comment:数据类型值" json:"c_data_type_value"`
	CStatus        int    `gorm:"column:c_status;type:tinyint;default:1;comment:状态(1启用)" json:"c_status"`
	CMtime         string `gorm:"column:c_mtime;type:datetime;comment:修改时间" json:"c_mtime"`
	CPageUUID      string `gorm:"column:c_page_uuid;type:varchar(64);comment:页面UUID" json:"c_page_uuid"`
	CCreator       string `gorm:"column:c_creator;type:varchar(255);comment:创建人" json:"c_creator"`
	CVersion       string `gorm:"column:c_version;type:varchar(20);comment:版本号" json:"c_version"`
	CVersionType   string `gorm:"column:c_version_type;type:varchar(20);comment:版本类型" json:"c_version_type"`
	CPageClass     string `gorm:"column:c_page_class;type:varchar(255);default:'';comment:页面分类" json:"c_page_class"`
	CRelationID    string `gorm:"column:c_relation_id;type:varchar(64);default:'';comment:关联ID" json:"c_relation_id"`
	CIsPatch       bool   `gorm:"column:c_is_patch;type:tinyint(1);default:0;comment:是否补丁" json:"c_is_patch"`
	CType          string `gorm:"column:c_type;type:varchar(50);default:'default';comment:配置类型" json:"c_type"`
	CConfig        string `gorm:"column:c_config;type:longtext;comment:配置内容(JSON字符串)" json:"c_config"`
}

// HistoryOutput 历史记录表结构
type HistoryOutput struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"-"`

	CModifier      string `gorm:"column:c_modifier;type:varchar(255);comment:修改人" json:"c_modifier"`
	CStrategyID    string `gorm:"column:c_strategy_id;type:varchar(64);not null;comment:策略ID" json:"c_strategy_id"`
	CFamilyID      string `gorm:"column:c_family_id;type:varchar(64);comment:家族ID" json:"c_family_id"`
	CParentID      string `gorm:"column:c_parent_id;type:varchar(64);comment:父ID" json:"c_parent_id"`
	CDataType      string `gorm:"column:c_data_type;type:varchar(50);comment:数据类型" json:"c_data_type"`
	CDataTypeValue string `gorm:"column:c_data_type_value;type:varchar(50);comment:数据类型值" json:"c_data_type_value"`
	CStatus        int    `gorm:"column:c_status;type:tinyint;default:1;comment:状态(1启用)" json:"c_status"`
	CMtime         string `gorm:"column:c_mtime;type:datetime;comment:修改时间" json:"c_mtime"`
	CPageUUID      string `gorm:"column:c_page_uuid;type:varchar(64);comment:页面UUID" json:"c_page_uuid"`
	CCreator       string `gorm:"column:c_creator;type:varchar(255);comment:创建人" json:"c_creator"`
	CVersion       string `gorm:"column:c_version;type:varchar(20);comment:版本号" json:"c_version"`
	CVersionType   string `gorm:"column:c_version_type;type:varchar(20);comment:版本类型" json:"c_version_type"`
	CPageClass     string `gorm:"column:c_page_class;type:varchar(255);default:'';comment:页面分类" json:"c_page_class"`
	CRelationID    string `gorm:"column:c_relation_id;type:varchar(64);default:'';comment:关联ID" json:"c_relation_id"`
	CIsPatch       bool   `gorm:"column:c_is_patch;type:tinyint(1);default:0;comment:是否补丁" json:"c_is_patch"`
	CType          string `gorm:"column:c_type;type:varchar(50);default:'default';comment:配置类型" json:"c_type"`
	CConfig        string `gorm:"column:c_config;type:longtext;comment:配置内容(JSON字符串)" json:"c_config"`

	RecordStatus string `gorm:"column:record_status;type:enum('current','historical');default:'current'"`
}

type TChangeLogs struct {
	ID             uint   `gorm:"primaryKey;autoIncrement;comment:自增主键"`
	TableName      string `gorm:"column:c_table_name;type:varchar(255);not null;index;comment:表名"`
	CFamilyID      string `gorm:"column:c_family_id;type:varchar(64);comment:家族ID" json:"c_family_id"`
	CDataTypeValue string `gorm:"column:c_data_type_value;type:varchar(50);comment:数据类型值" json:"c_data_type_value"`
	ChangeUser     string `gorm:"type:varchar(255);not null;comment:变更用户"`
	ChangeLog      string `gorm:"type:longtext;comment:变更详情(JSON格式)"`
	CheckTime      int64  `gorm:"type:bigint;not null;comment:检查时间戳"`
}
