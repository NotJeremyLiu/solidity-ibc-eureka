// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { ERC20 } from "@openzeppelin/token/ERC20/ERC20.sol";
import { IICS20Transfer } from "../interfaces/IICS20Transfer.sol";
import { Ownable } from "@openzeppelin/access/Ownable.sol";
import { IIBCERC20 } from "../interfaces/IIBCERC20.sol";
import { IEscrow } from "../interfaces/IEscrow.sol";
import { ICS20Lib } from "../utils/ICS20Lib.sol";

contract IBCERC20 is IIBCERC20, ERC20, Ownable {
    /// @notice The full IBC denom path for this token
    ICS20Lib.Denom private _denom;
    /// @notice The escrow contract address
    IEscrow private immutable ESCROW;

    constructor(
        IICS20Transfer owner_,
        IEscrow escrow_,
        bytes32 denomID_,
        ICS20Lib.Denom memory denom_
    )
        // TODO: Was there something I was supposed to be using instead of encodePacked?
        ERC20(string(abi.encodePacked(denomID_)), denom_.base)
        Ownable(address(owner_))
    {
        _denom = denom_;
        ESCROW = escrow_;
    }

    /// @inheritdoc IIBCERC20
    function fullDenom() public view returns (ICS20Lib.Denom memory) {
        return _denom;
    }

    /// @inheritdoc IIBCERC20
    function mint(uint256 amount) external onlyOwner {
        _mint(address(ESCROW), amount);
    }

    /// @inheritdoc IIBCERC20
    function burn(uint256 amount) external onlyOwner {
        _burn(address(ESCROW), amount);
    }
}
