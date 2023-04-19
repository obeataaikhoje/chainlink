package monitoring

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/riferrei/srclient"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-relay/pkg/monitoring/config"
)

type SchemaRegistry interface {
	// EnsureSchema handles three cases when pushing a schema spec to the SchemaRegistry:
	// 1. when the schema with a given subject does not exist, it will create it.
	// 2. if a schema with the given subject already exists but the spec is different, it will update it and bump the version.
	// 3. if the schema exists and the spec is the same, it will not do anything.
	EnsureSchema(subject, spec string) (Schema, error)
}

type schemaRegistry struct {
	backend srclient.ISchemaRegistryClient
	log     Logger
}

func NewSchemaRegistry(cfg config.SchemaRegistry, log Logger) SchemaRegistry {
	backend := srclient.CreateSchemaRegistryClient(cfg.URL)
	if cfg.Username != "" && cfg.Password != "" {
		backend.SetCredentials(cfg.Username, cfg.Password)
	}
	return &schemaRegistry{backend, log}
}

func (s *schemaRegistry) EnsureSchema(subject, spec string) (Schema, error) {
	existingSchema, err := s.backend.GetLatestSchema(subject)
	if err != nil && !isNotFoundErr(err) {
		return nil, fmt.Errorf("failed to read schema for subject '%s': %w", subject, err)
	}
	if err != nil && isNotFoundErr(err) {
		s.log.Infow("creating new schema", "subject", subject)
		newSchema, schemaErr := s.backend.CreateSchema(subject, spec, srclient.Avro)
		if schemaErr != nil {
			return nil, fmt.Errorf("unable to create new schema with subject '%s': %w", subject, schemaErr)
		}
		return wrapSchema{subject, newSchema}, nil
	}
	isEqualSchemas, errInIsEqualJSON := isEqualJSON(existingSchema.Schema(), spec)
	if errInIsEqualJSON != nil {
		return nil, fmt.Errorf("failed to compare schama in registry with local schema: %w", errInIsEqualJSON)
	}
	if isEqualSchemas {
		s.log.Infow("using existing schema", "subject", subject)
		return wrapSchema{subject, existingSchema}, nil
	}
	s.log.Infow("updating schema", "subject", subject)
	updatedSchema, err := s.backend.CreateSchema(subject, spec, srclient.Avro)
	if err != nil {
		return nil, fmt.Errorf("unable to update schema with subject '%s': %w", subject, err)
	}
	return wrapSchema{subject, updatedSchema}, nil
}

// Helpers

func isNotFoundErr(err error) bool {
	if strings.HasPrefix(err.Error(), "Subject not found") { // for mock schema registry
		return true
	}
	var srErr srclient.Error
	if errors.As(err, &srErr) && srErr.Code == 40401 { // for the actual schema registry api.
		return true
	}
	return false
}

func isEqualJSON(a, b string) (bool, error) {
	var aUntyped, bUntyped interface{}

	if err := json.Unmarshal([]byte(a), &aUntyped); err != nil {
		return false, fmt.Errorf("failed to unmarshal first avro schema: %w", err)
	}
	if err := json.Unmarshal([]byte(b), &bUntyped); err != nil {
		return false, fmt.Errorf("failed to unmarshal second avro schema: %w", err)
	}

	return assert.ObjectsAreEqual(aUntyped, bUntyped), nil
}
