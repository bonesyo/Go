package turbo-happiness

import (
    "log"
    "image"
    "github.com/cloud6/px/processors"
    "github.com/cloud6/px/px"
)

type (
    Options struct {
        Width  float64 `json:"width" cmd:"width"`
        Height float64 `json:"height" cmd:"height"`
    }
)

func Placejames(img *px.Image, opts interface{}, conf *px.Config) (image.Image, interface{}, error){

    log.Println("YAY Resize")
    return nil, nil, nil
}

func init() {
    processors.RegisterModule("placejames", Placejames, &Options{})
}
