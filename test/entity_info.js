const EntityInfo = artifacts.require("EntityInfo");

contract("EntityInfo", _ => {

    it("Create product type", async () => {
        const instance = await EntityInfo.deployed();
        await instance.createProductType("UVA-0", "{JSON}");
        assert.equal(await instance.existProductType("UVA-0"), true);
    });

    it("Create company", async () => {
        const instance = await EntityInfo.deployed();
        await instance.createCompany("COM-0", "{JSON}");
        assert.equal(await instance.existCompany("COM-0"), true);
    });

    it("Create container", async () => {
        const instance = await EntityInfo.deployed();
        await instance.createContainer("CTR-0", "{JSON}");
        assert.equal(await instance.existContainer("CTR-0"), true);
    });
});
