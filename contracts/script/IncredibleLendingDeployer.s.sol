// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";
import {LoanCoordinator, ERC20} from "landingprotocol/src/LoanCoordinator.sol";

import {BLSPublicKeyCompendium} from "@eigenlayer-middleware/src/BLSPublicKeyCompendium.sol";
import "@eigenlayer-middleware/src/BLSRegistryCoordinatorWithIndices.sol" as blsregcoord;
import {BLSPubkeyRegistry, IBLSPubkeyRegistry} from "@eigenlayer-middleware/src/BLSPubkeyRegistry.sol";
import {IndexRegistry, IIndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry, IStakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import {IVoteWeigher} from "@eigenlayer-middleware/src/interfaces/IVoteWeigher.sol";

import {IncredibleLendingServiceManager, IServiceManager} from "../src/IncredibleLendingServiceManager.sol";
import {IncredibleLendingTaskManager} from "../src/IncredibleLendingTaskManager.sol";
import {IIncredibleLendingTaskManager} from "../src/IIncredibleLendingTaskManager.sol";
import "../src/ERC20Mock.sol";
import {OnchainDepthOracle} from "src/OnChainDepthOracle.sol";
import {IncredibleLendingProtocol, ILoanCoordinator} from "src/IncredibleLendingProtocol.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/CredibleLendingDeployer.s.sol:CredibleLendingDeployer --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract CredibleLendingDeployer is Script, Utils {
    // DEPLOYMENT CONSTANTS
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 100;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_DURATION_BLOCKS = 0;
    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
    address public constant TASK_GENERATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    // ERC20 and Strategy: we need to deploy this erc20, create a strategy for it, and whitelist this strategy in the strategymanager

    ERC20Mock public erc20Mock;
    StrategyBaseTVLLimits public erc20MockStrategy;

    // Credible Lending contracts
    ProxyAdmin public credibleLendingProxyAdmin;
    PauserRegistry public credibleLendingPauserReg;

    blsregcoord.BLSRegistryCoordinatorWithIndices public registryCoordinator;
    blsregcoord.IBLSRegistryCoordinatorWithIndices public registryCoordinatorImplementation;

    IBLSPubkeyRegistry public blsPubkeyRegistry;
    IBLSPubkeyRegistry public blsPubkeyRegistryImplementation;

    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;

    IncredibleLendingServiceManager public credibleLendingServiceManager;
    IServiceManager public credibleLendingServiceManagerImplementation;

    IncredibleLendingTaskManager public credibleLendingTaskManager;
    IIncredibleLendingTaskManager public credibleLendingTaskManagerImplementation;

    //
    OnchainDepthOracle public onchainDepthOracle;
    IncredibleLendingProtocol public incredibleLendingProtocol;
    ILoanCoordinator public loanCoordinator;
    ERC20 public mockCollateral;
    ERC20 public mockWETH;

    function run() external {
        uint256 deployer = vm.envUint("DEPLOYER_KEY");

        // Eigenlayer contracts
        string memory eigenlayerDeployedContracts = readOutput("eigenlayer_deployment_output");
        string memory sharedAvsDeployedContracts = readOutput("shared_avs_contracts_deployment_output");
        IStrategyManager strategyManager =
            IStrategyManager(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.strategyManager"));
        IDelegationManager delegationManager =
            IDelegationManager(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.delegation"));
        ISlasher slasher = ISlasher(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.slasher"));
        ProxyAdmin eigenLayerProxyAdmin =
            ProxyAdmin(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.eigenLayerProxyAdmin"));
        PauserRegistry eigenLayerPauserReg =
            PauserRegistry(stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.eigenLayerPauserReg"));
        BLSPublicKeyCompendium pubkeyCompendium =
            BLSPublicKeyCompendium(stdJson.readAddress(sharedAvsDeployedContracts, ".blsPublicKeyCompendium"));
        StrategyBaseTVLLimits baseStrategyImplementation = StrategyBaseTVLLimits(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.baseStrategyImplementation")
        );

        address credibleLendingCommunityMultisig = msg.sender;
        address credibleLendingPauser = msg.sender;

        vm.startBroadcast(deployer);
        // Lending contracts
        mockCollateral = ERC20(address(new ERC20Mock()));
        mockWETH = ERC20(address(new ERC20Mock()));
        loanCoordinator = new LoanCoordinator();
        onchainDepthOracle = new OnchainDepthOracle(mockWETH);
        incredibleLendingProtocol = new IncredibleLendingProtocol(
            new string[](0),
            mockWETH,
            mockCollateral,
            loanCoordinator,
            TASK_GENERATOR_ADDR);

        _deployErc20AndStrategyAndWhitelistStrategy(
            eigenLayerProxyAdmin, eigenLayerPauserReg, baseStrategyImplementation, strategyManager
        );
        _deployCredibleLendingContracts(
            strategyManager,
            delegationManager,
            slasher,
            erc20MockStrategy,
            pubkeyCompendium,
            credibleLendingCommunityMultisig,
            credibleLendingPauser
        );

        // set tm
        incredibleLendingProtocol.setTaskManager(credibleLendingTaskManager);

        vm.stopBroadcast();
    }

    function _deployErc20AndStrategyAndWhitelistStrategy(
        ProxyAdmin eigenLayerProxyAdmin,
        PauserRegistry eigenLayerPauserReg,
        StrategyBaseTVLLimits baseStrategyImplementation,
        IStrategyManager strategyManager
    ) internal {
        erc20Mock = new ERC20Mock();
        // TODO(samlaf): any reason why we are using the strategybase with tvl limits instead of just using strategybase?
        // the maxPerDeposit and maxDeposits below are just arbitrary values.
        erc20MockStrategy = StrategyBaseTVLLimits(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        StrategyBaseTVLLimits.initialize.selector,
                        1 ether, // maxPerDeposit
                        100 ether, // maxDeposits
                        IERC20(erc20Mock),
                        eigenLayerPauserReg
                    )
                )
            )
        );
        IStrategy[] memory strats = new IStrategy[](1);
        strats[0] = erc20MockStrategy;
        strategyManager.addStrategiesToDepositWhitelist(strats);
    }

    function _deployCredibleLendingContracts(
        IStrategyManager strategyManager,
        IDelegationManager delegationManager,
        ISlasher slasher,
        IStrategy strat,
        BLSPublicKeyCompendium pubkeyCompendium,
        address credibleLendingCommunityMultisig,
        address credibleLendingPauser
    ) internal {
        // Adding this as a temporary fix to make the rest of the script work with a single strategy
        // since it was originally written to work with an array of strategies
        IStrategy[1] memory deployedStrategyArray = [strat];
        uint256 numStrategies = deployedStrategyArray.length;

        // deploy proxy admin for ability to upgrade proxy contracts
        credibleLendingProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        {
            address[] memory pausers = new address[](2);
            pausers[0] = credibleLendingPauser;
            pausers[1] = credibleLendingCommunityMultisig;
            credibleLendingPauserReg = new PauserRegistry(
                pausers,
                credibleLendingCommunityMultisig
            );
        }

        EmptyContract emptyContract = new EmptyContract();

        // hard-coded inputs

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        credibleLendingServiceManager = IncredibleLendingServiceManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(credibleLendingProxyAdmin),
                    ""
                )
            )
        );
        credibleLendingTaskManager = IncredibleLendingTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(credibleLendingProxyAdmin),
                    ""
                )
            )
        );
        registryCoordinator = blsregcoord.BLSRegistryCoordinatorWithIndices(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(credibleLendingProxyAdmin),
                    ""
                )
            )
        );
        blsPubkeyRegistry = IBLSPubkeyRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(credibleLendingProxyAdmin),
                    ""
                )
            )
        );
        indexRegistry = IIndexRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(credibleLendingProxyAdmin),
                    ""
                )
            )
        );
        stakeRegistry = IStakeRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(credibleLendingProxyAdmin),
                    ""
                )
            )
        );

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        {
            stakeRegistryImplementation = new StakeRegistry(
                registryCoordinator,
                strategyManager,
                credibleLendingServiceManager
            );

            // set up a quorum with each strategy that needs to be set up
            uint96[] memory minimumStakeForQuorum = new uint96[](numStrategies);
            IVoteWeigher.StrategyAndWeightingMultiplier[][] memory strategyAndWeightingMultipliers =
            new IVoteWeigher.StrategyAndWeightingMultiplier[][](
                    numStrategies
                );
            for (uint256 i = 0; i < numStrategies; i++) {
                strategyAndWeightingMultipliers[i] = new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
                strategyAndWeightingMultipliers[i][0] = IVoteWeigher.StrategyAndWeightingMultiplier({
                    strategy: deployedStrategyArray[i],
                    // setting this to 1 ether since the divisor is also 1 ether
                    // therefore this allows an operator to register with even just 1 token
                    // see ./eigenlayer-contracts/src/contracts/middleware/VoteWeigherBase.sol#L81
                    //    weight += uint96(sharesAmount * strategyAndMultiplier.multiplier / WEIGHTING_DIVISOR);
                    multiplier: 1 ether
                });
            }

            credibleLendingProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))),
                address(stakeRegistryImplementation),
                abi.encodeWithSelector(
                    StakeRegistry.initialize.selector, minimumStakeForQuorum, strategyAndWeightingMultipliers
                )
            );
        }

        registryCoordinatorImplementation = new blsregcoord.BLSRegistryCoordinatorWithIndices(
            slasher,
            credibleLendingServiceManager,
            blsregcoord.IStakeRegistry(address(stakeRegistry)),
            blsregcoord.IBLSPubkeyRegistry(address(blsPubkeyRegistry)),
            blsregcoord.IIndexRegistry(address(indexRegistry))
        );

        {
            blsregcoord.IBLSRegistryCoordinatorWithIndices.OperatorSetParam[] memory operatorSetParams =
            new blsregcoord.IBLSRegistryCoordinatorWithIndices.OperatorSetParam[](
                    numStrategies
                );
            for (uint256 i = 0; i < numStrategies; i++) {
                // hard code these for now
                operatorSetParams[i] = blsregcoord.IBLSRegistryCoordinatorWithIndices.OperatorSetParam({
                    maxOperatorCount: 10000,
                    kickBIPsOfOperatorStake: 15000,
                    kickBIPsOfTotalStake: 100
                });
            }
            credibleLendingProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(registryCoordinator))),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    blsregcoord.BLSRegistryCoordinatorWithIndices.initialize.selector,
                    // we set churnApprover and ejector to communityMultisig because we don't need them
                    credibleLendingCommunityMultisig,
                    credibleLendingCommunityMultisig,
                    operatorSetParams,
                    credibleLendingPauserReg,
                    // 0 initialPausedStatus means everything unpaused
                    0
                )
            );
        }

        blsPubkeyRegistryImplementation = new BLSPubkeyRegistry(
            registryCoordinator,
            pubkeyCompendium
        );

        credibleLendingProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(blsPubkeyRegistry))), address(blsPubkeyRegistryImplementation)
        );

        indexRegistryImplementation = new IndexRegistry(registryCoordinator);

        credibleLendingProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(indexRegistry))), address(indexRegistryImplementation)
        );

        credibleLendingServiceManagerImplementation = new IncredibleLendingServiceManager(
            registryCoordinator,
            slasher,
            credibleLendingTaskManager
        );
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        credibleLendingProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(credibleLendingServiceManager))),
            address(credibleLendingServiceManagerImplementation),
            abi.encodeWithSelector(
                credibleLendingServiceManager.initialize.selector,
                credibleLendingPauserReg,
                credibleLendingCommunityMultisig
            )
        );

        credibleLendingTaskManagerImplementation = new IncredibleLendingTaskManager(
            registryCoordinator,
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        credibleLendingProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(credibleLendingTaskManager))),
            address(credibleLendingTaskManagerImplementation),
            abi.encodeWithSelector(
                credibleLendingTaskManager.initialize.selector,
                credibleLendingPauserReg,
                incredibleLendingProtocol,
                onchainDepthOracle,
                credibleLendingCommunityMultisig,
                AGGREGATOR_ADDR,
                TASK_GENERATOR_ADDR,
                QUORUM_THRESHOLD_PERCENTAGE
            )
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(deployed_addresses, "erc20Mock", address(erc20Mock));
        vm.serializeAddress(deployed_addresses, "erc20MockStrategy", address(erc20MockStrategy));
        vm.serializeAddress(deployed_addresses, "credibleLendingServiceManager", address(credibleLendingServiceManager));
        vm.serializeAddress(
            deployed_addresses,
            "credibleLendingServiceManagerImplementation",
            address(credibleLendingServiceManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "credibleLendingTaskManager", address(credibleLendingTaskManager));
        vm.serializeAddress(
            deployed_addresses,
            "credibleLendingTaskManagerImplementation",
            address(credibleLendingTaskManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "registryCoordinator", address(registryCoordinator));
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses, "registryCoordinatorImplementation", address(registryCoordinatorImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);

        writeOutput(finalJson, "credible_lending_avs_deployment_output");
    }
}
