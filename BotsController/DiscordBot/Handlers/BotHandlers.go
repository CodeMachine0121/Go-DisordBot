package Handlers

import (
	"log"
	"os"

	"golang-discord-bot/BotsController/GlobalSetting"

	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/vault/api"
)

type BotHandlers struct {
	BotConfig DiscordBotConfig
}

func ErrorHandle(err error) {
	if err != nil {
		log.Println("something is error, pls contact support!")
		log.Fatal(err)
		return
	}
}

type DiscordBotConfig struct {
	Token string
}

func GetConfigWithToken(vaultToken string) DiscordBotConfig {
	secret := GetConnectionWithVault(vaultToken)

	data := secret.Data["data"].(map[string]interface{})
	for key, value := range data {
		if key == "token" {
			return DiscordBotConfig{Token: value.(string)}
		}
	}
	return DiscordBotConfig{}
}

func GetConnectionWithVault(vaultToken string) *api.Secret {
	client, err := api.NewClient(&api.Config{
		Address: GlobalSetting.VaultAddress,
	})
	ErrorHandle(err)

	client.SetToken(vaultToken)

	secret, err := client.Logical().Read(GlobalSetting.VaultTokenPath)
	ErrorHandle(err)
	return secret
}

func GetVaultToken() string {
	vaultToken, err := os.ReadFile("./setting.json")
	ErrorHandle(err)

	log.Println("Get Vault Token: " + string(vaultToken))
	return string(vaultToken)
}

func SendMessage(session *discordgo.Session, channelID string, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	ErrorHandle(err)
}
