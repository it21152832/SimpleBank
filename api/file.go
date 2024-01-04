package api

import (
	"net/http"

	db "new/learning/user/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createFileRequest struct {
	FileName string `json:"file_name" binding:"required"`
	Owner    string `json:"owner" binding:"required`
}

func (server *Server) CreateFile(ctx *gin.Context) {
	var req createFileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFileParams{
		FileName:   req.FileName,
		Owner:      req.Owner,
		ChunkCount: 0,
	}

	File, err := server.store.CreateFile(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, File)
}

// type getFileRequest struct {
// 	Hash int64 `uri:"hash" binding:"required,min=1"`
// }

// func (server *Server) getFile(ctx *gin.Context) {
// 	var req getFileRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	File, err := server.store.GetFile(ctx, req.Hash)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, File)
// }

// type ListAccountRequest struct {
// 	pageID   int32 `form:"page_id" binding:"required,min=1"`
// 	pageSize int32 `form:"page_id" binding:"required,min=5,max=10"`
// }

// func (server *Server) ListAccount(ctx *gin.Context) {
// 	var req ListAccountRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	arg := db.ListAccountsParams{
// 		Limit:  req.pageSize,
// 		Offset: (req.pageID - 1) * req.pageSize,
// 	}

// 	accounts, err := server.store.ListAccounts(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return

// 	}

// 	ctx.JSON(http.StatusOK, accounts)
// }
