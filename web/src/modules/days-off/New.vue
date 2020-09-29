<template>
  <div>
    <day-off-form :item.sync="item" @save="create" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Routes } from '@/router/names';
import DayOffForm from '@/modules/days-off/Form.vue';
import { DayOff } from '@/types/logic';
import { DayOffService } from '@/modules/days-off/service';
import { showFailure } from '@/utils/dialog';

@Component({
  components: {
    DayOffForm,
  },
})
export default class NewDayOff extends Vue {
  service = new DayOffService();
  item: DayOff | null = null;

  async created() {
    this.$navigation.title = 'New Day Off';
    this.$navigation.parent = { name: Routes.DaysOff };
    this.item = {
      id: '',
      part: '',
      date: new Date(),
      note: '',
      userId: '',
    };
  }

  async create(item: DayOff | null) {
    if (!item) return;
    try {
      await this.service.save(item);
    } catch (e) {
      showFailure('Save failed', e.message);
    }
  }
}
</script>

<style scoped lang="scss"></style>
