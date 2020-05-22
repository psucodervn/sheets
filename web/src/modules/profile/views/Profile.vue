<template>
  <div>Profile</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component({})
export default class Profile extends Vue {
  created() {
    this.$navigation.title = 'Profile';
    this.$navigation.parent = null;
    this.$navigation.from = null;
  }
}
</script>

<style scoped lang="scss"></style>
