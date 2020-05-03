import Vue from 'vue';

import '@/styles/quasar.scss';
import '@quasar/extras/material-icons/material-icons.css';
import {
  ClosePopup,
  Dialog,
  Notify,
  QAvatar,
  QBadge,
  QBtn,
  QCard,
  QCardActions,
  QCardSection,
  QCheckbox,
  QDate,
  QDialog,
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
  QOptionGroup,
  QPage,
  QPageContainer,
  QPageSticky,
  QPopupProxy,
  QPullToRefresh,
  QRouteTab,
  QSelect,
  QSeparator,
  QSpace,
  QTab,
  QTable,
  QTabs,
  QTd,
  QTh,
  QTime,
  QToolbar,
  QToolbarTitle,
  QTooltip,
  QTr,
  Quasar,
} from 'quasar';

Vue.use(Quasar, {
  config: {
    dark: true,
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
    QBadge,
    QCard,
    QCardSection,
    QCardActions,
    QAvatar,
    QSeparator,
    QTooltip,
    QPageSticky,
    QTime,
    QDialog,
    QCheckbox,
    QOptionGroup,
  },
  directives: {
    ClosePopup,
  },
  plugins: {
    Notify,
    Dialog,
  },
});
