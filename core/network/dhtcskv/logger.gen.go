// Code generated by berty.tech/core/.scripts/generate-logger.sh

<<<<<<< HEAD:core/api/p2p/logger.gen.go
package p2p
||||||| merged common ancestors:core/network/p2p/protocol/provider/dht/logger.gen.go
package dht
=======
package dhtcskv
>>>>>>> feat(network): replace pubsub provider by dht-cskv (conversation broken on this commit):core/network/dhtcskv/logger.gen.go

import "go.uber.org/zap"

func logger() *zap.Logger {
<<<<<<< HEAD:core/api/p2p/logger.gen.go
	return zap.L().Named("core.api.p2p")
||||||| merged common ancestors:core/network/p2p/protocol/provider/dht/logger.gen.go
	return zap.L().Named("core.network.p2p.protocol.provider.dht")
=======
	return zap.L().Named("core.network.dhtcskv")
>>>>>>> feat(network): replace pubsub provider by dht-cskv (conversation broken on this commit):core/network/dhtcskv/logger.gen.go
}
