package password

// CheckPassword returns an error if the input string is an invalid password per the
// rules below. Otherwise, it returns nil.
//
// MUST have a capital letter
// MUST contain at least one number
// MUST contain a punctuation mark or mathematical symbol
// MUST NOT contain the word "password"
// MUST be longer than 7 characters and shorter than 31 characters
//
func CheckPassword(pass string) error {
	panic("implement me")
}