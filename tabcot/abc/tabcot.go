package abc

import "container/list"

type Row struct {
	Value []interface{}
}

type Rows struct {
	list.List
}

type IDataset interface {
	Append(row Row, tags ...string)
	AppendCol(header string, data ...interface{})
	Dict() interface{}
	Export(format string) ([]byte, error)
	Extend(rows []Row, tags ...string)
	Map(func(row *Row))
	Filter(tag string) IDataset
	Col(index string) []interface{}

	Insert(index int64, row interface{}, tag ...string)
	InsertCol(index string, col, header interface{})

	Load()
	LPop()
	LPush()
	LPushCol()
	Pop()
	RemoveDuplicates()
	RPop()
	RPush()
	RPushCol()
	Sort()
	Stack()
	StackCols()
	Subset()
	Transpose()
	Height()
	Width()
	Wipe()
}

type IDataBook interface {
	AddSheet()
	Export()
	Load()
	Size()
	Wipe()
}

type IFormat struct {
}
