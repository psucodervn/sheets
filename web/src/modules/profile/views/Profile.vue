<template>
  <q-pull-to-refresh @refresh="fetchData">
    <div class="q-pa-sm">
      <p>Name: {{ currentUser.name }}</p>
      <p>Email: {{ currentUser.email || '[not set]' }}</p>
      <p>Sheet Name: {{ currentUser.sheetName || '[not set]' }}</p>
      <p>Jira Name: {{ currentUser.jiraName || '[not set]' }}</p>
      <p>Telegram ID: {{ currentUser.telegramId || '[not set]' }}</p>
      <div v-if="!currentUser.telegramId">
        <q-separator spaced />
        <q-btn
          color="blue-9"
          label="Integrate With Telegram"
          v-if="!telegramLink"
          @click="generateTelegramLink"
          class="full-width"
        ></q-btn>
        <p v-else>
          Connect to Telegram:<br />
          <a :href="telegramLink" target="_blank">{{ telegramLink }}</a>
        </p>
      </div>
      <q-separator spaced />
      <q-btn
        @click="logout"
        label="Logout"
        class="full-width"
        color="red-5"
      ></q-btn>
    </div>
  </q-pull-to-refresh>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { ProfileModule } from '@/store';

@Component({})
export default class Profile extends Vue {
  telegramLink: string | null = null;

  async fetchData(done?: Function) {
    try {
      await ProfileModule.fetchMe();
    } catch (e) {
      this.$q.notify({
        message: `Fetch data failed: ${e.message}`,
        type: 'negative',
      });
    } finally {
      if (done) {
        done();
      }
    }
  }

  async created() {
    this.$navigation.title = 'Profile';
    await this.fetchData();
  }

  get currentUser() {
    return ProfileModule.currentUser || {};
  }

  async logout() {
    await ProfileModule.logout();
    await this.$router.push('/');
  }

  async generateTelegramLink() {
    try {
      this.telegramLink = await ProfileModule.generateTelegramLink();
    } catch (e) {
      this.$q.notify({
        message: e.message,
        type: 'negative',
      });
    }
  }
}
</script>

<style scoped lang="scss"></style>
