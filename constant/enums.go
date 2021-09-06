package constant

type PRODUCT_STATUS uint8

const (
	PRODUCT_STATUS_USE PRODUCT_STATUS = iota + 1 // value --> 1 可用
	PRODUCT_STATUS_DIS                           // 禁用
	PRODUCT_STATUS_DEL                           // 删除
)
