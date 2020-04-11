<template>
  <q-page class="q-px-sm">
  </q-page>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator';
  import IssueTable from '@/modules/point/components/IssueTable.vue';

  @Component({
    components: {
      IssueTable,
    },
  })
  export default class Issues extends Vue {
    get name() {
      return this.$route.params.name;
    }

    mounted() {
      this.$navigation.title = `${this.name}'s Issues`;
    }
  }
</script>

<style lang="scss" scoped>

</style>
