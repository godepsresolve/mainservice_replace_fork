package mainservice

import (
	"fmt"
	"github.com/godepsresolve/corelib"
	"github.com/godepsresolve/helperlib"
	"github.com/godepsresolve/plugin"
	"github.com/godepsresolve/wraplib"
)

func Run() {
	format := "%s@v%s -> %s"
	input := "HelloWorld"
	fmt.Println("wrap:" + fmt.Sprintf(format, pkg, version, wraplib.Wrap(input)))
	fmt.Println("helper:" + fmt.Sprintf(format, pkg, version, helperlib.ProvideAssistance(input)))
	fmt.Println("plugin:" + fmt.Sprintf(format, pkg, version, plugin.MakeSometingPluggable(input)))
	fmt.Println("core:" + fmt.Sprintf(format, pkg, version, corelib.Format(input)))
}
