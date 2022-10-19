package transport

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

// Completer shows hints for entering commands
func (c *Commander) Completer(d prompt.Document) []prompt.Suggest {
	var sug []prompt.Suggest

	if strings.Contains(d.Text, registration) {
		sug = []prompt.Suggest{
			{Text: registration, Description: registrationDesc},
		}
	} else if strings.Contains(d.Text, authentication) {
		sug = []prompt.Suggest{
			{Text: authentication, Description: authenticationDesc},
		}
	} else if strings.Contains(d.Text, create) {
		sug = getDataTypeText(create, d.Text)
	} else if strings.Contains(d.Text, read) {
		sug = getDataTypeText(read, d.Text)
	} else if strings.Contains(d.Text, update) {
		sug = getDataTypeText(update, d.Text)
	} else if strings.Contains(d.Text, del) {
		sug = getDataTypeText(del, d.Text)
	} else {
		sug = []prompt.Suggest{
			{Text: registration, Description: registrationDesc},
			{Text: authentication, Description: authenticationDesc},
			{Text: create, Description: createDesc},
			{Text: read, Description: readDesc},
			{Text: update, Description: updateDesc},
			{Text: del, Description: delDesc},
		}
	}
	return prompt.FilterHasPrefix(sug, d.GetWordBeforeCursor(), true)
}

// getDataTypeText defines the data type
func getDataTypeText(command, text string) []prompt.Suggest {
	var sug []prompt.Suggest

	if strings.Contains(text, loginPassDataType) {
		sug = []prompt.Suggest{
			{Text: fmt.Sprint(command, " ", loginPassDataType), Description: getDesc(command, loginPassDataTypeDesc)},
		}
	} else if strings.Contains(text, textDataType) {
		sug = []prompt.Suggest{
			{Text: fmt.Sprint(command, " ", textDataType), Description: getDesc(command, textDataTypeDesc)},
		}
	} else if strings.Contains(text, binaryDataType) {
		sug = []prompt.Suggest{
			{Text: fmt.Sprint(command, " ", binaryDataType), Description: getDesc(command, binaryDataTypeDesc)},
		}
	} else if strings.Contains(text, bankCardDataType) {
		sug = []prompt.Suggest{
			{Text: fmt.Sprint(command, " ", bankCardDataType), Description: getDesc(command, bankCardDataTypeDesc)},
		}
	} else {
		sug = []prompt.Suggest{
			{Text: loginPassDataType, Description: loginPassDataTypeDesc},
			{Text: textDataType, Description: textDataTypeDesc},
			{Text: binaryDataType, Description: binaryDataTypeDesc},
			{Text: bankCardDataType, Description: bankCardDataTypeDesc},
		}
	}
	return sug
}

func getDesc(command, dataTypeDesc string) string {
	switch command {
	case update:
		return fmt.Sprintf("%s %s", "id", dataTypeDesc)
	case del:
		return "id"
	default:
		return dataTypeDesc
	}
}
