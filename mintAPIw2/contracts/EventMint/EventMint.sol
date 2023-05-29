// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@chainlink/contracts/src/v0.8/interfaces/VRFCoordinatorV2Interface.sol";
import "@chainlink/contracts/src/v0.8/VRFConsumerBaseV2.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Shows is ERC721, Ownable, VRFConsumerBaseV2 {
    event ShowCreated(
        uint256 indexed id,
        string name,
        uint256 date,
        string place,
        string[] specialPrices
    );

    event TicketCreated(
        uint256 indexed id,
        uint256 showId,
        string metadata,
        uint256 price
    );

    struct Show {
        string name;
        uint256 date;
        string place;
        string[] specialPrices;
    }

    struct Ticket {
        uint256 showId;
        string metadata;
        uint256 price;
    }

    uint256 public nextTicketId;
    uint256 public nextShowId;
    mapping(uint256 => Show) public shows;
    mapping(uint256 => Ticket) public tickets;
    VRFCoordinatorV2Interface private immutable i_vrfCoordinator;

    constructor(
        address vrfCoordinator
    ) VRFConsumerBaseV2(vrfCoordinator) ERC721("EventTicket", "ETKT") {
        nextShowId = 1;
        i_vrfCoordinator = VRFCoordinatorV2Interface(vrfCoordinator);
    }

    function createShow(
        string memory name,
        uint256 date,
        string memory place,
        string[] memory specialPrices
    ) public onlyOwner returns (uint256 showId) {
        showId = nextShowId++;
        shows[showId] = Show(name, date, place, specialPrices);
        emit ShowCreated(showId, name, date, place, specialPrices);
        return showId;
    }

    function createTicket(
        uint256 showId,
        string memory metadata,
        uint256 price
    ) public onlyOwner returns (uint256 ticketId) {
        ticketId = nextTicketId++;
        tickets[ticketId] = Ticket(showId, metadata, price);
        emit TicketCreated(ticketId, showId, metadata, price);
        return ticketId;
    }

    function getShow(
        uint256 showId
    )
        public
        view
        returns (
            string memory name,
            uint256 date,
            string memory place,
            string[] memory specialPrices
        )
    {
        return (
            shows[showId].name,
            shows[showId].date,
            shows[showId].place,
            shows[showId].specialPrices
        );
    }

    function getTicket(
        uint256 ticketId
    )
        public
        view
        returns (uint256 showId, string memory metadata, uint256 price)
    {
        return (
            tickets[ticketId].showId,
            tickets[ticketId].metadata,
            tickets[ticketId].price
        );
    }

    function _mint(address to, uint256 tokenId) internal override {
        super._mint(to, tokenId);
    }

    function fulfillRandomWords(
        uint256 requestId,
        uint256[] memory randomWords
    ) internal override {
        
    }
}
