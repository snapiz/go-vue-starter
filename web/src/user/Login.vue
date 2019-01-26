<template>
  <Page>
    <div id="login">
      <h3>Login with your account</h3>
      <sui-form @submit.prevent="login">
        <sui-form-field>
          <label>Email</label>
          <sui-input
            v-model.trim="$v.email.$model"
            icon="user"
            @focus="error = ''"
            icon-position="left"
          />
          <div class="error" v-if="$v.email.$dirty && !$v.email.required">
            Field is required
          </div>
          <div class="error" v-if="$v.email.$dirty && !$v.email.email">
            Must be en email
          </div>
        </sui-form-field>

        <sui-form-field>
          <label>Password</label>
          <sui-input
            type="password"
            v-model.trim="$v.password.$model"
            icon="lock"
            icon-position="left"
            @focus="error = ''"
          />
          <div class="error" v-if="$v.password.$dirty && !$v.password.required">
            Field is required
          </div>
          <div
            class="error"
            v-if="
              $v.password.$dirty &&
                (!$v.password.minLength || !$v.password.maxLength)
            "
          >
            Must be between 8 and 20 characters length
          </div>
        </sui-form-field>
        <div class="error" v-if="error">{{ error }}</div>
        <router-link to="/register">I don't have an account</router-link>
        <sui-button type="submit">Login</sui-button>
      </sui-form>
      <sui-divider horizontal>Or</sui-divider>
      <div class="socials">
        <sui-button icon="google" @click="loginOauth2('google')"
          >Google</sui-button
        >
        <sui-button icon="facebook" @click="loginOauth2('facebook')"
          >Facebook</sui-button
        >
      </div>
    </div>
  </Page>
</template>

<style scoped>
#login {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
}
.ui.form {
  background-color: white;
  padding: 15px;
}
.ui.button {
  width: 100%;
  margin-top: 15px;
}
.socials,
.ui.form,
.ui.horizontal.divider {
  min-width: 300px;
}
.error {
  color: red;
}
</style>

<script>
import {
  required,
  minLength,
  maxLength,
  email
} from "vuelidate/lib/validators";

import axios from "axios";
import Page from "@/common/Page.vue";
import RouterLink from "@/common/RouterLink.vue";
import router from "@/router";
import apollo from "@/apollo";
import { getUrlParameter } from "@/utils";

const winProviderOptions = {
  google: "width=452,height=633",
  facebook: "width=580,height=400"
};

export default {
  name: "login",
  components: {
    Page,
    RouterLink
  },
  data: () => ({
    email: "",
    password: "",
    error: ""
  }),
  validations: {
    email: {
      required,
      email
    },
    password: {
      required,
      minLength: minLength(8),
      maxLength: maxLength(20)
    }
  },
  async beforeRouteEnter() {
    try {
      await axios.post("/auth/logout");
      await apollo.resetStore();
    } catch (error) {
      console.log("%cError logout", "color: orange;", error.message);
    }
  },
  methods: {
    async login() {
      this.$v.$touch();

      if (this.$v.$invalid) {
        return;
      }

      try {
        const { email, password } = this;

        await axios.post("/auth/local", { email, password });
        await apollo.resetStore();

        router.context.history.push(getUrlParameter("redirect") || "/");
      } catch (error) {
        this.error = error.response.data;
      }
    },
    async loginOauth2(provider) {
      try {
        await apollo.resetStore();

        window.open(`/auth/${provider}`, "", winProviderOptions[provider]);
      } catch (error) {
        console.error(error);
      }
    }
  }
};
</script>
