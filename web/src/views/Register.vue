<template>
  <Page>
    <div class="register">
      <form v-on:submit="register">
        <label>Email</label>
        <input type="text" v-model.trim="$v.email.$model" v-on:focus="error = '';">
        <div class="error" v-if="$v.email.$dirty && !$v.email.required">Field is required</div>
        <div class="error" v-if="$v.email.$dirty && !$v.email.email">Must be en email</div>

        <label>Password</label>
        <input type="password" v-model.trim="$v.password.$model" v-on:focus="error = '';">
        <div class="error" v-if="$v.password.$dirty && !$v.password.required">Field is required</div>
        <div
          class="error"
          v-if="$v.password.$dirty && (!$v.password.minLength || !$v.password.maxLength)"
        >Must be between 8 and 20 characters length</div>

        <button type="submit">Register</button>
        <router-link tag="button" to="/login">Login</router-link>
        <div class="error" v-if="error">{{error}}</div>
      </form>
    </div>
  </Page>
</template>

<style scoped>
form {
  display: inline-block;
  margin: 0 auto;
}
input,
form button {
  display: block;
  margin: 15px 0;
  padding: 5px;
}
form button {
  float: right;
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
import Page from "@/components/Page.vue";
import router from "@/router";
import apollo from "@/apollo";

export default {
  name: "register",
  components: {
    Page
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
  async beforeRouteEnter(to, from, next) {
    try {
      await axios.post("/auth/logout");
      await apollo.resetStore();

      next();
    } catch (error) {
      console.log("%cError logout", "color: orange;", error.message);
    }
  },
  methods: {
    async register(e) {
      e.preventDefault();

      this.$v.$touch();

      if (this.$v.$invalid) {
        return;
      }

      try {
        const { email, password } = this;

        await axios.post("/auth/local/register", { email, password });
        await apollo.resetStore();
        router.push("/");
      } catch (error) {
        this.error = error.response.data;
      }
    }
  }
};
</script>
