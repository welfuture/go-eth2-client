// Copyright © 2023 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

// VersionedBlockRequest contains a versioned signed beacon block request.
type VersionedBlockRequest struct {
	Version   spec.DataVersion
	Bellatrix *bellatrix.SignedBeaconBlock
	Capella   *capella.SignedBeaconBlock
	Deneb     *deneb.SignedBeaconBlock
	Electra   *electra.SignedBeaconBlock
}

// Slot returns the slot of the signed beacon block.
func (v *VersionedBlockRequest) Slot() (phase0.Slot, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil {
			return 0, ErrDataMissing
		}

		return v.Bellatrix.Message.Slot, nil
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil {
			return 0, ErrDataMissing
		}

		return v.Capella.Message.Slot, nil
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil {
			return 0, ErrDataMissing
		}

		return v.Deneb.Message.Slot, nil
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil {
			return 0, ErrDataMissing
		}

		return v.Electra.Message.Slot, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// ExecutionBlockHash returns the block hash of the beacon block.
func (v *VersionedBlockRequest) ExecutionBlockHash() (phase0.Hash32, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil ||
			v.Bellatrix.Message.Body == nil ||
			v.Bellatrix.Message.Body.ExecutionPayload == nil {
			return phase0.Hash32{}, ErrDataMissing
		}

		return v.Bellatrix.Message.Body.ExecutionPayload.BlockHash, nil
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil ||
			v.Capella.Message.Body == nil ||
			v.Capella.Message.Body.ExecutionPayload == nil {
			return phase0.Hash32{}, ErrDataMissing
		}

		return v.Capella.Message.Body.ExecutionPayload.BlockHash, nil
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil ||
			v.Deneb.Message.Body == nil ||
			v.Deneb.Message.Body.ExecutionPayload == nil {
			return phase0.Hash32{}, ErrDataMissing
		}

		return v.Deneb.Message.Body.ExecutionPayload.BlockHash, nil
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil ||
			v.Electra.Message.Body == nil ||
			v.Electra.Message.Body.ExecutionPayload == nil {
			return phase0.Hash32{}, ErrDataMissing
		}

		return v.Electra.Message.Body.ExecutionPayload.BlockHash, nil
	default:
		return phase0.Hash32{}, ErrUnsupportedVersion
	}
}

// Attestations returns the attestations of the beacon block.
func (v *VersionedBlockRequest) Attestations() ([]spec.VersionedAttestation, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil ||
			v.Bellatrix.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttestations := make([]spec.VersionedAttestation, len(v.Bellatrix.Message.Body.Attestations))
		for i, attestation := range v.Bellatrix.Message.Body.Attestations {
			versionedAttestations[i] = spec.VersionedAttestation{
				Version:   spec.DataVersionBellatrix,
				Bellatrix: attestation,
			}
		}

		return versionedAttestations, nil
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil ||
			v.Capella.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttestations := make([]spec.VersionedAttestation, len(v.Capella.Message.Body.Attestations))
		for i, attestation := range v.Capella.Message.Body.Attestations {
			versionedAttestations[i] = spec.VersionedAttestation{
				Version: spec.DataVersionCapella,
				Capella: attestation,
			}
		}

		return versionedAttestations, nil
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil ||
			v.Deneb.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttestations := make([]spec.VersionedAttestation, len(v.Deneb.Message.Body.Attestations))
		for i, attestation := range v.Deneb.Message.Body.Attestations {
			versionedAttestations[i] = spec.VersionedAttestation{
				Version: spec.DataVersionDeneb,
				Deneb:   attestation,
			}
		}

		return versionedAttestations, nil
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil ||
			v.Electra.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttestations := make([]spec.VersionedAttestation, len(v.Electra.Message.Body.Attestations))
		for i, attestation := range v.Electra.Message.Body.Attestations {
			versionedAttestations[i] = spec.VersionedAttestation{
				Version: spec.DataVersionElectra,
				Electra: attestation,
			}
		}

		return versionedAttestations, nil
	default:
		return nil, ErrUnsupportedVersion
	}
}

// Root returns the root of the beacon block.
func (v *VersionedBlockRequest) Root() (phase0.Root, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Bellatrix.Message.HashTreeRoot()
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Capella.Message.HashTreeRoot()
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Deneb.Message.HashTreeRoot()
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Electra.Message.HashTreeRoot()
	default:
		return phase0.Root{}, ErrUnsupportedVersion
	}
}

// BodyRoot returns the body root of the beacon block.
func (v *VersionedBlockRequest) BodyRoot() (phase0.Root, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil ||
			v.Bellatrix.Message.Body == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Bellatrix.Message.Body.HashTreeRoot()
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil ||
			v.Capella.Message.Body == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Capella.Message.Body.HashTreeRoot()
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil ||
			v.Deneb.Message.Body == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Deneb.Message.Body.HashTreeRoot()
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil ||
			v.Electra.Message.Body == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Electra.Message.Body.HashTreeRoot()
	default:
		return phase0.Root{}, ErrUnsupportedVersion
	}
}

