// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgApproveLoan } from "./types/loan/loan/tx";
import { MsgRequestLoan } from "./types/loan/loan/tx";
import { MsgRepayLoan } from "./types/loan/loan/tx";
import { MsgLiquidateLoan } from "./types/loan/loan/tx";


export { MsgApproveLoan, MsgRequestLoan, MsgRepayLoan, MsgLiquidateLoan };

type sendMsgApproveLoanParams = {
  value: MsgApproveLoan,
  fee?: StdFee,
  memo?: string
};

type sendMsgRequestLoanParams = {
  value: MsgRequestLoan,
  fee?: StdFee,
  memo?: string
};

type sendMsgRepayLoanParams = {
  value: MsgRepayLoan,
  fee?: StdFee,
  memo?: string
};

type sendMsgLiquidateLoanParams = {
  value: MsgLiquidateLoan,
  fee?: StdFee,
  memo?: string
};


type msgApproveLoanParams = {
  value: MsgApproveLoan,
};

type msgRequestLoanParams = {
  value: MsgRequestLoan,
};

type msgRepayLoanParams = {
  value: MsgRepayLoan,
};

type msgLiquidateLoanParams = {
  value: MsgLiquidateLoan,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgApproveLoan({ value, fee, memo }: sendMsgApproveLoanParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgApproveLoan: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgApproveLoan({ value: MsgApproveLoan.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgApproveLoan: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRequestLoan({ value, fee, memo }: sendMsgRequestLoanParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRequestLoan: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRequestLoan({ value: MsgRequestLoan.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRequestLoan: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRepayLoan({ value, fee, memo }: sendMsgRepayLoanParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRepayLoan: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRepayLoan({ value: MsgRepayLoan.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRepayLoan: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgLiquidateLoan({ value, fee, memo }: sendMsgLiquidateLoanParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgLiquidateLoan: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgLiquidateLoan({ value: MsgLiquidateLoan.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgLiquidateLoan: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgApproveLoan({ value }: msgApproveLoanParams): EncodeObject {
			try {
				return { typeUrl: "/loan.loan.MsgApproveLoan", value: MsgApproveLoan.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgApproveLoan: Could not create message: ' + e.message)
			}
		},
		
		msgRequestLoan({ value }: msgRequestLoanParams): EncodeObject {
			try {
				return { typeUrl: "/loan.loan.MsgRequestLoan", value: MsgRequestLoan.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRequestLoan: Could not create message: ' + e.message)
			}
		},
		
		msgRepayLoan({ value }: msgRepayLoanParams): EncodeObject {
			try {
				return { typeUrl: "/loan.loan.MsgRepayLoan", value: MsgRepayLoan.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRepayLoan: Could not create message: ' + e.message)
			}
		},
		
		msgLiquidateLoan({ value }: msgLiquidateLoanParams): EncodeObject {
			try {
				return { typeUrl: "/loan.loan.MsgLiquidateLoan", value: MsgLiquidateLoan.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgLiquidateLoan: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			LoanLoan: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;