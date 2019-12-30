export interface IUser {
  name: string;
}

export interface IBalance {
  value: number;
}

export interface IUserBalance {
  user: IUser;
  balance: IBalance;
}

export class User implements IUser {
  name: string;

  constructor(name: string) {
    this.name = name;
  }
}

export class Balance implements IBalance {
  value: number;

  constructor(value: number) {
    this.value = value;
  }
}

export class UserBalance implements IUserBalance {
  user: IUser;
  balance: IBalance;

  constructor(user: IUser, balance: IBalance) {
    this.user = user;
    this.balance = balance;
  }
}
