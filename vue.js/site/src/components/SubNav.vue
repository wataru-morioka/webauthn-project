<template lang='pug'>
  div#nav
    HeaderNav
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import HeaderNav from '@/components/HeaderNav.vue';

@Component({
  components: {
    HeaderNav,
  },
})
export default class SubNav extends Vue {}
</script>

<style scoped lang='scss'>
#nav {
  position: fixed;
  display: flex;
  flex-direction: column;
  width: 100%;
  overflow: hidden;
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
