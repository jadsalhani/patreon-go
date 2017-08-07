package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchCampaign(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/oauth2/api/current_user/campaigns", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, fetchCampaignResp)
	})

	resp, err := client.FetchCampaign()
	require.NoError(t, err)

	require.Equal(t, 1, len(resp.Data))
	require.Equal(t, "campaign", resp.Data[0].Type)
	require.Equal(t, "278915", resp.Data[0].Id)
	require.Equal(t, "new podcasting experience - Podsync", resp.Data[0].Attributes.CreationName)
	require.Equal(t, 8, resp.Data[0].Attributes.CreationCount)
	require.True(t, resp.Data[0].Attributes.DisplayPatronGoals)
	require.NotEmpty(t, resp.Data[0].Attributes.ImageSmallURL)
	require.NotEmpty(t, resp.Data[0].Attributes.ImageURL)
	require.True(t, resp.Data[0].Attributes.IsChargedImmediately)
	require.True(t, resp.Data[0].Attributes.IsMonthly)
	require.True(t, resp.Data[0].Attributes.IsNsfw)
	require.True(t, resp.Data[0].Attributes.IsPlural)
	require.Equal(t, 123121, resp.Data[0].Attributes.PatronCount)
	require.Equal(t, "month", resp.Data[0].Attributes.PayPerName)
	require.Equal(t, 12321312, resp.Data[0].Attributes.PledgeSum)
	require.NotEmpty(t, resp.Data[0].Attributes.Summary)
	require.NotEmpty(t, resp.Data[0].Attributes.PledgeURL)
	require.NotEmpty(t, resp.Data[0].Attributes.ThanksMsg)
}

const fetchCampaignResp = `
{
    "data": [
        {
            "attributes": {
                "created_at": "2016-02-02T19:58:18+00:00",
                "creation_count": 8,
                "creation_name": "new podcasting experience - Podsync",
                "discord_server_id": null,
                "display_patron_goals": true,
                "earnings_visibility": null,
                "image_small_url": "https://c10.patreon.com/3/eyJoIjoxMjgwLCJ3IjoxMjgwfQ%3D%3D/patreon-user/AS2N2NrZauWDuhVcuua87P7QtSOdCtPWiazP99SpvJWWHn8d4GvZI56AHqTn94g2_large_2.png?token-time=2145916800&token-hash=VNcCmkq6bOjbjizpwNHePu3aNqujVMKJYdfaPxoz3_c%3D",
                "image_url": "https://c10.patreon.com/3/eyJ3IjoxOTIwfQ%3D%3D/patreon-user/AS2N2NrZauWDuhVcuua87P7QtSOdCtPWiazP99SpvJWWHn8d4GvZI56AHqTn94g2_large_2.png?token-time=2145916800&token-hash=4KxOxPVCtGwPskLYr8BZGyZW94VwAKbD7j9RDHcvf0E%3D",
                "is_charged_immediately": true,
                "is_monthly": true,
                "is_nsfw": true,
                "is_plural": true,
                "main_video_embed": "",
                "main_video_url": "",
                "one_liner": null,
                "outstanding_payment_amount_cents": 0,
                "patron_count": 123121,
                "pay_per_name": "month",
                "pledge_sum": 12321312,
                "pledge_url": "/bePatron?c=278915",
                "published_at": "2016-02-02T20:11:19+00:00",
                "summary": "<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> - is a simple, free service that lets you listen to any YouTube / Vimeo channels, playlists or user videos in podcast format.<br><br><strong>Idea:</strong><br>Podcast applications have a rich functionality for content delivery - automatic download of new episodes, remembering last played position, sync between devices and offline listening. This functionality is not available on YouTube and Vimeo. So the aim of\u00a0<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> is to make your life easier and enable you to view/listen to content on any device in podcast client.<br><br>It's my hobby project, so to continue to support and improve it, I need your help. Your money will go into paying my server bills and adding new features.<br><br>",
                "thanks_embed": "",
                "thanks_msg": "You are awesome!",
                "thanks_video_url": null
            },
            "id": "278915",
            "type": "campaign"
        }
    ]
}
`
