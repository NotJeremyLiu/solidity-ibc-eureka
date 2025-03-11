package main

import (
	"fmt"
	"math/big"
	"os"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/srdtrk/solidity-ibc-eureka/e2e/v8/cmd/utils"
	"github.com/srdtrk/solidity-ibc-eureka/e2e/v8/types/erc20"

	"github.com/cosmos/solidity-ibc-eureka/abigen/skipeurekahandler"
)

func TransferFromEth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from-eth-to-cosmos [amount] [erc20-contract-address] [to-address]", // TODO: Better name???
		Short: "Send a transfer from Ethereum to Cosmos",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			// get args
			amountStr := args[0]
			transferAmount, ok := new(big.Int).SetString(amountStr, 10)
			if !ok {
				return fmt.Errorf("invalid amount: %s", amountStr)
			}
			erc20Address := ethcommon.HexToAddress(args[1])
			to := args[2]

			// get flags
			ethRPC, _ := cmd.Flags().GetString(FlagEthRPC)
			if ethRPC == "" {
				return fmt.Errorf("eth rpc flag not set")
			}

			ics20Str, _ := cmd.Flags().GetString(FlagIcs20Address)
			if ics20Str == "" {
				return fmt.Errorf("ics20 address flag not set")
			}
			// ics20Address := ethcommon.HexToAddress(ics20Str)

			ethPrivateKeyStr := os.Getenv(EnvEthPrivateKey)
			if ethPrivateKeyStr == "" {
				return fmt.Errorf("ETH_PRIVATE_KEY env var not set")
			}

			sourceClientID, _ := cmd.Flags().GetString(FlagSourceClientID)
			if sourceClientID == "" {
				return fmt.Errorf("source client flag not set")
			}

			// transferWithCallbacksMemo, _ := cmd.Flags().GetBool(FlagTransferWithCallbacksMemo)

			// Set up everything needed to send the transfer
			ethClient, err := ethclient.Dial(ethRPC)
			if err != nil {
				return fmt.Errorf("failed to connect to Ethereum RPC: %w", err)
			}

			// ics20Contract, err := ics20transfer.NewContract(ics20Address, ethClient)

			erc20Contract, err := erc20.NewContract(erc20Address, ethClient)

			ethChainID, err := ethClient.ChainID(ctx)
			if err != nil {
				return fmt.Errorf("failed to get chain id: %w", err)
			}
			ethPrivKey := utils.EthPrivateKeyFromHex(ethPrivateKeyStr)
			ethereumUserAddress := crypto.PubkeyToAddress(ethPrivKey.PublicKey)

			// Approve ICS20 contract to spend ERC20
			// TODO: Consider if we should query Permit2, so we don't have to do this every time 🤔
			txOpts := utils.GetTransactOpts(ctx, ethClient, ethChainID, ethPrivKey)
			txOpts.GasPrice = nil
			if IsVerbose(cmd) {
				fmt.Printf("Approve TransactOpts: %+v\n", txOpts)
			}
			tx, err := erc20Contract.Approve(txOpts, ethcommon.HexToAddress("0x92470162374A6D185758982356833d1aFfFd3b03"), transferAmount)
			if err != nil {
				return fmt.Errorf("approve tx call failed: %w", err)
			}
			receipt := utils.GetTxReciept(ctx, ethClient, tx.Hash())
			if receipt != nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
				return fmt.Errorf("approve tx unsuccessful (%s) %+v", tx.Hash().String(), receipt)
			} else if receipt == nil {
				cmd.Printf("Approve TX (%s) was not confirmed within time limit, but it was also not rejected. We'll continue anyway.\n", tx.Hash().String())
			}

			fmt.Printf("Approved ICS20 contract (%s) to spend ERC20 (%s) from %s\n", ethcommon.HexToAddress("0x92470162374A6D185758982356833d1aFfFd3b03"), erc20Address.Hex(), ethereumUserAddress.Hex())

			timeout := uint64(time.Now().Add(1 * time.Hour).Unix())
			// sendTransferMsg := ics20transfer.IICS20TransferMsgsSendTransferMsg{
			// 	Denom:            erc20Address,
			// 	Amount:           transferAmount,
			// 	Receiver:         to,
			// 	SourceClient:     sourceClientID,
			// 	DestPort:         "transfer",
			// 	TimeoutTimestamp: timeout,
			// 	Memo:             "",
			// }
			// if transferWithCallbacksMemo {
			// 	sendTransferMsg.Memo = `
			// 	{
			// 		"dest_callback": {
			// 			"address": "lom1nc5tatafv6eyq7llkr2gv50ff9e22mnf70qgjlv737ktmt4eswrqfufmpf"
			// 		},
			// 		"wasm": {
			// 			"contract": "lom14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s9fver0",
			// 			"msg": {
			// 				"execute_swap_operations": {
			// 					"operations": [
			// 						{
			// 							"mantra_swap": {
			// 								"pool_identifier": "",
			// 								"token_in_denom": "ibc/CA99FBBE626DB68588B2BB86750F0A95762B8822E82B82D4AB9654083E54DE5D",
			// 								"token_out_denom": "slbtc"
			// 							}
			// 						}
			// 					],
			// 					"minimum_receive": "1",
			// 					"receiver": "lom1nc5tatafv6eyq7llkr2gv50ff9e22mnf70qgjlv737ktmt4eswrqfufmpf"
			// 				}
			// 			}
			// 		}
			// 	}`
			// }

			txOpts = utils.GetTransactOpts(ctx, ethClient, ethChainID, ethPrivKey)
			txOpts.GasPrice = nil
			if IsVerbose(cmd) {
				fmt.Printf("SendTransfer TransactOpts: %+v\n", txOpts)
			}

			skipEurekaHandler, err := skipeurekahandler.NewEurekaHandler(ethcommon.HexToAddress("0x92470162374A6D185758982356833d1aFfFd3b03"), ethClient)
			if err != nil {
				return fmt.Errorf("failed to create skip eureka handler: %w", err)
			}

			tx, err = skipEurekaHandler.LombardTransfer(txOpts, transferAmount, skipeurekahandler.IEurekaHandlerTransferParams{
				Token:            erc20Address,
				Recipient:        to,
				SourceClient:     sourceClientID,
				DestPort:         "transfer",
				TimeoutTimestamp: timeout,
				Memo: `
				{
					"dest_callback": {
						"address": "lom1k9z7fvd3307ynapnk5gjffn706mz8qck502f7qwgt2a2a35kgqlqr0mc4j"
					},
					"wasm": {
						"contract": "lom1rqw2gr0ajp0ajrluftj0kut70kvhu32yj5are47hqmdfs6fyrcas9tlrl9",
						"msg": {
							"swap_and_action": {
								"user_swap": {
									"swap_exact_asset_in": {
										"swap_venue_name": "testnet-ledger-lbtc-convert",
										"operations": [
											{
												"pool": "",
												"denom_in": "ibc/CA99FBBE626DB68588B2BB86750F0A95762B8822E82B82D4AB9654083E54DE5D",
												"denom_out": "slbtc"
											}
										]
									}
								},
								"min_asset": {
									"native": {
										"denom": "slbtc",
										"amount": "1"
									}
								},
								"timeout_timestamp": 1741755711295674400,
								"post_swap_action": {
									"ibc_transfer": {
										"ibc_info": {
											"source_channel": "channel-0",
											"receiver": "bbn18g8eu8x2pj63tmt620lhllxkzy5ayrk7z36z98",
											"memo": "",
											"recover_address": "lom18g8eu8x2pj63tmt620lhllxkzy5ayrk7pm6seh"
										}
									}
								},
								"affiliates": []
							}
						}
					}
				}
			`,
			}, skipeurekahandler.IEurekaHandlerFees{
				RelayFee: big.NewInt(0),
				// TODO: Set this
				// RelayFeeRecipient: ethcommon.HexToAddress("0x0000000000000000000000000000000000000000"),
				QuoteExpiry: timeout,
			})
			if err != nil {
				return fmt.Errorf("send transfer tx unsuccessful\nerr: %w", err)
			}

			// tx, err = ics20Contract.SendTransfer(txOpts, sendTransferMsg)
			// if err != nil {
			// 	fmt.Printf("tx %+v\n", tx)
			// 	return fmt.Errorf("send transfer tx unsuccessful\nmsg %+v\nerr: %w", sendTransferMsg, err)
			// }
			receipt = utils.GetTxReciept(ctx, ethClient, tx.Hash())
			if receipt != nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
				return fmt.Errorf("send transfer tx (%s) unsuccessful %+v", tx.Hash().String(), receipt)
			} else if receipt == nil {
				cmd.Printf("Send transfer TX (%s) was not confirmed within time limit, but it was also not rejected. Please check the transaction in an explorer to verify the success.\n", tx.Hash().String())

			}

			fmt.Printf("Transfer sent from %s to %s of %d %s with tx hash %s\n", ethereumUserAddress.Hex(), to, transferAmount, erc20Address.Hex(), tx.Hash().Hex())

			return nil
		},
	}

	AddEthFlags(cmd)
	cmd.Flags().String(FlagSourceClientID, MockTendermintClientID, "Tendermint Client ID on Ethereum")
	cmd.Flags().Bool(FlagTransferWithCallbacksMemo, false, "Transfer with callbacks forwarding memo")

	return cmd
}
