<template>
  <q-card :style="{ 'min-width': `${minWidth}px` }">
    <q-card-section class="bg-grey-9 text-uppercase"
      >{{ title }}
    </q-card-section>
    <q-separator />
    <div class="row">
      <div v-for="u in users" class="col-xs-6 col-sm-4">
        <q-checkbox v-model="u.selected" :label="u.label" />
      </div>
    </div>
    <q-separator />
    <q-card-actions align="right">
      <q-btn
        v-close-popup
        flat
        color="primary"
        label="Cancel"
        @click="dismiss"
      />
      <q-btn v-close-popup color="primary" label="OK" @click="update" />
    </q-card-actions>
  </q-card>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue, Watch } from 'vue-property-decorator';
import { IUser } from '@/model/user';
import { dom } from 'quasar';

interface Item {
  value: string;
  label: string;
  selected: boolean;
}

@Component({})
export default class UserPicker extends Vue {
  @Prop({ type: Array, required: true }) candidates!: IUser[];
  @Prop({ type: String, default: 'users' }) title!: string;
  users: Item[] = [];

  @Watch('candidates', { immediate: true })
  setUsers() {
    this.users = this.candidates.map(u => ({
      value: u.id,
      label: u.name,
      selected: false,
    }));
  }

  @Emit('input')
  dismiss() {
    return [];
  }

  @Emit('input')
  update() {
    return this.users.filter(u => u.selected).map(u => u.value);
  }

  get minWidth() {
    const el = window.document.getElementById('app')!;
    return Math.min(480, dom.width(el) - 40);
  }
}
</script>

<style scoped lang="scss"></style>
