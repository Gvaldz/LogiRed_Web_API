package infrastructure

import (
	"LogiredAPIWeb/src/internal/users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	CreateUserController     	*controllers.CreateUserController
	GetAllUsersController    	*controllers.GetAllUsersController
	GetUserController        	*controllers.GetByUserIDController
	UpdateUserController     	*controllers.UpdateUserController
	UpdatePasswordController 	*controllers.UpdatePasswordController
	RessetPasswordController 	*controllers.RessetPasswordController
	DeleteUserController     	*controllers.DeleteUserController
	AuthMiddleware           	gin.HandlerFunc
}

func NewUserRoutes(
	createUserController 		*controllers.CreateUserController,
	getAllUsersController 		*controllers.GetAllUsersController,
	getUserController 			*controllers.GetByUserIDController,
	updateUserController 		*controllers.UpdateUserController,
	updatePasswordController 	*controllers.UpdatePasswordController,
	ressetPasswordController 	*controllers.RessetPasswordController,
	deleteUserController 		*controllers.DeleteUserController,
	authMiddleware 				gin.HandlerFunc,
) *UserRoutes {
	return &UserRoutes{
		CreateUserController:     createUserController,
		GetAllUsersController:    getAllUsersController,
		GetUserController:        getUserController,
		UpdateUserController:     updateUserController,
		UpdatePasswordController: updatePasswordController,
		RessetPasswordController: ressetPasswordController,
		DeleteUserController:     deleteUserController,
		AuthMiddleware:           authMiddleware,
	}
}

func (r *UserRoutes) AttachRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", r.CreateUserController.Create)
		userGroup.GET("", r.GetAllUsersController.GetAll)
		userGroup.GET("/:id", r.GetUserController.GetByUserID)

		userGroup.PUT("/password-reset", r.RessetPasswordController.RessetPassword)

		protected := userGroup.Group("")
		protected.Use(r.AuthMiddleware)
		{
			protected.PUT("/:id", r.UpdateUserController.UpdateUser)

			protected.PUT("/update-password", r.UpdatePasswordController.UpdatePassword)

			protected.DELETE("/:id", r.DeleteUserController.Delete)
		}
	}
}
