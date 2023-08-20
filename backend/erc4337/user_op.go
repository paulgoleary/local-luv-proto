package erc4337

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/paulgoleary/local-luv-proto/chain"
	"github.com/paulgoleary/local-luv-proto/crypto"
	"github.com/stackup-wallet/stackup-bundler/pkg/userop"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"github.com/umbracle/ethgo/wallet"
	"math/big"
)

// for Mumbai, Polygon Mainnet, ETH mainnet, ...
var DefaultEntryPoint = ethgo.HexToAddress("0x5ff137d4b0fdcd49dca30c7cf57e578a026d2789")
var DefaultChainId = big.NewInt(137) // Polygon mainnet for now...

var abiExec, _ = abi.NewMethod("function execute(address to, uint256 value, bytes data)")

func makeExecute(toAddr ethgo.Address, value *big.Int, m *abi.Method, args ...interface{}) (enc []byte, err error) {
	if enc, err = m.Encode(args); err == nil {
		return abiExec.Encode([]interface{}{toAddr, value, enc})
	}
	return
}

var mintMethod, _ = abi.NewMethod("function mint(address sender, uint256 amount)")
var approveMethod, _ = abi.NewMethod("function approve(address spender, uint256 amount) external returns (bool)")
var withdrawToMethod, _ = abi.NewMethod("function withdrawTo(address account, uint256 amount) external returns (bool)")

var DefaultInitCodeGas = big.NewInt(300_000)
var DefaultMintGasLimit = big.NewInt(200_000)
var DefaultApproveGasLimit = big.NewInt(200_000)
var DefaultWithdrawToGasLimit = big.NewInt(200_000)

func makeBaseOp(nonce *big.Int, owner, sender ethgo.Address, callGasLimit, maxFeePerGas *big.Int, callData []byte) (op *userop.UserOperation, err error) {
	var initCode []byte

	if nonce.Int64() == 0 {
		if initCode, err = MakeDefaultInitCode(owner); err != nil {
			return
		}
	}

	vGasLimit := new(big.Int).Add(big.NewInt(150_000), DefaultInitCodeGas)
	opData := map[string]any{
		"sender":               sender.String(),
		"nonce":                nonce,
		"initCode":             hexutil.Encode(initCode),
		"callData":             hexutil.Encode(callData),
		"callGasLimit":         callGasLimit,
		"verificationGasLimit": vGasLimit,
		"maxFeePerGas":         maxFeePerGas,
		"maxPriorityFeePerGas": "0x0",
		"paymasterAndData":     "0x",
		"preVerificationGas":   big.NewInt(100_000),
		"signature":            "0x00",
	}
	op, err = userop.New(opData)
	return
}

func UserOpMint(nonce *big.Int, owner, sender, mintTargetAddr, toAddr ethgo.Address, amt *big.Int) (*userop.UserOperation, error) {
	if callData, err := makeExecute(mintTargetAddr, big.NewInt(0), mintMethod, toAddr, amt); err != nil {
		return nil, err
	} else {
		return makeBaseOp(nonce, owner, sender, DefaultMintGasLimit, big.NewInt(2_000_000_000), callData)
	}
}

func UserOpApprove(nonce *big.Int, owner, sender, targetAddr, spender ethgo.Address, amt *big.Int) (*userop.UserOperation, error) {
	if callData, err := makeExecute(targetAddr, big.NewInt(0), approveMethod, spender, amt); err != nil {
		return nil, err
	} else {
		return makeBaseOp(nonce, owner, sender, DefaultApproveGasLimit, big.NewInt(2_000_000_000), callData)
	}
}

func UserOpWithdrawTo(nonce *big.Int, owner, sender, targetAddr, toAddr ethgo.Address, amt *big.Int) (*userop.UserOperation, error) {
	if callData, err := makeExecute(targetAddr, big.NewInt(0), withdrawToMethod, toAddr, amt); err != nil {
		return nil, err
	} else {
		return makeBaseOp(nonce, owner, sender, DefaultWithdrawToGasLimit, big.NewInt(2_000_000_000), callData)
	}
}

func UserOpSeal(op *userop.UserOperation, chainId *big.Int, k *chain.EcdsaKey) (*userop.UserOperation, error) {
	opHash := op.GetUserOpHash(common.Address(DefaultEntryPoint), chainId)
	opEthHash := crypto.EthSignedMessageHash(opHash.Bytes())
	if sig, err := crypto.Sign(k.SK, opEthHash); err != nil {
		return nil, err
	} else {
		sig[64] += 27
		op.Signature = sig
	}
	return op, nil
}

func UserOpEcrecover(op *userop.UserOperation, chainId *big.Int) (ethgo.Address, error) {
	opHash := op.GetUserOpHash(common.Address(DefaultEntryPoint), chainId)
	opEthHash := crypto.EthSignedMessageHash(opHash.Bytes())
	return wallet.Ecrecover(opEthHash, op.Signature)
}
