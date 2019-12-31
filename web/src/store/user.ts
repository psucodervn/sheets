import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators';
import { IUserBalance } from '@/model/user';
import { ApiUrls } from '@/constants/apis';

@Module({
  name: 'user',
  namespaced: true,
})
export class UserStore extends VuexModule {
  users: IUserBalance[] = [];

  @Action({ commit: 'setUsers', rawError: true })
  async fetchUsers() {
    const result = await fetch(ApiUrls.FetchUsers);
    return await result.json();
  }

  @Mutation
  async setUsers(users: IUserBalance[]) {
    this.users = users;
  }
}
