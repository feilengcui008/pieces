package main

// demo for calculating the layer id of docker images

import "fmt"
import "crypto/sha256"
import "encoding/hex"
import "github.com/docker/docker/layer"
import "github.com/docker/docker/image"
import "io/ioutil"

func getDigestOfTwoLayer(l string, r string) string {
	hash := sha256.New()
	hash.Write([]byte(l + " " + r))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return "sha256:" + mdStr
}

func main() {
	diffID := []layer.DiffID{
		layer.DiffID("sha256:040ba7b9591ca6c75584e37195149facf6906d98e3597a9b2d4c1e3889aff633"),
		layer.DiffID("sha256:e15f8eeda399274bc7931ca0b2bb3bba68cbf56ac92e9a39e377ba26c876d037"),
		layer.DiffID("sha256:a9ee34f9e4e26a310dd0760d7da1a663ab8df3be7456bbfd8387be5ea54b5cfd"),
		layer.DiffID("sha256:0c291dc95357ed2e19b992b76eb66ebb05b4b068544589b81f2cfcc18d43fcf5"),
		layer.DiffID("sha256:f215f043863effefce8e88153268a12a65413473c5209c69343d11c6a507a54c"),
	}
	origin := getDigestOfTwoLayer(string(diffID[0]), string(diffID[1]))
	fmt.Println(origin)
	for i := 2; i <= 4; i++ {
		origin = getDigestOfTwoLayer(origin, string(diffID[i]))
		fmt.Println(origin)
	}
	fmt.Println("=======")
	// user layer api
	fmt.Println(layer.CreateChainID(diffID[0:1]))
	fmt.Println(layer.CreateChainID(diffID[0:2]))
	fmt.Println(layer.CreateChainID(diffID[0:3]))
	fmt.Println(layer.CreateChainID(diffID[0:4]))
	fmt.Println(layer.CreateChainID(diffID))
	fmt.Println("=======")
	// from file
	data, _ := ioutil.ReadFile("image.json")
	img, _ := image.NewFromJSON(data)
	fmt.Println(img.RootFS.DiffIDs)
	fmt.Println(img.RootFS.ChainID())
}
