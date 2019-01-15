package svc // import "ultre.me/calcbiz/svc"

import (
	"context"
	"fmt"

	tpyo "github.com/tpyolang/tpyo-cli"
	"ultre.me/calcbiz/api"
	"ultre.me/calcbiz/pkg/crew"
	"ultre.me/calcbiz/pkg/numberinfo"
	"ultre.me/calcbiz/pkg/random"
	"ultre.me/calcbiz/pkg/soundcloud"
	"ultre.me/kryptos"
)

type Options struct {
	SoundcloudUserID   int
	SoundcloudClientID string
}

type svc struct {
	opts       Options
	soundcloud soundcloud.Soundcloud
	// dashboard
}

func New(opts Options) (api.ServerServer, error) {
	svc := &svc{opts: opts}
	svc.soundcloud = soundcloud.New(
		opts.SoundcloudClientID,
		uint64(opts.SoundcloudUserID),
	)
	// svc.dashboard := calcdashboard.New()
	// svc.dashboard.SetSoundCloud(&soundcloud)
	return svc, nil
}

func (svc *svc) Ping(_ context.Context, input *api.Void) (*api.Pong, error) {
	return &api.Pong{Pong: "pong"}, nil
}

func (svc *svc) KryptosEncrypt(_ context.Context, input *api.KryptosInput) (*api.KryptosOutput, error) {
	return &api.KryptosOutput{
		To: kryptos.Encrypt(input.From),
	}, nil
}

func (svc *svc) KryptosDecrypt(_ context.Context, input *api.KryptosInput) (*api.KryptosOutput, error) {
	return &api.KryptosOutput{
		To: kryptos.Decrypt(input.From),
	}, nil
}

func (svc *svc) TpyoEnocde(_ context.Context, input *api.TpyoEnocdeIpunt) (*api.TpyoEnocdeOuptut, error) {
	enedocr := tpyo.NewTpyo()
	return &api.TpyoEnocdeOuptut{
		To: enedocr.Enocde(input.Form),
	}, nil
}

func (svc *svc) Dashboard(_ context.Context, input *api.Void) (*api.DashboardOutput, error) {
	/*
		r.Get("/dashboard/random", func(w http.ResponseWriter, r *http.Request) {
			dashboard, err := dashboard.Random()
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"result": dashboard,
				})
			}
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Crew(_ context.Context, input *api.Void) (*crew.Crew, error) {
	return &crew.CALC, nil
}

func (svc *svc) Numberinfo(_ context.Context, input *api.NumberinfoInput) (*api.NumberinfoOutput, error) {
	// FIXME: validate: input.Number is mandatory
	facts := map[string]string{}
	for k, v := range numberinfo.New(float64(input.Number)).All() {
		facts[k] = fmt.Sprintf("%v", v)
	}
	return &api.NumberinfoOutput{Facts: facts}, nil
}

func (svc *svc) Recettator(_ context.Context, input *api.RecettatorInput) (*api.RecettatorOutput, error) {
	/*
		r.Get("/recettator/json/:seed", func(w http.ResponseWriter, r *http.Request) {
			seed, err := strconv.ParseInt(c.Param("seed"), 10, 64)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": fmt.Sprintf("Invalid seed: %v (%v)", c.Param("seed"), err),
				})
				return
			}

			rctt := recettator.New(seed)
			rctt.SetSettings(recettator.Settings{
				MainIngredients:      2,
				SecondaryIngredients: 2,
				Steps:                5,
			})

			output := rctt.ToMap()

			c.JSON(http.StatusOK, gin.H{
				"result": output,
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Moijaime(_ context.Context, input *api.Void) (*api.MoijaimeOutput, error) {
	/*
		r.Get("/moijaime", func(w http.ResponseWriter, r *http.Request) {
			phrases := []string{}
			for i := 0; i < 20; i++ {
				phrases = append(phrases, moijaime.Generate())
			}
			c.JSON(http.StatusOK, gin.H{
				"result": phrases,
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SpreadshirtRandom(_ context.Context, input *api.Void) (*api.SpreadshirtRandomOutput, error) {
	/*
		r.Get("/spreadshirt/random", func(w http.ResponseWriter, r *http.Request) {
			c.JSON(http.StatusOK, gin.H{
				"result": calcspreadshirt.GetRandomProduct(250, 250),
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) SpreadshirtAll(_ context.Context, input *api.Void) (*api.SpreadshirtAllOutput, error) {
	/*
		r.Get("/spreadshirt/all", func(w http.ResponseWriter, r *http.Request) {
			c.JSON(http.StatusOK, gin.H{
				"result": calcspreadshirt.GetAllProducts(250, 250),
			})
		})
	*/
	return nil, fmt.Errorf("not implemented")
}

func (svc *svc) Wotd(_ context.Context, input *api.Void) (*api.WotdOutput, error) {
	return &api.WotdOutput{
		Word: random.WOTD(),
	}, nil
}

func (svc *svc) AlternateLogo(_ context.Context, input *api.Void) (*api.AlternateLogoOutput, error) {
	return &api.AlternateLogoOutput{
		Path: random.AlternateLogo(),
	}, nil
}

func (svc *svc) SoundcloudMe(_ context.Context, input *api.Void) (*soundcloud.User, error) {
	return svc.soundcloud.Me()
}

func (svc *svc) SoundcloudPlaylists(_ context.Context, input *api.Void) (*soundcloud.Playlists, error) {
	return svc.soundcloud.GetPlaylists()
}

func (svc *svc) SoundcloudPlaylist(_ context.Context, input *api.SoundcloudPlaylistInput) (*soundcloud.Playlist, error) {
	if input.PlaylistId < 1 { // pick random
		return svc.soundcloud.GetRandomPlaylist()
	}
	return svc.soundcloud.GetPlaylist(input.PlaylistId)
}

func (svc *svc) SoundcloudTracks(_ context.Context, input *api.Void) (*soundcloud.Tracks, error) {
	return svc.soundcloud.GetTracks()
}

func (svc *svc) SoundcloudTrack(_ context.Context, input *api.SoundcloudTrackInput) (*soundcloud.Track, error) {
	if input.TrackId < 1 { // pick random
		return svc.soundcloud.GetRandomTrack()
	}
	return svc.soundcloud.GetTrack(input.TrackId)
}