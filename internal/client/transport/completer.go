package transport

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func (s *Commander) Completer(d prompt.Document) []prompt.Suggest {
	var sug []prompt.Suggest

	if strings.Contains(d.Text, registration) {
		sug = []prompt.Suggest{
			{Text: registration, Description: registrationDesc},
		}
	} else if strings.Contains(d.Text, authentication) {
		sug = []prompt.Suggest{
			{Text: authentication, Description: "user authentication"},
		}
	} else if d.Text == exit {
		sug = []prompt.Suggest{
			{Text: exit, Description: "exit program"},
		}
	} else {
		sug = []prompt.Suggest{
			{Text: registration, Description: registrationDesc},
			{Text: authentication, Description: "user authentication"},
			{Text: exit, Description: "exit program"},
		}
	}
	return prompt.FilterHasPrefix(sug, d.GetWordBeforeCursor(), true)
}
