package main

import (
	"fmt"
	"github.com/drillbits/go-amesh/amesh"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

func getTopographyMap(client *amesh.Client) image.Image {
	resp, err := client.TopographyMap.Get()
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
	return resp.Image
}

func getPlaceNameMap(client *amesh.Client) image.Image {
	resp, err := client.PlaceNameMap.Get()
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
	return resp.Image
}

func getMesh(client *amesh.Client) image.Image {
	resp, err := client.Mesh.Get(&amesh.MeshGetOptions{})
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}
	return resp.Image
}

func writeJPGMap(img image.Image, name string) {
	f, _ := os.Create(name)
	defer f.Close()
	jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
}

func writePNGMap(img image.Image, name string) {
	f, _ := os.Create(name)
	defer f.Close()
	png.Encode(f, img)
}

func writeGIFMap(img image.Image, name string) {
	f, _ := os.Create(name)
	defer f.Close()
	gif.Encode(f, img, &gif.Options{})
}

func conv2PNG(img image.Image) {
	tmp, _ := ioutil.TempFile(os.TempDir(), "convtmp")
	defer tmp.Close()
	png.Encode(tmp, img)
}

func main() {
	client := amesh.NewClient(nil)

	mapImage := getTopographyMap(client)
	// writeJPGMap(mapImage, "map.jpg")

	mskImage := getPlaceNameMap(client)
	// writePNGMap(mskImage, "mask.png")

	meshImage := getMesh(client)
	// writeGIFMap(meshImage, "mesh.gif")

	conv2PNG(mapImage)
	conv2PNG(meshImage)

	area := amesh.GetAreaOrWhole(4) // 新宿
	dst := image.NewRGBA(image.Rect(0, 0, area.Width, area.Height))
	bounds := dst.Bounds()
	p := image.Pt(area.Left, area.Top)
	draw.Draw(dst, bounds, mapImage, p, draw.Src)
	draw.Draw(dst, bounds, meshImage, p, draw.Over)
	draw.Draw(dst, bounds, mskImage, p, draw.Over)

	ameshMap, _ := os.Create("amesh.png")
	defer ameshMap.Close()
	png.Encode(ameshMap, dst)
}
