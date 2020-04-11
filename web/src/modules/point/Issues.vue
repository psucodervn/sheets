<template>
  <issue-table :issues="issues"/>
</template>

<script lang="ts">
  import { Component, Vue, Watch } from 'vue-property-decorator';
  import IssueTable from '@/modules/point/components/IssueTable.vue';
  import { PointModule } from '@/store';

  @Component({
    components: {
      IssueTable,
    },
  })
  export default class Issues extends Vue {
    displayName = 'V';

    get name() {
      return this.$route.params.name;
    }

    get issues() {
      const u = PointModule.users.find(u => u.name === this.name);
      if (!u) return [];
      this.displayName = u.displayName;
      return u.issues;
    }

    @Watch('displayName', { immediate: true })
    setHeader() {
      this.$navigation.title = `${this.displayName}'s Issues`;
    }
  }
</script>

<style lang="scss" scoped>

</style>
