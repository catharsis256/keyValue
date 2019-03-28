package core

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_validateModeState(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name      string
		args      args
		wantState ModeState
		wantErr   bool
	}{
		// TODO: Add test cases.
		{"Should return error if it was passed an empty string",
			args{""},
			ModeState(0),
			true},

		{"Should return error if it was passed an empty string",
			args{"undefined"},
			ModeState(0),
			true},

		{"Should return error if it was passed an empty string",
			args{"PublicKeyValidation"},
			PublicKeyValidation,
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotState, err := validateModeState(tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateModeState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotState != tt.wantState {
				t.Errorf("validateModeState() = %v, want %v", gotState, tt.wantState)
			}
		})
	}
}

func TestModeState_validatePath(t *testing.T) {
	type args struct {
		paths []string
	}
	var mState = PublicKeyValidation
	tests := []struct {
		name        string
		m           *ModeState
		args        args
		wantIsValid bool
		wantErr     bool
	}{
		{"Should return true if it was not passed any argument",
			&mState,
			args{},
			true,
			false},

		{"Should return true if it was passed empty array",
			&mState,
			args{[]string{}},
			true,
			false},

		{"Should return false if it was passed an array with at least one empty path",
			&mState,
			args{[]string{""}},
			false,
			true},

		{"Should return false if it was passed an array with an inappropriate format element",
			&mState,
			args{[]string{"/home/user/Documents/foo.log some fuu"}},
			false,
			true},

		{"Should return false if it was passed an array with an inappropriate format element",
			&mState,
			args{[]string{"/home/user/Documents/foo.log"}},
			false,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsValid, err := tt.m.validatePath(tt.args.paths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModeState.validatePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("ModeState.validatePath() = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}

func Test_rawInputData_validate(t *testing.T) {
	tmpfile, err := ioutil.TempFile(os.TempDir(), "example.pem")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	type fields struct {
		mode           string
		modeState      ModeState
		privateKeyPath string
		publicKeyPath  string
	}
	tests := []struct {
		name       string
		fields     fields
		wantPassed bool
		wantErr    bool
	}{
		{"Should return false and an error if validation Mode state did not pass",
			fields{mode: "", privateKeyPath: ""},
			false,
			true},

		{"Should return false and an error if it was passed a valid Mode but PK path is empty",
			fields{mode: "PublicKeyValidation", privateKeyPath: ""},
			false,
			true},

		{"Should return false and an error if it was passed a valid Mode but PK path is empty",
			fields{mode: "PublicKeyValidation", privateKeyPath: tmpfile.Name()},
			true,
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RawInputData{
				Mode:       tt.fields.mode,
				State:      tt.fields.modeState,
				PrivateKey: tt.fields.privateKeyPath,
				PublicKey:  tt.fields.publicKeyPath,
			}
			gotPassed, err := r.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("RawInputData.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPassed != tt.wantPassed {
				t.Errorf("RawInputData.validate() = %v, want %v", gotPassed, tt.wantPassed)
			}
		})
	}
}
