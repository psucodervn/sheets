<template>
  <q-input outlined dense v-model="datetime" readonly>
    <template v-slot:prepend>
      <q-icon name="event" class="cursor-pointer">
        <q-popup-proxy transition-show="scale" transition-hide="scale">
          <q-date v-model="datetime" mask="YYYY-MM-DD HH:mm" today-btn />
        </q-popup-proxy>
      </q-icon>
    </template>

    <template v-slot:append>
      <q-icon name="access_time" class="cursor-pointer">
        <q-popup-proxy transition-show="scale" transition-hide="scale">
          <q-time v-model="datetime" mask="YYYY-MM-DD HH:mm" format24h />
        </q-popup-proxy>
      </q-icon>
    </template>
  </q-input>
</template>

<script lang="ts">
import { Component, PropSync, Vue } from 'vue-property-decorator';
import { date } from 'quasar';

@Component
export default class TimeInput extends Vue {
  @PropSync('time', { required: true }) t!: Date;

  get datetime() {
    return date.formatDate(this.t, 'YYYY-MM-DD HH:mm');
  }

  set datetime(val: string) {
    this.t = date.extractDate(val, 'YYYY-MM-DD HH:mm');
  }
}
</script>
