package server

import (
	"fmt"
	"log"
	"time"
	"github.com/chilts/sid"
	"github.com/gin-gonic/gin"

	consul "github.com/micro/go-micro/registry/consul"	
	"github.com/micro/go-micro/web"
)


func Start() web.Service {

	consl := consul.NewRegistry()

	web.DefaultAddress=":8080"
	web.DefaultName="GPDBNode"
	
	service := web.NewService(
		web.Id(fmt.Sprintf("%s-%s", web.DefaultName, sid.Id())),			
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
	)	

	service.Init(func(o *web.Options) {
		o.Registry = consl
	})

	return service
}

func RegistryRouter(service web.Service,router *gin.Engine) {
	service.Handle("/",router)
	log.Printf("Agent Server Dispatch Root Path")
	// Run server
	if err := service.Run(); err != nil {	
		log.Printf("Agent Server error")	
		log.Fatal(err)
	}
}