package utils

import (
  "golang.org/x/text/language"
  "golang.org/x/text/message"
)

var Formatter = message.NewPrinter(language.English)
