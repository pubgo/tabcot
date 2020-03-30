package cmds

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/tabcot/version"
	"github.com/pubgo/xcmd/xcmd"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

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

// Service service name
const Service = "tabcot"

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	xenv.Cfg.Service = Service
	xenv.Cfg.Version = version.Version

	// cmd.Flags().StringVarP(p *string, name, shorthand string, value string, usage string)
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

		_wfs := xerror.PanicErr(os.OpenFile(_out, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)).(*os.File)
		defer _wfs.Close()
		_wfs.Seek(0, io.SeekEnd)

		w := csv.NewWriter(_wfs)
		w.Comma = ','
		w.UseCRLF = true

		_head := map[string]bool{}
		var dt []map[string]gjson.Result
		for _, d := range gjson.ParseBytes(xerror.PanicBytes(ioutil.ReadFile(_in))).Array() {
			_d := d.Map()
			dt = append(dt, _d)

			for k := range _d {
				_head[k] = true
			}
		}

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
