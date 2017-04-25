package input_string

import "strings"

func Clear_input(text string) string {
  return strings.Trim(text, "\n")
}
