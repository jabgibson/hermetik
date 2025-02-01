package hermetik

import (
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		cipher  []byte
		subject []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "simple encryption",
			args: args{
				cipher:  []byte{120, 121, 122, 49, 50, 51},
				subject: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100},
			},
			want:    []byte{96, 94, 102, 29, 33, 211, 111, 104, 108, 29, 22},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.cipher, tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		cipher  []byte
		subject []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "simple decryption",
			args: args{
				cipher:  []byte{120, 121, 122, 49, 50, 51},
				subject: []byte{96, 94, 102, 29, 33, 211, 111, 104, 108, 29, 22},
			},
			want:    []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.cipher, tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encrypt(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "encrypt mode integrity test positive",
			args: args{
				b: 129,
			},
			want: 1,
		},
		{
			name: "encrypt mode integrity test negative",
			args: args{
				b: 1,
			},
			want: 129,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encrypt(tt.args.b); got != tt.want {
				t.Errorf("encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
