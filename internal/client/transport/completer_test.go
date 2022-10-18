package transport

import (
	"fmt"
	"testing"

	"github.com/c-bata/go-prompt"
	"github.com/stretchr/testify/assert"
)

func TestCommander_Completer(t *testing.T) {
	commander := NewCommander(nil, nil, nil)

	tests := []struct {
		name string
		d    prompt.Document
		want []prompt.Suggest
	}{
		{
			name: "registration",
			d: prompt.Document{
				Text: fmt.Sprintf("%s u1 u1", registration),
			},
			want: []prompt.Suggest{
				{Text: registration, Description: registrationDesc},
			},
		},
		{
			name: "authentication",
			d: prompt.Document{
				Text: fmt.Sprintf("%s u1 u1", authentication),
			},
			want: []prompt.Suggest{
				{Text: authentication, Description: authenticationDesc},
			},
		},
		{
			name: "create_lp",
			d: prompt.Document{
				Text: fmt.Sprintf("%s %s login password", create, loginPassDataType),
			},
			want: []prompt.Suggest{
				{Text: fmt.Sprint(create, " ", loginPassDataType), Description: getDesc(create, loginPassDataTypeDesc)},
			},
		},
		{
			name: "update_text",
			d: prompt.Document{
				Text: fmt.Sprintf("%s %s text", update, textDataType),
			},
			want: []prompt.Suggest{
				{Text: fmt.Sprint(update, " ", textDataType), Description: getDesc(update, textDataTypeDesc)},
			},
		},
		{
			name: "delete_binary",
			d: prompt.Document{
				Text: fmt.Sprintf("%s %s binary", del, binaryDataType),
			},
			want: []prompt.Suggest{
				{Text: fmt.Sprint(del, " ", binaryDataType), Description: getDesc(del, binaryDataTypeDesc)},
			},
		},

		{
			name: "read_bankCard",
			d: prompt.Document{
				Text: fmt.Sprintf("%s %s 6666", read, bankCardDataType),
			},
			want: []prompt.Suggest{
				{Text: fmt.Sprint(read, " ", bankCardDataType), Description: getDesc(read, bankCardDataTypeDesc)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := commander.Completer(tt.d)
			assert.Equal(t, tt.want, resp)
		})
	}
}

func Test_getDesc(t *testing.T) {
	tests := []struct {
		name         string
		command      string
		dataTypeDesc string
		want         string
	}{
		{
			name:         "read",
			command:      read,
			dataTypeDesc: readDesc,
			want:         readDesc,
		},
		{
			name:         "update",
			command:      update,
			dataTypeDesc: updateDesc,
			want:         fmt.Sprintf("%s %s", "id", updateDesc),
		},
		{
			name:         "create",
			command:      create,
			dataTypeDesc: createDesc,
			want:         createDesc,
		},
		{
			name:         "delete",
			command:      del,
			dataTypeDesc: delDesc,
			want:         "id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := getDesc(tt.command, tt.dataTypeDesc)
			assert.Equal(t, tt.want, resp)
		})
	}
}
