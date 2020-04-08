import { Action, Module, Mutation, MutationAction, VuexModule } from 'vuex-module-decorators';
import { ApiUrls } from '@/constants/apis';
import { IUserPoint } from '@/model/point';
import { Month, Months } from '@/constants/datetime';

@Module({
  name: 'point',
  namespaced: true,
})
export class PointStore extends VuexModule {
  users: IUserPoint[] = [];
  month: Month = Months[new Date().getUTCMonth()];
  year: number = new Date().getFullYear();

  @Action({ commit: 'setPoints', rawError: true })
  async fetchPoints(params: { year: number, month: number }) {
    const { year, month } = params;
    const result = await fetch(`${ApiUrls.FetchPoints}?year=${year}&month=${month}`);
    return (await result.json()).data;
  }

  @Mutation
  async setPoints(userPoints: IUserPoint[]) {
    this.users = userPoints;
  }

  @MutationAction({ mutate: ['year'] })
  async setYear(year: number) {
    return { year };
  }

  @MutationAction({ mutate: ['month'] })
  async setMonth(month: Month) {
    return { month };
  }
}
