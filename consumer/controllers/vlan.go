package controllers

import (
	"github.com/gin-gonic/gin"
)

type VlanSpec struct {
	Name      string `json:"name"`
	Ipaddress string `json:"ipaddress"`
	Netmask   string `json:"netmask"`
	Lower     int64  `json:"lower"`
	Tag       int64  `json:"tag"`
}

type VlanResponse struct {
}

func CreateVlanInterface() {

}

func DeleteVlanInterface() {
}

func GetVlanInterface() {
}

func GetVlanInterfaces(ctx *gin.Context) {
}
