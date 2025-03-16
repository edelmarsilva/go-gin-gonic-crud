package route

import (
	"crud-go/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {

	//Utiliza r como ponteiro para o pacote Gin Gonic
	//r recebeu o atribuito GET e define a url para acesso, seguida uma paramentro que deve ser passado o userId
	r.GET("/getUserById:userId", controller.FindUserById)
	r.GET("/getUserByEmail:userEmail", controller.FindUserByEmail)

	//Metodo post para criar usuario nao recebe parametro
	r.POST("/createUser", controller.CreateUser)

	//Metodo put para receber ID do usuario que deseja realizar update
	r.PUT("/updateUser:userId", controller.UpdateUser)

	//Metodo para deletar um usuario
	r.DELETE("/deleteUser:userId", controller.DeleteUserById)

}
