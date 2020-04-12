import { RawLocation } from 'vue-router';

let installed = false;
let _Vue: any;

export default class Navigation {
  static install(Vue: any): void {
    if (installed && _Vue === Vue) return;
    installed = true;
    _Vue = Vue;

    Vue.prototype.$navigation = new Vue({
      data: {
        title: '',
        to: null,
      },
    });
  }
}


declare class CNavigation {
  title: string;
  to: RawLocation | null;
}

declare module 'vue/types/vue' {
  interface Vue {
    $navigation: CNavigation,
  }
}
