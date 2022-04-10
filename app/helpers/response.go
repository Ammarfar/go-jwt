package helpers

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   bool        `json:"error"`
}

func ResponseSuccess(message string, data ...interface{}) response {
	return response{
		Message: message,
		Data:    Data(data),
		Error:   false,
	}
}

func ResponseError(message string, data ...interface{}) response {
	return response{
		Message: message,
		Data:    Data(data),
		Error:   true,
	}
}

func Data(data []interface{}) interface{} {
	if data != nil {
		return data[0]
	} else {
		return nil
	}
}
