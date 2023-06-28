package validations

// func TestIsValidPhoneNumber(t *testing.T) {
// 	type args struct {
// 		phoneNumber string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{"valid", args{phoneNumber: "559999999999"}, false},
// 		{"valid", args{phoneNumber: "aaaaaaaaa"}, true},
// 		{"valid", args{phoneNumber: "1111"}, true},
// 		{"valid", args{phoneNumber: "1111111111111111111111111111111111111111111"}, true},
// 		{"valid", args{phoneNumber: ""}, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := IsValidPhoneNumber(tt.args.phoneNumber); (err != nil) != tt.wantErr {
// 				t.Errorf("IsValidPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
