const common = require("./common.js");

// Contracts
const Auth = artifacts.require("Auth");
const EntityInfo = artifacts.require("EntityInfo");
const Traceability = artifacts.require("Traceability");

contract("TraceabilityLines", _ => {

    /*let authInstance;
    let entityInfoInstance;
    let traceabilityInstance;

    // Mocha hook
    beforeEach("Init", async () => {
        authInstance = await Auth.new();
        entityInfoInstance = await EntityInfo.new();
        traceabilityInstance = await Traceability.new(authInstance.address, entityInfoInstance.address);
    });

    it("SAUVIGNON BLANC MOSTO BLANCO GRANEL", async () => {
        // Required entities
        await entityInfoInstance.createProductType("SAUVIGNON_BLANC_UVA", "{JSON}");
        await entityInfoInstance.createProductType("SAUVIGNON_BLANC_MOSTO", "{JSON}");
        await entityInfoInstance.createCompany("EMPRESA_ENTRADA", "{JSON}");
        await entityInfoInstance.createCompany("MI_EMPRESA", "{JSON}");
        await entityInfoInstance.createCompany("EMPRESA_SALIDA", "{JSON}");
        await entityInfoInstance.createContainer("DESCARGADERO_X", "{JSON}");
        await entityInfoInstance.createContainer("A10", "{JSON}");
        await entityInfoInstance.createContainer("CC2", "{JSON}");

        // Transition 1
        let json = {
            "Recepción uvas (Kg)": 9000
        }
        await traceabilityInstance.productEntry(
            7200,
            "SAUVIGNON_BLANC_UVA",
            "EMPRESA_ENTRADA",
            "DESCARGADERO_X",
            JSON.stringify(json)
        );

        // Transition 2 (losses)
        json = {
            "Peso uvas (Kg)": 8500,
            "GAP (%Vol)": 12.4,
            "Densidad (g/ml)": 1.088
        }
        await traceabilityInstance.productProcessing(
            1,
            400,
            "SAUVIGNON_BLANC_UVA",
            "MI_EMPRESA",
            "DESCARGADERO_X",
            JSON.stringify(json)
        );

        // Transition 3
        json = {
            "GAP (%Vol)": 14.2
        }
        await traceabilityInstance.productProcessing(
            1,
            0,
            "SAUVIGNON_BLANC_MOSTO",
            "MI_EMPRESA",
            "A10",
            JSON.stringify(json)
        );

        // Transition 4
        json = {
            "GAP (%Vol)": 14.2
        }
        await traceabilityInstance.productOutput(
            1,
            "LN_1", // Internal lot number
            "EMPRESA_SALIDA",
            "CC2",
            JSON.stringify(json)
        );

        // Traceability
        await common.getTraceability(traceabilityInstance, authInstance, "LN_1");
    });

    it("MOSTO SAUVIGNON BLANC A VINO BLANCO JOVEN GRANEL", async () => {
        // Required entities
        await entityInfoInstance.createProductType("SAUVIGNON_BLANC_MOSTO", "{JSON}");
        await entityInfoInstance.createProductType("SAUVIGNON_BLANC_VINO", "{JSON}");
        await entityInfoInstance.createCompany("EMPRESA_ENTRADA", "{JSON}");
        await entityInfoInstance.createCompany("MI_EMPRESA", "{JSON}");
        await entityInfoInstance.createCompany("EMPRESA_SALIDA", "{JSON}");
        await entityInfoInstance.createContainer("D5", "{JSON}");
        await entityInfoInstance.createContainer("D6", "{JSON}");
        await entityInfoInstance.createContainer("CC3", "{JSON}");

        // Transition 1
        let json = {
            "GAP (%Vol)": 12.6
        }
        await traceabilityInstance.productEntry(
            7000,
            "SAUVIGNON_BLANC_MOSTO",
            "EMPRESA_ENTRADA",
            "D5",
            JSON.stringify(json)
        );

        // Transition 2
        json = {
            "GAP (%Vol)": 12.6
        }
        await traceabilityInstance.productProcessing(
            1,
            0,
            "SAUVIGNON_BLANC_VINO",
            "MI_EMPRESA",
            "D6",
            JSON.stringify(json)
        );

        // Transition 3 (internal transition along with the next one)
        json = {}
        await traceabilityInstance.productPartition(
            1,
            3500,
            "MI_EMPRESA",
            "CC3",
            JSON.stringify(json)
        );

        // Transition 4
        json = {
            "Densidad (g/ml)": 1.093,
            "GAP (%Vol)": 12.8
        }
        await traceabilityInstance.productOutput(
            3,
            "LN_3", // Internal lot number
            "EMPRESA_SALIDA",
            "CC3",
            JSON.stringify(json)
        );

        // Traceability
        await common.getTraceability(traceabilityInstance, authInstance, "LN_3");
    });

    it("MOSTO SAUVIGNON BLANC A VINO BLANCO JOVEN EMBOTELLADO", async () => {
        // Required entities
        await entityInfoInstance.createProductType("SAUVIGNON_BLANC_MOSTO", "{JSON}");
        await entityInfoInstance.createProductType("SAUVIGNON_BLANC_VINO", "{JSON}");
        await entityInfoInstance.createProductType("SAUVIGNON_BLANC_BOTELLA", "{JSON}");
        await entityInfoInstance.createCompany("EMPRESA_ENTRADA", "{JSON}");
        await entityInfoInstance.createCompany("MI_EMPRESA", "{JSON}");
        await entityInfoInstance.createCompany("EMPRESA_SALIDA", "{JSON}");
        await entityInfoInstance.createContainer("D10", "{JSON}");
        await entityInfoInstance.createContainer("D11", "{JSON}");
        await entityInfoInstance.createContainer("DE", "{JSON}");
        await entityInfoInstance.createContainer("EMBOTELLADORA_X", "{JSON}");
        await entityInfoInstance.createContainer("CAMIÓN_X", "{JSON}");

        // Transition 1
        let json = {
            "GAP (%Vol)": 12.6
        }
        await traceabilityInstance.productEntry(
            7000,
            "SAUVIGNON_BLANC_MOSTO",
            "EMPRESA_ENTRADA",
            "D10",
            JSON.stringify(json)
        );

        // Transition 2
        json = {
            "GAP (%Vol)": 12.6
        }
        await traceabilityInstance.productProcessing(
            1,
            0,
            "SAUVIGNON_BLANC_VINO",
            "MI_EMPRESA",
            "D11",
            JSON.stringify(json)
        );

        // Transition 3
        json = {}
        await traceabilityInstance.productPartition(
            1,
            1500,
            "MI_EMPRESA",
            "DE",
            JSON.stringify(json)
        );

        // Transition 4
        json = {
            "Botellas": 2000,
            "GAP (%Vol)": 13
        }
        await traceabilityInstance.productProcessing(
            3,
            0,
            "SAUVIGNON_BLANC_BOTELLA",
            "MI_EMPRESA",
            "EMBOTELLADORA_X",
            JSON.stringify(json)
        );

        // Transition 5 (internal transition along with the previous one)
        json = {
            "Botellas": 2000,
            "GAP (%Vol)": 13
        }
        await traceabilityInstance.productOutput(
            3,
            "L2243",
            "EMPRESA_SALIDA",
            "CAMIÓN_X",
            JSON.stringify(json)
        );

        // Traceability
        await common.getTraceability(traceabilityInstance, authInstance, "L2243");
    });*/
});
