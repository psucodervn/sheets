<template>
  <div class="q-pa-sm">
    <p>Name: {{ currentUser.name }}</p>
    <p>Email: {{ currentUser.email || '[not set]' }}</p>
    <p>Sheet Name: {{ currentUser.sheetName || '[not set]' }}</p>
    <p>Jira Name: {{ currentUser.jiraName || '[not set]' }}</p>
    <q-separator spaced />
    <q-btn
      @click="logout"
      label="Logout"
      class="full-width"
      color="red-5"
    ></q-btn>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { ProfileModule } from '@/store';

@Component({})
export default class Profile extends Vue {
  async created() {
    this.$navigation.title = 'Profile';
    await ProfileModule.fetchMe();
  }

  get currentUser() {
    return ProfileModule.currentUser || {};
  }

  async logout() {
    await ProfileModule.logout();
    await this.$router.push('/');
  }
}
</script>

<style scoped lang="scss"></style>
