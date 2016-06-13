package archive

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type LoginConfig struct {
	Registry   string
	Username   string
	Password   string
	OutputFile string
}

func (c LoginConfig) CreateConfig() []byte {
	loginConfig := map[string]map[string]map[string]string{}
	loginConfig["auths"] = make(map[string]map[string]string)
	loginConfig["auths"]["https://index.docker.io/v1/"] = make(map[string]string)
	authString := fmt.Sprintf("%s:%s", c.Username, c.Password)
	loginConfig["auths"][c.Registry]["auth"] = base64.StdEncoding.EncodeToString([]byte(authString))
	data, _ := json.MarshalIndent(loginConfig, "", "    ")
	return data
}

// Create the archive
func (c LoginConfig) CreateArchive() {
	filedata := c.CreateConfig()
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	var files = []struct {
		Name string
		Body []byte
		Mode int64
		Type byte
	}{
		{".docker", []byte{}, 0755, tar.TypeDir},
		{".docker/config.json", filedata, 0600, tar.TypeReg},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name:     file.Name,
			Mode:     file.Mode,
			Typeflag: file.Type,
			Size:     int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write(file.Body); err != nil {
			log.Fatalln(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	var f *os.File
	var err error
	if c.OutputFile != "-" {
		f, err = os.Create(c.OutputFile)
		defer f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		f = os.Stdout
	}
	w := gzip.NewWriter(f)
	defer w.Close()
	w.Write(buf.Bytes())
}
