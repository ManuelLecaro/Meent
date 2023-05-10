const EventTicket = artifacts.require('EventTicket');
const truffleAssert = require('truffle-assertions');

contract('EventTicket', (accounts) => {
    let contract;
    const owner = accounts[0];
    const buyer = accounts[1];
    const price = web3.utils.toWei('1', 'ether');
    const totalTickets = 10;

    beforeEach(async () => {
        contract = await EventTicket.new({ from: owner });
    });

    it('should be able to create an event', async () => {
        const eventName = 'Concert';
        await contract.createEvent(eventName, price, totalTickets, { from: owner });

        const event = await contract.events(1);
        assert.equal(event.eventName, eventName);
        assert.equal(event.price, price);
        assert.equal(event.totalTickets, totalTickets);
    });

    it('should be able to buy a ticket', async () => {
        const eventName = 'Concert';
        await contract.createEvent(eventName, price, totalTickets, { from: owner });

        await contract.buyTicket(1, { from: buyer, value: price });

        const event = await contract.events(1);
        assert.equal(event.ticketsSold, 1);

        const ownerOfTicket = await contract.ownerOf(1000000);
        assert.equal(ownerOfTicket, buyer);
    });

    it('should be able to refund a ticket', async () => {
        const eventName = 'Concert';
        await contract.createEvent(eventName, price, totalTickets, { from: owner });
        await contract.buyTicket(1, { from: buyer, value: price });

        await contract.refundTicket(1, buyer, 1000000, { from: owner });

        const event = await contract.events(1);
        assert.equal(event.ticketsSold, 0);

        await truffleAssert.reverts(contract.ownerOf(1000000));
    });
});
