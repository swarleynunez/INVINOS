var TraceabilityContract = artifacts.require("Traceability");

module.exports = async function (deployer) {
    await deployer.deploy(TraceabilityContract);
};
