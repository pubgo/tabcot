package tabcot

import (
	"github.com/pubgo/tabcot/tabcot/abc"
)

type JsonTTT struct {
}

type Tags struct {
}

func Separator(text string, sep string) abc.Row {
	return abc.Row{}
}

func Row(value ...interface{}) abc.Row {
	return abc.Row{Value: value}
}

//detect_format
//import_set
type Dataset struct {
	Data        *abc.Rows
	Tags        map[string][]*abc.Row
	Headers     map[string]string // 名字和类型
	Title       string
	Separators  []string
	_formatters []string
}

type DataBook struct {
}

func NewDataset() *Dataset {
	return &Dataset{}
}

func init() {

}
