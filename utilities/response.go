package utilities

func OK(data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": 200,
		"data":       data,
	}
}

func KO(statusCode int, message string) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": statusCode,
		"error":      message,
	}
}
