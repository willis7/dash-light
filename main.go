package main

import (
	"github.com/savaki/go.hue"
	"github.com/spf13/viper"
	"github.com/willis7/arp"
	"log"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}
}

func LightToggle(light *hue.Light) func() {
	return func() {
		attr, err := light.GetLightAttributes()
		if err != nil {
			log.Print(err)
		}

		log.Println("Toggle light")
		switch attr.State.On {
		case true:
			light.Off()
		case false:
			light.On()
		}
	}
}

func main() {
	ipAddr := viper.GetString("bridge.ipAddr")
	username := viper.GetString("bridge.username")
	lampName := viper.GetString("device.1.name")
	mac := viper.GetString("network.mac")
	nic := viper.GetString("network.nic")

	bridge := hue.NewBridge(ipAddr, username)
	light, err := bridge.FindLightByName(lampName)
	if err != nil {
		log.Print(err)
	}

	lamp1 := arp.ActionerFunc(LightToggle(light))
	devs := []arp.Device{{lampName, mac, lamp1}}
	arp.Sniff(devs, nic)
}