// ParentRoot returns the parent root of the beacon block.
func (v *VersionedBlockRequest) ParentRoot() (phase0.Root, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Bellatrix.Message.ParentRoot, nil
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Capella.Message.ParentRoot, nil
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Deneb.Message.ParentRoot, nil
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Electra.Message.ParentRoot, nil
	default:
		return phase0.Root{}, ErrUnsupportedVersion
	}
}

// StateRoot returns the state root of the beacon block.
func (v *VersionedBlockRequest) StateRoot() (phase0.Root, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Bellatrix.Message.StateRoot, nil
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Capella.Message.StateRoot, nil
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Deneb.Message.StateRoot, nil
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil {
			return phase0.Root{}, ErrDataMissing
		}

		return v.Electra.Message.StateRoot, nil
	default:
		return phase0.Root{}, ErrUnsupportedVersion
	}
}

// AttesterSlashings returns the attester slashings of the beacon block.
func (v *VersionedBlockRequest) AttesterSlashings() ([]spec.VersionedAttesterSlashing, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil ||
			v.Bellatrix.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttesterSlashings := make([]spec.VersionedAttesterSlashing, len(v.Bellatrix.Message.Body.AttesterSlashings))
		for i, attesterSlashing := range v.Bellatrix.Message.Body.AttesterSlashings {
			versionedAttesterSlashings[i] = spec.VersionedAttesterSlashing{
				Version:   spec.DataVersionBellatrix,
				Bellatrix: attesterSlashing,
			}
		}

		return versionedAttesterSlashings, nil
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil ||
			v.Capella.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttesterSlashings := make([]spec.VersionedAttesterSlashing, len(v.Capella.Message.Body.AttesterSlashings))
		for i, attesterSlashing := range v.Capella.Message.Body.AttesterSlashings {
			versionedAttesterSlashings[i] = spec.VersionedAttesterSlashing{
				Version: spec.DataVersionCapella,
				Capella: attesterSlashing,
			}
		}

		return versionedAttesterSlashings, nil
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil ||
			v.Deneb.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttesterSlashings := make([]spec.VersionedAttesterSlashing, len(v.Deneb.Message.Body.AttesterSlashings))
		for i, attesterSlashing := range v.Deneb.Message.Body.AttesterSlashings {
			versionedAttesterSlashings[i] = spec.VersionedAttesterSlashing{
				Version: spec.DataVersionDeneb,
				Deneb:   attesterSlashing,
			}
		}

		return versionedAttesterSlashings, nil
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil ||
			v.Electra.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttesterSlashings := make([]spec.VersionedAttesterSlashing, len(v.Electra.Message.Body.AttesterSlashings))
		for i, attesterSlashing := range v.Electra.Message.Body.AttesterSlashings {
			versionedAttesterSlashings[i] = spec.VersionedAttesterSlashing{
				Version: spec.DataVersionElectra,
				Electra: attesterSlashing,
			}
		}

		return versionedAttesterSlashings, nil
	default:
		return nil, ErrUnsupportedVersion
	}
}

// ProposerSlashings returns the proposer slashings of the beacon block.
func (v *VersionedBlockRequest) ProposerSlashings() ([]*phase0.ProposerSlashing, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil ||
			v.Bellatrix.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Bellatrix.Message.Body.ProposerSlashings, nil
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil ||
			v.Capella.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Capella.Message.Body.ProposerSlashings, nil
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil ||
			v.Deneb.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Deneb.Message.Body.ProposerSlashings, nil
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil ||
			v.Electra.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Electra.Message.Body.ProposerSlashings, nil
	default:
		return nil, ErrUnsupportedVersion
	}
}

// SyncAggregate returns the sync aggregate of the beacon block.
func (v *VersionedBlockRequest) SyncAggregate() (*altair.SyncAggregate, error) {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil ||
			v.Bellatrix.Message == nil ||
			v.Bellatrix.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Bellatrix.Message.Body.SyncAggregate, nil
	case spec.DataVersionCapella:
		if v.Capella == nil ||
			v.Capella.Message == nil ||
			v.Capella.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Capella.Message.Body.SyncAggregate, nil
	case spec.DataVersionDeneb:
		if v.Deneb == nil ||
			v.Deneb.Message == nil ||
			v.Deneb.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Deneb.Message.Body.SyncAggregate, nil
	case spec.DataVersionElectra:
		if v.Electra == nil ||
			v.Electra.Message == nil ||
			v.Electra.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Electra.Message.Body.SyncAggregate, nil
	default:
		return nil, ErrUnsupportedVersion
	}
}

// String returns a string version of the structure.
func (v *VersionedBlockRequest) String() string {
	switch v.Version {
	case spec.DataVersionBellatrix:
		if v.Bellatrix == nil {
			return ""
		}

		return v.Bellatrix.String()
	case spec.DataVersionCapella:
		if v.Capella == nil {
			return ""
		}

		return v.Capella.String()
	case spec.DataVersionDeneb:
		if v.Deneb == nil {
			return ""
		}

		return v.Deneb.String()
	case spec.DataVersionElectra:
		if v.Electra == nil {
			return ""
		}

		return v.Electra.String()
	default:
		return "unsupported version"
	}
}
