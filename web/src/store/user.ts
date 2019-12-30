import { Action, getModule, Module, Mutation, VuexModule } from 'vuex-module-decorators';
import { IUser, IUserBalance } from '@/model/user';
import { ApiUrls } from '@/constants/apis';
import store from '@/store/index';

@Module({
  dynamic: true,
  store: store,
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

export const UserModule = getModule(UserStore);
store.registerModule('user', UserStore, { preserveState: true });
