import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators';
import { ApiUrls } from '@/constants/apis';
import { IUserPoint } from '@/model/point';

@Module({
  name: 'point',
  namespaced: true,
})
export class PointStore extends VuexModule {
  users: IUserPoint[] = [];

  @Action({ commit: 'setPoints', rawError: true })
  async fetchPoints() {
    const result = await fetch(ApiUrls.FetchPoints);
    return (await result.json()).data;
  }

  @Mutation
  async setPoints(userPoints: IUserPoint[]) {
    this.users = userPoints;
  }
}
