// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { IICS26RouterMsgs } from "./msgs/IICS26RouterMsgs.sol";
import { IICS20TransferMsgs } from "./msgs/IICS20TransferMsgs.sol";

import { IICS20Errors } from "./errors/IICS20Errors.sol";
import { IEscrow } from "./interfaces/IEscrow.sol";
import { IIBCApp } from "./interfaces/IIBCApp.sol";
import { IERC20 } from "@openzeppelin-contracts/token/ERC20/IERC20.sol";
import { IICS20Transfer } from "./interfaces/IICS20Transfer.sol";
import { IICS26Router } from "./interfaces/IICS26Router.sol";

import { ReentrancyGuardTransientUpgradeable } from
    "@openzeppelin-upgradeable/utils/ReentrancyGuardTransientUpgradeable.sol";
import { SafeERC20 } from "@openzeppelin-contracts/token/ERC20/utils/SafeERC20.sol";
import { MulticallUpgradeable } from "@openzeppelin-upgradeable/utils/MulticallUpgradeable.sol";
import { ICS20Lib } from "./utils/ICS20Lib.sol";
import { IBCERC20 } from "./utils/IBCERC20.sol";
import { Escrow } from "./utils/Escrow.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { UUPSUpgradeable } from "@openzeppelin-contracts/proxy/utils/UUPSUpgradeable.sol";
import { IIBCUUPSUpgradeable } from "./interfaces/IIBCUUPSUpgradeable.sol";

using SafeERC20 for IERC20;

/*
 * Things not handled yet:
 * - Separate escrow balance tracking
 * - Related to escrow ^: invariant checking (where to implement that?)
 */
