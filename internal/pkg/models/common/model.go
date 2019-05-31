package common

// 分页条件
type PageWhereOrder struct {
	Order string
	Where string
	Value []interface{}
}