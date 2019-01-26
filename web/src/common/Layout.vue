<template>
  <div id="layout">
    <div class="navbar">
      <div class="nav-container">
        <sui-menu secondary>
          <sui-menu-item>
            <img src="../assets/logo.png" />
          </sui-menu-item>
          <router-link is="sui-menu-item" link to="/">{{
            "common_layout_nav_home" | t
          }}</router-link>
          <router-link is="sui-menu-item" link to="/about">{{
            "common_layout_nav_about" | t
          }}</router-link>
          <sui-menu-menu position="right">
            <sui-dropdown v-if="me" :icon="null" item>
              <sui-image
                v-if="me.picture"
                v-bind:src="me.picture"
                class="nav-avatar"
                avatar
              />
              <span class="avatar-c cavatar" v-if="!me.picture">
                {{ (me.displayName && me.displayName[0]) || "J" }}
              </span>
              <sui-dropdown-menu>
                <router-link is="sui-menu-item" link to="/me">{{
                  "common_layout_nav_me" | t
                }}</router-link>
                <router-link is="sui-menu-item" link to="/login">{{
                  "common_layout_nav_logout" | t
                }}</router-link>
              </sui-dropdown-menu>
            </sui-dropdown>
            <router-link is="sui-menu-item" link v-if="!me" to="/login">{{
              "common_layout_nav_login" | t
            }}</router-link>
          </sui-menu-menu>
        </sui-menu>
      </div>
    </div>
    <div class="nav-container">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>
.navbar {
  background-color: white;
}
.nav-container {
  max-width: 992px;
  margin: 0 auto;
}
.avatar-c {
  font-size: 1.5em;
  color: white;
  background-color: #dedcdc;
  border-radius: 50%;
  text-align: center;
  padding: 4px 0;
}
.ui.avatar.image,
.avatar-c {
  width: 32px;
  height: 32px;
}
</style>

<script>
import RouterLink from "./RouterLink.vue";

export default {
  name: "Layout",
  props: {
    me: Object
  },
  components: {
    RouterLink
  }
};
</script>
