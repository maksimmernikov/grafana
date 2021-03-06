package notifiers

import (
	"testing"

	"github.com/maksimmernikov/grafana/pkg/components/simplejson"
	m "github.com/maksimmernikov/grafana/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSlackNotifier(t *testing.T) {
	Convey("Slack notifier tests", t, func() {

		Convey("Parsing alert notification from settings", func() {
			Convey("empty settings should return error", func() {
				json := `{ }`

				settingsJSON, _ := simplejson.NewJson([]byte(json))
				model := &m.AlertNotification{
					Name:     "ops",
					Type:     "slack",
					Settings: settingsJSON,
				}

				_, err := NewSlackNotifier(model)
				So(err, ShouldNotBeNil)
			})

			Convey("from settings", func() {
				json := `
				{
          "url": "http://google.com"
				}`

				settingsJSON, _ := simplejson.NewJson([]byte(json))
				model := &m.AlertNotification{
					Name:     "ops",
					Type:     "slack",
					Settings: settingsJSON,
				}

				not, err := NewSlackNotifier(model)
				slackNotifier := not.(*SlackNotifier)

				So(err, ShouldBeNil)
				So(slackNotifier.Name, ShouldEqual, "ops")
				So(slackNotifier.Type, ShouldEqual, "slack")
				So(slackNotifier.Url, ShouldEqual, "http://google.com")
				So(slackNotifier.Recipient, ShouldEqual, "")
				So(slackNotifier.Username, ShouldEqual, "")
				So(slackNotifier.IconEmoji, ShouldEqual, "")
				So(slackNotifier.IconUrl, ShouldEqual, "")
				So(slackNotifier.Mention, ShouldEqual, "")
				So(slackNotifier.Token, ShouldEqual, "")
			})

			Convey("from settings with Recipient, Username, IconEmoji, IconUrl, Mention, and Token", func() {
				json := `
				{
          "url": "http://google.com",
          "recipient": "#ds-opentsdb",
          "username": "Grafana Alerts",
          "icon_emoji": ":smile:",
          "icon_url": "https://grafana.com/img/fav32.png",
          "mention": "@carl",
          "token": "xoxb-XXXXXXXX-XXXXXXXX-XXXXXXXXXX"
				}`

				settingsJSON, _ := simplejson.NewJson([]byte(json))
				model := &m.AlertNotification{
					Name:     "ops",
					Type:     "slack",
					Settings: settingsJSON,
				}

				not, err := NewSlackNotifier(model)
				slackNotifier := not.(*SlackNotifier)

				So(err, ShouldBeNil)
				So(slackNotifier.Name, ShouldEqual, "ops")
				So(slackNotifier.Type, ShouldEqual, "slack")
				So(slackNotifier.Url, ShouldEqual, "http://google.com")
				So(slackNotifier.Recipient, ShouldEqual, "#ds-opentsdb")
				So(slackNotifier.Username, ShouldEqual, "Grafana Alerts")
				So(slackNotifier.IconEmoji, ShouldEqual, ":smile:")
				So(slackNotifier.IconUrl, ShouldEqual, "https://grafana.com/img/fav32.png")
				So(slackNotifier.Mention, ShouldEqual, "@carl")
				So(slackNotifier.Token, ShouldEqual, "xoxb-XXXXXXXX-XXXXXXXX-XXXXXXXXXX")
			})
		})
	})
}
