<template>
  <div>
    Profile
    <q-btn @click="logout" label="Logout"></q-btn>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { ProfileModule } from '@/store';

@Component({})
export default class Profile extends Vue {
  created() {
    this.$navigation.title = 'Profile';
    this.$navigation.parent = null;
    this.$navigation.from = null;
  }

  async logout() {
    await ProfileModule.logout();
    await this.$router.push('/');
  }
}
</script>

<style scoped lang="scss"></style>
