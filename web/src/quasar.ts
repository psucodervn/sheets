import Vue from 'vue';

import './styles/quasar.scss';
import '@quasar/extras/material-icons/material-icons.css';
import {
  QBtn,
  QDrawer,
  QFooter,
  QHeader,
  QIcon,
  QItem,
  QItemLabel,
  QItemSection,
  QLayout,
  QList,
  QPage,
  QPageContainer,
  QRouteTab,
  QTab,
  QTabs,
  QToolbar,
  QToolbarTitle,
  Quasar,
} from 'quasar';

Vue.use(Quasar, {
  config: {
    dark: 'off',
  },
  components: {
    QLayout,
    QHeader,
    QFooter,
    QDrawer,
    QPageContainer,
    QPage,
    QToolbar,
    QToolbarTitle,
    QBtn,
    QIcon,
    QList,
    QItem,
    QItemSection,
    QItemLabel,
    QTabs,
    QTab,
    QRouteTab,
  },
  directives: {},
  plugins: {},
});
