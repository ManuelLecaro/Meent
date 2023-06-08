// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "../node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "../node_modules/@openzeppelin/contracts/utils/Strings.sol";
import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

contract EventTicket is ERC721, Ownable {
    using Strings for uint256;

    struct Event {
        uint256 eventId;
        string eventName;
        uint256 price;
        uint256 totalTickets;
        uint256 ticketsSold;
        uint256 totalRewards;
    }

    uint256 private nextEventId = 1;
    string private _baseURIextended; // base URI for metadata
    mapping(uint256 => Event) public events;
    mapping(uint256 => address) public eventOwners;
    mapping (uint256 => string) private _tokenURIs;
    
    constructor() ERC721("EventTicket", "ETKT") {}


    function createEvent(
        string memory eventName,
        uint256 price,
        uint256 totalTickets,
        address realOwner,
        uint256 totalRewards
    ) external onlyOwner returns (uint256) {
        events[nextEventId] = Event(
            nextEventId,
            eventName,
            price,
            totalTickets,
            0,
            totalRewards
        );
        eventOwners[nextEventId] = realOwner;
        uint256 currentEventId = nextEventId;
        nextEventId++;
        return currentEventId;
    }

    function updateEvent(
        uint256 eventId,
        string memory eventName,
        uint256 price,
        uint256 totalTickets
    ) external onlyOwner {
        require(events[eventId].eventId != 0, "Event not found");
        events[eventId].eventName = eventName;
        events[eventId].price = price;
        events[eventId].totalTickets = totalTickets;
    }

    function buyTicket(uint256 eventId) external payable {
        require(events[eventId].eventId != 0, "Event not found");
        require(msg.value >= events[eventId].price, "Insufficient funds");
        require(
            events[eventId].ticketsSold < events[eventId].totalTickets,
            "No more tickets available"
        );
        // agregar asignacion de reward aleatorio a ticket con Chainlink

        // llamar a la API que crea el JSON y la imagen del QR. El API debe crear el JSON de metadata e incluir el url de la imagen del QR.

        // solo necesito aquí el URI de la metadata JSON

        uint256 ticketId = eventId * 1000000 + events[eventId].ticketsSold;
        _safeMint(msg.sender, ticketId);
        events[eventId].ticketsSold++;

        if (msg.value > events[eventId].price) {
            payable(msg.sender).transfer(msg.value - events[eventId].price);
        }
    }

    function refundTicket(
        uint256 eventId,
        address buyer,
        uint256 ticketId
    ) external onlyOwner {
        require(events[eventId].eventId != 0, "Event not found");
        require(ownerOf(ticketId) == buyer, "Buyer does not own the ticket");

        _burn(ticketId);
        events[eventId].ticketsSold--;
        payable(buyer).transfer(events[eventId].price);
    }

    function transferEventOwnership(
        uint256 eventId,
        address newOwner
    ) external onlyOwner {
        require(events[eventId].eventId != 0, "Event not found");
        eventOwners[eventId] = newOwner;
    }

    function tokenURI(
        uint256 tokenId
    ) public view virtual override returns (string memory) {
        require(
            _exists(tokenId),
            "ERC721Metadata: URI query for nonexistent token"
        );

        string memory baseURI = _baseURI();
        return
            bytes(baseURI).length > 0
                ? string(abi.encodePacked(baseURI, tokenId.toString()))
                : tokenId.toString();
    }

    function setBaseURI(string memory baseURI_) external onlyOwner() {
        _baseURIextended = baseURI_;
    }

    function _setTokenURI(uint256 tokenId, string memory _tokenURI) external{
        require(_exists(tokenId), "ERC721Metadata: URI set of nonexistent token");
        _tokenURIs[tokenId] = _tokenURI;
    }
    
    function _baseURI() internal view virtual override returns (string memory) {
        return _baseURIextended;
    }
}
