import Vue from "vue";

import "@/styles/quasar.scss";
import "@quasar/extras/material-icons/material-icons.css";
import {
  QBadge,
  QBtn,
  QDate,
  QDrawer,
  QFooter,
  QHeader,
  QIcon,
  QInput,
  QItem,
  QItemLabel,
  QItemSection,
  QLayout,
  QList,
  QPage,
  QPageContainer,
  QPopupProxy,
  QPullToRefresh,
  QRouteTab,
  QSelect,
  QSpace,
  QTab,
  QTable,
  QTabs,
  QTd,
  QTh,
  QToolbar,
  QToolbarTitle,
  QTr,
  Quasar
} from "quasar";

Vue.use(Quasar, {
  config: {
    dark: true
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
    QTable,
    QTd,
    QTr,
    QTh,
    QSelect,
    QSpace,
    QDate,
    QInput,
    QPopupProxy,
    QPullToRefresh,
    QBadge
  },
  directives: {},
  plugins: {}
});
