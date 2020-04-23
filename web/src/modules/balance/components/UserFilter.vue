<template>
  <q-select
    :options="allItems"
    class="col"
    dense
    label="Users"
    outlined
    v-model="items"
    multiple
    clearable
    emit-value
  ></q-select>
</template>

<script lang="ts">
import { Component, Prop, PropSync, Vue } from 'vue-property-decorator';
import { IUser } from '@/model/user';

interface SelectItem {
  value: string;
  label: string;
}

@Component({})
export default class UserFilter extends Vue {
  @Prop({ type: Array, required: true }) users!: IUser[];
  @PropSync('selectedUsers', { required: true })
  items!: IUser[];

  get allItems() {
    return this.users
      .map((u: IUser): SelectItem => ({ label: u.name, value: u.name }))
      .sort((a, b) => a.label.localeCompare(b.label));
  }
}
</script>

<style scoped lang="scss"></style>
