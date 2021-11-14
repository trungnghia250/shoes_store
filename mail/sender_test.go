package mail

import (
	"testing"

	"gopkg.in/gomail.v2"

	"github.com/stretchr/testify/assert"
)

func TestWithTo(t *testing.T) {
	type args struct {
		to []string
		m  *Message
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "#1 Test join array string",
			args: args{
				to: []string{"andrew.nguyen@mobiclix.com", "ahihi@gmail.com"},
				m: &Message{
					mailMessage: gomail.NewMessage(),
				},
			},
			want: []string{"andrew.nguyen@mobiclix.com", "ahihi@gmail.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WithTo(tt.args.to)
			got(tt.args.m)
			to := tt.args.m.mailMessage.GetHeader("To")
			assert.Equal(t, tt.want, to)

		})
	}
}
