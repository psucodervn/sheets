<template>
  <div class="q-pt-sm">
    <div class="row">
      <div class="col q-py-sm label">
        Payers
        <template v-if="users.length > 0"
          >(total: {{ totalValueText }})</template
        >
      </div>
      <div class="col col-grow">
        <q-btn
          rounded
          flat
          icon="add_circle_outline"
          text-color="green"
          @click="payerDialog = !payerDialog"
          v-if="users.length && payerCandidates.length"
        />
      </div>
    </div>
    <div class="row">
      <div v-if="!users.length" class="col col-xs-12 q-pa-xs">
        <q-btn
          outline
          label="No payers. Click to add"
          icon-right="add"
          class="full-width"
          color="green"
          @click="payerDialog = !payerDialog"
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
        <q-input
          bg-color="green-10"
          standout=""
          dense
          :prefix="`${getName(u.id)}: `"
          suffix="vnđ"
          :value="u.value"
          @input="val => updateValue(idx, val)"
          type="number"
          pattern="[0-9]*"
          input-style="text-align: right"
          @focus="evt => evt.target.select && evt.target.select()"
        >
          <template v-slot:append>
            <q-btn
              dense
              rounded
              flat
              icon="clear"
              color="red"
              @click="() => remove(u.id)"
            />
          </template>
        </q-input>
      </div>
    </div>
    <q-dialog v-model="payerDialog" persistent>
      <user-picker
        :candidates="payerCandidates"
        @input="onAddPayers"
        title="Choose payers to add:"
      />
    </q-dialog>
  </div>
</template>

<script lang="ts">
import { Component, Prop, PropSync, Vue } from 'vue-property-decorator';
import { ITransactionUser } from '@/modules/balance/types/transaction';
import { BalanceModule } from '@/store';
import UserPicker from '@/modules/balance/components/transactions/UserPicker.vue';
import formatter from '@/utils/formatter';
import { IUser } from '@/model/user';

@Component({
  components: { UserPicker },
})
export default class ListPayers extends Vue {
  @PropSync('payers', { type: Array, required: true })
  users!: ITransactionUser[];
  @PropSync('value', { type: Number, required: true }) totalValue!: number;
  payerDialog = false;

  getName(id: string) {
    return BalanceModule.userIds[id]!.name;
  }

  get payerCandidates(): IUser[] {
    return BalanceModule.users.filter(u =>
      this.users.every(tu => tu.id !== u.id)
    );
  }

  onAddPayers(ids: string[]) {
    const adds = ids.map(
      (id: string): ITransactionUser => {
        const u = this.payerCandidates.find(u => u.id === id)!;
        return { id: u.id, value: 0 };
      }
    )!;
    this.users.push(...adds);
  }

  remove(id: string) {
    const idx = this.users.findIndex(u => u.id === id);
    this.users.splice(idx, 1);
  }

  updateValue(idx: number, val: number | string) {
    this.users[idx].value = Number(val) || 0;
  }

  get totalValueText(): string {
    return formatter.currency(this.totalValue) + ' vnđ';
  }
}
</script>

<style scoped lang="scss"></style>
