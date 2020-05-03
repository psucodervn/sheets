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
        :split-type.sync="tx.splitType"
        @update="setChanges"
      />
      <div class="row">
        <span class="q-py-sm label">Change Preview</span>
      </div>
      <div class="row q-pb-xs">
        <template v-for="(c, name) in tx.changes">
          <q-badge
            outline
            :color="c.value > 0 ? 'green' : c.value < 0 ? 'red' : 'grey'"
            class="q-my-xs q-mr-sm q-pa-sm"
            :label="
              `${c.name}: ${c.value > 0 ? '+' : ''}${formatValue(c.value)}`
            "
          />
        </template>
      </div>
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
      <q-card-actions align="right" class="q-pa-none q-pt-sm">
        <q-btn
          label="Save"
          icon="save"
          color="primary"
          @click="save"
          :loading="saving"
        ></q-btn>
      </q-card-actions>
    </q-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, PropSync, Vue, Watch } from 'vue-property-decorator';
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
  @Prop({ type: Boolean, default: false }) isEdit!: boolean;
  saving = false;

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

  validate() {
    if (!this.tx.summary) return 'Summary is required';
    if (!this.tx.payers.length) return 'Payers were required';
    if (!this.tx.participants.length) return 'Participants required';
  }

  async save() {
    const msg = this.validate();
    if (msg) {
      this.$q.notify({
        caption: 'Validate failed',
        message: msg,
        type: 'negative',
      });
      return;
    }
    try {
      this.saving = true;
      if (this.isEdit) {
        await BalanceModule.submitEditTransaction({ tx: this.tx });
      } else {
        await BalanceModule.submitNewTransaction({ tx: this.tx });
      }
      this.$q.notify({
        caption: 'Save succeed',
        message: `Transaction '${this.tx.summary}' was saved`,
        type: 'positive',
      });
    } catch (e) {
      this.$q.notify({
        caption: 'Save failed',
        message: e.message,
        type: 'negative',
      });
    } finally {
      this.saving = false;
    }
  }
}
</script>

<style scoped lang="scss">
::v-deep .label {
  color: grey;
}
</style>
