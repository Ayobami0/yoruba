package token

import "unicode"

func IsDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func IsAlpha(ch byte) bool {
	if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || ch == '_' {
		return true
	}
	return false
}

func IsSpace(ch byte) bool {
  return unicode.IsSpace(rune(ch))
}
