<template>
  <div class="q-py-xs bar row justify-between items-stretch">
    <q-btn @click="goBack" class="btn" rounded>
      <q-icon name="arrow_back_ios" v-if="canBack"/>
    </q-btn>
    <div class="title ellipsis">
      {{ title }}
    </div>
    <div class="btn">
    </div>
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
      return this.$navigation.to !== null;
    }

    goBack() {
      if (this.canBack) {
        this.$router.push(this.$navigation.to!);
      }
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
    margin: auto 0;
  }
</style>
