package utils

const(
	MinSize = "minSize"
	MinLowercase = "minLowercase"
	MinUppercase = "minUppercase"
	MinDigit = "minDigit"
	MinSpecialChars = "minSpecialChars"
	NoRepeated = "noRepeated"
	SpecialCharacterNotFound = "special character not found"

	FailedtoCheckPassword = "failed to check password"

	HttpPort = ":8080"
	APIStopped = "server stopped successfully"
)

var(
	ApiStartMsg = "API listening on port %s \n"
	ApiStartErr = "failed to start server: %v"
	ApiForcedShutdown = "server forced to shutdown: %s"
)