<template>
  <Layout>
    <div class="me">
      <form @submit.prevent="update">
        <label>Display name</label>
        <input type="text" v-model.trim="$v.displayName.$model" @focus="error = '';">
        <div
          class="error"
          v-if="$v.displayName.$dirty && !$v.displayName.required"
        >Field is required</div>
        <div
          class="error"
          v-if="$v.displayName.$dirty && !$v.displayName.alphaNum"
        >Must be alpha numeric</div>
        <div
          class="error"
          v-if="$v.displayName.$dirty && (!$v.displayName.minLength || !$v.displayName.maxLength)"
        >Must be between 3 and 50 characters length</div>

        <label>Picture</label>
        <input type="text" v-model.trim="$v.picture.$model" @focus="error = '';">
        <div class="error" v-if="$v.picture.$dirty && !$v.picture.url">Must be valid URL</div>

        <div class="error" v-if="error">{{error}}</div>
        <button type="submit">Update</button>
      </form>
      <div>
        <h3>Change password</h3>
      </div>
      <form @submit.prevent="changePassword">
        <label v-if="me.hasPassword">Current password</label>
        <input
          v-if="me.hasPassword"
          type="password"
          v-model.trim="$v.currentPassword.$model"
          @focus="errorPassword = '';"
        >
        <div
          class="error"
          v-if="$v.currentPassword.$dirty && (!$v.currentPassword.minLength || !$v.currentPassword.maxLength)"
        >Must be between 8 and 20 characters length</div>

        <label>New password</label>
        <input type="password" v-model.trim="$v.password.$model" @focus="errorPassword = '';">
        <div class="error" v-if="$v.password.$dirty && !$v.password.required">Field is required</div>
        <div
          class="error"
          v-if="$v.password.$dirty && (!$v.password.minLength || !$v.password.maxLength)"
        >Must be between 8 and 20 characters length</div>

        <div class="error" v-if="errorPassword">{{errorPassword}}</div>
        <button type="submit">Change password</button>
      </form>
    </div>
  </Layout>
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
import { required, minLength, maxLength, url } from "vuelidate/lib/validators";
import gql from "graphql-tag";
import { pick } from "lodash";

import Layout from "@/components/Layout.vue";
import { alphaNum } from "@/validators";
import { getGraphQLError } from "@/utils";
import apollo from "@/apollo";

export default {
  name: "me",
  components: {
    Layout
  },
  data() {
    return {
      ...pick(this.$route.params.$data.me, "displayName", "picture"),
      password: "",
      currentPassword: "",
      error: "",
      errorPassword: ""
    };
  },
  validations: {
    displayName: {
      required,
      alphaNum,
      minLength: minLength(3),
      maxLength: maxLength(50)
    },
    picture: {
      url
    },
    password: {
      required,
      minLength: minLength(8),
      maxLength: maxLength(20)
    },
    currentPassword: {
      minLength: minLength(8),
      maxLength: maxLength(20)
    },
    updateForm: ["displayName", "picture"],
    passwordForm: ["password", "currentPassword"]
  },
  methods: {
    async update() {
      this.$v.updateForm.$touch();

      if (this.$v.updateForm.$invalid) {
        return;
      }

      try {
        let { displayName, picture } = this;

        if (!picture) {
          picture = "";
        }

        await apollo.mutate({
          mutation: gql`
            mutation($input: UpdateUserInput!) {
              updateUser(input: $input) {
                user {
                  id
                  displayName
                  picture
                }
              }
            }
          `,
          variables: {
            input: {
              displayName,
              picture,
              clientMutationId: `Me_${new Date().getTime()}`
            }
          }
        });
      } catch (error) {
        this.error = getGraphQLError(error);
      }
    },
    async changePassword() {
      this.$v.passwordForm.$touch();

      if (this.$v.passwordForm.$invalid) {
        return;
      }

      try {
        let { password, currentPassword } = this;

        if (!currentPassword) {
          currentPassword = "";
        }

        await apollo.mutate({
          mutation: gql`
            mutation($input: ChangePasswordInput!) {
              changePassword(input: $input) {
                user {
                  id
                  hasPassword
                }
              }
            }
          `,
          variables: {
            input: {
              password,
              currentPassword,
              clientMutationId: `Me_${new Date().getTime()}`
            }
          }
        });

        this.me.hasPassword = true;
        this.password = "";
        this.currentPassword = "";
        this.$v.$reset();
      } catch (error) {
        this.errorPassword = getGraphQLError(error);
      }
    }
  }
};
</script>
