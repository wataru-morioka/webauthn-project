<template lang='pug'>
  div#app
    router-view
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import jQuert from 'jQuery';

@Component
export default class App extends Vue {
  private beforeCreate() {
    this.$store.dispatch('checkLoginStatus');
  }
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #ffffff70;
}
</style>
