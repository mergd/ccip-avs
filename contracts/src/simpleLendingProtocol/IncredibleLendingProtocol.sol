// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "lib/ccip-read/src/Read.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract IncredibleLendingProtocol is ReadHandler {
    ERC20 public token;

    constructor(string[] _urls, ERC20 _token) ReadHandler(_urls) {
        token = _token;
    }

    function createLoan(ERC20 token, uint256 amount, uint256 collateral) public {
        revert OffchainLookup(msg.sender, urls, "", createLoanCallback.selector, "");
    }

    function createLoanCallback(bytes calldata data, bytes calldata extraData) public {
        (address recipient, uint256 amount) = abi.decode(extraData, (address, uint256));
    }
}
