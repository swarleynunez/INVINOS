const Traceability = artifacts.require("Traceability");
const EntityInfo = artifacts.require("EntityInfo");

contract("Traceability", _ => {

    // ID:1
    // ID:2 ID:3
    //      ID:4 ID:5
    //           ID:6 ID:7

    it("Product entry (1)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        await entityInfoInstance.createProductType("UVA-0", "{JSON}");
        await entityInfoInstance.createCompany("COM-0", "{JSON}");
        await entityInfoInstance.createContainer("CTR-0", "{JSON}");
        await traceabilityInstance.productEntry(999, "UVA-0", "COM-0", "CTR-0", "{JSON}");

        const product = await traceabilityInstance.getProduct(1);
        assert.equal(product.quantity, 999);
        assert.lengthOf(product.tv.transitions, 1);
        const transition = product.tv.transitions[0];
        assert.equal(transition.typeID, 0);
        assert.equal(transition.productTypeID, "UVA-0");
        assert.equal(transition.companyID, "COM-0");
        assert.equal(transition.containerID, "CTR-0");
        assert.equal(transition.info, "{JSON}");
    });

    it("Product processing/movement (1)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        await entityInfoInstance.createProductType("UVA-1", "{JSON}");
        await entityInfoInstance.createCompany("COM-1", "{JSON}");
        await entityInfoInstance.createContainer("CTR-1", "{JSON}");
        await traceabilityInstance.productProcessing(1, "UVA-1", "COM-1", "CTR-1", "{JSON}");

        const product = await traceabilityInstance.getProduct(1);
        assert.lengthOf(product.tv.transitions, 2);
        const transition = product.tv.transitions[1];
        assert.equal(transition.typeID, 1);
        assert.equal(transition.productTypeID, "UVA-1");
        assert.equal(transition.companyID, "COM-1");
        assert.equal(transition.containerID, "CTR-1");
        assert.equal(transition.info, "{JSON}");
    });

    it("Product partition (1)", async () => {
        const traceabilityInstance = await Traceability.deployed();

        await traceabilityInstance.productPartition(1, 666, "UVA-1", "COM-1", "CTR-1", "{JSON}");

        const p1 = await traceabilityInstance.getProduct(2);
        assert.equal(p1.quantity, 333);
        assert.lengthOf(p1.tv.transitions, 1);
        assert.equal(p1.tv.prevProductID, 1);
        const p2 = await traceabilityInstance.getProduct(3);
        assert.equal(p2.quantity, 666);
        assert.lengthOf(p2.tv.transitions, 1);
        assert.equal(p1.tv.prevProductID, 1);
        assert.equal(await traceabilityInstance.isProductAvailable(1), false);
        assert.equal(await traceabilityInstance.isProductAvailable(2), true);
        assert.equal(await traceabilityInstance.isProductAvailable(3), true);
    });

    it("Product partition (3)", async () => {
        const traceabilityInstance = await Traceability.deployed();

        await traceabilityInstance.productPartition(3, 333, "UVA-0", "COM-0", "CTR-0", "{JSON}");

        const p1 = await traceabilityInstance.getProduct(4);
        const t1 = p1.tv.transitions[0];
        assert.equal(t1.typeID, 2);
        assert.equal(t1.productTypeID, "UVA-1");
        assert.equal(t1.companyID, "COM-1");
        assert.equal(t1.containerID, "CTR-1");
        assert.equal(t1.info, "{JSON}");
        const p2 = await traceabilityInstance.getProduct(5);
        const t2 = p2.tv.transitions[0];
        assert.equal(t2.typeID, 2);
        assert.equal(t2.productTypeID, "UVA-0");
        assert.equal(t2.companyID, "COM-0");
        assert.equal(t2.containerID, "CTR-0");
        assert.equal(t2.info, "{JSON}");
        assert.equal(await traceabilityInstance.isProductAvailable(3), false);
        assert.equal(await traceabilityInstance.isProductAvailable(4), true);
        assert.equal(await traceabilityInstance.isProductAvailable(5), true);
    });

    it("Product processing/movement and partition (5)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        // Processing/movement
        await traceabilityInstance.productProcessing(5, "UVA-1", "COM-1", "CTR-1", "{JSON}");
        await entityInfoInstance.createProductType("UVA-2", "{JSON}");
        await entityInfoInstance.createCompany("COM-2", "{JSON}");
        await entityInfoInstance.createContainer("CTR-2", "{JSON}");
        await traceabilityInstance.productProcessing(5, "UVA-2", "COM-2", "CTR-2", "{JSON}");

        let p0 = await traceabilityInstance.getProduct(5);
        assert.lengthOf(p0.tv.transitions, 3);

        // Partition
        await traceabilityInstance.productPartition(5, 111, "UVA-0", "COM-0", "CTR-0", "{JSON}");

        p0 = await traceabilityInstance.getProduct(5);
        assert.lengthOf(p0.tv.transitions, 3);
        assert.equal(p0.partitioned, true);
    });

    it("Get traceability (7)", async () => {
        const traceabilityInstance = await Traceability.deployed();
    });
});
