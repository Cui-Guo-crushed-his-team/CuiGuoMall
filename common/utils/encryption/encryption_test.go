package encryption

import "testing"

func TestMd5(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "user_id+user_trait",
			args: args{
				data: "123456789" + "192.168.1.1",
			},
			want: "f5d2bd3affa41216454dd6f3af45f9f9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.data); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}
