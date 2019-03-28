package core

import (
	"errors"
	"flag"
	"fmt"
)

func MakeInputData() (iData *InputData, err error) {
	//var iData InputData
	arguments := makeRawArgumentString()
	fmt.Printf("%+v\n", *arguments)

	if passed, err := arguments.validate(); !passed {
		fmt.Println("raw input data validation did not passed")
		return nil, err
	}

	arguments.constructInputData()

	return nil, errors.New("undefined")
}

func (r *RawInputData) constructInputData() (iData *InputData, err error) {
	iData = &InputData{}
	return nil, nil
}

//parse Argument String
func makeRawArgumentString() *RawInputData {
	rInputData := &RawInputData{}

	flag.StringVar(&rInputData.Mode, "Mode", "PublicKeyValidation", "What would you like to verify? [PrivateToPublic,...]")
	flag.StringVar(&rInputData.PrivateKey, "PrivateKey", "", "Path to private key")
	flag.StringVar(&rInputData.PublicKey, "PublicKey", "", "Path to public key")

	flag.Parse()

	return rInputData
}

func (r *RawInputData) validate() (passed bool, err error) {
	state, err := validateModeState(r.Mode)
	if err != nil {
		return
	}

	r.State = state

	switch state {
	case PublicKeyValidation:
		passed, err = state.validatePath(r.PrivateKey)
	default:
		err = errors.New("unknown program Mode type")
	}

	return passed, err
}

func validateModeState(mode string) (state ModeState, err error) {
	if len(mode) == 0 {
		fmt.Println("Program Mode string is empty. It has to be one one of the ", ProgramState)
		return state, errors.New("program Mode string is empty")
	}

	if pState, ok := ProgramState[mode]; ok {
		return pState, nil
	}

	return state, errors.New("program Mode has not been found in the available set")
}

func (m *ModeState) validatePath(paths ...string) (isValid bool, err error) {
	// this is because it has to catch errors of regexp compilation
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New(fmt.Sprintf("%v", rec))
			return
		}
	}()

	lpType := getRegExpByOS()

	for _, path := range paths {

		if len(path) == 0 {
			fmt.Println("Program Mode string is empty. It has to be one one of the ", ProgramState)
			return false, errors.New("program Mode string is empty")
		}

		if isPathValid := lpType.RegExpValidation(path); !isPathValid {
			fmt.Println("RegExpValidation did not pass for ", path)
			return false, errors.New("file path format is not appropriate")
		}

		if isFileValid := FileValidation(path); !isFileValid {
			fmt.Println("FileValidation did not pass for ", path)
			return false, errors.New("file is broken or invalid")
		}

	}

	return true, nil
}

func getRegExpByOS() RegExpStore {
	if IsItRunInPosix() {
		return WinPathRegex
	}
	return PosixPathRegex
}
