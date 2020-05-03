<template>
  <div class="q-pt-sm">
    <div class="row">
      <div class="col q-py-sm label">Participants</div>
      <div class="col col-grow">
        <q-option-group
          :options="splitOptions"
          v-model="split"
          inline
          dense
          class="float-left option"
          v-if="$q.screen.gt.xs && users.length"
        ></q-option-group>
        <q-btn
          rounded
          flat
          icon="add_circle_outline"
          text-color="green"
          @click="participantDialog = !participantDialog"
          v-if="users.length && participantCandidates.length"
        />
      </div>
    </div>
    <div class="row">
      <q-option-group
        :options="splitOptions"
        v-model="split"
        inline
        dense
        class="q-pa-sm option"
        v-if="!$q.screen.gt.xs && users.length"
      ></q-option-group>
      <div v-if="!users.length" class="col col-xs-12 q-pa-xs">
        <q-btn
          outline
          label="No participants. Click to add"
          icon-right="add"
          class="full-width"
          color="green"
          @click="participantDialog = !participantDialog"
        ></q-btn>
      </div>
      <div
        class="col-xs-12 q-py-xs"
        :class="{
          'col-sm-6': users.length > 1,
          'q-pr-sm-xs': idx % 2 === 0,
          'q-pl-sm-xs': idx % 2 === 1,
        }"
        v-for="(u, idx) in users"
      >
        <participant-input
          :split-option="split"
          :participant="u"
          @remove="() => remove(u.id)"
          @update="setValues"
        />
      </div>
    </div>
    <q-dialog v-model="participantDialog" persistent>
      <user-picker
        :candidates="participantCandidates"
        @input="onAddParticipants"
        title="Choose participants to add:"
      />
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { Component, Prop, PropSync, Vue, Watch } from 'vue-property-decorator';
import {
  ESplitOption,
  ITransactionUser,
} from '@/modules/balance/types/transaction';
import UserPicker from '@/modules/balance/components/transactions/UserPicker.vue';
import { BalanceModule } from '@/store';
import ParticipantInput from '@/modules/balance/components/transactions/ParticipantInput.vue';

@Component({
  components: { ParticipantInput, UserPicker },
})
export default class ListParticipants extends Vue {
  @PropSync('participants', { type: Array, required: true })
  users!: ITransactionUser[];
  @Prop({ type: Number, required: true }) totalValue!: number;
  @PropSync('splitType', { required: true }) split!: ESplitOption;

  participantDialog = false;

  splitOptions = [
    { label: 'Equal', value: ESplitOption.Equal },
    { label: 'Ratio', value: ESplitOption.Ratio },
    { label: 'Custom', value: ESplitOption.Custom, disable: true },
  ];

  get participantCandidates() {
    return BalanceModule.users
      .filter(u => this.users.every(tu => tu.id !== u.id))
      .sort((a, b) => a.name.localeCompare(b.name));
  }

  onAddParticipants(ids: string[]) {
    const adds = ids.map(
      (id: string): ITransactionUser => {
        const u = this.participantCandidates.find(u => u.id === id)!;
        return { id: u.id, value: 0, name: u.name, percent: 1 };
      }
    )!;
    this.users.push(...adds);
  }

  getName(id: string) {
    return BalanceModule.userIds[id]!.name;
  }

  remove(id: string) {
    const idx = this.users.findIndex(u => u.id === id);
    this.users.splice(idx, 1);
  }

  updateValue(idx: number, val: number | string) {
    this.users[idx].value = Number(val) || 0;
  }

  @Watch('users.length')
  @Watch('totalValue')
  @Watch('split', { immediate: true })
  setValues() {
    if (!this.users.length) return;
    if (this.split === ESplitOption.Equal) {
      const v = (this.totalValue / this.users.length).toFixed(0);
      for (let i = 0; i < this.users.length; i++) {
        this.users[i].value = Number(v);
      }
    } else if (this.split === ESplitOption.Ratio) {
      const sumPercent = this.users.reduce((p, u) => p + (u.percent || 1), 0);
      for (let i = 0; i < this.users.length; i++) {
        this.users[i].value =
          ((this.users[i].percent || 1) / sumPercent) * this.totalValue;
      }
    }
    this.$emit('update');
  }
}
</script>

<style scoped lang="scss">
.option {
  padding-top: 7px;
  padding-right: 15px;
  font-size: 0.9em;
}
</style>
