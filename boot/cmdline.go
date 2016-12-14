package boot

import (
	"bytes"
	"flag"
	"html/template"
	"log"
	"strings"

	. "ActivedRouter/global"
)

///解析命令行参数
func parseCmdline() {
	runmode := flag.String("runmode", "", "runmode must be server or client or reserveproxy")
	flag.Parse()
	//如果没设置运行模式
	if *runmode == "" {
		//log.Println(UsageTemplate)
	} else {
		if strings.ToLower(*runmode) != ServerMode &&
			strings.ToLower(*runmode) != ClientMode &&
			strings.ToLower(*runmode) != ReserveProxyMode &&
			strings.ToLower(*runmode) != MixMode {
			t, _ := template.New("info").Parse(UsageRunmodeTemplate)
			buffer := &bytes.Buffer{}
			t.Execute(buffer, struct{ Msg string }{Msg: "runmode参数错误,参考 ActiveRouter --runmode=Server或Client或Reserveproxy 或Mix"})
			log.Println(string(buffer.Bytes()))
		} else {
			//set run mode
			RunMode = *runmode
		}
	}
	//EXIT:
	//	os.Exit(0)
}
