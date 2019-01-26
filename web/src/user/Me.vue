<template>
  <Layout v-bind:me="me">
    <div id="me">
      <sui-divider horizontal>Profile</sui-divider>
      <sui-form id="profile-form" @submit.prevent="update">
        <sui-form-field>
          <label>Display name</label>
          <sui-input
            v-model.trim="$v.displayName.$model"
            icon="user"
            icon-position="left"
            @focus="error = ''"
          />
          <div
            class="error"
            v-if="$v.displayName.$dirty && !$v.displayName.required"
          >
            Field is required
          </div>
          <div
            class="error"
            v-if="$v.displayName.$dirty && !$v.displayName.alphaNum"
          >
            Must be alpha numeric
          </div>
          <div
            class="error"
            v-if="
              $v.displayName.$dirty &&
                (!$v.displayName.minLength || !$v.displayName.maxLength)
            "
          >
            Must be between 3 and 50 characters length
          </div>
        </sui-form-field>

        <sui-form-field>
          <label>Picture</label>
          <sui-input
            v-model.trim="$v.picture.$model"
            icon="image"
            icon-position="left"
            @focus="error = ''"
          />
          <div class="error" v-if="$v.picture.$dirty && !$v.picture.url">
            Must be valid URL
          </div>
        </sui-form-field>

        <div class="error" v-if="error">{{ error }}</div>
        <sui-button type="submit">Update</sui-button>
      </sui-form>
      <sui-divider horizontal>Change password</sui-divider>
      <sui-form @submit.prevent="changePassword">
        <sui-form-field>
          <label v-if="me.hasPassword">Current password</label>
          <sui-input
            v-if="me.hasPassword"
            type="password"
            v-model.trim="$v.currentPassword.$model"
            icon="lock"
            icon-position="left"
            @focus="errorPassword = ''"
          />
          <div
            class="error"
            v-if="
              $v.currentPassword.$dirty &&
                (!$v.currentPassword.minLength || !$v.currentPassword.maxLength)
            "
          >
            Must be between 8 and 20 characters length
          </div>
        </sui-form-field>

        <sui-form-field>
          <label>New password</label>
          <sui-input
            type="password"
            v-model.trim="$v.password.$model"
            icon="lock"
            icon-position="left"
            @focus="errorPassword = ''"
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

        <div class="error" v-if="errorPassword">{{ errorPassword }}</div>
        <sui-button type="submit">Change password</sui-button>
      </sui-form>
    </div>
  </Layout>
</template>

<style scoped>
#me {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: calc(100vh - 60px);
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
#profile-form {
  margin-bottom: 15px;
}
.error {
  color: red;
}
</style>

<script>
import { required, minLength, maxLength, url } from "vuelidate/lib/validators";
import gql from "graphql-tag";
import { pick } from "lodash";

import Layout from "@/common/Layout.vue";
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
      ...pick(this.me, "displayName", "picture"),
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
