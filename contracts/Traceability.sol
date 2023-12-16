// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./Counters.sol";
import "./Auth.sol";
import "./EntityInfo.sol";

contract Traceability {
    // External libraries
    using Counters for Counters.Counter;

    // Smart contract deployer/owner
    address public owner;

    // Complementary smart contracts
    Auth public authContract;
    EntityInfo public entityInfoContract;

    // Data structures
    struct Product {
        uint quantity;
        TraceabilityVector tv;
        bool partitioned;
        bool completed;
    }

    struct TraceabilityVector {
        Transition[] transitions;
        uint prevProductID; // Optional
    }

    struct Transition {
        uint typeID;
        uint lostQuantity; // Optional
        string productTypeID;
        string companyID;
        string containerID;
        uint createdAt;
        string info; // Encoded
    }

    // Data variables
    mapping(uint => Product) public products;
    mapping(string => uint) public lotNumbers;

    // Other variables
    Counters.Counter private productID;

    // Init
    constructor() {
        owner = msg.sender;

        // Complementary smart contract instances
        authContract = new Auth(owner);
        entityInfoContract = new EntityInfo(owner);
    }

    // Modifiers
    modifier onlyOwner() {
        require(
            msg.sender == owner,
            "Only the contract owner can call this function"
        );
        _;
    }

    /////////////
    // Setters //
    /////////////
    function productEntry(
        uint _quantity,
        string memory _productTypeID,
        string memory _companyID,
        string memory _containerID,
        string memory _info
    ) public onlyOwner {
        require(_quantity > 0, "Quantity of 0 not allowed");
        require(
            entityInfoContract.existProductType(_productTypeID),
            "Product type does not exist"
        );
        require(
            entityInfoContract.existCompany(_companyID),
            "Company does not exist"
        );
        require(
            entityInfoContract.existContainer(_containerID),
            "Container does not exist"
        );

        // Update product counter (avoiding ID 0)
        productID.increment();

        // Create transition object
        Transition memory t = createTransition(
            0, // 0 = PRODUCT_ENTRY
            0, // Losses
            _productTypeID,
            _companyID,
            _containerID,
            _info
        );

        // Create product object
        products[productID.current()].quantity = _quantity;
        products[productID.current()].tv.transitions.push(t);
    }

    // Internal processing/movements (e.g. bottling)
    function productProcessing(
        uint _productID,
        uint _lostQuantity,
        string memory _productTypeID,
        string memory _containerID,
        string memory _info
    ) public onlyOwner {
        require(isProductAvailable(_productID), "Product is not available");
        require(
            _lostQuantity <= getProductQuantity(_productID),
            "Not enough product available"
        );
        require(
            entityInfoContract.existProductType(_productTypeID),
            "Product type does not exist"
        );
        require(
            entityInfoContract.existContainer(_containerID),
            "Container does not exist"
        );

        // Create transition object
        uint lastIndex = products[_productID].tv.transitions.length - 1;
        Transition memory t = createTransition(
            1, // 1 = PRODUCT_PROCESSING
            _lostQuantity, // Losses
            _productTypeID,
            products[_productID].tv.transitions[lastIndex].companyID,
            _containerID,
            _info
        );

        // Update product object
        products[_productID].tv.transitions.push(t);
    }

    function productPartition(
        uint _productID,
        uint _quantity,
        string memory _containerID,
        string memory _info
    ) public onlyOwner {
        require(isProductAvailable(_productID), "Product is not available");
        require(_quantity > 0, "Quantity of 0 not allowed");
        require(
            _quantity < getProductQuantity(_productID),
            "Not enough product available"
        );
        require(
            entityInfoContract.existContainer(_containerID),
            "Container does not exist"
        );

        // Update old product
        products[_productID].partitioned = true;

        // Create transition 1
        uint lastIndex = products[_productID].tv.transitions.length - 1;
        Transition memory t1 = createTransition(
            2, // 2 = PRODUCT_PARTITION
            0, // Losses
            products[_productID].tv.transitions[lastIndex].productTypeID,
            products[_productID].tv.transitions[lastIndex].companyID,
            products[_productID].tv.transitions[lastIndex].containerID,
            products[_productID].tv.transitions[lastIndex].info
        );

        // Create product 1
        productID.increment();
        products[productID.current()].quantity =
            getProductQuantity(_productID) -
            _quantity;
        products[productID.current()].tv.transitions.push(t1);
        products[productID.current()].tv.prevProductID = _productID;

        // Create transition 2
        Transition memory t2 = createTransition(
            2, // 2 = PRODUCT_PARTITION
            0, // Losses
            products[_productID].tv.transitions[lastIndex].productTypeID,
            products[_productID].tv.transitions[lastIndex].companyID,
            _containerID,
            _info
        );

        // Create product 2
        productID.increment();
        products[productID.current()].quantity = _quantity;
        products[productID.current()].tv.transitions.push(t2);
        products[productID.current()].tv.prevProductID = _productID;
    }

    // Packaging and sales
    function productOutput(
        uint _productID,
        string memory _lotNumber,
        string memory _companyID,
        string memory _containerID,
        string memory _info
    ) public onlyOwner {
        require(isProductAvailable(_productID), "Product is not available");
        require(
            isLotNumberAvailable(_lotNumber),
            "Lot number is not available"
        );
        require(
            entityInfoContract.existCompany(_companyID),
            "Company does not exist"
        );
        require(
            entityInfoContract.existContainer(_containerID),
            "Container does not exist"
        );

        // Complete/close product
        products[_productID].completed = true;

        // Link a final lot number to the product
        lotNumbers[_lotNumber] = _productID;

        // Create transition object
        uint lastIndex = products[_productID].tv.transitions.length - 1;
        Transition memory t = createTransition(
            3, // 3 = PRODUCT_OUTPUT
            0, // Losses
            products[_productID].tv.transitions[lastIndex].productTypeID,
            _companyID,
            _containerID,
            _info
        );

        // Update product object
        products[_productID].tv.transitions.push(t);
    }

    /////////////
    // Getters //
    /////////////
    function getProductQuantity(uint _productID) public view returns (uint) {
        uint totalLosses;
        for (uint i = 0; i < products[_productID].tv.transitions.length; i++) {
            totalLosses += products[_productID].tv.transitions[i].lostQuantity;
        }

        return products[_productID].quantity - totalLosses;
    }

    /////////////
    // Helpers //
    /////////////
    function isProductAvailable(uint _productID) public view returns (bool) {
        return
            products[_productID].tv.transitions.length > 0 &&
            !products[_productID].partitioned &&
            !products[_productID].completed;
    }

    function isLotNumberAvailable(
        string memory _lotNumber
    ) public view returns (bool) {
        return lotNumbers[_lotNumber] == 0;
    }

    function createTransition(
        uint _typeID,
        uint _lostQuantity,
        string memory _productTypeID,
        string memory _companyID,
        string memory _containerID,
        string memory _info
    ) private view returns (Transition memory) {
        Transition memory t;
        t.typeID = _typeID;
        t.lostQuantity = _lostQuantity;
        t.productTypeID = _productTypeID;
        t.companyID = _companyID;
        t.containerID = _containerID;
        t.createdAt = block.timestamp;
        t.info = _info;

        return t;
    }
}
