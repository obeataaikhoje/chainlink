package monitoring

import (
	"encoding/json"
	"fmt"

	"github.com/linkedin/goavro/v2"

	"github.com/smartcontractkit/chainlink-relay/pkg/monitoring/avro"
)

// This files contains Avro schemas for encoding message to be published to kafka.
// See https://avro.apache.org/docs/current/spec.html#schemas
// Note that changes to these schemas need to be both forward and backward compatible, ie. FULL_TRANSITIVE.
// See https://docs.confluent.io/platform/current/schema-registry/avro.html

var transmissionAvroSchema = avro.Record("transmission", avro.Opts{Namespace: "link.chain.ocr2"}, avro.Fields{
	avro.Field("block_number", avro.Opts{Doc: "uint64 big endian"}, avro.Bytes),
	avro.Field("block_number_uint64", avro.Opts{Default: avro.NullValue}, avro.Union{
		avro.Null,
		avro.Decimal("transmission_block_number", 32, 78, 0),
	}),
	avro.Field("answer", avro.Opts{}, avro.Record("answer", avro.Opts{}, avro.Fields{
		avro.Field("data", avro.Opts{Doc: "*big.Int"}, avro.Bytes),
		avro.Field("data_uint256", avro.Opts{Default: avro.NullValue, Doc: "string version of data"}, avro.Union{
			avro.Null,
			avro.Decimal("transmission_data", 32, 78, 0),
		}),
		avro.Field("timestamp", avro.Opts{Doc: "uint32"}, avro.Long),
		// These fields are made "optional" for FULL_TRANSITIVE compatibility, but they should be set in all cases.
		avro.Field("config_digest", avro.Opts{Doc: "[32]byte encoded as base64", Default: avro.NullValue}, avro.Union{avro.Null, avro.String}),
		avro.Field("epoch", avro.Opts{Doc: "uint32", Default: avro.NullValue}, avro.Union{avro.Null, avro.Long}),
		avro.Field("round", avro.Opts{Doc: "uint8", Default: avro.NullValue}, avro.Union{avro.Null, avro.Int}),
	})),
	// These field is "optional" for FULL_TRANSITIVE compatibility, but it should be set in all cases.
	avro.Field("chain_config", avro.Opts{Default: avro.NullValue, Doc: "required!"}, avro.Union{
		avro.Null,
		avro.Record("chain_config", avro.Opts{}, avro.Fields{
			avro.Field("network_name", avro.Opts{}, avro.String),
			avro.Field("network_id", avro.Opts{}, avro.String),
			avro.Field("chain_id", avro.Opts{}, avro.String),
		}),
	}),
	avro.Field("solana_chain_config", avro.Opts{Doc: "deprecated! in favour of chain_config"}, avro.Record("solana_chain_config", avro.Opts{}, avro.Fields{
		avro.Field("network_name", avro.Opts{}, avro.String),
		avro.Field("network_id", avro.Opts{}, avro.String),
		avro.Field("chain_id", avro.Opts{}, avro.String),
	})),
	avro.Field("feed_config", avro.Opts{}, avro.Record("feed_config", avro.Opts{}, avro.Fields{
		avro.Field("feed_name", avro.Opts{}, avro.String),
		avro.Field("feed_path", avro.Opts{}, avro.String),
		avro.Field("symbol", avro.Opts{}, avro.String),
		avro.Field("heartbeat_sec", avro.Opts{}, avro.Long),
		avro.Field("contract_type", avro.Opts{}, avro.String),
		avro.Field("contract_status", avro.Opts{}, avro.String),
		avro.Field("contract_address", avro.Opts{Doc: "[32]byte"}, avro.Bytes),
		avro.Field("contract_address_string", avro.Opts{Default: avro.NullValue}, avro.Union{avro.Null, avro.String}),
		// These field is "required" for FULL_TRANSITIVE compatibility, but they are deprecated and they should be set to a zero value.
		avro.Field("transmissions_account", avro.Opts{Doc: "[32]byte deprecated!"}, avro.Bytes),
		avro.Field("state_account", avro.Opts{Doc: "[32]byte deprecated!"}, avro.Bytes),
	})),
	// These field is "optional" for FULL_TRANSITIVE compatibility, but it should be set in all cases.
	avro.Field("link_balance", avro.Opts{Default: avro.NullValue, Doc: "required!"}, avro.Union{
		avro.Null,
		avro.Bytes,
	}),
	avro.Field("link_balance_uint256", avro.Opts{Default: avro.NullValue}, avro.Union{
		avro.Null,
		avro.Decimal("transmission_link_balance", 32, 78, 0),
	}),
})

