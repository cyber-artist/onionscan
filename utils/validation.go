package utils

func IsOnion(identifier string) bool {
	// TODO: At some point we will want to support i2p
	/*if len(identifier) >= 22 && strings.HasSuffix(identifier, ".onion") {
		matched, _ := regexp.MatchString(`(^|\.)[a-z2-7]{16}\.onion$`, identifier)
		return matched
	}*/
	return true
}
