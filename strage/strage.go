package strage

import (
	"bytes"
	"cloud.google.com/go/storage"
	"fmt"
	"golang.org/x/net/context"
)

type Worker struct {
	bucketName string
}

func New(bucketName string) Worker {

	return Worker{
		bucketName: bucketName,
	}
}

func (wo Worker) DummyFunc() {
	fmt.Println("Hello World!")
}
func (wo Worker) Put(path string, data []byte) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	w := client.Bucket(wo.bucketName).Object(path).NewWriter(ctx)

	if n, err := w.Write(data); err != nil {
		return err
	} else if n != len(data) {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	if err := client.Close(); err != nil {
		return err
	}
	return nil
}

func (wo Worker) Get(path string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	r, err := client.Bucket(wo.bucketName).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}
	if err := r.Close(); err != nil {
		return nil, err
	}

	if err := client.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