var configSetSimplifiedAvroSchema = avro.Record("config_set_simplified", avro.Opts{Namespace: "link.chain.ocr2"}, avro.Fields{
	avro.Field("config_digest", avro.Opts{Doc: "[32]byte encoded as base64"}, avro.String),
	avro.Field("block_number", avro.Opts{Doc: "uint64 big endian"}, avro.Bytes),
	avro.Field("block_number_uint64", avro.Opts{Default: avro.NullValue}, avro.Union{
		avro.Null,
		avro.Decimal("config_block_number", 32, 78, 0),
	}),
	avro.Field("signers", avro.Opts{Doc: "json encoded array of base64-encoded signing keys"}, avro.String),
	avro.Field("transmitters", avro.Opts{Doc: "json encoded array of base64-encoded transmission keys"}, avro.String),
	avro.Field("f", avro.Opts{Doc: "uint8"}, avro.Int),
	avro.Field("delta_progress", avro.Opts{Doc: "uint64 big endian"}, avro.Bytes),
	avro.Field("delta_progress_uint64", avro.Opts{Default: avro.NullValue}, avro.Union{
		avro.Null,
		avro.Decimal("config_delta_progress", 32, 78, 0),
	}),
	avro.Field("delta_resend", avro.Opts{Doc: "uint64 big endian"}, avro.Bytes),
	avro.Field("delta_resend_uint64", avro.Opts{Default: avro.NullValue}, avro.Union{
		avro.Null,
		avro.Decimal("config_delta_resend", 32, 78, 0),
	}),
	avro.Field("delta_round", avro.Opts{Doc: "uint64 big endian"}, avro.Bytes),
	avro.Field("delta_round_uint64", avro.Opts{Default: avro.NullValue}, avro.Union{
		avro.Null,
		avro.Decimal("config_delta_round", 32, 78, 0),
	}),
	avro.Field("delta_grace", avro.Opts{Doc: "uint64 big endian"}, avro.Bytes),
	avro.Field("delta_grace_uint64", avro.Opts{Default: avro.NullValue}, avro.Union{
		avro.Null,
		avro.Decimal("config_delta_grace", 32, 78, 0),
	}),
	avro.Field("delta_stage", avro.Opts{Doc: "uint64 big endian"}, avro.Bytes),
	avro.Field("delta_stage_uint64", avro.Opts{Default: avro.NullValue}, avro.Union{
		avro.Null,
		avro.Decimal("config_delta_stage", 32, 78, 0),
	}),
	avro.Field("r_max", avro.Opts{Doc: "uint32"}, avro.Long),
	avro.Field("s", avro.Opts{Doc: "json encoded aray of ints"}, avro.String),
	avro.Field("oracles", avro.Opts{Doc: "json encoded list of oracles"}, avro.String),
	avro.Field("feed_state_account", avro.Opts{Doc: "[32]byte"}, avro.String),
})

var (
	// Avro schemas to sync with the registry
	TransmissionAvroSchema        string
	ConfigSetSimplifiedAvroSchema string

	// These codecs are used in tests
	transmissionCodec        *goavro.Codec
	configSetSimplifiedCodec *goavro.Codec
)

func init() {
	var err error
	var buf []byte

	buf, err = json.Marshal(transmissionAvroSchema)
	if err != nil {
		panic(fmt.Errorf("failed to generate Avro schema for transmission: %w", err))
	}
	TransmissionAvroSchema = string(buf)
	transmissionCodec, err = goavro.NewCodec(TransmissionAvroSchema)
	if err != nil {
		panic(fmt.Errorf("failed to parse Avro schema for the latest transmission: %w", err))
	}

	buf, err = json.Marshal(configSetSimplifiedAvroSchema)
	if err != nil {
		panic(fmt.Errorf("failed to generate Avro schema for configSimplified: %w", err))
	}
	ConfigSetSimplifiedAvroSchema = string(buf)
	configSetSimplifiedCodec, err = goavro.NewCodec(ConfigSetSimplifiedAvroSchema)
	if err != nil {
		panic(fmt.Errorf("failed to parse Avro schema for the latest configSetSimplified: %w", err))
	}

	// These codecs are used in tests but not in main, so the linter complains.
	_ = transmissionCodec
	_ = configSetSimplifiedCodec
}
