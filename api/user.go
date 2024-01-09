package api

import (
	"net/http"

	db "new/learning/user/db/sqlc"
	"new/learning/user/util"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword,err := util.HashPassword(req.Password)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:    req.Username,
		HashedPassword: hashedPassword,
		FullName: req.FullName,
		Email: req.Email,
	}

	User, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok{
			switch pqErr.Code.Name(){
			case "foreign_key_violation","unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, User)
}

// type getAccountRequest struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) getAccount(ctx *gin.Context) {
// 	var req getAccountRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	account, err := server.store.GetAccount(ctx, req.ID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, account)
// }

// // type ListAccountRequest struct {
// // 	pageID   int32 `form:"page_id" binding:"required,min=1"`
// // 	pageSize int32 `form:"page_id" binding:"required,min=5,max=10"`
// // }

// // func (server *Server) ListAccount(ctx *gin.Context) {
// // 	var req ListAccountRequest
// // 	if err := ctx.ShouldBindQuery(&req); err != nil {
// // 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// // 		return
// // 	}
// // 	arg := db.ListAccountsParams{
// // 		Limit:  req.pageSize,
// // 		Offset: (req.pageID - 1) * req.pageSize,
// // 	}

// // 	accounts, err := server.store.ListAccounts(ctx, arg)
// // 	if err != nil {
// // 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// // 		return

// // 	}

// // 	ctx.JSON(http.StatusOK, accounts)
// // }
