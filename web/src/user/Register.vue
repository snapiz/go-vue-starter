<template>
  <Page>
    <div id="register">
      <h3>Create your new account</h3>
      <sui-form @submit.prevent="register">
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
        <router-link to="/login">I already have an account</router-link>
        <sui-button type="submit">Register</sui-button>
      </sui-form>
    </div>
  </Page>
</template>

<style scoped>
#register {
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

export default {
  name: "register",
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
    async register() {
      this.$v.$touch();

      if (this.$v.$invalid) {
        return;
      }

      try {
        const { email, password } = this;

        await axios.post("/auth/local/register", { email, password });
        await apollo.resetStore();

        router.context.history.push("/");
      } catch (error) {
        this.error = error.response.data;
      }
    }
  }
};
</script>
