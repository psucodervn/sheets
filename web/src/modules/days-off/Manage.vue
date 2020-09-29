<template>
  <q-pull-to-refresh @refresh="fetchData" style="padding-bottom: 70px">
    <day-off-table :items="items" @remove="remove" />
    <add-button @click="add" />
  </q-pull-to-refresh>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { DayOff } from '@/types/logic';
import { DayOffService } from '@/modules/days-off/service';
import { showFailure, showSuccess } from '@/utils/dialog';
import DayOffTable from '@/modules/days-off/Table.vue';
import { Routes } from '@/router/names';

@Component({
  components: {
    DayOffTable,
  },
})
export default class DaysOff extends Vue {
  service = new DayOffService();
  items: DayOff[] = [];
  loading = false;

  created() {
    this.$navigation.title = 'Days Off';
    this.$navigation.from = null;
    this.$navigation.parent = null;
  }

  async mounted() {
    await this.fetchData();
  }

  async fetchData() {
    try {
      this.loading = true;
      this.items = await this.service.list();
    } catch (e) {
      showFailure('List failed: ' + e.message);
    } finally {
      this.loading = false;
    }
  }

  async add() {
    await this.$router.push({ name: Routes.DaysOffNew });
  }

  async remove(d: DayOff) {
    try {
      this.loading = true;
      await this.service.remove(d.id);
      await this.fetchData();
      showSuccess('Remove succeed');
    } catch (e) {
      showFailure('Remove failed', e.message);
    } finally {
      this.loading = false;
    }
  }
}
</script>

<style scoped lang="scss"></style>
