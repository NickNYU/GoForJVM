package rtda

// JVM 运行时数据区的操作单位是Slot
// 除去 Long，Double 是 two slot， reference和其他基本类型都是 one slot
type Slot struct {
	num int32
	ref *Object
}
