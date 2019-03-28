package core

import (
	"fmt"
)

var VerificationDone = make(chan bool, 1)
var ProgramState map[string]ModeState

func Process() {
	var verificationDone = false
	defer func() {
		VerificationDone <- verificationDone
	}()

	data, e := MakeInputData()
	if e != nil {
		fmt.Println("Program input parameter are not appropriate")
	} else {
		fmt.Printf("%+v", *data)
		verificationDone = true
	}

	//privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//if err != nil {
	//	panic(err)
	//}
	//
	//msg := "hello, world"
	//hash := sha256.Sum256([]byte(msg))
	//
	//r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)
	//
	//valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	//fmt.Println("signature verified:", valid)

}

func init() {
	ProgramState = map[string]ModeState{
		"PublicKeyValidation": PublicKeyValidation,
	}
}
