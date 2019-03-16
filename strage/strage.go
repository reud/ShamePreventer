package strage

import (
	"bytes"
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

func Put(bucket, path string, data []byte) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	w := client.Bucket(bucket).Object(path).NewWriter(ctx)

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

func Get(bucket, path string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	r, err := client.Bucket(bucket).Object(path).NewReader(ctx)
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
