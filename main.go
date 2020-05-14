
package main


import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/psychedelicnekopunch/go-slack/app/infrastructure"
)


func main()  {
	c := infrastructure.NewConfig()
	api := slack.New(c.Slack.Token)

	// WebSocket
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		// fmt.Print("Event Received: \n")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			// fmt.Println("Infos:", ev.Info)
			// fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID
			// rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C2147483705"))

		case *slack.MessageEvent:
			// 新しいメッセージが送信された場合 ClientMsgID が付与される
			if fmt.Sprintf("%v", ev.ClientMsgID) == "" {
				break
			}
			// fmt.Printf("Message: %v\n", ev)
			// fmt.Printf("Message: %v\n", ev.ClientMsgID)
			// fmt.Printf("%v\n", ev.Channel)
			rtm.SendMessage(rtm.NewOutgoingMessage("せやな", fmt.Sprintf("%v", ev.Channel)))
		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			// fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.DesktopNotificationEvent:
			// fmt.Printf("Desktop Notification: %v\n", ev)
		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
