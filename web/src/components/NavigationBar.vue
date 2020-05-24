<template>
  <div class="q-pb-xs bar row justify-between items-stretch">
    <q-btn @click="goBack" class="btn" rounded v-if="canBack">
      <q-icon name="arrow_back_ios" />
    </q-btn>
    <div class="btn" v-else></div>
    <div class="title ellipsis vertical-middle">
      {{ title }}
    </div>
    <div class="btn"></div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component({})
export default class NavigationBar extends Vue {
  get title(): string {
    return this.$navigation.title;
  }

  get canBack(): boolean {
    return !!this.$navigation.from || !this.$route.meta.root;
  }

  goBack() {
    if (!this.canBack) return;
    if (this.$navigation.from) return this.$router.push(this.$navigation.from);
    return this.$router.push(this.$navigation.parent!);
  }
}
</script>

<style lang="scss" scoped>
.bar {
  /*border: #21BA45 dashed 1px;*/

  .btn {
    width: 38px;

    i {
      width: 15px;
    }
  }
}

.title {
  margin: 0 auto;
  line-height: 36px;
}
</style>
