<template>
  <q-page class="q-pa-md flex flex-center row">
    <div class="col-xs-12 col-sm-8 col-md-4 q-pa-sm text-center">
      <p class="text-center text-h6">Login</p>
      <q-form :disabled="!canLoginByPassword">
        <q-input
          dense
          v-model="form.email"
          outlined
          label="Email"
          type="email"
          class="q-my-sm"
          autocomplete="current-email"
          :disable="!canLoginByPassword"
        />
        <q-input
          dense
          v-model="form.password"
          outlined
          label="Password"
          type="password"
          class="q-my-sm"
          autocomplete="current-password"
          :disable="!canLoginByPassword"
        />
        <q-btn
          color="primary"
          class="full-width"
          @click="login"
          :disable="!canLoginByPassword"
        >
          Submit
        </q-btn>
      </q-form>
      <p class="text-grey q-my-sm">- OR -</p>
      <q-btn color="red-5 full-width" @click="authenticate('google')">
        Login by @vzota.com.vn email
      </q-btn>
    </div>
  </q-page>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { ProfileModule } from '@/store';

@Component({})
export default class Login extends Vue {
  form = {
    email: '',
    password: '',
  };
  canLoginByPassword = false;

  async login() {
    try {
      await ProfileModule.login(this.form);
      const redirect = String(this.$route.query.redirect || '/');
      await this.$router.push({ path: redirect });
    } catch (e) {
      this.$q.notify({
        message: 'Login failed: ' + String(e.message),
        type: 'negative',
        position: 'top',
      });
    }
  }

  async authenticate(provider: string) {
    try {
      await ProfileModule.authenticate({ provider });
      const redirect = String(this.$route.query.redirect) || '/';
      await this.$router.push({ path: redirect });
    } catch (e) {
      this.$q.notify({
        message: 'Login failed: ' + String(e.message),
        type: 'negative',
        position: 'top',
      });
    }
  }
}
</script>

<style scoped lang="scss"></style>
