<template>
  <q-form @submit.prevent.stop="save">
    <q-card class="q-my-sm">
      <q-card-section class="row q-px-sm q-py-sm q-col-gutter-sm">
        <q-select
          label="Name"
          :options="users"
          v-model="_item.userId"
          dense
          options-dense
          map-options
          emit-value
          outlined
          class="col-sm-6 col-xs-12"
          :rules="[val => !!val || 'Name is required']"
        ></q-select>
        <date-picker
          :date.sync="_item.date"
          label="Date"
          class="col-sm-6 col-xs-12"
          hint="Format: DD/MM/YYYY"
          readonly
        />
        <q-select
          label="Time"
          :options="['All day', 'Morning', 'Afternoon']"
          v-model="_item.part"
          outlined
          dense
          options-dense
          class="col-sm-6 col-xs-12"
        ></q-select>
        <q-input
          v-model="_item.note"
          label="Note"
          dense
          outlined
          class="col-sm-6 col-xs-12"
        ></q-input>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn
          label="Save"
          icon="save"
          color="primary"
          type="submit"
          :loading="loading"
        ></q-btn>
      </q-card-actions> </q-card
  ></q-form>
</template>

<script lang="ts">
import { Component, Emit, Prop, PropSync, Vue } from 'vue-property-decorator';
import { DayOff } from '@/types/logic';
import { BalanceModule } from '@/store';

@Component({})
export default class DayOffForm extends Vue {
  @PropSync('item', { required: true }) _item!: DayOff;
  @Prop({ default: false }) loading!: boolean;

  mounted() {
    this._item.part = this._item.part || 'All day';
  }

  get users() {
    return BalanceModule.users
      .map(u => ({ label: u.name, value: u.id }))
      .sort((a, b) => a.label.localeCompare(b.label));
  }

  @Emit('save')
  save() {
    return this._item;
  }
}
</script>

<style scoped lang="scss"></style>
