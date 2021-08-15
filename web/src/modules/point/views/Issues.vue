<template>
  <div>
    <point-time-filter />
    <q-space class="q-my-sm" />
    <issue-table :issues="issues" />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import IssueTable from '@/modules/point/components/IssueTable.vue';
import { PointModule } from '@/store';
import { Routes } from '@/router/names';
import PointTimeFilter from '@/modules/point/components/PointTimeFilter.vue';
import formatter from '@/utils/formatter';

@Component({
  components: {
    PointTimeFilter,
    IssueTable,
  },
})
export default class Issues extends Vue {
  displayName = 'V';

  get name() {
    return this.$route.params.name;
  }

  get time() {
    return formatter.month(PointModule.month);
  }

  get issues() {
    const user = PointModule.users.find(u => u.name === this.name);
    if (!user) return [];
    this.displayName = user.displayName;
    return user.issues;
  }

  @Watch('displayName', { immediate: true })
  setHeader() {
    this.$navigation.title = `${this.displayName}'s Issues`;
  }

  @Watch('time', { immediate: true })
  fetchData() {
    PointModule.fetchPoints({ month: PointModule.month });
  }

  mounted() {
    this.$navigation.parent = { name: Routes.Point };
  }
}
</script>
