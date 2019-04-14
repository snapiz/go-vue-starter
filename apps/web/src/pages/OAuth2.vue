<template>
  <div>Loading...</div>
</template>

<script>
import gql from "graphql-tag";

const providers = {
  google: "GOOGLE",
  facebook: "FACEBOOK"
};

export default {
  created() {
    const { provider, code } = this.$route.params;

    this.$apollo
      .mutate({
        mutation: gql`
          mutation($input: LoginWithOAuth2Input!) {
            loginWithOAuth2(input: $input) {
              user {
                id
                username
                displayName
                picture
              }
            }
          }
        `,
        variables: {
          input: {
            provider: providers[provider],
            code,
            clientMutationId: `pages_OAuth2_${new Date().getTime()}`
          }
        }
      })
      .then(() => {
        this.$router.context.history.push("/avatar");
      })
      .catch(error => {
        console.log(error);
      });
  }
};
</script>
