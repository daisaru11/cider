package docker

import (
	"bytes"
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type clientImpl struct {
	cli *client.Client
}

func NewClientImpl() (Client, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &clientImpl{
		cli: cli,
	}, nil
}

func (c *clientImpl) PullImage(ctx context.Context, image string) error {
	log.Print("docker client pull image start")

	options := types.ImagePullOptions{}

	reader, err := c.cli.ImagePull(ctx, image, options)
	if nil != err {
		log.Println(err)
		return err
	}
	defer reader.Close()

	// if _, err := io.Copy(ioutil.Discard, reader); err != nil {
	// 	return err
	// }

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	s := buf.String()
	log.Println("docker client pull image result: ", s)

	return nil
}
