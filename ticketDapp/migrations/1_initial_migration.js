const Shows = artifacts.require("Shows");

module.exports = function(deployer) {
  const userAddress = "0x3B1bec8c3e746c7892A995e69ba9CC05F2e843bc";
  deployer.deploy(Shows, userAddress);
};