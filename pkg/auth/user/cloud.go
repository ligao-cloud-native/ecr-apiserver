package user

func VerifyCloudAccount(username, pwd string) bool {
	enableCloudAuth := true
	if enableCloudAuth {
		return cloudAuth()
	}
	return false
}

func cloudAuth() bool {
	return false
}
