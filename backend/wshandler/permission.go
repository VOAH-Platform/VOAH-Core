package wshandler

import (
	"fmt"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/checkperm"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type Message struct {
	Message string `json:"message" validate:"required"`
}
type CheckIDRequest struct {
	UserID string `json:"user-id" validate:"required,uuid4"`
}
type NewMessageResponse struct {
	Permissions []models.Permission `json:"permissions"`
}
type AliveCheckRequest struct {
	Alive bool `json:"writing"`
}

func PermissionWebsocket() func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		var alive bool = true
		var lastRecieved time.Time = time.Now()

		c.SetCloseHandler(func(code int, text string) error {
			alive = false
			return nil
		})
		var err error
		if err == nil {
			go func() {
				// check user access-token
				if err = c.WriteJSON(&Message{Message: "Send User ID in 15Sec"}); err != nil {
					alive = false
					c.Close()
					return
				}
				var checkUserID CheckIDRequest
				if errArr := validator.ReadJSONAndValidate(c, &checkUserID); errArr != nil {
					fmt.Println(errArr)
					alive = false
					c.Close()
					return
				}
				lastRecieved = time.Now()

				db := database.DB
				var user models.User
				if err = db.Where(&models.User{ID: uuid.MustParse(checkUserID.UserID)}).First(&user).Error; err != nil {
					fmt.Println(err)
					c.WriteJSON(&Message{Message: "User not found"})
					alive = false
					c.Close()
					return
				}

				// check user permissions
				go func() {
					for alive {
						permissions, err := checkperm.GetUserPermissionArr(&user)
						if err != nil {
							alive = false
							c.Close()
							return
						}
						if err = c.WriteJSON(&NewMessageResponse{Permissions: permissions}); err != nil {
							alive = false
							c.Close()
							return
						}
						time.Sleep(time.Second * 5)
					}
				}()

				// check client is writing
				go func() {
					var aliveCheck AliveCheckRequest
					for alive {
						if err = c.ReadJSON(&aliveCheck); err != nil {
							alive = false
							c.Close()
							return
						}
						lastRecieved = time.Now()
					}
				}()
			}()
			WSTimeOut(c, &alive, &lastRecieved, 15)
		}
	}
}
