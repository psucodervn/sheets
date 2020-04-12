import { Action, Module, Mutation, MutationAction, VuexModule } from 'vuex-module-decorators';
import { ApiUrls } from '@/constants/apis';
import { IUserPoint } from '@/model/point';
import { Month } from '@/types/datetime';
import { toMonth } from '@/utils/datetime';

@Module({
  name: 'point',
  namespaced: true,
})
export class PointStore extends VuexModule {
  users: IUserPoint[] = [];
  month: Month = toMonth(new Date());

  @Action({ commit: 'setPoints', rawError: true })
  async fetchPoints(params: { month: Month }) {
    const { year, month } = params.month;
    const result = await fetch(`${ApiUrls.FetchPoints}?year=${year}&month=${month}`);
    return (await result.json()).data;
  }

  @Mutation
  async setPoints(userPoints: IUserPoint[]) {
    this.users = userPoints;
  }

  @MutationAction({ mutate: ['month'] })
  async setMonth(month: Month) {
    return { month };
  }
}
