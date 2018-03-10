package parser

// Token represents the token of the configuration.
type Token struct {
	tok      int
	lit      string
	filename string
	line     int
	column   int
}
