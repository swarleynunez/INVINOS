const common = require("./common.js");

// Contracts
const Auth = artifacts.require("Auth");
const EntityInfo = artifacts.require("EntityInfo");
const Traceability = artifacts.require("Traceability");

contract("Traceability", _ => {

    /*it("Product entry (1)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        await entityInfoInstance.createProductType("UVA-0", "{JSON}");
        await entityInfoInstance.createCompany("COM-0", "{JSON}");
        await entityInfoInstance.createContainer("CTR-0", "{JSON}");
        await traceabilityInstance.productEntry(999, "UVA-0", "COM-0", "CTR-0", "{JSON}");

        const product = await traceabilityInstance.products(1);
        const productQuantity = await traceabilityInstance.getProductQuantity(1);
        assert.equal(productQuantity, 999);
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
        await traceabilityInstance.productProcessing(1, 0, "UVA-1", "COM-1", "CTR-1", "{JSON}");

        const product = await traceabilityInstance.products(1);
        assert.lengthOf(product.tv.transitions, 2);
        const productQuantity = await traceabilityInstance.getProductQuantity(1);
        assert.equal(productQuantity, 999);
        const transition = product.tv.transitions[1];
        assert.equal(transition.typeID, 1);
        assert.equal(transition.productTypeID, "UVA-1");
        assert.equal(transition.companyID, "COM-1");
        assert.equal(transition.containerID, "CTR-1");
        assert.equal(transition.info, "{JSON}");
    });

    it("Product partition (1)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        await entityInfoInstance.createContainer("CTR-2", "{JSON}");
        await traceabilityInstance.productPartition(1, 666, "COM-1", "CTR-2", "{JSON}");

        const p1 = await traceabilityInstance.products(2);
        assert.equal(p1.quantity, 333);
        assert.lengthOf(p1.tv.transitions, 1);
        assert.equal(p1.tv.prevProductID, 1);
        const p2 = await traceabilityInstance.products(3);
        assert.equal(p2.quantity, 666);
        assert.lengthOf(p2.tv.transitions, 1);
        assert.equal(p1.tv.prevProductID, 1);
        assert.equal(await traceabilityInstance.isProductAvailable(1), false);
        assert.equal(await traceabilityInstance.isProductAvailable(2), true);
        assert.equal(await traceabilityInstance.isProductAvailable(3), true);
    });

    it("Product partition (3)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        await entityInfoInstance.createContainer("CTR-3", "{JSON}");
        await traceabilityInstance.productPartition(3, 333, "COM-1", "CTR-3", "{JSON}");

        const p1 = await traceabilityInstance.products(4);
        const t1 = p1.tv.transitions[0];
        assert.equal(t1.typeID, 2);
        assert.equal(t1.productTypeID, "UVA-1");
        assert.equal(t1.companyID, "COM-1");
        assert.equal(t1.containerID, "CTR-2");
        assert.equal(t1.info, "{JSON}");
        const p2 = await traceabilityInstance.products(5);
        const t2 = p2.tv.transitions[0];
        assert.equal(t2.typeID, 2);
        assert.equal(t2.productTypeID, "UVA-1");
        assert.equal(t2.companyID, "COM-1");
        assert.equal(t2.containerID, "CTR-3");
        assert.equal(t2.info, "{JSON}");
        assert.equal(await traceabilityInstance.isProductAvailable(3), false);
        assert.equal(await traceabilityInstance.isProductAvailable(4), true);
        assert.equal(await traceabilityInstance.isProductAvailable(5), true);
    });

    it("Product processing/movement (5) and partition (5)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        // Processing/movement
        await entityInfoInstance.createProductType("UVA-2", "{JSON}");
        await traceabilityInstance.productProcessing(5, 0, "UVA-2", "COM-1", "CTR-0", "{JSON}");
        await entityInfoInstance.createProductType("UVA-3", "{JSON}");
        await entityInfoInstance.createContainer("CTR-4", "{JSON}");
        await traceabilityInstance.productProcessing(5, 0, "UVA-3", "COM-1", "CTR-4", "{JSON}");

        let p0 = await traceabilityInstance.products(5);
        assert.lengthOf(p0.tv.transitions, 3);

        // Partition
        await entityInfoInstance.createContainer("CTR-5", "{JSON}");
        await traceabilityInstance.productPartition(5, 111, "COM-1", "CTR-5", "{JSON}");

        p0 = await traceabilityInstance.products(5);
        assert.lengthOf(p0.tv.transitions, 3);
        assert.equal(p0.partitioned, true);
    });

    it("Product output (7)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        await entityInfoInstance.createCompany("COM-2", "{JSON}");
        await entityInfoInstance.createContainer("CTR-6", "{JSON}");
        await traceabilityInstance.productOutput(7, "LN-7", "COM-2", "CTR-6", "{JSON}");

        const product = await traceabilityInstance.products(7);
        const transition = product.tv.transitions[1];
        assert.lengthOf(product.tv.transitions, 2);
        assert.equal(product.partitioned, false);
        assert.equal(product.completed, true);
        assert.equal(await traceabilityInstance.lotNumbers("LN-7"), 7);
        assert.equal(transition.typeID, 3);
        assert.equal(transition.productTypeID, "UVA-3");
        assert.equal(transition.companyID, "COM-2");
        assert.equal(transition.containerID, "CTR-6");
        assert.equal(transition.info, "{JSON}");

        // Traceability
        await common.getTraceability(traceabilityInstance, await Auth.deployed(), "LN-7");
    });

    it("Product processings/movements with losses (2)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        // Processing/movement 1
        await entityInfoInstance.createContainer("CTR-7", "{JSON}");
        await traceabilityInstance.productProcessing(2, 111, "UVA-1", "COM-1", "CTR-7", "{JSON}");

        let product = await traceabilityInstance.products(2);
        assert.equal(product.quantity, 333);
        assert.equal(await traceabilityInstance.getProductQuantity(2), 222);

        // Processing/movement 2
        await entityInfoInstance.createContainer("CTR-8", "{JSON}");
        await traceabilityInstance.productProcessing(2, 111, "UVA-2", "COM-1", "CTR-8", "{JSON}");

        product = await traceabilityInstance.products(2);
        assert.equal(product.quantity, 333);
        assert.equal(await traceabilityInstance.getProductQuantity(2), 111);
    });

    it("Product partition (2) and output (9)", async () => {
        const traceabilityInstance = await Traceability.deployed();
        const entityInfoInstance = await EntityInfo.deployed();

        // Partition
        await entityInfoInstance.createContainer("CTR-9", "{JSON}");
        await traceabilityInstance.productPartition(2, 61, "COM-1", "CTR-9", "{JSON}");

        const p1 = await traceabilityInstance.products(8);
        assert.equal(p1.quantity, 50);
        assert.equal(await traceabilityInstance.getProductQuantity(8), 50);

        const p2 = await traceabilityInstance.products(9);
        assert.equal(p2.quantity, 61);
        assert.equal(await traceabilityInstance.getProductQuantity(9), 61);

        // Output
        await entityInfoInstance.createCompany("COM-3", "{JSON}");
        await entityInfoInstance.createContainer("CTR-10", "{JSON}");
        await traceabilityInstance.productOutput(9, "LN-9", "COM-3", "CTR-10", "{JSON}");

        // Traceability
        await common.getTraceability(traceabilityInstance, await Auth.deployed(), "LN-9");
    });*/
});
