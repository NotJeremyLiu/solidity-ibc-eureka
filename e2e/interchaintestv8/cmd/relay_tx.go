package main

import (
	"context"
	"encoding/hex"
	"os"
	"strings"

	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/gogoproto/proto"
	"github.com/spf13/cobra"
	"github.com/srdtrk/solidity-ibc-eureka/e2e/v8/cmd/utils"
	relayertypes "github.com/srdtrk/solidity-ibc-eureka/e2e/v8/types/relayer"

	clientTx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmosTx "github.com/cosmos/cosmos-sdk/types/tx"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	ibctypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
)

func RelayTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relay-tx [txHash]",
		Short: "Relay a transaction (currently only from eth to cosmos)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			// Get args
			txHash := args[0]

			var err error
			if strings.HasPrefix(txHash, "0x") {
				err = relayFromEthToCosmos(ctx, cmd, txHash)
			} else {
				err = relayFromCosmosToEth(ctx, cmd, txHash)
			}

			if err != nil {
				return err
			}

			return nil
		},
	}

	AddEthFlags(cmd)
	AddCosmosFlags(cmd)
	cmd.Flags().String(FlagEthChainID, DefaultEthChainID, "Ethereum Chain ID number (e.g. 11155111)")
	cmd.Flags().String(FlagCosmosClientIDOnEth, "", "target client id of the cosmos client on ethereum (used for relaying from cosmos to eth)")
	cmd.Flags().String(FlagEthClientIDOnCosmos, "", "target client id of the ethereum client on cosmos (used for relaying from eth to cosmos)")

	return cmd
}

