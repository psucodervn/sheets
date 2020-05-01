interface ITransactionUser {
  id: string;
  value: number;
  name?: string;
  text?: string;
  percent?: number;
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
  value: number;
  payers: ITransactionUser[];
  participants: ITransactionUser[];
  changes: TTransactionChanges;
}
type TTransactionNew = Omit<ITransaction, 'id'>;

enum ESplitOption {
  Equal,
  Ratio,
  Custom,
}

export {
  ITransactionUser,
  ITransactionChange,
  TTransactionChanges,
  ITransaction,
  TTransactionNew,
  ESplitOption,
};
