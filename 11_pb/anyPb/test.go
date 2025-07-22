package main

import (
	"fmt"
	"log"
	"strconv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// TOutput 定义模拟的protobuf消息
type TOutput struct {
	FamilyId      string `protobuf:"bytes,1,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"`
	IsPatch       int32  `protobuf:"varint,2,opt,name=is_patch,json=isPatch,proto3" json:"is_patch,omitempty"`
	StrategyId    string `protobuf:"bytes,3,opt,name=strategy_id,json=strategyId,proto3" json:"strategy_id,omitempty"`
	DataTypeValue string `protobuf:"bytes,4,opt,name=data_type_value,json=dataTypeValue,proto3" json:"data_type_value,omitempty"`
	Version       string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Type          string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	ParentId      string `protobuf:"bytes,7,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Config        string `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty"`
	DataType      string `protobuf:"bytes,9,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	PageClass     string `protobuf:"bytes,10,opt,name=page_class,json=pageClass,proto3" json:"page_class,omitempty"`
	PageUuid      string `protobuf:"bytes,11,opt,name=page_uuid,json=pageUuid,proto3" json:"page_uuid,omitempty"`
	ModifyTime    string `protobuf:"bytes,12,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
}

// AnyList 定义模拟的AnyList结构
type AnyList struct {
	Items      []*anypb.Any `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	LastModify int64        `protobuf:"varint,2,opt,name=last_modify,json=lastModify,proto3" json:"last_modify,omitempty"`
}

// GetMVLCacheKey 模拟的工具函数
func GetMVLCacheKey(table, familyId, dataTypeValue string) string {
	return fmt.Sprintf("%s_%s_%s", table, familyId, dataTypeValue)
}

// OutputMapConvert 重构后的函数
func OutputMapConvert(fieldMapList []map[string]string) map[string]*AnyList {
	outputList := make(map[string]*AnyList, len(fieldMapList))

	for _, fieldMap := range fieldMapList {
		if fieldMap == nil {
			continue
		}

		// 类型转换
		patchInt, _ := strconv.Atoi(fieldMap["c_is_patch"])
		lastModify, _ := strconv.ParseInt(fieldMap["c_mtime"], 10, 64)

		// 构造TOutput对象
		output := &TOutput{
			FamilyId:      fieldMap["c_family_id"],
			IsPatch:       int32(patchInt),
			StrategyId:    fieldMap["c_strategy_id"],
			DataTypeValue: fieldMap["c_data_type_value"],
			Version:       fieldMap["c_version"],
			Type:          fieldMap["c_type"],
			ParentId:      fieldMap["c_parent_id"],
			Config:        fieldMap["c_config"],
			DataType:      fieldMap["c_data_type"],
			PageClass:     fieldMap["c_page_class"],
			PageUuid:      fieldMap["c_page_uuid"],
			ModifyTime:    fieldMap["c_mtime"],
		}

		// 转换为Any类型
		outputAny, err := anypb.New(output)
		if err != nil {
			log.Printf("anypb.New error: %v", err)
			continue
		}

		// 构造AnyList
		anyList := &AnyList{
			Items:      []*anypb.Any{outputAny},
			LastModify: lastModify,
		}

		// 生成缓存键并存储
		cacheKey := GetMVLCacheKey("t_output", output.FamilyId, output.DataTypeValue)
		outputList[cacheKey] = anyList
	}

	return outputList
}

func test() {
	// 构造测试数据
	testData := []map[string]string{
		{
			"c_family_id":       "family_001",
			"c_is_patch":        "1",
			"c_strategy_id":     "strategy_100",
			"c_data_type_value": "type_a",
			"c_version":         "v2.1.0",
			"c_type":            "output",
			"c_parent_id":       "parent_500",
			"c_config":          `{"size":10,"color":"red"}`,
			"c_data_type":       "json",
			"c_page_class":      "premium",
			"c_page_uuid":       "uuid_9k2j3h",
			"c_mtime":           "1630000000",
		},
		{
			"c_family_id":       "family_002",
			"c_is_patch":        "0",
			"c_strategy_id":     "strategy_200",
			"c_data_type_value": "type_b",
			"c_version":         "v1.5.0",
			"c_type":            "output",
			"c_parent_id":       "parent_600",
			"c_config":          `{"size":15,"color":"blue"}`,
			"c_data_type":       "json",
			"c_page_class":      "standard",
			"c_page_uuid":       "uuid_8h7j6k",
			"c_mtime":           "1631000000",
		},
	}

	// 调用函数
	result := OutputMapConvert(testData)

	// 打印结果
	fmt.Println("=== Conversion Result ===")
	for key, anyList := range result {
		fmt.Printf("Cache Key: %s\n", key)
		fmt.Printf("Last Modify: %d\n", anyList.LastModify)

		for _, item := range anyList.Items {
			var output TOutput
			if err := anypb.UnmarshalTo(item, &output, proto.UnmarshalOptions{}); err != nil {
				fmt.Printf("Unmarshal error: %v\n", err)
				continue
			}

			fmt.Printf("Output: %+v\n", output)
		}
		fmt.Println("----------------------")
	}
}
