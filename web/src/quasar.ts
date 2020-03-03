import Vue from 'vue';

import './styles/quasar.scss';
import '@quasar/extras/material-icons/material-icons.css';
import {
  Quasar,
  QLayout,
  QHeader,
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
  QItemLabel, QTabs, QTab, QRouteTab, QFooter,
} from 'quasar';

Vue.use(Quasar, {
  config: {
    dark: 'auto',
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
