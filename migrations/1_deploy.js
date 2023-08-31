var AuthContract = artifacts.require("Auth");
var EntityInfoContract = artifacts.require("EntityInfo");
var TraceabilityContract = artifacts.require("Traceability");

module.exports = async function (deployer) {
  await deployer.deploy(AuthContract);
  await deployer.deploy(EntityInfoContract);
  await deployer.deploy(TraceabilityContract, AuthContract.address, EntityInfoContract.address);
};
