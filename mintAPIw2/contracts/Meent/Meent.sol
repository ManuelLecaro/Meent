// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";

contract Meent is ERC721, Ownable, ChainlinkClient {
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
    event EventCreated(uint256 eventId);

    // required for Chainlink External API Calls
    using Chainlink for Chainlink.Request;
    uint256 public apiresponse;
    bytes32 private jobId;
    uint256 private fee;
    event RequestTicketCreation(bytes32 indexed requestId, uint256 apiresponse);

    
    constructor() ERC721("Meent", "MEENT") {
        setChainlinkToken(0x779877A7B0D9E8603169DdbD7836e478b4624789);
        setChainlinkOracle(0x6090149792dAAeE9D1D568c9f9a6F6B46AA29eFD);
        jobId = "ca98366cc7314957b8c012c72f05aeeb";
        fee = (1 * LINK_DIVISIBILITY) / 10; // 0,1 LINK on testnet. Must be modified on mainnet
    }


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
        emit EventCreated(currentEventId);
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
        // TODO: Add random number generation for ticket rewards.

        uint256 ticketId = eventId * 1000000 + events[eventId].ticketsSold;
        _safeMint(msg.sender, ticketId);
        
        // Once minted, call to API to create metadata and set TokenURI
        requestTicketMetadataCreation(ticketId);

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

    // Function to retrieve metadata (used by Opensea or Metamask)
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

    function requestTicketMetadataCreation(uint256 ticketId) internal returns (bytes32 requestId) {
        Chainlink.Request memory req = buildChainlinkRequest(
            jobId,
            address(this),
            this.fulfill.selector
        );

        // Set the URL to perform the GET request on
        // TODO: send recently created ticketId to API 
        req.add(
            "get",
            "https://<<APIURL>>/"
        );

        // request.add("path", "RAW.ETH.USD.VOLUME24HOUR"); // Chainlink nodes prior to 1.0.0 support this format
        req.add("path", "RAW,ETH,USD,VOLUME24HOUR"); // Chainlink nodes 1.0.0 and later support this format

        // Multiply the result by 1000000000000000000 to remove decimals
        int256 timesAmount = 10 ** 18;
        req.addInt("times", timesAmount);

        // Sends the request
        return sendChainlinkRequest(req, fee);
    }

    // Receive Response from API (from Chainlink node)
    function fulfill(
        bytes32 _requestId,
        uint256 _volume
    ) public recordChainlinkFulfillment(_requestId) {
        emit RequestTicketCreation(_requestId, _volume);
        apiresponse = _volume;
    }

    function withdrawLink() public onlyOwner {
        LinkTokenInterface link = LinkTokenInterface(chainlinkTokenAddress());
        require(
            link.transfer(msg.sender, link.balanceOf(address(this))),
            "Unable to transfer"
        );
    }
}