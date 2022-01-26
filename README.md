<h1 align="center">Sonr</h1>

<div align="center">
  :trident: :dolphin: :godmode: :trident:
</div>
<div align="center">
  <strong>The Official Sonr project source code</strong>
</div>
<div align="center">
  A <code>easy-to-use</code> framework for building immersive decentralized applications.
</div>
<br />

<div align="center">
  <!-- Stability -->
    <img alt="CodeFactor Grade" src="https://img.shields.io/codefactor/grade/github/sonr-io/sonr-go/master?style=for-the-badge">
  <!-- NPM version -->
  <a href="https://godoc.org/github.com/sonr-io/sonr-go">
  <img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge" />
  </a>
  <!-- Test Coverage -->
  <a href="https://codecov.io/github/sonr-io/sonr-go">
<img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/sonr-io/sonr-go?label=TLOC&style=for-the-badge">
  </a>
  <!-- Downloads -->
<img alt="Twitter Follow" src="https://img.shields.io/twitter/follow/sonrprotocol?color=%2300ACEE&label=🐦 SonrProtocol&style=for-the-badge">
</div>

<div align="center">
  <h3>
    <a href="https://sonr.io">
      Home
    </a>
    <span> | </span>
    <a href="https://discord.gg/tjWMfvQZ7b">
      Discord
    </a>
    <span> | </span>
    <a href="https://github.com/sonr-io/sonr-go/wiki">
      Wiki
    </a>
    <span> | </span>
      <!-- <span> | </span> -->
    <a href="https://github.com/sonr-io/sonr-go/issues">
      Issues
    </a>
  </h3>
</div>

<div align="center">
  <sub>The most comprehensive framework for the DWeb. Built with ❤︎ by the
  <a href="mailto:team@sonr.io">Sonr Team</a> and
  <a href="https://github.com/sonr-io/sonr/graphs/contributors">
    contributors
  </a>
</div>

## Table of Contents

- [About](#about)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Version History](#version-history)
- [Contributors](#contributors)
- [Acknowledgments](#acknowledgments)
- [License](#license)

## About

Sonr is building the most immersive DWeb experience for both Users and Developers alike. We believe the best way to onboard the next billion users is to create a cohesive end-to-end platform that’s composable and interoperable with all existing protocols.

For this we built our Networking layer in [Libp2p](“https://libp2p.io”) and our Layer 1 Blockchain with [Starport](“https://starport.com”). Our network comprises of two separate nodes: [Highway](“https://github.com/sonr-io/highway”) and [Motor](“https://github.com/sonr-io/motor”), which each have a specific use case on the network. In order to maximize the onboarding experience, we developed our own [Wallet](“https://github.com/sonr-io/wallet) which has value out of the gate!

<img src="https://camo.githubusercontent.com/1c3eb2fc698e088b15bec07168ad4e037ac2f5c4469c91a311a1038b5b702966/68747470733a2f2f646f63732e736f6e722e696f2f7e2f66696c65732f76302f622f676974626f6f6b2d782d70726f642e61707073706f742e636f6d2f6f2f73706163657325324638784859417a3845707652436a67336873674d5525324675706c6f616473253246326869324f6c50524b78566b51327a3269454d582532466f70656e67726170682e706e673f616c743d6d6564696126746f6b656e3d35643764383431302d663533632d343462312d383264612d356331316431616237373735"/>

## Getting Started

IPFS has established a critical piece of infrastructure that is used by the majority of DWeb protocols to different degrees. Libp2p is part of the IPFS stack and is used for the underlying peer-to-peer network for discovery, routing and data exchange.

Peers in the network can dial other peers in the network to exchange messages using various transports, like QUIC, TCP, WebSocket, and Bluetooth. Modular design of the libp2p framework enables it to build drivers for other transports. Peers can run on any device, as a cloud service, mobile application or in the browser and talk to each other as long as they are connected through the same libp2p network.

### Dependencies

- [Golang](https://go.dev)
- [Libp2p](https://libp2p.io)
- [Starport](https://starport.com)

### Configuration

This project is a pseudo-monorepo, meaning it has a single root directory and all of its packages are in subdirectories. The structure is as follows:

```text
/app             ->        Exported Starport app
/cmd             ->        CLI commands
/common          ->        Commonly used packages
  └─ device      ->        +   Node Device utilities
  └─ discover    ->        +   Peer discovery
  └─ exchange    ->        +   Exchange data with other nodes
  └─ transmit    ->        +   Direct data transmission
/core            ->        Fundamental components of the network.
  └─ bucket      ->        +   A collection of objects or blobs
  └─ channel     ->        +   Topic pub/sub messaging and Realtime K/V Store
  └─ node        ->        +   Central Interface to manage user's peer on the network
  └─ object      ->        +   A registered schema type on the blockchain
  └─ service     ->        +   Interface to interact with a deployed or authorized service
/did             ->        Sonr DID Method Utility Package
/docs            ->        Documentation.
/highway         ->        The Guardian Node on the network.
/motor           ->        The User Node on the network.
/proto           ->        Cosmos SDK Protocol Definitions
/testutil        ->        Blockchain test utilities.
/types           ->        Sonr Protocol Definitions and buf.build config
/vue             ->        Vue.js frontend for Cosmos SDK
/x               ->        Implementation of Cosmos-Sonr Schemas
```

## Installation

<!-- ### Framework

### Plugins -->

### CLI

To install the latest version of the Sonr blockchain node's binary, execute the following command on your machine:

```shell
curl https://sonr.ws/sonr@latest! | sudo bash
```

### Release

To install the latest version of the Sonr blockchain node's binary, execute the following command on your machine:

```shell
curl https://sonr.ws/sonr@latest! | sudo bash
```

## Usage

To launch the Sonr Blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Start the Blockchain

```shell
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Run the Flutter Frontend

Starport has scaffolded a Flutter-based mobile app in the `flutter` directory. Run the following commands to install dependencies and start the app:

```shell
cd flutter
flutter pub get
flutter run
```

### Run the Vue.js Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```text
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Version History

### v0.0.1

- Implement Sonr Blockchain client into Motor and Highway nodes
- Add Highway Service and Swagger generated clients and documentation
- Create `Registry` for storing `DIDDocument` in the Blockchain

## Contributors

> Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated!**

### Authors

- [Prad Nukala](https://github.com/prnk28)
- [Ian Judd](https://github.com/ikjudd)

### Submitting a PR

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Acknowledgments

Tools, libraries, and frameworks that make the Sonr project possible:

- [Libp2p](https://libp2p.io/)
- [Cosmos](https://www.cosmos.network/)
- [Handshake](https://handshake.org/)

## License

This project facilitated under **Sonr Inc.** is distributed under the **GPLv3 License**. See `LICENSE.md` for more information.
