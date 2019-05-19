package main

import (
	"fmt"
	"log"
	"os"
	"plugin"

	"github.com/plugins/v1/core"
)

type PluginMgr struct {
	plugins []core.Plugin
}

func NewPluginMgr() *PluginMgr {
	p := &PluginMgr{}
	return p
}

func (pmgr *PluginMgr) registerPlugin(p core.Plugin) {
	pmgr.plugins = append(pmgr.plugins, p)
}

func (pmgr *PluginMgr) GetPlugin(pluginName string) core.Plugin {
	for _, plguin := range pmgr.GetPlugins() {
		if plguin.GetName() == pluginName {
			return plguin
		}
	}
	return nil
}

func (pmgr *PluginMgr) GetPlugins() []core.Plugin {
	return pmgr.plugins
}

func main() {
	mods := []string{"./plugins/mail/mail.so", "./plugins/contact/contact.so"}
	pmgr := NewPluginMgr()

	for _, mod := range mods {
		// load module
		// 1. open the so file to load the symbols
		plug, err := plugin.Open(mod)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 2. look up a symbol (an exported function or variable)
		// in this case, variable Greeter
		symGreeter, err := plug.Lookup("Plugin")
		if err != nil {
			fmt.Println("Error:-", err)
			os.Exit(1)
		}

		// 3. Assert that loaded symbol is of a desired type
		// in this case interface type Greeter (defined above)
		greeter, ok := symGreeter.(core.Plugin)
		if !ok {
			fmt.Println("unexpected type from module symbol")
			os.Exit(1)
		}

		// 4. use the module
		pmgr.registerPlugin(greeter)
	}

	pg := pmgr.GetPlugin("mail")
	log.Println("pg:-", pg)
	log.Println(pg.GetName())

	pgs := pmgr.GetPlugin("contact")
	log.Println("pg:-", pgs)
	log.Println(pgs.GetName())
}
