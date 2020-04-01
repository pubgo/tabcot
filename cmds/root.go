package cmds

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/tabcot/version"
	"github.com/pubgo/xcmd/xcmd"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"golang.org/x/crypto/ssh/terminal"
)

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}

// MapKeys 获取map的keys
func MapKeys(data interface{}) interface{} {
	vdt := reflect.ValueOf(data)
	if !vdt.IsValid() || vdt.IsNil() || vdt.Kind() != reflect.Map || vdt.Len() == 0 {
		panic("MapKeys input error")
	}

	_keys := vdt.MapKeys()
	_rst := reflect.MakeSlice(reflect.SliceOf(_keys[0].Type()), 0, len(_keys))
	_rst = reflect.Append(_rst, _keys...)
	return _rst.Interface()
}

// IsURL 检查是否为url地址
func IsURL(u string) bool {
	return strings.HasPrefix(u, "http")
}

// Service service name
const Service = "tabcot"

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	xenv.Cfg.Service = Service
	xenv.Cfg.Version = version.Version

	var expr string = ""

	cmd.PersistentFlags().StringVarP(&expr, "expr", "k", expr, "json path")
	cmd.Example = "tabcot input.json output.csv"

	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		defer xerror.RespErr(&err)

		_in := "input.json"
		_out := "output.csv"

		switch len(args) {
		case 1:
			_in = args[0]
		case 2:
			_in = args[0]
			_out = args[1]
		}

		var _inData []byte

		if !terminal.IsTerminal(0) {
			// fmt.Println(isInputFromPipe())
			_inData = xerror.PanicBytes(ioutil.ReadAll(os.Stdin))
		} else {
			if IsURL(_in) {
				resp := xerror.PanicErr(http.Get(_in)).(*http.Response)
				_inData = xerror.PanicBytes(ioutil.ReadAll(resp.Body))
			} else {
				_inData = xerror.PanicBytes(ioutil.ReadFile(_in))
			}
		}

		_head := map[string]bool{}
		var dt []map[string]gjson.Result
		var _data gjson.Result

		_data = gjson.ParseBytes(_inData)
		if expr != "" {
			_data = _data.Get(expr)
		}

		for _, d := range _data.Array() {
			_d := d.Map()
			dt = append(dt, _d)

			for k := range _d {
				_head[k] = true
			}
		}

		_wfs := xerror.PanicErr(os.OpenFile(_out, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)).(*os.File)
		defer _wfs.Close()
		_wfs.Seek(0, io.SeekEnd)

		w := csv.NewWriter(_wfs)
		w.Comma = ','
		w.UseCRLF = true

		var _head1 = MapKeys(_head).([]string)
		xerror.Panic(w.Write(_head1))

		for i := 0; i < len(dt); i++ {
			var _dt1 []string
			for j := 0; j < len(_head1); j++ {
				_dt1 = append(_dt1, dt[i][_head1[j]].String())
			}
			xerror.Panic(w.Write(_dt1))
		}

		w.Flush()
		return
	}
})
