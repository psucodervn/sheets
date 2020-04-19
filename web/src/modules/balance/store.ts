import { Action, Module, Mutation, VuexModule } from "vuex-module-decorators";
import api from "@/utils/api";
import { ITransaction } from "@/modules/balance/types/transaction";

@Module({
  name: "balance",
  namespaced: true
})
export default class BalanceStore extends VuexModule {
  transactions: ITransaction[] = [];

  @Action({ rawError: true })
  async fetchTransactions() {
    const res = await api.get<ITransaction[]>("/balance/transactions");
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
