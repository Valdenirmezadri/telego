package api_test

import (
	"testing"

	"github.com/davilag/telego/api"
	"github.com/davilag/telego/kind"
)

func TestMessage_GetKind(t *testing.T) {
	location := api.Location{
		Longitude: 0.00000,
		Latitude:  0.00000,
	}
	contact := api.Contact{
		UserID: 1,
	}

	photo := []api.PhotoSize{
		{
			FileID: "1",
		},
	}

	sticker := api.Sticker{
		FileID: "1",
	}

	video := api.Video{
		FileID: "1",
	}

	voice := api.Voice{
		FileID: "1",
	}

	videoNote := api.VideoNote{
		FileID: "1",
	}
	tests := []struct {
		m api.Message
		k kind.Kind
	}{
		{
			m: api.Message{
				Location: &location,
			},
			k: kind.Location,
		},
		{
			m: api.Message{
				Contact: &contact,
			},
			k: kind.Contact,
		},
		{
			m: api.Message{
				Photo: &photo,
			},
			k: kind.Photo,
		},
		{
			m: api.Message{
				Sticker: &sticker,
			},
			k: kind.Sticker,
		},
		{
			m: api.Message{
				Video: &video,
			},
			k: kind.Video,
		},
		{
			m: api.Message{
				Voice: &voice,
			},
			k: kind.VoiceMessage,
		},
		{
			m: api.Message{
				VideoNote: &videoNote,
			},
			k: kind.VideoNote,
		},
		{
			m: api.Message{},
			k: kind.Text,
		},
	}

	for _, test := range tests {
		testK := test.m.GetKind()
		if test.k != testK {
			t.Fatalf("Expected to be %s, but got %s", string(test.k), string(testK))
		}
	}
}
