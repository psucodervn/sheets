import { RawLocation } from 'vue-router';

export default class Navigation {
  static installed = false;

  static install(Vue: any): void {
    if (this.installed) return;
    this.installed = true;

    const nav = new Vue({
      data: {
        title: '',
        from: null,
        parent: null,
      },
    });
    Vue.$navigation = nav;
    Vue.prototype.$navigation = nav;
  }
}

declare class CNavigation {
  title: string;
  from: RawLocation | null;
  parent: RawLocation | null;
}

declare module 'vue/types/vue' {
  interface Vue {
    $navigation: CNavigation;
  }

  interface VueConstructor {
    $navigation: CNavigation;
  }
}
