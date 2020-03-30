import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators';
import { IUser } from '@/model/user';
import { ApiUrls } from '@/constants/apis';

@Module({
  name: 'user',
  namespaced: true,
})
export class UserStore extends VuexModule {
  users: IUser[] = [];

  @Action({ commit: 'setUsers', rawError: true })
  async fetchUsers() {
    const result = await fetch(ApiUrls.FetchUsers);
    return (await result.json()).data;
  }

  @Mutation
  async setUsers(users: IUser[]) {
    this.users = users;
  }
}
