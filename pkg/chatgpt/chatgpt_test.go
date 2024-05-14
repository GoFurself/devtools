package chatgpt_test

import (
	"fmt"
	"testing"

	"github.com/GoFurself/devtools/pkg/chatgpt"
)

func TestStuff(t *testing.T) {

	cc := chatgpt.NewChatCompletion(
		"sk-proj-DziSlJJpn0OXyNUCMbWgT3BlbkFJLNUYjsBYNq6npcx7oWlf",
		chatgpt.ModelGPT4,
		2000,
		chatgpt.HTTPRequestHandler,
		&chatgpt.JsonMarshalHandler{},
	)

	point := `
	Kahvilatoiminta ja erilaiset aktiviteetit sekä tapahtumat Laivapuistossa sijaitsevassa uudessa kahvilassa. Tuotevalikoimaan kuuluvat sesongin mukaan jäätelöannokset, salaattiannokset, hedelmäsalaatit, leivät ja sämpylät, hot dogit, nuudeliannokset, kahvi, tee, smoothiet, virvokkeet jne. Vohveleiden myynti joko itse tai mahdollisuus antaa esim. nuorelle pöytä vohveleiden myyntiä varten. Toimintaa suunnitellaan yhteistyössä vuokranantajan eli Vaasan kaupungin kanssa, esim. skeittilautojen, skuuttien ja kypäröiden vuokrausta, näyttelytoimintaa, tapahtumia, kierrätystoimintaa (vaihtohylly kirjoille, leluille, kasveille ja taideteoksille). 
				`

	cc.AddMessage(chatgpt.ChatGPTRoleAssistant, "Seuraavaksi annan sinulle liikeidean lyhyesti. Pyytäisin sinua kommentoimaan sisältöä yritysneuvojan näkökulmasta. Esitä konkreettisia näkemyksiä, rakentavaa kritiikkiä ja ideoita. ")
	cc.AddMessage(chatgpt.ChatGPTRoleUser, point)

	response, err := cc.HandleRequest()
	if err != nil {
		fmt.Println("chatgpt return a error:", err)
		return
	}

	// Do something with the response
	fmt.Println(response.Choices[0].Message.Content)

}
