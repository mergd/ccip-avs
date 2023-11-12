// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./IIncredibleLendingTaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

/**
 * @title Primary entrypoint for procuring services from IncredibleLending.
 * @author Layr Labs, Inc.
 */
contract IncredibleLendingServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    IIncredibleLendingTaskManager public immutable incredibleLendingTaskManager;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyIncredibleLendingTaskManager() {
        require(
            msg.sender == address(incredibleLendingTaskManager),
            "onlyIncredibleLendingTaskManager: not from credible lending task manager"
        );
        _;
    }

    constructor(
        IBLSRegistryCoordinatorWithIndices _registryCoordinator,
        ISlasher _slasher,
        IIncredibleLendingTaskManager _incredibleLendingTaskManager
    ) ServiceManagerBase(_registryCoordinator, _slasher) {
        incredibleLendingTaskManager = _incredibleLendingTaskManager;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(address operatorAddr) external override onlyIncredibleLendingTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
