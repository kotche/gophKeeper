package transport

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommander_getMetaInfo(t *testing.T) {
	commander := NewCommander(nil, nil, nil)

	tests := []struct {
		name        string
		in          string
		indMetaWant int
		metaWant    string
	}{
		{
			name:        "registration",
			in:          fmt.Sprintf("%s u1 p1 %s meta info", registration, metaInfo),
			indMetaWant: 3,
			metaWant:    "meta info",
		},
		{
			name:        "registration_no_meta",
			in:          fmt.Sprintf("%s u1 p1", registration),
			indMetaWant: -1,
			metaWant:    "",
		},
		{
			name:        "text_create_meta",
			in:          fmt.Sprintf("%s %s some text %s meta info", create, textDataType, metaInfo),
			indMetaWant: 4,
			metaWant:    "meta info",
		},
		{
			name:        "login_pass_update_meta",
			in:          fmt.Sprintf("%s %s l1 p1 %s meta info", update, textDataType, metaInfo),
			indMetaWant: 4,
			metaWant:    "meta info",
		},
		{
			name:        "login_pass_update_no_meta",
			in:          fmt.Sprintf("%s %s l1 p1", update, textDataType),
			indMetaWant: -1,
			metaWant:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blocks := strings.Split(tt.in, " ")
			ind, meta := commander.getMetaInfo(tt.in, blocks)
			assert.Equal(t, tt.indMetaWant, ind)
			assert.Equal(t, tt.metaWant, meta)
		})
	}
}
