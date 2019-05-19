package core

type Plugin interface {
	InitPlugin()
	GetName() string
	GetDisplayName() string
	GetSettingsBase() string
	BidSharedComponent() int
	GetSharedComponent() int
}
