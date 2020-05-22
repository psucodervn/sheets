<template>
  <component :is="layout" id="app"> </component>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import SimpleLayout from '@/layouts/SimpleLayout.vue';
import DefaultLayout from '@/layouts/DefaultLayout.vue';

@Component({
  components: {
    DefaultLayout,
    SimpleLayout,
  },
})
export default class App extends Vue {
  get layout() {
    const layout = this.$route.meta.layout;
    return layout || 'default-layout';
  }
}
</script>

<style lang="scss" scoped>
#app {
  font-family: 'JetBrains Mono', monospace;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: 0 auto;
  min-width: 360px;
  max-width: 640px;
}
</style>

<!--<template v-for="(_, slot) of $scopedSlots" v-slot:[slot]="scope">-->
<!--  <slot :name="slot" v-bind="scope" />-->
<!--</template>-->
