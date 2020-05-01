<template>
  <q-input
    v-if="splitOption === ESplitOption.Equal"
    readonly
    standout=""
    dense
    :prefix="`${user.name}: `"
    suffix="vnđ"
    :value="formatText(user.value)"
    type="text"
    :input-class="['text-right', 'q-pr-sm']"
    @focus="evt => evt.target.select && evt.target.select()"
  >
    <template v-slot:append>
      <q-btn dense rounded flat icon="clear" color="red" @click="remove" />
    </template>
  </q-input>
  <q-input
    v-else-if="splitOption === ESplitOption.Ratio"
    type="number"
    standout=""
    dense
    :value="user.percent"
    @input="val => updatePercent(user, val)"
    :prefix="`${user.name}: `"
    pattern="[0-9]*"
    input-class="text-right"
  >
    <template v-slot:append>
      <q-btn
        dense
        rounded
        flat
        icon="add"
        @click="user.percent++"
        color="grey"
      />
      <q-btn
        dense
        rounded
        flat
        icon="remove"
        @click="user.percent > 1 && user.percent--"
        color="grey"
      />
      <q-btn dense rounded flat icon="clear" color="red" @click="remove" />
    </template>
  </q-input>
  <q-input
    v-else
    type="number"
    standout=""
    dense
    v-model="user.value"
    :prefix="`${user.name}: `"
    pattern="[0-9]*"
    suffix="vnđ"
    input-class="text-right"
  >
    <template v-slot:append>
      <q-btn dense rounded flat icon="clear" color="red" @click="remove" />
    </template>
  </q-input>
</template>

<script lang="ts">
import { Component, Emit, Prop, PropSync, Vue } from 'vue-property-decorator';
import {
  ESplitOption,
  ITransactionUser,
} from '@/modules/balance/types/transaction';
import formatter from '@/utils/formatter';

@Component
export default class ParticipantInput extends Vue {
  @PropSync('participant', { type: Object, required: true })
  user!: ITransactionUser;
  @Prop({ type: Number, required: true }) splitOption!: ESplitOption;
  ESplitOption = ESplitOption;

  @Emit('remove')
  remove() {}

  formatText(val: number | string) {
    return formatter.currency(Number(val));
  }

  updatePercent(u: ITransactionUser, val: number) {
    u.percent = Math.max(1, val);
  }
}
</script>
