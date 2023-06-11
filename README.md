# Meent


![Architecture](https://raw.githubusercontent.com/ManuelLecaro/Meent/main/docs/architecture.jpg)

This repository contains the API and smart contract components of Meent, a platform that helps web2 ticket portals to adapt to web3 technologies to provide a secure and transparent ticketing ecosystem for live events.

## Architecture

The Meent platform is built on a robust architecture that ensures seamless integration and enhanced security. The following components make up the system:

- **API**: The API, developed in Golang, serves as the bridge between web2 ticketing platforms and the blockchain. It enables event organizers to effortlessly sell tickets and integrates with the smart contract for secure transaction handling.

- **Smart Contract**: The smart contract, developed using Truffle, manages the transfer of ticket values and leverages the power of blockchain for enhanced security and transparency.

## Prerequisites

Before running the API and deploying the smart contract, ensure that you have the following prerequisites installed:

- Golang
- Truffle

## Getting Started

To get started, follow these steps:

1. Clone this repository: `git clone https://github.com/ManuelLecaro/Meent.git`
2. Navigate to the project directory: `cd Meent`
3. go to truffle project `cd ticketDapp`
4. Install the necessary dependencies: `npm install`
5. Deploy the smart contract: `truffle migrate`
6. go to api project `cd mintAPIw2`
7. Fill the configurations on config.toml
8. Start the API server: `docker-compose up`

## API Endpoints

The API provides the following endpoints:

- `/event`: Retrieves a list of events
- `/event/:id`: Retrieves a specific event by ID
- `/event/:id/mint`: Mint a ticket event by ID

For detailed information on the API endpoints and request/response formats, please refer to the API documentation.

## Frontend

You can download the frontend from this repo: ` git clone https://github.com/daniel0ar/MeentFrontend`
Follow its readme instructions and deploy it on local

## Contributing

We welcome contributions from the community to improve and enhance the Meent platform. If you'd like to contribute, please follow our [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).

