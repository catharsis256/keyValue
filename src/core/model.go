package core

//go:generate string -type=ModeState
type ModeState uint8

type InputData struct {
	Mode       ModeState
	PrivateKey string
	PublicKey  string
}
type RawInputData struct {
	Mode       string
	State      ModeState
	PrivateKey string
	PublicKey  string
}
