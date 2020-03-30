<template>
  <q-layout view="hHh lpR fFf">
    <q-page-container>
      <router-view/>
    </q-page-container>
    <q-footer>
      <q-tabs dense>
        <q-route-tab exact inline-label label="Balance" to="/"/>
        <q-route-tab
          exact
          inline-label
          label="Points"
          to="/points"
        />
      </q-tabs>
    </q-footer>
  </q-layout>
</template>

<script>
  import { openURL } from 'quasar';

  export default {
    name: 'Layout',
    data() {
      return {
        leftDrawerOpen: this.$q.platform.is.desktop,
        navs: [
          {
            label: 'Todo',
            icon: 'list',
            to: '/',
          },
          {
            label: 'Settings',
            icon: 'settings',
            to: '/settings',
          },
        ],
      };
    },
    methods: {
      openURL,
    },
  };
</script>

<style lang="scss">
  .q-drawer {
    .q-router-link--exact-active {
      color: white !important;
    }
  }
</style>
