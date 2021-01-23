<template lang='pug'>
  div#management
    BackImage
    div.subject
      p(style='animation-delay: 1s')
        span M
      p(style='animation-delay: 1.1s')
        span A
      p(style='animation-delay: 1.2s')
        span N
      p(style='animation-delay: 1.3s')
        span A
      p(style='animation-delay: 1.4s')
        span G
      p(style='animation-delay: 1.5s')
        span E
      p(style='animation-delay: 1.6s')
        span M
      p(style='animation-delay: 1.7s')
        span E
      p(style='animation-delay: 1.8s')
        span N
      p(style='animation-delay: 1.9s')
        span T
    div.nav
      ManagementHeaderNav
    router-view
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import BackImage from '@/components/BackImage.vue';
import FooterNav from '@/components/FooterNav.vue';
import ManagementHeaderNav from '@/components/ManagementHeaderNav.vue';
import SubMenu from '@/components/SubMenu.vue';
import jQuery from 'jQuery';

@Component({
  components: {
    BackImage,
    FooterNav,
    ManagementHeaderNav,
    SubMenu,
  },
})
export default class Management extends Vue {
}
</script>

<style scoped lang='scss'>
#management {
  position: relative;
  animation-name: body-fadein;
  animation-duration: 2s;
}

.subject {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  width: auto;
  height: 500px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  margin: auto;
  font-size: 50px;
  transform: translateY(0px);
  opacity: .2;
  z-index: -3;

   p {
    animation-name: content-fadein;
    animation-duration: 2s;
    margin-bottom: 0;
  }
}
.nav {
  position: fixed;
  display: flex;
  flex-direction: column;
  width: 100%;
  overflow: hidden;
  min-height: 100%;
  animation-name: content-fadein;
  animation-duration: 2s;
  will-change: transform;
}

@media screen and (max-width: 768px){
  .subject {
    font-size: 30px;
  }
}

@keyframes body-fadein {
  from {
      opacity: 0;
  }
  to {
      opacity: 1;
  }
}

@keyframes content-fadein {
  from {
      opacity: 0;
      transform: translateY(40px);
  }
  to {
      // opacity: .2;
      transform: translateY(0);
  }
}
</style>