package lighthouse

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestReport(t *testing.T) {
	Convey("Unmarshal report", t, func() {
		b, err := ioutil.ReadFile("report.json")
		So(err, ShouldBeNil)
		var report *Report
		err = json.Unmarshal(b, &report)
		So(err, ShouldBeNil)
		So(report, ShouldNotBeNil)
	})
}

func TestDevtoolsLog(t *testing.T) {
	Convey("Unmarshal devtools log", t, func() {
		b, err := ioutil.ReadFile("devtoolslog.json")
		So(err, ShouldBeNil)
		var log []LogItem
		err = json.Unmarshal(b, &log)
		So(err, ShouldBeNil)
		So(log, ShouldNotBeNil)
	})
}
