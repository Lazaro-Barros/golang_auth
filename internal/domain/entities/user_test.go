package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	type args struct {
		email           string
		inputedPassword string
		name            string
		phoneNumber     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no error",
			args: args{
				email:           "test@test.com",
				inputedPassword: "1!Abcdefs",
				name:            "name test",
				phoneNumber:     "9999999999999",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewUser(tt.args.email, tt.args.inputedPassword, tt.args.name, tt.args.phoneNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserFunctions(t *testing.T) {
	user, _ := NewUser("mocked1@user.com", "!@#123Mockedpassword", "mocker user 1", "9999999999999")
	assert.Equal(t, "mocked1@user.com", user.Email())
	assert.NotEmpty(t, user.Id())
	assert.Equal(t, "9999999999999", user.PhoneNumber())
	assert.Equal(t, "mocker user 1", user.Name())
	assert.NotEqual(t, "!@#123Mockedpassword", user.HashedPassword())
}
