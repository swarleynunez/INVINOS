// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract EntityInfo {
    // Smart contract deployer/owner
    address public owner;

    // Data structures
    struct Entity {
        string info; // Encoded
        bool created;
    }

    // Data variables
    mapping(string => Entity) public productTypes;
    mapping(string => Entity) public companies;
    mapping(string => Entity) public containers;

    // Init
    constructor(address _owner) {
        owner = _owner;
    }

    // Modifiers
    modifier onlyOwner() {
        require(
            msg.sender == owner,
            "Only the contract owner can call this function"
        );
        _;
    }

    // Setters
    function createProductType(
        string memory _id,
        string memory _info
    ) public onlyOwner {
        require(!existProductType(_id), "Product type already created");

        // Create product type
        productTypes[_id].info = _info;
        productTypes[_id].created = true;
    }

    function createCompany(
        string memory _id,
        string memory _info
    ) public onlyOwner {
        require(!existCompany(_id), "Company already created");

        // Create company
        companies[_id].info = _info;
        companies[_id].created = true;
    }

    function createContainer(
        string memory _id,
        string memory _info
    ) public onlyOwner {
        require(!existContainer(_id), "Container already created");

        // Create container
        containers[_id].info = _info;
        containers[_id].created = true;
    }

    // Getters

    // Helpers
    function existProductType(string memory _id) public view returns (bool) {
        return productTypes[_id].created;
    }

    function existCompany(string memory _id) public view returns (bool) {
        return companies[_id].created;
    }

    function existContainer(string memory _id) public view returns (bool) {
        return containers[_id].created;
    }
}
