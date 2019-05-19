package main

const (
	PLUGIN_NAME         = "contact"
	PLUGIN_DISPLAYNAME  = "Contact"
	PLUGIN_SETTINGSBASE = "contact"
)

type plugin string

func (p *plugin) InitPlugin() {
}

func (p *plugin) GetName() string {
	return PLUGIN_NAME
}

func (p *plugin) GetDisplayName() string {
	return PLUGIN_DISPLAYNAME
}

func (p *plugin) GetSettingsBase() string {
	return PLUGIN_SETTINGSBASE
}

func (p *plugin) BidSharedComponent() int {
	return -1
}

func (p *plugin) GetSharedComponent() int {
	return -1
}

var Plugin plugin
