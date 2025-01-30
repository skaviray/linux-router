package api

import (
	"context"
	"database/sql"
	"fmt"
	"gateway-router/db/sqlc"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VlanSpec struct {
	Name      string `json:"name"`
	Ipaddress string `json:"ipaddress"`
	Netmask   string `json:"netmask"`
	Lower     string `json:"lower"`
	Tag       int64  `json:"tag"`
}

func CreateVlanInterface(ctx *gin.Context) {
	interfaceArgs := sqlc.CreateVlanParams{}
	if err := ctx.ShouldBind(&interfaceArgs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	iface, err := sqlc.Db.GetInterface(context.Background(), interfaceArgs.Lower)
	if err == sql.ErrNoRows {
		msg := fmt.Sprintf("unable to find the interface with id %d", interfaceArgs.Lower)
		ctx.JSON(http.StatusNotFound, gin.H{"message": msg})
		return
	} else if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	vlanGetParams := sqlc.GetVlanByLowerAndTagParams{
		Tag:   interfaceArgs.Tag,
		Lower: interfaceArgs.Lower,
	}
	vlans, _ := sqlc.Db.GetVlanByLowerAndTag(context.Background(), vlanGetParams)
	if len(vlans) != 0 {
		msg := fmt.Sprintf("vlan interface already exists with tag %d", interfaceArgs.Tag)
		ctx.JSON(http.StatusConflict, gin.H{"message": msg})
		return
	}
	vlanIface, err := sqlc.Db.CreateVlan(context.Background(), interfaceArgs)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	vlanSpec := VlanSpec{
		Name:      vlanIface.Name,
		Ipaddress: vlanIface.Ipaddress,
		Netmask:   vlanIface.Netmask,
		Tag:       vlanIface.Tag,
		Lower:     iface.Name,
	}
	message := TaskMessage{
		TaskID:   vlanIface.ID,
		TaskType: "vlan",
		Data:     vlanSpec,
		Action:   "create",
	}
	if err := SendToQueue(message); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vlanIface)
}

func DeleteVlanInterface(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "id is required"})
		return
	}
	intid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	idInt64 := int64(intid)
	vlanIface, err := sqlc.Db.GetVlan(context.Background(), idInt64)
	if err == sql.ErrNoRows {
		msg := fmt.Sprintf("unable to find the vlan %d", idInt64)
		ctx.JSON(http.StatusNotFound, gin.H{"message": msg})
		return
	} else if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	iface, err := sqlc.Db.GetInterface(context.TODO(), vlanIface.Lower)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	vlanSpec := VlanSpec{
		Name:      vlanIface.Name,
		Ipaddress: vlanIface.Ipaddress,
		Netmask:   vlanIface.Netmask,
		Tag:       vlanIface.Tag,
		Lower:     iface.Name,
	}
	message := TaskMessage{
		TaskID:   vlanIface.ID,
		TaskType: "vlan",
		Data:     vlanSpec,
		Action:   "delete",
	}
	if err := SendToQueue(message); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// if err := sqlc.Db.DeleteVlan(context.Background(), idInt64); err != nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	// 	return
	// }
	msg := fmt.Sprintf("successfully deleted vlan %d", idInt64)
	ctx.JSON(http.StatusCreated, gin.H{"message": msg})
}

func GetVlanInterface(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "id is required"})
		return
	}
	intid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	idInt64 := int64(intid)
	vlan, err := sqlc.Db.GetVlan(context.Background(), idInt64)
	if err == sql.ErrNoRows {
		msg := fmt.Sprintf("unable to find the vlan %d", idInt64)
		ctx.JSON(http.StatusNotFound, gin.H{"message": msg})
		return
	} else if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vlan)
}

func GetVlanInterfaces(ctx *gin.Context) {
	args := sqlc.ListVlansParams{
		Limit:  5,
		Offset: 0,
	}
	vlans, err := sqlc.Db.ListVlans(context.Background(), args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vlans)
}
