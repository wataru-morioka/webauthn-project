<template lang='pug'>
  div#project
    BackImage
    SubNav
    div.subject
      p(style='animation-delay: 1s')
        span P
      p(style='animation-delay: 1.1s')
        span R
      p(style='animation-delay: 1.2s')
        span O
      p(style='animation-delay: 1.3s')
        span J
      p(style='animation-delay: 1.4s')
        span E
      p(style='animation-delay: 1.5s')
        span C
      p(style='animation-delay: 1.6s')
        span T
    div#sub-menu-wrap
      SubMenu
    video#project-video(controls playsinline)
      source(src='../assets/jager.mp4' type='video/mp4')
    ProjectContent
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import BackImage from '@/components/BackImage.vue';
import SubNav from '@/components/SubNav.vue';
import SubMenu from '@/components/SubMenu.vue';
import ProjectContent from '@/components/ProjectContent.vue';

@Component({
  components: {
    BackImage,
    SubNav,
    SubMenu,
    ProjectContent,
  },
})
export default class Project extends Vue {}
</script>

<style scoped lang='scss'>
#project {
  position: relative;
  height: 100vw;
}

::-webkit-scrollbar {
  display: none;
  -webkit-appearance: none;
}

#sub-menu-wrap {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  flex-direction: column;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: auto;
  width: 100px;
  height: 270px;
  overflow-y: scroll;
  -webkit-overflow-scrolling: touch;
  z-index: 2;
}

video {
  opacity: 0;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
  width: 40%;
  transition: 1s;
  z-index: -10;
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
  font-size: 60px;
  transform: translateY(0px);
  opacity: .2;
  z-index: -3;

  p {
    animation-name: content-fadein;
    animation-duration: 2s;
    margin-left: 1px;
    margin-right: 1px;
    margin-bottom: 0;
  }
}

@media screen and (max-width: 768px){
  .subject {
    font-size: 30px;
  }

  video {
    width: 100%;
  }

  #sub-menu-wrap {
    display: none;
  }
}

@keyframes content-fadein {
  from {
      opacity: 0;
      transform: translateY(40px);
  }
  to {
      transform: translateY(0);
  }
}
</style>
