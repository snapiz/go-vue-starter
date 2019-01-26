<template>
  <Layout v-bind:me="me">
    <div id="me">
      <sui-divider horizontal>{{ "user_me_header_1" | t }}</sui-divider>
      <sui-form id="profile-form" @submit.prevent="update">
        <sui-form-field>
          <label>{{ "user_me_form_label_display_name" | t }}</label>
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
            {{ "form_error_required" | t }}
          </div>
          <div
            class="error"
            v-if="$v.displayName.$dirty && !$v.displayName.alphaNum"
          >
            {{ "form_error_alpha_numeric" | t }}
          </div>
          <div
            class="error"
            v-if="
              $v.displayName.$dirty &&
                (!$v.displayName.minLength || !$v.displayName.maxLength)
            "
          >
            {{ "form_error_between" | t({ min: 3, max: 50 }) }}
          </div>
        </sui-form-field>

        <sui-form-field>
          <label>{{ "user_me_form_label_picture" | t }}</label>
          <sui-input
            v-model.trim="$v.picture.$model"
            icon="image"
            icon-position="left"
            @focus="error = ''"
          />
          <div class="error" v-if="$v.picture.$dirty && !$v.picture.url">
            {{ "form_error_url" | t }}
          </div>
        </sui-form-field>

        <div class="error" v-if="error">{{ error | t }}</div>
        <sui-button type="submit">{{ "user_me_submit" | t }}</sui-button>
      </sui-form>
      <sui-divider horizontal>{{ "user_me_header_2" | t }}</sui-divider>
      <sui-form @submit.prevent="changePassword">
        <sui-form-field>
          <label v-if="me.hasPassword">{{
            "user_me_form_label_current_password" | t
          }}</label>
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
            {{ "form_error_between" | t({ min: 8, max: 20 }) }}
          </div>
        </sui-form-field>

        <sui-form-field>
          <label>{{ "user_me_form_label_new_password" | t }}</label>
          <sui-input
            type="password"
            v-model.trim="$v.password.$model"
            icon="lock"
            icon-position="left"
            @focus="errorPassword = ''"
          />
          <div class="error" v-if="$v.password.$dirty && !$v.password.required">
            {{ "form_error_required" | t }}
          </div>
          <div
            class="error"
            v-if="
              $v.password.$dirty &&
                (!$v.password.minLength || !$v.password.maxLength)
            "
          >
            {{ "form_error_between" | t({ min: 8, max: 20 }) }}
          </div>
        </sui-form-field>

        <div class="error" v-if="errorPassword">{{ errorPassword | t }}</div>
        <sui-button type="submit">{{ "user_me_submit" | t }}</sui-button>
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
  width: 300px;
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
