import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators';
import { ITransaction } from '@/modules/balance/types/transaction';
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
    if (res.success) {
      this.setTransactions({ transactions: res.data! });
    }
    return res.data;
  }

  @Mutation
  setTransactions(param: { transactions: ITransaction[] }) {
    this.transactions = param.transactions;
  }
}
