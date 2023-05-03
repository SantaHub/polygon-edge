package server

import (
	"github.com/SantaHub/polygon-edge/chain"
	"github.com/SantaHub/polygon-edge/consensus"
	consensusDev "github.com/SantaHub/polygon-edge/consensus/dev"
	consensusDummy "github.com/SantaHub/polygon-edge/consensus/dummy"
	consensusIBFT "github.com/SantaHub/polygon-edge/consensus/ibft"
	consensusPolyBFT "github.com/SantaHub/polygon-edge/consensus/polybft"
	"github.com/SantaHub/polygon-edge/secrets"
	"github.com/SantaHub/polygon-edge/secrets/awsssm"
	"github.com/SantaHub/polygon-edge/secrets/gcpssm"
	"github.com/SantaHub/polygon-edge/secrets/hashicorpvault"
	"github.com/SantaHub/polygon-edge/secrets/local"
	"github.com/SantaHub/polygon-edge/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
