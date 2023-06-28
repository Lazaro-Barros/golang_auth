package validations

import "testing"

func TestIsValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"valid email", args{email: "test@test.com"}, false},
		{"empty email", args{email: ""}, true},
		{"invalid email", args{email: "@emailinvalid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsValidEmail(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("IsValidEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