func relayFromEthToCosmos(ctx context.Context, cmd *cobra.Command, txHashHexStr string) error {
	fmt.Println("Relaying from Ethereum to Cosmos")

	// get the flags we need
	cosmosRPC, _ := cmd.Flags().GetString(FlagCosmosRPC)
	if cosmosRPC == "" {
		return fmt.Errorf("cosmos-rpc flag not set")
	}
	cosmosChainID, _ := cmd.Flags().GetString(FlagCosmosChainID)
	if cosmosChainID == "" {
		return fmt.Errorf("cosmos-chain-id flag not set")
	}

	targetClientID, _ := cmd.Flags().GetString(FlagEthClientIDOnCosmos)
	if targetClientID == "" {
		targetClientID = MockEthClientID
	}

	cosmosGrpcAddress, _ := cmd.Flags().GetString(FlagCosmosGRPC)
	if cosmosGrpcAddress == "" {
		return fmt.Errorf("cosmos-grpc flag not set")
	}

	ethChainID, _ := cmd.Flags().GetString(FlagEthChainID)
	if ethChainID == "" {
		return fmt.Errorf("eth chain id flag not set")
	}

	// Set up everything we need to relay
	// db := dbm.NewMemDB()
	// app := simapp.NewUnitTestSimApp(log.NewNopLogger(), db, nil, false, simtestutil.EmptyAppOptions{}, nil)

	cosmosRelayerPrivateKeyStr := os.Getenv(EnvRelayerWallet)
	if cosmosRelayerPrivateKeyStr == "" {
		return fmt.Errorf("%s env var not set", EnvRelayerWallet)
	}
	// cosmosRelayerPrivateKey, err := utils.CosmosPrivateKeyFromHex(cosmosRelayerPrivateKeyStr)
	// if err != nil {
	// 	return err
	// }

	privKeyBytes, err := hex.DecodeString(cosmosRelayerPrivateKeyStr)
	if err != nil {
		return fmt.Errorf("Error decoding private key hex, not adding signer info: %v", err)
	}
	cosmosRelayerPrivateKey := &secp256k1.PrivKey{}
	if err := cosmosRelayerPrivateKey.UnmarshalAmino(privKeyBytes); err != nil {
		return fmt.Errorf("Error unmarshaling private key bytes, not adding signer info: %v", err)
	}

	// cosmosAddress := sdk.AccAddress(cosmosRelayerPrivateKey.PubKey().Address())

	grpcConn, err := utils.GetTLSGRPC(cosmosGrpcAddress)
	if err != nil {
		return err
	}
	authQueryClient := authtypes.NewQueryClient(grpcConn)
	txClient := cosmosTx.NewServiceClient(grpcConn)
	protoCodec := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
	txConfig := authTx.NewTxConfig(protoCodec, authTx.DefaultSignModes)

	txHash := ethcommon.HexToHash(txHashHexStr)

	relayerClient, err := GetRelayerClient(RelayerURL)
	if err != nil {
		return err
	}

	resp, err := relayerClient.RelayByTx(ctx, &relayertypes.RelayByTxRequest{
		SrcChain:       ethChainID,
		DstChain:       cosmosChainID,
		SourceTxIds:    [][]byte{txHash.Bytes()},
		TargetClientId: targetClientID,
	})
	if err != nil {
		return fmt.Errorf("failed to relayed tx: %w", err)
	}

	// Extract messages from the response (cosmos specific)
	var txBody txtypes.TxBody
	if err := proto.Unmarshal(resp.Tx, &txBody); err != nil {
		return err
	}

	fmt.Println("txBody: ", txBody)

	if len(txBody.Messages) == 0 {
		return fmt.Errorf("no messages to relay")
	}

	fmt.Println("txBody.Messages: ", txBody.Messages)
	fmt.Println("Len(txBody.Messages): ", len(txBody.Messages))

	// var msgs []sdk.Msg
	// for _, msg := range txBody.Messages {
	// 	var sdkMsg sdk.Msg
	// 	if err := app.InterfaceRegistry().UnpackAny(msg, &sdkMsg); err != nil {
	// 		return err
	// 	}

	// 	msgs = append(msgs, sdkMsg)
	// }

	var msgs []sdk.Msg
	for _, msg := range txBody.Messages {
		if msg.TypeUrl == "/ibc.core.channel.v2.MsgRecvPacket" {
			receiveMsg := &ibctypes.MsgRecvPacket{}
			if err := proto.Unmarshal(msg.Value, receiveMsg); err != nil {
				return err
			}
			msgs = append(msgs, receiveMsg)
		}
	}

	fmt.Println("msgs: ", msgs)

	fmt.Println("msg: ", msgs[0].String())

	bech32Address := "lom127tlxptdqt0pe25mq5cf8ju68lsa9yxynn90rj"

	// Get account for sequence and account number
	// accountClient := accounttypes.NewQueryClient(grpcConn)
	// accountRes, err := accountClient.AccountInfo(ctx, &accounttypes.QueryAccountInfoRequest{Address: bech32Address})
	// if err != nil {
	// 	return fmt.Errorf("failed to get account info: %w", err)
	// }
	// fmt.Println("AccountRes: ", accountRes)

	accReq := &authtypes.QueryAccountRequest{
		Address: bech32Address,
	}
	acc, err := authQueryClient.Account(ctx, accReq)
	if err != nil {
		return fmt.Errorf("failed to get account info: %w", err)
	}
	account := authtypes.BaseAccount{}
	if err := account.Unmarshal(acc.Account.Value); err != nil {
		return fmt.Errorf("failed to unmarshal account: %w", err)
	}

	txBuilder := txConfig.NewTxBuilder()

	//txBuilder := app.TxConfig().NewTxBuilder()
	txBuilder.SetGasLimit(2000000)
	txBuilder.SetMsgs(msgs...)
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin("ulom", 200000000)))

	// sigV2 := signing.SignatureV2{
	// 	PubKey: cosmosRelayerPrivateKey.PubKey(),
	// 	Data: &signing.SingleSignatureData{
	// 		//SignMode:  signing.SignMode(app.TxConfig().SignModeHandler().DefaultMode()),
	// 		SignMode:  signing.SignMode(txConfig.SignModeHandler().DefaultMode()),
	// 		Signature: nil,
	// 	},
	// 	Sequence: accountRes.Info.Sequence,
	// }

	signerData := authsigning.SignerData{
		ChainID:       cosmosChainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
		PubKey:        cosmosRelayerPrivateKey.PubKey(),
		Address:       bech32Address,
	}
	sigV2, err := clientTx.SignWithPrivKey(
		context.Background(),
		signing.SignMode(txConfig.SignModeHandler().DefaultMode()),
		signerData,
		txBuilder,
		cosmosRelayerPrivateKey,
		txConfig,
		account.GetSequence(),
	)
	if err != nil {
		return fmt.Errorf("failed to sign with priv key: %w", err)
	}
	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return fmt.Errorf("failed to set signature: %w", err)
	}

	// signerData := xauthsigning.SignerData{
	// 	Address:       bech32Address,
	// 	ChainID:       cosmosChainID,
	// 	AccountNumber: accountRes.Info.AccountNumber,
	// }
	// sigV2, err = tx.SignWithPrivKey(
	// 	ctx,
	// 	signing.SignMode(txConfig.SignModeHandler().DefaultMode()),
	// 	signerData,
	// 	txBuilder,
	// 	cosmosRelayerPrivateKey,
	// 	txConfig,
	// 	accountRes.Info.Sequence,
	// )
	// if err != nil {
	// 	return fmt.Errorf("failed to sign with priv key: %w", err)
	// }
	// err = txBuilder.SetSignatures(sigV2)
	// if err != nil {
	// 	return fmt.Errorf("failed to set signature: %w", err)
	// }

	// Generated Protobuf-encoded bytes.
	txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return fmt.Errorf("failed to encode tx: %w", err)
	}

	// txClient := txtypes.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		ctx,
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return fmt.Errorf("failed to broadcast tx: %w", err)
	}
	if grpcRes.TxResponse.Code != 0 {
		return fmt.Errorf("tx failed with code %d: %+v", grpcRes.TxResponse.Code, grpcRes.TxResponse)
	}

	fmt.Printf("Tx relayed successfully with hash %s\n", grpcRes.TxResponse.TxHash)
	rpcTxURL := cosmosRPC + "/tx?hash=0x" + grpcRes.TxResponse.TxHash
	fmt.Printf("Find full event logs here: %s\n", rpcTxURL)

	return nil
}

