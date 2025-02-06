const truffleAssert = require('truffle-assertions');
const common = require("./common.js");

// Contracts
const Auth = artifacts.require("Auth");
const EntityInfo = artifacts.require("EntityInfo");
const Traceability = artifacts.require("Traceability");

contract("Traceability (PRUEBAS FUNCIONALES)", _ => {

    let tinst;

    // Mocha hook
    beforeEach("Init", async () => {
        tinst = await Traceability.new();
    });

    it("E5.6-1: Validar empresa (EXISTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "EMPRESA1";
        console.log("ID empresa:", id);

        await einst.createCompany(id, "{JSON}");
        assert.equal(await einst.existCompany(id), true);
    });

    it("E5.6-1: Validar empresa (NO EXISTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "ABCDEFG";
        console.log("ID empresa:", id);

        assert.equal(await einst.existCompany(id), false);
    });

    it("E5.6-2: Añadir empresa (NO CREADA)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "EMPRESA1";
        console.log("ID empresa:", id);

        assert.equal(await einst.existCompany(id), false);
        await einst.createCompany(id, "{JSON}");
    });

    it("E5.6-2: Añadir empresa (SIN DATOS)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "";
        const error = "Empty company data";
        console.log("ID empresa:", id);
        console.log("Error esperado:", error);

        assert.equal(await einst.existCompany(id), false);
        await truffleAssert.fails(einst.createCompany(id, ""), truffleAssert.ErrorType.REVERT, error);
        assert.equal(await einst.existCompany(id), false);
    });

    it("E5.6-2: Añadir empresa (YA CREADA)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "EMPRESA1";
        const error = "Company already created";
        console.log("ID empresa:", id);
        console.log("Error esperado:", error);

        await einst.createCompany(id, "{JSON}");
        await truffleAssert.fails(einst.createCompany(id, "{JSON}"), truffleAssert.ErrorType.REVERT, error);
    });

    it("E5.6-3: Validar contenedor (EXISTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "CONTENEDOR1";
        console.log("ID contenedor:", id);

        await einst.createContainer(id, "{JSON}");
        assert.equal(await einst.existContainer(id), true);
    });

    it("E5.6-3: Validar contenedor (NO EXISTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "ABCDEFG";
        console.log("ID contenedor:", id);

        assert.equal(await einst.existContainer(id), false);
    });

    it("E5.6-4: Añadir contenedor (NO CREADO)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "CONTENEDOR1";
        console.log("ID contenedor:", id);

        assert.equal(await einst.existContainer(id), false);
        await einst.createContainer(id, "{JSON}");
    });

    it("E5.6-4: Añadir contenedor (SIN DATOS)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "";
        const error = "Empty container data";
        console.log("ID contenedor:", id);
        console.log("Error esperado:", error);

        assert.equal(await einst.existContainer(id), false);
        await truffleAssert.fails(einst.createContainer(id, ""), truffleAssert.ErrorType.REVERT, error);
        assert.equal(await einst.existContainer(id), false);
    });

    it("E5.6-4: Añadir contenedor (YA CREADO)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "CONTENEDOR1";
        const error = "Container already created";
        console.log("ID contenedor:", id);
        console.log("Error esperado:", error);

        await einst.createContainer(id, "{JSON}");
        await truffleAssert.fails(einst.createContainer(id, "{JSON}"), truffleAssert.ErrorType.REVERT, error);
    });

    it("E5.6-5: Validar tipo de producto (EXISTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "GRANEL1";
        console.log("ID tipo de producto:", id);

        await einst.createProductType(id, "{JSON}");
        assert.equal(await einst.existProductType(id), true);
    });

    it("E5.6-5: Validar tipo de producto (NO EXISTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "ABCDEFG";
        console.log("ID tipo de producto:", id);

        assert.equal(await einst.existProductType(id), false);
    });

    it("E5.6-6: Añadir tipo de producto (NO CREADO)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "GRANEL1";
        console.log("ID tipo de producto:", id);

        assert.equal(await einst.existProductType(id), false);
        await einst.createProductType(id, "{JSON}");
    });

    it("E5.6-6: Añadir tipo de producto (SIN DATOS)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "";
        const error = "Empty product type data";
        console.log("ID tipo de producto:", id);
        console.log("Error esperado:", error);

        assert.equal(await einst.existProductType(id), false);
        await truffleAssert.fails(einst.createProductType(id, ""), truffleAssert.ErrorType.REVERT, error);
        assert.equal(await einst.existProductType(id), false);
    });

    it("E5.6-6: Añadir tipo de producto (YA CREADO)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = "GRANEL1";
        const error = "Product type already created";
        console.log("ID tipo de producto:", id);
        console.log("Error esperado:", error);

        await einst.createProductType(id, "{JSON}");
        await truffleAssert.fails(einst.createProductType(id, "{JSON}"), truffleAssert.ErrorType.REVERT, error);
    });

    it("E5.6-7: Añadir transición (DATOS CORRECTOS)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const quantity = 1000;
        const idPT = "GRANEL1";
        const idCom = "EMPRESA1";
        const idCtr = "CONTENEDOR1";
        console.log("Cantidad de producto:", quantity);
        console.log("ID tipo de producto:", idPT);
        console.log("ID empresa:", idCom);
        console.log("ID contenedor:", idCtr);

        await einst.createProductType(idPT, "{JSON}");
        await einst.createCompany(idCom, "{JSON}");
        await einst.createContainer(idCtr, "{JSON}");
        await tinst.productEntry(quantity, idPT, idCom, idCtr, "{JSON}");

        const p = await tinst.products(1);
        assert.lengthOf(p.tv.transitions, 1);
    });

    it("E5.6-7: Añadir transición (DATOS INCORRECTOS)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const quantity = 0;
        const idPT = "ABCDEFG";
        const idCom = "ABCDEFG";
        const idCtr = "ABCDEFG";
        const error1 = "Quantity of 0 not allowed";
        const error2 = "Product type does not exist";
        const error3 = "Company does not exist";
        const error4 = "Container does not exist";
        console.log("Cantidad de producto:", quantity);
        console.log("ID tipo de producto:", idPT);
        console.log("ID empresa:", idCom);
        console.log("ID contenedor:", idCtr);
        console.log("Errores esperados:");
        console.log("   ", error1);
        console.log("   ", error2);
        console.log("   ", error3);
        console.log("   ", error4);

        await truffleAssert.fails(tinst.productEntry(quantity, "", "", "", ""), truffleAssert.ErrorType.REVERT, error1);
        await truffleAssert.fails(tinst.productEntry(1000, idPT, "", "", ""), truffleAssert.ErrorType.REVERT, error2);
        const idPT2 = "GRANEL1";
        await einst.createProductType(idPT2, "{JSON}");
        await truffleAssert.fails(tinst.productEntry(1000, idPT2, idCom, "", ""), truffleAssert.ErrorType.REVERT, error3);
        const idCom2 = "EMPRESA1";
        await einst.createCompany(idCom2, "{JSON}");
        await truffleAssert.fails(tinst.productEntry(1000, idPT2, idCom2, idCtr, ""), truffleAssert.ErrorType.REVERT, error4);
    });

    it("E5.6-8: Obtener transición (EXISTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = 1;
        const quantity = 1000;
        const idPT = "GRANEL1";
        const idCom = "EMPRESA1";
        const idCtr = "CONTENEDOR1";
        console.log("ID producto/transición:", id);
        console.log("Cantidad de producto:", quantity);
        console.log("ID tipo de producto:", idPT);
        console.log("ID empresa:", idCom);
        console.log("ID contenedor:", idCtr);

        await einst.createProductType(idPT, "{JSON}");
        await einst.createCompany(idCom, "{JSON}");
        await einst.createContainer(idCtr, "{JSON}");
        await tinst.productEntry(1000, idPT, idCom, idCtr, "{JSON}");

        const p = await tinst.products(id);
        const pq = await tinst.getProductQuantity(id);
        assert.equal(pq, quantity);
        const t = p.tv.transitions[0];
        assert.equal(t.typeID, 0);
        assert.equal(t.productTypeID, idPT);
        assert.equal(t.companyID, idCom);
        assert.equal(t.containerID, idCtr);
        assert.equal(t.info, "{JSON}");
    });

    it("E5.6-8: Obtener transición (NO EXISTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());

        const id = 999999;
        const quantity = 0;
        const idPT = "ABCDEFG";
        const idCom = "ABCDEFG";
        const idCtr = "ABCDEFG";
        console.log("ID producto/transición:", id);
        console.log("Cantidad de producto:", quantity);
        console.log("ID tipo de producto:", idPT);
        console.log("ID empresa:", idCom);
        console.log("ID contenedor:", idCtr);

        const p = await tinst.products(id);
        assert.lengthOf(p.tv.transitions, 0);
    });

    it("E5.6-9: Obtener trazabilidad (EXISTE LOTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());
        const ainst = await Auth.at(await tinst.authContract());

        const id = 1;
        const quantity = 1000;
        const idPT = "GRANEL1";
        const idCom = "EMPRESA1";
        const idCtr = "CONTENEDOR1";
        const lot = "LT-1234";
        console.log("ID producto:", id);
        console.log("Número de lote:", lot);
        console.log("Transiciones esperadas:");

        await einst.createProductType(idPT, "{JSON}");
        await einst.createCompany(idCom, "{JSON}");
        await einst.createContainer(idCtr, "{JSON}");
        await tinst.productEntry(quantity, idPT, idCom, idCtr, "{JSON}");
        await tinst.productOutput(1, lot, idCom, idCtr, "{JSON}");

        await common.getTraceability(tinst, ainst, lot);

        const p = await tinst.products(id);
        const pq = await tinst.getProductQuantity(id);
        assert.equal(pq, quantity);
        const t1 = p.tv.transitions[0];
        assert.equal(t1.typeID, 0);
        assert.equal(t1.productTypeID, idPT);
        assert.equal(t1.companyID, idCom);
        assert.equal(t1.containerID, idCtr);
        assert.equal(t1.info, "{JSON}");
        const t2 = p.tv.transitions[1];
        assert.equal(t2.typeID, 3);
        assert.equal(t2.productTypeID, idPT);
        assert.equal(t2.companyID, idCom);
        assert.equal(t2.containerID, idCtr);
        assert.equal(t2.info, "{JSON}");
    });

    it("E5.6-9: Obtener trazabilidad (NO EXISTE LOTE)", async () => {
        const einst = await EntityInfo.at(await tinst.entityInfoContract());
        const ainst = await Auth.at(await tinst.authContract());

        const lot = "LT-999999";
        const error = "Unknown lot number";
        console.log("Número de lote:", lot);
        console.log("Error esperado:", error);
        process.stdout.write("Error lanzado: ");
        await common.getTraceability(tinst, ainst, lot);
    });
});
