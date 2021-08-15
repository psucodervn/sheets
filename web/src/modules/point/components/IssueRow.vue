<template>
  <div class="q-py-xs ellipsis">
    <span style="display: inline-block; min-width: 100px">
      <a
        :href="`https://pm.vzota.com.vn/browse/${issue.key}`"
        class="link"
        target="_blank"
        rel="noreferrer"
      >
        <q-badge
          :color="color"
          :label="issue.key"
          class="q-pa-xs text-bold"
          outline
        />
      </a>
      <q-badge
        :label="`+${point}`"
        class="q-pa-xs q-ml-xs text-bold"
        color="green"
        outline
        v-if="point >= 0"
      />
    </span>
    <span>
      {{ issue.summary }}
      <q-tooltip
        anchor="top middle"
        content-class="bg-blue-grey-9 text-white shadow-2"
        content-style="font-size: 1em; border: black solid 1px"
        self="bottom middle"
      >
        {{ `${issue.key}: ${issue.summary} (${issue.status})` }}
      </q-tooltip>
    </span>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { IIssue } from '@/model/point';

@Component({
  components: {},
})
export default class IssueRow extends Vue {
  @Prop({ type: Object, required: true }) issue!: IIssue;

  get point(): number {
    if (this.issue.status.toUpperCase() === 'DONE') {
      return +this.issue.point;
    }
    return -1;
  }

  get color(): string {
    switch (this.issue.status.toUpperCase()) {
      case 'DONE':
        return 'green';
      default:
        return 'blue';
    }
  }
}
</script>

<style></style>