contract ICS20Transfer is
    IICS20Errors,
    IICS20Transfer,
    IIBCApp,
    ReentrancyGuardTransientUpgradeable,
    MulticallUpgradeable,
    UUPSUpgradeable
{
    /// @notice Storage of the ICS20Transfer contract
    /// @dev It's implemented on a custom ERC-7201 namespace to reduce the risk of storage collisions when using with
    /// upgradeable contracts.
    /// @param escrow The escrow contract
    /// @param ibcDenomContracts Mapping of non-native denoms to their respective IBCERC20 contracts
    /// @param ics26Router The ICS26Router contract
    /// @custom:storage-location erc7201:ibc.storage.ICS20Transfer
    struct ICS20TransferStorage {
        IEscrow escrow;
        mapping(bytes32 => IBCERC20) ibcDenomContracts;
        IICS26Router ics26Router;
    }

    /// @notice ERC-7201 slot for the ICS20Transfer storage
    /// @dev keccak256(abi.encode(uint256(keccak256("ibc.storage.ICS20Transfer")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ICS20TRANSFER_STORAGE_SLOT =
        0x823f7a8ea9ae6df0eb03ec5e1682d7a2839417ad8a91774118e6acf2e8d2f800;

    /// @dev This contract is meant to be deployed by a proxy, so the constructor is not used
    constructor() {
        _disableInitializers();
    }

    /// @notice Initializes the contract instead of a constructor
    /// @dev Meant to be called only once from the proxy
    /// @param ics26Router The ICS26Router contract address
    function initialize(address ics26Router) public initializer {
        __ReentrancyGuardTransient_init();
        __Multicall_init();

        ICS20TransferStorage storage $ = _getICS20TransferStorage();

        $.escrow = new Escrow(address(this));
        $.ics26Router = IICS26Router(ics26Router);
    }

    /// @inheritdoc IICS20Transfer
    function escrow() public view returns (address) {
        return address(_getEscrow());
    }

    /// @inheritdoc IICS20Transfer
    function ibcERC20Contract(IICS20TransferMsgs.Denom calldata denom) external view returns (address) {
        bytes32 denomID = ICS20Lib.getDenomIdentifier(denom);
        address contractAddress = address(_getICS20TransferStorage().ibcDenomContracts[denomID]);
        require(contractAddress != address(0), ICS20DenomNotFound(denom));
        return contractAddress;
    }

    /// @inheritdoc IICS20Transfer
    function newMsgSendPacketV2(
        address sender,
        IICS20TransferMsgs.SendTransferMsg calldata msg_
    )
        external
        view
        override
        returns (IICS26RouterMsgs.MsgSendPacket memory)
    {
        return ICS20Lib.newMsgSendPacketV2(sender, msg_);
    }

    /// @inheritdoc IICS20Transfer
    function sendTransfer(IICS20TransferMsgs.SendTransferMsg calldata msg_) external override returns (uint32) {
        return _getICS26Router().sendPacket(ICS20Lib.newMsgSendPacketV2(_msgSender(), msg_));
    }

    /// @inheritdoc IIBCApp
    function onSendPacket(OnSendPacketCallback calldata msg_) external onlyRouter nonReentrant {
        require(
            keccak256(bytes(msg_.payload.version)) == keccak256(bytes(ICS20Lib.ICS20_VERSION)),
            ICS20UnexpectedVersion(ICS20Lib.ICS20_VERSION, msg_.payload.version)
        );

        IICS20TransferMsgs.FungibleTokenPacketDataV2 memory packetData =
            abi.decode(msg_.payload.value, (IICS20TransferMsgs.FungibleTokenPacketDataV2));

        (bool valid, string memory error) = ICS20Lib.validatePacketData(packetData);
        require(valid, ICS20InvalidPacketData(error));

        address sender = ICS20Lib.mustHexStringToAddress(packetData.sender);

        // only the sender in the payload or this contract (sendTransfer) can send the packet
        require(msg_.sender == sender || msg_.sender == address(this), ICS20UnauthorizedPacketSender(msg_.sender));

        for (uint256 i = 0; i < packetData.tokens.length; i++) {
            IICS20TransferMsgs.Token memory token = packetData.tokens[i];
            require(token.amount > 0, ICS20InvalidAmount(token.amount));

            (bool returningToSource, address erc20Address) =
                _getSendingERC20Address(msg_.payload.sourcePort, msg_.sourceClient, token.denom);

            // transfer the tokens to us (requires the allowance to be set)
            _transferFrom(sender, escrow(), erc20Address, token.amount);

            if (returningToSource) {
                // if the token is returning to source, it is an IBCERC20 and we must burn the token (not keep it in
                // escrow it)
                IBCERC20(erc20Address).burn(token.amount);
            }
        }
    }

    /// @inheritdoc IIBCApp
    function onRecvPacket(OnRecvPacketCallback calldata msg_) external onlyRouter nonReentrant returns (bytes memory) {
        // TODO: Figure out if should actually error out, or if just error acking is enough (#112)

        // Since this function mostly returns acks, also when it fails, the ics26router (the caller) will log the ack
        if (keccak256(bytes(msg_.payload.version)) != keccak256(bytes(ICS20Lib.ICS20_VERSION))) {
            return ICS20Lib.errorAck(abi.encodePacked("unexpected version: ", msg_.payload.version));
        }

        IICS20TransferMsgs.FungibleTokenPacketDataV2 memory packetData =
            abi.decode(msg_.payload.value, (IICS20TransferMsgs.FungibleTokenPacketDataV2));

        if (packetData.forwarding.hops.length > 0) {
            return ICS20Lib.errorAck("unsupported feature: forwarding on receive");
        }

        (bool valid, string memory error) = ICS20Lib.validatePacketData(packetData);
        if (!valid) {
            return ICS20Lib.errorAck(abi.encodePacked("invalid packet data: ", bytes(error)));
        }

        (bool receiverConvertSuccess, address receiver) = Strings.tryParseAddress(packetData.receiver);
        if (!receiverConvertSuccess) {
            return ICS20Lib.errorAck(abi.encodePacked("invalid receiver: ", packetData.receiver));
        }

        for (uint256 i = 0; i < packetData.tokens.length; i++) {
            IICS20TransferMsgs.Token memory token = packetData.tokens[i];

            // This is the prefix that would have been prefixed to the denomination
            // on sender chain IF and only if the token originally came from the
            // receiving chain.
            //
            // NOTE: We use SourcePort and SourceChannel here, because the counterparty
            // chain would have prefixed with DestPort and DestChannel when originally
            // receiving this token.
            bool returningToOrigin = ICS20Lib.hasPrefix(token.denom, msg_.payload.sourcePort, msg_.sourceClient);

            address erc20Address;
            if (returningToOrigin) {
                // we are the origin source of this token: it is either an IBCERC20 or a "native" ERC20:
                // remove the first hop to unwind the trace
                IICS20TransferMsgs.Denom memory newDenom = ICS20Lib.removeHop(token.denom);

                if (newDenom.trace.length == 0) {
                    // we are the origin source and the token is a native token
                    erc20Address = ICS20Lib.mustHexStringToAddress(newDenom.base);
                } else {
                    // we are the origin source and the token must be an IBCERC20 (since it has a trace left):
                    bytes32 denomID = ICS20Lib.getDenomIdentifier(newDenom);
                    erc20Address = address(_getICS20TransferStorage().ibcDenomContracts[denomID]);
                }
            } else {
                // we are not origin source, i.e. sender chain is the origin source: add denom trace and mint vouchers
                IICS20TransferMsgs.Denom memory newDenom = ICS20Lib.addHop(
                    token.denom,
                    IICS20TransferMsgs.Hop({ portId: msg_.payload.destPort, clientId: msg_.destinationClient })
                );

                erc20Address = _findOrCreateERC20Address(newDenom);
                IBCERC20(erc20Address).mint(token.amount);
            }

            // transfer the tokens to the receiver
            // solhint-disable-next-line multiple-sends
            _getEscrow().send(IERC20(erc20Address), receiver, token.amount);
        }

        return ICS20Lib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON;
    }

    /// @inheritdoc IIBCApp
    function onAcknowledgementPacket(OnAcknowledgementPacketCallback calldata msg_) external onlyRouter nonReentrant {
        if (keccak256(msg_.acknowledgement) != ICS20Lib.KECCAK256_SUCCESSFUL_ACKNOWLEDGEMENT_JSON) {
            IICS20TransferMsgs.FungibleTokenPacketDataV2 memory packetData =
                abi.decode(msg_.payload.value, (IICS20TransferMsgs.FungibleTokenPacketDataV2));

            _refundTokens(msg_.payload.sourcePort, msg_.sourceClient, packetData);
        }
    }

    /// @inheritdoc IIBCApp
    function onTimeoutPacket(OnTimeoutPacketCallback calldata msg_) external onlyRouter nonReentrant {
        IICS20TransferMsgs.FungibleTokenPacketDataV2 memory packetData =
            abi.decode(msg_.payload.value, (IICS20TransferMsgs.FungibleTokenPacketDataV2));
        _refundTokens(msg_.payload.sourcePort, msg_.sourceClient, packetData);
    }

    /// @notice Refund the tokens to the sender
    /// @param sourcePort The source port of the packet
    /// @param sourceClient The source client of the packet
    /// @param packetData The packet data
    function _refundTokens(
        string calldata sourcePort,
        string calldata sourceClient,
        IICS20TransferMsgs.FungibleTokenPacketDataV2 memory packetData
    )
        private
    {
        address refundee = ICS20Lib.mustHexStringToAddress(packetData.sender);

        for (uint256 i = 0; i < packetData.tokens.length; i++) {
            IICS20TransferMsgs.Token memory token = packetData.tokens[i];

            (bool returningToSource, address erc20Address) =
                _getSendingERC20Address(sourcePort, sourceClient, token.denom);

            if (returningToSource) {
                // if the token was returning to source, it was burned on send, so we mint it back now
                IBCERC20(erc20Address).mint(token.amount);
            }

            // solhint-disable-next-line multiple-sends
            _getEscrow().send(IERC20(erc20Address), refundee, token.amount);
        }
    }

    function _getSendingERC20Address(
        string calldata sourcePort,
        string calldata sourceClient,
        IICS20TransferMsgs.Denom memory denom
    )
        private
        view
        returns (bool returningToSource, address erc20Address)
    {
        bytes32 denomID = ICS20Lib.getDenomIdentifier(denom);

        // if the denom is prefixed by the port and channel on which we are sending
        // the token, then we must be returning the token back to the chain they originated from
        returningToSource = ICS20Lib.hasPrefix(denom, sourcePort, sourceClient);
        if (returningToSource) {
            // receiving chain is source of the token, so we've received and mapped this token before
            erc20Address = address(_getICS20TransferStorage().ibcDenomContracts[denomID]);
        } else {
            // the receiving chain is not the source of the token, so the token is either a native token
            // or we are a middle chain and the token was minted (and mapped) here.
            // NOTE: We check if the token is mapped _first_, to avoid a scenario where someone has a base denom
            // that is an address on their chain, and we would parse it as an address and fail to find the
            // mapped contract.
            address ibcDenomContract = address(_getICS20TransferStorage().ibcDenomContracts[denomID]);
            if (ibcDenomContract != address(0)) {
                erc20Address = ibcDenomContract;
            } else {
                // the token is not mapped, so the token must be native
                erc20Address = ICS20Lib.mustHexStringToAddress(denom.base);
            }
        }
        require(erc20Address != address(0), ICS20DenomNotFound(denom));

        return (returningToSource, erc20Address);
    }

    /// @notice Transfer tokens from sender to receiver
    /// @param sender The sender of the tokens
    /// @param receiver The receiver of the tokens
    /// @param tokenContract The address of the token contract
    /// @param amount The amount of tokens to transfer
    function _transferFrom(address sender, address receiver, address tokenContract, uint256 amount) private {
        // we snapshot current balance of this token
        uint256 ourStartingBalance = IERC20(tokenContract).balanceOf(receiver);

        IERC20(tokenContract).safeTransferFrom(sender, receiver, amount);

        // check what this particular ERC20 implementation actually gave us, since it doesn't
        // have to be at all related to the _amount
        uint256 actualEndingBalance = IERC20(tokenContract).balanceOf(receiver);

        uint256 expectedEndingBalance = ourStartingBalance + amount;
        // a very strange ERC20 may trigger this condition, if we didn't have this we would
        // underflow, so it's mostly just an error message printer
        require(
            actualEndingBalance > ourStartingBalance && actualEndingBalance == expectedEndingBalance,
            ICS20UnexpectedERC20Balance(expectedEndingBalance, actualEndingBalance)
        );
    }

    /// @notice Finds a contract in the foreign mapping, or creates a new IBCERC20 contract
    /// @notice This function will never return address(0)
    /// @param denom The denom to find or create the contract for
    /// @return The address of the erc20 contract
    function _findOrCreateERC20Address(IICS20TransferMsgs.Denom memory denom) private returns (address) {
        ICS20TransferStorage storage $ = _getICS20TransferStorage();

        // check if denom already has a foreign registered contract
        bytes32 denomID = ICS20Lib.getDenomIdentifier(denom);
        address erc20Contract = address($.ibcDenomContracts[denomID]);
        if (erc20Contract == address(0)) {
            // nothing exists, so we create new erc20 contract and register it in the mapping
            IBCERC20 ibcERC20 = new IBCERC20(this, $.escrow, denom);

            $.ibcDenomContracts[denomID] = ibcERC20;
            erc20Contract = address(ibcERC20);
        }

        return erc20Contract;
    }

    /// @notice Returns the storage of the ICS20Transfer contract
    function _getICS20TransferStorage() private pure returns (ICS20TransferStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := ICS20TRANSFER_STORAGE_SLOT
        }
    }

    /// @inheritdoc UUPSUpgradeable
    function _authorizeUpgrade(address) internal virtual override {
        address ics26Router = address(_getICS26Router());
        require(IIBCUUPSUpgradeable(ics26Router).isAdmin(_msgSender()), ICS20Unauthorized(_msgSender()));
    }

    function _getEscrow() private view returns (IEscrow) {
        return _getICS20TransferStorage().escrow;
    }

    function _getICS26Router() private view returns (IICS26Router) {
        return _getICS20TransferStorage().ics26Router;
    }

    modifier onlyRouter() {
        require(_msgSender() == address(_getICS26Router()), ICS20Unauthorized(_msgSender()));
        _;
    }
}
