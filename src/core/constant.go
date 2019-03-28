package core

type RegExpStore string

const (
	WinPathRegex   RegExpStore = `^(?:[a-zA-Z]\:(\\|\/)|file\:\/\/|\\\\|\.(\/|\\))([^\\\/\:\*\?\<\>\"\|]+(\\|\/){0,1})+$`
	PosixPathRegex RegExpStore = `^((\.\./|[a-zA-Z0-9_/\-\.\\])*\.[a-zA-Z0-9]+)$`
)

const Windows = "windows"

const (
	_ ModeState = iota
	PublicKeyValidation
)
