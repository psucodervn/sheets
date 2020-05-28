<template>
  <div>
    <p class="q-pa-sm">Report time range:</p>
    <q-input
      :value="range.label"
      dense
      input-class="text-center"
      outlined
      readonly
    >
      <template v-slot:append>
        <q-icon class="cursor-pointer" name="event" ref="dateIcon">
          <q-popup-proxy
            ref="qDateProxy"
            transition-hide="scale"
            transition-show="scale"
          >
            <q-date
              v-model="startDate"
              @input="onChangeTimeFrom"
              subtitle="Start date of week"
              first-day-of-week="1"
            />
          </q-popup-proxy>
        </q-icon>
      </template>
    </q-input>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { TimeRange } from '@/types/datetime';
import { date } from 'quasar';

@Component({})
export default class ReportTimeFilter extends Vue {
  @Prop({ type: TimeRange, required: true }) range!: TimeRange;
  startDate!: string;

  created() {
    this.startDate = date.formatDate(this.range.from, 'YYYY/MM/DD');
  }

  onChangeTimeFrom(val: Date) {
    this.range.from = new Date(val);
    this.range.to = date.addToDate(this.range.from, { days: 13 });
    (this.$refs.qDateProxy as any).hide();
    this.$emit('input', this.range);
  }
}
</script>

<style lang="scss" scoped></style>
