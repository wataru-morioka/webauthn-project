<template lang='pug'>
  div#nav
    HeaderNav
    CenterNav
    FooterNav
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import HeaderNav from '@/components/HeaderNav.vue';
import CenterNav from '@/components/CenterNav.vue';
import FooterNav from '@/components/FooterNav.vue';

@Component({
  components: {
    HeaderNav,
    CenterNav,
    FooterNav,
  },
})
export default class MainNav extends Vue {}
</script>

<style scoped lang='scss'>
#nav {
  position: relative;
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100vh;
  min-height: 100%;
  animation-name: fadein;
  animation-duration: 2s;
  will-change: transform;
}

@keyframes fadein {
  from {
      opacity: 0;
      transform: translateY(20px);
  }
  to {
      opacity: 1;
      transform: translateY(0);
  }
}
</style>