func relayFromCosmosToEth(ctx context.Context, cmd *cobra.Command, txHash string) error {
	fmt.Println("Relaying from Cosmos to Ethereum")
	txHashBz, err := hex.DecodeString(txHash)
	if err != nil {
		return err
	}

	// get the flags we need
	targetClientID, _ := cmd.Flags().GetString(FlagCosmosClientIDOnEth)
	if targetClientID == "" {
		targetClientID = MockTendermintClientID
	}

	cosmosChainID, _ := cmd.Flags().GetString(FlagCosmosChainID)
	if cosmosChainID == "" {
		return fmt.Errorf("cosmos-chain-id flag not set")
	}

	ethChainID, _ := cmd.Flags().GetString(FlagEthChainID)
	if ethChainID == "" {
		return fmt.Errorf("eth chain id flag not set")
	}

	ethRPC, _ := cmd.Flags().GetString(FlagEthRPC)
	if ethRPC == "" {
		return fmt.Errorf("eth-rpc flag not set")
	}
	ics26AddressStr, _ := cmd.Flags().GetString(FlagIcs26Address)
	if ics26AddressStr == "" {
		return fmt.Errorf("ics26-address flag not set")
	}

	ethPrivateKeyStr := os.Getenv(EnvEthPrivateKey)
	if ethPrivateKeyStr == "" {
		return fmt.Errorf("ETH_PRIVATE_KEY env var not set")
	}

	ethPrivKey := utils.EthPrivateKeyFromHex(ethPrivateKeyStr)

	// set up everything we need to relay
	ics26Address := ethcommon.HexToAddress(ics26AddressStr)
	ethClient, err := ethclient.Dial(ethRPC)
	if err != nil {
		return fmt.Errorf("failed to dial eth client: %w", err)
	}

	relayerClient, err := GetRelayerClient(RelayerURL)
	if err != nil {
		return err
	}

	ethChainIDBigInt, err := ethClient.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chain id: %w", err)
	}

	resp, err := relayerClient.RelayByTx(ctx, &relayertypes.RelayByTxRequest{
		SrcChain:       cosmosChainID,
		DstChain:       ethChainID,
		SourceTxIds:    [][]byte{txHashBz},
		TargetClientId: targetClientID,
	})
	if err != nil {
		return fmt.Errorf("failed to relayed tx: %w", err)
	}

	txOpts := utils.GetTransactOpts(ctx, ethClient, ethChainIDBigInt, ethPrivKey)
	if IsVerbose(cmd) {
		fmt.Printf("TransactOpts: %+v\n", txOpts)
	}

	unsignedTx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{
		ChainID:   ethChainIDBigInt,
		Nonce:     txOpts.Nonce.Uint64(),
		GasTipCap: txOpts.GasTipCap,
		GasFeeCap: txOpts.GasFeeCap,
		Gas:       700_000,
		To:        &ics26Address,
		Data:      resp.Tx,
	})

	signedTx, err := txOpts.Signer(txOpts.From, unsignedTx)
	if err != nil {
		return fmt.Errorf("failed to sign tx: %w", err)
	}

	err = ethClient.SendTransaction(ctx, signedTx)
	if err != nil {
		return fmt.Errorf("failed to send tx: %w", err)
	}

	receipt := utils.GetTxReciept(ctx, ethClient, signedTx.Hash())
	if receipt != nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("relay tx unsuccessful (%s) %+v", signedTx.Hash().String(), receipt)
	} else if receipt == nil {
		cmd.Printf("Relay TX (%s) was not confirmed within time limit, but it was also not rejected. Please check the transaction in an explorer to verify the success.\n", signedTx.Hash().String())
	}
	fmt.Printf("Receipt: %+v\n", receipt)

	fmt.Printf("Tx relayed successfully with hash %s\n", signedTx.Hash().String())

	return nil
}

// GetGRPCClient returns a gRPC client for the relayer.
func GetRelayerClient(grpcAddr string) (relayertypes.RelayerServiceClient, error) {
	conn, err := utils.GetTLSGRPC(grpcAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server for relayer: %w", err)
	}

	return relayertypes.NewRelayerServiceClient(conn), nil
}
