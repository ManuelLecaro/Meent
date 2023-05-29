// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0  <0.9.0;

import '@openzeppelin/contracts/token/ERC721/ERC721.sol';
import '@openzeppelin/contracts/utils/Strings.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

contract EventTicket is ERC721, Ownable {
    using Strings for uint256;

    struct Event {
        uint256 eventId;
        string eventName;
        uint256 price;
        uint256 totalTickets;
        uint256 ticketsSold;
    }

    uint256 private nextEventId = 1;
    mapping(uint256 => Event) public events;
    mapping(uint256 => address) public eventOwners;

    constructor() ERC721("EventTicket", "ETKT") {}


    function createEvent(
        string memory eventName,
        uint256 price,
        uint256 totalTickets
    ) external onlyOwner {
        events[nextEventId] = Event(
            nextEventId,
            eventName,
            price,
            totalTickets,
            0
        );
        eventOwners[nextEventId] = msg.sender;
        nextEventId++;
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
                : "";
    }
}
