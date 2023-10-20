module.exports.getTraceability = async function (traceabilityInstance, authInstance, lotNumber) {
    console.log("----- Traceability of \"" + lotNumber + "\" -----");

    // Get product ID
    let currentProductID = (await traceabilityInstance.lotNumbers(lotNumber)).toNumber();
    if (currentProductID > 0) {
        let totalTransitions = [];
        let product = await traceabilityInstance.products(currentProductID);

        // For each product in the traceability vector
        do {
            // Get product's transitions
            let totalLosses = 0;
            let tmpTransitions = [];
            for (const transition of [...product.tv.transitions]) {
                totalLosses += Number(transition.lostQuantity);

                // Custom transition object
                const newTransition = {
                    ...transition,
                    currentProductID: currentProductID,
                    quantity: product.quantity - totalLosses
                };
                tmpTransitions.push(newTransition);
            }

            // Sort and save product's transitions
            totalTransitions = totalTransitions.concat(tmpTransitions.reverse());

            // Get previous product in the traceability vector
            currentProductID = product.tv.prevProductID;
            product = await traceabilityInstance.products(currentProductID);
        } while (currentProductID > 0);

        // Print traceability vector
        for (const transition of totalTransitions.reverse()) {
            console.log(
                "PRODUCT:" + transition.currentProductID, "|",
                (await authInstance.transitionTypes(transition.typeID)).info, "|",
                transition.quantity, "|",
                Number(transition.lostQuantity), "|",
                transition.productTypeID, "|",
                transition.companyID, "|",
                transition.containerID, "|",
                parseTimestamp(transition.createdAt), "|",
                transition.info);
        }
    } else {
        console.log("Unknown lot number");
    }
}

function parseTimestamp(timestamp) {
    const date = new Date(Number(timestamp + "000"));
    return date.getDay() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear();
}
