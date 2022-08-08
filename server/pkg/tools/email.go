package tools

// MakeVerifyEmailLink ...
func MakeVerifyEmailLink(email string) string {
	return "https://moyzavod.com/verify?email=" + email + "&code=" + PasswordHash(email)
}
