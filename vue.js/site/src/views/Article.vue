<template lang='pug'>
  div#about
    BackImage
    SubNav
    div.subject
      p(style='animation-delay: 1s')
        span A
      p(style='animation-delay: 1.1s')
        span R
      p(style='animation-delay: 1.2s')
        span T
      p(style='animation-delay: 1.3s')
        span I
      p(style='animation-delay: 1.4s')
        span C
      p(style='animation-delay: 1.5s')
        span L
      p(style='animation-delay: 1.6s')
        span E
    ArticleContent
</template>

<script lang='ts'>
import { Component, Vue } from 'vue-property-decorator';
import BackImage from '@/components/BackImage.vue';
import SubNav from '@/components/SubNav.vue';
import ArticleContent from '@/components/ArticleContent.vue';

@Component({
  components: {
    BackImage,
    SubNav,
    ArticleContent,
  },
})
export default class Article extends Vue {}
</script>

<style scoped lang='scss'>
#about {
  position: relative;
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
