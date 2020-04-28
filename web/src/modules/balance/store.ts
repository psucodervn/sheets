import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators';
import {
  ITransaction,
  TTransactionChanges,
} from '@/modules/balance/types/transaction';
import { Vue } from 'vue-property-decorator';
import { IUser } from '@/model/user';

@Module({
  name: 'balance',
  namespaced: true,
})
export default class BalanceStore extends VuexModule {
  transactions: ITransaction[] = [];
  users: IUser[] = [];

  @Action({ commit: 'setUsers', rawError: true })
  async fetchUsers() {
    const res = await Vue.$api.get<IUser[]>('/users');
    if (!res.success) throw new Error(res.message);
    return res.data!;
  }

  @Mutation
  async setUsers(users: IUser[]) {
    this.users = users;
  }

  @Action({ rawError: true })
  async fetchTransactions() {
    const res = await Vue.$api.get<ITransaction[]>('/balance/transactions');
    let transactions: ITransaction[] = [];
    if (res.success) {
      transactions = res.data!.map(tx => {
        const changes: TTransactionChanges = {};
        for (const u of tx.senders) {
          if (!changes[u.name]) changes[u.name] = { value: 0 };
          changes[u.name].value += u.value;
        }
        for (const u of tx.receivers) {
          if (!changes[u.name]) changes[u.name] = { value: 0 };
          changes[u.name].value -= u.value;
        }
        return { ...tx, changes };
      });
      this.setTransactions({ transactions });
    }
    return transactions;
  }

  @Mutation
  setTransactions(param: { transactions: ITransaction[] }) {
    this.transactions = param.transactions;
  }
}
