package validations

import "testing"

func TestIsValidPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{password: "!1ValidPassword"}, false},
		{"invalid", args{password: "invalidpassowrd"}, true},
		{"invalid", args{password: "invalid"}, true},
		{"invalid", args{password: "!1Inva"}, true},
		{"invalid", args{password: "Invalidpassowrd"}, true},
		{"invalid", args{password: "123Invalidpassowrd"}, true},
		{"invalid", args{password: "!@#Invalidpassowrd"}, true},
		{"invalid", args{password: "23!@#INVALID"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsValidPassword(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("IsValidPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
