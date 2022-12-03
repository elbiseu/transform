package transform

import (
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"io"
	"log"
)

func Scale(img image.Image, x, y int) (image.Image, error) {
	rgba := image.NewRGBA(image.Rect(0, 0, x, y))
	draw.NearestNeighbor.Scale(rgba, rgba.Rect, img, img.Bounds(), draw.Over, nil)
	r, w := io.Pipe()
	defer r.Close()
	defer w.Close()
	go func() {
		if err := jpeg.Encode(w, rgba, &jpeg.Options{Quality: 100}); err != nil {
			log.Print(err)
			return
		}
	}()
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}
