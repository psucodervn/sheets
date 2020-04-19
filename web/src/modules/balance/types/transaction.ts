interface ITransactionUser {
  id: string;
  name: string;
  value: number;
}

interface ITransaction {
  id: string;
  time: Date;
  description: string;
  summary: string;
  totalValue: number;
  senders: ITransactionUser[];
  receivers: ITransactionUser[];
}

export { ITransactionUser, ITransaction };
