package turbohappiness

import (
    "log"
    "image"
    "github.com/cloud6/ducking-octo-lana/processors"
    "github.com/cloud6/ducking-octo-lana/img"
    "github.com/cloud6/ducking-octo-lana/config"
)

type (
    Options struct {
        Width  float64 `json:"width" cmd:"width"`
        Height float64 `json:"height" cmd:"height"`
    }
)

func Placejames(img *img.Image, opts interface{}, conf *config.Config) (image.Image, interface{}, error){

    log.Println("YAY Resize")
    return nil, nil, nil
}

func init() {
    processors.RegisterModule("placejames", Placejames, &Options{})
}
