# IPFS Clubnet

IPFS Clubnet is a tool for creating and participating in specialized self discovering subgroups of IPFS nodes that have extended functionality.

- Deterministic published tokens called "badges" identify new nodes to the clubnet
- Nodes publish badges and a peer list
- Existing clubnet nodes find new nodes by doing a `dht findprovs` on the badge's hash
- Nodes publish the peers they know about
- Nodes spider other peers peerlist to augment their own peerlist
