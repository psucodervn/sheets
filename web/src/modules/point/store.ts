import { Action, Module, Mutation, MutationAction, VuexModule } from 'vuex-module-decorators';
import { IUserPoint } from '@/model/point';
import { Month, TimeRange } from '@/types/datetime';
import { toMonth } from '@/utils/datetime';
import { Vue } from 'vue-property-decorator';
import { ITransaction } from '@/modules/balance/types/transaction';

@Module({
  name: 'point',
  namespaced: true,
})
export default class PointStore extends VuexModule {
  users: IUserPoint[] = [];
  month: Month = toMonth(new Date());

  @Action({ commit: 'setPoints', rawError: true })
  async fetchPoints(params: { month: Month }) {
    const { year, month } = params.month;
    const res = await Vue.$api.get<ITransaction[]>('/points', {
      params: { year, month },
    });
    if (!res.success) throw new Error(res.message);
    return res.data!;
  }

  @Mutation
  async setPoints(userPoints: IUserPoint[]) {
    this.users = userPoints;
  }

  @MutationAction({ mutate: ['month'] })
  async setMonth(month: Month) {
    return { month };
  }

  @Action({ rawError: true })
  async fetchReport(params: { range: TimeRange }) {
    const { from, to } = params.range;
    from.toISOString();
    const res = await Vue.$api.get<IUserPoint[]>('/report', {
      params: { from: from.toISOString(), to: to.toISOString() },
    });
    if (!res.success) throw new Error(res.message);
    return res.data!;
  }
}
