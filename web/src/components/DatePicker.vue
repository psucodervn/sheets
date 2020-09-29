<template>
  <q-input
    dense
    outlined
    :value="dateModel"
    mask="##/##/####"
    :label="label"
    @input="onInput"
    v-bind="$attrs"
  >
    <template v-slot:append>
      <q-icon name="event" class="cursor-pointer">
        <q-popup-proxy
          ref="qDateProxy"
          transition-show="scale"
          transition-hide="scale"
        >
          <q-date
            :value="dateModel"
            mask="DD/MM/YYYY"
            @input="onInputFromQDate"
          />
        </q-popup-proxy>
      </q-icon>
    </template>
  </q-input>
</template>

<script lang="ts">
import { Component, Prop, PropSync, Vue } from 'vue-property-decorator';
import { date as dateUtils } from 'quasar';

@Component({})
export default class DatePicker extends Vue {
  @Prop({ default: undefined }) label!: string;
  @PropSync('date', { required: true }) _date!: Date;

  mounted() {
    this._date = new Date(
      this._date.getFullYear(),
      this._date.getMonth(),
      this._date.getDate()
    );
  }

  get dateModel() {
    return dateUtils.formatDate(this._date, 'DD/MM/YYYY');
  }

  onInput(val: string) {
    // d = dateUtils.addToDate(d, { minutes: -d.getTimezoneOffset() });
    this._date = dateUtils.extractDate(val, 'DD/MM/YYYY');
  }

  onInputFromQDate(
    value: string,
    reason: string,
    details: { day: number; month: number; year: number }
  ) {
    this._date = new Date(details.year, details.month - 1, details.day);
    // eslint-disable-next-line @typescript-eslint/ban-ts-ignore
    // @ts-ignore
    this.$refs.qDateProxy.hide();
  }
}
</script>

<style scoped lang="scss"></style>
