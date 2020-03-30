export interface IBalance {
  value: number;
}

export interface IUser {
  name: string;
  balance: IBalance;
}

export class User implements IUser {
  name: string;
  balance: IBalance;

  constructor(name: string, balance: Balance) {
    this.name = name;
    this.balance = balance;
  }
}

export class Balance implements IBalance {
  value: number;

  constructor(value: number) {
    this.value = value;
  }
}
