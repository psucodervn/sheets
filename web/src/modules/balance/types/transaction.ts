interface ITransactionUser {
  id: string;
  name: string;
  value: number;
}

interface ITransactionChange {
  value: number;
  [_: string]: any;
}

type TTransactionChanges = Record<string, ITransactionChange>;

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

export {
  ITransactionUser,
  ITransactionChange,
  TTransactionChanges,
  ITransaction,
};
