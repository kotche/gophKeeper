package transport

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

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
	} else if strings.Contains(d.Text, delete) {
		sug = getDataTypeText(delete, d.Text)
	} else {
		sug = []prompt.Suggest{
			{Text: registration, Description: registrationDesc},
			{Text: authentication, Description: authenticationDesc},
			{Text: create, Description: createDesc},
			{Text: read, Description: readDesc},
			{Text: update, Description: updateDesc},
			{Text: delete, Description: deleteDesc},
		}
	}
	return prompt.FilterHasPrefix(sug, d.GetWordBeforeCursor(), true)
}

func getDataTypeText(command, text string) []prompt.Suggest {
	var sug []prompt.Suggest

	if strings.Contains(text, loginPassDataType) {
		sug = []prompt.Suggest{
			{Text: fmt.Sprint(command, " ", loginPassDataType), Description: loginPassDataTypeDesc},
		}
	} else if strings.Contains(text, textDataType) {
		sug = []prompt.Suggest{
			{Text: fmt.Sprint(command, " ", textDataType), Description: textDataTypeDesc},
		}
	} else if strings.Contains(text, binaryDataType) {
		sug = []prompt.Suggest{
			{Text: fmt.Sprint(command, " ", binaryDataType), Description: binaryDataTypeDesc},
		}
	} else if strings.Contains(text, bankCardDataType) {
		sug = []prompt.Suggest{
			{Text: fmt.Sprint(command, " ", bankCardDataType), Description: bankCardDataTypeDesc},
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
