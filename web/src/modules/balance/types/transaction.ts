interface ITransactionUser {
  id: string;
  name: string;
  value: number;
}

type TTransactionChanges = Record<string, number>;

interface ITransaction {
  id: string;
  time: Date;
  description: string;
  summary: string;
  totalValue: number;
  senders: ITransactionUser[];
  receivers: ITransactionUser[];
  changes: TTransactionChanges;
}

export { ITransactionUser, TTransactionChanges, ITransaction };
