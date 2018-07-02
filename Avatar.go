package AvatarGenerator

import (
	"os"
	"image/png"
	"image"
	"fmt"
	"crypto/sha256"
	"image/color"
)
const (
	IMAGE_DIMENSION = 256
)


func GenerateAvatar(email string, ip string, user string)  {
	if email == "" || ip == "" || user == "" {
		print("You must provide an email, ip address and username \n")
		os.Exit(1)
	}
	bArray := sha256.Sum256([]byte(email + ip + user))


	img := image.NewNRGBA(image.Rect(0, 0, IMAGE_DIMENSION, IMAGE_DIMENSION))
	bLen := len(bArray)
	loop := IMAGE_DIMENSION / bLen
	for i := 0; i < loop; i++ {
		for j := 0; j < loop; j++ {
			wStart := i * bLen
			hStart := j * bLen
			v := bArray[(i*loop+j)%bLen]
			picker := (i*loop + j) % 3
			var r, g, b uint8
			if picker == 0 {
				r = 255
				g = v
				b = 0
			}
			if picker == 1 {
				r = 0
				g = 255
				b = v
			}
			if picker == 2 {
				r = v
				g = 0
				b = 255
			}

			for m := 0; m < bLen; m++ {
				for n := 0; n < bLen; n++ {
					img.Set(wStart+m, hStart+n, color.RGBA{r, g, b, 255})
				}
			}
		}
	}

	f, _ := os.Create("avatarme.png")
	defer f.Close()
	png.Encode(f, img)
	fmt.Println("Successfully saved image.")

}
