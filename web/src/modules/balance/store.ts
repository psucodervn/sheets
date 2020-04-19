import { Action, Module, VuexModule } from "vuex-module-decorators";
import api from "@/utils/api";
import { ITransaction } from "@/modules/balance/types/transaction";

@Module({
  name: "balance",
  namespaced: true
})
export default class BalanceStore extends VuexModule {
  @Action({ rawError: true })
  async fetchTransactions() {
    const res = await api.get<ITransaction[]>("/balance/transactions");
    return res.data;
  }
}
