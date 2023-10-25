package validator

import (
	"github.com/gofiber/contrib/websocket"
)

func ReadJSONAndValidate(c *websocket.Conn, v interface{}) []ValidationErrorResponse {
	if err := c.ReadJSON(v); err != nil {
		return []ValidationErrorResponse{
			{
				Error:       true,
				FailedField: "body",
				Tag:         err.Error(),
			},
		}
	}
	if errArr := VOAHValidator.Validate(v); len(errArr) != 0 {
		return errArr
	}
	return nil
}
