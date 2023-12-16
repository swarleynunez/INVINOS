// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract Auth {
    // Smart contract deployer/owner
    address public owner;

    // Data structures
    struct TransitionType {
        string info; // Encoded
        bool created;
    }

    // Data variables
    mapping(uint => TransitionType) public transitionTypes;

    // Init
    constructor(address _owner) {
        owner = _owner;

        // Create transition types
        transitionTypes[0].info = "PRODUCT_ENTRY";
        transitionTypes[0].created = true;

        transitionTypes[1].info = "PRODUCT_PROCESSING";
        transitionTypes[1].created = true;

        transitionTypes[2].info = "PRODUCT_PARTITION";
        transitionTypes[2].created = true;

        transitionTypes[3].info = "PRODUCT_OUTPUT";
        transitionTypes[3].created = true;
    }

    // Setters

    // Getters

    // Helpers
}
