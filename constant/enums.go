package constant

//产品状态
type PRODUCT_STATUS uint8

const (
	PRODUCT_STATUS_USE PRODUCT_STATUS = iota + 1 // value --> 1 可用
	PRODUCT_STATUS_DIS                           // 禁用
	PRODUCT_STATUS_DEL                           // 删除
)

// 产品类型
type PRODUCT_TYPE uint8

const (
	PRODUCT_TYPE_AGR   PRODUCT_TYPE = iota + 1 // value --> 1 //农产品
	PRODUCT_TYPE_SIMF                          // 加工品
	PRODUCT_TYPE_FIN                           // 成品
	PRODUCT_TYPE_STORE                         // 上架
)

//产品种类状态
type PRODUCT_SPEC_STATUS uint8

const (
	PRODUCT_SPEC_STATUS_USE PRODUCT_SPEC_STATUS = iota + 1 // value --> 1 可用
	PRODUCT_SPEC_STATUS_DIS                                // 禁用
	PRODUCT_SPEC_STATUS_DEL                                // 删除
)

//产品使用范围
type PRODUCT_SPEC_RANGE uint8

const (
	PRODUCT_SPEC_RANGE_ALL  PRODUCT_SPEC_RANGE = iota + 1 // value --> 1 适用所有用户
	PRODUCT_SPEC_RANGE_USER                               // 适用当前用户
)

// 土地使用状态
type LAND_STATUS uint8

const (
	LAND_STATUS_USE LAND_STATUS = iota + 1 // value --> 1 可用
	LAND_STATUS_DIS                        // 禁用
	LAND_STATUS_DEL                        // 删除
)

// 资源类型
type RESOURCE_TYPE uint8

const (
	RESOURCE_TYPE_PROD  PRODUCT_TYPE = iota + 1 // value --> 1 //产品图片
	RESOURCE_TYPE_PLANT                         // 种植信息
	RESOURCE_TYPE_QUA                           // 质检文件
	RESOURCE_TYPE_ORG                           // 组织文件
)

// 资源状态
type RESOURCE_STATUS uint8

const (
	RESOURCE_STATUS_USE RESOURCE_STATUS = iota + 1 // value --> 1 可用
	RESOURCE_STATUS_DIS                            // 禁用
	RESOURCE_STATUS_DEL                            // 删除
)

// 种植类型
type PLANTLOG_TYPE uint8

const (
	PLANTLOG_TYPE_PLANT PLANTLOG_TYPE = iota + 1 // value --> 1 种植
	PLANTLOG_TYPE_IRR                            //灌溉
	PLANTLOG_TYPE_FERT                           //施肥
	PLANTLOG_TYPE_MEDI                           //用药
	PLANTLOG_TYPE_FRUIT                          //结果
	PLANTLOG_TYPE_PICK                           //采摘
	PLANTLOG_TYPE_WEE                            //除草
	PLANTLOG_TYPE_SEED                           //选种
)

// 质检状态
type QUALITY_STATUS uint8

const (
	QUALITY_STATUS_PASS QUALITY_STATUS = iota + 1
	QUALITY_STATUS_FAILED
)

// 组织状态
type ORG_STATUS uint8

const (
	ORG_STATUS_PASS ORG_STATUS = iota + 1
	ORG_STATUS_PENGING
	ORG_STATUS_FAILED
)

//种植类型状态
type PLANT_SPEC_STATUS uint8

const (
	PLANT_SPEC_STATUS_USE PLANT_SPEC_STATUS = iota + 1 // value --> 1 可用
	PLANT_SPEC_STATUS_DIS                              // 禁用
	PLANT_SPEC_STATUS_DEL                              // 删除
)

//种植使用范围
type PLANT_SPEC_RANGE uint8

const (
	PLANT_SPEC_RANGE_ALL  PLANT_SPEC_RANGE = iota + 1 // value --> 1 适用所有用户
	PLANT_SPEC_RANGE_USER                             // 适用当前用户
)

//溯源流程类型
type TRACE_PROCESS_TYPE uint8

const (
	TRACE_PROCESS_TYPE_PRO     TRACE_PROCESS_TYPE = iota + 1 // value --> 1 产品
	TRACE_PROCESS_TYPE_PLANT                                 // 种植
	TRACE_PROCESS_TYPE_LANDCO                                // 土地采集
	TRACE_PROCESS_TYPE_FACTORY                               // 工厂流程
	TRACE_PROCESS_TYPE_QA                                    // 质检流程
	TRACE_PROCESS_TYPE_PWH                                   // 成品入库

)

func (this TRACE_PROCESS_TYPE) String() string {
	switch this {
	case TRACE_PROCESS_TYPE_PRO:
		return "产品创建"
	case TRACE_PROCESS_TYPE_PLANT:
		return "种植"
	case TRACE_PROCESS_TYPE_LANDCO:
		return "土地采集"
	case TRACE_PROCESS_TYPE_FACTORY:
		return "工厂流程"
	case TRACE_PROCESS_TYPE_QA:
		return "质检流程"
	case TRACE_PROCESS_TYPE_PWH:
		return "成品入库"
	default:
		return "unknown"
	}
}

//溯源流程类型
type TRACE_CHAIN_STATUS uint8

const (
	TRACE_CHAIN_STATUS_APPLY   TRACE_CHAIN_STATUS = iota + 1 // value --> 1 申请上链
	TRACE_CHAIN_STATUS_COMMIT                                // 上链提交
	TRACE_CHAIN_STATUS_PENDING                               // 上链中
	TRACE_CHAIN_STATUS_RECEIPT                               // 上链成功
	TRACE_CHAIN_STATUS_FAILED                                // 上链失败
	TRACE_CHAIN_STATUS_NONEED                                // 无需上链
)

func (this TRACE_CHAIN_STATUS) String() string {
	switch this {
	case TRACE_CHAIN_STATUS_APPLY:
		return "申请上链"
	case TRACE_CHAIN_STATUS_COMMIT:
		return "上链提交"
	case TRACE_CHAIN_STATUS_PENDING:
		return "上链中"
	case TRACE_CHAIN_STATUS_RECEIPT:
		return "上链成功"
	case TRACE_CHAIN_STATUS_FAILED:
		return "上链失败"
	case TRACE_CHAIN_STATUS_NONEED:
		return "无需上链"
	default:
		return "unknown"
	}
}

// 加工信息使用范围
type FACTORY_PROCESS_RANGE uint8

const (
	FACTORY_PROCESS_RANGE_ALL  FACTORY_PROCESS_RANGE = iota + 1 // value --> 1 适用所有用户
	FACTORY_PROCESS_RANGE_USER                                  // 适用当前用户
)

// 加工信息状态
type FACTORY_PROCESS_STATUS uint8

const (
	FACTORY_PROCESS_STATUS_USE FACTORY_PROCESS_STATUS = iota + 1 // value --> 1 可用
	FACTORY_PROCESS_STATUS_DIS                                   // 禁用
	FACTORY_PROCESS_STATUS_DEL                                   // 删除
)
