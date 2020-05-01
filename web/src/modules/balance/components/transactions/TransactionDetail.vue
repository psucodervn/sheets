<template>
  <div>
    <q-card class="q-pa-sm">
      <div class="row">
        <div class="col-xs-12 col-sm-grow q-pr-sm-xs q-pt-sm">
          <span class=" label">Summary</span>
          <q-input
            outlined
            dense
            v-model="tx.summary"
            class="q-pt-sm"
            placeholder="Transaction's summary"
          />
        </div>
        <div class="col-xs-12 col-sm-auto q-pl-sm-xs q-pt-sm">
          <span class=" label">Time</span>
          <time-input :time.sync="tx.time" class="q-pt-sm" />
        </div>
      </div>
      <list-payers :payers="tx.payers" :value.sync="tx.value" />
      <list-participants
        :participants="tx.participants"
        :total-value="tx.value"
      />
      <div class="row q-pt-sm">
        <span class="q-py-sm label">Note</span>
        <q-input
          autogrow
          outlined
          dense
          class="full-width"
          placeholder="Transaction's note"
          v-model="tx.description"
        />
      </div>
      <div class="row q-pt-sm">
        <span class="q-py-sm label">Change Preview</span>
      </div>
      <div class="row q-pb-xs">
        <template v-for="(c, name) in tx.changes">
          <q-badge
            v-if="c.value >= 0"
            outline
            color="green"
            class="q-ma-xs q-pa-xs"
            :class="{ blur: c.filteredOut }"
            :label="`${c.name}: +${formatValue(c.value)}`"
          />
          <q-badge
            v-else
            outline
            color="red"
            class="q-ma-xs q-pa-xs"
            :class="{ blur: c.filteredOut }"
            :label="`${c.name}: ${formatValue(c.value)}`"
          />
        </template>
      </div>
      <q-separator />
      <q-card-actions align="right" class="q-pa-none q-pt-sm">
        <q-btn v-if="!isEdit" label="Add" icon="add" color="primary"></q-btn>
        <q-btn v-else label="Save" icon="save" color="primary"></q-btn>
      </q-card-actions>
    </q-card>
  </div>
</template>

<script lang="ts">
import { Component, PropSync, Vue, Watch } from 'vue-property-decorator';
import { ITransaction } from '@/modules/balance/types/transaction';
import TimeInput from '@/modules/balance/components/transactions/TimeInput.vue';
import { BalanceModule } from '@/store';
import UserPicker from '@/modules/balance/components/transactions/UserPicker.vue';
import ListPayers from '@/modules/balance/components/transactions/ListPayers.vue';
import ListParticipants from '@/modules/balance/components/transactions/ListParticipants.vue';
import formatter from '@/utils/formatter';

@Component({
  components: { ListParticipants, ListPayers, UserPicker, TimeInput },
})
export default class TransactionDetail extends Vue {
  @PropSync('transaction', { type: Object, required: true }) tx!: ITransaction;
  isEdit = false;

  get users() {
    return BalanceModule.users.sort((a, b) => a.name.localeCompare(b.name));
  }

  @Watch('tx.payers', { immediate: true, deep: true })
  setTotalValue() {
    this.tx.value = this.tx.payers.reduce((p, u) => p + u.value, 0);
  }

  @Watch('tx.payers', { deep: true, immediate: true })
  @Watch('tx.participants', { deep: true, immediate: true })
  async setChanges() {
    this.tx.changes = await BalanceModule.calcChanges(this.tx);
  }

  formatValue(val: number) {
    return formatter.currency(val);
  }
}
</script>

<style scoped lang="scss">
::v-deep .label {
  color: grey;
}
</style>
