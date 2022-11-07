package server

import (
	"fmt"
	"log"
	"net/http"
	"selfhelp-iptables/config"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

func StartServer() {
	cfg := config.GetConfig()
	//开启go routine
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HelloServer)
	router.HandleFunc("/subscription", GetSubscription)
	router.HandleFunc("/api/add", AddWhitelist)
	router.HandleFunc("/api/ban/{ip}", AddBlackList)
	router.HandleFunc("/api/list", ShowWhitelist)
	router.HandleFunc("/api/listb", ShowBlacklist)
	router.HandleFunc("/api/log", GetLogs)
	router.HandleFunc("/api/reset", Reset)
	router.HandleFunc("/api/vnstat", Vnstat)
	router.HandleFunc("/api/record", GetRecords)
	router.HandleFunc("/api/remove/{ip}", RemoveWhitelist)
	router.HandleFunc("/api/unban/{ip}", RemoveBlacklist)

	fmt.Println("Server start Port:"+cfg.ListenPort+" UserKey:"+cfg.UserKey+" AdminKey:"+cfg.AdminKey, "\n输入help查看控制台命令帮助")
	color.Unset()
	err := http.ListenAndServe(cfg.ListenIP+":"+cfg.ListenPort, router)
	if err != nil {
		log.Fatal("Server error: " + err.Error())
	}
}
