// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgSetName } from "./types/nameservice/nameservice/tx";
import { MsgBuyName } from "./types/nameservice/nameservice/tx";
import { MsgDeleteName } from "./types/nameservice/nameservice/tx";


export { MsgSetName, MsgBuyName, MsgDeleteName };

type sendMsgSetNameParams = {
  value: MsgSetName,
  fee?: StdFee,
  memo?: string
};

type sendMsgBuyNameParams = {
  value: MsgBuyName,
  fee?: StdFee,
  memo?: string
};

type sendMsgDeleteNameParams = {
  value: MsgDeleteName,
  fee?: StdFee,
  memo?: string
};


type msgSetNameParams = {
  value: MsgSetName,
};

type msgBuyNameParams = {
  value: MsgBuyName,
};

type msgDeleteNameParams = {
  value: MsgDeleteName,
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
		
		async sendMsgSetName({ value, fee, memo }: sendMsgSetNameParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSetName: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSetName({ value: MsgSetName.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSetName: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgBuyName({ value, fee, memo }: sendMsgBuyNameParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgBuyName: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgBuyName({ value: MsgBuyName.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgBuyName: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeleteName({ value, fee, memo }: sendMsgDeleteNameParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeleteName: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeleteName({ value: MsgDeleteName.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeleteName: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgSetName({ value }: msgSetNameParams): EncodeObject {
			try {
				return { typeUrl: "/nameservice.nameservice.MsgSetName", value: MsgSetName.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSetName: Could not create message: ' + e.message)
			}
		},
		
		msgBuyName({ value }: msgBuyNameParams): EncodeObject {
			try {
				return { typeUrl: "/nameservice.nameservice.MsgBuyName", value: MsgBuyName.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgBuyName: Could not create message: ' + e.message)
			}
		},
		
		msgDeleteName({ value }: msgDeleteNameParams): EncodeObject {
			try {
				return { typeUrl: "/nameservice.nameservice.MsgDeleteName", value: MsgDeleteName.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeleteName: Could not create message: ' + e.message)
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
			NameserviceNameservice: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;