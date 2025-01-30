package api

import (
	"database/sql"
	"gateway-router/db/sqlc"
	"gateway-router/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VxlanTunnelParams struct {
	Name      string `json:"name" binding:"required"`
	Tag       int64  `json:"tag" binding:"required,min=1,max=16777215"`
	TunnelIp  string `json:"tunnel_ip" binding:"required,cidr"`
	LocalIp   string `json:"local_ip" binding:"required,ipv4"`
	RemoteIp  string `json:"remote_ip" binding:"required,ipv4"`
	RemoteMac string `json:"remote_mac" binding:"required,mac"`
}

func (server *Server) CreateVxlanTunnel(ctx *gin.Context) {
	var tunnelArgs VxlanTunnelParams
	if err := ctx.ShouldBind(&tunnelArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	args := sqlc.CreateVxlanTunnelParams{
		Name:      tunnelArgs.Name,
		Tag:       tunnelArgs.Tag,
		TunnelIp:  tunnelArgs.TunnelIp,
		LocalIp:   tunnelArgs.LocalIp,
		RemoteIp:  tunnelArgs.RemoteIp,
		RemoteMac: tunnelArgs.RemoteMac,
	}
	tunnel, err := server.queries.CreateVxlanTunnel(ctx, args)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	message := TaskMessage{
		TaskID:   tunnel.ID,
		TaskType: "vxlan_tunnel",
		Data:     tunnel,
		Action:   "create",
	}
	if err := SendToQueue(message); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tunnel)
}

type VxlanGetParams struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) DeleteVxlanTunnel(ctx *gin.Context) {
	var tunnelArgs VxlanGetParams
	if err := ctx.ShouldBindUri(&tunnelArgs); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	vxlanIface, err := server.queries.GetVxlanTunnel(ctx, tunnelArgs.Id)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	message := TaskMessage{
		TaskID:   vxlanIface.ID,
		TaskType: "vxlan_tunnel",
		Data:     vxlanIface,
		Action:   "delete",
	}
	if err := SendToQueue(message); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// msg := fmt.Sprintf("successfully deleted vxlan tunnel %d", tunnelArgs.Id)
	ctx.JSON(http.StatusCreated, nil)
}

func (server *Server) GetVxlanTunnel(ctx *gin.Context) {
	var tunnelArgs VxlanGetParams
	if err := ctx.ShouldBindUri(&tunnelArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	tunnel, err := server.queries.GetVxlanTunnel(ctx, tunnelArgs.Id)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	} else if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tunnel)
}

type ListTunnelsParams struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListVxlanTunnel(ctx *gin.Context) {
	var tunnelArgs ListTunnelsParams
	if err := ctx.ShouldBindQuery(&tunnelArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	args := sqlc.ListVxlanTunnelParams{
		Limit:  tunnelArgs.PageSize,
		Offset: (tunnelArgs.PageId - 1) * tunnelArgs.PageSize,
	}
	tunnels, err := server.queries.ListVxlanTunnel(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tunnels)
}
