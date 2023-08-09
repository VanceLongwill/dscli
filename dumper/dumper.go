package dumper

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Encoder interface {
	Encode(v any) error
}

type Dumper struct {
	encoder Encoder
	client  *datastore.Client

	batchSize  int
	entityName string
}

type Config struct {
	Client     *datastore.Client
	Encoder    Encoder
	BatchSize  int
	EntityName string
}

func New(cfg Config) *Dumper {
	return &Dumper{
		encoder:    cfg.Encoder,
		client:     cfg.Client,
		batchSize:  cfg.BatchSize,
		entityName: cfg.EntityName,
	}
}

func (d *Dumper) Dump(ctx context.Context) error {
	var offset int
	for {
		q := datastore.NewQuery(d.entityName).
			Limit(d.batchSize).
			Offset(offset)
		var requests []Entity
		keys, err := d.client.GetAll(ctx, q, &requests)
		if err != nil {
			return err
		}
		for _, r := range requests {
			err := d.encoder.Encode(r)
			if err != nil {
				return err
			}
		}
		if len(keys) < d.batchSize {
			break
		}
		offset += d.batchSize
	}
	return nil
}
